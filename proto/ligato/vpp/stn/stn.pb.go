// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ligato/vpp/stn/stn.proto

package vpp_stn

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

type Rule struct {
	IpAddress            string   `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	Interface            string   `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rule) Reset()         { *m = Rule{} }
func (m *Rule) String() string { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()    {}
func (*Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_f76a9726e8a218b9, []int{0}
}

func (m *Rule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rule.Unmarshal(m, b)
}
func (m *Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rule.Marshal(b, m, deterministic)
}
func (m *Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rule.Merge(m, src)
}
func (m *Rule) XXX_Size() int {
	return xxx_messageInfo_Rule.Size(m)
}
func (m *Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_Rule proto.InternalMessageInfo

func (m *Rule) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *Rule) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func init() {
	proto.RegisterType((*Rule)(nil), "ligato.vpp.stn.Rule")
}

func init() { proto.RegisterFile("ligato/vpp/stn/stn.proto", fileDescriptor_f76a9726e8a218b9) }

var fileDescriptor_f76a9726e8a218b9 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0xc9, 0x4c, 0x4f,
	0x2c, 0xc9, 0xd7, 0x2f, 0x2b, 0x28, 0xd0, 0x2f, 0x2e, 0xc9, 0x03, 0x61, 0xbd, 0x82, 0xa2, 0xfc,
	0x92, 0x7c, 0x21, 0x3e, 0x88, 0x8c, 0x5e, 0x59, 0x41, 0x81, 0x5e, 0x71, 0x49, 0x9e, 0x92, 0x33,
	0x17, 0x4b, 0x50, 0x69, 0x4e, 0xaa, 0x90, 0x2c, 0x17, 0x57, 0x66, 0x41, 0x7c, 0x62, 0x4a, 0x4a,
	0x51, 0x6a, 0x71, 0xb1, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x67, 0x66, 0x81, 0x23, 0x44,
	0x40, 0x48, 0x86, 0x8b, 0x33, 0x33, 0xaf, 0x24, 0xb5, 0x28, 0x2d, 0x31, 0x39, 0x55, 0x82, 0x09,
	0x2a, 0x0b, 0x13, 0x70, 0xb2, 0x88, 0x32, 0x4b, 0xcf, 0xd7, 0x83, 0x9a, 0x9c, 0x09, 0xb6, 0x56,
	0x37, 0x31, 0x3d, 0x35, 0xaf, 0x44, 0xbf, 0xcc, 0x58, 0x1f, 0x6c, 0xaf, 0x3e, 0xaa, 0x83, 0xac,
	0xcb, 0x0a, 0x0a, 0xe2, 0x8b, 0x4b, 0xf2, 0x92, 0xd8, 0xc0, 0xb2, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xa0, 0xe2, 0x6c, 0x64, 0xb1, 0x00, 0x00, 0x00,
}