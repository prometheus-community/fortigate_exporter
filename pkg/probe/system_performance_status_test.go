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
	# HELP fortigate_system_performance_status_cpu_cores_idle_ratio Percentage of time that the CPU was idle and the system did not have an outstanding disk I/O request.
	# TYPE fortigate_system_performance_status_cpu_cores_idle_ratio gauge
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_0",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_1",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_10",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_11",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_12",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_13",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_14",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_15",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_16",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_17",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_18",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_19",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_2",vdom="root"} 0.53
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_20",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_21",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_22",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_23",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_24",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_25",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_26",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_27",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_28",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_29",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_3",vdom="root"} 0.51
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_30",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_31",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_4",vdom="root"} 0.99
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_5",vdom="root"} 0.98
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_6",vdom="root"} 0.92
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_7",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_8",vdom="root"} 0.92
	fortigate_system_performance_status_cpu_cores_idle_ratio{label="core_9",vdom="root"} 1
	# HELP fortigate_system_performance_status_cpu_cores_iowait_ratio Percentage of time that the CPU was idle during which the system had an outstanding disk I/O request.
	# TYPE fortigate_system_performance_status_cpu_cores_iowait_ratio gauge
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_0",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_1",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_10",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_11",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_12",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_13",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_14",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_15",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_16",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_17",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_18",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_19",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_2",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_20",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_21",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_22",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_23",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_24",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_25",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_26",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_27",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_28",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_29",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_3",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_30",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_31",vdom="root"} 1
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_4",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_5",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_6",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_7",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_8",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_iowait_ratio{label="core_9",vdom="root"} 0
	# HELP fortigate_system_performance_status_cpu_cores_nice_ratio Percentage of CPU utilization that occurred while executing at the user level with nice priority.
	# TYPE fortigate_system_performance_status_cpu_cores_nice_ratio gauge
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_0",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_1",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_10",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_11",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_12",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_13",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_14",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_15",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_16",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_17",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_18",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_19",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_2",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_20",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_21",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_22",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_23",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_24",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_25",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_26",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_27",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_28",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_29",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_3",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_30",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_31",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_4",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_5",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_6",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_7",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_8",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_nice_ratio{label="core_9",vdom="root"} 0
	# HELP fortigate_system_performance_status_cpu_cores_system_ratio Percentage of CPU utilization that occurred while executing at the system level.
	# TYPE fortigate_system_performance_status_cpu_cores_system_ratio gauge
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_0",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_1",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_10",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_11",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_12",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_13",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_14",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_15",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_16",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_17",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_18",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_19",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_2",vdom="root"} 0.2
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_20",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_21",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_22",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_23",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_24",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_25",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_26",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_27",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_28",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_29",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_3",vdom="root"} 0.18
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_30",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_31",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_4",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_5",vdom="root"} 0.01
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_6",vdom="root"} 0.01
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_7",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_8",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_system_ratio{label="core_9",vdom="root"} 0
	# HELP fortigate_system_performance_status_cpu_cores_user_ratio Percentage of CPU utilization that occurred at the user level.
	# TYPE fortigate_system_performance_status_cpu_cores_user_ratio gauge
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_0",vdom="root"} 0.
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_1",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_10",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_11",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_12",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_13",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_14",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_15",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_16",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_17",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_18",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_19",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_2",vdom="root"} 0.27
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_20",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_21",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_22",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_23",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_24",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_25",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_26",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_27",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_28",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_29",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_3",vdom="root"} 0.31
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_30",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_31",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_4",vdom="root"} 0.01
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_5",vdom="root"} 0.01
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_6",vdom="root"} 0.07
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_7",vdom="root"} 0
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_8",vdom="root"} 0.08
	fortigate_system_performance_status_cpu_cores_user_ratio{label="core_9",vdom="root"} 0
	# HELP fortigate_system_performance_status_cpu_idle_ratio Percentage of time that the CPU or CPUs were idle and the system did not have an outstanding disk I/O request.
	# TYPE fortigate_system_performance_status_cpu_idle_ratio gauge
	fortigate_system_performance_status_cpu_idle_ratio{vdom="root"} 0.97
	# HELP fortigate_system_performance_status_cpu_iowait_ratio Percentage of time that the CPU or CPUs were idle during which the system had an outstanding disk I/O request.
	# TYPE fortigate_system_performance_status_cpu_iowait_ratio gauge
	fortigate_system_performance_status_cpu_iowait_ratio{vdom="root"} 0
	# HELP fortigate_system_performance_status_cpu_nice_ratio Percentage of CPU utilization that occurred while executing at the user level with nice priority.
	# TYPE fortigate_system_performance_status_cpu_nice_ratio gauge
	fortigate_system_performance_status_cpu_nice_ratio{vdom="root"} 0
	# HELP fortigate_system_performance_status_cpu_system_ratio Percentage of CPU utilization that occurred while executing at the system level.
	# TYPE fortigate_system_performance_status_cpu_system_ratio gauge
	fortigate_system_performance_status_cpu_system_ratio{vdom="root"} 0.01
	# HELP fortigate_system_performance_status_cpu_user_ratio Percentage of CPU utilization that occurred at the user level.
	# TYPE fortigate_system_performance_status_cpu_user_ratio gauge
	fortigate_system_performance_status_cpu_user_ratio{vdom="root"} 0.02
	# HELP fortigate_system_performance_status_mem_free_bytes All the memory in RAM that is not being used for anything (even caches), in bytes.
	# TYPE fortigate_system_performance_status_mem_free_bytes gauge
	fortigate_system_performance_status_mem_free_bytes{vdom="root"} 361498
	# HELP fortigate_system_performance_status_mem_freeable_bytes Freeable buffers/caches memory, in bytes.
	# TYPE fortigate_system_performance_status_mem_freeable_bytes gauge
	fortigate_system_performance_status_mem_freeable_bytes{vdom="root"} 5299
	# HELP fortigate_system_performance_status_mem_bytes_total All the installed memory in RAM, in bytes.
	# TYPE fortigate_system_performance_status_mem_bytes_total gauge
	fortigate_system_performance_status_mem_bytes_total{vdom="root"} 50727878
	# HELP fortigate_system_performance_status_mem_used_bytes Memory are being used, in bytes.
	# TYPE fortigate_system_performance_status_mem_used_bytes gauge
	fortigate_system_performance_status_mem_used_bytes{vdom="root"} 140480780
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}
