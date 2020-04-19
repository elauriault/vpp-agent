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

package vpp_lb

import (
	"go.ligato.io/vpp-agent/v3/pkg/models"
)

// ModuleName is the module name used for models.
const ModuleName = "vpp.lb"

var (
	ModelLBGlobal = models.Register(&LBGlobal{}, models.Spec{
		Module:  ModuleName,
		Type:    "lb-global",
		Version: "v2",
	})

	ModelLBVip = models.Register(&LBVip{}, models.Spec{
		Module:  ModuleName,
		Type:    "lb-vip",
		Version: "v2",
	}, models.WithNameTemplate(
		"{{.Prefix}}"+
			"{{if and .Port .Protocol}}/{{.Protocol}}/{{.Port}}{{end}}",
	))

	ModelLBAs = models.Register(&LBAs{}, models.Spec{
		Module:  ModuleName,
		Type:    "lb-as",
		Version: "v2",
	}, models.WithNameTemplate(
		"{{.Prefix}}"+
			"{{if and .Port .Protocol}}/{{.Protocol}}/{{.Port}}{{end}}"+
			"/{{.Address}}",
	))
)

func GlobalLBKey() string {
	return models.Key(&LBGlobal{})
}

func LBVipKey(prefix string, protocol, port uint32) string {
	return models.Key(&LBVip{
		Prefix:   prefix,
		Protocol: protocol,
		Port:     port,
	})
}

func LBAsKey(prefix string, protocol, port uint32, address string) string {
	return models.Key(&LBAs{
		Prefix:   prefix,
		Protocol: protocol,
		Port:     port,
		Address:  address,
	})
}
