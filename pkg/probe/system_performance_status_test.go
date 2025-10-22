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
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestSystemPerformanceStatus(t *testing.T) {
	c := newFakeClient()
	c.prepare("api/v2/monitor/system/performance/status", "testdata/system-performance-status.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !testProbe(probeSystemPerformanceStatus, c, r) {
		t.Errorf("probeSystemPerformanceStatus() returned non-success")
	}

	em := `
	# HELP fortigate_system_performance_status_cpu_cores_idle Percentage of time that the CPU was idle and the system did not have an outstanding disk I/O request.
	# TYPE fortigate_system_performance_status_cpu_cores_idle gauge
	fortigate_system_performance_status_cpu_cores_idle{label="core_0",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_idle{label="core_1",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_idle{label="core_2",vdom="root"} 0
	# HELP fortigate_system_performance_status_cpu_cores_iowait Percentage of time that the CPU was idle during which the system had an outstanding disk I/O request.
	# TYPE fortigate_system_performance_status_cpu_cores_iowait gauge
	fortigate_system_performance_status_cpu_cores_iowait{label="core_0",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait{label="core_1",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait{label="core_2",vdom="root"} 0
	# HELP fortigate_system_performance_status_cpu_cores_nice Percentage of CPU utilization that occurred while executing at the user level with nice priority.
	# TYPE fortigate_system_performance_status_cpu_cores_nice gauge
	fortigate_system_performance_status_cpu_cores_nice{label="core_0",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice{label="core_1",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice{label="core_2",vdom="root"} 0
	# HELP fortigate_system_performance_status_cpu_cores_system Percentage of CPU utilization that occurred while executing at the system level.
	# TYPE fortigate_system_performance_status_cpu_cores_system gauge
	fortigate_system_performance_status_cpu_cores_system{label="core_0",vdom="root"} 13
	fortigate_system_performance_status_cpu_cores_system{label="core_1",vdom="root"} 14
	fortigate_system_performance_status_cpu_cores_system{label="core_2",vdom="root"} 0
	# HELP fortigate_system_performance_status_cpu_cores_user Percentage of CPU utilization that occurred at the user level.
	# TYPE fortigate_system_performance_status_cpu_cores_user gauge
	fortigate_system_performance_status_cpu_cores_user{label="core_0",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user{label="core_1",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_user{label="core_2",vdom="root"} 2
	# HELP fortigate_system_performance_status_cpu_idle Percentage of time that the CPU or CPUs were idle and the system did not have an outstanding disk I/O request.
	# TYPE fortigate_system_performance_status_cpu_idle gauge
	fortigate_system_performance_status_cpu_idle{vdom="root"} 0
	# HELP fortigate_system_performance_status_cpu_iowait Percentage of time that the CPU or CPUs were idle during which the system had an outstanding disk I/O request.
	# TYPE fortigate_system_performance_status_cpu_iowait gauge
	fortigate_system_performance_status_cpu_iowait{vdom="root"} 0
	# HELP fortigate_system_performance_status_cpu_nice Percentage of CPU utilization that occurred while executing at the user level with nice priority.
	# TYPE fortigate_system_performance_status_cpu_nice gauge
	fortigate_system_performance_status_cpu_nice{vdom="root"} 0
	# HELP fortigate_system_performance_status_cpu_system Percentage of CPU utilization that occurred while executing at the system level.
	# TYPE fortigate_system_performance_status_cpu_system gauge
	fortigate_system_performance_status_cpu_system{vdom="root"} 0
	# HELP fortigate_system_performance_status_cpu_user Percentage of CPU utilization that occurred at the user level.
	# TYPE fortigate_system_performance_status_cpu_user gauge
	fortigate_system_performance_status_cpu_user{vdom="root"} 200
	# HELP fortigate_system_performance_status_mem_free All the memory in RAM that is not being used for anything (even caches), in bytes.
	# TYPE fortigate_system_performance_status_mem_free gauge
	fortigate_system_performance_status_mem_free{vdom="root"} 0
	# HELP fortigate_system_performance_status_mem_freeable Freeable buffers/caches memory, in bytes.
	# TYPE fortigate_system_performance_status_mem_freeable gauge
	fortigate_system_performance_status_mem_freeable{vdom="root"} 0
	# HELP fortigate_system_performance_status_mem_total All the installed memory in RAM, in bytes.
	# TYPE fortigate_system_performance_status_mem_total gauge
	fortigate_system_performance_status_mem_total{vdom="root"} 0
	# HELP fortigate_system_performance_status_mem_used Memory are being used, in bytes.
	# TYPE fortigate_system_performance_status_mem_used gauge
	fortigate_system_performance_status_mem_used{vdom="root"} 0
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}