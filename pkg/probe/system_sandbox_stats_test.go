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
	# HELP fortigate_sandbox_stats_clean_total Number of clean files
	# TYPE fortigate_sandbox_stats_clean_total counter
	fortigate_sandbox_stats_clean_total 45120
	# HELP fortigate_sandbox_stats_detected_total Number of detected files
	# TYPE fortigate_sandbox_stats_detected_total counter
	fortigate_sandbox_stats_detected_total 10
	# HELP fortigate_sandbox_stats_risk_high_total Number of high risk files detected
	# TYPE fortigate_sandbox_stats_risk_high_total counter
	fortigate_sandbox_stats_risk_high_total 5
	# HELP fortigate_sandbox_stats_risk_low_total Number of low risk files detected
	# TYPE fortigate_sandbox_stats_risk_low_total counter
	fortigate_sandbox_stats_risk_low_total 3
	# HELP fortigate_sandbox_stats_risk_medium_total Number of medium risk files detected
	# TYPE fortigate_sandbox_stats_risk_medium_total counter
	fortigate_sandbox_stats_risk_medium_total 2
	# HELP fortigate_sandbox_stats_submitted_total Number of submitted files
	# TYPE fortigate_sandbox_stats_submitted_total counter
	fortigate_sandbox_stats_submitted_total 45130
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}
