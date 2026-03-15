// Copyright The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package probe

import (
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestSystemInterfaceTransceivers(t *testing.T) {
	c := newFakeClient()
	c.prepare("api/v2/monitor/system/interface/transceivers", "testdata/interface-transceivers.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !testProbe(probeSystemInterfaceTransceivers, c, r) {
		t.Errorf("probeSystemInterfaceTransceivers() returned non-success")
	}

	em := `
	# HELP fortigate_inteface_transceivers_info Interface transceivers information
	# TYPE fortigate_inteface_transceivers_info gauge
	fortigate_inteface_transceivers_info{interface="ha1",type="SFP/SFP+/SFP28",vendor="FORTINET",vendorpartnumber="FTL",vendorserialnumber="U00000"} 1
	fortigate_inteface_transceivers_info{interface="ha2",type="SFP/SFP+/SFP28",vendor="FORTINET",vendorpartnumber="FTL",vendorserialnumber="U00000"} 1
	fortigate_inteface_transceivers_info{interface="port33",type="QSFP/QSFP+",vendor="FORTINET",vendorpartnumber="FTL",vendorserialnumber="U00000"} 1
	fortigate_inteface_transceivers_info{interface="port34",type="QSFP/QSFP+",vendor="FORTINET",vendorpartnumber="FTL",vendorserialnumber="U00000"} 1
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}
