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

func TestSystemSensorInfo(t *testing.T) {
	c := newFakeClient()
	c.prepare("api/v2/monitor/system/sensor-info", "testdata/system-sensor-info.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !testProbe(probeSystemSensorInfo, c, r) {
		t.Errorf("probeSystemSensorInfo() returned non-success")
	}

	em := `
	# HELP fortigate_sensor_alarm_status Sensor alarm status
	# TYPE fortigate_sensor_alarm_status gauge
	fortigate_sensor_alarm_status{name="+12V"} 0
	fortigate_sensor_alarm_status{name="+3.3VSB"} 0
	fortigate_sensor_alarm_status{name="+3.3VSB_SMC"} 0
	fortigate_sensor_alarm_status{name="3VDD"} 0
	fortigate_sensor_alarm_status{name="CPU 0 Core 0"} 0
	fortigate_sensor_alarm_status{name="CPU 0 Core 1"} 0
	fortigate_sensor_alarm_status{name="CPU 0 Core 2"} 0
	fortigate_sensor_alarm_status{name="CPU 0 Core 3"} 0
	fortigate_sensor_alarm_status{name="CPU 0 Core 4"} 0
	fortigate_sensor_alarm_status{name="CPU 0 Core 5"} 0
	fortigate_sensor_alarm_status{name="CPU 0 Core 6"} 0
	fortigate_sensor_alarm_status{name="CPU 0 Core 7"} 0
	fortigate_sensor_alarm_status{name="CPU 1 Core 0"} 0
	fortigate_sensor_alarm_status{name="CPU 1 Core 1"} 0
	fortigate_sensor_alarm_status{name="CPU 1 Core 2"} 0
	fortigate_sensor_alarm_status{name="CPU 1 Core 3"} 0
	fortigate_sensor_alarm_status{name="CPU 1 Core 4"} 0
	fortigate_sensor_alarm_status{name="CPU 1 Core 5"} 0
	fortigate_sensor_alarm_status{name="CPU 1 Core 6"} 0
	fortigate_sensor_alarm_status{name="CPU 1 Core 7"} 0
	fortigate_sensor_alarm_status{name="CPU0 PVCCIN"} 0
	fortigate_sensor_alarm_status{name="CPU1 PVCCIN"} 0
	fortigate_sensor_alarm_status{name="DTS CPU0"} 0
	fortigate_sensor_alarm_status{name="DTS CPU1"} 0
	fortigate_sensor_alarm_status{name="FAN1"} 0
	fortigate_sensor_alarm_status{name="FAN2"} 0
	fortigate_sensor_alarm_status{name="FAN3"} 0
	fortigate_sensor_alarm_status{name="FAN4"} 0
	fortigate_sensor_alarm_status{name="FAN5"} 0
	fortigate_sensor_alarm_status{name="FAN6"} 0
	fortigate_sensor_alarm_status{name="MAC_1.025V"} 0
	fortigate_sensor_alarm_status{name="MAC_AVS 1V"} 0
	fortigate_sensor_alarm_status{name="P1V05_PCH"} 0
	fortigate_sensor_alarm_status{name="P3V3_AUX"} 0
	fortigate_sensor_alarm_status{name="PS1 Fan 1"} 0
	fortigate_sensor_alarm_status{name="PS1 Status"} 0
	fortigate_sensor_alarm_status{name="PS1 Temp"} 0
	fortigate_sensor_alarm_status{name="PS1 VIN"} 0
	fortigate_sensor_alarm_status{name="PS1 VOUT_12V"} 0
	fortigate_sensor_alarm_status{name="PS2 Fan 1"} 0
	fortigate_sensor_alarm_status{name="PS2 Status"} 1
	fortigate_sensor_alarm_status{name="PS2 Temp"} 0
	fortigate_sensor_alarm_status{name="PS2 VIN"} 0
	fortigate_sensor_alarm_status{name="PS2 VOUT_12V"} 0
	fortigate_sensor_alarm_status{name="PVCCIO"} 0
	fortigate_sensor_alarm_status{name="PVDDQ AB"} 0
	fortigate_sensor_alarm_status{name="PVDDQ EF"} 0
	fortigate_sensor_alarm_status{name="PVTT AB"} 0
	fortigate_sensor_alarm_status{name="PVTT CD"} 0
	fortigate_sensor_alarm_status{name="PVTT GH"} 0
	fortigate_sensor_alarm_status{name="TD1"} 0
	fortigate_sensor_alarm_status{name="TD2"} 0
	fortigate_sensor_alarm_status{name="TD3"} 0
	fortigate_sensor_alarm_status{name="TD4"} 0
	fortigate_sensor_alarm_status{name="TS1"} 0
	fortigate_sensor_alarm_status{name="TS2"} 0
	fortigate_sensor_alarm_status{name="TS3"} 0
	fortigate_sensor_alarm_status{name="TS4"} 0
	fortigate_sensor_alarm_status{name="TS5"} 0
	fortigate_sensor_alarm_status{name="VCC1.15V"} 0
	fortigate_sensor_alarm_status{name="VCC2.5V"} 0
	fortigate_sensor_alarm_status{name="VCC3V3"} 0
	fortigate_sensor_alarm_status{name="VCC5V"} 0
	# HELP fortigate_sensor_fan_rpm Sensor fan rotation speed in RPM
	# TYPE fortigate_sensor_fan_rpm gauge
	fortigate_sensor_fan_rpm{name="FAN1"} 2900
	fortigate_sensor_fan_rpm{name="FAN2"} 2400
	fortigate_sensor_fan_rpm{name="FAN3"} 3000
	fortigate_sensor_fan_rpm{name="FAN4"} 2500
	fortigate_sensor_fan_rpm{name="FAN5"} 2900
	fortigate_sensor_fan_rpm{name="FAN6"} 2600
	fortigate_sensor_fan_rpm{name="PS1 Fan 1"} 4096
	fortigate_sensor_fan_rpm{name="PS2 Fan 1"} 4224
	# HELP fortigate_sensor_temperature_celsius Sensor temperature in degree celsius
	# TYPE fortigate_sensor_temperature_celsius gauge
	fortigate_sensor_temperature_celsius{name="CPU 0 Core 0"} 40
	fortigate_sensor_temperature_celsius{name="CPU 0 Core 1"} 42
	fortigate_sensor_temperature_celsius{name="CPU 0 Core 2"} 42
	fortigate_sensor_temperature_celsius{name="CPU 0 Core 3"} 41
	fortigate_sensor_temperature_celsius{name="CPU 0 Core 4"} 43
	fortigate_sensor_temperature_celsius{name="CPU 0 Core 5"} 41
	fortigate_sensor_temperature_celsius{name="CPU 0 Core 6"} 44
	fortigate_sensor_temperature_celsius{name="CPU 0 Core 7"} 43
	fortigate_sensor_temperature_celsius{name="CPU 1 Core 0"} 41
	fortigate_sensor_temperature_celsius{name="CPU 1 Core 1"} 42
	fortigate_sensor_temperature_celsius{name="CPU 1 Core 2"} 42
	fortigate_sensor_temperature_celsius{name="CPU 1 Core 3"} 41
	fortigate_sensor_temperature_celsius{name="CPU 1 Core 4"} 43
	fortigate_sensor_temperature_celsius{name="CPU 1 Core 5"} 41
	fortigate_sensor_temperature_celsius{name="CPU 1 Core 6"} 44
	fortigate_sensor_temperature_celsius{name="CPU 1 Core 7"} 43
	fortigate_sensor_temperature_celsius{name="DTS CPU0"} 47
	fortigate_sensor_temperature_celsius{name="DTS CPU1"} 49
	fortigate_sensor_temperature_celsius{name="PS1 Temp"} 25
	fortigate_sensor_temperature_celsius{name="PS2 Temp"} 25
	fortigate_sensor_temperature_celsius{name="TD1"} 31
	fortigate_sensor_temperature_celsius{name="TD2"} 38
	fortigate_sensor_temperature_celsius{name="TD3"} 27
	fortigate_sensor_temperature_celsius{name="TD4"} 30
	fortigate_sensor_temperature_celsius{name="TS1"} 31
	fortigate_sensor_temperature_celsius{name="TS2"} 31
	fortigate_sensor_temperature_celsius{name="TS3"} 32
	fortigate_sensor_temperature_celsius{name="TS4"} 32
	fortigate_sensor_temperature_celsius{name="TS5"} 31
	# HELP fortigate_sensor_thresholds Sensor threasholds
	# TYPE fortigate_sensor_thresholds gauge
	fortigate_sensor_thresholds{name="+12V",threshold="LowerCrit"} 10.307
	fortigate_sensor_thresholds{name="+12V",threshold="LowerNonCrit"} 10.897
	fortigate_sensor_thresholds{name="+12V",threshold="LowerNonRec"} 9.953
	fortigate_sensor_thresholds{name="+12V",threshold="UpperCrit"} 13.729
	fortigate_sensor_thresholds{name="+12V",threshold="UpperNonCrit"} 12.962
	fortigate_sensor_thresholds{name="+12V",threshold="UpperNonRec"} 14.083
	fortigate_sensor_thresholds{name="+3.3VSB",threshold="LowerCrit"} 2.928
	fortigate_sensor_thresholds{name="+3.3VSB",threshold="LowerNonCrit"} 3.072
	fortigate_sensor_thresholds{name="+3.3VSB",threshold="LowerNonRec"} 2.88
	fortigate_sensor_thresholds{name="+3.3VSB",threshold="UpperCrit"} 3.648
	fortigate_sensor_thresholds{name="+3.3VSB",threshold="UpperNonCrit"} 3.552
	fortigate_sensor_thresholds{name="+3.3VSB",threshold="UpperNonRec"} 3.696
	fortigate_sensor_thresholds{name="+3.3VSB_SMC",threshold="LowerCrit"} 2.928
	fortigate_sensor_thresholds{name="+3.3VSB_SMC",threshold="LowerNonCrit"} 3.072
	fortigate_sensor_thresholds{name="+3.3VSB_SMC",threshold="LowerNonRec"} 2.88
	fortigate_sensor_thresholds{name="+3.3VSB_SMC",threshold="UpperCrit"} 3.648
	fortigate_sensor_thresholds{name="+3.3VSB_SMC",threshold="UpperNonCrit"} 3.552
	fortigate_sensor_thresholds{name="+3.3VSB_SMC",threshold="UpperNonRec"} 3.696
	fortigate_sensor_thresholds{name="3VDD",threshold="LowerCrit"} 2.928
	fortigate_sensor_thresholds{name="3VDD",threshold="LowerNonCrit"} 3.072
	fortigate_sensor_thresholds{name="3VDD",threshold="LowerNonRec"} 2.88
	fortigate_sensor_thresholds{name="3VDD",threshold="UpperCrit"} 3.648
	fortigate_sensor_thresholds{name="3VDD",threshold="UpperNonCrit"} 3.552
	fortigate_sensor_thresholds{name="3VDD",threshold="UpperNonRec"} 3.696
	fortigate_sensor_thresholds{name="CPU 0 Core 0",threshold="UpperCrit"} 100
	fortigate_sensor_thresholds{name="CPU 0 Core 0",threshold="UpperNonCrit"} 95
	fortigate_sensor_thresholds{name="CPU 0 Core 0",threshold="UpperNonRec"} 105
	fortigate_sensor_thresholds{name="CPU 0 Core 1",threshold="UpperCrit"} 100
	fortigate_sensor_thresholds{name="CPU 0 Core 1",threshold="UpperNonCrit"} 95
	fortigate_sensor_thresholds{name="CPU 0 Core 1",threshold="UpperNonRec"} 105
	fortigate_sensor_thresholds{name="CPU 0 Core 3",threshold="UpperCrit"} 100
	fortigate_sensor_thresholds{name="CPU 0 Core 3",threshold="UpperNonCrit"} 95
	fortigate_sensor_thresholds{name="CPU 0 Core 3",threshold="UpperNonRec"} 105
	fortigate_sensor_thresholds{name="CPU 0 Core 4",threshold="UpperCrit"} 100
	fortigate_sensor_thresholds{name="CPU 0 Core 4",threshold="UpperNonCrit"} 95
	fortigate_sensor_thresholds{name="CPU 0 Core 4",threshold="UpperNonRec"} 105
	fortigate_sensor_thresholds{name="CPU 0 Core 5",threshold="UpperCrit"} 100
	fortigate_sensor_thresholds{name="CPU 0 Core 5",threshold="UpperNonCrit"} 95
	fortigate_sensor_thresholds{name="CPU 0 Core 5",threshold="UpperNonRec"} 105
	fortigate_sensor_thresholds{name="CPU 0 Core 6",threshold="UpperCrit"} 100
	fortigate_sensor_thresholds{name="CPU 0 Core 6",threshold="UpperNonCrit"} 95
	fortigate_sensor_thresholds{name="CPU 0 Core 6",threshold="UpperNonRec"} 105
	fortigate_sensor_thresholds{name="CPU 0 Core 7",threshold="UpperCrit"} 100
	fortigate_sensor_thresholds{name="CPU 0 Core 7",threshold="UpperNonCrit"} 95
	fortigate_sensor_thresholds{name="CPU 0 Core 7",threshold="UpperNonRec"} 105
	fortigate_sensor_thresholds{name="CPU 1 Core 0",threshold="UpperCrit"} 100
	fortigate_sensor_thresholds{name="CPU 1 Core 0",threshold="UpperNonCrit"} 95
	fortigate_sensor_thresholds{name="CPU 1 Core 0",threshold="UpperNonRec"} 105
	fortigate_sensor_thresholds{name="CPU 1 Core 1",threshold="UpperCrit"} 100
	fortigate_sensor_thresholds{name="CPU 1 Core 1",threshold="UpperNonCrit"} 95
	fortigate_sensor_thresholds{name="CPU 1 Core 1",threshold="UpperNonRec"} 105
	fortigate_sensor_thresholds{name="CPU 1 Core 2",threshold="UpperCrit"} 100
	fortigate_sensor_thresholds{name="CPU 1 Core 2",threshold="UpperNonCrit"} 95
	fortigate_sensor_thresholds{name="CPU 1 Core 2",threshold="UpperNonRec"} 105
	fortigate_sensor_thresholds{name="CPU 1 Core 3",threshold="UpperCrit"} 100
	fortigate_sensor_thresholds{name="CPU 1 Core 3",threshold="UpperNonCrit"} 95
	fortigate_sensor_thresholds{name="CPU 1 Core 3",threshold="UpperNonRec"} 105
	fortigate_sensor_thresholds{name="CPU 1 Core 4",threshold="UpperCrit"} 100
	fortigate_sensor_thresholds{name="CPU 1 Core 4",threshold="UpperNonCrit"} 95
	fortigate_sensor_thresholds{name="CPU 1 Core 4",threshold="UpperNonRec"} 105
	fortigate_sensor_thresholds{name="CPU 1 Core 5",threshold="UpperCrit"} 100
	fortigate_sensor_thresholds{name="CPU 1 Core 5",threshold="UpperNonCrit"} 95
	fortigate_sensor_thresholds{name="CPU 1 Core 5",threshold="UpperNonRec"} 105
	fortigate_sensor_thresholds{name="CPU 1 Core 6",threshold="UpperCrit"} 100
	fortigate_sensor_thresholds{name="CPU 1 Core 6",threshold="UpperNonCrit"} 95
	fortigate_sensor_thresholds{name="CPU 1 Core 6",threshold="UpperNonRec"} 105
	fortigate_sensor_thresholds{name="CPU 1 Core 7",threshold="UpperCrit"} 100
	fortigate_sensor_thresholds{name="CPU 1 Core 7",threshold="UpperNonCrit"} 95
	fortigate_sensor_thresholds{name="CPU 1 Core 7",threshold="UpperNonRec"} 105
	fortigate_sensor_thresholds{name="CPU0 PVCCIN",threshold="LowerCrit"} 1.344
	fortigate_sensor_thresholds{name="CPU0 PVCCIN",threshold="LowerNonCrit"} 1.392
	fortigate_sensor_thresholds{name="CPU0 PVCCIN",threshold="LowerNonRec"} 1.312
	fortigate_sensor_thresholds{name="CPU0 PVCCIN",threshold="UpperCrit"} 2.032
	fortigate_sensor_thresholds{name="CPU0 PVCCIN",threshold="UpperNonCrit"} 1.968
	fortigate_sensor_thresholds{name="CPU0 PVCCIN",threshold="UpperNonRec"} 2.064
	fortigate_sensor_thresholds{name="CPU1 PVCCIN",threshold="LowerCrit"} 1.344
	fortigate_sensor_thresholds{name="CPU1 PVCCIN",threshold="LowerNonCrit"} 1.392
	fortigate_sensor_thresholds{name="CPU1 PVCCIN",threshold="LowerNonRec"} 1.312
	fortigate_sensor_thresholds{name="CPU1 PVCCIN",threshold="UpperCrit"} 2.032
	fortigate_sensor_thresholds{name="CPU1 PVCCIN",threshold="UpperNonCrit"} 1.968
	fortigate_sensor_thresholds{name="CPU1 PVCCIN",threshold="UpperNonRec"} 2.064
	fortigate_sensor_thresholds{name="DTS CPU0",threshold="UpperCrit"} 100
	fortigate_sensor_thresholds{name="DTS CPU0",threshold="UpperNonCrit"} 95
	fortigate_sensor_thresholds{name="DTS CPU0",threshold="UpperNonRec"} 105
	fortigate_sensor_thresholds{name="DTS CPU1",threshold="UpperCrit"} 100
	fortigate_sensor_thresholds{name="DTS CPU1",threshold="UpperNonCrit"} 95
	fortigate_sensor_thresholds{name="DTS CPU1",threshold="UpperNonRec"} 105
	fortigate_sensor_thresholds{name="FAN1",threshold="LowerCrit"} 1000
	fortigate_sensor_thresholds{name="FAN1",threshold="LowerNonCrit"} 1500
	fortigate_sensor_thresholds{name="FAN1",threshold="LowerNonRec"} 500
	fortigate_sensor_thresholds{name="FAN1",threshold="UpperCrit"} 11000
	fortigate_sensor_thresholds{name="FAN1",threshold="UpperNonCrit"} 10000
	fortigate_sensor_thresholds{name="FAN1",threshold="UpperNonRec"} 12000
	fortigate_sensor_thresholds{name="FAN2",threshold="LowerCrit"} 1000
	fortigate_sensor_thresholds{name="FAN2",threshold="LowerNonCrit"} 1500
	fortigate_sensor_thresholds{name="FAN2",threshold="LowerNonRec"} 500
	fortigate_sensor_thresholds{name="FAN2",threshold="UpperCrit"} 11000
	fortigate_sensor_thresholds{name="FAN2",threshold="UpperNonCrit"} 10000
	fortigate_sensor_thresholds{name="FAN2",threshold="UpperNonRec"} 12000
	fortigate_sensor_thresholds{name="FAN3",threshold="LowerCrit"} 1000
	fortigate_sensor_thresholds{name="FAN3",threshold="LowerNonCrit"} 1500
	fortigate_sensor_thresholds{name="FAN3",threshold="LowerNonRec"} 500
	fortigate_sensor_thresholds{name="FAN3",threshold="UpperCrit"} 11000
	fortigate_sensor_thresholds{name="FAN3",threshold="UpperNonCrit"} 10000
	fortigate_sensor_thresholds{name="FAN3",threshold="UpperNonRec"} 12000
	fortigate_sensor_thresholds{name="FAN4",threshold="LowerCrit"} 1000
	fortigate_sensor_thresholds{name="FAN4",threshold="LowerNonCrit"} 1500
	fortigate_sensor_thresholds{name="FAN4",threshold="LowerNonRec"} 500
	fortigate_sensor_thresholds{name="FAN4",threshold="UpperCrit"} 11000
	fortigate_sensor_thresholds{name="FAN4",threshold="UpperNonCrit"} 10000
	fortigate_sensor_thresholds{name="FAN4",threshold="UpperNonRec"} 12000
	fortigate_sensor_thresholds{name="FAN5",threshold="LowerCrit"} 1000
	fortigate_sensor_thresholds{name="FAN5",threshold="LowerNonCrit"} 1500
	fortigate_sensor_thresholds{name="FAN5",threshold="LowerNonRec"} 500
	fortigate_sensor_thresholds{name="FAN5",threshold="UpperCrit"} 11000
	fortigate_sensor_thresholds{name="FAN5",threshold="UpperNonCrit"} 10000
	fortigate_sensor_thresholds{name="FAN5",threshold="UpperNonRec"} 12000
	fortigate_sensor_thresholds{name="FAN6",threshold="LowerCrit"} 1000
	fortigate_sensor_thresholds{name="FAN6",threshold="LowerNonCrit"} 1500
	fortigate_sensor_thresholds{name="FAN6",threshold="LowerNonRec"} 500
	fortigate_sensor_thresholds{name="FAN6",threshold="UpperCrit"} 11000
	fortigate_sensor_thresholds{name="FAN6",threshold="UpperNonCrit"} 10000
	fortigate_sensor_thresholds{name="FAN6",threshold="UpperNonRec"} 12000
	fortigate_sensor_thresholds{name="MAC_1.025V",threshold="LowerCrit"} 0.9192
	fortigate_sensor_thresholds{name="MAC_1.025V",threshold="LowerNonCrit"} 0.9486
	fortigate_sensor_thresholds{name="MAC_1.025V",threshold="LowerNonRec"} 0.8996
	fortigate_sensor_thresholds{name="MAC_1.025V",threshold="UpperCrit"} 1.1348
	fortigate_sensor_thresholds{name="MAC_1.025V",threshold="UpperNonCrit"} 1.1054
	fortigate_sensor_thresholds{name="MAC_1.025V",threshold="UpperNonRec"} 1.1544
	fortigate_sensor_thresholds{name="MAC_AVS 1V",threshold="LowerCrit"} 0.892
	fortigate_sensor_thresholds{name="MAC_AVS 1V",threshold="LowerNonCrit"} 0.9214
	fortigate_sensor_thresholds{name="MAC_AVS 1V",threshold="LowerNonRec"} 0.8724
	fortigate_sensor_thresholds{name="MAC_AVS 1V",threshold="UpperCrit"} 1.1076
	fortigate_sensor_thresholds{name="MAC_AVS 1V",threshold="UpperNonCrit"} 1.0782
	fortigate_sensor_thresholds{name="MAC_AVS 1V",threshold="UpperNonRec"} 1.1272
	fortigate_sensor_thresholds{name="P1V05_PCH",threshold="LowerCrit"} 0.864
	fortigate_sensor_thresholds{name="P1V05_PCH",threshold="LowerNonCrit"} 0.912
	fortigate_sensor_thresholds{name="P1V05_PCH",threshold="LowerNonRec"} 0.816
	fortigate_sensor_thresholds{name="P1V05_PCH",threshold="UpperCrit"} 1.248
	fortigate_sensor_thresholds{name="P1V05_PCH",threshold="UpperNonCrit"} 1.2
	fortigate_sensor_thresholds{name="P1V05_PCH",threshold="UpperNonRec"} 1.296
	fortigate_sensor_thresholds{name="P3V3_AUX",threshold="LowerCrit"} 2.9703
	fortigate_sensor_thresholds{name="P3V3_AUX",threshold="LowerNonCrit"} 3.0844
	fortigate_sensor_thresholds{name="P3V3_AUX",threshold="LowerNonRec"} 2.9051
	fortigate_sensor_thresholds{name="P3V3_AUX",threshold="UpperCrit"} 3.6223
	fortigate_sensor_thresholds{name="P3V3_AUX",threshold="UpperNonCrit"} 3.5245
	fortigate_sensor_thresholds{name="P3V3_AUX",threshold="UpperNonRec"} 3.6875
	fortigate_sensor_thresholds{name="PS1 Fan 1",threshold="LowerCrit"} 256
	fortigate_sensor_thresholds{name="PS1 Fan 1",threshold="LowerNonCrit"} 384
	fortigate_sensor_thresholds{name="PS1 Fan 1",threshold="LowerNonRec"} 128
	fortigate_sensor_thresholds{name="PS1 Fan 1",threshold="UpperCrit"} 25600
	fortigate_sensor_thresholds{name="PS1 Fan 1",threshold="UpperNonCrit"} 24576
	fortigate_sensor_thresholds{name="PS1 Fan 1",threshold="UpperNonRec"} 26880
	fortigate_sensor_thresholds{name="PS1 Temp",threshold="UpperCrit"} 65
	fortigate_sensor_thresholds{name="PS1 Temp",threshold="UpperNonCrit"} 60
	fortigate_sensor_thresholds{name="PS1 Temp",threshold="UpperNonRec"} 70
	fortigate_sensor_thresholds{name="PS1 VIN",threshold="LowerCrit"} 36
	fortigate_sensor_thresholds{name="PS1 VIN",threshold="LowerNonCrit"} 38
	fortigate_sensor_thresholds{name="PS1 VIN",threshold="LowerNonRec"} 34
	fortigate_sensor_thresholds{name="PS1 VIN",threshold="UpperCrit"} 300
	fortigate_sensor_thresholds{name="PS1 VIN",threshold="UpperNonCrit"} 292
	fortigate_sensor_thresholds{name="PS1 VIN",threshold="UpperNonRec"} 306
	fortigate_sensor_thresholds{name="PS1 VOUT_12V",threshold="LowerCrit"} 9.512
	fortigate_sensor_thresholds{name="PS1 VOUT_12V",threshold="LowerNonCrit"} 10.016
	fortigate_sensor_thresholds{name="PS1 VOUT_12V",threshold="LowerNonRec"} 9.197
	fortigate_sensor_thresholds{name="PS1 VOUT_12V",threshold="UpperCrit"} 14.867
	fortigate_sensor_thresholds{name="PS1 VOUT_12V",threshold="UpperNonCrit"} 14.048
	fortigate_sensor_thresholds{name="PS1 VOUT_12V",threshold="UpperNonRec"} 15.245
	fortigate_sensor_thresholds{name="PS2 Fan 1",threshold="LowerCrit"} 256
	fortigate_sensor_thresholds{name="PS2 Fan 1",threshold="LowerNonCrit"} 384
	fortigate_sensor_thresholds{name="PS2 Fan 1",threshold="LowerNonRec"} 128
	fortigate_sensor_thresholds{name="PS2 Fan 1",threshold="UpperCrit"} 25600
	fortigate_sensor_thresholds{name="PS2 Fan 1",threshold="UpperNonCrit"} 24576
	fortigate_sensor_thresholds{name="PS2 Fan 1",threshold="UpperNonRec"} 26880
	fortigate_sensor_thresholds{name="PS2 Temp",threshold="UpperCrit"} 65
	fortigate_sensor_thresholds{name="PS2 Temp",threshold="UpperNonCrit"} 60
	fortigate_sensor_thresholds{name="PS2 Temp",threshold="UpperNonRec"} 70
	fortigate_sensor_thresholds{name="PS2 VIN",threshold="LowerCrit"} 36
	fortigate_sensor_thresholds{name="PS2 VIN",threshold="LowerNonCrit"} 38
	fortigate_sensor_thresholds{name="PS2 VIN",threshold="LowerNonRec"} 34
	fortigate_sensor_thresholds{name="PS2 VIN",threshold="UpperCrit"} 300
	fortigate_sensor_thresholds{name="PS2 VIN",threshold="UpperNonCrit"} 292
	fortigate_sensor_thresholds{name="PS2 VIN",threshold="UpperNonRec"} 306
	fortigate_sensor_thresholds{name="PS2 VOUT_12V",threshold="LowerCrit"} 9.512
	fortigate_sensor_thresholds{name="PS2 VOUT_12V",threshold="LowerNonCrit"} 10.016
	fortigate_sensor_thresholds{name="PS2 VOUT_12V",threshold="LowerNonRec"} 9.197
	fortigate_sensor_thresholds{name="PS2 VOUT_12V",threshold="UpperCrit"} 14.867
	fortigate_sensor_thresholds{name="PS2 VOUT_12V",threshold="UpperNonCrit"} 14.048
	fortigate_sensor_thresholds{name="PS2 VOUT_12V",threshold="UpperNonRec"} 15.245
	fortigate_sensor_thresholds{name="PVCCIO",threshold="LowerCrit"} 0.864
	fortigate_sensor_thresholds{name="PVCCIO",threshold="LowerNonCrit"} 0.896
	fortigate_sensor_thresholds{name="PVCCIO",threshold="LowerNonRec"} 0.848
	fortigate_sensor_thresholds{name="PVCCIO",threshold="UpperCrit"} 1.184
	fortigate_sensor_thresholds{name="PVCCIO",threshold="UpperNonCrit"} 1.152
	fortigate_sensor_thresholds{name="PVCCIO",threshold="UpperNonRec"} 1.2
	fortigate_sensor_thresholds{name="PVDDQ AB",threshold="LowerCrit"} 1.088
	fortigate_sensor_thresholds{name="PVDDQ AB",threshold="LowerNonCrit"} 1.12
	fortigate_sensor_thresholds{name="PVDDQ AB",threshold="LowerNonRec"} 1.056
	fortigate_sensor_thresholds{name="PVDDQ AB",threshold="UpperCrit"} 1.312
	fortigate_sensor_thresholds{name="PVDDQ AB",threshold="UpperNonCrit"} 1.28
	fortigate_sensor_thresholds{name="PVDDQ AB",threshold="UpperNonRec"} 1.344
	fortigate_sensor_thresholds{name="PVDDQ EF",threshold="LowerCrit"} 1.088
	fortigate_sensor_thresholds{name="PVDDQ EF",threshold="LowerNonCrit"} 1.12
	fortigate_sensor_thresholds{name="PVDDQ EF",threshold="LowerNonRec"} 1.056
	fortigate_sensor_thresholds{name="PVDDQ EF",threshold="UpperCrit"} 1.312
	fortigate_sensor_thresholds{name="PVDDQ EF",threshold="UpperNonCrit"} 1.28
	fortigate_sensor_thresholds{name="PVDDQ EF",threshold="UpperNonRec"} 1.344
	fortigate_sensor_thresholds{name="PVTT AB",threshold="LowerCrit"} 0.528
	fortigate_sensor_thresholds{name="PVTT AB",threshold="LowerNonCrit"} 0.56
	fortigate_sensor_thresholds{name="PVTT AB",threshold="LowerNonRec"} 0.512
	fortigate_sensor_thresholds{name="PVTT AB",threshold="UpperCrit"} 0.656
	fortigate_sensor_thresholds{name="PVTT AB",threshold="UpperNonCrit"} 0.64
	fortigate_sensor_thresholds{name="PVTT AB",threshold="UpperNonRec"} 0.672
	fortigate_sensor_thresholds{name="PVTT CD",threshold="LowerCrit"} 0.528
	fortigate_sensor_thresholds{name="PVTT CD",threshold="LowerNonCrit"} 0.56
	fortigate_sensor_thresholds{name="PVTT CD",threshold="LowerNonRec"} 0.512
	fortigate_sensor_thresholds{name="PVTT CD",threshold="UpperCrit"} 0.656
	fortigate_sensor_thresholds{name="PVTT CD",threshold="UpperNonCrit"} 0.64
	fortigate_sensor_thresholds{name="PVTT CD",threshold="UpperNonRec"} 0.672
	fortigate_sensor_thresholds{name="PVTT GH",threshold="LowerCrit"} 0.528
	fortigate_sensor_thresholds{name="PVTT GH",threshold="LowerNonCrit"} 0.56
	fortigate_sensor_thresholds{name="PVTT GH",threshold="LowerNonRec"} 0.512
	fortigate_sensor_thresholds{name="PVTT GH",threshold="UpperCrit"} 0.656
	fortigate_sensor_thresholds{name="PVTT GH",threshold="UpperNonCrit"} 0.64
	fortigate_sensor_thresholds{name="PVTT GH",threshold="UpperNonRec"} 0.672
	fortigate_sensor_thresholds{name="TD1",threshold="UpperCrit"} 70
	fortigate_sensor_thresholds{name="TD1",threshold="UpperNonCrit"} 65
	fortigate_sensor_thresholds{name="TD1",threshold="UpperNonRec"} 75
	fortigate_sensor_thresholds{name="TD2",threshold="UpperCrit"} 75
	fortigate_sensor_thresholds{name="TD2",threshold="UpperNonCrit"} 70
	fortigate_sensor_thresholds{name="TD2",threshold="UpperNonRec"} 80
	fortigate_sensor_thresholds{name="TD3",threshold="UpperCrit"} 65
	fortigate_sensor_thresholds{name="TD3",threshold="UpperNonCrit"} 60
	fortigate_sensor_thresholds{name="TD3",threshold="UpperNonRec"} 70
	fortigate_sensor_thresholds{name="TD4",threshold="UpperCrit"} 65
	fortigate_sensor_thresholds{name="TD4",threshold="UpperNonCrit"} 60
	fortigate_sensor_thresholds{name="TD4",threshold="UpperNonRec"} 70
	fortigate_sensor_thresholds{name="TS1",threshold="UpperCrit"} 70
	fortigate_sensor_thresholds{name="TS1",threshold="UpperNonCrit"} 65
	fortigate_sensor_thresholds{name="TS1",threshold="UpperNonRec"} 75
	fortigate_sensor_thresholds{name="TS2",threshold="UpperCrit"} 70
	fortigate_sensor_thresholds{name="TS2",threshold="UpperNonCrit"} 65
	fortigate_sensor_thresholds{name="TS2",threshold="UpperNonRec"} 75
	fortigate_sensor_thresholds{name="TS3",threshold="UpperCrit"} 70
	fortigate_sensor_thresholds{name="TS3",threshold="UpperNonCrit"} 65
	fortigate_sensor_thresholds{name="TS3",threshold="UpperNonRec"} 75
	fortigate_sensor_thresholds{name="TS4",threshold="UpperCrit"} 70
	fortigate_sensor_thresholds{name="TS4",threshold="UpperNonCrit"} 65
	fortigate_sensor_thresholds{name="TS4",threshold="UpperNonRec"} 75
	fortigate_sensor_thresholds{name="TS5",threshold="UpperCrit"} 70
	fortigate_sensor_thresholds{name="TS5",threshold="UpperNonCrit"} 65
	fortigate_sensor_thresholds{name="TS5",threshold="UpperNonRec"} 75
	fortigate_sensor_thresholds{name="VCC1.15V",threshold="LowerCrit"} 1.0307
	fortigate_sensor_thresholds{name="VCC1.15V",threshold="LowerNonCrit"} 1.0699
	fortigate_sensor_thresholds{name="VCC1.15V",threshold="LowerNonRec"} 1.0111
	fortigate_sensor_thresholds{name="VCC1.15V",threshold="UpperCrit"} 1.2659
	fortigate_sensor_thresholds{name="VCC1.15V",threshold="UpperNonCrit"} 1.2365
	fortigate_sensor_thresholds{name="VCC1.15V",threshold="UpperNonRec"} 1.2855
	fortigate_sensor_thresholds{name="VCC2.5V",threshold="LowerCrit"} 2.2709
	fortigate_sensor_thresholds{name="VCC2.5V",threshold="LowerNonCrit"} 2.3447
	fortigate_sensor_thresholds{name="VCC2.5V",threshold="LowerNonRec"} 2.2094
	fortigate_sensor_thresholds{name="VCC2.5V",threshold="UpperCrit"} 2.7383
	fortigate_sensor_thresholds{name="VCC2.5V",threshold="UpperNonCrit"} 2.6645
	fortigate_sensor_thresholds{name="VCC2.5V",threshold="UpperNonRec"} 2.7875
	fortigate_sensor_thresholds{name="VCC3V3",threshold="LowerCrit"} 2.9703
	fortigate_sensor_thresholds{name="VCC3V3",threshold="LowerNonCrit"} 3.0844
	fortigate_sensor_thresholds{name="VCC3V3",threshold="LowerNonRec"} 2.9051
	fortigate_sensor_thresholds{name="VCC3V3",threshold="UpperCrit"} 3.6223
	fortigate_sensor_thresholds{name="VCC3V3",threshold="UpperNonCrit"} 3.5245
	fortigate_sensor_thresholds{name="VCC3V3",threshold="UpperNonRec"} 3.6875
	fortigate_sensor_thresholds{name="VCC5V",threshold="LowerCrit"} 4.46
	fortigate_sensor_thresholds{name="VCC5V",threshold="LowerNonCrit"} 4.607
	fortigate_sensor_thresholds{name="VCC5V",threshold="LowerNonRec"} 4.362
	fortigate_sensor_thresholds{name="VCC5V",threshold="UpperCrit"} 5.538
	fortigate_sensor_thresholds{name="VCC5V",threshold="UpperNonCrit"} 5.391
	fortigate_sensor_thresholds{name="VCC5V",threshold="UpperNonRec"} 5.636
	# HELP fortigate_sensor_voltage_volts Sensor voltage in volts
	# TYPE fortigate_sensor_voltage_volts gauge
	fortigate_sensor_voltage_volts{name="+12V"} 12.077
	fortigate_sensor_voltage_volts{name="+3.3VSB"} 3.264
	fortigate_sensor_voltage_volts{name="+3.3VSB_SMC"} 3.264
	fortigate_sensor_voltage_volts{name="3VDD"} 3.264
	fortigate_sensor_voltage_volts{name="CPU0 PVCCIN"} 1.792
	fortigate_sensor_voltage_volts{name="CPU1 PVCCIN"} 1.792
	fortigate_sensor_voltage_volts{name="MAC_1.025V"} 1.027
	fortigate_sensor_voltage_volts{name="MAC_AVS 1V"} 0.99
	fortigate_sensor_voltage_volts{name="P1V05_PCH"} 1.008
	fortigate_sensor_voltage_volts{name="P3V3_AUX"} 3.3126
	fortigate_sensor_voltage_volts{name="PS1 VIN"} 224
	fortigate_sensor_voltage_volts{name="PS1 VOUT_12V"} 12.032
	fortigate_sensor_voltage_volts{name="PS2 VIN"} 226
	fortigate_sensor_voltage_volts{name="PS2 VOUT_12V"} 12.032
	fortigate_sensor_voltage_volts{name="PVCCIO"} 1.04
	fortigate_sensor_voltage_volts{name="PVDDQ AB"} 1.2
	fortigate_sensor_voltage_volts{name="PVDDQ EF"} 1.2
	fortigate_sensor_voltage_volts{name="PVTT AB"} 0.592
	fortigate_sensor_voltage_volts{name="PVTT CD"} 0.592
	fortigate_sensor_voltage_volts{name="PVTT GH"} 0.592
	fortigate_sensor_voltage_volts{name="VCC1.15V"} 1.1581
	fortigate_sensor_voltage_volts{name="VCC2.5V"} 2.5169
	fortigate_sensor_voltage_volts{name="VCC3V3"} 3.3126
	fortigate_sensor_voltage_volts{name="VCC5V"} 4.999
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}
