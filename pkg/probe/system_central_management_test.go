// Copyright 2025 The Prometheus Authors
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

func TestSystemCentralManagementStatus(t *testing.T) {
	c := newFakeClient()
	c.prepare("api/v2/monitor/system/central-management/status", "testdata/system-central-management-status.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !testProbe(probeSystemCentralManagementStatus, c, r) {
		t.Errorf("probeSystemCentralManagementStatus() returned non-success")
	}

	em := `
	# HELP fortigate_system_central_management_mode Operating mode of the central management. (Normal = 1, Backup = 2)
	# TYPE fortigate_system_central_management_mode gauge
	fortigate_system_central_management_mode{mgmt_ip="127.0.0.1",mgmt_port="0",pendfortman="12.329845.45k3",server="HA-TEST",sn="121748"} 1
	# HELP fortigate_system_central_management_registration_status Status of the registration from FortiGate to the central management server. (unknown = -1, in_progress = 1, registered = 0, unregistered = 2)
	# TYPE fortigate_system_central_management_registration_status gauge
	fortigate_system_central_management_registration_status{mgmt_ip="127.0.0.1",mgmt_port="0",pendfortman="12.329845.45k3",server="HA-TEST",sn="121748"} -1
	# HELP fortigate_system_central_management_status Status of the connection from FortiGate to the central management server. (down = 0, up = 1, handshake = 2)
	# TYPE fortigate_system_central_management_status gauge
	fortigate_system_central_management_status{mgmt_ip="127.0.0.1",mgmt_port="0",pendfortman="12.329845.45k3",server="HA-TEST",sn="121748"} 0
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}