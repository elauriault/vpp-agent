//  Copyright (c) 2018 Cisco and/or its affiliates.
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

package vpp_lb_test

import (
	"testing"

	lb "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/lb"
)

func TestLBVipKey(t *testing.T) {
	tests := []struct {
		name        string
		prefix      string
		protocol    uint32
		port        uint32
		expectedKey string
	}{
		{
			name:        "valid VIP case: no protocol/port",
			prefix:      "10.0.0.1",
			protocol:    0,
			port:        0,
			expectedKey: "config/vpp/lb/v2/lb-vip/10.0.0.1",
		},
		{
			name:        "valid VIP case: protocol/port",
			prefix:      "10.0.0.1",
			protocol:    6,
			port:        80,
			expectedKey: "config/vpp/lb/v2/lb-vip/10.0.0.1/6/80",
		},
		{
			name:        "invalid VIP case: no protocol/port",
			prefix:      "<invalid>",
			protocol:    0,
			port:        0,
			expectedKey: "config/vpp/lb/v2/lb-vip/<invalid>",
		},
		{
			name:        "invalid VIP case: protocol/port",
			prefix:      "10.0.0.1",
			protocol:    0,
			port:        80,
			expectedKey: "config/vpp/lb/v2/lb-vip/10.0.0.1",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			key := lb.LBVipKey(test.prefix, test.protocol, test.port)
			if key != test.expectedKey {
				t.Errorf("failed for: lbName=%s\n"+
					"expected key:\n\t%q\ngot key:\n\t%q",
					test.name, test.expectedKey, key)
			}
		})
	}
}

func TestLBAsKey(t *testing.T) {
	tests := []struct {
		name        string
		prefix      string
		protocol    uint32
		port        uint32
		address     string
		expectedKey string
	}{
		{
			name:        "valid AS case: no protocol/port",
			prefix:      "10.0.0.1",
			protocol:    0,
			port:        0,
			address:     "1.2.3.4",
			expectedKey: "config/vpp/lb/v2/lb-as/10.0.0.1/1.2.3.4",
		},
		{
			name:        "valid AS case: protocol/port",
			prefix:      "10.0.0.1",
			protocol:    6,
			port:        80,
			address:     "1.2.3.4",
			expectedKey: "config/vpp/lb/v2/lb-as/10.0.0.1/6/80/1.2.3.4",
		},
		{
			name:        "invalid AS case: bad prefix",
			prefix:      "<invalid>",
			protocol:    0,
			port:        0,
			address:     "1.2.3.4",
			expectedKey: "config/vpp/lb/v2/lb-as/<invalid>/1.2.3.4",
		},
		{
			name:        "invalid AS case: bad address protocol/port",
			prefix:      "10.0.0.1",
			protocol:    6,
			port:        80,
			address:     "<invalid>",
			expectedKey: "config/vpp/lb/v2/lb-as/10.0.0.1/6/80/<invalid>",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			key := lb.LBAsKey(test.prefix, test.protocol, test.port, test.address)
			if key != test.expectedKey {
				t.Errorf("failed for: lbName=%s\n"+
					"expected key:\n\t%q\ngot key:\n\t%q",
					test.name, test.expectedKey, key)
			}
		})
	}
}
