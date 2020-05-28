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
	"strconv"
	"strings"

	"github.com/pkg/errors"

	ip_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2001/ip_types"
	lbba "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2001/lb"
	lb "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/lb"
)

func (h *LbVppHandler) SetLBConf(address string, buckets uint32, timeout uint32) (err error) {
	Addr, err := ipToAddress(address)
	// h.log.Warnf("DEBUG_STUFF : SetLBConf %v %v %v %v %v", address, Addr, buckets, timeout, err)
	if err != nil {
		return errors.Errorf("unable to parse source address %s", address)
	}
	reply := &lbba.LbConfReply{}
	switch Addr.Af {
	case ip_types.ADDRESS_IP4:
		// Addr.Un.GetIP4 doesn't return expected value
		ip4, err := ipTo4Address(address)
		req := &lbba.LbConf{
			FlowTimeout:          timeout,
			IP4SrcAddress:        ip4,
			StickyBucketsPerCore: buckets,
		}
		// h.log.Warnf("DEBUG_STUFF : SetLBConf before VPP request %v\n", &lbba.LbConf{FlowTimeout: timeout, IP4SrcAddress: Addr.Un.GetIP4(), StickyBucketsPerCore: buckets})
		// h.log.Warnf("DEBUG_STUFF : SetLBConf before VPP request %v\n", req)
		if err = h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
			//		h.log.Warnf("DEBUG_STUFF : Inside SetLBConf with error, reply from vpp : %v", reply)
			return err
		}
	case ip_types.ADDRESS_IP6:
		req := &lbba.LbConf{
			FlowTimeout:          timeout,
			IP6SrcAddress:        Addr.Un.GetIP6(),
			StickyBucketsPerCore: buckets,
		}
		if err = h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
			//		h.log.Warnf("DEBUG_STUFF : Inside SetLBConf with error, reply from vpp : %v", reply)
			return err
		}
	}

	//	h.log.Warnf("DEBUG_STUFF : Inside SetLBConf, no error reply from vpp : %v", reply)
	return nil

}

func (h *LbVppHandler) AddLBVip(prefix string, protocol uint8, port uint16, encap lb.LBVip_Encap, dscp uint8, srv_type lb.LBVip_SrvType, targetport, nodeport uint16, len uint32) error {
	return h.UpdateLBVip(prefix, protocol, port, encap, dscp, srv_type, targetport, nodeport, len, false)
}

func (h *LbVppHandler) DelLBVip(prefix string, protocol uint8, port uint16) error {
	return h.UpdateLBVip(prefix, protocol, port, 0, 0, 0, 0, 0, 0, true)
}

func (h *LbVppHandler) UpdateLBVip(prefix string, protocol uint8, port uint16, encap lb.LBVip_Encap, dscp uint8, srv_type lb.LBVip_SrvType, targetport, nodeport uint16, len uint32, del bool) error {

	var e lbba.LbEncapType
	var s lbba.LbSrvType

	switch encap {
	case lb.LBVip_GRE4:
		e = lbba.LB_API_ENCAP_TYPE_GRE4
	case lb.LBVip_GRE6:
		e = lbba.LB_API_ENCAP_TYPE_GRE6
	case lb.LBVip_NAT4:
		e = lbba.LB_API_ENCAP_TYPE_NAT4
	case lb.LBVip_NAT6:
		e = lbba.LB_API_ENCAP_TYPE_NAT6
	case lb.LBVip_L3DSR:
		e = lbba.LB_API_ENCAP_TYPE_L3DSR
	}

	switch srv_type {
	case lb.LBVip_CLUSTERIP:
		s = lbba.LB_API_SRV_TYPE_CLUSTERIP
	case lb.LBVip_NODEPORT:
		s = lbba.LB_API_SRV_TYPE_NODEPORT
	}

	Pfx, err := toPrefix(prefix, h)
	// h.log.Warnf("DEBUG_STUFF : UpdateLBVip %v %v %v %v", prefix, encap, len, err)
	// h.log.Warnf("DEBUG_STUFF : UpdateLBVip toPrefix %v", Pfx)
	if err != nil {
		return errors.Errorf("unable to parse source address %s", prefix)
	}

	// Need to set a minimal len
	if len < 1024 {
		len = 1024
	}

	req := &lbba.LbAddDelVip{
		Pfx:                 Pfx,
		Protocol:            protocol,
		Port:                port,
		Encap:               e,
		Dscp:                dscp,
		Type:                s,
		TargetPort:          targetport,
		NodePort:            nodeport,
		NewFlowsTableLength: len,
		IsDel:               del,
	}
	// dscp                uint8
	// type                lbsrvtype
	// targetport          uint16
	// nodeport            uint16
	reply := &lbba.LbAddDelVipReply{}
	// h.log.Warnf("DEBUG_STUFF : Inside UpdateLBVip req: %v", req)

	if err := h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		// h.log.Warnf("DEBUG_STUFF : Inside UpdateLBVip with error, reply from vpp : %v", reply)
		return err
	}
	// h.log.Warnf("DEBUG_STUFF : Inside UpdateLBVip, no error reply from vpp : %v", reply)
	return nil
}

func (h *LbVppHandler) AddLBAs(prefix string, protocol uint8, port uint16, address string) error {
	return h.UpdateLBAs(prefix, protocol, port, address, false, false)
}

func (h *LbVppHandler) DelLBAs(prefix string, protocol uint8, port uint16, address string) error {
	return h.UpdateLBAs(prefix, protocol, port, address, false, true)
}

func (h *LbVppHandler) UpdateLBAs(prefix string, protocol uint8, port uint16, address string, flush, del bool) error {

	Pfx, err := toPrefix(prefix, h)
	// h.log.Warnf("DEBUG_STUFF : UpdateLBAs %v %v %v %v", prefix, protocol, port, err)
	// h.log.Warnf("DEBUG_STUFF : UpdateLBAs toPrefix %v", Pfx)
	if err != nil {
		return errors.Errorf("unable to parse CIDR address %s", prefix)
	}

	Addr, err := ipToAddress(address)
	// h.log.Warnf("DEBUG_STUFF : UpdateLBAs %v %v %v %v", prefix, protocol, port, err)
	// h.log.Warnf("DEBUG_STUFF : UpdateLBAs toPrefix %v", Addr)
	if err != nil {
		return errors.Errorf("unable to parse AS address %s", prefix)
	}

	req := &lbba.LbAddDelAs{
		Pfx:       Pfx,
		Protocol:  protocol,
		Port:      port,
		AsAddress: Addr,
		IsDel:     flush,
		IsFlush:   del,
	}
	reply := &lbba.LbAddDelAsReply{}
	// h.log.Warnf("DEBUG_STUFF : Inside UpdateLBAs req: %v", req)

	if err := h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		// h.log.Warnf("DEBUG_STUFF : Inside UpdateLBAs with error, reply from vpp : %v", reply)
		return err
	}
	// h.log.Warnf("DEBUG_STUFF : Inside UpdateLBAs, no error reply from vpp : %v", reply)
	return nil
}

func ipTo4Address(ipStr string) (addr lbba.IP4Address, err error) {
	netIP := net.ParseIP(ipStr)
	if netIP == nil {
		return lbba.IP4Address{}, fmt.Errorf("invalid IP: %q", ipStr)
	}
	if ip4 := netIP.To4(); ip4 != nil {
		var ip4Addr lbba.IP4Address
		copy(ip4Addr[:], netIP.To4())
		addr = ip4Addr
	} else {
		return lbba.IP4Address{}, fmt.Errorf("required IPv4, provided: %q", ipStr)
	}
	return
}

func ipToAddress(ipStr string) (addr lbba.Address, err error) {
	var (
		a  []byte
		b  [16]byte
		af ip_types.AddressFamily
	)
	address := net.ParseIP(ipStr)
	a = address.To16()
	if address == nil {
		return lbba.Address{}, fmt.Errorf("invalid IP: %q", ipStr)
	}
	if ip := address.To4(); ip != nil {
		copy(b[12:16], a[12:16])
		af = ip_types.ADDRESS_IP4
	} else if ip := address.To16(); ip != nil {
		copy(b[0:16], a[0:16])
		af = ip_types.ADDRESS_IP6
	} else {
		return lbba.Address{}, fmt.Errorf("required IPv4, provided: %q", ipStr)
	}
	// fmt.Printf("DEBUG_STUFF : ipToAddress Returning : %v\n", ip_types.Address{Af: af, Un: ip_types.AddressUnion{XXX_UnionData: b}})
	return ip_types.Address{
		Af: af,
		Un: ip_types.AddressUnion{XXX_UnionData: b},
	}, nil
}

func toPrefix(ipStr string, h *LbVppHandler) (addr lbba.AddressWithPrefix, err error) {
	var (
		a    []byte
		b    [16]byte
		t    int
		mask uint8
	)
	address, _, e := net.ParseCIDR(ipStr)
	// h.log.Warnf("DEBUG_STUFF : toPrefix ParseCIDR %v %v %v", address, n, e)
	if e != nil {
		address = net.ParseIP(ipStr)
		if address == nil {
			// h.log.Warnf("DEBUG_STUFF : toPrefix ParseIp %v %v %v", address, n, e)
			return lbba.AddressWithPrefix{}, fmt.Errorf("required IPv4 prefix, provided: %q", ipStr)
		}
		e = nil
	}
	a = address.To16()
	// h.log.Warnf("DEBUG_STUFF : AddLBVip toPrefix a: %v", a)
	s := strings.Split(ipStr, "/")
	if address == nil {
		return ip_types.AddressWithPrefix{}, e
	}
	if ip := address.To4(); ip != nil {
		// h.log.Warnf("DEBUG_STUFF : AddLBVip IPv4 %v", s)
		if len(s) == 1 {
			mask = 128
		} else {
			t, _ = strconv.Atoi(s[1])
			mask = uint8(t)
			mask += 96
		}
		copy(b[12:16], a[12:16])
		// h.log.Warnf("DEBUG_STUFF : AddLBVip toPrefix b: %v", b)
		return ip_types.AddressWithPrefix{
			Len: mask,
			Address: lbba.Address{
				Af: 0,
				Un: ip_types.AddressUnion{XXX_UnionData: b},
			},
		}, nil
	}
	if ip := address.To16(); ip != nil {
		// h.log.Warnf("DEBUG_STUFF : AddLBVip IPv6 %v", s)
		if len(s) == 1 {
			mask = 128
		} else {
			t, _ = strconv.Atoi(s[1])
			mask = uint8(t)
		}
		copy(b[0:16], a[0:16])
		// h.log.Warnf("DEBUG_STUFF : AddLBVip toPrefix b: %v", b)
		return ip_types.AddressWithPrefix{
			Len: mask,
			Address: lbba.Address{
				Af: 1,
				Un: ip_types.AddressUnion{XXX_UnionData: b},
			},
		}, nil
	}
	return lbba.AddressWithPrefix{}, e
}
