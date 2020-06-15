// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /usr/share/vpp/api/core/gre.api.json

/*
Package gre is a generated VPP binary API for 'gre' module.

It consists of:
	  4 messages
	  2 services
*/
package gre

import (
	"bytes"
	"context"
	"io"
	"strconv"

	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "gre"
	// APIVersion is the API version of this module.
	APIVersion = "1.0.1"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x83d2c14f
)

// GreAddDelTunnel represents VPP binary API message 'gre_add_del_tunnel'.
type GreAddDelTunnel struct {
	IsAdd      uint8
	IsIPv6     uint8
	TunnelType uint8
	Instance   uint32
	SrcAddress []byte `struc:"[16]byte"`
	DstAddress []byte `struc:"[16]byte"`
	OuterFibID uint32
	SessionID  uint16
}

func (m *GreAddDelTunnel) Reset()                        { *m = GreAddDelTunnel{} }
func (*GreAddDelTunnel) GetMessageName() string          { return "gre_add_del_tunnel" }
func (*GreAddDelTunnel) GetCrcString() string            { return "9f03ede2" }
func (*GreAddDelTunnel) GetMessageType() api.MessageType { return api.RequestMessage }

// GreAddDelTunnelReply represents VPP binary API message 'gre_add_del_tunnel_reply'.
type GreAddDelTunnelReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (m *GreAddDelTunnelReply) Reset()                        { *m = GreAddDelTunnelReply{} }
func (*GreAddDelTunnelReply) GetMessageName() string          { return "gre_add_del_tunnel_reply" }
func (*GreAddDelTunnelReply) GetCrcString() string            { return "fda5941f" }
func (*GreAddDelTunnelReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// GreTunnelDetails represents VPP binary API message 'gre_tunnel_details'.
type GreTunnelDetails struct {
	SwIfIndex  uint32
	Instance   uint32
	IsIPv6     uint8
	TunnelType uint8
	SrcAddress []byte `struc:"[16]byte"`
	DstAddress []byte `struc:"[16]byte"`
	OuterFibID uint32
	SessionID  uint16
}

func (m *GreTunnelDetails) Reset()                        { *m = GreTunnelDetails{} }
func (*GreTunnelDetails) GetMessageName() string          { return "gre_tunnel_details" }
func (*GreTunnelDetails) GetCrcString() string            { return "1a12b8c1" }
func (*GreTunnelDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// GreTunnelDump represents VPP binary API message 'gre_tunnel_dump'.
type GreTunnelDump struct {
	SwIfIndex uint32
}

func (m *GreTunnelDump) Reset()                        { *m = GreTunnelDump{} }
func (*GreTunnelDump) GetMessageName() string          { return "gre_tunnel_dump" }
func (*GreTunnelDump) GetCrcString() string            { return "529cb13f" }
func (*GreTunnelDump) GetMessageType() api.MessageType { return api.RequestMessage }

func init() {
	api.RegisterMessage((*GreAddDelTunnel)(nil), "gre.GreAddDelTunnel")
	api.RegisterMessage((*GreAddDelTunnelReply)(nil), "gre.GreAddDelTunnelReply")
	api.RegisterMessage((*GreTunnelDetails)(nil), "gre.GreTunnelDetails")
	api.RegisterMessage((*GreTunnelDump)(nil), "gre.GreTunnelDump")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*GreAddDelTunnel)(nil),
		(*GreAddDelTunnelReply)(nil),
		(*GreTunnelDetails)(nil),
		(*GreTunnelDump)(nil),
	}
}

// RPCService represents RPC service API for gre module.
type RPCService interface {
	DumpGreTunnel(ctx context.Context, in *GreTunnelDump) (RPCService_DumpGreTunnelClient, error)
	GreAddDelTunnel(ctx context.Context, in *GreAddDelTunnel) (*GreAddDelTunnelReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpGreTunnel(ctx context.Context, in *GreTunnelDump) (RPCService_DumpGreTunnelClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpGreTunnelClient{stream}
	return x, nil
}

type RPCService_DumpGreTunnelClient interface {
	Recv() (*GreTunnelDetails, error)
}

type serviceClient_DumpGreTunnelClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpGreTunnelClient) Recv() (*GreTunnelDetails, error) {
	m := new(GreTunnelDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) GreAddDelTunnel(ctx context.Context, in *GreAddDelTunnel) (*GreAddDelTunnelReply, error) {
	out := new(GreAddDelTunnelReply)
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
