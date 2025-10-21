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
	"log"

	"github.com/prometheus-community/fortigate_exporter/pkg/http"
	"github.com/prometheus/client_golang/prometheus"
)

func probeSystemSandboxConnection(c http.FortiHTTP, meta *TargetMetadata) ([]prometheus.Metric, bool) {
	var (
		connectionStatus = prometheus.NewDesc(
			"fortigate_sandbox_connection_status",
			"Sandbox connection status, (unreachable=0, reachable=1, disabled=-1)",
			[]string{"type"}, nil,
		)
	)

	type SystemSandboxConnection struct {
		Status string `json:"status"`
		Type   string `json:"type"`
	}

	type SystemSandboxConnectionResult struct {
		Results []SystemSandboxConnection `json:"results"`
	}
	var res SystemSandboxConnectionResult
	if err := c.Get("api/v2/monitor/system/sandbox/connection","", &res); err != nil {
		log.Printf("Warning: %v", err)
		return nil, false
	}

	m := []prometheus.Metric{}
	for _, r := range res.Results {
	switch r.Status {
		case "unreachable":
			m = append(m, prometheus.MustNewConstMetric(connectionStatus, prometheus.GaugeValue, 0, r.Type))
		case "reachable":
			m = append(m, prometheus.MustNewConstMetric(connectionStatus, prometheus.GaugeValue, 1, r.Type))
		case "disabled":
			m = append(m, prometheus.MustNewConstMetric(connectionStatus, prometheus.GaugeValue, -1, r.Type))
		}
	}
	return m, true
}