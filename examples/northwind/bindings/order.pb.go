// Code generated by protoc-gen-go.
// source: order.proto
// DO NOT EDIT!

package northwind

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "artisan"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Order struct {
	Id             string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Customer       *Customer    `protobuf:"bytes,2,opt,name=customer" json:"customer,omitempty"`
	Employee       *Employee    `protobuf:"bytes,3,opt,name=employee" json:"employee,omitempty"`
	OrderDate      string       `protobuf:"bytes,4,opt,name=order_date,json=orderDate" json:"order_date,omitempty"`
	RequiredDate   string       `protobuf:"bytes,5,opt,name=required_date,json=requiredDate" json:"required_date,omitempty"`
	ShippedDate    string       `protobuf:"bytes,6,opt,name=shipped_date,json=shippedDate" json:"shipped_date,omitempty"`
	Shipper        *Shipper     `protobuf:"bytes,7,opt,name=shipper" json:"shipper,omitempty"`
	Freight        float32      `protobuf:"fixed32,8,opt,name=freight" json:"freight,omitempty"`
	ShipName       string       `protobuf:"bytes,9,opt,name=ship_name,json=shipName" json:"ship_name,omitempty"`
	ShipAddress    string       `protobuf:"bytes,10,opt,name=ship_address,json=shipAddress" json:"ship_address,omitempty"`
	ShipCity       string       `protobuf:"bytes,11,opt,name=ship_city,json=shipCity" json:"ship_city,omitempty"`
	ShipRegion     string       `protobuf:"bytes,12,opt,name=ship_region,json=shipRegion" json:"ship_region,omitempty"`
	ShipPostalCode string       `protobuf:"bytes,13,opt,name=ship_postal_code,json=shipPostalCode" json:"ship_postal_code,omitempty"`
	ShipCountry    string       `protobuf:"bytes,14,opt,name=ship_country,json=shipCountry" json:"ship_country,omitempty"`
	Items          []*OrderItem `protobuf:"bytes,15,rep,name=items" json:"items,omitempty"`
	Total          float32      `protobuf:"fixed32,16,opt,name=total" json:"total,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetCustomer() *Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

func (m *Order) GetEmployee() *Employee {
	if m != nil {
		return m.Employee
	}
	return nil
}

func (m *Order) GetOrderDate() string {
	if m != nil {
		return m.OrderDate
	}
	return ""
}

func (m *Order) GetRequiredDate() string {
	if m != nil {
		return m.RequiredDate
	}
	return ""
}

func (m *Order) GetShippedDate() string {
	if m != nil {
		return m.ShippedDate
	}
	return ""
}

func (m *Order) GetShipper() *Shipper {
	if m != nil {
		return m.Shipper
	}
	return nil
}

func (m *Order) GetFreight() float32 {
	if m != nil {
		return m.Freight
	}
	return 0
}

func (m *Order) GetShipName() string {
	if m != nil {
		return m.ShipName
	}
	return ""
}

func (m *Order) GetShipAddress() string {
	if m != nil {
		return m.ShipAddress
	}
	return ""
}

func (m *Order) GetShipCity() string {
	if m != nil {
		return m.ShipCity
	}
	return ""
}

func (m *Order) GetShipRegion() string {
	if m != nil {
		return m.ShipRegion
	}
	return ""
}

func (m *Order) GetShipPostalCode() string {
	if m != nil {
		return m.ShipPostalCode
	}
	return ""
}

func (m *Order) GetShipCountry() string {
	if m != nil {
		return m.ShipCountry
	}
	return ""
}

func (m *Order) GetItems() []*OrderItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Order) GetTotal() float32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type OrderItem struct {
	UnitPrice float32 `protobuf:"fixed32,1,opt,name=unit_price,json=unitPrice" json:"unit_price,omitempty"`
	Quantity  uint32  `protobuf:"varint,2,opt,name=quantity" json:"quantity,omitempty"`
	Discount  float32 `protobuf:"fixed32,3,opt,name=discount" json:"discount,omitempty"`
}

func (m *OrderItem) Reset()                    { *m = OrderItem{} }
func (m *OrderItem) String() string            { return proto.CompactTextString(m) }
func (*OrderItem) ProtoMessage()               {}
func (*OrderItem) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *OrderItem) GetUnitPrice() float32 {
	if m != nil {
		return m.UnitPrice
	}
	return 0
}

func (m *OrderItem) GetQuantity() uint32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *OrderItem) GetDiscount() float32 {
	if m != nil {
		return m.Discount
	}
	return 0
}

func init() {
	proto.RegisterType((*Order)(nil), "northwind.Order")
	proto.RegisterType((*OrderItem)(nil), "northwind.OrderItem")
}

func init() { proto.RegisterFile("order.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x51, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x95, 0x6e, 0xbb, 0x6d, 0x26, 0x6d, 0x58, 0x99, 0x45, 0xb2, 0x8a, 0x04, 0x61, 0x79,
	0x20, 0x20, 0xd4, 0x95, 0x96, 0x13, 0xb0, 0x85, 0x07, 0x5e, 0x60, 0x65, 0x0e, 0x10, 0x99, 0xd8,
	0x6c, 0x2d, 0x35, 0x76, 0xd6, 0x71, 0x84, 0x7a, 0x0b, 0x2e, 0xc3, 0x3b, 0x17, 0xe0, 0x34, 0x5c,
	0x00, 0x79, 0x1c, 0xa7, 0x95, 0x78, 0xaa, 0xfc, 0xff, 0xdf, 0x74, 0xfe, 0x99, 0x0c, 0x64, 0xc6,
	0x0a, 0x69, 0x37, 0xad, 0x35, 0xce, 0x90, 0x54, 0x1b, 0xeb, 0x76, 0x3f, 0x94, 0x16, 0xeb, 0x27,
	0xdc, 0x3a, 0xd5, 0x71, 0x7d, 0x3d, 0xfc, 0x06, 0x62, 0x9d, 0xd7, 0x7d, 0xe7, 0x4c, 0x13, 0x2b,
	0xd6, 0xb9, 0x6c, 0xda, 0xbd, 0x39, 0x48, 0x39, 0xbc, 0x57, 0xdd, 0x4e, 0xb5, 0x6d, 0xb4, 0xaf,
	0xfe, 0x4c, 0x61, 0xf6, 0xc5, 0x37, 0x20, 0x39, 0x4c, 0x94, 0xa0, 0x49, 0x91, 0x94, 0x29, 0x9b,
	0x28, 0x41, 0xae, 0x61, 0x11, 0xff, 0x8a, 0x4e, 0x8a, 0xa4, 0xcc, 0x6e, 0x1e, 0x6f, 0xc6, 0xee,
	0x9b, 0xed, 0x60, 0xb1, 0x11, 0xf2, 0x05, 0xb1, 0x17, 0x3d, 0xfb, 0xaf, 0xe0, 0xe3, 0x60, 0xb1,
	0x11, 0x22, 0x2f, 0x01, 0x70, 0xb6, 0x4a, 0x70, 0x27, 0xe9, 0xd4, 0x77, 0xbe, 0x9d, 0xfe, 0xfc,
	0x45, 0x13, 0x96, 0xa2, 0xfe, 0x81, 0x3b, 0x49, 0x5e, 0xc3, 0xca, 0xca, 0x87, 0x5e, 0x59, 0x29,
	0x02, 0x37, 0x3b, 0xe1, 0x96, 0xd1, 0x42, 0xf4, 0x15, 0x2c, 0xc3, 0x70, 0x03, 0x79, 0x7e, 0x42,
	0x66, 0x83, 0x83, 0xe0, 0x5b, 0x98, 0x0f, 0x5b, 0xa0, 0x73, 0x0c, 0x4a, 0x4e, 0x82, 0x7e, 0x0d,
	0x0e, 0x8b, 0x08, 0x79, 0x06, 0xf3, 0xef, 0x56, 0xaa, 0xfb, 0x9d, 0xa3, 0x8b, 0x22, 0x29, 0x27,
	0xb7, 0xd3, 0xdf, 0x7f, 0x69, 0xc2, 0xa2, 0x48, 0x9e, 0x42, 0xea, 0xd1, 0x4a, 0xf3, 0x46, 0xd2,
	0x14, 0xf7, 0xb7, 0xf0, 0xc2, 0x67, 0xde, 0x48, 0xf2, 0x22, 0x64, 0xaa, 0xb8, 0x10, 0x56, 0x76,
	0x1d, 0x05, 0xf4, 0x31, 0xcd, 0xfb, 0x20, 0x8d, 0xf5, 0xb5, 0x72, 0x07, 0x9a, 0x1d, 0xeb, 0xb7,
	0xca, 0x1d, 0xc8, 0x73, 0x40, 0xb6, 0xb2, 0xf2, 0x5e, 0x19, 0x4d, 0x97, 0x68, 0x83, 0x97, 0x18,
	0x2a, 0xa4, 0x84, 0x0b, 0x04, 0x5a, 0xd3, 0x39, 0xbe, 0xaf, 0x6a, 0x23, 0x24, 0x5d, 0x21, 0x95,
	0x7b, 0xfd, 0x0e, 0xe5, 0xad, 0x11, 0xc7, 0x28, 0xb5, 0xe9, 0xb5, 0xb3, 0x07, 0x9a, 0x1f, 0xa3,
	0x6c, 0x83, 0x44, 0xde, 0xc0, 0x4c, 0x39, 0xd9, 0x74, 0xf4, 0x51, 0x71, 0x56, 0x66, 0x37, 0x97,
	0x27, 0x6b, 0xc1, 0x23, 0xf9, 0xe4, 0x64, 0xc3, 0x02, 0x42, 0x2e, 0x61, 0xe6, 0x8c, 0xe3, 0x7b,
	0x7a, 0xe1, 0x97, 0xc2, 0xc2, 0xe3, 0x4a, 0x43, 0x3a, 0x92, 0xfe, 0x03, 0xf7, 0x5a, 0xb9, 0xaa,
	0xb5, 0xaa, 0x96, 0x78, 0x5a, 0x71, 0x79, 0xa9, 0xd7, 0xef, 0xbc, 0x4c, 0xd6, 0xb0, 0x78, 0xe8,
	0xb9, 0x76, 0x7e, 0x7a, 0x7f, 0x67, 0x2b, 0x36, 0xbe, 0x49, 0x01, 0x0b, 0xa1, 0x3a, 0x0c, 0x8c,
	0x27, 0x15, 0xcb, 0x47, 0xf5, 0xdb, 0x39, 0x9e, 0xf1, 0xbb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xf4, 0xce, 0xd8, 0xd8, 0x26, 0x03, 0x00, 0x00,
}