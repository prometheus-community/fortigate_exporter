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

	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus-community/fortigate_exporter/pkg/http"
)

func probeSystemCentralManagementStatus(c http.FortiHTTP, _ *TargetMetadata) ([]prometheus.Metric, bool) {
	var (
		mode = prometheus.NewDesc(
			"fortigate_system_central_management_mode",
			"Operating mode of the central management.",
			[]string{"mode", "server", "mgmt_ip", "mgmt_port", "sn", "pendfortman"}, nil,
		)
		status = prometheus.NewDesc(
			"fortigate_system_central_management_status",
			"Status of the connection from FortiGate to the central management server.",
			[]string{"status", "server", "mgmt_ip", "mgmt_port", "sn", "pendfortman"}, nil,
		)
		registrationStatus = prometheus.NewDesc(
			"fortigate_system_central_management_registration_status",
			"Status of the registration from FortiGate to the central management server.",
			[]string{"status", "server", "mgmt_ip", "mgmt_port", "sn", "pendfortman"}, nil,
		)
	)

	type centralManagementStatus struct {
		Mode       string  `json:"mode"`
		Server     string  `json:"server"`
		Status     string  `json:"status"`
		RegStat    string  `json:"registration_status"`
		MgmtIP     string  `json:"mgmt_ip"`
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
	var normal, backup, down, up, handshake, inProgress, registered, unregistered, defaultValue float64
	if res.Result.Mode == "normal" {
		normal = 1
	} else {
		backup = 1
	}
	switch res.Result.Status {
	case "down":
		down = 1
	case "up":
		up = 1
	case "handshake":
		handshake = 1
	}
	switch res.Result.RegStat {
	case "in_progress":
		inProgress = 1
	case "registered":
		registered = 1
	case "unregistered":
		unregistered = 1
	default:
		defaultValue = 1
	}
	m = append(m, prometheus.MustNewConstMetric(mode, prometheus.GaugeValue, normal, "normal", res.Result.Server, res.Result.MgmtIP, strconv.FormatFloat(res.Result.MgmtPort, 'f', -1, 64), res.Result.Sn, res.Result.PenFortMan))
	m = append(m, prometheus.MustNewConstMetric(mode, prometheus.GaugeValue, backup, "backup", res.Result.Server, res.Result.MgmtIP, strconv.FormatFloat(res.Result.MgmtPort, 'f', -1, 64), res.Result.Sn, res.Result.PenFortMan))
	m = append(m, prometheus.MustNewConstMetric(status, prometheus.GaugeValue, down, "down", res.Result.Server, res.Result.MgmtIP, strconv.FormatFloat(res.Result.MgmtPort, 'f', -1, 64), res.Result.Sn, res.Result.PenFortMan))
	m = append(m, prometheus.MustNewConstMetric(status, prometheus.GaugeValue, up, "up", res.Result.Server, res.Result.MgmtIP, strconv.FormatFloat(res.Result.MgmtPort, 'f', -1, 64), res.Result.Sn, res.Result.PenFortMan))
	m = append(m, prometheus.MustNewConstMetric(status, prometheus.GaugeValue, handshake, "handshake", res.Result.Server, res.Result.MgmtIP, strconv.FormatFloat(res.Result.MgmtPort, 'f', -1, 64), res.Result.Sn, res.Result.PenFortMan))
	m = append(m, prometheus.MustNewConstMetric(registrationStatus, prometheus.GaugeValue, inProgress, "inprogress", res.Result.Server, res.Result.MgmtIP, strconv.FormatFloat(res.Result.MgmtPort, 'f', -1, 64), res.Result.Sn, res.Result.PenFortMan))
	m = append(m, prometheus.MustNewConstMetric(registrationStatus, prometheus.GaugeValue, registered, "registered", res.Result.Server, res.Result.MgmtIP, strconv.FormatFloat(res.Result.MgmtPort, 'f', -1, 64), res.Result.Sn, res.Result.PenFortMan))
	m = append(m, prometheus.MustNewConstMetric(registrationStatus, prometheus.GaugeValue, unregistered, "unregistered", res.Result.Server, res.Result.MgmtIP, strconv.FormatFloat(res.Result.MgmtPort, 'f', -1, 64), res.Result.Sn, res.Result.PenFortMan))
	m = append(m, prometheus.MustNewConstMetric(registrationStatus, prometheus.GaugeValue, defaultValue, "unknown", res.Result.Server, res.Result.MgmtIP, strconv.FormatFloat(res.Result.MgmtPort, 'f', -1, 64), res.Result.Sn, res.Result.PenFortMan))

	return m, true
}
