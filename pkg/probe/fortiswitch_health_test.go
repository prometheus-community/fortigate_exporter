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

func TestSwitchHealth(t *testing.T) {
	c := newFakeClient()
	c.prepare("api/v2/monitor/switch-controller/managed-switch/health", "testdata/fsw-health.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !testProbe(probeSwitchHealth, c, r) {
		t.Errorf("probeSwitchHealth() returned non-success")
	}

	em := `
	# HELP fortiswitch_health_performance_stats_cpu_idle Fortiswitch CPU idle
	# TYPE fortiswitch_health_performance_stats_cpu_idle gauge
   	fortiswitch_health_performance_stats_cpu_idle{fortiswitch="FS00000000000024",unit="%%",vdom="root"} 100
   	fortiswitch_health_performance_stats_cpu_idle{fortiswitch="FS00000000000027",unit="%%",vdom="root"} 100
   	fortiswitch_health_performance_stats_cpu_idle{fortiswitch="FS00000000000030",unit="%%",vdom="root"} 100
   	fortiswitch_health_performance_stats_cpu_idle{fortiswitch="FS00000000000038",unit="%%",vdom="root"} 100
   	# HELP fortiswitch_health_performance_stats_cpu_nice Fortiswitch CPU nice usage
   	# TYPE fortiswitch_health_performance_stats_cpu_nice gauge
	fortiswitch_health_performance_stats_cpu_nice{fortiswitch="FS00000000000024",unit="%%",vdom="root"} 0
   	fortiswitch_health_performance_stats_cpu_nice{fortiswitch="FS00000000000027",unit="%%",vdom="root"} 0
   	fortiswitch_health_performance_stats_cpu_nice{fortiswitch="FS00000000000030",unit="%%",vdom="root"} 0
   	fortiswitch_health_performance_stats_cpu_nice{fortiswitch="FS00000000000038",unit="%%",vdom="root"} 0
   	# HELP fortiswitch_health_performance_stats_cpu_system Fortiswitch CPU system usage
   	# TYPE fortiswitch_health_performance_stats_cpu_system gauge
   	fortiswitch_health_performance_stats_cpu_system{fortiswitch="FS00000000000024",unit="%%",vdom="root"} 0
   	fortiswitch_health_performance_stats_cpu_system{fortiswitch="FS00000000000027",unit="%%",vdom="root"} 0
   	fortiswitch_health_performance_stats_cpu_system{fortiswitch="FS00000000000030",unit="%%",vdom="root"} 0
   	fortiswitch_health_performance_stats_cpu_system{fortiswitch="FS00000000000038",unit="%%",vdom="root"} 0
   	# HELP fortiswitch_health_performance_stats_cpu_user Fortiswitch CPU user usage
   	# TYPE fortiswitch_health_performance_stats_cpu_user gauge
   	fortiswitch_health_performance_stats_cpu_user{fortiswitch="FS00000000000024",unit="%%",vdom="root"} 0
   	fortiswitch_health_performance_stats_cpu_user{fortiswitch="FS00000000000027",unit="%%",vdom="root"} 0
   	fortiswitch_health_performance_stats_cpu_user{fortiswitch="FS00000000000030",unit="%%",vdom="root"} 0
   	fortiswitch_health_performance_stats_cpu_user{fortiswitch="FS00000000000038",unit="%%",vdom="root"} 0
   	# HELP fortiswitch_health_summary_cpu Boolean indicator if CPU health is good
   	# TYPE fortiswitch_health_summary_cpu gauge
   	fortiswitch_health_summary_cpu{fortiswitch="FS00000000000024",rating="good",value="0",vdom="root"} 1
   	fortiswitch_health_summary_cpu{fortiswitch="FS00000000000027",rating="good",value="0",vdom="root"} 1
   	fortiswitch_health_summary_cpu{fortiswitch="FS00000000000030",rating="good",value="0",vdom="root"} 1
   	fortiswitch_health_summary_cpu{fortiswitch="FS00000000000038",rating="good",value="0",vdom="root"} 1
   	# HELP fortiswitch_health_summary_memory Boolean indicator if Memory health is good
   	# TYPE fortiswitch_health_summary_memory gauge
   	fortiswitch_health_summary_memory{fortiswitch="FS00000000000024",rating="good",value="10",vdom="root"} 1
   	fortiswitch_health_summary_memory{fortiswitch="FS00000000000027",rating="good",value="15",vdom="root"} 1
   	fortiswitch_health_summary_memory{fortiswitch="FS00000000000030",rating="good",value="50",vdom="root"} 1
   	fortiswitch_health_summary_memory{fortiswitch="FS00000000000038",rating="good",value="32",vdom="root"} 1
   	# HELP fortiswitch_health_summary_temperature Boolean indicator if Temperature health is good
   	# TYPE fortiswitch_health_summary_temperature gauge
   	fortiswitch_health_summary_temperature{fortiswitch="FS00000000000024",rating="good",value="49",vdom="root"} 1
   	fortiswitch_health_summary_temperature{fortiswitch="FS00000000000027",rating="good",value="46",vdom="root"} 1
   	fortiswitch_health_summary_temperature{fortiswitch="FS00000000000030",rating="good",value="40",vdom="root"} 1
   	fortiswitch_health_summary_temperature{fortiswitch="FS00000000000038",rating="good",value="42",vdom="root"} 1
   	# HELP fortiswitch_health_summary_uptime Boolean indicator if Uptime is good
   	# TYPE fortiswitch_health_summary_uptime gauge
   	fortiswitch_health_summary_uptime{fortiswitch="FS00000000000024",rating="good",value="39289680",vdom="root"} 1
   	fortiswitch_health_summary_uptime{fortiswitch="FS00000000000027",rating="good",value="39289740",vdom="root"} 1
   	fortiswitch_health_summary_uptime{fortiswitch="FS00000000000030",rating="good",value="26612880",vdom="root"} 1
   	fortiswitch_health_summary_uptime{fortiswitch="FS00000000000038",rating="good",value="26612580",vdom="root"} 1
   	# HELP fortiswitch_health_temperature Temperature per switch sensor
   	# TYPE fortiswitch_health_temperature gauge
   	fortiswitch_health_temperature{fortiswitch="FS00000000000024",module="sensor1(CPU  Board Temp)",unit="celsius",vdom="root"} 41.937
   	fortiswitch_health_temperature{fortiswitch="FS00000000000024",module="sensor2(MAIN Board Temp1)",unit="celsius",vdom="root"} 63.875
   	fortiswitch_health_temperature{fortiswitch="FS00000000000024",module="sensor3(MAIN Board Temp2)",unit="celsius",vdom="root"} 51.312
   	fortiswitch_health_temperature{fortiswitch="FS00000000000024",module="sensor4(MAIN Board Temp3)",unit="celsius",vdom="root"} 38.687
   	fortiswitch_health_temperature{fortiswitch="FS00000000000027",module="sensor1(CPU  Board Temp)",unit="celsius",vdom="root"} 39
   	fortiswitch_health_temperature{fortiswitch="FS00000000000027",module="sensor2(MAIN Board Temp1)",unit="celsius",vdom="root"} 60.625
   	fortiswitch_health_temperature{fortiswitch="FS00000000000027",module="sensor3(MAIN Board Temp2)",unit="celsius",vdom="root"} 48.937
   	fortiswitch_health_temperature{fortiswitch="FS00000000000027",module="sensor4(MAIN Board Temp3)",unit="celsius",vdom="root"} 36.062
   	fortiswitch_health_temperature{fortiswitch="FS00000000000030",module="sensor1(CPU  Board Temp)",unit="celsius",vdom="root"} 33.875
   	fortiswitch_health_temperature{fortiswitch="FS00000000000030",module="sensor2(MAIN Board Temp1)",unit="celsius",vdom="root"} 53.75
   	fortiswitch_health_temperature{fortiswitch="FS00000000000030",module="sensor3(MAIN Board Temp2)",unit="celsius",vdom="root"} 41
   	fortiswitch_health_temperature{fortiswitch="FS00000000000030",module="sensor4(MAIN Board Temp3)",unit="celsius",vdom="root"} 30.25
   	fortiswitch_health_temperature{fortiswitch="FS00000000000038",module="sensor1(CPU  Board Temp)",unit="celsius",vdom="root"} 35.437
   	fortiswitch_health_temperature{fortiswitch="FS00000000000038",module="sensor2(MAIN Board Temp1)",unit="celsius",vdom="root"} 55.625
   	fortiswitch_health_temperature{fortiswitch="FS00000000000038",module="sensor3(MAIN Board Temp2)",unit="celsius",vdom="root"} 43.125
   	fortiswitch_health_temperature{fortiswitch="FS00000000000038",module="sensor4(MAIN Board Temp3)",unit="celsius",vdom="root"} 32.312
    `
	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}
