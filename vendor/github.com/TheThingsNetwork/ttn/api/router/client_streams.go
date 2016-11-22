// Copyright © 2016 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package router

import (
	"io"
	"time"

	"github.com/TheThingsNetwork/ttn/api"
	"github.com/TheThingsNetwork/ttn/api/gateway"
	"github.com/TheThingsNetwork/ttn/utils/backoff"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// GatewayStream interface
type GatewayStream interface {
	Close()
}

type gatewayStream struct {
	closing bool
	ctx     api.Logger
	client  RouterClientForGateway
}

// DefaultBufferSize indicates the default send and receive buffer sizes
var DefaultBufferSize = 10

// GatewayStatusStream for sending gateway statuses
type GatewayStatusStream interface {
	GatewayStream
	Send(*gateway.Status) error
}

// NewMonitoredGatewayStatusStream starts and monitors a GatewayStatusStream
func NewMonitoredGatewayStatusStream(client RouterClientForGateway) GatewayStatusStream {
	s := &gatewayStatusStream{
		ch:  make(chan *gateway.Status, DefaultBufferSize),
		err: make(chan error),
	}
	s.client = client
	s.ctx = client.GetLogger()

	go func() {
		var retries int

		for {
			// Session channels
			ch := make(chan *gateway.Status)
			errCh := make(chan error)

			// Session client
			client, err := s.client.GatewayStatus()
			if err != nil {
				s.ctx.WithError(err).Warn("Could not start GatewayStatus stream, retrying...")
				time.Sleep(backoff.Backoff(retries))
				retries++
				continue
			}
			retries = 0

			s.ctx.Debug("Started GatewayStatus stream")

			// Receive errors
			go func() {
				empty := new(empty.Empty)
				if err := client.RecvMsg(empty); err != nil {
					errCh <- err
				}
				close(errCh)
			}()

			// Send
			go func() {
				for status := range ch {
					s.ctx.Debug("Sending GatewayStatus message")
					if err := client.Send(status); err != nil {
						s.ctx.WithError(err).Warn("Error sending GatewayStatus message")
						break
					}
				}
			}()

			// Monitoring
			var mErr error

		monitor:
			for {
				select {
				case mErr = <-errCh:
					break monitor
				case msg, ok := <-s.ch:
					if !ok {
						break monitor // channel closed
					}
					ch <- msg
				}
			}

			close(ch)
			_, mErr = client.CloseAndRecv()

			if mErr == nil || mErr == io.EOF || grpc.Code(mErr) == codes.Canceled {
				s.ctx.Debug("Stopped GatewayStatus stream")
				if s.closing {
					break
				}
			} else {
				s.ctx.WithError(mErr).Warn("Error in GatewayStatus stream")
			}
			time.Sleep(backoff.Backoff(retries))
			retries++
		}
	}()

	return s
}

type gatewayStatusStream struct {
	gatewayStream
	ch  chan *gateway.Status
	err chan error
}

func (s *gatewayStatusStream) Send(status *gateway.Status) error {
	select {
	case s.ch <- status:
	default:
		s.ctx.Warn("Dropping GatewayStatus message, buffer full")
	}
	return nil
}

func (s *gatewayStatusStream) Close() {
	s.closing = true
	close(s.ch)
}

// UplinkStream for sending uplink messages
type UplinkStream interface {
	GatewayStream
	Send(*UplinkMessage) error
}

// NewMonitoredUplinkStream starts and monitors a UplinkStream
func NewMonitoredUplinkStream(client RouterClientForGateway) UplinkStream {
	s := &uplinkStream{
		ch:  make(chan *UplinkMessage, DefaultBufferSize),
		err: make(chan error),
	}
	s.client = client
	s.ctx = client.GetLogger()

	go func() {
		var retries int

		for {
			// Session channels
			ch := make(chan *UplinkMessage)
			errCh := make(chan error)

			// Session client
			client, err := s.client.Uplink()
			if err != nil {
				s.ctx.WithError(err).Warn("Could not start Uplink stream, retrying...")
				time.Sleep(backoff.Backoff(retries))
				retries++
				continue
			}
			retries = 0

			s.ctx.Debug("Started Uplink stream")

			// Receive errors
			go func() {
				empty := new(empty.Empty)
				if err := client.RecvMsg(empty); err != nil {
					errCh <- err
				}
				close(errCh)
			}()

			// Send
			go func() {
				for message := range ch {
					s.ctx.Debug("Sending Uplink message")
					if err := client.Send(message); err != nil {
						s.ctx.WithError(err).Warn("Error sending Uplink message")
						break
					}
				}
			}()

			// Monitoring
			var mErr error

		monitor:
			for {
				select {
				case mErr = <-errCh:
					break monitor
				case msg, ok := <-s.ch:
					if !ok {
						break monitor // channel closed
					}
					ch <- msg
				}
			}

			close(ch)
			_, mErr = client.CloseAndRecv()

			if mErr == nil || mErr == io.EOF || grpc.Code(mErr) == codes.Canceled {
				s.ctx.Debug("Stopped Uplink stream")
				if s.closing {
					break
				}
			} else {
				s.ctx.WithError(mErr).Warn("Error in Uplink stream")
			}
			time.Sleep(backoff.Backoff(retries))
			retries++
		}
	}()

	return s
}

type uplinkStream struct {
	gatewayStream
	ch  chan *UplinkMessage
	err chan error
}

func (s *uplinkStream) Send(message *UplinkMessage) error {
	select {
	case s.ch <- message:
	default:
		s.ctx.Warn("Dropping Uplink message, buffer full")
	}
	return nil
}

func (s *uplinkStream) Close() {
	s.closing = true
	close(s.ch)
}

// DownlinkStream for sending downlink messages
type DownlinkStream interface {
	GatewayStream
	Channel() <-chan *DownlinkMessage
}

// NewMonitoredDownlinkStream starts and monitors a DownlinkStream
func NewMonitoredDownlinkStream(client RouterClientForGateway) DownlinkStream {
	s := &downlinkStream{
		ch:  make(chan *DownlinkMessage, DefaultBufferSize),
		err: make(chan error),
	}
	s.client = client
	s.ctx = client.GetLogger()

	go func() {
		var client Router_SubscribeClient
		var err error
		var retries int
		var message *DownlinkMessage

		for {
			client, s.cancel, err = s.client.Subscribe()
			if err != nil {
				s.ctx.WithError(err).Warn("Could not start Downlink stream, retrying...")
				time.Sleep(backoff.Backoff(retries))
				retries++
				continue
			}
			retries = 0

			s.ctx.Debug("Started Downlink stream")

			for {
				message, err = client.Recv()
				if message != nil {
					s.ctx.Debug("Receiving Downlink message")
					select {
					case s.ch <- message:
					default:
						s.ctx.Warn("Dropping Downlink message, buffer full")
					}
				}
				if err != nil {
					break
				}
			}

			if err == nil || err == io.EOF || grpc.Code(err) == codes.Canceled {
				s.ctx.Debug("Stopped Downlink stream")
				if s.closing {
					break
				}
			} else {
				s.ctx.WithError(err).Warn("Error in Downlink stream")
			}
			time.Sleep(backoff.Backoff(retries))
			retries++
		}

		close(s.ch)
	}()
	return s
}

type downlinkStream struct {
	gatewayStream
	cancel context.CancelFunc
	ch     chan *DownlinkMessage
	err    chan error
}

func (s *downlinkStream) Close() {
	s.closing = true
	if s.cancel != nil {
		s.cancel()
	}
}

func (s *downlinkStream) Channel() <-chan *DownlinkMessage {
	return s.ch
}
