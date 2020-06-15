// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ligato/govppmux/metrics.proto

package govppmux

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

type Metrics struct {
	ChannelsCreated      uint64   `protobuf:"varint,1,opt,name=channels_created,json=channelsCreated,proto3" json:"channels_created,omitempty"`
	ChannelsOpen         uint64   `protobuf:"varint,2,opt,name=channels_open,json=channelsOpen,proto3" json:"channels_open,omitempty"`
	RequestsSent         uint64   `protobuf:"varint,3,opt,name=requests_sent,json=requestsSent,proto3" json:"requests_sent,omitempty"`
	RequestsDone         uint64   `protobuf:"varint,4,opt,name=requests_done,json=requestsDone,proto3" json:"requests_done,omitempty"`
	RequestsFail         uint64   `protobuf:"varint,5,opt,name=requests_fail,json=requestsFail,proto3" json:"requests_fail,omitempty"`
	RequestReplies       uint64   `protobuf:"varint,6,opt,name=request_replies,json=requestReplies,proto3" json:"request_replies,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metrics) Reset()         { *m = Metrics{} }
func (m *Metrics) String() string { return proto.CompactTextString(m) }
func (*Metrics) ProtoMessage()    {}
func (*Metrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_a54ec93fe2fe8fda, []int{0}
}

func (m *Metrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metrics.Unmarshal(m, b)
}
func (m *Metrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metrics.Marshal(b, m, deterministic)
}
func (m *Metrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metrics.Merge(m, src)
}
func (m *Metrics) XXX_Size() int {
	return xxx_messageInfo_Metrics.Size(m)
}
func (m *Metrics) XXX_DiscardUnknown() {
	xxx_messageInfo_Metrics.DiscardUnknown(m)
}

var xxx_messageInfo_Metrics proto.InternalMessageInfo

func (m *Metrics) GetChannelsCreated() uint64 {
	if m != nil {
		return m.ChannelsCreated
	}
	return 0
}

func (m *Metrics) GetChannelsOpen() uint64 {
	if m != nil {
		return m.ChannelsOpen
	}
	return 0
}

func (m *Metrics) GetRequestsSent() uint64 {
	if m != nil {
		return m.RequestsSent
	}
	return 0
}

func (m *Metrics) GetRequestsDone() uint64 {
	if m != nil {
		return m.RequestsDone
	}
	return 0
}

func (m *Metrics) GetRequestsFail() uint64 {
	if m != nil {
		return m.RequestsFail
	}
	return 0
}

func (m *Metrics) GetRequestReplies() uint64 {
	if m != nil {
		return m.RequestReplies
	}
	return 0
}

func init() {
	proto.RegisterType((*Metrics)(nil), "ligato.govppmux.Metrics")
}

func init() {
	proto.RegisterFile("ligato/govppmux/metrics.proto", fileDescriptor_a54ec93fe2fe8fda)
}

var fileDescriptor_a54ec93fe2fe8fda = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xd0, 0xb1, 0x4b, 0xc4, 0x30,
	0x14, 0xc7, 0x71, 0xaa, 0xe7, 0x09, 0x41, 0xad, 0x74, 0xca, 0x22, 0x88, 0x0e, 0xea, 0x60, 0x33,
	0xdc, 0x22, 0xba, 0xa9, 0xb8, 0x89, 0x70, 0x6e, 0x2e, 0x25, 0xf6, 0x9e, 0x35, 0x90, 0x7b, 0xef,
	0x99, 0xbc, 0x2b, 0xfe, 0xdb, 0xfe, 0x07, 0x62, 0xd3, 0x08, 0xbd, 0x2d, 0xf9, 0xbe, 0xcf, 0xf4,
	0x53, 0x27, 0xde, 0x75, 0x56, 0xc8, 0x74, 0xd4, 0x33, 0xaf, 0x37, 0xdf, 0x66, 0x0d, 0x12, 0x5c,
	0x1b, 0x6b, 0x0e, 0x24, 0x54, 0x95, 0xe9, 0x5c, 0xe7, 0xf3, 0xd9, 0x4f, 0xa1, 0xf6, 0x9f, 0x13,
	0xa9, 0xae, 0xd4, 0x71, 0xfb, 0x69, 0x11, 0xc1, 0xc7, 0xa6, 0x0d, 0x60, 0x05, 0x56, 0xba, 0x38,
	0x2d, 0x2e, 0x67, 0xcb, 0x32, 0xf7, 0x87, 0x94, 0xab, 0x73, 0x75, 0xf8, 0x4f, 0x89, 0x01, 0xf5,
	0xce, 0xe0, 0x0e, 0x72, 0x7c, 0x61, 0xc0, 0x3f, 0x14, 0xe0, 0x6b, 0x03, 0x51, 0x62, 0x13, 0x01,
	0x45, 0xef, 0x26, 0x94, 0xe3, 0x2b, 0xa0, 0x4c, 0xd0, 0x8a, 0x10, 0xf4, 0x6c, 0x8a, 0x1e, 0x09,
	0x61, 0x82, 0x3e, 0xac, 0xf3, 0x7a, 0x6f, 0x8a, 0x9e, 0xac, 0xf3, 0xd5, 0x85, 0x2a, 0xc7, 0x7f,
	0x13, 0x80, 0xbd, 0x83, 0xa8, 0xe7, 0x03, 0x3b, 0x1a, 0xf3, 0x32, 0xd5, 0xfb, 0xdb, 0xb7, 0x9b,
	0x8e, 0xea, 0x71, 0x09, 0x47, 0xa6, 0x67, 0xbe, 0xb6, 0x1d, 0xa0, 0x98, 0x7e, 0x61, 0x86, 0x9d,
	0xcc, 0xd6, 0x8a, 0x77, 0xf9, 0xf1, 0x3e, 0x1f, 0xee, 0x8b, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xdc, 0xc0, 0xc5, 0xb4, 0x68, 0x01, 0x00, 0x00,
}
