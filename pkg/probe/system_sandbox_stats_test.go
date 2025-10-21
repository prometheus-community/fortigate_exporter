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

func TestSystemSandboxStats(t *testing.T) {
	c := newFakeClient()
	c.prepare("api/v2/monitor/system/sandbox/stats", "testdata/system-sandbox-stats.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !testProbe(probeSystemSandboxStats, c, r) {
		t.Errorf("probeSystemSandboxStats() returned non-success")
	}

	em := `
	# HELP fortigate_sandbox_stats_clean Number of clean files
	# TYPE fortigate_sandbox_stats_clean counter
	fortigate_sandbox_stats_clean 45120
	# HELP fortigate_sandbox_stats_detected Number of detected files
	# TYPE fortigate_sandbox_stats_detected counter
	fortigate_sandbox_stats_detected 10
	# HELP fortigate_sandbox_stats_risk_high Number of high risk files detected
	# TYPE fortigate_sandbox_stats_risk_high counter
	fortigate_sandbox_stats_risk_high 5
	# HELP fortigate_sandbox_stats_risk_low Number of low risk files detected
	# TYPE fortigate_sandbox_stats_risk_low counter
	fortigate_sandbox_stats_risk_low 3
	# HELP fortigate_sandbox_stats_risk_medium Number of medium risk files detected
	# TYPE fortigate_sandbox_stats_risk_medium counter
	fortigate_sandbox_stats_risk_medium 2
	# HELP fortigate_sandbox_stats_submitted Number of submitted files
	# TYPE fortigate_sandbox_stats_submitted counter
	fortigate_sandbox_stats_submitted 45130
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}