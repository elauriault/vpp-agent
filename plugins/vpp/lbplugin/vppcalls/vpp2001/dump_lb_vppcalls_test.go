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

package vpp2001_test

import (
	"net"
	"strconv"
	"strings"
	"testing"

	. "github.com/onsi/gomega"
	"go.ligato.io/cn-infra/v2/logging/logrus"
	ip_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2001/ip_types"
	lbba "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2001/lb"
	vpp_vpe "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2001/vpe"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/ifplugin/ifaceidx"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/lbplugin/vppcalls"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/lbplugin/vppcalls/vpp2001"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/vppmock"
)

func TestLBVipDump(t *testing.T) {
	ctx, lbHandler, _ := lbTestSetup(t)
	defer ctx.TeardownTestCtx()

	ctx.MockVpp.MockReply(
		&lbba.LbVipDetails{
			Vip: lbba.LbVip{
				Pfx:      toPrefix("2.3.4.5"),
				Protocol: 0,
				Port:     0,
			},
			Encap:           lbba.LB_API_ENCAP_TYPE_GRE4,
			Dscp:            0,
			SrvType:         0,
			TargetPort:      0,
			FlowTableLength: 1024,
		},
		&lbba.LbVipDetails{
			Vip: lbba.LbVip{
				Pfx:      toPrefix("10.0.0.1/24"),
				Protocol: 6,
				Port:     80,
			},
			Encap:           lbba.LB_API_ENCAP_TYPE_GRE4,
			Dscp:            0,
			SrvType:         0,
			TargetPort:      0,
			FlowTableLength: 1024,
		},
		&lbba.LbVipDetails{
			Vip: lbba.LbVip{
				Pfx:      toPrefix("2600:2600:2600:2600::1/64"),
				Protocol: 6,
				Port:     80,
			},
			Encap:           lbba.LB_API_ENCAP_TYPE_GRE6,
			Dscp:            0,
			SrvType:         0,
			TargetPort:      0,
			FlowTableLength: 1024,
		},
	)
	ctx.MockVpp.MockReply(&vpp_vpe.ControlPingReply{})

	vips, err := lbHandler.LBVipDump()
	Expect(err).To(Succeed())

	Expect(vips).To(HaveLen(3))

	Expect(vips[0].Prefix).To(Equal("2.3.4.5/32"))
	Expect(vips[0].Protocol).To(BeEquivalentTo(0))
	Expect(vips[0].Port).To(BeEquivalentTo(0))
	Expect(vips[0].Encap).To(BeEquivalentTo(0))
	Expect(vips[0].Dscp).To(BeEquivalentTo(0))
	Expect(vips[0].SrvType).To(BeEquivalentTo(0))
	Expect(vips[0].TargetPort).To(BeEquivalentTo(0))
	Expect(vips[0].NewLen).To(BeEquivalentTo(1024))

	Expect(vips[1].Prefix).To(Equal("10.0.0.1/24"))
	Expect(vips[1].Protocol).To(BeEquivalentTo(6))
	Expect(vips[1].Port).To(BeEquivalentTo(80))
	Expect(vips[1].Encap).To(BeEquivalentTo(0))
	Expect(vips[1].Dscp).To(BeEquivalentTo(0))
	Expect(vips[1].SrvType).To(BeEquivalentTo(0))
	Expect(vips[1].TargetPort).To(BeEquivalentTo(0))
	Expect(vips[1].NewLen).To(BeEquivalentTo(1024))

	Expect(vips[2].Prefix).To(Equal("2600:2600:2600:2600::1/64"))
	Expect(vips[2].Protocol).To(BeEquivalentTo(6))
	Expect(vips[2].Port).To(BeEquivalentTo(80))
	Expect(vips[2].Encap).To(BeEquivalentTo(1))
	Expect(vips[2].Dscp).To(BeEquivalentTo(0))
	Expect(vips[2].SrvType).To(BeEquivalentTo(0))
	Expect(vips[2].TargetPort).To(BeEquivalentTo(0))
	Expect(vips[2].NewLen).To(BeEquivalentTo(1024))

}

func TestLBAsDump(t *testing.T) {
	ctx, lbHandler, _ := lbTestSetup(t)
	defer ctx.TeardownTestCtx()

	ctx.MockVpp.MockReply(
		&lbba.LbAsDetails{
			Vip: lbba.LbVip{
				Pfx:      toPrefix("2.3.4.5"),
				Protocol: 0,
				Port:     0,
			},
			AppSrv: ipToAddress("10.0.0.1"),
		},
		&lbba.LbAsDetails{
			Vip: lbba.LbVip{
				Pfx:      toPrefix("10.0.0.1/24"),
				Protocol: 6,
				Port:     80,
			},
			AppSrv: ipToAddress("10.0.0.2"),
		},
		&lbba.LbAsDetails{
			Vip: lbba.LbVip{
				Pfx:      toPrefix("2600:2600:2600:2600::1/64"),
				Protocol: 0,
				Port:     0,
			},
			AppSrv: ipToAddress("2601:2601:2601:2601::1"),
		},
	)
	ctx.MockVpp.MockReply(&vpp_vpe.ControlPingReply{})

	as, err := lbHandler.LBAsDump()
	Expect(err).To(Succeed())

	Expect(as).To(HaveLen(3))

	Expect(as[0].Prefix).To(Equal("2.3.4.5/32"))
	Expect(as[0].Protocol).To(BeEquivalentTo(0))
	Expect(as[0].Port).To(BeEquivalentTo(0))
	Expect(as[0].Address).To(Equal("10.0.0.1"))

	Expect(as[1].Prefix).To(Equal("10.0.0.1/24"))
	Expect(as[1].Protocol).To(BeEquivalentTo(6))
	Expect(as[1].Port).To(BeEquivalentTo(80))
	Expect(as[1].Address).To(Equal("10.0.0.2"))

	Expect(as[2].Prefix).To(Equal("2600:2600:2600:2600::1/64"))
	Expect(as[2].Protocol).To(BeEquivalentTo(0))
	Expect(as[2].Port).To(BeEquivalentTo(0))
	Expect(as[2].Address).To(Equal("2601:2601:2601:2601::1"))
}

func lbTestSetup(t *testing.T) (*vppmock.TestCtx, vppcalls.LbVppAPI, ifaceidx.IfaceMetadataIndexRW) {
	ctx := vppmock.SetupTestCtx(t)
	logger := logrus.NewLogger("test-log")
	ifIndexes := ifaceidx.NewIfaceIndex(logger, "lb-if-idx")
	lbHandler := vpp2001.NewLbVppHandler(ctx.MockChannel, ifIndexes, logrus.DefaultLogger())
	return ctx, lbHandler, ifIndexes
}

func ipToAddress(ipStr string) (addr lbba.Address) {
	var (
		a  []byte
		b  [16]byte
		af ip_types.AddressFamily
	)
	address := net.ParseIP(ipStr)
	a = address.To16()
	if address == nil {
		return lbba.Address{}
	}
	if ip := address.To4(); ip != nil {
		copy(b[12:16], a[12:16])
		af = ip_types.ADDRESS_IP4
	} else if ip := address.To16(); ip != nil {
		copy(b[0:16], a[0:16])
		af = ip_types.ADDRESS_IP6
	} else {
		return lbba.Address{}
	}
	return ip_types.Address{
		Af: af,
		Un: ip_types.AddressUnion{XXX_UnionData: b},
	}
}

func toPrefix(ipStr string) (addr lbba.AddressWithPrefix) {
	var (
		a    []byte
		b    [16]byte
		t    int
		mask uint8
	)
	address, _, e := net.ParseCIDR(ipStr)
	if e != nil {
		address = net.ParseIP(ipStr)
		if address == nil {
			return
		}
		e = nil
	}
	a = address.To16()
	s := strings.Split(ipStr, "/")
	if address == nil {
		return ip_types.AddressWithPrefix{}
	}
	if ip := address.To4(); ip != nil {
		if len(s) == 1 {
			mask = 128
		} else {
			t, _ = strconv.Atoi(s[1])
			mask = uint8(t)
			mask += 96
		}
		copy(b[12:16], a[12:16])
		return ip_types.AddressWithPrefix{
			Len: mask,
			Address: lbba.Address{
				Af: 0,
				Un: ip_types.AddressUnion{XXX_UnionData: b},
			},
		}
	}
	if ip := address.To16(); ip != nil {
		if len(s) == 1 {
			mask = 128
		} else {
			t, _ = strconv.Atoi(s[1])
			mask = uint8(t)
		}
		copy(b[0:16], a[0:16])
		return ip_types.AddressWithPrefix{
			Len: mask,
			Address: lbba.Address{
				Af: 1,
				Un: ip_types.AddressUnion{XXX_UnionData: b},
			},
		}
	}
	return lbba.AddressWithPrefix{}
}
