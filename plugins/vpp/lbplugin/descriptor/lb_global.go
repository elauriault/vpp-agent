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
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"go.ligato.io/cn-infra/v2/logging"

	"go.ligato.io/vpp-agent/v3/pkg/models"
	kvs "go.ligato.io/vpp-agent/v3/plugins/kvscheduler/api"
	vpp_ifdescriptor "go.ligato.io/vpp-agent/v3/plugins/vpp/ifplugin/descriptor"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/lbplugin/descriptor/adapter"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/lbplugin/vppcalls"
	lb "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/lb"
)

const (
	LBGlobalDescriptorName = "vpp-lb-global"
)

// A list of non-retriable errors:
var (
	ErrSrcUndefined = errors.New("Neither Ip4SrcAddress or Ip6SrcAddress is defined")
	ErrSrcAmbiguous = errors.New("Both Ip4SrcAddress and Ip6SrcAddress are defined")
)

// LBGlobalDescriptor teaches KVScheduler how to configure global options for
// VPP LB.
type LBGlobalDescriptor struct {
	log       logging.Logger
	lbHandler vppcalls.LbVppAPI

	defaultGlobalCfg *lb.LBGlobal
}

// NewLBGlobalDescriptor creates a new instance of the LBGlobal descriptor.
func NewLBGlobalDescriptor(lbHandler vppcalls.LbVppAPI, log logging.PluginLogger) *kvs.KVDescriptor {
	ctx := &LBGlobalDescriptor{
		lbHandler:        lbHandler,
		log:              log.NewLogger("lb-global-descriptor"),
		defaultGlobalCfg: lbHandler.DefaultLBGlobalConfig(),
	}
	// ctx.log.Warnf("DEBUG_STUFF : NewLBGlobalDescriptor")

	typedDescr := &adapter.LBGlobalDescriptor{
		Name:                 LBGlobalDescriptorName,
		NBKeyPrefix:          lb.ModelLBGlobal.KeyPrefix(),
		ValueTypeName:        lb.ModelLBGlobal.ProtoName(),
		KeySelector:          lb.ModelLBGlobal.IsKeyValid,
		ValueComparator:      ctx.EquivalentLBGlobal,
		Validate:             ctx.Validate,
		Create:               ctx.Create,
		Delete:               ctx.Delete,
		Update:               ctx.Update,
		Retrieve:             ctx.Retrieve,
		DerivedValues:        ctx.DerivedValues,
		RetrieveDependencies: []string{vpp_ifdescriptor.InterfaceDescriptorName},
	}
	return adapter.NewLBGlobalDescriptor(typedDescr)
}

// EquivalentLBGlobal compares two LB global configs for equality.
func (d *LBGlobalDescriptor) EquivalentLBGlobal(key string, oldGlobalCfg, newGlobalCfg *lb.LBGlobal) bool {
	// d.log.Warnf("DEBUG_STUFF : EquivalentLBGlobal %v %v %v", key, oldGlobalCfg, newGlobalCfg)
	if oldGlobalCfg.GetIp4SrcAddress() != newGlobalCfg.GetIp4SrcAddress() {
		return false
	}
	if oldGlobalCfg.Ip6SrcAddress != newGlobalCfg.Ip6SrcAddress {
		return false
	}
	if oldGlobalCfg.Bucket != newGlobalCfg.Bucket {
		return false
	}
	if oldGlobalCfg.Timeout != newGlobalCfg.Timeout {
		return false
	}

	return true
}

// Validate validates VPP LB global configuration.
func (d *LBGlobalDescriptor) Validate(key string, globalCfg *lb.LBGlobal) error {
	// d.log.Warnf("DEBUG_STUFF : Validate %v %v", key, globalCfg)

	srcUndefined := kvs.NewInvalidValueError(ErrSrcUndefined, "src_address")
	srcAmbiguous := kvs.NewInvalidValueError(ErrSrcAmbiguous, "src_address")
	if globalCfg.Ip4SrcAddress == "" && globalCfg.Ip6SrcAddress == "" {
		return srcUndefined
	}
	if globalCfg.Ip4SrcAddress != "" && globalCfg.Ip6SrcAddress != "" {
		return srcAmbiguous
	}

	return nil
}

// Create applies LB global options.
func (d *LBGlobalDescriptor) Create(key string, globalCfg *lb.LBGlobal) (metadata interface{}, err error) {
	// d.log.Warnf("DEBUG_STUFF : Create %v %v", key, globalCfg)
	return d.Update(key, d.defaultGlobalCfg, globalCfg, nil)
}

// Delete sets LB global options back to the defaults.
func (d *LBGlobalDescriptor) Delete(key string, globalCfg *lb.LBGlobal, metadata interface{}) error {
	// d.log.Warnf("DEBUG_STUFF : Delete %v %v", key, globalCfg)
	_, err := d.Update(key, globalCfg, d.defaultGlobalCfg, metadata)
	return err
}

// Update updates LB global options.
func (d *LBGlobalDescriptor) Update(key string, oldGlobalCfg, newGlobalCfg *lb.LBGlobal, oldMetadata interface{}) (newMetadata interface{}, err error) {
	// d.log.Warnf("DEBUG_STUFF : Update %v %v %v", key, oldGlobalCfg, newGlobalCfg)

	d.lbHandler.SetLBConf(newGlobalCfg.Ip4SrcAddress, newGlobalCfg.Bucket, newGlobalCfg.Timeout)
	// update virtual reassembly for IPv4
	// if !proto.Equal(d.getVirtualReassembly(oldGlobalCfg), d.getVirtualReassembly(newGlobalCfg)) {
	// 	if err = d.lbHandler.SetVirtualReassemblyIPv4(d.getVirtualReassembly(newGlobalCfg)); err != nil {
	// 		err = errors.Errorf("failed to set NAT virtual reassembly for IPv4: %v", err)
	// 		d.log.Error(err)
	// 		return nil, err
	// 	}
	// }

	return nil, nil
}

// Retrieve returns the current NAT44 global configuration.
func (d *LBGlobalDescriptor) Retrieve(correlate []adapter.LBGlobalKVWithMetadata) ([]adapter.LBGlobalKVWithMetadata, error) {
	// d.log.Warnf("DEBUG_STUFF : Retrieve")
	// Note: either this descriptor (deprecated) or separate interface / address pool descriptors
	// can retrieve NAT interfaces / address pools, never both of them. Use correlate to decide.
	// d.UseDeprecatedAPI = false
	// for _, g := range correlate {
	// 	if len(g.Value.NatInterfaces) > 0 || len(g.Value.AddressPool) > 0 {
	// 		d.UseDeprecatedAPI = true
	// 	}
	// }

	globalCfg, err := d.lbHandler.LBGlobalConfigDump()
	if err != nil {
		d.log.Error(err)
		return nil, err
	}

	origin := kvs.FromNB
	if proto.Equal(globalCfg, d.defaultGlobalCfg) {
		origin = kvs.FromSB
	}

	retrieved := []adapter.LBGlobalKVWithMetadata{{
		Key:    models.Key(globalCfg),
		Value:  globalCfg,
		Origin: origin,
	}}

	return retrieved, nil
}

// DerivedValues derives:
//   - nat.NatAddress for every IP address to be added into the NAT address pool,
//   - nat.NatInterface for every interface with assigned NAT configuration.
func (d *LBGlobalDescriptor) DerivedValues(key string, globalCfg *lb.LBGlobal) (derValues []kvs.KeyValuePair) {
	// d.log.Warnf("DEBUG_STUFF : DerivedValues %v %v", key, globalCfg)
	// NAT addresses
	// for _, natAddr := range globalCfg.AddressPool {
	// 	derValues = append(derValues, kvs.KeyValuePair{
	// 		Key:   nat.DerivedAddressNAT44Key(natAddr.Address, natAddr.TwiceNat),
	// 		Value: natAddr,
	// 	})
	// }
	// // NAT interfaces
	// for _, natIface := range globalCfg.NatInterfaces {
	// 	derValues = append(derValues, kvs.KeyValuePair{
	// 		Key:   nat.DerivedInterfaceNAT44Key(natIface.Name, natIface.IsInside),
	// 		Value: natIface,
	// 	})
	// }
	return derValues
}
