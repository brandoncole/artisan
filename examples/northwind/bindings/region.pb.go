// Code generated by protoc-gen-go.
// source: region.proto
// DO NOT EDIT!

package northwind

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

func init() {
	proto.RegisterType((*Region)(nil), "northwind.Region")
}

func init() { proto.RegisterFile("region.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 98 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4a, 0x4d, 0xcf,
	0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcc, 0xcb, 0x2f, 0x2a, 0xc9, 0x28,
	0xcf, 0xcc, 0x4b, 0x51, 0xb2, 0xe2, 0x62, 0x0b, 0x02, 0x4b, 0x09, 0xf1, 0x71, 0x31, 0x65, 0xa6,
	0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x29, 0x70, 0x71, 0xa7, 0xa4,
	0x16, 0x27, 0x17, 0x65, 0x16, 0x94, 0x64, 0xe6, 0xe7, 0x49, 0x30, 0x81, 0x25, 0x90, 0x85, 0x92,
	0xd8, 0xc0, 0xa6, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x86, 0xf3, 0x00, 0x1b, 0x5d, 0x00,
	0x00, 0x00,
}
