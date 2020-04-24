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
	"testing"

	. "github.com/onsi/gomega"
	lbba "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2001/lb"
	lb "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/lb"
)

func TestAddLBVip(t *testing.T) {
	ctx, lbHandler, _ := lbTestSetup(t)
	defer ctx.TeardownTestCtx()

	ctx.MockVpp.MockReply(&lbba.LbAddDelVipReply{})
	err := lbHandler.AddLBVip("1.2.3.4", 0, 0, lb.LBVip_GRE4, 0, 0, 0, 0, 1024)
	Expect(err).ShouldNot(HaveOccurred())

	msg, ok := ctx.MockChannel.Msg.(*lbba.LbAddDelVip)
	Expect(ok).To(BeTrue())
	Expect(msg).ToNot(BeNil())
	Expect(msg.Pfx).To(Equal(toPrefix("1.2.3.4/32")))
	Expect(msg.Encap).To(BeEquivalentTo(0))
	Expect(msg.NewFlowsTableLength).To(BeEquivalentTo(1024))

	ctx.MockVpp.MockReply(&lbba.LbAddDelVipReply{})
	err = lbHandler.AddLBVip("1.2.3.4/24", 6, 80, lb.LBVip_GRE4, 0, 0, 0, 0, 0)
	Expect(err).ShouldNot(HaveOccurred())

	msg, ok = ctx.MockChannel.Msg.(*lbba.LbAddDelVip)
	Expect(ok).To(BeTrue())
	Expect(msg).ToNot(BeNil())
	Expect(msg.Pfx).To(Equal(toPrefix("1.2.3.4/24")))
	Expect(msg.Encap).To(BeEquivalentTo(0))
	Expect(msg.Protocol).To(BeEquivalentTo(6))
	Expect(msg.Port).To(BeEquivalentTo(80))
	Expect(msg.NewFlowsTableLength).To(BeEquivalentTo(1024))

	ctx.MockVpp.MockReply(&lbba.LbAddDelVipReply{})
	err = lbHandler.AddLBVip("2600:2600:2600:2600::1/64", 6, 80, lb.LBVip_GRE6, 0, 0, 0, 0, 1024)
	Expect(err).ShouldNot(HaveOccurred())

	msg, ok = ctx.MockChannel.Msg.(*lbba.LbAddDelVip)
	Expect(ok).To(BeTrue())
	Expect(msg).ToNot(BeNil())
	Expect(msg.Pfx).To(Equal(toPrefix("2600:2600:2600:2600::1/64")))
	Expect(msg.Encap).To(BeEquivalentTo(1))
	Expect(msg.Protocol).To(BeEquivalentTo(6))
	Expect(msg.Port).To(BeEquivalentTo(80))
	Expect(msg.NewFlowsTableLength).To(BeEquivalentTo(1024))

}

func TestAddLBVipError(t *testing.T) {
	ctx, lbHandler, _ := lbTestSetup(t)
	defer ctx.TeardownTestCtx()

	// Invalid arguments
	ctx.MockVpp.MockReply(&lbba.LbAddDelVipReply{})
	err := lbHandler.AddLBVip("1.2.3", 0, 0, lb.LBVip_GRE4, 0, 0, 0, 0, 1024)
	Expect(err).Should(HaveOccurred())

}

func TestAddLBVipRetval(t *testing.T) {
	ctx, lbHandler, _ := lbTestSetup(t)
	defer ctx.TeardownTestCtx()

	// Invalid arguments
	ctx.MockVpp.MockReply(&lbba.LbAddDelVipReply{
		Retval: 1,
	})
	err := lbHandler.AddLBVip("1.2.3.4/24", 6, 80, lb.LBVip_GRE4, 0, 0, 0, 0, 0)
	Expect(err).Should(HaveOccurred())

}

func TestAddLbAs(t *testing.T) {
	ctx, lbHandler, _ := lbTestSetup(t)
	defer ctx.TeardownTestCtx()

	ctx.MockVpp.MockReply(&lbba.LbAddDelAsReply{})
	err := lbHandler.AddLBAs("1.2.3.4", 0, 0, "10.1.1.1")
	Expect(err).ShouldNot(HaveOccurred())

	msg, ok := ctx.MockChannel.Msg.(*lbba.LbAddDelAs)
	Expect(ok).To(BeTrue())
	Expect(msg).ToNot(BeNil())
	Expect(msg.Pfx).To(Equal(toPrefix("1.2.3.4")))
	Expect(msg.Protocol).To(BeEquivalentTo(0))
	Expect(msg.Port).To(BeEquivalentTo(0))
	Expect(msg.AsAddress).To(Equal(ipToAddress("10.1.1.1")))

	ctx.MockVpp.MockReply(&lbba.LbAddDelAsReply{})
	err = lbHandler.AddLBAs("1.2.3.4/24", 6, 80, "10.1.1.1")
	Expect(err).ShouldNot(HaveOccurred())

	msg, ok = ctx.MockChannel.Msg.(*lbba.LbAddDelAs)
	Expect(ok).To(BeTrue())
	Expect(msg).ToNot(BeNil())
	Expect(msg.Pfx).To(Equal(toPrefix("1.2.3.4/24")))
	Expect(msg.Protocol).To(BeEquivalentTo(6))
	Expect(msg.Port).To(BeEquivalentTo(80))
	Expect(msg.AsAddress).To(Equal(ipToAddress("10.1.1.1")))

	ctx.MockVpp.MockReply(&lbba.LbAddDelAsReply{})
	err = lbHandler.AddLBAs("2600:2600:2600:2600::1/64", 6, 80, "2601:2601:2601:2601::1")
	Expect(err).ShouldNot(HaveOccurred())
	_, _ = ctx.MockChannel.Msg.(*lbba.LbAddDelAsReply)

	msg, ok = ctx.MockChannel.Msg.(*lbba.LbAddDelAs)
	Expect(ok).To(BeTrue())
	Expect(msg).ToNot(BeNil())
	Expect(msg.Pfx).To(Equal(toPrefix("2600:2600:2600:2600::1/64")))
	Expect(msg.Protocol).To(BeEquivalentTo(6))
	Expect(msg.Port).To(BeEquivalentTo(80))
	Expect(msg.AsAddress).To(Equal(ipToAddress("2601:2601:2601:2601::1")))

}

func TestAddLBAsError(t *testing.T) {
	ctx, lbHandler, _ := lbTestSetup(t)
	defer ctx.TeardownTestCtx()

	// Invalid arguments
	ctx.MockVpp.MockReply(&lbba.LbAddDelAsReply{})
	err := lbHandler.AddLBAs("1.2.3.4/24", 6, 80, "10.1.1")
	Expect(err).Should(HaveOccurred())
}

func TestAddLBAsRetval(t *testing.T) {
	ctx, lbHandler, _ := lbTestSetup(t)
	defer ctx.TeardownTestCtx()

	ctx.MockVpp.MockReply(&lbba.LbAddDelAsReply{
		Retval: 1,
	})
	err := lbHandler.AddLBAs("1.2.3.4/24", 6, 80, "10.1.1")
	Expect(err).Should(HaveOccurred())

}

func TestLBConf(t *testing.T) {
	var b [16]byte
	var c [4]byte
	ctx, lbHandler, _ := lbTestSetup(t)
	defer ctx.TeardownTestCtx()

	ctx.MockVpp.MockReply(&lbba.LbConfReply{})
	err := lbHandler.SetLBConf("2600:2600:2600:2600::1", 4, 4)
	Expect(err).ShouldNot(HaveOccurred())

	msg, ok := ctx.MockChannel.Msg.(*lbba.LbConf)
	Expect(ok).To(BeTrue())
	Expect(msg).ToNot(BeNil())

	i := ipToAddress("2600:2600:2600:2600::1")
	copy(b[0:16], i.Un.XXX_UnionData[0:16])
	Expect(msg.IP6SrcAddress).To(BeEquivalentTo(b))
	Expect(msg.StickyBucketsPerCore).To(BeEquivalentTo(4))
	Expect(msg.FlowTimeout).To(BeEquivalentTo(4))

	ctx.MockVpp.MockReply(&lbba.LbConfReply{})
	err = lbHandler.SetLBConf("1.2.3.4", 4, 4)
	Expect(err).ShouldNot(HaveOccurred())

	msg, ok = ctx.MockChannel.Msg.(*lbba.LbConf)
	Expect(ok).To(BeTrue())
	Expect(msg).ToNot(BeNil())

	i = ipToAddress("1.2.3.4")
	copy(c[0:4], i.Un.XXX_UnionData[12:16])
	Expect(msg.IP4SrcAddress).To(BeEquivalentTo(c))
	Expect(msg.StickyBucketsPerCore).To(BeEquivalentTo(4))
	Expect(msg.FlowTimeout).To(BeEquivalentTo(4))
}

func TestAddLBConfError(t *testing.T) {
	ctx, lbHandler, _ := lbTestSetup(t)
	defer ctx.TeardownTestCtx()

	// Invalid arguments
	ctx.MockVpp.MockReply(&lbba.LbConfReply{})
	err := lbHandler.SetLBConf("1.1.1.", 4, 4)
	Expect(err).Should(HaveOccurred())

}

func TestAddLBConfRetval(t *testing.T) {
	ctx, lbHandler, _ := lbTestSetup(t)
	defer ctx.TeardownTestCtx()

	ctx.MockVpp.MockReply(&lbba.LbConfReply{
		Retval: 1,
	})
	err := lbHandler.SetLBConf("1.1.1.1", 4, 4)
	Expect(err).Should(HaveOccurred())

}
