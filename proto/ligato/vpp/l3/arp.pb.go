// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ligato/vpp/l3/arp.proto

package vpp_l3

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

type ARPEntry struct {
	Interface            string   `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
	IpAddress            string   `protobuf:"bytes,2,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	PhysAddress          string   `protobuf:"bytes,3,opt,name=phys_address,json=physAddress,proto3" json:"phys_address,omitempty"`
	Static               bool     `protobuf:"varint,4,opt,name=static,proto3" json:"static,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ARPEntry) Reset()         { *m = ARPEntry{} }
func (m *ARPEntry) String() string { return proto.CompactTextString(m) }
func (*ARPEntry) ProtoMessage()    {}
func (*ARPEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_667541aaa11bae35, []int{0}
}

func (m *ARPEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ARPEntry.Unmarshal(m, b)
}
func (m *ARPEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ARPEntry.Marshal(b, m, deterministic)
}
func (m *ARPEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ARPEntry.Merge(m, src)
}
func (m *ARPEntry) XXX_Size() int {
	return xxx_messageInfo_ARPEntry.Size(m)
}
func (m *ARPEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ARPEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ARPEntry proto.InternalMessageInfo

func (m *ARPEntry) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *ARPEntry) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *ARPEntry) GetPhysAddress() string {
	if m != nil {
		return m.PhysAddress
	}
	return ""
}

func (m *ARPEntry) GetStatic() bool {
	if m != nil {
		return m.Static
	}
	return false
}

func init() {
	proto.RegisterType((*ARPEntry)(nil), "ligato.vpp.l3.ARPEntry")
}

func init() {
	proto.RegisterFile("ligato/vpp/l3/arp.proto", fileDescriptor_667541aaa11bae35)
}

var fileDescriptor_667541aaa11bae35 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0xc9, 0x4c, 0x4f,
	0x2c, 0xc9, 0xd7, 0x2f, 0x2b, 0x28, 0xd0, 0xcf, 0x31, 0xd6, 0x4f, 0x2c, 0x2a, 0xd0, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x85, 0x48, 0xe8, 0x95, 0x15, 0x14, 0xe8, 0xe5, 0x18, 0x2b, 0x35,
	0x31, 0x72, 0x71, 0x38, 0x06, 0x05, 0xb8, 0xe6, 0x95, 0x14, 0x55, 0x0a, 0xc9, 0x70, 0x71, 0x66,
	0xe6, 0x95, 0xa4, 0x16, 0xa5, 0x25, 0x26, 0xa7, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x21,
	0x04, 0x84, 0x64, 0xb9, 0xb8, 0x32, 0x0b, 0xe2, 0x13, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x25,
	0x98, 0xa0, 0xd2, 0x05, 0x8e, 0x10, 0x01, 0x21, 0x45, 0x2e, 0x9e, 0x82, 0x8c, 0xca, 0x62, 0xb8,
	0x02, 0x66, 0xb0, 0x02, 0x6e, 0x90, 0x18, 0x4c, 0x89, 0x18, 0x17, 0x5b, 0x71, 0x49, 0x62, 0x49,
	0x66, 0xb2, 0x04, 0x8b, 0x02, 0xa3, 0x06, 0x47, 0x10, 0x94, 0xe7, 0x64, 0x16, 0x65, 0x92, 0x9e,
	0xaf, 0x07, 0x75, 0x58, 0x26, 0xd8, 0xd1, 0xba, 0x89, 0xe9, 0xa9, 0x79, 0x25, 0xfa, 0x65, 0xc6,
	0xfa, 0x60, 0x67, 0xeb, 0xa3, 0x78, 0xc7, 0xba, 0xac, 0xa0, 0x20, 0x3e, 0xc7, 0x38, 0x89, 0x0d,
	0x2c, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x20, 0x67, 0xcc, 0xdd, 0xed, 0x00, 0x00, 0x00,
}
