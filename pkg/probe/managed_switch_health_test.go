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
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestProbeManagedSwitchHealth(t *testing.T) {
	c := newFakeClient()
	c.prepare("api/v2/monitor/switch-controller/managed-switch/health", "testdata/managed-switch-health.jsonnet")
	meta := &TargetMetadata{
		VersionMajor: 7,
		VersionMinor: 4,
	}
	r := prometheus.NewPedanticRegistry()
	if !testProbeWithMetadata(probeManagedSwitchHealth, c, meta, r) {
		t.Errorf("probeManagedSwitchHealth() returned non-success")
	}

	em := `
		# HELP fortigate_managed_switch_cpu_mode_ratio CPU time split by mode
		# TYPE fortigate_managed_switch_cpu_mode_ratio gauge
		fortigate_managed_switch_cpu_mode_ratio{mode="idle",switch_name="FS00000000000024",vdom="root"} 1
		fortigate_managed_switch_cpu_mode_ratio{mode="nice",switch_name="FS00000000000024",vdom="root"} 0
		fortigate_managed_switch_cpu_mode_ratio{mode="system",switch_name="FS00000000000024",vdom="root"} 0
		fortigate_managed_switch_cpu_mode_ratio{mode="user",switch_name="FS00000000000024",vdom="root"} 0
		fortigate_managed_switch_cpu_mode_ratio{mode="idle",switch_name="FS00000000000027",vdom="root"} 1
		fortigate_managed_switch_cpu_mode_ratio{mode="nice",switch_name="FS00000000000027",vdom="root"} 0
		fortigate_managed_switch_cpu_mode_ratio{mode="system",switch_name="FS00000000000027",vdom="root"} 0
		fortigate_managed_switch_cpu_mode_ratio{mode="user",switch_name="FS00000000000027",vdom="root"} 0
		# HELP fortigate_managed_switch_cpu_usage_ratio Current CPU usage of the switch
		# TYPE fortigate_managed_switch_cpu_usage_ratio gauge
		fortigate_managed_switch_cpu_usage_ratio{switch_name="FS00000000000024",vdom="root"} 0
		fortigate_managed_switch_cpu_usage_ratio{switch_name="FS00000000000027",vdom="root"} 0
		# HELP fortigate_managed_switch_health_good FortiOS health rating per check, good=1 otherwise 0
		# TYPE fortigate_managed_switch_health_good gauge
		fortigate_managed_switch_health_good{check="cpu",switch_name="FS00000000000024",vdom="root"} 1
		fortigate_managed_switch_health_good{check="memory",switch_name="FS00000000000024",vdom="root"} 1
		fortigate_managed_switch_health_good{check="overall",switch_name="FS00000000000024",vdom="root"} 1
		fortigate_managed_switch_health_good{check="temperature",switch_name="FS00000000000024",vdom="root"} 1
		fortigate_managed_switch_health_good{check="uptime",switch_name="FS00000000000024",vdom="root"} 1
		fortigate_managed_switch_health_good{check="cpu",switch_name="FS00000000000027",vdom="root"} 1
		fortigate_managed_switch_health_good{check="memory",switch_name="FS00000000000027",vdom="root"} 1
		fortigate_managed_switch_health_good{check="overall",switch_name="FS00000000000027",vdom="root"} 1
		fortigate_managed_switch_health_good{check="temperature",switch_name="FS00000000000027",vdom="root"} 1
		fortigate_managed_switch_health_good{check="uptime",switch_name="FS00000000000027",vdom="root"} 1
		# HELP fortigate_managed_switch_memory_usage_ratio Current memory usage of the switch
		# TYPE fortigate_managed_switch_memory_usage_ratio gauge
		fortigate_managed_switch_memory_usage_ratio{switch_name="FS00000000000024",vdom="root"} 0.1
		fortigate_managed_switch_memory_usage_ratio{switch_name="FS00000000000027",vdom="root"} 0.15
		# HELP fortigate_managed_switch_temperature_celsius Temperature reported by a switch sensor
		# TYPE fortigate_managed_switch_temperature_celsius gauge
		fortigate_managed_switch_temperature_celsius{module="sensor1(CPU  Board Temp)",switch_name="FS00000000000024",vdom="root"} 41.937
		fortigate_managed_switch_temperature_celsius{module="sensor2(MAIN Board Temp1)",switch_name="FS00000000000024",vdom="root"} 63.875
		fortigate_managed_switch_temperature_celsius{module="sensor3(MAIN Board Temp2)",switch_name="FS00000000000024",vdom="root"} 51.312
		fortigate_managed_switch_temperature_celsius{module="sensor4(MAIN Board Temp3)",switch_name="FS00000000000024",vdom="root"} 38.687
		fortigate_managed_switch_temperature_celsius{module="sensor1(CPU  Board Temp)",switch_name="FS00000000000027",vdom="root"} 39
		fortigate_managed_switch_temperature_celsius{module="sensor2(MAIN Board Temp1)",switch_name="FS00000000000027",vdom="root"} 60.625
		fortigate_managed_switch_temperature_celsius{module="sensor3(MAIN Board Temp2)",switch_name="FS00000000000027",vdom="root"} 48.937
		fortigate_managed_switch_temperature_celsius{module="sensor4(MAIN Board Temp3)",switch_name="FS00000000000027",vdom="root"} 36.062
		# HELP fortigate_managed_switch_uptime_seconds Time since the switch was last started
		# TYPE fortigate_managed_switch_uptime_seconds gauge
		fortigate_managed_switch_uptime_seconds{switch_name="FS00000000000024",vdom="root"} 3.928968e+07
		fortigate_managed_switch_uptime_seconds{switch_name="FS00000000000027",vdom="root"} 3.928974e+07
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}

func TestProbeManagedSwitchHealthFortiOS76(t *testing.T) {
	c := newFakeClient()
	c.prepare("api/v2/monitor/switch-controller/managed-switch/health-status", "testdata/managed-switch-health-76.jsonnet")
	meta := &TargetMetadata{
		VersionMajor: 7,
		VersionMinor: 6,
	}
	r := prometheus.NewPedanticRegistry()
	if !testProbeWithMetadata(probeManagedSwitchHealth, c, meta, r) {
		t.Errorf("probeManagedSwitchHealth() returned non-success")
	}

	em := `
		# HELP fortigate_managed_switch_cpu_mode_ratio CPU time split by mode
		# TYPE fortigate_managed_switch_cpu_mode_ratio gauge
		fortigate_managed_switch_cpu_mode_ratio{mode="idle",switch_name="FOO-SW-76",vdom="root"} 0.15
		fortigate_managed_switch_cpu_mode_ratio{mode="nice",switch_name="FOO-SW-76",vdom="root"} 0
		fortigate_managed_switch_cpu_mode_ratio{mode="system",switch_name="FOO-SW-76",vdom="root"} 0.15
		fortigate_managed_switch_cpu_mode_ratio{mode="user",switch_name="FOO-SW-76",vdom="root"} 0.7
		# HELP fortigate_managed_switch_cpu_usage_ratio Current CPU usage of the switch
		# TYPE fortigate_managed_switch_cpu_usage_ratio gauge
		fortigate_managed_switch_cpu_usage_ratio{switch_name="FOO-SW-76",vdom="root"} 0.85
		# HELP fortigate_managed_switch_health_good FortiOS health rating per check, good=1 otherwise 0
		# TYPE fortigate_managed_switch_health_good gauge
		fortigate_managed_switch_health_good{check="cpu",switch_name="FOO-SW-76",vdom="root"} 0
		fortigate_managed_switch_health_good{check="memory",switch_name="FOO-SW-76",vdom="root"} 1
		fortigate_managed_switch_health_good{check="overall",switch_name="FOO-SW-76",vdom="root"} 0
		fortigate_managed_switch_health_good{check="temperature",switch_name="FOO-SW-76",vdom="root"} 1
		fortigate_managed_switch_health_good{check="uptime",switch_name="FOO-SW-76",vdom="root"} 1
		# HELP fortigate_managed_switch_memory_usage_ratio Current memory usage of the switch
		# TYPE fortigate_managed_switch_memory_usage_ratio gauge
		fortigate_managed_switch_memory_usage_ratio{switch_name="FOO-SW-76",vdom="root"} 0.4
		# HELP fortigate_managed_switch_temperature_celsius Temperature reported by a switch sensor
		# TYPE fortigate_managed_switch_temperature_celsius gauge
		fortigate_managed_switch_temperature_celsius{module="sensor1(CPU Board Temp)",switch_name="FOO-SW-76",vdom="root"} 55.5
		# HELP fortigate_managed_switch_uptime_seconds Time since the switch was last started
		# TYPE fortigate_managed_switch_uptime_seconds gauge
		fortigate_managed_switch_uptime_seconds{switch_name="FOO-SW-76",vdom="root"} 3600
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}
