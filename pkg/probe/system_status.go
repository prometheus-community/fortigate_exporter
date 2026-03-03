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
	"fmt"
	"log"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus-community/fortigate_exporter/pkg/http"
)

func probeSystemStatus(c http.FortiHTTP, _ *TargetMetadata) ([]prometheus.Metric, bool) {
	mVersion := prometheus.NewDesc(
		"fortigate_version_info",
		"System version and build information",
		[]string{"serial", "version", "build", "model_name", "model_number", "model", "hostname"}, nil,
	)
	mLogDiskState := prometheus.NewDesc(
		"fortigate_system_status_log_disk_state",
		"System log disk availability state",
		[]string{"state"}, nil,
	)

	type systemResult struct {
		Name          string `json:"model_name"`
		Number        string `json:"model_number"`
		Model         string `json:"model"`
		Hostname      string `json:"hostname"`
		LogDiskStatus string `json:"log_disk_status"`
	}

	type systemStatus struct {
		Status  string
		Serial  string
		Version string
		Build   int64
		Results systemResult
	}
	var st systemStatus

	if err := c.Get("api/v2/monitor/system/status", "", &st); err != nil {
		log.Printf("Error: %v", err)
		return nil, false
	}

	m := []prometheus.Metric{
		prometheus.MustNewConstMetric(mLogDiskState, prometheus.GaugeValue, 0.0, "available"),
		prometheus.MustNewConstMetric(mLogDiskState, prometheus.GaugeValue, 0.0, "need_format"),
		prometheus.MustNewConstMetric(mLogDiskState, prometheus.GaugeValue, 0.0, "not_available"),
	}
	switch st.Results.LogDiskStatus {
	case "available":
		m[0] = prometheus.MustNewConstMetric(mLogDiskState, prometheus.GaugeValue, 1.0, st.Results.LogDiskStatus)
	case "need_format":
		m[1] = prometheus.MustNewConstMetric(mLogDiskState, prometheus.GaugeValue, 1.0, st.Results.LogDiskStatus)
	case "not_available":
		m[2] = prometheus.MustNewConstMetric(mLogDiskState, prometheus.GaugeValue, 1.0, st.Results.LogDiskStatus)
	}
	m = append(m, prometheus.MustNewConstMetric(mVersion, prometheus.GaugeValue, 1.0, st.Serial, st.Version, fmt.Sprintf("%d", st.Build), st.Results.Name, st.Results.Number, st.Results.Model, st.Results.Hostname))
	return m, true
}
