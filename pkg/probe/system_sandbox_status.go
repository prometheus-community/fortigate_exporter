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
	"strconv"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus-community/fortigate_exporter/pkg/fortigatehttpclient"
)

func probeSystemSandboxStatus(c fortigatehttpclient.FortiHTTP, _ *TargetMetadata) ([]prometheus.Metric, bool) {
	Count := prometheus.NewDesc(
		"fortigate_system_sandbox_status_signatures_count",
		"The number of signatures that have been loaded on the FortiSandbox.",
		[]string{"configured", "type", "cloud_region", "server", "malware_package_version", "signatures_loaded", "vdom"}, nil,
	)

	type SystemSandboxStatus struct {
		Configured bool    `json:"configured"`
		Type       string  `json:"type"`
		Cloud      string  `json:"cloud_region"`
		Server     string  `json:"server"`
		MPV        string  `json:"malware_package_version"`
		Loaded     bool    `json:"signatures_loaded"`
		Count      float64 `json:"signatures_count"`
	}

	type SystemSandboxStatusResult struct {
		Result SystemSandboxStatus `json:"results"`
		VDOM   string              `json:"vdom"`
	}

	var res SystemSandboxStatusResult
	if err := c.Get("api/v2/monitor/system/sandbox/status", "", &res); err != nil {
		log.Printf("Warning: %v", err)
		return nil, false
	}
	m := []prometheus.Metric{}
	cloud := "null"
	if res.Result.Cloud != "" {
		cloud = res.Result.Cloud
	}
	m = append(m, prometheus.MustNewConstMetric(
		Count,
		prometheus.GaugeValue,
		res.Result.Count,
		strconv.FormatBool(res.Result.Configured),
		res.Result.Type,
		cloud,
		res.Result.Server,
		res.Result.MPV,
		strconv.FormatBool(res.Result.Loaded),
		res.VDOM),
	)
	return m, true
}
