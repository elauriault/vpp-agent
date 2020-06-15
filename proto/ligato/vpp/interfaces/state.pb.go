// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ligato/vpp/interfaces/state.proto

package vpp_interfaces

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

type InterfaceState_Status int32

const (
	InterfaceState_UNKNOWN_STATUS InterfaceState_Status = 0
	InterfaceState_UP             InterfaceState_Status = 1
	InterfaceState_DOWN           InterfaceState_Status = 2
	InterfaceState_DELETED        InterfaceState_Status = 3
)

var InterfaceState_Status_name = map[int32]string{
	0: "UNKNOWN_STATUS",
	1: "UP",
	2: "DOWN",
	3: "DELETED",
}

var InterfaceState_Status_value = map[string]int32{
	"UNKNOWN_STATUS": 0,
	"UP":             1,
	"DOWN":           2,
	"DELETED":        3,
}

func (x InterfaceState_Status) String() string {
	return proto.EnumName(InterfaceState_Status_name, int32(x))
}

func (InterfaceState_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_89357752b24c297d, []int{1, 0}
}

type InterfaceState_Duplex int32

const (
	InterfaceState_UNKNOWN_DUPLEX InterfaceState_Duplex = 0
	InterfaceState_HALF           InterfaceState_Duplex = 1
	InterfaceState_FULL           InterfaceState_Duplex = 2
)

var InterfaceState_Duplex_name = map[int32]string{
	0: "UNKNOWN_DUPLEX",
	1: "HALF",
	2: "FULL",
}

var InterfaceState_Duplex_value = map[string]int32{
	"UNKNOWN_DUPLEX": 0,
	"HALF":           1,
	"FULL":           2,
}

func (x InterfaceState_Duplex) String() string {
	return proto.EnumName(InterfaceState_Duplex_name, int32(x))
}

func (InterfaceState_Duplex) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_89357752b24c297d, []int{1, 1}
}

type InterfaceNotification_NotifType int32

const (
	InterfaceNotification_UNKNOWN  InterfaceNotification_NotifType = 0
	InterfaceNotification_UPDOWN   InterfaceNotification_NotifType = 1
	InterfaceNotification_COUNTERS InterfaceNotification_NotifType = 2
)

var InterfaceNotification_NotifType_name = map[int32]string{
	0: "UNKNOWN",
	1: "UPDOWN",
	2: "COUNTERS",
}

var InterfaceNotification_NotifType_value = map[string]int32{
	"UNKNOWN":  0,
	"UPDOWN":   1,
	"COUNTERS": 2,
}

func (x InterfaceNotification_NotifType) String() string {
	return proto.EnumName(InterfaceNotification_NotifType_name, int32(x))
}

func (InterfaceNotification_NotifType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_89357752b24c297d, []int{2, 0}
}

type InterfaceStats struct {
	Name                 string                          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Rx                   *InterfaceStats_CombinedCounter `protobuf:"bytes,2,opt,name=rx,proto3" json:"rx,omitempty"`
	Tx                   *InterfaceStats_CombinedCounter `protobuf:"bytes,3,opt,name=tx,proto3" json:"tx,omitempty"`
	RxUnicast            *InterfaceStats_CombinedCounter `protobuf:"bytes,4,opt,name=rx_unicast,json=rxUnicast,proto3" json:"rx_unicast,omitempty"`
	RxMulticast          *InterfaceStats_CombinedCounter `protobuf:"bytes,5,opt,name=rx_multicast,json=rxMulticast,proto3" json:"rx_multicast,omitempty"`
	RxBroadcast          *InterfaceStats_CombinedCounter `protobuf:"bytes,6,opt,name=rx_broadcast,json=rxBroadcast,proto3" json:"rx_broadcast,omitempty"`
	TxUnicast            *InterfaceStats_CombinedCounter `protobuf:"bytes,7,opt,name=tx_unicast,json=txUnicast,proto3" json:"tx_unicast,omitempty"`
	TxMulticast          *InterfaceStats_CombinedCounter `protobuf:"bytes,8,opt,name=tx_multicast,json=txMulticast,proto3" json:"tx_multicast,omitempty"`
	TxBroadcast          *InterfaceStats_CombinedCounter `protobuf:"bytes,9,opt,name=tx_broadcast,json=txBroadcast,proto3" json:"tx_broadcast,omitempty"`
	RxError              uint64                          `protobuf:"varint,12,opt,name=rx_error,json=rxError,proto3" json:"rx_error,omitempty"`
	TxError              uint64                          `protobuf:"varint,13,opt,name=tx_error,json=txError,proto3" json:"tx_error,omitempty"`
	RxNoBuf              uint64                          `protobuf:"varint,10,opt,name=rx_no_buf,json=rxNoBuf,proto3" json:"rx_no_buf,omitempty"`
	RxMiss               uint64                          `protobuf:"varint,11,opt,name=rx_miss,json=rxMiss,proto3" json:"rx_miss,omitempty"`
	Drops                uint64                          `protobuf:"varint,14,opt,name=drops,proto3" json:"drops,omitempty"`
	Punts                uint64                          `protobuf:"varint,15,opt,name=punts,proto3" json:"punts,omitempty"`
	Ip4                  uint64                          `protobuf:"varint,16,opt,name=ip4,proto3" json:"ip4,omitempty"`
	Ip6                  uint64                          `protobuf:"varint,17,opt,name=ip6,proto3" json:"ip6,omitempty"`
	Mpls                 uint64                          `protobuf:"varint,18,opt,name=mpls,proto3" json:"mpls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *InterfaceStats) Reset()         { *m = InterfaceStats{} }
func (m *InterfaceStats) String() string { return proto.CompactTextString(m) }
func (*InterfaceStats) ProtoMessage()    {}
func (*InterfaceStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_89357752b24c297d, []int{0}
}

func (m *InterfaceStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfaceStats.Unmarshal(m, b)
}
func (m *InterfaceStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfaceStats.Marshal(b, m, deterministic)
}
func (m *InterfaceStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfaceStats.Merge(m, src)
}
func (m *InterfaceStats) XXX_Size() int {
	return xxx_messageInfo_InterfaceStats.Size(m)
}
func (m *InterfaceStats) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfaceStats.DiscardUnknown(m)
}

var xxx_messageInfo_InterfaceStats proto.InternalMessageInfo

func (m *InterfaceStats) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InterfaceStats) GetRx() *InterfaceStats_CombinedCounter {
	if m != nil {
		return m.Rx
	}
	return nil
}

func (m *InterfaceStats) GetTx() *InterfaceStats_CombinedCounter {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *InterfaceStats) GetRxUnicast() *InterfaceStats_CombinedCounter {
	if m != nil {
		return m.RxUnicast
	}
	return nil
}

func (m *InterfaceStats) GetRxMulticast() *InterfaceStats_CombinedCounter {
	if m != nil {
		return m.RxMulticast
	}
	return nil
}

func (m *InterfaceStats) GetRxBroadcast() *InterfaceStats_CombinedCounter {
	if m != nil {
		return m.RxBroadcast
	}
	return nil
}

func (m *InterfaceStats) GetTxUnicast() *InterfaceStats_CombinedCounter {
	if m != nil {
		return m.TxUnicast
	}
	return nil
}

func (m *InterfaceStats) GetTxMulticast() *InterfaceStats_CombinedCounter {
	if m != nil {
		return m.TxMulticast
	}
	return nil
}

func (m *InterfaceStats) GetTxBroadcast() *InterfaceStats_CombinedCounter {
	if m != nil {
		return m.TxBroadcast
	}
	return nil
}

func (m *InterfaceStats) GetRxError() uint64 {
	if m != nil {
		return m.RxError
	}
	return 0
}

func (m *InterfaceStats) GetTxError() uint64 {
	if m != nil {
		return m.TxError
	}
	return 0
}

func (m *InterfaceStats) GetRxNoBuf() uint64 {
	if m != nil {
		return m.RxNoBuf
	}
	return 0
}

func (m *InterfaceStats) GetRxMiss() uint64 {
	if m != nil {
		return m.RxMiss
	}
	return 0
}

func (m *InterfaceStats) GetDrops() uint64 {
	if m != nil {
		return m.Drops
	}
	return 0
}

func (m *InterfaceStats) GetPunts() uint64 {
	if m != nil {
		return m.Punts
	}
	return 0
}

func (m *InterfaceStats) GetIp4() uint64 {
	if m != nil {
		return m.Ip4
	}
	return 0
}

func (m *InterfaceStats) GetIp6() uint64 {
	if m != nil {
		return m.Ip6
	}
	return 0
}

func (m *InterfaceStats) GetMpls() uint64 {
	if m != nil {
		return m.Mpls
	}
	return 0
}

type InterfaceStats_CombinedCounter struct {
	Packets              uint64   `protobuf:"varint,1,opt,name=packets,proto3" json:"packets,omitempty"`
	Bytes                uint64   `protobuf:"varint,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InterfaceStats_CombinedCounter) Reset()         { *m = InterfaceStats_CombinedCounter{} }
func (m *InterfaceStats_CombinedCounter) String() string { return proto.CompactTextString(m) }
func (*InterfaceStats_CombinedCounter) ProtoMessage()    {}
func (*InterfaceStats_CombinedCounter) Descriptor() ([]byte, []int) {
	return fileDescriptor_89357752b24c297d, []int{0, 0}
}

func (m *InterfaceStats_CombinedCounter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfaceStats_CombinedCounter.Unmarshal(m, b)
}
func (m *InterfaceStats_CombinedCounter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfaceStats_CombinedCounter.Marshal(b, m, deterministic)
}
func (m *InterfaceStats_CombinedCounter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfaceStats_CombinedCounter.Merge(m, src)
}
func (m *InterfaceStats_CombinedCounter) XXX_Size() int {
	return xxx_messageInfo_InterfaceStats_CombinedCounter.Size(m)
}
func (m *InterfaceStats_CombinedCounter) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfaceStats_CombinedCounter.DiscardUnknown(m)
}

var xxx_messageInfo_InterfaceStats_CombinedCounter proto.InternalMessageInfo

func (m *InterfaceStats_CombinedCounter) GetPackets() uint64 {
	if m != nil {
		return m.Packets
	}
	return 0
}

func (m *InterfaceStats_CombinedCounter) GetBytes() uint64 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

type InterfaceState struct {
	Name                 string                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	InternalName         string                     `protobuf:"bytes,2,opt,name=internal_name,json=internalName,proto3" json:"internal_name,omitempty"`
	Type                 Interface_Type             `protobuf:"varint,3,opt,name=type,proto3,enum=ligato.vpp.interfaces.Interface_Type" json:"type,omitempty"`
	IfIndex              uint32                     `protobuf:"varint,4,opt,name=if_index,json=ifIndex,proto3" json:"if_index,omitempty"`
	AdminStatus          InterfaceState_Status      `protobuf:"varint,5,opt,name=admin_status,json=adminStatus,proto3,enum=ligato.vpp.interfaces.InterfaceState_Status" json:"admin_status,omitempty"`
	OperStatus           InterfaceState_Status      `protobuf:"varint,6,opt,name=oper_status,json=operStatus,proto3,enum=ligato.vpp.interfaces.InterfaceState_Status" json:"oper_status,omitempty"`
	LastChange           int64                      `protobuf:"varint,7,opt,name=last_change,json=lastChange,proto3" json:"last_change,omitempty"`
	PhysAddress          string                     `protobuf:"bytes,8,opt,name=phys_address,json=physAddress,proto3" json:"phys_address,omitempty"`
	Speed                uint64                     `protobuf:"varint,9,opt,name=speed,proto3" json:"speed,omitempty"`
	Mtu                  uint32                     `protobuf:"varint,10,opt,name=mtu,proto3" json:"mtu,omitempty"`
	Duplex               InterfaceState_Duplex      `protobuf:"varint,11,opt,name=duplex,proto3,enum=ligato.vpp.interfaces.InterfaceState_Duplex" json:"duplex,omitempty"`
	Statistics           *InterfaceState_Statistics `protobuf:"bytes,100,opt,name=statistics,proto3" json:"statistics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *InterfaceState) Reset()         { *m = InterfaceState{} }
func (m *InterfaceState) String() string { return proto.CompactTextString(m) }
func (*InterfaceState) ProtoMessage()    {}
func (*InterfaceState) Descriptor() ([]byte, []int) {
	return fileDescriptor_89357752b24c297d, []int{1}
}

func (m *InterfaceState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfaceState.Unmarshal(m, b)
}
func (m *InterfaceState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfaceState.Marshal(b, m, deterministic)
}
func (m *InterfaceState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfaceState.Merge(m, src)
}
func (m *InterfaceState) XXX_Size() int {
	return xxx_messageInfo_InterfaceState.Size(m)
}
func (m *InterfaceState) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfaceState.DiscardUnknown(m)
}

var xxx_messageInfo_InterfaceState proto.InternalMessageInfo

func (m *InterfaceState) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InterfaceState) GetInternalName() string {
	if m != nil {
		return m.InternalName
	}
	return ""
}

func (m *InterfaceState) GetType() Interface_Type {
	if m != nil {
		return m.Type
	}
	return Interface_UNDEFINED_TYPE
}

func (m *InterfaceState) GetIfIndex() uint32 {
	if m != nil {
		return m.IfIndex
	}
	return 0
}

func (m *InterfaceState) GetAdminStatus() InterfaceState_Status {
	if m != nil {
		return m.AdminStatus
	}
	return InterfaceState_UNKNOWN_STATUS
}

func (m *InterfaceState) GetOperStatus() InterfaceState_Status {
	if m != nil {
		return m.OperStatus
	}
	return InterfaceState_UNKNOWN_STATUS
}

func (m *InterfaceState) GetLastChange() int64 {
	if m != nil {
		return m.LastChange
	}
	return 0
}

func (m *InterfaceState) GetPhysAddress() string {
	if m != nil {
		return m.PhysAddress
	}
	return ""
}

func (m *InterfaceState) GetSpeed() uint64 {
	if m != nil {
		return m.Speed
	}
	return 0
}

func (m *InterfaceState) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *InterfaceState) GetDuplex() InterfaceState_Duplex {
	if m != nil {
		return m.Duplex
	}
	return InterfaceState_UNKNOWN_DUPLEX
}

func (m *InterfaceState) GetStatistics() *InterfaceState_Statistics {
	if m != nil {
		return m.Statistics
	}
	return nil
}

type InterfaceState_Statistics struct {
	InPackets            uint64   `protobuf:"varint,1,opt,name=in_packets,json=inPackets,proto3" json:"in_packets,omitempty"`
	InBytes              uint64   `protobuf:"varint,2,opt,name=in_bytes,json=inBytes,proto3" json:"in_bytes,omitempty"`
	OutPackets           uint64   `protobuf:"varint,3,opt,name=out_packets,json=outPackets,proto3" json:"out_packets,omitempty"`
	OutBytes             uint64   `protobuf:"varint,4,opt,name=out_bytes,json=outBytes,proto3" json:"out_bytes,omitempty"`
	DropPackets          uint64   `protobuf:"varint,5,opt,name=drop_packets,json=dropPackets,proto3" json:"drop_packets,omitempty"`
	PuntPackets          uint64   `protobuf:"varint,6,opt,name=punt_packets,json=puntPackets,proto3" json:"punt_packets,omitempty"`
	Ipv4Packets          uint64   `protobuf:"varint,7,opt,name=ipv4_packets,json=ipv4Packets,proto3" json:"ipv4_packets,omitempty"`
	Ipv6Packets          uint64   `protobuf:"varint,8,opt,name=ipv6_packets,json=ipv6Packets,proto3" json:"ipv6_packets,omitempty"`
	InNobufPackets       uint64   `protobuf:"varint,9,opt,name=in_nobuf_packets,json=inNobufPackets,proto3" json:"in_nobuf_packets,omitempty"`
	InMissPackets        uint64   `protobuf:"varint,10,opt,name=in_miss_packets,json=inMissPackets,proto3" json:"in_miss_packets,omitempty"`
	InErrorPackets       uint64   `protobuf:"varint,11,opt,name=in_error_packets,json=inErrorPackets,proto3" json:"in_error_packets,omitempty"`
	OutErrorPackets      uint64   `protobuf:"varint,12,opt,name=out_error_packets,json=outErrorPackets,proto3" json:"out_error_packets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InterfaceState_Statistics) Reset()         { *m = InterfaceState_Statistics{} }
func (m *InterfaceState_Statistics) String() string { return proto.CompactTextString(m) }
func (*InterfaceState_Statistics) ProtoMessage()    {}
func (*InterfaceState_Statistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_89357752b24c297d, []int{1, 0}
}

func (m *InterfaceState_Statistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfaceState_Statistics.Unmarshal(m, b)
}
func (m *InterfaceState_Statistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfaceState_Statistics.Marshal(b, m, deterministic)
}
func (m *InterfaceState_Statistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfaceState_Statistics.Merge(m, src)
}
func (m *InterfaceState_Statistics) XXX_Size() int {
	return xxx_messageInfo_InterfaceState_Statistics.Size(m)
}
func (m *InterfaceState_Statistics) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfaceState_Statistics.DiscardUnknown(m)
}

var xxx_messageInfo_InterfaceState_Statistics proto.InternalMessageInfo

func (m *InterfaceState_Statistics) GetInPackets() uint64 {
	if m != nil {
		return m.InPackets
	}
	return 0
}

func (m *InterfaceState_Statistics) GetInBytes() uint64 {
	if m != nil {
		return m.InBytes
	}
	return 0
}

func (m *InterfaceState_Statistics) GetOutPackets() uint64 {
	if m != nil {
		return m.OutPackets
	}
	return 0
}

func (m *InterfaceState_Statistics) GetOutBytes() uint64 {
	if m != nil {
		return m.OutBytes
	}
	return 0
}

func (m *InterfaceState_Statistics) GetDropPackets() uint64 {
	if m != nil {
		return m.DropPackets
	}
	return 0
}

func (m *InterfaceState_Statistics) GetPuntPackets() uint64 {
	if m != nil {
		return m.PuntPackets
	}
	return 0
}

func (m *InterfaceState_Statistics) GetIpv4Packets() uint64 {
	if m != nil {
		return m.Ipv4Packets
	}
	return 0
}

func (m *InterfaceState_Statistics) GetIpv6Packets() uint64 {
	if m != nil {
		return m.Ipv6Packets
	}
	return 0
}

func (m *InterfaceState_Statistics) GetInNobufPackets() uint64 {
	if m != nil {
		return m.InNobufPackets
	}
	return 0
}

func (m *InterfaceState_Statistics) GetInMissPackets() uint64 {
	if m != nil {
		return m.InMissPackets
	}
	return 0
}

func (m *InterfaceState_Statistics) GetInErrorPackets() uint64 {
	if m != nil {
		return m.InErrorPackets
	}
	return 0
}

func (m *InterfaceState_Statistics) GetOutErrorPackets() uint64 {
	if m != nil {
		return m.OutErrorPackets
	}
	return 0
}

type InterfaceNotification struct {
	Type                 InterfaceNotification_NotifType `protobuf:"varint,1,opt,name=type,proto3,enum=ligato.vpp.interfaces.InterfaceNotification_NotifType" json:"type,omitempty"`
	State                *InterfaceState                 `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *InterfaceNotification) Reset()         { *m = InterfaceNotification{} }
func (m *InterfaceNotification) String() string { return proto.CompactTextString(m) }
func (*InterfaceNotification) ProtoMessage()    {}
func (*InterfaceNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_89357752b24c297d, []int{2}
}

func (m *InterfaceNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfaceNotification.Unmarshal(m, b)
}
func (m *InterfaceNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfaceNotification.Marshal(b, m, deterministic)
}
func (m *InterfaceNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfaceNotification.Merge(m, src)
}
func (m *InterfaceNotification) XXX_Size() int {
	return xxx_messageInfo_InterfaceNotification.Size(m)
}
func (m *InterfaceNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfaceNotification.DiscardUnknown(m)
}

var xxx_messageInfo_InterfaceNotification proto.InternalMessageInfo

func (m *InterfaceNotification) GetType() InterfaceNotification_NotifType {
	if m != nil {
		return m.Type
	}
	return InterfaceNotification_UNKNOWN
}

func (m *InterfaceNotification) GetState() *InterfaceState {
	if m != nil {
		return m.State
	}
	return nil
}

func init() {
	proto.RegisterEnum("ligato.vpp.interfaces.InterfaceState_Status", InterfaceState_Status_name, InterfaceState_Status_value)
	proto.RegisterEnum("ligato.vpp.interfaces.InterfaceState_Duplex", InterfaceState_Duplex_name, InterfaceState_Duplex_value)
	proto.RegisterEnum("ligato.vpp.interfaces.InterfaceNotification_NotifType", InterfaceNotification_NotifType_name, InterfaceNotification_NotifType_value)
	proto.RegisterType((*InterfaceStats)(nil), "ligato.vpp.interfaces.InterfaceStats")
	proto.RegisterType((*InterfaceStats_CombinedCounter)(nil), "ligato.vpp.interfaces.InterfaceStats.CombinedCounter")
	proto.RegisterType((*InterfaceState)(nil), "ligato.vpp.interfaces.InterfaceState")
	proto.RegisterType((*InterfaceState_Statistics)(nil), "ligato.vpp.interfaces.InterfaceState.Statistics")
	proto.RegisterType((*InterfaceNotification)(nil), "ligato.vpp.interfaces.InterfaceNotification")
}

func init() {
	proto.RegisterFile("ligato/vpp/interfaces/state.proto", fileDescriptor_89357752b24c297d)
}

var fileDescriptor_89357752b24c297d = []byte{
	// 961 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdd, 0x6e, 0xe2, 0x46,
	0x14, 0x0e, 0x3f, 0x31, 0x70, 0x0c, 0xc4, 0x3b, 0xea, 0xaa, 0x6e, 0xaa, 0xaa, 0x09, 0x55, 0xaa,
	0xa8, 0x6a, 0x61, 0x95, 0xdd, 0x22, 0x55, 0xb9, 0x4a, 0x02, 0x51, 0xb7, 0x25, 0x04, 0x19, 0xd0,
	0xae, 0x7a, 0x63, 0x19, 0x3c, 0x64, 0x47, 0x85, 0xb1, 0xe5, 0x19, 0x23, 0xe7, 0xcd, 0xfa, 0x20,
	0xbd, 0xe8, 0x4b, 0xf4, 0x1d, 0xaa, 0x39, 0x83, 0x8d, 0x59, 0x45, 0x4a, 0x54, 0xee, 0xe6, 0x7c,
	0xe7, 0xfb, 0x3e, 0x7b, 0xce, 0x9c, 0x33, 0x1a, 0x38, 0x5d, 0xb2, 0x07, 0x4f, 0x06, 0x9d, 0x75,
	0x18, 0x76, 0x18, 0x97, 0x34, 0x5a, 0x78, 0x73, 0x2a, 0x3a, 0x42, 0x7a, 0x92, 0xb6, 0xc3, 0x28,
	0x90, 0x01, 0x79, 0xad, 0x29, 0xed, 0x75, 0x18, 0xb6, 0xb7, 0x94, 0xe3, 0xb3, 0xa7, 0x95, 0xd9,
	0x52, 0xab, 0x5b, 0xff, 0x54, 0xa0, 0xf9, 0x3e, 0xc5, 0xc6, 0xd2, 0x93, 0x82, 0x10, 0x28, 0x73,
	0x6f, 0x45, 0xed, 0xc2, 0x49, 0xe1, 0xbc, 0xe6, 0xe0, 0x9a, 0xf4, 0xa1, 0x18, 0x25, 0x76, 0xf1,
	0xa4, 0x70, 0x6e, 0x5e, 0xfc, 0xdc, 0x7e, 0xf2, 0x8b, 0xed, 0x5d, 0x9b, 0xf6, 0x4d, 0xb0, 0x9a,
	0x31, 0x4e, 0xfd, 0x9b, 0x20, 0x56, 0x09, 0xa7, 0x18, 0x25, 0xca, 0x46, 0x26, 0x76, 0x69, 0x2f,
	0x1b, 0x99, 0x90, 0x09, 0x40, 0x94, 0xb8, 0x31, 0x67, 0x73, 0x4f, 0x48, 0xbb, 0xbc, 0x8f, 0x5d,
	0x2d, 0x4a, 0xa6, 0xda, 0x87, 0x7c, 0x84, 0x7a, 0x94, 0xb8, 0xab, 0x78, 0x29, 0xb5, 0xef, 0xe1,
	0x3e, 0xbe, 0x66, 0x94, 0xdc, 0xa5, 0x4e, 0x1b, 0xe7, 0x59, 0x14, 0x78, 0x3e, 0x3a, 0x1b, 0x7b,
	0x3a, 0x5f, 0xa7, 0x4e, 0xaa, 0x12, 0x72, 0x5b, 0x89, 0xca, 0x5e, 0x95, 0x90, 0xf9, 0x4a, 0xc8,
	0x7c, 0x25, 0xaa, 0x7b, 0xfd, 0xaf, 0xdc, 0xad, 0x84, 0xcc, 0x57, 0xa2, 0xb6, 0xa7, 0xf3, 0xb6,
	0x12, 0x5f, 0x41, 0x35, 0x4a, 0x5c, 0x1a, 0x45, 0x41, 0x64, 0xd7, 0x4f, 0x0a, 0xe7, 0x65, 0xa7,
	0x12, 0x25, 0x7d, 0x15, 0xaa, 0x94, 0x4c, 0x53, 0x0d, 0x9d, 0x92, 0x9b, 0xd4, 0x31, 0xd4, 0xa2,
	0xc4, 0xe5, 0x81, 0x3b, 0x8b, 0x17, 0x36, 0xa4, 0xb2, 0x61, 0x70, 0x1d, 0x2f, 0xc8, 0x97, 0x50,
	0x51, 0xfd, 0xc0, 0x84, 0xb0, 0x4d, 0xcc, 0x18, 0x51, 0x72, 0xc7, 0x84, 0x20, 0x5f, 0xc0, 0xa1,
	0x1f, 0x05, 0xa1, 0xb0, 0x9b, 0x08, 0xeb, 0x40, 0xa1, 0x61, 0xcc, 0xa5, 0xb0, 0x8f, 0x34, 0x8a,
	0x01, 0xb1, 0xa0, 0xc4, 0xc2, 0x77, 0xb6, 0x85, 0x98, 0x5a, 0x6a, 0xa4, 0x6b, 0xbf, 0x4a, 0x91,
	0xae, 0x1a, 0xb8, 0x55, 0xb8, 0x14, 0x36, 0x41, 0x08, 0xd7, 0xc7, 0x57, 0x70, 0xf4, 0xd9, 0x76,
	0x89, 0x0d, 0x95, 0xd0, 0x9b, 0xff, 0x49, 0xa5, 0xc0, 0xd1, 0x2c, 0x3b, 0x69, 0xa8, 0x3e, 0x3d,
	0x7b, 0x94, 0x54, 0xe0, 0x80, 0x96, 0x1d, 0x1d, 0xb4, 0xfe, 0xad, 0x7e, 0x36, 0xda, 0xf4, 0xc9,
	0xd1, 0xfe, 0x0e, 0x1a, 0x58, 0x72, 0xee, 0x2d, 0x5d, 0x4c, 0x16, 0x31, 0x59, 0x4f, 0xc1, 0xa1,
	0x22, 0xfd, 0x02, 0x65, 0xf9, 0x18, 0x52, 0x1c, 0xdd, 0xe6, 0xc5, 0xd9, 0x73, 0xe7, 0xd5, 0x9e,
	0x3c, 0x86, 0xd4, 0x41, 0x89, 0xaa, 0x3e, 0x5b, 0xb8, 0x8c, 0xfb, 0x34, 0xc1, 0x51, 0x6d, 0x38,
	0x15, 0xb6, 0x78, 0xaf, 0x42, 0x72, 0x0f, 0x75, 0xcf, 0x5f, 0x31, 0xee, 0xaa, 0xfb, 0x2c, 0x16,
	0x38, 0x71, 0xcd, 0x8b, 0x1f, 0x5f, 0xd2, 0x0d, 0xb4, 0x3d, 0x46, 0x8d, 0x63, 0xa2, 0x83, 0x0e,
	0xc8, 0x1d, 0x98, 0x41, 0x48, 0xa3, 0xd4, 0xcf, 0xf8, 0x1f, 0x7e, 0xa0, 0x0c, 0x36, 0x76, 0xdf,
	0x82, 0xb9, 0xf4, 0x84, 0x74, 0xe7, 0x9f, 0x3c, 0xfe, 0x40, 0x71, 0xbc, 0x4a, 0x0e, 0x28, 0xe8,
	0x06, 0x11, 0x72, 0x0a, 0xf5, 0xf0, 0xd3, 0xa3, 0x70, 0x3d, 0xdf, 0x8f, 0xa8, 0x10, 0x38, 0x28,
	0x35, 0xc7, 0x54, 0xd8, 0x95, 0x86, 0xd4, 0xd9, 0x88, 0x90, 0x52, 0x1f, 0x5b, 0xbd, 0xec, 0xe8,
	0x40, 0x35, 0xc1, 0x4a, 0xc6, 0xd8, 0x71, 0x0d, 0x47, 0x2d, 0x49, 0x0f, 0x0c, 0x3f, 0x0e, 0x97,
	0x34, 0xc1, 0x66, 0x7b, 0xf1, 0x5f, 0xf7, 0x50, 0xe3, 0x6c, 0xb4, 0x64, 0x04, 0xa0, 0xf6, 0xce,
	0x84, 0x64, 0x73, 0x61, 0xfb, 0x38, 0x5d, 0x6f, 0x5e, 0xbe, 0x7f, 0xad, 0x73, 0x72, 0x1e, 0xc7,
	0x7f, 0x95, 0x00, 0xb6, 0x29, 0xf2, 0x0d, 0x00, 0xe3, 0xee, 0x6e, 0x1f, 0xd6, 0x18, 0x1f, 0x6d,
	0x3a, 0x51, 0x1d, 0x36, 0x77, 0xf3, 0xcd, 0x58, 0x61, 0xfc, 0x5a, 0x85, 0xaa, 0x98, 0x41, 0x2c,
	0x33, 0x69, 0x09, 0xb3, 0x10, 0xc4, 0x32, 0xd5, 0x7e, 0x0d, 0x35, 0x45, 0xd0, 0xe2, 0x32, 0xa6,
	0xab, 0x41, 0x2c, 0xb5, 0xfa, 0x14, 0xea, 0x6a, 0xcc, 0x32, 0xf9, 0x21, 0xe6, 0x4d, 0x85, 0xa5,
	0x7a, 0x75, 0x18, 0x31, 0xdf, 0x7e, 0xc1, 0xd0, 0x14, 0x85, 0xe5, 0x28, 0x2c, 0x5c, 0xbf, 0xcb,
	0x28, 0x15, 0x4d, 0x51, 0xd8, 0x2e, 0xa5, 0x9b, 0x51, 0xaa, 0x19, 0xa5, 0x9b, 0x52, 0xce, 0xc1,
	0x62, 0xdc, 0xe5, 0xc1, 0x2c, 0x5e, 0x64, 0x34, 0x7d, 0xba, 0x4d, 0xc6, 0x87, 0x0a, 0x4e, 0x99,
	0xdf, 0xc3, 0x11, 0xe3, 0x78, 0x85, 0x64, 0x44, 0x7d, 0xc9, 0x34, 0x18, 0x57, 0x57, 0xc9, 0xae,
	0x23, 0xde, 0x50, 0x19, 0xd1, 0x4c, 0x1d, 0xf1, 0xa6, 0x4a, 0x99, 0x3f, 0xc0, 0x2b, 0x55, 0xa4,
	0x5d, 0xaa, 0xbe, 0xef, 0x8e, 0x82, 0x58, 0xe6, 0xb9, 0xad, 0x4b, 0x30, 0x36, 0x8d, 0x4c, 0xa0,
	0x39, 0x1d, 0xfe, 0x3e, 0xbc, 0xff, 0x30, 0x74, 0xc7, 0x93, 0xab, 0xc9, 0x74, 0x6c, 0x1d, 0x10,
	0x03, 0x8a, 0xd3, 0x91, 0x55, 0x20, 0x55, 0x28, 0xf7, 0xee, 0x3f, 0x0c, 0xad, 0x22, 0x31, 0xa1,
	0xd2, 0xeb, 0x0f, 0xfa, 0x93, 0x7e, 0xcf, 0x2a, 0xb5, 0xde, 0x80, 0xa1, 0x7b, 0x2b, 0x2f, 0xee,
	0x4d, 0x47, 0x83, 0xfe, 0x47, 0xeb, 0x40, 0x89, 0x7e, 0xbd, 0x1a, 0xdc, 0x6a, 0xf9, 0xed, 0x74,
	0x30, 0xb0, 0x8a, 0xad, 0xbf, 0x0b, 0xf0, 0x3a, 0xeb, 0xa9, 0x61, 0x20, 0xd9, 0x82, 0xcd, 0x3d,
	0xc9, 0x02, 0x4e, 0x7e, 0xdb, 0xdc, 0x1e, 0x05, 0xec, 0xec, 0xee, 0x73, 0xfd, 0x98, 0xd7, 0xb6,
	0x31, 0xc8, 0x5d, 0x27, 0x97, 0x70, 0x88, 0xaf, 0x9f, 0xcd, 0x63, 0xe4, 0xec, 0x45, 0xcd, 0xed,
	0x68, 0x4d, 0xeb, 0x02, 0x6a, 0x99, 0x9f, 0xda, 0xee, 0x66, 0x5f, 0xd6, 0x01, 0x01, 0x30, 0xa6,
	0x23, 0xac, 0x43, 0x81, 0xd4, 0xa1, 0x7a, 0x73, 0x3f, 0x1d, 0x4e, 0xfa, 0xce, 0xd8, 0x2a, 0x5e,
	0xdf, 0xfe, 0xd1, 0x7b, 0x08, 0xd2, 0xaf, 0x30, 0x7c, 0x50, 0xfd, 0xe4, 0x3d, 0x50, 0x2e, 0x3b,
	0xeb, 0xb7, 0x1d, 0x7c, 0x44, 0x75, 0x9e, 0x7c, 0x6a, 0x5d, 0xae, 0xc3, 0xd0, 0xdd, 0x86, 0x33,
	0x03, 0xb9, 0x6f, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xc8, 0xdd, 0x20, 0xd3, 0x09, 0x00,
	0x00,
}
