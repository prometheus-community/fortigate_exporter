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

func probeSystemSandboxStatus(c http.FortiHTTP, meta *TargetMetadata) ([]prometheus.Metric, bool) {
	var (
		signatureCount = prometheus.NewDesc(
			"fortigate_sandbox_status_signature_count",
			"Sandbox signature counts",
			[]string{"server", "region", "version", "type"}, nil,
		)
	)

	type SystemSandboxStatus struct {
		Server  string  `json:"server"`
		Type    string  `json:"type"`
		Region  string  `json:"cloud_region"`
		Version string  `json:"malware_package_version"`
		Count   float64 `json:"signatures_count"`
	}

	type SystemSandboxStatusResult struct {
		Results []SystemSandboxStatus `json:"results"`
	}

	var res SystemSandboxStatusResult
	if err := c.Get("api/v2/monitor/system/sandbox/status","", &res); err != nil {
		log.Printf("Warning: %v", err)
		return nil, false
	}

	m := []prometheus.Metric{}
	for _, r := range res.Results {
		m = append(m, prometheus.MustNewConstMetric(signatureCount, prometheus.GaugeValue, r.Count, r.Server, r.Region, r.Version, r.Type))
	}

	return m, true
}