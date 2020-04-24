//  Copyright (c) 2019 Cisco and/or its affiliates.
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
	lb "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/lb"
)

func (h *LbVppHandler) LBGlobalConfigDump() *lb.LBGlobal {
	// h.log.Warnf("DEBUG_STUFF : LBGlobalConfigDump : curently no API in VPP lb plugin")
	return nil
}

func (h *LbVppHandler) DefaultLBGlobalConfig() *lb.LBGlobal {
	// h.log.Warnf("DEBUG_STUFF : DefaultLBGlobalConfig : returns defaults")
	return &lb.LBGlobal{
		Ip4SrcAddress: "255.255.255.255",
		Ip6SrcAddress: "",
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
		msg := &lbba.LbVipDetails{}
		stop, err := reqContext.ReceiveReply(msg)
		if err != nil {
			return nil, fmt.Errorf("failed to dump VIPs: %v", err)
		}
		if stop {
			break
		}
		// h.log.Warnf("DEBUG_STUFF : LBVipDump : %v", msg)

		vip := &lb.LBVip{
			Prefix:     PrefixToString(msg.Vip.Pfx),
			Protocol:   uint32(msg.Vip.Protocol),
			Port:       uint32(msg.Vip.Port),
			Encap:      lb.LBVip_Encap(msg.Encap),
			Dscp:       uint32(msg.Dscp),
			SrvType:    lb.LBVip_SrvType(msg.SrvType),
			TargetPort: uint32(msg.TargetPort),
			NewLen:     uint32(msg.FlowTableLength),
		}
		vipLst = append(vipLst, vip)
	}

	return vipLst, nil
}

func (h *LbVppHandler) LBAsDump() ([]*lb.LBAs, error) {
	var asLst []*lb.LBAs

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
		// h.log.Warnf("DEBUG_STUFF : LBAsDump : %v", msg)
		as := &lb.LBAs{
			Prefix:   PrefixToString(msg.Vip.Pfx),
			Protocol: uint32(msg.Vip.Protocol),
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
		str = net.IP(ip[12:16]).String()
	case ip_types.ADDRESS_IP6:
		ip := p.Un.GetIP6()
		str = net.IP(ip[:]).String()
	}

	return str
}

func PrefixToString(p lbba.AddressWithPrefix) string {
	var str string

	switch af := p.Address.Af; af {
	case ip_types.ADDRESS_IP4:
		ip := p.Address.Un.GetIP6()
		str = fmt.Sprintf("%v/%d", net.IP(ip[12:16]).String(), p.Len-96)
	case ip_types.ADDRESS_IP6:
		ip := p.Address.Un.GetIP6()
		str = fmt.Sprintf("%v/%d", net.IP(ip[:]).String(), p.Len)
	}

	return str
}
