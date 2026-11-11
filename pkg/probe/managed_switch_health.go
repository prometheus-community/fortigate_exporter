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

func probeManagedSwitchHealth(c fortigatehttpclient.FortiHTTP, meta *TargetMetadata) ([]prometheus.Metric, bool) {
	var (
		healthGood = prometheus.NewDesc(
			"fortigate_managed_switch_health_good",
			"FortiOS health rating per check, good=1 otherwise 0",
			[]string{"vdom", "switch_name", "check"}, nil,
		)
		cpuUsage = prometheus.NewDesc(
			"fortigate_managed_switch_cpu_usage_ratio",
			"Current CPU usage of the switch",
			[]string{"vdom", "switch_name"}, nil,
		)
		cpuMode = prometheus.NewDesc(
			"fortigate_managed_switch_cpu_mode_ratio",
			"CPU time split by mode",
			[]string{"vdom", "switch_name", "mode"}, nil,
		)
		memoryUsage = prometheus.NewDesc(
			"fortigate_managed_switch_memory_usage_ratio",
			"Current memory usage of the switch",
			[]string{"vdom", "switch_name"}, nil,
		)
		uptime = prometheus.NewDesc(
			"fortigate_managed_switch_uptime_seconds",
			"Time since the switch was last started",
			[]string{"vdom", "switch_name"}, nil,
		)
		temperature = prometheus.NewDesc(
			"fortigate_managed_switch_temperature_celsius",
			"Temperature reported by a switch sensor",
			[]string{"vdom", "switch_name", "module"}, nil,
		)
	)

	type valueUnit struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}
	type rated struct {
		Value  float64 `json:"value"`
		Rating string  `json:"rating"`
	}
	type health struct {
		PerformanceStatus struct {
			CPU struct {
				User   valueUnit `json:"user"`
				System valueUnit `json:"system"`
				Nice   valueUnit `json:"nice"`
				Idle   valueUnit `json:"idle"`
			} `json:"cpu"`
		} `json:"performance-status"`
		Temperature []struct {
			Module string    `json:"module"`
			Status valueUnit `json:"status"`
		} `json:"temperature"`
		Summary struct {
			CPU         rated  `json:"cpu"`
			Memory      rated  `json:"memory"`
			Uptime      rated  `json:"uptime"`
			Temperature rated  `json:"temperature"`
			Overall     string `json:"overall"`
		} `json:"summary"`
	}
	type healthResponse []struct {
		Results map[string]health `json:"results"`
		VDOM    string            `json:"vdom"`
	}

	// FortiOS 7.6 deprecates /health in favour of /health-status.
	apiPath := "api/v2/monitor/switch-controller/managed-switch/health"
	if meta.VersionMajor > 7 || (meta.VersionMajor == 7 && meta.VersionMinor >= 6) {
		apiPath = "api/v2/monitor/switch-controller/managed-switch/health-status"
	}

	var response healthResponse
	if err := c.Get(apiPath, "vdom=*", &response); err != nil {
		log.Printf("Error: %v", err)
		return nil, false
	}

	good := func(rating string) float64 {
		if rating == "good" {
			return 1
		}
		return 0
	}

	var m []prometheus.Metric
	for _, rs := range response {
		for switchName, h := range rs.Results {
			s := h.Summary
			for _, check := range []struct {
				name   string
				rating string
			}{
				{"cpu", s.CPU.Rating},
				{"memory", s.Memory.Rating},
				{"uptime", s.Uptime.Rating},
				{"temperature", s.Temperature.Rating},
				{"overall", s.Overall},
			} {
				m = append(m, prometheus.MustNewConstMetric(healthGood, prometheus.GaugeValue, good(check.rating), rs.VDOM, switchName, check.name))
			}
			m = append(m, prometheus.MustNewConstMetric(cpuUsage, prometheus.GaugeValue, s.CPU.Value/100, rs.VDOM, switchName))
			m = append(m, prometheus.MustNewConstMetric(memoryUsage, prometheus.GaugeValue, s.Memory.Value/100, rs.VDOM, switchName))
			m = append(m, prometheus.MustNewConstMetric(uptime, prometheus.GaugeValue, s.Uptime.Value, rs.VDOM, switchName))

			cpu := h.PerformanceStatus.CPU
			for mode, v := range map[string]float64{
				"user":   cpu.User.Value,
				"system": cpu.System.Value,
				"nice":   cpu.Nice.Value,
				"idle":   cpu.Idle.Value,
			} {
				m = append(m, prometheus.MustNewConstMetric(cpuMode, prometheus.GaugeValue, v/100, rs.VDOM, switchName, mode))
			}

			for _, t := range h.Temperature {
				if t.Status.Unit != "celsius" {
					continue
				}
				m = append(m, prometheus.MustNewConstMetric(temperature, prometheus.GaugeValue, t.Status.Value, rs.VDOM, switchName, t.Module))
			}
		}
	}

	return m, true
}
