// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ligato/vpp/l3/l3.proto

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

type IPScanNeighbor_Mode int32

const (
	IPScanNeighbor_DISABLED IPScanNeighbor_Mode = 0
	IPScanNeighbor_IPV4     IPScanNeighbor_Mode = 1
	IPScanNeighbor_IPV6     IPScanNeighbor_Mode = 2
	IPScanNeighbor_BOTH     IPScanNeighbor_Mode = 3
)

var IPScanNeighbor_Mode_name = map[int32]string{
	0: "DISABLED",
	1: "IPV4",
	2: "IPV6",
	3: "BOTH",
}

var IPScanNeighbor_Mode_value = map[string]int32{
	"DISABLED": 0,
	"IPV4":     1,
	"IPV6":     2,
	"BOTH":     3,
}

func (x IPScanNeighbor_Mode) String() string {
	return proto.EnumName(IPScanNeighbor_Mode_name, int32(x))
}

func (IPScanNeighbor_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eb46f906a6f7c0e7, []int{1, 0}
}

type ProxyARP struct {
	Interfaces           []*ProxyARP_Interface `protobuf:"bytes,1,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
	Ranges               []*ProxyARP_Range     `protobuf:"bytes,2,rep,name=ranges,proto3" json:"ranges,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ProxyARP) Reset()         { *m = ProxyARP{} }
func (m *ProxyARP) String() string { return proto.CompactTextString(m) }
func (*ProxyARP) ProtoMessage()    {}
func (*ProxyARP) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb46f906a6f7c0e7, []int{0}
}

func (m *ProxyARP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyARP.Unmarshal(m, b)
}
func (m *ProxyARP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyARP.Marshal(b, m, deterministic)
}
func (m *ProxyARP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyARP.Merge(m, src)
}
func (m *ProxyARP) XXX_Size() int {
	return xxx_messageInfo_ProxyARP.Size(m)
}
func (m *ProxyARP) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyARP.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyARP proto.InternalMessageInfo

func (m *ProxyARP) GetInterfaces() []*ProxyARP_Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *ProxyARP) GetRanges() []*ProxyARP_Range {
	if m != nil {
		return m.Ranges
	}
	return nil
}

type ProxyARP_Interface struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProxyARP_Interface) Reset()         { *m = ProxyARP_Interface{} }
func (m *ProxyARP_Interface) String() string { return proto.CompactTextString(m) }
func (*ProxyARP_Interface) ProtoMessage()    {}
func (*ProxyARP_Interface) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb46f906a6f7c0e7, []int{0, 0}
}

func (m *ProxyARP_Interface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyARP_Interface.Unmarshal(m, b)
}
func (m *ProxyARP_Interface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyARP_Interface.Marshal(b, m, deterministic)
}
func (m *ProxyARP_Interface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyARP_Interface.Merge(m, src)
}
func (m *ProxyARP_Interface) XXX_Size() int {
	return xxx_messageInfo_ProxyARP_Interface.Size(m)
}
func (m *ProxyARP_Interface) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyARP_Interface.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyARP_Interface proto.InternalMessageInfo

func (m *ProxyARP_Interface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ProxyARP_Range struct {
	FirstIpAddr          string   `protobuf:"bytes,1,opt,name=first_ip_addr,json=firstIpAddr,proto3" json:"first_ip_addr,omitempty"`
	LastIpAddr           string   `protobuf:"bytes,2,opt,name=last_ip_addr,json=lastIpAddr,proto3" json:"last_ip_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProxyARP_Range) Reset()         { *m = ProxyARP_Range{} }
func (m *ProxyARP_Range) String() string { return proto.CompactTextString(m) }
func (*ProxyARP_Range) ProtoMessage()    {}
func (*ProxyARP_Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb46f906a6f7c0e7, []int{0, 1}
}

func (m *ProxyARP_Range) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyARP_Range.Unmarshal(m, b)
}
func (m *ProxyARP_Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyARP_Range.Marshal(b, m, deterministic)
}
func (m *ProxyARP_Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyARP_Range.Merge(m, src)
}
func (m *ProxyARP_Range) XXX_Size() int {
	return xxx_messageInfo_ProxyARP_Range.Size(m)
}
func (m *ProxyARP_Range) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyARP_Range.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyARP_Range proto.InternalMessageInfo

func (m *ProxyARP_Range) GetFirstIpAddr() string {
	if m != nil {
		return m.FirstIpAddr
	}
	return ""
}

func (m *ProxyARP_Range) GetLastIpAddr() string {
	if m != nil {
		return m.LastIpAddr
	}
	return ""
}

type IPScanNeighbor struct {
	Mode                 IPScanNeighbor_Mode `protobuf:"varint,1,opt,name=mode,proto3,enum=ligato.vpp.l3.IPScanNeighbor_Mode" json:"mode,omitempty"`
	ScanInterval         uint32              `protobuf:"varint,2,opt,name=scan_interval,json=scanInterval,proto3" json:"scan_interval,omitempty"`
	MaxProcTime          uint32              `protobuf:"varint,3,opt,name=max_proc_time,json=maxProcTime,proto3" json:"max_proc_time,omitempty"`
	MaxUpdate            uint32              `protobuf:"varint,4,opt,name=max_update,json=maxUpdate,proto3" json:"max_update,omitempty"`
	ScanIntDelay         uint32              `protobuf:"varint,5,opt,name=scan_int_delay,json=scanIntDelay,proto3" json:"scan_int_delay,omitempty"`
	StaleThreshold       uint32              `protobuf:"varint,6,opt,name=stale_threshold,json=staleThreshold,proto3" json:"stale_threshold,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *IPScanNeighbor) Reset()         { *m = IPScanNeighbor{} }
func (m *IPScanNeighbor) String() string { return proto.CompactTextString(m) }
func (*IPScanNeighbor) ProtoMessage()    {}
func (*IPScanNeighbor) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb46f906a6f7c0e7, []int{1}
}

func (m *IPScanNeighbor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPScanNeighbor.Unmarshal(m, b)
}
func (m *IPScanNeighbor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPScanNeighbor.Marshal(b, m, deterministic)
}
func (m *IPScanNeighbor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPScanNeighbor.Merge(m, src)
}
func (m *IPScanNeighbor) XXX_Size() int {
	return xxx_messageInfo_IPScanNeighbor.Size(m)
}
func (m *IPScanNeighbor) XXX_DiscardUnknown() {
	xxx_messageInfo_IPScanNeighbor.DiscardUnknown(m)
}

var xxx_messageInfo_IPScanNeighbor proto.InternalMessageInfo

func (m *IPScanNeighbor) GetMode() IPScanNeighbor_Mode {
	if m != nil {
		return m.Mode
	}
	return IPScanNeighbor_DISABLED
}

func (m *IPScanNeighbor) GetScanInterval() uint32 {
	if m != nil {
		return m.ScanInterval
	}
	return 0
}

func (m *IPScanNeighbor) GetMaxProcTime() uint32 {
	if m != nil {
		return m.MaxProcTime
	}
	return 0
}

func (m *IPScanNeighbor) GetMaxUpdate() uint32 {
	if m != nil {
		return m.MaxUpdate
	}
	return 0
}

func (m *IPScanNeighbor) GetScanIntDelay() uint32 {
	if m != nil {
		return m.ScanIntDelay
	}
	return 0
}

func (m *IPScanNeighbor) GetStaleThreshold() uint32 {
	if m != nil {
		return m.StaleThreshold
	}
	return 0
}

type DHCPProxy struct {
	SourceIpAddress      string                  `protobuf:"bytes,1,opt,name=source_ip_address,json=sourceIpAddress,proto3" json:"source_ip_address,omitempty"`
	RxVrfId              uint32                  `protobuf:"varint,2,opt,name=rx_vrf_id,json=rxVrfId,proto3" json:"rx_vrf_id,omitempty"`
	Servers              []*DHCPProxy_DHCPServer `protobuf:"bytes,4,rep,name=servers,proto3" json:"servers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *DHCPProxy) Reset()         { *m = DHCPProxy{} }
func (m *DHCPProxy) String() string { return proto.CompactTextString(m) }
func (*DHCPProxy) ProtoMessage()    {}
func (*DHCPProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb46f906a6f7c0e7, []int{2}
}

func (m *DHCPProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DHCPProxy.Unmarshal(m, b)
}
func (m *DHCPProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DHCPProxy.Marshal(b, m, deterministic)
}
func (m *DHCPProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DHCPProxy.Merge(m, src)
}
func (m *DHCPProxy) XXX_Size() int {
	return xxx_messageInfo_DHCPProxy.Size(m)
}
func (m *DHCPProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_DHCPProxy.DiscardUnknown(m)
}

var xxx_messageInfo_DHCPProxy proto.InternalMessageInfo

func (m *DHCPProxy) GetSourceIpAddress() string {
	if m != nil {
		return m.SourceIpAddress
	}
	return ""
}

func (m *DHCPProxy) GetRxVrfId() uint32 {
	if m != nil {
		return m.RxVrfId
	}
	return 0
}

func (m *DHCPProxy) GetServers() []*DHCPProxy_DHCPServer {
	if m != nil {
		return m.Servers
	}
	return nil
}

type DHCPProxy_DHCPServer struct {
	VrfId                uint32   `protobuf:"varint,1,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	IpAddress            string   `protobuf:"bytes,2,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DHCPProxy_DHCPServer) Reset()         { *m = DHCPProxy_DHCPServer{} }
func (m *DHCPProxy_DHCPServer) String() string { return proto.CompactTextString(m) }
func (*DHCPProxy_DHCPServer) ProtoMessage()    {}
func (*DHCPProxy_DHCPServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb46f906a6f7c0e7, []int{2, 0}
}

func (m *DHCPProxy_DHCPServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DHCPProxy_DHCPServer.Unmarshal(m, b)
}
func (m *DHCPProxy_DHCPServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DHCPProxy_DHCPServer.Marshal(b, m, deterministic)
}
func (m *DHCPProxy_DHCPServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DHCPProxy_DHCPServer.Merge(m, src)
}
func (m *DHCPProxy_DHCPServer) XXX_Size() int {
	return xxx_messageInfo_DHCPProxy_DHCPServer.Size(m)
}
func (m *DHCPProxy_DHCPServer) XXX_DiscardUnknown() {
	xxx_messageInfo_DHCPProxy_DHCPServer.DiscardUnknown(m)
}

var xxx_messageInfo_DHCPProxy_DHCPServer proto.InternalMessageInfo

func (m *DHCPProxy_DHCPServer) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *DHCPProxy_DHCPServer) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func init() {
	proto.RegisterEnum("ligato.vpp.l3.IPScanNeighbor_Mode", IPScanNeighbor_Mode_name, IPScanNeighbor_Mode_value)
	proto.RegisterType((*ProxyARP)(nil), "ligato.vpp.l3.ProxyARP")
	proto.RegisterType((*ProxyARP_Interface)(nil), "ligato.vpp.l3.ProxyARP.Interface")
	proto.RegisterType((*ProxyARP_Range)(nil), "ligato.vpp.l3.ProxyARP.Range")
	proto.RegisterType((*IPScanNeighbor)(nil), "ligato.vpp.l3.IPScanNeighbor")
	proto.RegisterType((*DHCPProxy)(nil), "ligato.vpp.l3.DHCPProxy")
	proto.RegisterType((*DHCPProxy_DHCPServer)(nil), "ligato.vpp.l3.DHCPProxy.DHCPServer")
}

func init() {
	proto.RegisterFile("ligato/vpp/l3/l3.proto", fileDescriptor_eb46f906a6f7c0e7)
}

var fileDescriptor_eb46f906a6f7c0e7 = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xdf, 0x8a, 0x9b, 0x40,
	0x14, 0xc6, 0x6b, 0x62, 0xb2, 0xf1, 0xe4, 0xcf, 0xa6, 0x03, 0x2d, 0x12, 0x08, 0x4d, 0xdd, 0x42,
	0x97, 0x42, 0x15, 0xe2, 0x36, 0x37, 0xa5, 0x17, 0x49, 0x53, 0x58, 0xa1, 0xdb, 0x8a, 0x49, 0xf7,
	0xa2, 0x37, 0x32, 0xab, 0x93, 0x44, 0x50, 0x67, 0x98, 0x71, 0xc5, 0x7d, 0xae, 0xbe, 0x4b, 0x9f,
	0xa4, 0x0f, 0x50, 0x1c, 0x75, 0x9b, 0x2c, 0xf4, 0xee, 0xf0, 0xcd, 0xef, 0xf3, 0x3b, 0x9e, 0x99,
	0x03, 0x2f, 0xe3, 0x68, 0x8f, 0x33, 0x6a, 0xe5, 0x8c, 0x59, 0xb1, 0x6d, 0xc5, 0xb6, 0xc9, 0x38,
	0xcd, 0x28, 0x1a, 0x56, 0xba, 0x99, 0x33, 0x66, 0xc6, 0xb6, 0xf1, 0x47, 0x81, 0x9e, 0xcb, 0x69,
	0xf1, 0xb0, 0xf4, 0x5c, 0xb4, 0x04, 0x88, 0xd2, 0x8c, 0xf0, 0x1d, 0x0e, 0x88, 0xd0, 0x95, 0x59,
	0xfb, 0xb2, 0x3f, 0x7f, 0x6d, 0x9e, 0x18, 0xcc, 0x06, 0x36, 0x9d, 0x86, 0xf4, 0x8e, 0x4c, 0xe8,
	0x03, 0x74, 0x39, 0x4e, 0xf7, 0x44, 0xe8, 0x2d, 0x69, 0x9f, 0xfe, 0xcf, 0xee, 0x95, 0x94, 0x57,
	0xc3, 0x93, 0x57, 0xa0, 0x3d, 0x7e, 0x0f, 0x21, 0x50, 0x53, 0x9c, 0x10, 0x5d, 0x99, 0x29, 0x97,
	0x9a, 0x27, 0xeb, 0xc9, 0x0d, 0x74, 0xa4, 0x03, 0x19, 0x30, 0xdc, 0x45, 0x5c, 0x64, 0x7e, 0xc4,
	0x7c, 0x1c, 0x86, 0xbc, 0xa6, 0xfa, 0x52, 0x74, 0xd8, 0x32, 0x0c, 0x39, 0x9a, 0xc1, 0x20, 0xc6,
	0x47, 0x48, 0x4b, 0x22, 0x50, 0x6a, 0x15, 0x61, 0xfc, 0x6a, 0xc1, 0xc8, 0x71, 0x37, 0x01, 0x4e,
	0xbf, 0x91, 0x68, 0x7f, 0xb8, 0xa3, 0x1c, 0x2d, 0x40, 0x4d, 0x68, 0x58, 0xa5, 0x8e, 0xe6, 0xc6,
	0x93, 0xbe, 0x4f, 0x61, 0xf3, 0x86, 0x86, 0xc4, 0x93, 0x3c, 0xba, 0x80, 0xa1, 0x08, 0x70, 0xea,
	0xcb, 0x21, 0xe4, 0x38, 0x96, 0x69, 0x43, 0x6f, 0x50, 0x8a, 0x4e, 0xad, 0x95, 0x5d, 0x27, 0xb8,
	0xf0, 0x19, 0xa7, 0x81, 0x9f, 0x45, 0x09, 0xd1, 0xdb, 0x12, 0xea, 0x27, 0xb8, 0x70, 0x39, 0x0d,
	0xb6, 0x51, 0x42, 0xd0, 0x14, 0xa0, 0x64, 0xee, 0x59, 0x88, 0x33, 0xa2, 0xab, 0x12, 0xd0, 0x12,
	0x5c, 0xfc, 0x90, 0x02, 0x7a, 0x03, 0xa3, 0x26, 0xc7, 0x0f, 0x49, 0x8c, 0x1f, 0xf4, 0xce, 0x49,
	0xd0, 0xba, 0xd4, 0xd0, 0x5b, 0x38, 0x17, 0x19, 0x8e, 0x89, 0x9f, 0x1d, 0x38, 0x11, 0x07, 0x1a,
	0x87, 0x7a, 0x57, 0x62, 0x23, 0x29, 0x6f, 0x1b, 0xd5, 0x98, 0x83, 0x5a, 0xfe, 0x04, 0x1a, 0x40,
	0x6f, 0xed, 0x6c, 0x96, 0xab, 0xaf, 0x5f, 0xd6, 0xe3, 0x67, 0xa8, 0x07, 0xaa, 0xe3, 0xde, 0x5e,
	0x8d, 0x95, 0xba, 0x5a, 0x8c, 0x5b, 0x65, 0xb5, 0xfa, 0xbe, 0xbd, 0x1e, 0xb7, 0x8d, 0xdf, 0x0a,
	0x68, 0xeb, 0xeb, 0xcf, 0xae, 0xbc, 0x44, 0xf4, 0x0e, 0x9e, 0x0b, 0x7a, 0xcf, 0x03, 0xd2, 0xcc,
	0x99, 0x08, 0x51, 0xdf, 0xc6, 0x79, 0x75, 0x50, 0x0d, 0x9b, 0x08, 0x81, 0x26, 0xa0, 0xf1, 0xc2,
	0xcf, 0xf9, 0xce, 0x8f, 0xc2, 0x7a, 0x40, 0x67, 0xbc, 0xb8, 0xe5, 0x3b, 0x27, 0x44, 0x9f, 0xe0,
	0x4c, 0x10, 0x9e, 0x13, 0x2e, 0x74, 0x55, 0xbe, 0x99, 0x8b, 0x27, 0xb3, 0x7f, 0x8c, 0x94, 0xd5,
	0x46, 0xb2, 0x5e, 0xe3, 0x99, 0xac, 0x00, 0xfe, 0xc9, 0xe8, 0x05, 0x74, 0xeb, 0x14, 0x45, 0xa6,
	0x74, 0x72, 0x99, 0x31, 0x05, 0x38, 0x6a, 0xb2, 0x7a, 0x0f, 0x5a, 0xd4, 0xb4, 0xb7, 0x5a, 0xfc,
	0xbc, 0xda, 0xd3, 0x26, 0x35, 0x92, 0x4b, 0xf3, 0x1e, 0xef, 0x49, 0x9a, 0x59, 0xb9, 0x6d, 0xc9,
	0xbd, 0xb1, 0x4e, 0xd6, 0xe9, 0x63, 0xce, 0x98, 0x1f, 0xdb, 0x77, 0x5d, 0x79, 0x66, 0xff, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0xad, 0x6b, 0xa6, 0xb1, 0x6d, 0x03, 0x00, 0x00,
}
