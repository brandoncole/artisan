// Code generated by protoc-gen-go.
// source: shipper.proto
// DO NOT EDIT!

package northwind

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/brandoncole/artisan/artisan/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Shipper struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Phone string `protobuf:"bytes,3,opt,name=phone" json:"phone,omitempty"`
}

func (m *Shipper) Reset()                    { *m = Shipper{} }
func (m *Shipper) String() string            { return proto.CompactTextString(m) }
func (*Shipper) ProtoMessage()               {}
func (*Shipper) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *Shipper) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Shipper) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Shipper) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type ShipperFK struct {
	// Types that are valid to be assigned to Shipper:
	//	*ShipperFK_Id
	//	*ShipperFK_Instance
	Shipper isShipperFK_Shipper `protobuf_oneof:"shipper"`
}

func (m *ShipperFK) Reset()                    { *m = ShipperFK{} }
func (m *ShipperFK) String() string            { return proto.CompactTextString(m) }
func (*ShipperFK) ProtoMessage()               {}
func (*ShipperFK) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

type isShipperFK_Shipper interface {
	isShipperFK_Shipper()
}

type ShipperFK_Id struct {
	Id string `protobuf:"bytes,1,opt,name=id,oneof"`
}
type ShipperFK_Instance struct {
	Instance *Shipper `protobuf:"bytes,2,opt,name=instance,oneof"`
}

func (*ShipperFK_Id) isShipperFK_Shipper()       {}
func (*ShipperFK_Instance) isShipperFK_Shipper() {}

func (m *ShipperFK) GetShipper() isShipperFK_Shipper {
	if m != nil {
		return m.Shipper
	}
	return nil
}

func (m *ShipperFK) GetId() string {
	if x, ok := m.GetShipper().(*ShipperFK_Id); ok {
		return x.Id
	}
	return ""
}

func (m *ShipperFK) GetInstance() *Shipper {
	if x, ok := m.GetShipper().(*ShipperFK_Instance); ok {
		return x.Instance
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ShipperFK) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ShipperFK_OneofMarshaler, _ShipperFK_OneofUnmarshaler, _ShipperFK_OneofSizer, []interface{}{
		(*ShipperFK_Id)(nil),
		(*ShipperFK_Instance)(nil),
	}
}

func _ShipperFK_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ShipperFK)
	// shipper
	switch x := m.Shipper.(type) {
	case *ShipperFK_Id:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Id)
	case *ShipperFK_Instance:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Instance); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ShipperFK.Shipper has unexpected type %T", x)
	}
	return nil
}

func _ShipperFK_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ShipperFK)
	switch tag {
	case 1: // shipper.id
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Shipper = &ShipperFK_Id{x}
		return true, err
	case 2: // shipper.instance
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Shipper)
		err := b.DecodeMessage(msg)
		m.Shipper = &ShipperFK_Instance{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ShipperFK_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ShipperFK)
	// shipper
	switch x := m.Shipper.(type) {
	case *ShipperFK_Id:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Id)))
		n += len(x.Id)
	case *ShipperFK_Instance:
		s := proto.Size(x.Instance)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Shipper)(nil), "northwind.Shipper")
	proto.RegisterType((*ShipperFK)(nil), "northwind.ShipperFK")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Shippers service

type ShippersClient interface {
}

type shippersClient struct {
	cc *grpc.ClientConn
}

func NewShippersClient(cc *grpc.ClientConn) ShippersClient {
	return &shippersClient{cc}
}

// Server API for Shippers service

type ShippersServer interface {
}

func RegisterShippersServer(s *grpc.Server, srv ShippersServer) {
	s.RegisterService(&_Shippers_serviceDesc, srv)
}

var _Shippers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "northwind.Shippers",
	HandlerType: (*ShippersServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "shipper.proto",
}

func init() { proto.RegisterFile("shipper.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0xce, 0xc8, 0x2c,
	0x28, 0x48, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcc, 0xcb, 0x2f, 0x2a, 0xc9,
	0x28, 0xcf, 0xcc, 0x4b, 0x91, 0x12, 0x4d, 0x2c, 0x2a, 0xc9, 0x2c, 0x4e, 0xcc, 0xd3, 0x87, 0xd2,
	0x10, 0x15, 0x4a, 0xfe, 0x5c, 0xec, 0xc1, 0x10, 0x2d, 0x42, 0x22, 0x5c, 0x4c, 0x99, 0x29, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x4e, 0x2c, 0x0d, 0x5b, 0x25, 0x18, 0x83, 0x98, 0x32, 0x53, 0x84,
	0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x98, 0x40, 0xe2, 0x41, 0x60, 0xb6, 0x90, 0x14,
	0x17, 0x6b, 0x41, 0x46, 0x7e, 0x5e, 0xaa, 0x04, 0x33, 0x44, 0xf1, 0x83, 0xcf, 0x12, 0x8c, 0x41,
	0x10, 0x21, 0xa5, 0x28, 0x2e, 0x4e, 0xa8, 0x81, 0x6e, 0xde, 0x42, 0x02, 0x08, 0x23, 0x3d, 0x18,
	0xc0, 0xc6, 0x19, 0x70, 0x71, 0x64, 0xe6, 0x15, 0x97, 0x24, 0xe6, 0x25, 0x43, 0x8c, 0xe4, 0x36,
	0x12, 0xd2, 0x83, 0x3b, 0x52, 0x0f, 0xaa, 0xd3, 0x83, 0x21, 0x08, 0xae, 0xca, 0x89, 0x93, 0x8b,
	0x1d, 0xea, 0x29, 0x23, 0x79, 0x2e, 0x0e, 0xa8, 0x8a, 0x62, 0x29, 0xe1, 0x43, 0x4f, 0x64, 0xf9,
	0x61, 0x7e, 0x81, 0x8a, 0x26, 0xb1, 0x81, 0x3d, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x45,
	0xea, 0x29, 0x36, 0x07, 0x01, 0x00, 0x00,
}
