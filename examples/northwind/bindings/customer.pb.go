// Code generated by protoc-gen-go.
// source: customer.proto
// DO NOT EDIT!

/*
Package northwind is a generated protocol buffer package.

It is generated from these files:
	customer.proto
	employee.proto
	image.proto
	order.proto
	product.proto
	region.proto
	shipper.proto
	supplier.proto
	territory.proto

It has these top-level messages:
	Customer
	Employee
	Image
	Order
	OrderItem
	Product
	Region
	Shipper
	Supplier
	Territory
*/
package northwind

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "artisan"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Customer struct {
	Id           string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	ContactName  string `protobuf:"bytes,3,opt,name=contact_name,json=contactName" json:"contact_name,omitempty"`
	ContactTitle string `protobuf:"bytes,4,opt,name=contact_title,json=contactTitle" json:"contact_title,omitempty"`
	Address      string `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	City         string `protobuf:"bytes,6,opt,name=city" json:"city,omitempty"`
	Region       string `protobuf:"bytes,7,opt,name=region" json:"region,omitempty"`
	PostalCode   string `protobuf:"bytes,8,opt,name=postal_code,json=postalCode" json:"postal_code,omitempty"`
	Country      string `protobuf:"bytes,9,opt,name=country" json:"country,omitempty"`
	Phone        string `protobuf:"bytes,10,opt,name=phone" json:"phone,omitempty"`
	Fax          string `protobuf:"bytes,11,opt,name=fax" json:"fax,omitempty"`
	Email        string `protobuf:"bytes,12,opt,name=email" json:"email,omitempty"`
}

func (m *Customer) Reset()                    { *m = Customer{} }
func (m *Customer) String() string            { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()               {}
func (*Customer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Customer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Customer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Customer) GetContactName() string {
	if m != nil {
		return m.ContactName
	}
	return ""
}

func (m *Customer) GetContactTitle() string {
	if m != nil {
		return m.ContactTitle
	}
	return ""
}

func (m *Customer) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Customer) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Customer) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *Customer) GetPostalCode() string {
	if m != nil {
		return m.PostalCode
	}
	return ""
}

func (m *Customer) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Customer) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Customer) GetFax() string {
	if m != nil {
		return m.Fax
	}
	return ""
}

func (m *Customer) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*Customer)(nil), "northwind.Customer")
}

func init() { proto.RegisterFile("customer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x3f, 0x4e, 0x33, 0x31,
	0x10, 0xc5, 0x95, 0x4d, 0xb2, 0xc9, 0xce, 0xe6, 0x4b, 0x61, 0xe9, 0x8b, 0x46, 0x69, 0xf8, 0xd7,
	0x50, 0x41, 0xc1, 0x0d, 0x48, 0x4f, 0x11, 0xd1, 0x47, 0xc6, 0x36, 0xc4, 0xd2, 0xae, 0x67, 0xe5,
	0x9d, 0x08, 0xd2, 0x71, 0x1c, 0x2e, 0xc0, 0x19, 0xb8, 0x16, 0xf2, 0x78, 0x57, 0xa2, 0xf2, 0xbc,
	0xdf, 0x7b, 0x4f, 0x4f, 0x32, 0xac, 0xcd, 0xa9, 0x67, 0x6a, 0x5d, 0xbc, 0xeb, 0x22, 0x31, 0xa9,
	0x2a, 0x50, 0xe4, 0xe3, 0xbb, 0x0f, 0x76, 0xfb, 0x5f, 0x47, 0xf6, 0xbd, 0x0e, 0xf7, 0xc3, 0x9b,
	0x13, 0xd7, 0x3f, 0x05, 0x2c, 0x77, 0x43, 0x49, 0xad, 0xa1, 0xf0, 0x16, 0x27, 0x97, 0x93, 0xdb,
	0x6a, 0x5f, 0x78, 0xab, 0x14, 0xcc, 0x82, 0x6e, 0x1d, 0x16, 0x42, 0xe4, 0x56, 0x57, 0xb0, 0x32,
	0x14, 0x58, 0x1b, 0x3e, 0x88, 0x37, 0x15, 0xaf, 0x1e, 0xd8, 0x53, 0x8a, 0xdc, 0xc0, 0xbf, 0x31,
	0xc2, 0x9e, 0x1b, 0x87, 0x33, 0xc9, 0x8c, 0xbd, 0xe7, 0xc4, 0x14, 0xc2, 0x42, 0x5b, 0x1b, 0x5d,
	0xdf, 0xe3, 0x5c, 0xec, 0x51, 0xa6, 0x55, 0xe3, 0xf9, 0x8c, 0x65, 0x5e, 0x4d, 0xb7, 0xda, 0x40,
	0x19, 0xdd, 0x9b, 0xa7, 0x80, 0x0b, 0xa1, 0x83, 0x52, 0x17, 0x50, 0x77, 0xd4, 0xb3, 0x6e, 0x0e,
	0x86, 0xac, 0xc3, 0xa5, 0x98, 0x90, 0xd1, 0x8e, 0xac, 0xcc, 0x18, 0x3a, 0x05, 0x8e, 0x67, 0xac,
	0xf2, 0xcc, 0x20, 0xd5, 0x16, 0xe6, 0xdd, 0x91, 0x82, 0x43, 0x48, 0xfc, 0x71, 0xf6, 0xf5, 0x8d,
	0x93, 0x7d, 0x46, 0x6a, 0x03, 0xd3, 0x57, 0xfd, 0x81, 0xf5, 0x1f, 0x27, 0x81, 0xd4, 0x71, 0xad,
	0xf6, 0x0d, 0xae, 0xb2, 0xf3, 0x29, 0x1d, 0x41, 0x2f, 0xa5, 0x7c, 0xe8, 0xc3, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x89, 0x80, 0xb5, 0x17, 0x84, 0x01, 0x00, 0x00,
}
