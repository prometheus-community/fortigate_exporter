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
	"fmt"
	"log"

	"github.com/prometheus-community/fortigate_exporter/pkg/http"
	"github.com/prometheus/client_golang/prometheus"
)

func probeSwitchHealth(c http.FortiHTTP, meta *TargetMetadata) ([]prometheus.Metric, bool) {
	var (
		mSumCPU = prometheus.NewDesc(
			"fortiswitch_health_summary_cpu",
			"Boolean indicator if CPU health is good",
			[]string{"value", "rating", "fortiswitch", "vdom"}, nil,
		)
		mSumMem = prometheus.NewDesc(
			"fortiswitch_health_summary_memory",
			"Boolean indicator if Memory health is good",
			[]string{"value", "rating", "fortiswitch", "vdom"}, nil,
		)
		mSumUpTime = prometheus.NewDesc(
			"fortiswitch_health_summary_uptime",
			"Boolean indicator if Uptime is good",
			[]string{"value", "rating", "fortiswitch", "vdom"}, nil,
		)
		mSumTemp = prometheus.NewDesc(
			"fortiswitch_health_summary_temperature",
			"Boolean indicator if Temperature health is good",
			[]string{"value", "rating", "fortiswitch", "vdom"}, nil,
		)
		mTemp = prometheus.NewDesc(
			"fortiswitch_health_temperature",
			"Temperature per switch sensor",
			[]string{"unit", "module", "fortiswitch", "vdom"}, nil,
		)
		mCpuUser = prometheus.NewDesc(
			"fortiswitch_health_performance_stats_cpu_user",
			"Fortiswitch CPU user usage",
			[]string{"unit", "fortiswitch", "vdom"}, nil,
		)
		mCpuSystem = prometheus.NewDesc(
			"fortiswitch_health_performance_stats_cpu_system",
			"Fortiswitch CPU system usage",
			[]string{"unit", "fortiswitch", "vdom"}, nil,
		)
		mCpuIdle = prometheus.NewDesc(
			"fortiswitch_health_performance_stats_cpu_idle",
			"Fortiswitch CPU idle",
			[]string{"unit", "fortiswitch", "vdom"}, nil,
		)
		mCpuNice = prometheus.NewDesc(
			"fortiswitch_health_performance_stats_cpu_nice",
			"Fortiswitch CPU nice usage",
			[]string{"unit", "fortiswitch", "vdom"}, nil,
		)
	)
	type Sum struct {
		Value  float64 `json:"value"`
		Rating string  `json:"rating"`
	}
	type Status struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}
	type Uptime struct {
		Days    Status `json:"days"`
		Hours   Status `json:"hours"`
		Minutes Status `json:"minutes"`
	}
	type Network struct {
		In1Min  Status `json:"in-1min"`
		In10Min Status `json:"in-10min"`
		In30Min Status `json:"in-30min"`
	}
	type Memory struct {
		Used Status `json:"used"`
	}
	type CPU struct {
		User   Status `json:"user"`
		System Status `json:"system"`
		Nice   Status `json:"nice"`
		Idle   Status `json:"idle"`
	}
	type PerformanceStatus struct {
		CPU     CPU     `json:"cpu"`
		Memory  Memory  `json:"memory"`
		Network Network `json:"network"`
		Uptime  Uptime  `json:"uptime"`
	}
	type Temperature struct {
		Module string
		Status Status
	}
	type Summary struct {
		Overall     string `json:"overall"`
		CPU         Sum
		Memory      Sum
		Uptime      Sum
		Temperature Sum
	}
	type Poe struct {
		Value    float64 `json:"value"`
		MaxValue float64 `json:"max_value"`
		Unit     string  `json:"unit"`
	}
	type Results struct {
		PerformanceStatus PerformanceStatus `json:"performance-status"`
		Temperature       []Temperature     `json:"temperature"`
		Summary           Summary           `json:"summary"`
		Poe               Poe               `json:"poe"`
	}

	type swResponse struct {
		Results map[string]Results `json:"results"`
		Vdom    string
	}

	var apiPath string

	if meta.VersionMajor > 7 || (meta.VersionMajor == 7 && meta.VersionMinor >= 6) {
		apiPath = "api/v2/monitor/switch-controller/managed-switch/health-status"
	} else {
		apiPath = "api/v2/monitor/switch-controller/managed-switch/health"
	}

	var r swResponse

	if err := c.Get(apiPath, "vdom=root", &r); err != nil {
		log.Printf("Error: %v", err)
		return nil, false
	}

	m := []prometheus.Metric{}

	for fswitch, hr := range r.Results {

		var cpuGood float64
		if hr.Summary.CPU.Rating == "good" {
			cpuGood = 1
		} else {
			cpuGood = 0
		}
		m = append(m, prometheus.MustNewConstMetric(mSumCPU, prometheus.GaugeValue, cpuGood, fmt.Sprintf("%.0f", hr.Summary.CPU.Value), hr.Summary.CPU.Rating, fswitch, r.Vdom))

		var memGood float64
		if hr.Summary.Memory.Rating == "good" {
			memGood = 1
		} else {
			memGood = 0
		}
		m = append(m, prometheus.MustNewConstMetric(mSumMem, prometheus.GaugeValue, memGood, fmt.Sprintf("%0.f", hr.Summary.Memory.Value), hr.Summary.Memory.Rating, fswitch, r.Vdom))

		var uptimeGood float64
		if hr.Summary.Uptime.Rating == "good" {
			uptimeGood = 1
		} else {
			uptimeGood = 0
		}
		m = append(m, prometheus.MustNewConstMetric(mSumUpTime, prometheus.GaugeValue, uptimeGood, fmt.Sprintf("%0.f", hr.Summary.Uptime.Value), hr.Summary.Uptime.Rating, fswitch, r.Vdom))

		var tempGood float64
		if hr.Summary.Temperature.Rating == "good" {
			tempGood = 1
		} else {
			tempGood = 0
		}
		m = append(m, prometheus.MustNewConstMetric(mSumTemp, prometheus.GaugeValue, tempGood, fmt.Sprintf("%0.f", hr.Summary.Temperature.Value), hr.Summary.Temperature.Rating, fswitch, r.Vdom))

		for _, ts := range hr.Temperature {
			m = append(m, prometheus.MustNewConstMetric(mTemp, prometheus.GaugeValue, ts.Status.Value, ts.Status.Unit, ts.Module, fswitch, r.Vdom))
		}

		CpuUnit := hr.PerformanceStatus.CPU.System.Unit

		m = append(m, prometheus.MustNewConstMetric(mCpuUser, prometheus.GaugeValue, hr.PerformanceStatus.CPU.User.Value, CpuUnit, fswitch, r.Vdom))
		m = append(m, prometheus.MustNewConstMetric(mCpuNice, prometheus.GaugeValue, hr.PerformanceStatus.CPU.Nice.Value, CpuUnit, fswitch, r.Vdom))
		m = append(m, prometheus.MustNewConstMetric(mCpuSystem, prometheus.GaugeValue, hr.PerformanceStatus.CPU.System.Value, CpuUnit, fswitch, r.Vdom))
		m = append(m, prometheus.MustNewConstMetric(mCpuIdle, prometheus.GaugeValue, hr.PerformanceStatus.CPU.Idle.Value, CpuUnit, fswitch, r.Vdom))
	}

	return m, true
}
