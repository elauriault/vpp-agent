// Copyright (c) 2020 Ubisoft Entertainment and/or its affiliates.
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

//go:generate descriptor-adapter --descriptor-name LBGlobal --value-type *vpp_lb.LBGlobal --import "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/lb" --output-dir "descriptor"
//go:generate descriptor-adapter --descriptor-name LBVip --value-type *vpp_lb.LBVip --import "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/lb" --output-dir "descriptor"
//go:generate descriptor-adapter --descriptor-name LBAs --value-type *vpp_lb.LBAs --import "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/lb" --output-dir "descriptor"

package lbplugin

import (
	"github.com/pkg/errors"
	"go.ligato.io/cn-infra/v2/health/statuscheck"
	"go.ligato.io/cn-infra/v2/infra"

	"go.ligato.io/vpp-agent/v3/plugins/govppmux"
	"go.ligato.io/vpp-agent/v3/plugins/kvscheduler"
	kvs "go.ligato.io/vpp-agent/v3/plugins/kvscheduler/api"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/ifplugin"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/lbplugin/descriptor"
	"go.ligato.io/vpp-agent/v3/plugins/vpp/lbplugin/vppcalls"

	_ "go.ligato.io/vpp-agent/v3/plugins/vpp/lbplugin/vppcalls/vpp2001"
)

func init() {
	kvscheduler.AddNonRetryableError(vppcalls.ErrGlobalConfigDumpNotImplemented)
}

// LBPlugin configures VPP LB rules using GoVPP.
type LBPlugin struct {
	Deps

	// handlers
	lbHandler vppcalls.LbVppAPI

	// descriptors
	lbDescriptor *descriptor.LBGlobalDescriptor
}

// Deps lists dependencies of the LB plugin.
type Deps struct {
	infra.PluginDeps
	KVScheduler kvs.KVScheduler
	VPP         govppmux.API
	IfPlugin    ifplugin.API
	StatusCheck statuscheck.PluginStatusWriter // optional
}

// Init registers LB-related descriptors.
func (p *LBPlugin) Init() (err error) {
	if !p.VPP.IsPluginLoaded("lb") {
		p.Log.Warnf("VPP plugin LB was disabled by VPP")
		return nil
	}

	// init handlers
	p.lbHandler = vppcalls.CompatibleLbVppHandler(p.VPP, p.IfPlugin.GetInterfaceIndex(), p.Log)
	if p.lbHandler == nil {
		return errors.New("lbHandler is not available")
	}

	// init and register LB descriptor
	lbGlobalDescriptor := descriptor.NewLBGlobalDescriptor(p.lbHandler, p.Log)
	lbVipsDescriptor := descriptor.NewLBVipDescriptor(p.lbHandler, p.Log)
	lbAsDescriptor := descriptor.NewLBAsDescriptor(p.lbHandler, p.Log)
	p.Log.Warnf("{}", lbGlobalDescriptor.NBKeyPrefix)
	err = p.Deps.KVScheduler.RegisterKVDescriptor(
		lbGlobalDescriptor,
		lbVipsDescriptor,
		lbAsDescriptor,
	)
	if err != nil {
		return err
	}

	return nil
}

// AfterInit registers plugin with StatusCheck.
func (p *LBPlugin) AfterInit() error {
	if p.StatusCheck != nil {
		p.StatusCheck.Register(p.PluginName, nil)
	}
	return nil
}
