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

func probeSystemPerformanceStatus(c http.FortiHTTP, meta *TargetMetadata) ([]prometheus.Metric, bool) {
	var (
		cpuCoresUser = prometheus.NewDesc(
			"fortigate_system_performance_status_cpu_cores_user",
			"Percentage of CPU utilization that occurred at the user level.",
			[]string{"label", "vdom"}, nil,
		)
		cpuCoresSystem = prometheus.NewDesc(
			"fortigate_system_performance_status_cpu_cores_system",
			"Percentage of CPU utilization that occurred while executing at the system level.",
			[]string{"label", "vdom"}, nil,
		)
		cpuCoresNice = prometheus.NewDesc(
			"fortigate_system_performance_status_cpu_cores_nice",
			"Percentage of CPU utilization that occurred while executing at the user level with nice priority.",
			[]string{"label", "vdom"}, nil,
		)
		cpuCoresIdle = prometheus.NewDesc(
			"fortigate_system_performance_status_cpu_cores_idle",
			"Percentage of time that the CPU was idle and the system did not have an outstanding disk I/O request.",
			[]string{"label", "vdom"}, nil,
		)
		cpuCoresIowait = prometheus.NewDesc(
			"fortigate_system_performance_status_cpu_cores_iowait",
			"Percentage of time that the CPU was idle during which the system had an outstanding disk I/O request.",
			[]string{"label", "vdom"}, nil,
		)
		cpuUser = prometheus.NewDesc(
			"fortigate_system_performance_status_cpu_user",
			"Percentage of CPU utilization that occurred at the user level.",
			[]string{"vdom"}, nil,
		)
		cpuSystem = prometheus.NewDesc(
			"fortigate_system_performance_status_cpu_system",
			"Percentage of CPU utilization that occurred while executing at the system level.",
			[]string{"vdom"}, nil,
		)
		cpuNice = prometheus.NewDesc(
			"fortigate_system_performance_status_cpu_nice",
			"Percentage of CPU utilization that occurred while executing at the user level with nice priority.",
			[]string{"vdom"}, nil,
		)
		cpuIdle = prometheus.NewDesc(
			"fortigate_system_performance_status_cpu_idle",
			"Percentage of time that the CPU or CPUs were idle and the system did not have an outstanding disk I/O request.",
			[]string{"vdom"}, nil,
		)
		cpuIowait = prometheus.NewDesc(
			"fortigate_system_performance_status_cpu_iowait",
			"Percentage of time that the CPU or CPUs were idle during which the system had an outstanding disk I/O request.",
			[]string{"vdom"}, nil,
		)
		memTotal = prometheus.NewDesc(
			"fortigate_system_performance_status_mem_total",
			"All the installed memory in RAM, in bytes.",
			[]string{"vdom"}, nil,
		)
		memUsed = prometheus.NewDesc(
			"fortigate_system_performance_status_mem_used",
			"Memory are being used, in bytes.",
			[]string{"vdom"}, nil,
		)
		memFree = prometheus.NewDesc(
			"fortigate_system_performance_status_mem_free",
			"All the memory in RAM that is not being used for anything (even caches), in bytes.",
			[]string{"vdom"}, nil,
		)
		memFreeable = prometheus.NewDesc(
			"fortigate_system_performance_status_mem_freeable",
			"Freeable buffers/caches memory, in bytes.",
			[]string{"vdom"}, nil,
		)
	)

	type SystemPerformanceStatusCores struct {
		User   int `json:"user"`
		System int `json:"system"`
		Nice   int `json:"nice"`
		Idle   int `json:"idle"`
		Iowait int `json:"iowait"`
	}

	type SystemPerformanceStatusCpu struct {
		Cores  []SystemPerformanceStatusCores `json:"cores"`
		User   int                            `json:"user"`
		System int                            `json:"system"`
		Nice   int                            `json:"nice"`
		Idle   int                            `json:"idle"`
		Iowait int                            `json:"iowait"`
	}

	type SystemPerformanceStatusMem struct {
		Total    int `json:"total"`
		Used     int `json:"used"`
		Free     int `json:"free"`
		Freeable int `json:"freeable"`
	}

	type SystemPerformanceStatus struct {
		Cpu SystemPerformanceStatusCpu `json:"cpu"`
		Mem SystemPerformanceStatusMem `json:"mem"`
	}

	type SystemPerformanceStatusResult struct {
		Results SystemPerformanceStatus `json:"results"`
		VDOM    string                  `json:"vdom"`
	}

	var res SystemPerformanceStatusResult
	if err := c.Get("api/v2/monitor/system/performance/status", "", &res); err != nil {
		log.Printf("Error: %v", err)
		return nil, false
	}
	m := []prometheus.Metric{}
	var core_num string
	for i, core := range res.Results.Cpu.Cores {
		core_num = "core_" + strconv.Itoa(i)
		m = append(m, prometheus.MustNewConstMetric(cpuCoresUser, prometheus.GaugeValue, float64(core.User), core_num, res.VDOM))
		m = append(m, prometheus.MustNewConstMetric(cpuCoresSystem, prometheus.GaugeValue, float64(core.System), core_num, res.VDOM))
		m = append(m, prometheus.MustNewConstMetric(cpuCoresNice, prometheus.GaugeValue, float64(core.Nice), core_num, res.VDOM))
		m = append(m, prometheus.MustNewConstMetric(cpuCoresIdle, prometheus.GaugeValue, float64(core.Idle), core_num, res.VDOM))
		m = append(m, prometheus.MustNewConstMetric(cpuCoresIowait, prometheus.GaugeValue, float64(core.Iowait), core_num, res.VDOM))
	}
	m = append(m, prometheus.MustNewConstMetric(cpuUser,prometheus.GaugeValue, float64(res.Results.Cpu.User), res.VDOM))
	m = append(m, prometheus.MustNewConstMetric(cpuSystem,prometheus.GaugeValue, float64(res.Results.Cpu.System), res.VDOM))
	m = append(m, prometheus.MustNewConstMetric(cpuNice,prometheus.GaugeValue, float64(res.Results.Cpu.Nice), res.VDOM))
	m = append(m, prometheus.MustNewConstMetric(cpuIdle,prometheus.GaugeValue, float64(res.Results.Cpu.Idle), res.VDOM))
	m = append(m, prometheus.MustNewConstMetric(cpuIowait,prometheus.GaugeValue, float64(res.Results.Cpu.Iowait), res.VDOM))
	m = append(m, prometheus.MustNewConstMetric(memTotal,prometheus.GaugeValue, float64(res.Results.Mem.Total), res.VDOM))
	m = append(m, prometheus.MustNewConstMetric(memUsed,prometheus.GaugeValue, float64(res.Results.Mem.Used), res.VDOM))
	m = append(m, prometheus.MustNewConstMetric(memFree,prometheus.GaugeValue, float64(res.Results.Mem.Free), res.VDOM))
	m = append(m, prometheus.MustNewConstMetric(memFreeable,prometheus.GaugeValue, float64(res.Results.Mem.Freeable), res.VDOM))
	return m, true
}