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

func TestSandboxStatus(t *testing.T) {
	c := newFakeClient()
	c.prepare("api/v2/monitor/system/sandbox/status", "testdata/system-sandbox-status.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !testProbe(probeSystemSandboxStatus, c, r) {
		t.Errorf("probeSystemSandboxStatus() returned non-success")
	}

	em := `
	# HELP fortigate_system_sandbox_status_signatures_count The number of signatures that have been loaded on the FortiSandbox.
	# TYPE fortigate_system_sandbox_status_signatures_count gauge
	fortigate_system_sandbox_status_signatures_count{cloud_region="null",configured="true",malware_package_version="5.125",server="0.0.0.0",signatures_loaded="false",type="appliance",vdom="root"} 0
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}
