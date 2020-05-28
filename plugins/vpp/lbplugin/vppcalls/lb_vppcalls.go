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

package vppcalls

import (
	"errors"

	govppapi "git.fd.io/govpp.git/api"
	"go.ligato.io/cn-infra/v2/logging"

	"go.ligato.io/vpp-agent/v3/plugins/vpp"

	"go.ligato.io/vpp-agent/v3/plugins/vpp/ifplugin/ifaceidx"
	lb "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/lb"
)

var (
	// ErrGlobalConfigDumpNotImplemented is used for LbVppAPI handlers that are missing implementation.
	ErrGlobalConfigDumpNotImplemented = errors.New("global config dump not implemented")
)

// LbVppAPI provides methods for managing LB
type LbVppAPI interface {
	LbVppRead

	SetLBConf(address string, buckets uint32, timeout uint32) error
	AddLBVip(prefix string, protocol uint8, port uint16, encap lb.LBVip_Encap, dscp uint8, srv_type lb.LBVip_SrvType, targetport, nodeport uint16, len uint32) error
	DelLBVip(prefix string, protocol uint8, port uint16) error
	AddLBAs(prefix string, protocol uint8, port uint16, address string) error
	DelLBAs(prefix string, protocol uint8, port uint16, address string) error
}

// LbVppRead provides read methods for LB
type LbVppRead interface {
	DefaultLBGlobalConfig() *lb.LBGlobal
	LBGlobalConfigDump() (*lb.LBGlobal, error)
	LBVipDump() ([]*lb.LBVip, error)
	LBAsDump() ([]*lb.LBAs, error)
}

var handler = vpp.RegisterHandler(vpp.HandlerDesc{
	Name:       "lb",
	HandlerAPI: (*LbVppAPI)(nil),
})

type NewHandlerFunc func(ch govppapi.Channel, ifIdx ifaceidx.IfaceMetadataIndex, log logging.Logger) LbVppAPI

func AddLbHandlerVersion(version vpp.Version, msgs []govppapi.Message, h NewHandlerFunc) {
	handler.AddVersion(vpp.HandlerVersion{
		Version: version,
		Check: func(c vpp.Client) error {
			ch, err := c.NewAPIChannel()
			if err != nil {
				return err
			}
			return ch.CheckCompatiblity(msgs...)
		},
		NewHandler: func(c vpp.Client, a ...interface{}) vpp.HandlerAPI {
			ch, err := c.NewAPIChannel()
			if err != nil {
				return err
			}
			return h(ch, a[0].(ifaceidx.IfaceMetadataIndex), a[1].(logging.Logger))
		},
	})
}

func CompatibleLbVppHandler(c vpp.Client, ifIdx ifaceidx.IfaceMetadataIndex, log logging.Logger) LbVppAPI {
	if v := handler.FindCompatibleVersion(c); v != nil {
		return v.NewHandler(c, ifIdx, log).(LbVppAPI)
	}
	return nil
}
