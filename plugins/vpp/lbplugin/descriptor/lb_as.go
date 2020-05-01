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
	LBAsDescriptorName = "vpp-lb-as"
)

// A list of non-retriable errors:
var (
	// ErrSrcUndefined = errors.New("Neither Ip4SrcAddress or Ip6SrcAddress is defined")
	// ErrSrcAmbiguous = errors.New("Both Ip4SrcAddress and Ip6SrcAddress are defined")
	ErrAS4 = errors.New("v4 AS requires a v4 VIP")
	ErrAS6 = errors.New("v6 AS requires a v6 VIP")
)

// LBAsDescriptor teaches KVScheduler how to configure global options for
// VPP LB.
type LBAsDescriptor struct {
	log       logging.Logger
	lbHandler vppcalls.LbVppAPI
}

// NewLBAsDescriptor creates a new instance of the LBGlobal descriptor.
func NewLBAsDescriptor(lbHandler vppcalls.LbVppAPI, log logging.PluginLogger) *kvs.KVDescriptor {
	ctx := &LBAsDescriptor{
		lbHandler: lbHandler,
		log:       log.NewLogger("lb-as-descriptor"),
	}
	// ctx.log.Warnf("DEBUG_STUFF : NewLBAsDescriptor")

	typedDescr := &adapter.LBAsDescriptor{
		Name:          LBAsDescriptorName,
		NBKeyPrefix:   lb.ModelLBAs.KeyPrefix(),
		ValueTypeName: lb.ModelLBAs.ProtoName(),
		KeySelector:   lb.ModelLBAs.IsKeyValid,
		Validate:      ctx.Validate,
		Create:        ctx.Create,
		Delete:        ctx.Delete,
		Retrieve:      ctx.Retrieve,
		Dependencies:  ctx.Dependencies,
		RetrieveDependencies: []string{
			// refresh the pool of allocated IP addresses first
			vpp_ifdescriptor.InterfaceDescriptorName,
			LBVipDescriptorName},
	}

	return adapter.NewLBAsDescriptor(typedDescr)
}

// Validate validates AS configuration.
func (d *LBAsDescriptor) Validate(key string, as *lb.LBAs) error {
	// d.log.Warnf("DEBUG_STUFF : Validate %v %v", key, vip)

	switch p := as.GetProtocol(); p {
	case 0:
		if as.GetPort() != 0 {
			return ErrProtocol
		}
	case 6:
		if as.GetPort() == 0 {
			return ErrPort
		}
	case 17:
		if as.GetPort() == 0 {
			return ErrPort
		}
	}

	if is_v4(as.GetPrefix()) {
		if !is_v4(as.GetAddress()) {
			return ErrAS4
		}
	}
	if is_v6(as.GetPrefix()) {
		if !is_v6(as.GetAddress()) {
			return ErrAS6
		}
	}

	return nil

}

// Create adds new AS
func (d *LBAsDescriptor) Create(key string, as *lb.LBAs) (metadata interface{}, err error) {
	// d.log.Warnf("DEBUG_STUFF : Create %v %v", key, as)
	d.lbHandler.AddLBAs(as.GetPrefix(), uint8(as.GetProtocol()), uint16(as.GetPort()), as.GetAddress())
	return nil, nil
}

// Delete removed AS
func (d *LBAsDescriptor) Delete(key string, as *lb.LBAs, metadata interface{}) error {
	// d.log.Warnf("DEBUG_STUFF : Delete %v %v", key, as)
	d.lbHandler.DelLBAs(as.GetPrefix(), uint8(as.GetProtocol()), uint16(as.GetPort()), as.GetAddress())
	return nil
}

// Retrieve returns the current AS.

func (d *LBAsDescriptor) Retrieve(correlate []adapter.LBAsKVWithMetadata) (retrieved []adapter.LBAsKVWithMetadata, err error) {
	// d.log.Warnf("DEBUG_STUFF : Retrieve As")
	as_list, err := d.lbHandler.LBAsDump()
	if err != nil {
		return nil, err
	}
	for _, as := range as_list {
		retrieved = append(retrieved, adapter.LBAsKVWithMetadata{
			Key:    lb.LBAsKey(as.GetPrefix(), as.GetProtocol(), as.GetPort(), as.GetAddress()),
			Value:  as,
			Origin: kvs.FromNB,
		})
		// d.log.Warnf("DEBUG_STUFF : Inside Retrieve As : %v", as)
	}
	return
}

func (d *LBAsDescriptor) Dependencies(key string, as *lb.LBAs) []kvs.Dependency {
	return []kvs.Dependency{
		{
			Label: lbVipDependency,
			Key:   lb.LBVipKey(as.GetPrefix(), as.GetProtocol(), as.GetPort()),
		},
	}
}
