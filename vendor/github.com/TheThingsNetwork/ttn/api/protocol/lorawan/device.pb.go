// Code generated by protoc-gen-gogo.
// source: github.com/TheThingsNetwork/ttn/api/protocol/lorawan/device.proto
// DO NOT EDIT!

/*
	Package lorawan is a generated protocol buffer package.

	It is generated from these files:
		github.com/TheThingsNetwork/ttn/api/protocol/lorawan/device.proto

	It has these top-level messages:
		DeviceIdentifier
		Device
*/
package lorawan

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_TheThingsNetwork_ttn_core_types "github.com/TheThingsNetwork/ttn/core/types"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DeviceIdentifier struct {
	// The AppEUI is a unique, 8 byte identifier for the application a device belongs to.
	AppEui *github_com_TheThingsNetwork_ttn_core_types.AppEUI `protobuf:"bytes,1,opt,name=app_eui,json=appEui,proto3,customtype=github.com/TheThingsNetwork/ttn/core/types.AppEUI" json:"app_eui,omitempty"`
	// The DevEUI is a unique, 8 byte identifier for the device.
	DevEui *github_com_TheThingsNetwork_ttn_core_types.DevEUI `protobuf:"bytes,2,opt,name=dev_eui,json=devEui,proto3,customtype=github.com/TheThingsNetwork/ttn/core/types.DevEUI" json:"dev_eui,omitempty"`
}

func (m *DeviceIdentifier) Reset()                    { *m = DeviceIdentifier{} }
func (m *DeviceIdentifier) String() string            { return proto.CompactTextString(m) }
func (*DeviceIdentifier) ProtoMessage()               {}
func (*DeviceIdentifier) Descriptor() ([]byte, []int) { return fileDescriptorDevice, []int{0} }

type Device struct {
	// The AppEUI is a unique, 8 byte identifier for the application a device belongs to.
	AppEui *github_com_TheThingsNetwork_ttn_core_types.AppEUI `protobuf:"bytes,1,opt,name=app_eui,json=appEui,proto3,customtype=github.com/TheThingsNetwork/ttn/core/types.AppEUI" json:"app_eui,omitempty"`
	// The DevEUI is a unique, 8 byte identifier for the device.
	DevEui *github_com_TheThingsNetwork_ttn_core_types.DevEUI `protobuf:"bytes,2,opt,name=dev_eui,json=devEui,proto3,customtype=github.com/TheThingsNetwork/ttn/core/types.DevEUI" json:"dev_eui,omitempty"`
	// The AppID is a unique identifier for the application a device belongs to. It can contain lowercase letters, numbers, - and _.
	AppId string `protobuf:"bytes,3,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	// The DevID is a unique identifier for the device. It can contain lowercase letters, numbers, - and _.
	DevId string `protobuf:"bytes,4,opt,name=dev_id,json=devId,proto3" json:"dev_id,omitempty"`
	// The DevAddr is a dynamic, 4 byte session address for the device.
	DevAddr *github_com_TheThingsNetwork_ttn_core_types.DevAddr `protobuf:"bytes,5,opt,name=dev_addr,json=devAddr,proto3,customtype=github.com/TheThingsNetwork/ttn/core/types.DevAddr" json:"dev_addr,omitempty"`
	// The NwkSKey is a 16 byte session key that is known by the device and the network. It is used for routing and MAC related functionality.
	// This key is negotiated during the OTAA join procedure, or statically configured using ABP.
	NwkSKey *github_com_TheThingsNetwork_ttn_core_types.NwkSKey `protobuf:"bytes,6,opt,name=nwk_s_key,json=nwkSKey,proto3,customtype=github.com/TheThingsNetwork/ttn/core/types.NwkSKey" json:"nwk_s_key,omitempty"`
	// The AppSKey is a 16 byte session key that is known by the device and the application. It is used for payload encryption.
	// This key is negotiated during the OTAA join procedure, or statically configured using ABP.
	AppSKey *github_com_TheThingsNetwork_ttn_core_types.AppSKey `protobuf:"bytes,7,opt,name=app_s_key,json=appSKey,proto3,customtype=github.com/TheThingsNetwork/ttn/core/types.AppSKey" json:"app_s_key,omitempty"`
	// The AppKey is a 16 byte static key that is known by the device and the application. It is used for negotiating session keys (OTAA).
	AppKey *github_com_TheThingsNetwork_ttn_core_types.AppKey `protobuf:"bytes,8,opt,name=app_key,json=appKey,proto3,customtype=github.com/TheThingsNetwork/ttn/core/types.AppKey" json:"app_key,omitempty"`
	// FCntUp is the uplink frame counter for a device session.
	FCntUp uint32 `protobuf:"varint,9,opt,name=f_cnt_up,json=fCntUp,proto3" json:"f_cnt_up,omitempty"`
	// FCntDown is the downlink frame counter for a device session.
	FCntDown uint32 `protobuf:"varint,10,opt,name=f_cnt_down,json=fCntDown,proto3" json:"f_cnt_down,omitempty"`
	// The DisableFCntCheck option disables the frame counter check. Disabling this makes the device vulnerable to replay attacks, but makes ABP slightly easier.
	DisableFCntCheck bool `protobuf:"varint,11,opt,name=disable_f_cnt_check,json=disableFCntCheck,proto3" json:"disable_f_cnt_check,omitempty"`
	// The Uses32BitFCnt option indicates that the device keeps track of full 32 bit frame counters. As only the 16 lsb are actually transmitted, the 16 msb will have to be inferred.
	Uses32BitFCnt bool `protobuf:"varint,12,opt,name=uses32_bit_f_cnt,json=uses32BitFCnt,proto3" json:"uses32_bit_f_cnt,omitempty"`
	// The ActivationContstraints are used to allocate a device address for a device.
	// There are different prefixes for `otaa`, `abp`, `world`, `local`, `private`, `testing`.
	ActivationConstraints string `protobuf:"bytes,13,opt,name=activation_constraints,json=activationConstraints,proto3" json:"activation_constraints,omitempty"`
	// When the device was last seen (Unix nanoseconds)
	LastSeen int64 `protobuf:"varint,21,opt,name=last_seen,json=lastSeen,proto3" json:"last_seen,omitempty"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptorDevice, []int{1} }

func init() {
	proto.RegisterType((*DeviceIdentifier)(nil), "lorawan.DeviceIdentifier")
	proto.RegisterType((*Device)(nil), "lorawan.Device")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DeviceManager service

type DeviceManagerClient interface {
	GetDevice(ctx context.Context, in *DeviceIdentifier, opts ...grpc.CallOption) (*Device, error)
	SetDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	DeleteDevice(ctx context.Context, in *DeviceIdentifier, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type deviceManagerClient struct {
	cc *grpc.ClientConn
}

func NewDeviceManagerClient(cc *grpc.ClientConn) DeviceManagerClient {
	return &deviceManagerClient{cc}
}

func (c *deviceManagerClient) GetDevice(ctx context.Context, in *DeviceIdentifier, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := grpc.Invoke(ctx, "/lorawan.DeviceManager/GetDevice", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceManagerClient) SetDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/lorawan.DeviceManager/SetDevice", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceManagerClient) DeleteDevice(ctx context.Context, in *DeviceIdentifier, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/lorawan.DeviceManager/DeleteDevice", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DeviceManager service

type DeviceManagerServer interface {
	GetDevice(context.Context, *DeviceIdentifier) (*Device, error)
	SetDevice(context.Context, *Device) (*google_protobuf.Empty, error)
	DeleteDevice(context.Context, *DeviceIdentifier) (*google_protobuf.Empty, error)
}

func RegisterDeviceManagerServer(s *grpc.Server, srv DeviceManagerServer) {
	s.RegisterService(&_DeviceManager_serviceDesc, srv)
}

func _DeviceManager_GetDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceManagerServer).GetDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lorawan.DeviceManager/GetDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceManagerServer).GetDevice(ctx, req.(*DeviceIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceManager_SetDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Device)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceManagerServer).SetDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lorawan.DeviceManager/SetDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceManagerServer).SetDevice(ctx, req.(*Device))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceManager_DeleteDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceManagerServer).DeleteDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lorawan.DeviceManager/DeleteDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceManagerServer).DeleteDevice(ctx, req.(*DeviceIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeviceManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lorawan.DeviceManager",
	HandlerType: (*DeviceManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDevice",
			Handler:    _DeviceManager_GetDevice_Handler,
		},
		{
			MethodName: "SetDevice",
			Handler:    _DeviceManager_SetDevice_Handler,
		},
		{
			MethodName: "DeleteDevice",
			Handler:    _DeviceManager_DeleteDevice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/TheThingsNetwork/ttn/api/protocol/lorawan/device.proto",
}

func (m *DeviceIdentifier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeviceIdentifier) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.AppEui != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDevice(dAtA, i, uint64(m.AppEui.Size()))
		n1, err := m.AppEui.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.DevEui != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDevice(dAtA, i, uint64(m.DevEui.Size()))
		n2, err := m.DevEui.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *Device) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Device) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.AppEui != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDevice(dAtA, i, uint64(m.AppEui.Size()))
		n3, err := m.AppEui.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.DevEui != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDevice(dAtA, i, uint64(m.DevEui.Size()))
		n4, err := m.DevEui.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if len(m.AppId) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintDevice(dAtA, i, uint64(len(m.AppId)))
		i += copy(dAtA[i:], m.AppId)
	}
	if len(m.DevId) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintDevice(dAtA, i, uint64(len(m.DevId)))
		i += copy(dAtA[i:], m.DevId)
	}
	if m.DevAddr != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintDevice(dAtA, i, uint64(m.DevAddr.Size()))
		n5, err := m.DevAddr.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.NwkSKey != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintDevice(dAtA, i, uint64(m.NwkSKey.Size()))
		n6, err := m.NwkSKey.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.AppSKey != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintDevice(dAtA, i, uint64(m.AppSKey.Size()))
		n7, err := m.AppSKey.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.AppKey != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintDevice(dAtA, i, uint64(m.AppKey.Size()))
		n8, err := m.AppKey.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	if m.FCntUp != 0 {
		dAtA[i] = 0x48
		i++
		i = encodeVarintDevice(dAtA, i, uint64(m.FCntUp))
	}
	if m.FCntDown != 0 {
		dAtA[i] = 0x50
		i++
		i = encodeVarintDevice(dAtA, i, uint64(m.FCntDown))
	}
	if m.DisableFCntCheck {
		dAtA[i] = 0x58
		i++
		if m.DisableFCntCheck {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Uses32BitFCnt {
		dAtA[i] = 0x60
		i++
		if m.Uses32BitFCnt {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.ActivationConstraints) > 0 {
		dAtA[i] = 0x6a
		i++
		i = encodeVarintDevice(dAtA, i, uint64(len(m.ActivationConstraints)))
		i += copy(dAtA[i:], m.ActivationConstraints)
	}
	if m.LastSeen != 0 {
		dAtA[i] = 0xa8
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintDevice(dAtA, i, uint64(m.LastSeen))
	}
	return i, nil
}

func encodeFixed64Device(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Device(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintDevice(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DeviceIdentifier) Size() (n int) {
	var l int
	_ = l
	if m.AppEui != nil {
		l = m.AppEui.Size()
		n += 1 + l + sovDevice(uint64(l))
	}
	if m.DevEui != nil {
		l = m.DevEui.Size()
		n += 1 + l + sovDevice(uint64(l))
	}
	return n
}

func (m *Device) Size() (n int) {
	var l int
	_ = l
	if m.AppEui != nil {
		l = m.AppEui.Size()
		n += 1 + l + sovDevice(uint64(l))
	}
	if m.DevEui != nil {
		l = m.DevEui.Size()
		n += 1 + l + sovDevice(uint64(l))
	}
	l = len(m.AppId)
	if l > 0 {
		n += 1 + l + sovDevice(uint64(l))
	}
	l = len(m.DevId)
	if l > 0 {
		n += 1 + l + sovDevice(uint64(l))
	}
	if m.DevAddr != nil {
		l = m.DevAddr.Size()
		n += 1 + l + sovDevice(uint64(l))
	}
	if m.NwkSKey != nil {
		l = m.NwkSKey.Size()
		n += 1 + l + sovDevice(uint64(l))
	}
	if m.AppSKey != nil {
		l = m.AppSKey.Size()
		n += 1 + l + sovDevice(uint64(l))
	}
	if m.AppKey != nil {
		l = m.AppKey.Size()
		n += 1 + l + sovDevice(uint64(l))
	}
	if m.FCntUp != 0 {
		n += 1 + sovDevice(uint64(m.FCntUp))
	}
	if m.FCntDown != 0 {
		n += 1 + sovDevice(uint64(m.FCntDown))
	}
	if m.DisableFCntCheck {
		n += 2
	}
	if m.Uses32BitFCnt {
		n += 2
	}
	l = len(m.ActivationConstraints)
	if l > 0 {
		n += 1 + l + sovDevice(uint64(l))
	}
	if m.LastSeen != 0 {
		n += 2 + sovDevice(uint64(m.LastSeen))
	}
	return n
}

func sovDevice(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDevice(x uint64) (n int) {
	return sovDevice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DeviceIdentifier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDevice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeviceIdentifier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceIdentifier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppEui", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_TheThingsNetwork_ttn_core_types.AppEUI
			m.AppEui = &v
			if err := m.AppEui.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevEui", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_TheThingsNetwork_ttn_core_types.DevEUI
			m.DevEui = &v
			if err := m.DevEui.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDevice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDevice
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Device) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDevice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Device: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Device: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppEui", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_TheThingsNetwork_ttn_core_types.AppEUI
			m.AppEui = &v
			if err := m.AppEui.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevEui", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_TheThingsNetwork_ttn_core_types.DevEUI
			m.DevEui = &v
			if err := m.DevEui.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DevId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_TheThingsNetwork_ttn_core_types.DevAddr
			m.DevAddr = &v
			if err := m.DevAddr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NwkSKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_TheThingsNetwork_ttn_core_types.NwkSKey
			m.NwkSKey = &v
			if err := m.NwkSKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppSKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_TheThingsNetwork_ttn_core_types.AppSKey
			m.AppSKey = &v
			if err := m.AppSKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_TheThingsNetwork_ttn_core_types.AppKey
			m.AppKey = &v
			if err := m.AppKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FCntUp", wireType)
			}
			m.FCntUp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FCntUp |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FCntDown", wireType)
			}
			m.FCntDown = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FCntDown |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisableFCntCheck", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DisableFCntCheck = bool(v != 0)
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uses32BitFCnt", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Uses32BitFCnt = bool(v != 0)
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActivationConstraints", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDevice
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ActivationConstraints = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 21:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastSeen", wireType)
			}
			m.LastSeen = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastSeen |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDevice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDevice
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDevice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDevice
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDevice
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthDevice
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDevice
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipDevice(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthDevice = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDevice   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/TheThingsNetwork/ttn/api/protocol/lorawan/device.proto", fileDescriptorDevice)
}

var fileDescriptorDevice = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x94, 0xcf, 0x6e, 0x13, 0x3d,
	0x14, 0xc5, 0xbf, 0xf9, 0x4a, 0x93, 0x8c, 0x69, 0x44, 0x65, 0xd4, 0xca, 0xa4, 0x28, 0x8d, 0xba,
	0x21, 0x9b, 0xce, 0x88, 0xfe, 0x81, 0x75, 0x9a, 0xa4, 0x28, 0x42, 0x54, 0x62, 0xda, 0x6e, 0xd8,
	0x8c, 0x9c, 0xf1, 0xcd, 0xc4, 0xca, 0xd4, 0xb6, 0x66, 0x3c, 0x33, 0xca, 0x03, 0xf0, 0x0e, 0x3c,
	0x07, 0x6f, 0xc0, 0x8e, 0x25, 0xeb, 0x2e, 0x2a, 0x54, 0x5e, 0x04, 0x79, 0x9c, 0x52, 0x14, 0x09,
	0x55, 0x64, 0xc5, 0xee, 0xce, 0x39, 0xc7, 0xbf, 0x6b, 0xc7, 0xf1, 0x45, 0xbd, 0x98, 0xeb, 0x69,
	0x3e, 0xf6, 0x22, 0x79, 0xe5, 0x5f, 0x4c, 0xe1, 0x62, 0xca, 0x45, 0x9c, 0x9d, 0x81, 0x2e, 0x65,
	0x3a, 0xf3, 0xb5, 0x16, 0x3e, 0x55, 0xdc, 0x57, 0xa9, 0xd4, 0x32, 0x92, 0x89, 0x9f, 0xc8, 0x94,
	0x96, 0x54, 0xf8, 0x0c, 0x0a, 0x1e, 0x81, 0x57, 0xe9, 0xb8, 0xbe, 0x50, 0x5b, 0x3b, 0xb1, 0x94,
	0x71, 0x02, 0x36, 0x3e, 0xce, 0x27, 0x3e, 0x5c, 0x29, 0x3d, 0xb7, 0xa9, 0xd6, 0xfe, 0x6f, 0x8d,
	0x62, 0x19, 0xcb, 0xfb, 0x94, 0xf9, 0xaa, 0x3e, 0xaa, 0xca, 0xc6, 0xf7, 0x3e, 0x3b, 0x68, 0x73,
	0x50, 0x75, 0x19, 0x31, 0x10, 0x9a, 0x4f, 0x38, 0xa4, 0xf8, 0x0c, 0xd5, 0xa9, 0x52, 0x21, 0xe4,
	0x9c, 0x38, 0x1d, 0xa7, 0xbb, 0x71, 0x72, 0x7c, 0x7d, 0xb3, 0xfb, 0xf2, 0xa1, 0x13, 0x44, 0x32,
	0x05, 0x5f, 0xcf, 0x15, 0x64, 0x5e, 0x4f, 0xa9, 0xe1, 0xe5, 0x28, 0xa8, 0x51, 0xa5, 0x86, 0x39,
	0x37, 0x3c, 0x06, 0x45, 0xc5, 0xfb, 0x7f, 0x25, 0xde, 0x00, 0x8a, 0x8a, 0xc7, 0xa0, 0x18, 0xe6,
	0x7c, 0xef, 0x63, 0x0d, 0xd5, 0xec, 0xa6, 0xff, 0xf5, 0xad, 0xe2, 0x2d, 0x64, 0xc8, 0x21, 0x67,
	0x64, 0xad, 0xe3, 0x74, 0xdd, 0x60, 0x9d, 0x2a, 0x35, 0x62, 0x46, 0x36, 0x6d, 0x38, 0x23, 0x8f,
	0xac, 0xcc, 0xa0, 0x18, 0x31, 0xfc, 0x1e, 0x35, 0x8c, 0x4c, 0x19, 0x4b, 0xc9, 0x7a, 0xd5, 0xfe,
	0xd5, 0xf5, 0xcd, 0xee, 0xc1, 0xdf, 0xb5, 0xef, 0x31, 0x96, 0x06, 0xe6, 0x14, 0xa6, 0xc0, 0x01,
	0x72, 0x45, 0x39, 0x0b, 0xb3, 0x70, 0x06, 0x73, 0x52, 0x5b, 0x89, 0x79, 0x56, 0xce, 0xce, 0xdf,
	0xc2, 0x3c, 0xa8, 0x0b, 0x5b, 0x18, 0xa6, 0x39, 0x94, 0x65, 0xd6, 0x57, 0x62, 0xf6, 0x94, 0xb2,
	0x4c, 0x6a, 0x8b, 0xbb, 0x8b, 0x34, 0xc4, 0xc6, 0xaa, 0x17, 0x69, 0x80, 0xe6, 0xe7, 0x36, 0x3c,
	0x82, 0x1a, 0x93, 0x30, 0x12, 0x3a, 0xcc, 0x15, 0x71, 0x3b, 0x4e, 0xb7, 0x19, 0xd4, 0x26, 0x7d,
	0xa1, 0x2f, 0x15, 0x7e, 0x8e, 0x90, 0x75, 0x98, 0x2c, 0x05, 0x41, 0x95, 0xd7, 0x30, 0xde, 0x40,
	0x96, 0x02, 0xef, 0xa3, 0xa7, 0x8c, 0x67, 0x74, 0x9c, 0x40, 0x68, 0x53, 0xd1, 0x14, 0xa2, 0x19,
	0x79, 0xdc, 0x71, 0xba, 0x8d, 0x60, 0x73, 0x61, 0x9d, 0xf6, 0x85, 0xee, 0x1b, 0x1d, 0xbf, 0x40,
	0x9b, 0x79, 0x06, 0xd9, 0xe1, 0x41, 0x38, 0xe6, 0xda, 0xae, 0x20, 0x1b, 0x55, 0xb6, 0x69, 0xf5,
	0x13, 0xae, 0x4d, 0x1a, 0x1f, 0xa3, 0x6d, 0x1a, 0x69, 0x5e, 0x50, 0xcd, 0xa5, 0x08, 0x23, 0x29,
	0x32, 0x9d, 0x52, 0x2e, 0x74, 0x46, 0x9a, 0xd5, 0x3f, 0x60, 0xeb, 0xde, 0xed, 0xdf, 0x9b, 0x78,
	0x07, 0xb9, 0x09, 0xcd, 0x74, 0x98, 0x01, 0x08, 0xb2, 0xd5, 0x71, 0xba, 0x6b, 0x41, 0xc3, 0x08,
	0xe7, 0x00, 0xe2, 0xe0, 0x8b, 0x83, 0x9a, 0xf6, 0x1d, 0xbc, 0xa3, 0x82, 0xc6, 0x90, 0xe2, 0xd7,
	0xc8, 0x7d, 0x03, 0x7a, 0xf1, 0x36, 0x9e, 0x79, 0x8b, 0x89, 0xe1, 0x2d, 0xbf, 0xf0, 0xd6, 0x93,
	0x25, 0x0b, 0x1f, 0x21, 0xf7, 0xfc, 0xd7, 0xc2, 0x65, 0xb7, 0xb5, 0xed, 0xd9, 0x91, 0xe3, 0xdd,
	0x0d, 0x13, 0x6f, 0x68, 0x46, 0x0e, 0xee, 0xa1, 0x8d, 0x01, 0x24, 0xa0, 0xe1, 0xe1, 0x8e, 0x7f,
	0x40, 0x9c, 0x9c, 0x7e, 0xbd, 0x6d, 0x3b, 0xdf, 0x6e, 0xdb, 0xce, 0xf7, 0xdb, 0xb6, 0xf3, 0xe9,
	0x47, 0xfb, 0xbf, 0x0f, 0x47, 0xab, 0x8c, 0xca, 0x71, 0xad, 0x52, 0x0e, 0x7f, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x5e, 0xdc, 0x9f, 0x7c, 0x69, 0x05, 0x00, 0x00,
}
