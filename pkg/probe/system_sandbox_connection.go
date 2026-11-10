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
	"log"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus-community/fortigate_exporter/pkg/fortigatehttpclient"
)

func probeSystemSandboxConnection(c fortigatehttpclient.FortiHTTP, _ *TargetMetadata) ([]prometheus.Metric, bool) {
	connectionStatus := prometheus.NewDesc(
		"fortigate_sandbox_connection_state",
		"Sandbox connection status",
		[]string{"sandbox_type", "status"}, nil,
	)

	type SystemSandboxConnection struct {
		Status string `json:"status"`
		Type   string `json:"type"`
	}

	type SystemSandboxConnectionResult struct {
		Results []SystemSandboxConnection `json:"results"`
	}
	var res SystemSandboxConnectionResult
	if err := c.Get("api/v2/monitor/system/sandbox/connection", "", &res); err != nil {
		log.Printf("Warning: %v", err)
		return nil, false
	}

	m := []prometheus.Metric{}
	for _, r := range res.Results {
		t := []prometheus.Metric{
			prometheus.MustNewConstMetric(connectionStatus, prometheus.GaugeValue, 0, r.Type, "unreachable"),
			prometheus.MustNewConstMetric(connectionStatus, prometheus.GaugeValue, 0, r.Type, "reachable"),
			prometheus.MustNewConstMetric(connectionStatus, prometheus.GaugeValue, 0, r.Type, "disabled"),
			prometheus.MustNewConstMetric(connectionStatus, prometheus.GaugeValue, 0, r.Type, "unauthorized"),
			prometheus.MustNewConstMetric(connectionStatus, prometheus.GaugeValue, 0, r.Type, "incompatible"),
			prometheus.MustNewConstMetric(connectionStatus, prometheus.GaugeValue, 0, r.Type, "unverified"),
		}
		switch r.Status {
		case "unreachable":
			t[0] = prometheus.MustNewConstMetric(connectionStatus, prometheus.GaugeValue, 1, r.Type, r.Status)
		case "reachable":
			t[1] = prometheus.MustNewConstMetric(connectionStatus, prometheus.GaugeValue, 1, r.Type, r.Status)
		case "disabled":
			t[2] = prometheus.MustNewConstMetric(connectionStatus, prometheus.GaugeValue, 1, r.Type, r.Status)
		case "unauthorized":
			t[3] = prometheus.MustNewConstMetric(connectionStatus, prometheus.GaugeValue, 1, r.Type, r.Status)
		case "incompatible":
			t[4] = prometheus.MustNewConstMetric(connectionStatus, prometheus.GaugeValue, 1, r.Type, r.Status)
		case "unverified":
			t[5] = prometheus.MustNewConstMetric(connectionStatus, prometheus.GaugeValue, 1, r.Type, r.Status)
		}
		m = append(m, t...)
	}
	return m, true
}
