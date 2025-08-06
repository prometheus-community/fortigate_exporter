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
			[]string{"label"}, nil,
		)
		cpuCoresSystem = prometheus.NewDesc(
			"fortigate_system_performance_status_cpu_cores_system",
			"Percentage of CPU utilization that occurred while executing at the system level.",
			[]string{"label"}, nil,
		)
		cpuCoresNice = prometheus.NewDesc(
			"fortigate_system_performance_status_cpu_cores_nice",
			"Percentage of CPU utilization that occurred while executing at the user level with nice priority.",
			[]string{"label"}, nil,
		)
		cpuCoresIdle = prometheus.NewDesc(
			"fortigate_system_performance_status_cpu_cores_idle",
			"Percentage of time that the CPU was idle and the system did not have an outstanding disk I/O request.",
			[]string{"label"}, nil,
		)
		cpuCoresIowait = prometheus.NewDesc(
			"fortigate_system_performance_status_cpu_cores_iowait",
			"Percentage of time that the CPU was idle during which the system had an outstanding disk I/O request.",
			[]string{"label"}, nil,
		)
		cpuUser = prometheus.NewDesc(
			"fortigate_system_performance_status_cpu_user",
			"Percentage of CPU utilization that occurred at the user level.",
			[]string{"label"}, nil,
		)
		cpuSystem = prometheus.NewDesc(
			"fortigate_system_performance_status_cpu_system",
			"Percentage of CPU utilization that occurred while executing at the system level.",
			[]string{"label"}, nil,
		)
		cpuNice = prometheus.NewDesc(
			"fortigate_system_performance_status_cpu_nice",
			"Percentage of CPU utilization that occurred while executing at the user level with nice priority.",
			[]string{"label"}, nil,
		)
		cpuIdle = prometheus.NewDesc(
			"fortigate_system_performance_status_cpu_idle",
			"Percentage of time that the CPU or CPUs were idle and the system did not have an outstanding disk I/O request.",
			[]string{"label"}, nil,
		)
		cpuIowait = prometheus.NewDesc(
			"fortigate_system_performance_status_cpu_iowait",
			"Percentage of time that the CPU or CPUs were idle during which the system had an outstanding disk I/O request.",
			[]string{"label"}, nil,
		)
		memTotal = prometheus.NewDesc(
			"fortigate_system_performance_status_mem_total",
			"All the installed memory in RAM, in bytes.",
			[]string{"label"}, nil,
		)
		memUsed = prometheus.NewDesc(
			"fortigate_system_performance_status_mem_used",
			"Memory are being used, in bytes.",
			[]string{"label"}, nil,
		)
		memFree = prometheus.NewDesc(
			"fortigate_system_performance_status_mem_free",
			"All the memory in RAM that is not being used for anything (even caches), in bytes.",
			[]string{"label"}, nil,
		)
		memFreeable = prometheus.NewDesc(
			"fortigate_system_performance_status_mem_freeable",
			"Freeable buffers/caches memory, in bytes.",
			[]string{"label"}, nil,
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
		Results []SystemPerformanceStatus `json:"results"`
	}

	var res SystemPerformanceStatusResult
	if err := c.Get("api/v2/monitor/system/performance/status", "", &res); err != nil {
		log.Printf("Error: %v", err)
		return nil, false
	}
	m := []prometheus.Metric{}
	var cpu_num, mem_num, core_num string
	for n, r := range res.Results {
		cpu_num = "cpu_" + strconv.Itoa(n)
		mem_num = "mem_" + strconv.Itoa(n)
		for i, core := range r.Cpu.Cores {
			core_num = "core_" + strconv.Itoa(i)
			m = append(m, prometheus.MustNewConstMetric(cpuCoresUser, prometheus.GaugeValue, float64(core.User), cpu_num + "_" + core_num))
			m = append(m, prometheus.MustNewConstMetric(cpuCoresSystem, prometheus.GaugeValue, float64(core.System), cpu_num + "_" + core_num))
			m = append(m, prometheus.MustNewConstMetric(cpuCoresNice, prometheus.GaugeValue, float64(core.Nice), cpu_num + "_" + core_num))
			m = append(m, prometheus.MustNewConstMetric(cpuCoresIdle, prometheus.GaugeValue, float64(core.Idle), cpu_num + "_" + core_num))
			m = append(m, prometheus.MustNewConstMetric(cpuCoresIowait, prometheus.GaugeValue, float64(core.Iowait), cpu_num + "_" + core_num))
		}
		m = append(m, prometheus.MustNewConstMetric(cpuUser,prometheus.GaugeValue, float64(r.Cpu.User), cpu_num))
		m = append(m, prometheus.MustNewConstMetric(cpuSystem,prometheus.GaugeValue, float64(r.Cpu.System), cpu_num))
		m = append(m, prometheus.MustNewConstMetric(cpuNice,prometheus.GaugeValue, float64(r.Cpu.Nice), cpu_num))
		m = append(m, prometheus.MustNewConstMetric(cpuIdle,prometheus.GaugeValue, float64(r.Cpu.Idle), cpu_num))
		m = append(m, prometheus.MustNewConstMetric(cpuIowait,prometheus.GaugeValue, float64(r.Cpu.Iowait), cpu_num))
		m = append(m, prometheus.MustNewConstMetric(memTotal,prometheus.GaugeValue, float64(r.Mem.Total), mem_num))
		m = append(m, prometheus.MustNewConstMetric(memUsed,prometheus.GaugeValue, float64(r.Mem.Used), mem_num))
		m = append(m, prometheus.MustNewConstMetric(memFree,prometheus.GaugeValue, float64(r.Mem.Free), mem_num))
		m = append(m, prometheus.MustNewConstMetric(memFreeable,prometheus.GaugeValue, float64(r.Mem.Freeable), mem_num))
	}

	return m, true
}