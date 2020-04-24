// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package descriptor

import (
	"errors"

	"go.ligato.io/cn-infra/v2/logging"

	kvs "go.ligato.io/vpp-agent/v3/plugins/kvscheduler/api"
	vpp_ifdescriptor "go.ligato.io/vpp-agent/v3/plugins/vpp/ifplugin/descriptor"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/lbplugin/descriptor/adapter"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/lbplugin/vppcalls"
	lb "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/lb"
)

const (
	LBVipDescriptorName = "vpp-lb-vip"
	// dependency labels
	lbVipDependency = "lbvip-exists"
)

// A list of non-retriable errors:
var (
	// ErrSrcUndefined = errors.New("Neither Ip4SrcAddress or Ip6SrcAddress is defined")
	// ErrSrcAmbiguous = errors.New("Both Ip4SrcAddress and Ip6SrcAddress are defined")
	ErrProtocol = errors.New("If protocol is 0(undefined), port should also be 0")
	ErrPort     = errors.New("If protocol is tcp(6) or udp(17), port should also not be 0")
	ErrDscp     = errors.New("If encap is L3DSR DSCP should be defined to a non 0 value")
	ErrNat      = errors.New("If encap is NAT4 or NAT6, Srv Type, Nodeport ans Targetport and should be defined to a non zero value")
	ErrLen      = errors.New("Len should be a power of 2")
)

// LBVipDescriptor teaches KVScheduler how to configure global options for
// VPP LB.
type LBVipDescriptor struct {
	log       logging.Logger
	lbHandler vppcalls.LbVppAPI
}

// NewLBVipDescriptor creates a new instance of the LBGlobal descriptor.
func NewLBVipDescriptor(lbHandler vppcalls.LbVppAPI, log logging.PluginLogger) *kvs.KVDescriptor {
	ctx := &LBVipDescriptor{
		lbHandler: lbHandler,
		log:       log.NewLogger("lb-vip-descriptor"),
	}
	// ctx.log.Warnf("DEBUG_STUFF : NewLBVipDescriptor")

	typedDescr := &adapter.LBVipDescriptor{
		Name:                 LBVipDescriptorName,
		NBKeyPrefix:          lb.ModelLBVip.KeyPrefix(),
		ValueTypeName:        lb.ModelLBVip.ProtoName(),
		KeySelector:          lb.ModelLBVip.IsKeyValid,
		Validate:             ctx.Validate,
		Create:               ctx.Create,
		Delete:               ctx.Delete,
		Retrieve:             ctx.Retrieve,
		RetrieveDependencies: []string{vpp_ifdescriptor.InterfaceDescriptorName},
	}
	return adapter.NewLBVipDescriptor(typedDescr)
}

// Validate validates VIP configuration.
func (d *LBVipDescriptor) Validate(key string, vip *lb.LBVip) error {
	// d.log.Warnf("DEBUG_STUFF : Validate %v %v", key, vip)

	switch p := vip.GetProtocol(); p {
	case 0:
		if vip.GetPort() != 0 {
			return ErrProtocol
		}
	case 6:
		if vip.GetPort() == 0 {
			return ErrPort
		}
	case 17:
		if vip.GetPort() == 0 {
			return ErrPort
		}
	}

	switch e := vip.GetEncap(); e {
	case lb.LBVip_GRE4:
	case lb.LBVip_GRE6:
	case lb.LBVip_L3DSR:
		if vip.GetDscp() == 0 {
			return ErrDscp
		}
	case lb.LBVip_NAT4, lb.LBVip_NAT6:
		if vip.GetSrvType() == 0 {
			return ErrNat
		}
		if vip.GetTargetPort() == 0 {
			return ErrNat
		}
		if vip.GetNodePort() == 0 {
			return ErrNat
		}
	}

	len := vip.GetNewLen()
	if len%2 != 0 {
		return ErrLen
	}

	return nil
}

// Create adds LB VIP
func (d *LBVipDescriptor) Create(key string, vip *lb.LBVip) (metadata interface{}, err error) {
	// d.log.Warnf("DEBUG_STUFF : Create %v %v", key, vip)
	d.lbHandler.AddLBVip(vip.GetPrefix(), uint8(vip.GetProtocol()), uint16(vip.GetPort()), vip.GetEncap(), uint8(vip.GetDscp()), vip.GetSrvType(), uint16(vip.GetTargetPort()), uint16(vip.GetNodePort()), vip.GetNewLen())
	return nil, nil
}

// Delete sets LB global options back to the defaults.
func (d *LBVipDescriptor) Delete(key string, vip *lb.LBVip, metadata interface{}) error {
	// d.log.Warnf("DEBUG_STUFF : Delete %v %v", key, vip)
	d.lbHandler.DelLBVip(vip.GetPrefix(), uint8(vip.GetProtocol()), uint16(vip.GetPort()))
	return nil
}

// Retrieve returns the current VIPS.

func (d *LBVipDescriptor) Retrieve(correlate []adapter.LBVipKVWithMetadata) (retrieved []adapter.LBVipKVWithMetadata, err error) {
	// d.log.Warnf("DEBUG_STUFF : Retrieve Vips")
	vips, err := d.lbHandler.LBVipDump()
	if err != nil {
		return nil, err
	}
	for _, vip := range vips {
		retrieved = append(retrieved, adapter.LBVipKVWithMetadata{
			Key:    lb.LBVipKey(vip.GetPrefix(), vip.GetProtocol(), vip.GetPort()),
			Value:  vip,
			Origin: kvs.FromNB,
		})
	}
	return
}
