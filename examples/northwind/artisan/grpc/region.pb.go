// Code generated by protoc-gen-go.
// source: region.proto
// DO NOT EDIT!

package northwind

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/brandoncole/artisan/artisan"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Region struct {
	Id          string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *Region) Reset()                    { *m = Region{} }
func (m *Region) String() string            { return proto.CompactTextString(m) }
func (*Region) ProtoMessage()               {}
func (*Region) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *Region) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Region) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type RegionFK struct {
	// Types that are valid to be assigned to Region:
	//	*RegionFK_Id
	//	*RegionFK_Instance
	Region isRegionFK_Region `protobuf_oneof:"region"`
}

func (m *RegionFK) Reset()                    { *m = RegionFK{} }
func (m *RegionFK) String() string            { return proto.CompactTextString(m) }
func (*RegionFK) ProtoMessage()               {}
func (*RegionFK) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

type isRegionFK_Region interface {
	isRegionFK_Region()
}

type RegionFK_Id struct {
	Id string `protobuf:"bytes,1,opt,name=id,oneof"`
}
type RegionFK_Instance struct {
	Instance *Region `protobuf:"bytes,2,opt,name=instance,oneof"`
}

func (*RegionFK_Id) isRegionFK_Region()       {}
func (*RegionFK_Instance) isRegionFK_Region() {}

func (m *RegionFK) GetRegion() isRegionFK_Region {
	if m != nil {
		return m.Region
	}
	return nil
}

func (m *RegionFK) GetId() string {
	if x, ok := m.GetRegion().(*RegionFK_Id); ok {
		return x.Id
	}
	return ""
}

func (m *RegionFK) GetInstance() *Region {
	if x, ok := m.GetRegion().(*RegionFK_Instance); ok {
		return x.Instance
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RegionFK) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RegionFK_OneofMarshaler, _RegionFK_OneofUnmarshaler, _RegionFK_OneofSizer, []interface{}{
		(*RegionFK_Id)(nil),
		(*RegionFK_Instance)(nil),
	}
}

func _RegionFK_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RegionFK)
	// region
	switch x := m.Region.(type) {
	case *RegionFK_Id:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Id)
	case *RegionFK_Instance:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Instance); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RegionFK.Region has unexpected type %T", x)
	}
	return nil
}

func _RegionFK_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RegionFK)
	switch tag {
	case 1: // region.id
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Region = &RegionFK_Id{x}
		return true, err
	case 2: // region.instance
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Region)
		err := b.DecodeMessage(msg)
		m.Region = &RegionFK_Instance{msg}
		return true, err
	default:
		return false, nil
	}
}

func _RegionFK_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RegionFK)
	// region
	switch x := m.Region.(type) {
	case *RegionFK_Id:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Id)))
		n += len(x.Id)
	case *RegionFK_Instance:
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
	proto.RegisterType((*Region)(nil), "northwind.Region")
	proto.RegisterType((*RegionFK)(nil), "northwind.RegionFK")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Regions service

type RegionsClient interface {
}

type regionsClient struct {
	cc *grpc.ClientConn
}

func NewRegionsClient(cc *grpc.ClientConn) RegionsClient {
	return &regionsClient{cc}
}

// Server API for Regions service

type RegionsServer interface {
}

func RegisterRegionsServer(s *grpc.Server, srv RegionsServer) {
	s.RegisterService(&_Regions_serviceDesc, srv)
}

var _Regions_serviceDesc = grpc.ServiceDesc{
	ServiceName: "northwind.Regions",
	HandlerType: (*RegionsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "region.proto",
}

func init() { proto.RegisterFile("region.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4a, 0x4d, 0xcf,
	0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcc, 0xcb, 0x2f, 0x2a, 0xc9, 0x28,
	0xcf, 0xcc, 0x4b, 0x91, 0x12, 0x4d, 0x2c, 0x2a, 0xc9, 0x2c, 0x4e, 0xcc, 0xd3, 0x87, 0xd2, 0x10,
	0x15, 0x4a, 0x0e, 0x5c, 0x6c, 0x41, 0x60, 0x1d, 0x42, 0x22, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x4e, 0x2c, 0x0d, 0x5b, 0x25, 0x18, 0x83, 0x98, 0x32, 0x53, 0x84, 0x14,
	0xb8, 0xb8, 0x53, 0x52, 0x8b, 0x93, 0x8b, 0x32, 0x0b, 0x4a, 0x32, 0xf3, 0xf3, 0x24, 0x98, 0x40,
	0xd2, 0x41, 0xc8, 0x42, 0x4a, 0xe1, 0x5c, 0x1c, 0x10, 0x13, 0xdc, 0xbc, 0x85, 0x04, 0x10, 0x66,
	0x78, 0x30, 0x80, 0xf5, 0xeb, 0x73, 0x71, 0x64, 0xe6, 0x15, 0x97, 0x24, 0xe6, 0x25, 0xa7, 0x82,
	0x35, 0x73, 0x1b, 0x09, 0xea, 0xc1, 0x1d, 0xa5, 0x07, 0xd1, 0xe8, 0xc1, 0x10, 0x04, 0x57, 0xe4,
	0xc4, 0xc1, 0xc5, 0x06, 0xf1, 0x82, 0x91, 0x2c, 0x17, 0x3b, 0x44, 0xbe, 0x58, 0x4a, 0xe8, 0xd0,
	0x13, 0x59, 0x3e, 0x98, 0xc3, 0x21, 0x82, 0x49, 0x6c, 0x60, 0x0f, 0x18, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xc9, 0x5c, 0x9d, 0xad, 0xf2, 0x00, 0x00, 0x00,
}