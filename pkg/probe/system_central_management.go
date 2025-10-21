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
	"strconv"

	"github.com/prometheus-community/fortigate_exporter/pkg/http"
	"github.com/prometheus/client_golang/prometheus"
)

func probeSystemCentralManagementStatus (c http.FortiHTTP, meta *TargetMetadata) ([]prometheus.Metric, bool) {
	var (
		mode = prometheus.NewDesc(
			"fortigate_system_central_management_mode",
			"Operating mode of the central management. (Normal = 1, Backup = 2)",
			[]string{"server", "mgmt_ip", "mgmt_port", "sn", "pendfortman"}, nil,
		)
		status = prometheus.NewDesc(
			"fortigate_system_central_management_status",
			"Status of the connection from FortiGate to the central management server. (down = 0, up = 1, handshake = 2)",
			[]string{"server", "mgmt_ip", "mgmt_port", "sn", "pendfortman"}, nil,
		)
		registration_status = prometheus.NewDesc(
			"fortigate_system_central_management_registration_status",
			"Status of the registration from FortiGate to the central management server. (unknown = -1, in_progress = 1, registered = 0, unregistered = 2)",
			[]string{"server", "mgmt_ip", "mgmt_port", "sn", "pendfortman"}, nil,
		)
	)

	type centralManagementStatus struct {
		Mode       string  `json:"mode"`
		Server     string  `json:"server"`
		Status     string  `json:"status"`
		RegStat    string  `json:"registration_status"`
		MgmtIp     string  `json:"mgmt_ip"`
		MgmtPort   float64 `json:"mgmt_port"`
		Sn         string  `json:"sn"`
		PenFortMan string  `json:"pending_fortimanager"`
	}

	type centralManagementStatusResult struct {
		Result centralManagementStatus `json:"results"`
	}

	var res centralManagementStatusResult
	if err := c.Get("api/v2/monitor/system/central-management/status", "skip_detect=true", &res); err != nil {
		log.Printf("Error: %v", err)
		return nil, false
	}

	m := []prometheus.Metric{}
    if res.Result.Mode == "normal" {
		m = append(m, prometheus.MustNewConstMetric(mode, prometheus.GaugeValue, 1, res.Result.Server, res.Result.MgmtIp, strconv.FormatFloat(res.Result.MgmtPort, 'f', -1, 64), res.Result.Sn, res.Result.PenFortMan))
	} else {
		m = append(m, prometheus.MustNewConstMetric(mode, prometheus.GaugeValue, 2, res.Result.Server, res.Result.MgmtIp, strconv.FormatFloat(res.Result.MgmtPort, 'f', -1, 64), res.Result.Sn, res.Result.PenFortMan))
	}
	switch res.Result.Status {
	case "down":
		m = append(m, prometheus.MustNewConstMetric(status, prometheus.GaugeValue, 0, res.Result.Server, res.Result.MgmtIp, strconv.FormatFloat(res.Result.MgmtPort, 'f', -1, 64), res.Result.Sn, res.Result.PenFortMan))
	case "up":
		m = append(m, prometheus.MustNewConstMetric(status, prometheus.GaugeValue, 1, res.Result.Server, res.Result.MgmtIp, strconv.FormatFloat(res.Result.MgmtPort, 'f', -1, 64), res.Result.Sn, res.Result.PenFortMan))
	case "handshake":
		m = append(m, prometheus.MustNewConstMetric(status, prometheus.GaugeValue, 2, res.Result.Server, res.Result.MgmtIp, strconv.FormatFloat(res.Result.MgmtPort, 'f', -1, 64), res.Result.Sn, res.Result.PenFortMan))
	}
	switch res.Result.RegStat {
	case "in_progress":
		m = append(m, prometheus.MustNewConstMetric(registration_status, prometheus.GaugeValue, 1, res.Result.Server, res.Result.MgmtIp, strconv.FormatFloat(res.Result.MgmtPort, 'f', -1, 64), res.Result.Sn, res.Result.PenFortMan))
	case "registered":
		m = append(m, prometheus.MustNewConstMetric(registration_status, prometheus.GaugeValue, 0, res.Result.Server, res.Result.MgmtIp, strconv.FormatFloat(res.Result.MgmtPort, 'f', -1, 64), res.Result.Sn, res.Result.PenFortMan))
	case "unregistered":
		m = append(m, prometheus.MustNewConstMetric(registration_status, prometheus.GaugeValue, 2, res.Result.Server, res.Result.MgmtIp, strconv.FormatFloat(res.Result.MgmtPort, 'f', -1, 64), res.Result.Sn, res.Result.PenFortMan))
	default:
		m = append(m, prometheus.MustNewConstMetric(registration_status, prometheus.GaugeValue, -1, res.Result.Server, res.Result.MgmtIp, strconv.FormatFloat(res.Result.MgmtPort, 'f', -1, 64), res.Result.Sn, res.Result.PenFortMan))
	}
	return m, true
}