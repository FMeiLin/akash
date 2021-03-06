// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/market/order.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/ovrclk/akash/x/deployment/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Order_State int32

const (
	// Prefix should start with 0 in enum. So declaring dummy state
	OrderStateInvalid Order_State = 0
	OrderOpen         Order_State = 1
	OrderMatched      Order_State = 2
	OrderClosed       Order_State = 3
)

var Order_State_name = map[int32]string{
	0: "invalid",
	1: "open",
	2: "matched",
	3: "closed",
}

var Order_State_value = map[string]int32{
	"invalid": 0,
	"open":    1,
	"matched": 2,
	"closed":  3,
}

func (x Order_State) String() string {
	return proto.EnumName(Order_State_name, int32(x))
}

func (Order_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_19ed549aa5a11244, []int{2, 0}
}

// MsgCloseOrder defines an SDK message for closing order
type MsgCloseOrder struct {
	OrderID OrderID `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"id" yaml:"id"`
}

func (m *MsgCloseOrder) Reset()         { *m = MsgCloseOrder{} }
func (m *MsgCloseOrder) String() string { return proto.CompactTextString(m) }
func (*MsgCloseOrder) ProtoMessage()    {}
func (*MsgCloseOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_19ed549aa5a11244, []int{0}
}
func (m *MsgCloseOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCloseOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCloseOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCloseOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCloseOrder.Merge(m, src)
}
func (m *MsgCloseOrder) XXX_Size() int {
	return m.Size()
}
func (m *MsgCloseOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCloseOrder.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCloseOrder proto.InternalMessageInfo

func (m *MsgCloseOrder) GetOrderID() OrderID {
	if m != nil {
		return m.OrderID
	}
	return OrderID{}
}

// OrderID stores owner and all other seq numbers
type OrderID struct {
	Owner github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=owner,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"owner" yaml:"owner"`
	DSeq  uint64                                        `protobuf:"varint,2,opt,name=dseq,proto3" json:"dseq" yaml:"dseq"`
	GSeq  uint32                                        `protobuf:"varint,3,opt,name=gseq,proto3" json:"gseq" yaml:"gseq"`
	OSeq  uint32                                        `protobuf:"varint,4,opt,name=oseq,proto3" json:"oseq" yaml:"oseq"`
}

func (m *OrderID) Reset()      { *m = OrderID{} }
func (*OrderID) ProtoMessage() {}
func (*OrderID) Descriptor() ([]byte, []int) {
	return fileDescriptor_19ed549aa5a11244, []int{1}
}
func (m *OrderID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrderID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrderID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderID.Merge(m, src)
}
func (m *OrderID) XXX_Size() int {
	return m.Size()
}
func (m *OrderID) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderID.DiscardUnknown(m)
}

var xxx_messageInfo_OrderID proto.InternalMessageInfo

func (m *OrderID) GetOwner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *OrderID) GetDSeq() uint64 {
	if m != nil {
		return m.DSeq
	}
	return 0
}

func (m *OrderID) GetGSeq() uint32 {
	if m != nil {
		return m.GSeq
	}
	return 0
}

func (m *OrderID) GetOSeq() uint32 {
	if m != nil {
		return m.OSeq
	}
	return 0
}

// Order stores orderID, state of order and other details
type Order struct {
	OrderID OrderID         `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"id" yaml:"id"`
	State   Order_State     `protobuf:"varint,2,opt,name=state,proto3,enum=akash.market.Order_State" json:"state" yaml:"state"`
	StartAt int64           `protobuf:"varint,3,opt,name=start_at,json=startAt,proto3" json:"start-at" yaml:"start-at"`
	Spec    types.GroupSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec" yaml:"spec"`
	CloseAt int64           `protobuf:"varint,5,opt,name=close_at,json=closeAt,proto3" json:"close-at" yaml:"close-at"`
}

func (m *Order) Reset()      { *m = Order{} }
func (*Order) ProtoMessage() {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_19ed549aa5a11244, []int{2}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Order.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return m.Size()
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetOrderID() OrderID {
	if m != nil {
		return m.OrderID
	}
	return OrderID{}
}

func (m *Order) GetState() Order_State {
	if m != nil {
		return m.State
	}
	return OrderStateInvalid
}

func (m *Order) GetStartAt() int64 {
	if m != nil {
		return m.StartAt
	}
	return 0
}

func (m *Order) GetSpec() types.GroupSpec {
	if m != nil {
		return m.Spec
	}
	return types.GroupSpec{}
}

func (m *Order) GetCloseAt() int64 {
	if m != nil {
		return m.CloseAt
	}
	return 0
}

// OrderFilters defines flags for order list filter
type OrderFilters struct {
	Owner github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=owner,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"owner" yaml:"owner"`
	DSeq  uint64                                        `protobuf:"varint,2,opt,name=dseq,proto3" json:"dseq" yaml:"dseq"`
	GSeq  uint32                                        `protobuf:"varint,3,opt,name=gseq,proto3" json:"gseq" yaml:"gseq"`
	OSeq  uint32                                        `protobuf:"varint,4,opt,name=oseq,proto3" json:"oseq" yaml:"oseq"`
	State Order_State                                   `protobuf:"varint,5,opt,name=state,proto3,casttype=Order_State" json:"state" yaml:"state"`
}

func (m *OrderFilters) Reset()         { *m = OrderFilters{} }
func (m *OrderFilters) String() string { return proto.CompactTextString(m) }
func (*OrderFilters) ProtoMessage()    {}
func (*OrderFilters) Descriptor() ([]byte, []int) {
	return fileDescriptor_19ed549aa5a11244, []int{3}
}
func (m *OrderFilters) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OrderFilters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OrderFilters.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OrderFilters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderFilters.Merge(m, src)
}
func (m *OrderFilters) XXX_Size() int {
	return m.Size()
}
func (m *OrderFilters) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderFilters.DiscardUnknown(m)
}

var xxx_messageInfo_OrderFilters proto.InternalMessageInfo

func (m *OrderFilters) GetOwner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *OrderFilters) GetDSeq() uint64 {
	if m != nil {
		return m.DSeq
	}
	return 0
}

func (m *OrderFilters) GetGSeq() uint32 {
	if m != nil {
		return m.GSeq
	}
	return 0
}

func (m *OrderFilters) GetOSeq() uint32 {
	if m != nil {
		return m.OSeq
	}
	return 0
}

func (m *OrderFilters) GetState() Order_State {
	if m != nil {
		return m.State
	}
	return 0
}

func init() {
	proto.RegisterEnum("akash.market.Order_State", Order_State_name, Order_State_value)
	proto.RegisterType((*MsgCloseOrder)(nil), "akash.market.MsgCloseOrder")
	proto.RegisterType((*OrderID)(nil), "akash.market.OrderID")
	proto.RegisterType((*Order)(nil), "akash.market.Order")
	proto.RegisterType((*OrderFilters)(nil), "akash.market.OrderFilters")
}

func init() { proto.RegisterFile("akash/market/order.proto", fileDescriptor_19ed549aa5a11244) }

var fileDescriptor_19ed549aa5a11244 = []byte{
	// 664 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x54, 0xbf, 0x4f, 0xdb, 0x40,
	0x14, 0xb6, 0x83, 0x43, 0xe0, 0x12, 0x4a, 0x6a, 0x15, 0x15, 0x42, 0x9b, 0x8b, 0x6e, 0xa8, 0xb2,
	0xe0, 0x48, 0xb0, 0x65, 0xa8, 0x1a, 0x17, 0x15, 0x81, 0x8a, 0x90, 0x4c, 0xa7, 0x2e, 0xd4, 0xf8,
	0x4e, 0xc6, 0xca, 0x8f, 0x33, 0xbe, 0x83, 0x96, 0xb5, 0x43, 0x55, 0x65, 0xea, 0xd0, 0xa1, 0x4b,
	0x24, 0xa4, 0xfe, 0x33, 0x8c, 0x2c, 0x95, 0x3a, 0x59, 0x55, 0x58, 0x50, 0xc6, 0x8c, 0x4c, 0xd5,
	0xbd, 0x73, 0x08, 0x48, 0xed, 0xd8, 0xa9, 0x93, 0xef, 0xbe, 0xef, 0x7d, 0xef, 0x3d, 0x7f, 0xf7,
	0xee, 0xd0, 0xb2, 0xdf, 0xf6, 0xc5, 0x51, 0xa3, 0xeb, 0x27, 0x6d, 0x26, 0x1b, 0x3c, 0xa1, 0x2c,
	0x71, 0xe2, 0x84, 0x4b, 0x6e, 0x97, 0x80, 0x71, 0x34, 0x53, 0x79, 0x14, 0xf2, 0x90, 0x03, 0xd1,
	0x50, 0x2b, 0x1d, 0x53, 0x79, 0xa2, 0xd5, 0x94, 0xc5, 0x1d, 0x7e, 0xd6, 0x65, 0x3d, 0xd9, 0x08,
	0x13, 0x7e, 0x12, 0x6b, 0x96, 0xb4, 0xd1, 0xc2, 0xae, 0x08, 0x5f, 0x76, 0xb8, 0x60, 0x7b, 0x2a,
	0xb1, 0xfd, 0x06, 0xcd, 0x41, 0x85, 0x83, 0x88, 0x2e, 0x9b, 0x35, 0xb3, 0x5e, 0x5c, 0x5f, 0x72,
	0xee, 0x56, 0x71, 0x20, 0x6c, 0x7b, 0xd3, 0x25, 0x17, 0x29, 0x36, 0x86, 0x29, 0x2e, 0x64, 0xc0,
	0x28, 0xc5, 0xb9, 0x88, 0x8e, 0x53, 0x3c, 0x7f, 0xe6, 0x77, 0x3b, 0x4d, 0x12, 0x51, 0xe2, 0x15,
	0x20, 0xd5, 0x36, 0x6d, 0x5a, 0xd7, 0xe7, 0xd8, 0x20, 0x5f, 0x73, 0x68, 0x12, 0x6d, 0xbf, 0x43,
	0x79, 0xfe, 0xbe, 0xc7, 0x12, 0x28, 0x52, 0x72, 0x77, 0x46, 0x29, 0xd6, 0xc0, 0x38, 0xc5, 0x25,
	0x9d, 0x01, 0xb6, 0xe4, 0x26, 0xc5, 0x6b, 0x61, 0x24, 0x8f, 0x4e, 0x0e, 0x9d, 0x80, 0x77, 0x1b,
	0x01, 0x17, 0x5d, 0x2e, 0xb2, 0xcf, 0x9a, 0xa0, 0xed, 0x86, 0x3c, 0x8b, 0x99, 0x70, 0x5a, 0x41,
	0xd0, 0xa2, 0x34, 0x61, 0x42, 0x78, 0x3a, 0x8f, 0xbd, 0x81, 0x2c, 0x2a, 0xd8, 0xf1, 0x72, 0xae,
	0x66, 0xd6, 0x2d, 0x17, 0x0f, 0x53, 0x6c, 0x6d, 0xee, 0xb3, 0xe3, 0x51, 0x8a, 0x01, 0x1f, 0xa7,
	0xb8, 0xa8, 0xeb, 0xa8, 0x1d, 0xf1, 0x00, 0x54, 0xa2, 0x50, 0x89, 0x66, 0x6a, 0x66, 0x7d, 0x41,
	0x8b, 0xb6, 0x32, 0x51, 0x78, 0x4f, 0x14, 0x6a, 0x51, 0x98, 0x89, 0xb8, 0x12, 0x59, 0x53, 0xd1,
	0x5e, 0x26, 0xe2, 0xf7, 0x44, 0x5c, 0x8b, 0xd4, 0xa7, 0x39, 0xf7, 0xed, 0x1c, 0x1b, 0xd7, 0xe7,
	0xd8, 0x24, 0x1f, 0x2d, 0x94, 0xff, 0x87, 0xe6, 0xdb, 0x3b, 0x28, 0x2f, 0xa4, 0x2f, 0x19, 0x38,
	0xf1, 0x60, 0x7d, 0xe5, 0x0f, 0x29, 0x9d, 0x7d, 0x15, 0xe0, 0xae, 0xa8, 0x53, 0x80, 0xd8, 0xe9,
	0x29, 0xc0, 0x96, 0x78, 0x1a, 0xb6, 0x9b, 0x68, 0x4e, 0x48, 0x3f, 0x91, 0x07, 0xbe, 0x04, 0x8f,
	0x66, 0x5c, 0x3c, 0x4a, 0xb1, 0xc6, 0xd6, 0x7c, 0x39, 0x4e, 0xf1, 0xe2, 0xad, 0x0c, 0x10, 0xe2,
	0x15, 0x60, 0xd9, 0x92, 0xf6, 0x6b, 0x64, 0x89, 0x98, 0x05, 0x60, 0x53, 0x71, 0x7d, 0x35, 0x6b,
	0x63, 0x3a, 0x98, 0xce, 0x96, 0x1a, 0xcc, 0xfd, 0x98, 0x05, 0xee, 0xaa, 0xfa, 0x3f, 0xe5, 0x9f,
	0x12, 0x4c, 0xfd, 0x53, 0x3b, 0xe2, 0x01, 0xa8, 0x3a, 0x09, 0xd4, 0xd8, 0xaa, 0x4e, 0xf2, 0xd3,
	0x4e, 0x00, 0xbb, 0xd7, 0xc9, 0x04, 0x21, 0x5e, 0x01, 0x96, 0x2d, 0x49, 0x3e, 0x99, 0x28, 0x0f,
	0x7f, 0x6c, 0x13, 0x54, 0x88, 0x7a, 0xa7, 0x7e, 0x27, 0xa2, 0x65, 0xa3, 0xb2, 0xd4, 0x1f, 0xd4,
	0x1e, 0x82, 0x1f, 0x40, 0x6e, 0x6b, 0xc2, 0x7e, 0x8c, 0x2c, 0x1e, 0xb3, 0x5e, 0xd9, 0xac, 0x2c,
	0xf4, 0x07, 0xb5, 0x79, 0x08, 0xd8, 0x8b, 0x59, 0xcf, 0x7e, 0x8a, 0x0a, 0x5d, 0x5f, 0x06, 0x47,
	0x8c, 0x96, 0x73, 0x95, 0x72, 0x7f, 0x50, 0x2b, 0x01, 0xb7, 0xab, 0x31, 0x7b, 0x15, 0xcd, 0x42,
	0x41, 0x5a, 0x9e, 0xa9, 0x2c, 0xf6, 0x07, 0xb5, 0x22, 0xb0, 0x70, 0xd7, 0x68, 0xc5, 0xfa, 0xfc,
	0xbd, 0x6a, 0xdc, 0x0e, 0x81, 0x41, 0x7e, 0xe4, 0x90, 0x56, 0xbf, 0x8a, 0x3a, 0x92, 0x25, 0xe2,
	0xbf, 0xbf, 0x20, 0xf6, 0xf3, 0xc9, 0xd8, 0xaa, 0xd3, 0xcd, 0xbb, 0xf5, 0xbf, 0xce, 0xe6, 0x4d,
	0x8a, 0xb5, 0xb5, 0x07, 0x70, 0x6c, 0xd9, 0xa8, 0xea, 0x37, 0xc7, 0x7d, 0x71, 0x31, 0xac, 0x9a,
	0x97, 0xc3, 0xaa, 0xf9, 0x6b, 0x58, 0x35, 0xbf, 0x5c, 0x55, 0x8d, 0xcb, 0xab, 0xaa, 0xf1, 0xf3,
	0xaa, 0x6a, 0xbc, 0x7d, 0x76, 0xc7, 0x34, 0x7e, 0x9a, 0x04, 0x9d, 0x76, 0x43, 0x3f, 0x95, 0x1f,
	0x26, 0x4f, 0x2d, 0x18, 0x77, 0x38, 0x0b, 0x2f, 0xe5, 0xc6, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x0e, 0x7a, 0x4e, 0x7f, 0x87, 0x05, 0x00, 0x00,
}

func (this *OrderID) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*OrderID)
	if !ok {
		that2, ok := that.(OrderID)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Owner, that1.Owner) {
		return false
	}
	if this.DSeq != that1.DSeq {
		return false
	}
	if this.GSeq != that1.GSeq {
		return false
	}
	if this.OSeq != that1.OSeq {
		return false
	}
	return true
}
func (m *MsgCloseOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCloseOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCloseOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.OrderID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *OrderID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OSeq != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.OSeq))
		i--
		dAtA[i] = 0x20
	}
	if m.GSeq != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.GSeq))
		i--
		dAtA[i] = 0x18
	}
	if m.DSeq != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.DSeq))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Order) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Order) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Order) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CloseAt != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.CloseAt))
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.Spec.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.StartAt != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.StartAt))
		i--
		dAtA[i] = 0x18
	}
	if m.State != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.OrderID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *OrderFilters) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderFilters) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OrderFilters) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.State != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x28
	}
	if m.OSeq != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.OSeq))
		i--
		dAtA[i] = 0x20
	}
	if m.GSeq != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.GSeq))
		i--
		dAtA[i] = 0x18
	}
	if m.DSeq != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.DSeq))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrder(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCloseOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.OrderID.Size()
	n += 1 + l + sovOrder(uint64(l))
	return n
}

func (m *OrderID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	if m.DSeq != 0 {
		n += 1 + sovOrder(uint64(m.DSeq))
	}
	if m.GSeq != 0 {
		n += 1 + sovOrder(uint64(m.GSeq))
	}
	if m.OSeq != 0 {
		n += 1 + sovOrder(uint64(m.OSeq))
	}
	return n
}

func (m *Order) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.OrderID.Size()
	n += 1 + l + sovOrder(uint64(l))
	if m.State != 0 {
		n += 1 + sovOrder(uint64(m.State))
	}
	if m.StartAt != 0 {
		n += 1 + sovOrder(uint64(m.StartAt))
	}
	l = m.Spec.Size()
	n += 1 + l + sovOrder(uint64(l))
	if m.CloseAt != 0 {
		n += 1 + sovOrder(uint64(m.CloseAt))
	}
	return n
}

func (m *OrderFilters) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	if m.DSeq != 0 {
		n += 1 + sovOrder(uint64(m.DSeq))
	}
	if m.GSeq != 0 {
		n += 1 + sovOrder(uint64(m.GSeq))
	}
	if m.OSeq != 0 {
		n += 1 + sovOrder(uint64(m.OSeq))
	}
	if m.State != 0 {
		n += 1 + sovOrder(uint64(m.State))
	}
	return n
}

func sovOrder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrder(x uint64) (n int) {
	return sovOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCloseOrder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCloseOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCloseOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OrderID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrder
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrder
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
func (m *OrderID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OrderID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DSeq", wireType)
			}
			m.DSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DSeq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GSeq", wireType)
			}
			m.GSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GSeq |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OSeq", wireType)
			}
			m.OSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OSeq |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrder
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrder
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
func (m *Order) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Order: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Order: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OrderID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= Order_State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartAt", wireType)
			}
			m.StartAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CloseAt", wireType)
			}
			m.CloseAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CloseAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrder
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrder
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
func (m *OrderFilters) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OrderFilters: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderFilters: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DSeq", wireType)
			}
			m.DSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DSeq |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GSeq", wireType)
			}
			m.GSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GSeq |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OSeq", wireType)
			}
			m.OSeq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OSeq |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= Order_State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOrder
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOrder
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
func skipOrder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrder
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
					return 0, ErrIntOverflowOrder
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowOrder
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
			if length < 0 {
				return 0, ErrInvalidLengthOrder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrder = fmt.Errorf("proto: unexpected end of group")
)
