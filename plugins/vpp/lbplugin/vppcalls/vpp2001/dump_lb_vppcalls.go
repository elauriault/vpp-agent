//  Copyright (c) 2020 Ubisoft Entertainment and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package vpp2001

import (
	"fmt"
	"net"

	ip_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2001/ip_types"
	lbba "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2001/lb"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/lbplugin/vppcalls"
	lb "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/lb"
)

func (h *LbVppHandler) LBGlobalConfigDump() (*lb.LBGlobal, error) {
	// h.log.Warnf("DEBUG_STUFF : LBGlobalConfigDump : curently no API in VPP lb plugin")
	return nil, vppcalls.ErrGlobalConfigDumpNotImplemented

}

func (h *LbVppHandler) DefaultLBGlobalConfig() *lb.LBGlobal {
	// h.log.Warnf("DEBUG_STUFF : DefaultLBGlobalConfig : returns defaults")
	return &lb.LBGlobal{
		Ip4SrcAddress: "255.255.255.255",
		Ip6SrcAddress: "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff",
		Bucket:        4,
		Timeout:       4,
	}
}

func uintToBool(value uint8) bool {
	if value == 0 {
		return false
	}
	return true
}

func (h *LbVppHandler) LBVipDump() ([]*lb.LBVip, error) {
	var vipLst []*lb.LBVip

	// h.log.Warnf("DEBUG_STUFF : Entering LBVipDump")
	req := &lbba.LbVipDump{}
	reqContext := h.callsChannel.SendMultiRequest(req)

	for {
		var proto uint32
		var encap lb.LBVip_Encap
		var srvtype lb.LBVip_SrvType
		msg := &lbba.LbVipDetails{}
		stop, err := reqContext.ReceiveReply(msg)
		if err != nil {
			return nil, fmt.Errorf("failed to dump VIPs: %v", err)
		}
		if stop {
			break
		}

		if msg.Vip.Protocol == ip_types.IP_API_PROTO_RESERVED {
			proto = 0
		} else {
			proto = uint32(msg.Vip.Protocol)
		}

		// lb_vip_dump passes the value of vip->type as encap

		encap = VipTypeToEncap(lbba.LbVipType(msg.Encap))

		switch msg.SrvType {
		case lbba.LB_API_SRV_TYPE_CLUSTERIP:
			srvtype = lb.LBVip_CLUSTERIP
		case lbba.LB_API_SRV_TYPE_NODEPORT:
			srvtype = lb.LBVip_NODEPORT
		}
		// h.log.Warnf("DEBUG_STUFF : LBVipDump : %v\n", msg)
		// h.log.Warnf("DEBUG_STUFF : LBVipDump encap : %v, %v, %v\n", msg.Encap, lbba.LbVipType(msg.Encap), encap)
		// h.log.Warnf("DEBUG_STUFF : LBVipDump encap : %d, %d\n", msg.Encap, encap)
		// h.log.Warnf("DEBUG_STUFF : LBVipDump proto : %v, %d\n", msg.Vip.Protocol, proto)

		vip := &lb.LBVip{
			Prefix:     PrefixToString(msg.Vip.Pfx),
			Protocol:   proto,
			Port:       uint32(msg.Vip.Port),
			Encap:      encap,
			Dscp:       uint32(msg.Dscp),
			SrvType:    srvtype,
			TargetPort: uint32(msg.TargetPort),
			NewLen:     uint32(msg.FlowTableLength),
		}
		vipLst = append(vipLst, vip)
	}

	return vipLst, nil
}

func (h *LbVppHandler) LBAsDump() ([]*lb.LBAs, error) {
	var asLst []*lb.LBAs
	var proto uint32

	// h.log.Warnf("DEBUG_STUFF : Entering LBAsDump")
	req := &lbba.LbAsDump{}
	reqContext := h.callsChannel.SendMultiRequest(req)

	for {
		msg := &lbba.LbAsDetails{}
		stop, err := reqContext.ReceiveReply(msg)
		if err != nil {
			return nil, fmt.Errorf("failed to dump ASs: %v", err)
		}
		if stop {
			break
		}

		if msg.Vip.Protocol == ip_types.IP_API_PROTO_RESERVED {
			proto = 0
		} else {
			proto = uint32(msg.Vip.Protocol)
		}

		// h.log.Warnf("DEBUG_STUFF : LBAsDump : %v", msg)
		as := &lb.LBAs{
			Prefix:   PrefixToString(msg.Vip.Pfx),
			Protocol: proto,
			Port:     uint32(msg.Vip.Port),
			Address:  AddressToString(msg.AppSrv),
		}
		asLst = append(asLst, as)
	}
	// h.log.Warnf("DEBUG_STUFF : LBVipDump : %v", asLst)

	return asLst, nil
}
func AddressToString(p ip_types.Address) string {
	var str string

	switch af := p.Af; af {
	case ip_types.ADDRESS_IP4:
		ip := p.Un.GetIP6()
		str = net.IP(ip[0:4]).String()
	case ip_types.ADDRESS_IP6:
		ip := p.Un.GetIP6()
		str = net.IP(ip[:]).String()
	}

	return str
}

func PrefixToString(p lbba.AddressWithPrefix) string {
	var str string

	ip := p.Address.Un.GetIP6()
	switch af := p.Address.Af; af {
	case ip_types.ADDRESS_IP4:
		if p.Len == 128 {
			str = fmt.Sprintf("%v", net.IP(ip[0:4]).String())
		} else {
			str = fmt.Sprintf("%v/%d", net.IP(ip[0:4]).String(), p.Len-96)
		}
	case ip_types.ADDRESS_IP6:
		if p.Len == 128 {
			str = fmt.Sprintf("%v", net.IP(ip[:]).String())
		} else {
			str = fmt.Sprintf("%v/%d", net.IP(ip[:]).String(), p.Len)
		}
	}
	// fmt.Printf("DEBUG_STUFF : PrefixToString : %v, %v\n", p, str)

	return str
}

func VipTypeToEncap(t lbba.LbVipType) (e lb.LBVip_Encap) {

	switch t {
	case lbba.LB_API_VIP_TYPE_IP4_GRE4:
		e = lb.LBVip_GRE4
	case lbba.LB_API_VIP_TYPE_IP4_GRE6:
		e = lb.LBVip_GRE6
	case lbba.LB_API_VIP_TYPE_IP6_GRE4:
		e = lb.LBVip_GRE4
	case lbba.LB_API_VIP_TYPE_IP6_GRE6:
		e = lb.LBVip_GRE6
	case lbba.LB_API_VIP_TYPE_IP4_L3DSR:
		e = lb.LBVip_L3DSR
	case lbba.LB_API_VIP_TYPE_IP4_NAT4:
		e = lb.LBVip_NAT4
	case lbba.LB_API_VIP_TYPE_IP6_NAT6:
		e = lb.LBVip_NAT6
	}

	return

}
