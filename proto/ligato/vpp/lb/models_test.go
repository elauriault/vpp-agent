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

/*func TestLBKey(t *testing.T) {
	tests := []struct {
		name         string
		lbInterface string
		lbIP        string
		expectedKey  string
	}{
		{
			name:         "valid LB case",
			lbInterface: "if1",
			lbIP:        "10.0.0.1",
			expectedKey:  "vpp/config/v2/lb/rule/if1/ip/10.0.0.1",
		},
		{
			name:         "invalid LB case (undefined interface)",
			lbInterface: "",
			lbIP:        "10.0.0.1",
			expectedKey:  "vpp/config/v2/lb/rule/<invalid>/ip/10.0.0.1",
		},
		{
			name:         "invalid LB case (undefined address)",
			lbInterface: "if1",
			lbIP:        "",
			expectedKey:  "vpp/config/v2/lb/rule/if1/ip/<invalid>",
		},
		{
			name:         "invalid LB case (IP address with mask provided)",
			lbInterface: "if1",
			lbIP:        "10.0.0.1/24",
			expectedKey:  "vpp/config/v2/lb/rule/if1/ip/<invalid>",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			key := lb.Key(test.lbInterface, test.lbIP)
			if key != test.expectedKey {
				t.Errorf("failed for: lbName=%s\n"+
					"expected key:\n\t%q\ngot key:\n\t%q",
					test.name, test.expectedKey, key)
			}
		})
	}
}*/

/*func TestParseLBKey(t *testing.T) {
	tests := []struct {
		name             string
		key              string
		expectedIfName   string
		expectedIP       string
		expectedIsLBKey bool
	}{
		{
			name:             "valid LB key",
			key:              "vpp/config/v2/lb/rule/if1/ip/10.0.0.1",
			expectedIfName:   "if1",
			expectedIP:       "10.0.0.1",
			expectedIsLBKey: true,
		},
		{
			name:             "invalid if",
			key:              "vpp/config/v2/lb/rule/<invalid>/ip/10.0.0.1",
			expectedIfName:   "<invalid>",
			expectedIP:       "10.0.0.1",
			expectedIsLBKey: true,
		},
		{
			name:             "invalid LB",
			key:              "vpp/config/v2/lb/rule/if1/ip/<invalid>",
			expectedIfName:   "if1",
			expectedIP:       "<invalid>",
			expectedIsLBKey: true,
		},
		{
			name:             "invalid all",
			key:              "vpp/config/v2/lb/rule/<invalid>/ip/<invalid>",
			expectedIfName:   "<invalid>",
			expectedIP:       "<invalid>",
			expectedIsLBKey: true,
		},
		{
			name:             "not LB key",
			key:              "vpp/config/v2/bd/bd1",
			expectedIfName:   "",
			expectedIP:       "",
			expectedIsLBKey: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ifName, ip, isLBKey := lb.ParseKey(test.key)
			if isLBKey != test.expectedIsLBKey {
				t.Errorf("expected isFIBKey: %v\tgot: %v", test.expectedIsLBKey, isLBKey)
			}
			if ifName != test.expectedIfName {
				t.Errorf("expected ifName: %s\tgot: %s", test.expectedIfName, ifName)
			}
			if ip != test.expectedIP {
				t.Errorf("expected IP: %s\tgot: %s", test.expectedIP, ip)
			}
		})
	}
}*/
