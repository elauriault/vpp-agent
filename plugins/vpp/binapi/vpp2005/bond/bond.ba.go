// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /usr/share/vpp/api/core/bond.api.json

/*
Package bond is a generated VPP binary API for 'bond' module.

It consists of:
	  8 enums
	  2 aliases
	 14 messages
	  7 services
*/
package bond

import (
	"bytes"
	"context"
	"io"
	"strconv"

	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"

	ethernet_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2005/ethernet_types"
	interface_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2005/interface_types"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "bond"
	// APIVersion is the API version of this module.
	APIVersion = "2.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0xa65a4a1e
)

// BondLbAlgo represents VPP binary API enum 'bond_lb_algo'.
type BondLbAlgo uint32

const (
	BOND_API_LB_ALGO_L2  BondLbAlgo = 0
	BOND_API_LB_ALGO_L34 BondLbAlgo = 1
	BOND_API_LB_ALGO_L23 BondLbAlgo = 2
	BOND_API_LB_ALGO_RR  BondLbAlgo = 3
	BOND_API_LB_ALGO_BC  BondLbAlgo = 4
	BOND_API_LB_ALGO_AB  BondLbAlgo = 5
)

var BondLbAlgo_name = map[uint32]string{
	0: "BOND_API_LB_ALGO_L2",
	1: "BOND_API_LB_ALGO_L34",
	2: "BOND_API_LB_ALGO_L23",
	3: "BOND_API_LB_ALGO_RR",
	4: "BOND_API_LB_ALGO_BC",
	5: "BOND_API_LB_ALGO_AB",
}

var BondLbAlgo_value = map[string]uint32{
	"BOND_API_LB_ALGO_L2":  0,
	"BOND_API_LB_ALGO_L34": 1,
	"BOND_API_LB_ALGO_L23": 2,
	"BOND_API_LB_ALGO_RR":  3,
	"BOND_API_LB_ALGO_BC":  4,
	"BOND_API_LB_ALGO_AB":  5,
}

func (x BondLbAlgo) String() string {
	s, ok := BondLbAlgo_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// BondMode represents VPP binary API enum 'bond_mode'.
type BondMode uint32

const (
	BOND_API_MODE_ROUND_ROBIN   BondMode = 1
	BOND_API_MODE_ACTIVE_BACKUP BondMode = 2
	BOND_API_MODE_XOR           BondMode = 3
	BOND_API_MODE_BROADCAST     BondMode = 4
	BOND_API_MODE_LACP          BondMode = 5
)

var BondMode_name = map[uint32]string{
	1: "BOND_API_MODE_ROUND_ROBIN",
	2: "BOND_API_MODE_ACTIVE_BACKUP",
	3: "BOND_API_MODE_XOR",
	4: "BOND_API_MODE_BROADCAST",
	5: "BOND_API_MODE_LACP",
}

var BondMode_value = map[string]uint32{
	"BOND_API_MODE_ROUND_ROBIN":   1,
	"BOND_API_MODE_ACTIVE_BACKUP": 2,
	"BOND_API_MODE_XOR":           3,
	"BOND_API_MODE_BROADCAST":     4,
	"BOND_API_MODE_LACP":          5,
}

func (x BondMode) String() string {
	s, ok := BondMode_name[uint32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

type IfStatusFlags = interface_types.IfStatusFlags

type IfType = interface_types.IfType

type LinkDuplex = interface_types.LinkDuplex

type MtuProto = interface_types.MtuProto

type RxMode = interface_types.RxMode

type SubIfFlags = interface_types.SubIfFlags

type InterfaceIndex = interface_types.InterfaceIndex

type MacAddress = ethernet_types.MacAddress

// BondCreate represents VPP binary API message 'bond_create'.
type BondCreate struct {
	ID           uint32
	UseCustomMac bool
	MacAddress   MacAddress
	Mode         BondMode
	Lb           BondLbAlgo
	NumaOnly     bool
}

func (m *BondCreate) Reset()                        { *m = BondCreate{} }
func (*BondCreate) GetMessageName() string          { return "bond_create" }
func (*BondCreate) GetCrcString() string            { return "48883c7e" }
func (*BondCreate) GetMessageType() api.MessageType { return api.RequestMessage }

// BondCreateReply represents VPP binary API message 'bond_create_reply'.
type BondCreateReply struct {
	Retval    int32
	SwIfIndex InterfaceIndex
}

func (m *BondCreateReply) Reset()                        { *m = BondCreateReply{} }
func (*BondCreateReply) GetMessageName() string          { return "bond_create_reply" }
func (*BondCreateReply) GetCrcString() string            { return "5383d31f" }
func (*BondCreateReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// BondDelete represents VPP binary API message 'bond_delete'.
type BondDelete struct {
	SwIfIndex InterfaceIndex
}

func (m *BondDelete) Reset()                        { *m = BondDelete{} }
func (*BondDelete) GetMessageName() string          { return "bond_delete" }
func (*BondDelete) GetCrcString() string            { return "f9e6675e" }
func (*BondDelete) GetMessageType() api.MessageType { return api.RequestMessage }

// BondDeleteReply represents VPP binary API message 'bond_delete_reply'.
type BondDeleteReply struct {
	Retval int32
}

func (m *BondDeleteReply) Reset()                        { *m = BondDeleteReply{} }
func (*BondDeleteReply) GetMessageName() string          { return "bond_delete_reply" }
func (*BondDeleteReply) GetCrcString() string            { return "e8d4e804" }
func (*BondDeleteReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// BondDetachSlave represents VPP binary API message 'bond_detach_slave'.
type BondDetachSlave struct {
	SwIfIndex InterfaceIndex
}

func (m *BondDetachSlave) Reset()                        { *m = BondDetachSlave{} }
func (*BondDetachSlave) GetMessageName() string          { return "bond_detach_slave" }
func (*BondDetachSlave) GetCrcString() string            { return "f9e6675e" }
func (*BondDetachSlave) GetMessageType() api.MessageType { return api.RequestMessage }

// BondDetachSlaveReply represents VPP binary API message 'bond_detach_slave_reply'.
type BondDetachSlaveReply struct {
	Retval int32
}

func (m *BondDetachSlaveReply) Reset()                        { *m = BondDetachSlaveReply{} }
func (*BondDetachSlaveReply) GetMessageName() string          { return "bond_detach_slave_reply" }
func (*BondDetachSlaveReply) GetCrcString() string            { return "e8d4e804" }
func (*BondDetachSlaveReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// BondEnslave represents VPP binary API message 'bond_enslave'.
type BondEnslave struct {
	SwIfIndex     InterfaceIndex
	BondSwIfIndex InterfaceIndex
	IsPassive     bool
	IsLongTimeout bool
}

func (m *BondEnslave) Reset()                        { *m = BondEnslave{} }
func (*BondEnslave) GetMessageName() string          { return "bond_enslave" }
func (*BondEnslave) GetCrcString() string            { return "076ecfa7" }
func (*BondEnslave) GetMessageType() api.MessageType { return api.RequestMessage }

// BondEnslaveReply represents VPP binary API message 'bond_enslave_reply'.
type BondEnslaveReply struct {
	Retval int32
}

func (m *BondEnslaveReply) Reset()                        { *m = BondEnslaveReply{} }
func (*BondEnslaveReply) GetMessageName() string          { return "bond_enslave_reply" }
func (*BondEnslaveReply) GetCrcString() string            { return "e8d4e804" }
func (*BondEnslaveReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// SwInterfaceBondDetails represents VPP binary API message 'sw_interface_bond_details'.
type SwInterfaceBondDetails struct {
	SwIfIndex     InterfaceIndex
	ID            uint32
	Mode          BondMode
	Lb            BondLbAlgo
	NumaOnly      bool
	ActiveSlaves  uint32
	Slaves        uint32
	InterfaceName string `struc:"[64]byte"`
}

func (m *SwInterfaceBondDetails) Reset()                        { *m = SwInterfaceBondDetails{} }
func (*SwInterfaceBondDetails) GetMessageName() string          { return "sw_interface_bond_details" }
func (*SwInterfaceBondDetails) GetCrcString() string            { return "f5ef2106" }
func (*SwInterfaceBondDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// SwInterfaceBondDump represents VPP binary API message 'sw_interface_bond_dump'.
type SwInterfaceBondDump struct{}

func (m *SwInterfaceBondDump) Reset()                        { *m = SwInterfaceBondDump{} }
func (*SwInterfaceBondDump) GetMessageName() string          { return "sw_interface_bond_dump" }
func (*SwInterfaceBondDump) GetCrcString() string            { return "51077d14" }
func (*SwInterfaceBondDump) GetMessageType() api.MessageType { return api.RequestMessage }

// SwInterfaceSetBondWeight represents VPP binary API message 'sw_interface_set_bond_weight'.
type SwInterfaceSetBondWeight struct {
	SwIfIndex InterfaceIndex
	Weight    uint32
}

func (m *SwInterfaceSetBondWeight) Reset()                        { *m = SwInterfaceSetBondWeight{} }
func (*SwInterfaceSetBondWeight) GetMessageName() string          { return "sw_interface_set_bond_weight" }
func (*SwInterfaceSetBondWeight) GetCrcString() string            { return "deb510a0" }
func (*SwInterfaceSetBondWeight) GetMessageType() api.MessageType { return api.RequestMessage }

// SwInterfaceSetBondWeightReply represents VPP binary API message 'sw_interface_set_bond_weight_reply'.
type SwInterfaceSetBondWeightReply struct {
	Retval int32
}

func (m *SwInterfaceSetBondWeightReply) Reset() { *m = SwInterfaceSetBondWeightReply{} }
func (*SwInterfaceSetBondWeightReply) GetMessageName() string {
	return "sw_interface_set_bond_weight_reply"
}
func (*SwInterfaceSetBondWeightReply) GetCrcString() string            { return "e8d4e804" }
func (*SwInterfaceSetBondWeightReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// SwInterfaceSlaveDetails represents VPP binary API message 'sw_interface_slave_details'.
type SwInterfaceSlaveDetails struct {
	SwIfIndex     InterfaceIndex
	InterfaceName string `struc:"[64]byte"`
	IsPassive     bool
	IsLongTimeout bool
	IsLocalNuma   bool
	Weight        uint32
}

func (m *SwInterfaceSlaveDetails) Reset()                        { *m = SwInterfaceSlaveDetails{} }
func (*SwInterfaceSlaveDetails) GetMessageName() string          { return "sw_interface_slave_details" }
func (*SwInterfaceSlaveDetails) GetCrcString() string            { return "3c4a0e23" }
func (*SwInterfaceSlaveDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// SwInterfaceSlaveDump represents VPP binary API message 'sw_interface_slave_dump'.
type SwInterfaceSlaveDump struct {
	SwIfIndex InterfaceIndex
}

func (m *SwInterfaceSlaveDump) Reset()                        { *m = SwInterfaceSlaveDump{} }
func (*SwInterfaceSlaveDump) GetMessageName() string          { return "sw_interface_slave_dump" }
func (*SwInterfaceSlaveDump) GetCrcString() string            { return "f9e6675e" }
func (*SwInterfaceSlaveDump) GetMessageType() api.MessageType { return api.RequestMessage }

func init() {
	api.RegisterMessage((*BondCreate)(nil), "bond.BondCreate")
	api.RegisterMessage((*BondCreateReply)(nil), "bond.BondCreateReply")
	api.RegisterMessage((*BondDelete)(nil), "bond.BondDelete")
	api.RegisterMessage((*BondDeleteReply)(nil), "bond.BondDeleteReply")
	api.RegisterMessage((*BondDetachSlave)(nil), "bond.BondDetachSlave")
	api.RegisterMessage((*BondDetachSlaveReply)(nil), "bond.BondDetachSlaveReply")
	api.RegisterMessage((*BondEnslave)(nil), "bond.BondEnslave")
	api.RegisterMessage((*BondEnslaveReply)(nil), "bond.BondEnslaveReply")
	api.RegisterMessage((*SwInterfaceBondDetails)(nil), "bond.SwInterfaceBondDetails")
	api.RegisterMessage((*SwInterfaceBondDump)(nil), "bond.SwInterfaceBondDump")
	api.RegisterMessage((*SwInterfaceSetBondWeight)(nil), "bond.SwInterfaceSetBondWeight")
	api.RegisterMessage((*SwInterfaceSetBondWeightReply)(nil), "bond.SwInterfaceSetBondWeightReply")
	api.RegisterMessage((*SwInterfaceSlaveDetails)(nil), "bond.SwInterfaceSlaveDetails")
	api.RegisterMessage((*SwInterfaceSlaveDump)(nil), "bond.SwInterfaceSlaveDump")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*BondCreate)(nil),
		(*BondCreateReply)(nil),
		(*BondDelete)(nil),
		(*BondDeleteReply)(nil),
		(*BondDetachSlave)(nil),
		(*BondDetachSlaveReply)(nil),
		(*BondEnslave)(nil),
		(*BondEnslaveReply)(nil),
		(*SwInterfaceBondDetails)(nil),
		(*SwInterfaceBondDump)(nil),
		(*SwInterfaceSetBondWeight)(nil),
		(*SwInterfaceSetBondWeightReply)(nil),
		(*SwInterfaceSlaveDetails)(nil),
		(*SwInterfaceSlaveDump)(nil),
	}
}

// RPCService represents RPC service API for bond module.
type RPCService interface {
	DumpSwInterfaceBond(ctx context.Context, in *SwInterfaceBondDump) (RPCService_DumpSwInterfaceBondClient, error)
	DumpSwInterfaceSlave(ctx context.Context, in *SwInterfaceSlaveDump) (RPCService_DumpSwInterfaceSlaveClient, error)
	BondCreate(ctx context.Context, in *BondCreate) (*BondCreateReply, error)
	BondDelete(ctx context.Context, in *BondDelete) (*BondDeleteReply, error)
	BondDetachSlave(ctx context.Context, in *BondDetachSlave) (*BondDetachSlaveReply, error)
	BondEnslave(ctx context.Context, in *BondEnslave) (*BondEnslaveReply, error)
	SwInterfaceSetBondWeight(ctx context.Context, in *SwInterfaceSetBondWeight) (*SwInterfaceSetBondWeightReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpSwInterfaceBond(ctx context.Context, in *SwInterfaceBondDump) (RPCService_DumpSwInterfaceBondClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpSwInterfaceBondClient{stream}
	return x, nil
}

type RPCService_DumpSwInterfaceBondClient interface {
	Recv() (*SwInterfaceBondDetails, error)
}

type serviceClient_DumpSwInterfaceBondClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpSwInterfaceBondClient) Recv() (*SwInterfaceBondDetails, error) {
	m := new(SwInterfaceBondDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpSwInterfaceSlave(ctx context.Context, in *SwInterfaceSlaveDump) (RPCService_DumpSwInterfaceSlaveClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpSwInterfaceSlaveClient{stream}
	return x, nil
}

type RPCService_DumpSwInterfaceSlaveClient interface {
	Recv() (*SwInterfaceSlaveDetails, error)
}

type serviceClient_DumpSwInterfaceSlaveClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpSwInterfaceSlaveClient) Recv() (*SwInterfaceSlaveDetails, error) {
	m := new(SwInterfaceSlaveDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) BondCreate(ctx context.Context, in *BondCreate) (*BondCreateReply, error) {
	out := new(BondCreateReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BondDelete(ctx context.Context, in *BondDelete) (*BondDeleteReply, error) {
	out := new(BondDeleteReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BondDetachSlave(ctx context.Context, in *BondDetachSlave) (*BondDetachSlaveReply, error) {
	out := new(BondDetachSlaveReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) BondEnslave(ctx context.Context, in *BondEnslave) (*BondEnslaveReply, error) {
	out := new(BondEnslaveReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SwInterfaceSetBondWeight(ctx context.Context, in *SwInterfaceSetBondWeight) (*SwInterfaceSetBondWeightReply, error) {
	out := new(SwInterfaceSetBondWeightReply)
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
