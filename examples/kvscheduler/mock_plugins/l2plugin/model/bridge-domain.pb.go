// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/bridge-domain.proto

package mock_l2

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BridgeDomain struct {
	Name                 string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Interfaces           []*BridgeDomain_Interface `protobuf:"bytes,10,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *BridgeDomain) Reset()         { *m = BridgeDomain{} }
func (m *BridgeDomain) String() string { return proto.CompactTextString(m) }
func (*BridgeDomain) ProtoMessage()    {}
func (*BridgeDomain) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d3fba68d05addcf, []int{0}
}

func (m *BridgeDomain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BridgeDomain.Unmarshal(m, b)
}
func (m *BridgeDomain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BridgeDomain.Marshal(b, m, deterministic)
}
func (m *BridgeDomain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BridgeDomain.Merge(m, src)
}
func (m *BridgeDomain) XXX_Size() int {
	return xxx_messageInfo_BridgeDomain.Size(m)
}
func (m *BridgeDomain) XXX_DiscardUnknown() {
	xxx_messageInfo_BridgeDomain.DiscardUnknown(m)
}

var xxx_messageInfo_BridgeDomain proto.InternalMessageInfo

func (m *BridgeDomain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BridgeDomain) GetInterfaces() []*BridgeDomain_Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

type BridgeDomain_Interface struct {
	Name                    string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	BridgedVirtualInterface bool     `protobuf:"varint,2,opt,name=bridged_virtual_interface,json=bridgedVirtualInterface,proto3" json:"bridged_virtual_interface,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *BridgeDomain_Interface) Reset()         { *m = BridgeDomain_Interface{} }
func (m *BridgeDomain_Interface) String() string { return proto.CompactTextString(m) }
func (*BridgeDomain_Interface) ProtoMessage()    {}
func (*BridgeDomain_Interface) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d3fba68d05addcf, []int{0, 0}
}

func (m *BridgeDomain_Interface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BridgeDomain_Interface.Unmarshal(m, b)
}
func (m *BridgeDomain_Interface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BridgeDomain_Interface.Marshal(b, m, deterministic)
}
func (m *BridgeDomain_Interface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BridgeDomain_Interface.Merge(m, src)
}
func (m *BridgeDomain_Interface) XXX_Size() int {
	return xxx_messageInfo_BridgeDomain_Interface.Size(m)
}
func (m *BridgeDomain_Interface) XXX_DiscardUnknown() {
	xxx_messageInfo_BridgeDomain_Interface.DiscardUnknown(m)
}

var xxx_messageInfo_BridgeDomain_Interface proto.InternalMessageInfo

func (m *BridgeDomain_Interface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BridgeDomain_Interface) GetBridgedVirtualInterface() bool {
	if m != nil {
		return m.BridgedVirtualInterface
	}
	return false
}

func init() {
	proto.RegisterType((*BridgeDomain)(nil), "mock.l2.BridgeDomain")
	proto.RegisterType((*BridgeDomain_Interface)(nil), "mock.l2.BridgeDomain.Interface")
}

func init() { proto.RegisterFile("model/bridge-domain.proto", fileDescriptor_1d3fba68d05addcf) }

var fileDescriptor_1d3fba68d05addcf = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xcd, 0x4f, 0x49,
	0xcd, 0xd1, 0x4f, 0x2a, 0xca, 0x4c, 0x49, 0x4f, 0xd5, 0x4d, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0xd3,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf, 0xcd, 0x4f, 0xce, 0xd6, 0xcb, 0x31, 0x52, 0x3a,
	0xc0, 0xc8, 0xc5, 0xe3, 0x04, 0x56, 0xe0, 0x02, 0x96, 0x17, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc,
	0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0xec, 0xb9, 0xb8, 0x32, 0xf3,
	0x4a, 0x52, 0x8b, 0xd2, 0x12, 0x93, 0x53, 0x8b, 0x25, 0xb8, 0x14, 0x98, 0x35, 0xb8, 0x8d, 0xe4,
	0xf5, 0xa0, 0x46, 0xe8, 0x21, 0x6b, 0xd7, 0xf3, 0x84, 0xa9, 0x0b, 0x42, 0xd2, 0x22, 0x15, 0xcd,
	0xc5, 0x09, 0x97, 0xc0, 0x6a, 0x83, 0x15, 0x97, 0x24, 0xc4, 0x99, 0x29, 0xf1, 0x65, 0x99, 0x45,
	0x25, 0xa5, 0x89, 0x39, 0xf1, 0x70, 0xed, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0xe2, 0x50,
	0x05, 0x61, 0x10, 0x79, 0xb8, 0x79, 0x49, 0x6c, 0x60, 0x2f, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xc3, 0xb2, 0x31, 0xa7, 0xef, 0x00, 0x00, 0x00,
}
