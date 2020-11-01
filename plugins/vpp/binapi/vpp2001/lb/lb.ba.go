// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /usr/share/vpp/api/plugins/lb.api.json

/*
Package lb is a generated VPP binary API for 'lb' module.

It consists of:
	 15 enums
	  6 aliases
	  7 types
	  1 union
	 16 messages
	  8 services
*/
package lb

import (
	bytes "bytes"
	context "context"
	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"
	io "io"
	strconv "strconv"

	interface_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2001/interface_types"
	ip_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2001/ip_types"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "lb"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x76c36b4c
)

type AddressFamily = ip_types.AddressFamily

type IfStatusFlags = interface_types.IfStatusFlags

type IfType = interface_types.IfType

type IPDscp = ip_types.IPDscp

type IPEcn = ip_types.IPEcn

type IPProto = ip_types.IPProto

// LbEncapType represents VPP binary API enum 'lb_encap_type'.
type LbEncapType uint32

const (
	LB_API_ENCAP_TYPE_GRE4  LbEncapType = 0
	LB_API_ENCAP_TYPE_GRE6  LbEncapType = 1
	LB_API_ENCAP_TYPE_L3DSR LbEncapType = 2
	LB_API_ENCAP_TYPE_NAT4  LbEncapType = 3
	LB_API_ENCAP_TYPE_NAT6  LbEncapType = 4
	LB_API_ENCAP_N_TYPES    LbEncapType = 5
)

var LbEncapType_name = map[uint32]string{
	0: "LB_API_ENCAP_TYPE_GRE4",
	1: "LB_API_ENCAP_TYPE_GRE6",
	2: "LB_API_ENCAP_TYPE_L3DSR",
	3: "LB_API_ENCAP_TYPE_NAT4",
	4: "LB_API_ENCAP_TYPE_NAT6",
	5: "LB_API_ENCAP_N_TYPES",
}

var LbEncapType_value = map[string]uint32{
	"LB_API_ENCAP_TYPE_GRE4":  0,
	"LB_API_ENCAP_TYPE_GRE6":  1,
	"LB_API_ENCAP_TYPE_L3DSR": 2,
	"LB_API_ENCAP_TYPE_NAT4":  3,
	"LB_API_ENCAP_TYPE_NAT6":  4,
	"LB_API_ENCAP_N_TYPES":    5,
}

func (x LbEncapType) String() string {
	s, ok := LbEncapType_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// LbLkpTypeT represents VPP binary API enum 'lb_lkp_type_t'.
type LbLkpTypeT uint32

const (
	LB_API_LKP_SAME_IP_PORT LbLkpTypeT = 0
	LB_API_LKP_DIFF_IP_PORT LbLkpTypeT = 1
	LB_API_LKP_ALL_PORT_IP  LbLkpTypeT = 2
	LB_API_LKP_N_TYPES      LbLkpTypeT = 3
)

var LbLkpTypeT_name = map[uint32]string{
	0: "LB_API_LKP_SAME_IP_PORT",
	1: "LB_API_LKP_DIFF_IP_PORT",
	2: "LB_API_LKP_ALL_PORT_IP",
	3: "LB_API_LKP_N_TYPES",
}

var LbLkpTypeT_value = map[string]uint32{
	"LB_API_LKP_SAME_IP_PORT": 0,
	"LB_API_LKP_DIFF_IP_PORT": 1,
	"LB_API_LKP_ALL_PORT_IP":  2,
	"LB_API_LKP_N_TYPES":      3,
}

func (x LbLkpTypeT) String() string {
	s, ok := LbLkpTypeT_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// LbNatProtocol represents VPP binary API enum 'lb_nat_protocol'.
type LbNatProtocol uint32

const (
	LB_API_NAT_PROTOCOL_UDP LbNatProtocol = 6
	LB_API_NAT_PROTOCOL_TCP LbNatProtocol = 23
	LB_API_NAT_PROTOCOL_ANY LbNatProtocol = 4.294967295e+09
)

var LbNatProtocol_name = map[uint32]string{
	6:               "LB_API_NAT_PROTOCOL_UDP",
	23:              "LB_API_NAT_PROTOCOL_TCP",
	4.294967295e+09: "LB_API_NAT_PROTOCOL_ANY",
}

var LbNatProtocol_value = map[string]uint32{
	"LB_API_NAT_PROTOCOL_UDP": 6,
	"LB_API_NAT_PROTOCOL_TCP": 23,
	"LB_API_NAT_PROTOCOL_ANY": 4.294967295e+09,
}

func (x LbNatProtocol) String() string {
	s, ok := LbNatProtocol_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// LbSrvType represents VPP binary API enum 'lb_srv_type'.
type LbSrvType uint32

const (
	LB_API_SRV_TYPE_CLUSTERIP LbSrvType = 0
	LB_API_SRV_TYPE_NODEPORT  LbSrvType = 1
	LB_API_SRV_N_TYPES        LbSrvType = 2
)

var LbSrvType_name = map[uint32]string{
	0: "LB_API_SRV_TYPE_CLUSTERIP",
	1: "LB_API_SRV_TYPE_NODEPORT",
	2: "LB_API_SRV_N_TYPES",
}

var LbSrvType_value = map[string]uint32{
	"LB_API_SRV_TYPE_CLUSTERIP": 0,
	"LB_API_SRV_TYPE_NODEPORT":  1,
	"LB_API_SRV_N_TYPES":        2,
}

func (x LbSrvType) String() string {
	s, ok := LbSrvType_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// LbVipType represents VPP binary API enum 'lb_vip_type'.
type LbVipType uint32

const (
	LB_API_VIP_TYPE_IP6_GRE6  LbVipType = 0
	LB_API_VIP_TYPE_IP6_GRE4  LbVipType = 1
	LB_API_VIP_TYPE_IP4_GRE6  LbVipType = 2
	LB_API_VIP_TYPE_IP4_GRE4  LbVipType = 3
	LB_API_VIP_TYPE_IP4_L3DSR LbVipType = 4
	LB_API_VIP_TYPE_IP4_NAT4  LbVipType = 5
	LB_API_VIP_TYPE_IP6_NAT6  LbVipType = 6
	LB_API_VIP_N_TYPES        LbVipType = 7
)

var LbVipType_name = map[uint32]string{
	0: "LB_API_VIP_TYPE_IP6_GRE6",
	1: "LB_API_VIP_TYPE_IP6_GRE4",
	2: "LB_API_VIP_TYPE_IP4_GRE6",
	3: "LB_API_VIP_TYPE_IP4_GRE4",
	4: "LB_API_VIP_TYPE_IP4_L3DSR",
	5: "LB_API_VIP_TYPE_IP4_NAT4",
	6: "LB_API_VIP_TYPE_IP6_NAT6",
	7: "LB_API_VIP_N_TYPES",
}

var LbVipType_value = map[string]uint32{
	"LB_API_VIP_TYPE_IP6_GRE6":  0,
	"LB_API_VIP_TYPE_IP6_GRE4":  1,
	"LB_API_VIP_TYPE_IP4_GRE6":  2,
	"LB_API_VIP_TYPE_IP4_GRE4":  3,
	"LB_API_VIP_TYPE_IP4_L3DSR": 4,
	"LB_API_VIP_TYPE_IP4_NAT4":  5,
	"LB_API_VIP_TYPE_IP6_NAT6":  6,
	"LB_API_VIP_N_TYPES":        7,
}

func (x LbVipType) String() string {
	s, ok := LbVipType_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

type LinkDuplex = interface_types.LinkDuplex

type MtuProto = interface_types.MtuProto

type RxMode = interface_types.RxMode

type SubIfFlags = interface_types.SubIfFlags

type AddressWithPrefix = ip_types.AddressWithPrefix

type InterfaceIndex = interface_types.InterfaceIndex

type IP4Address = ip_types.IP4Address

type IP4AddressWithPrefix = ip_types.IP4AddressWithPrefix

type IP6Address = ip_types.IP6Address

type IP6AddressWithPrefix = ip_types.IP6AddressWithPrefix

type Address = ip_types.Address

type IP4Prefix = ip_types.IP4Prefix

type IP6Prefix = ip_types.IP6Prefix

// LbVip represents VPP binary API type 'lb_vip'.
type LbVip struct {
	Pfx      AddressWithPrefix
	Protocol IPProto
	Port     uint16
}

func (*LbVip) GetTypeName() string { return "lb_vip" }

type Mprefix = ip_types.Mprefix

type Prefix = ip_types.Prefix

type PrefixMatcher = ip_types.PrefixMatcher

type AddressUnion = ip_types.AddressUnion

// LbAddDelAs represents VPP binary API message 'lb_add_del_as'.
type LbAddDelAs struct {
	Pfx       AddressWithPrefix
	Protocol  uint8
	Port      uint16
	AsAddress Address
	IsDel     bool
	IsFlush   bool
}

func (m *LbAddDelAs) Reset()                        { *m = LbAddDelAs{} }
func (*LbAddDelAs) GetMessageName() string          { return "lb_add_del_as" }
func (*LbAddDelAs) GetCrcString() string            { return "78628987" }
func (*LbAddDelAs) GetMessageType() api.MessageType { return api.RequestMessage }

// LbAddDelAsReply represents VPP binary API message 'lb_add_del_as_reply'.
type LbAddDelAsReply struct {
	Retval int32
}

func (m *LbAddDelAsReply) Reset()                        { *m = LbAddDelAsReply{} }
func (*LbAddDelAsReply) GetMessageName() string          { return "lb_add_del_as_reply" }
func (*LbAddDelAsReply) GetCrcString() string            { return "e8d4e804" }
func (*LbAddDelAsReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// LbAddDelIntfNat4 represents VPP binary API message 'lb_add_del_intf_nat4'.
type LbAddDelIntfNat4 struct {
	IsAdd     bool
	SwIfIndex InterfaceIndex
}

func (m *LbAddDelIntfNat4) Reset()                        { *m = LbAddDelIntfNat4{} }
func (*LbAddDelIntfNat4) GetMessageName() string          { return "lb_add_del_intf_nat4" }
func (*LbAddDelIntfNat4) GetCrcString() string            { return "47d6e753" }
func (*LbAddDelIntfNat4) GetMessageType() api.MessageType { return api.RequestMessage }

// LbAddDelIntfNat4Reply represents VPP binary API message 'lb_add_del_intf_nat4_reply'.
type LbAddDelIntfNat4Reply struct {
	Retval int32
}

func (m *LbAddDelIntfNat4Reply) Reset()                        { *m = LbAddDelIntfNat4Reply{} }
func (*LbAddDelIntfNat4Reply) GetMessageName() string          { return "lb_add_del_intf_nat4_reply" }
func (*LbAddDelIntfNat4Reply) GetCrcString() string            { return "e8d4e804" }
func (*LbAddDelIntfNat4Reply) GetMessageType() api.MessageType { return api.ReplyMessage }

// LbAddDelIntfNat6 represents VPP binary API message 'lb_add_del_intf_nat6'.
type LbAddDelIntfNat6 struct {
	IsAdd     bool
	SwIfIndex InterfaceIndex
}

func (m *LbAddDelIntfNat6) Reset()                        { *m = LbAddDelIntfNat6{} }
func (*LbAddDelIntfNat6) GetMessageName() string          { return "lb_add_del_intf_nat6" }
func (*LbAddDelIntfNat6) GetCrcString() string            { return "47d6e753" }
func (*LbAddDelIntfNat6) GetMessageType() api.MessageType { return api.RequestMessage }

// LbAddDelIntfNat6Reply represents VPP binary API message 'lb_add_del_intf_nat6_reply'.
type LbAddDelIntfNat6Reply struct {
	Retval int32
}

func (m *LbAddDelIntfNat6Reply) Reset()                        { *m = LbAddDelIntfNat6Reply{} }
func (*LbAddDelIntfNat6Reply) GetMessageName() string          { return "lb_add_del_intf_nat6_reply" }
func (*LbAddDelIntfNat6Reply) GetCrcString() string            { return "e8d4e804" }
func (*LbAddDelIntfNat6Reply) GetMessageType() api.MessageType { return api.ReplyMessage }

// LbAddDelVip represents VPP binary API message 'lb_add_del_vip'.
type LbAddDelVip struct {
	Pfx                 AddressWithPrefix
	Protocol            uint8
	Port                uint16
	Encap               LbEncapType
	Dscp                uint8
	Type                LbSrvType
	TargetPort          uint16
	NodePort            uint16
	NewFlowsTableLength uint32
	IsDel               bool
}

func (m *LbAddDelVip) Reset()                        { *m = LbAddDelVip{} }
func (*LbAddDelVip) GetMessageName() string          { return "lb_add_del_vip" }
func (*LbAddDelVip) GetCrcString() string            { return "d15b7ddc" }
func (*LbAddDelVip) GetMessageType() api.MessageType { return api.RequestMessage }

// LbAddDelVipReply represents VPP binary API message 'lb_add_del_vip_reply'.
type LbAddDelVipReply struct {
	Retval int32
}

func (m *LbAddDelVipReply) Reset()                        { *m = LbAddDelVipReply{} }
func (*LbAddDelVipReply) GetMessageName() string          { return "lb_add_del_vip_reply" }
func (*LbAddDelVipReply) GetCrcString() string            { return "e8d4e804" }
func (*LbAddDelVipReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// LbAsDetails represents VPP binary API message 'lb_as_details'.
type LbAsDetails struct {
	Vip        LbVip
	AppSrv     Address
	Flags      uint8
	InUseSince uint32
}

func (m *LbAsDetails) Reset()                        { *m = LbAsDetails{} }
func (*LbAsDetails) GetMessageName() string          { return "lb_as_details" }
func (*LbAsDetails) GetCrcString() string            { return "9c39f60e" }
func (*LbAsDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// LbAsDump represents VPP binary API message 'lb_as_dump'.
type LbAsDump struct {
	Pfx      AddressWithPrefix
	Protocol uint8
	Port     uint16
}

func (m *LbAsDump) Reset()                        { *m = LbAsDump{} }
func (*LbAsDump) GetMessageName() string          { return "lb_as_dump" }
func (*LbAsDump) GetCrcString() string            { return "1063f819" }
func (*LbAsDump) GetMessageType() api.MessageType { return api.RequestMessage }

// LbConf represents VPP binary API message 'lb_conf'.
type LbConf struct {
	IP4SrcAddress        IP4Address
	IP6SrcAddress        IP6Address
	StickyBucketsPerCore uint32
	FlowTimeout          uint32
}

func (m *LbConf) Reset()                        { *m = LbConf{} }
func (*LbConf) GetMessageName() string          { return "lb_conf" }
func (*LbConf) GetCrcString() string            { return "22ddb739" }
func (*LbConf) GetMessageType() api.MessageType { return api.RequestMessage }

// LbConfReply represents VPP binary API message 'lb_conf_reply'.
type LbConfReply struct {
	Retval int32
}

func (m *LbConfReply) Reset()                        { *m = LbConfReply{} }
func (*LbConfReply) GetMessageName() string          { return "lb_conf_reply" }
func (*LbConfReply) GetCrcString() string            { return "e8d4e804" }
func (*LbConfReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// LbFlushVip represents VPP binary API message 'lb_flush_vip'.
type LbFlushVip struct {
	Pfx      AddressWithPrefix
	Protocol uint8
	Port     uint16
}

func (m *LbFlushVip) Reset()                        { *m = LbFlushVip{} }
func (*LbFlushVip) GetMessageName() string          { return "lb_flush_vip" }
func (*LbFlushVip) GetCrcString() string            { return "1063f819" }
func (*LbFlushVip) GetMessageType() api.MessageType { return api.RequestMessage }

// LbFlushVipReply represents VPP binary API message 'lb_flush_vip_reply'.
type LbFlushVipReply struct {
	Retval int32
}

func (m *LbFlushVipReply) Reset()                        { *m = LbFlushVipReply{} }
func (*LbFlushVipReply) GetMessageName() string          { return "lb_flush_vip_reply" }
func (*LbFlushVipReply) GetCrcString() string            { return "e8d4e804" }
func (*LbFlushVipReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// LbVipDetails represents VPP binary API message 'lb_vip_details'.
type LbVipDetails struct {
	Vip             LbVip
	Encap           LbEncapType
	Dscp            IPDscp
	SrvType         LbSrvType
	TargetPort      uint16
	FlowTableLength uint16
}

func (m *LbVipDetails) Reset()                        { *m = LbVipDetails{} }
func (*LbVipDetails) GetMessageName() string          { return "lb_vip_details" }
func (*LbVipDetails) GetCrcString() string            { return "08f39bed" }
func (*LbVipDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// LbVipDump represents VPP binary API message 'lb_vip_dump'.
type LbVipDump struct {
	Pfx        AddressWithPrefix
	PfxMatcher PrefixMatcher
	Protocol   uint8
	Port       uint16
}

func (m *LbVipDump) Reset()                        { *m = LbVipDump{} }
func (*LbVipDump) GetMessageName() string          { return "lb_vip_dump" }
func (*LbVipDump) GetCrcString() string            { return "c7bcb124" }
func (*LbVipDump) GetMessageType() api.MessageType { return api.RequestMessage }

func init() {
	api.RegisterMessage((*LbAddDelAs)(nil), "lb.LbAddDelAs")
	api.RegisterMessage((*LbAddDelAsReply)(nil), "lb.LbAddDelAsReply")
	api.RegisterMessage((*LbAddDelIntfNat4)(nil), "lb.LbAddDelIntfNat4")
	api.RegisterMessage((*LbAddDelIntfNat4Reply)(nil), "lb.LbAddDelIntfNat4Reply")
	api.RegisterMessage((*LbAddDelIntfNat6)(nil), "lb.LbAddDelIntfNat6")
	api.RegisterMessage((*LbAddDelIntfNat6Reply)(nil), "lb.LbAddDelIntfNat6Reply")
	api.RegisterMessage((*LbAddDelVip)(nil), "lb.LbAddDelVip")
	api.RegisterMessage((*LbAddDelVipReply)(nil), "lb.LbAddDelVipReply")
	api.RegisterMessage((*LbAsDetails)(nil), "lb.LbAsDetails")
	api.RegisterMessage((*LbAsDump)(nil), "lb.LbAsDump")
	api.RegisterMessage((*LbConf)(nil), "lb.LbConf")
	api.RegisterMessage((*LbConfReply)(nil), "lb.LbConfReply")
	api.RegisterMessage((*LbFlushVip)(nil), "lb.LbFlushVip")
	api.RegisterMessage((*LbFlushVipReply)(nil), "lb.LbFlushVipReply")
	api.RegisterMessage((*LbVipDetails)(nil), "lb.LbVipDetails")
	api.RegisterMessage((*LbVipDump)(nil), "lb.LbVipDump")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*LbAddDelAs)(nil),
		(*LbAddDelAsReply)(nil),
		(*LbAddDelIntfNat4)(nil),
		(*LbAddDelIntfNat4Reply)(nil),
		(*LbAddDelIntfNat6)(nil),
		(*LbAddDelIntfNat6Reply)(nil),
		(*LbAddDelVip)(nil),
		(*LbAddDelVipReply)(nil),
		(*LbAsDetails)(nil),
		(*LbAsDump)(nil),
		(*LbConf)(nil),
		(*LbConfReply)(nil),
		(*LbFlushVip)(nil),
		(*LbFlushVipReply)(nil),
		(*LbVipDetails)(nil),
		(*LbVipDump)(nil),
	}
}

// RPCService represents RPC service API for lb module.
type RPCService interface {
	DumpLbAs(ctx context.Context, in *LbAsDump) (RPCService_DumpLbAsClient, error)
	DumpLbVip(ctx context.Context, in *LbVipDump) (RPCService_DumpLbVipClient, error)
	LbAddDelAs(ctx context.Context, in *LbAddDelAs) (*LbAddDelAsReply, error)
	LbAddDelIntfNat4(ctx context.Context, in *LbAddDelIntfNat4) (*LbAddDelIntfNat4Reply, error)
	LbAddDelIntfNat6(ctx context.Context, in *LbAddDelIntfNat6) (*LbAddDelIntfNat6Reply, error)
	LbAddDelVip(ctx context.Context, in *LbAddDelVip) (*LbAddDelVipReply, error)
	LbConf(ctx context.Context, in *LbConf) (*LbConfReply, error)
	LbFlushVip(ctx context.Context, in *LbFlushVip) (*LbFlushVipReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpLbAs(ctx context.Context, in *LbAsDump) (RPCService_DumpLbAsClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpLbAsClient{stream}
	return x, nil
}

type RPCService_DumpLbAsClient interface {
	Recv() (*LbAsDetails, error)
}

type serviceClient_DumpLbAsClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpLbAsClient) Recv() (*LbAsDetails, error) {
	m := new(LbAsDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpLbVip(ctx context.Context, in *LbVipDump) (RPCService_DumpLbVipClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpLbVipClient{stream}
	return x, nil
}

type RPCService_DumpLbVipClient interface {
	Recv() (*LbVipDetails, error)
}

type serviceClient_DumpLbVipClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpLbVipClient) Recv() (*LbVipDetails, error) {
	m := new(LbVipDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) LbAddDelAs(ctx context.Context, in *LbAddDelAs) (*LbAddDelAsReply, error) {
	out := new(LbAddDelAsReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) LbAddDelIntfNat4(ctx context.Context, in *LbAddDelIntfNat4) (*LbAddDelIntfNat4Reply, error) {
	out := new(LbAddDelIntfNat4Reply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) LbAddDelIntfNat6(ctx context.Context, in *LbAddDelIntfNat6) (*LbAddDelIntfNat6Reply, error) {
	out := new(LbAddDelIntfNat6Reply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) LbAddDelVip(ctx context.Context, in *LbAddDelVip) (*LbAddDelVipReply, error) {
	out := new(LbAddDelVipReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) LbConf(ctx context.Context, in *LbConf) (*LbConfReply, error) {
	out := new(LbConfReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) LbFlushVip(ctx context.Context, in *LbFlushVip) (*LbFlushVipReply, error) {
	out := new(LbFlushVipReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion1 // please upgrade the GoVPP api package

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = bytes.NewBuffer
var _ = context.Background
var _ = io.Copy
var _ = strconv.Itoa
var _ = struc.Pack