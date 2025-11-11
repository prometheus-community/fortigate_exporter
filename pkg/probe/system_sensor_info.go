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
	"reflect"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus-community/fortigate_exporter/pkg/http"
)

func probeSystemSensorInfo(c http.FortiHTTP, meta *TargetMetadata) ([]prometheus.Metric, bool) {
	var (
		sensorTemperature = prometheus.NewDesc(
			"fortigate_sensor_temperature_celsius",
			"Sensor temperature in degree celsius",
			[]string{"name"}, nil,
		)
		sensorFan = prometheus.NewDesc(
			"fortigate_sensor_fan_rpm",
			"Sensor fan rotation speed in RPM",
			[]string{"name"}, nil,
		)
		sensorVoltage = prometheus.NewDesc(
			"fortigate_sensor_voltage_volts",
			"Sensor voltage in volts",
			[]string{"name"}, nil,
		)
		sensorAlarm = prometheus.NewDesc(
			"fortigate_sensor_alarm_status",
			"Sensor alarm status",
			[]string{"name"}, nil,
		)
		sensorThresholds = prometheus.NewDesc(
			"fortigate_sensor_thresholds",
			"Sensor threasholds",
			[]string{"name", "threshold"}, nil,
		)
	)

	type SystemSensorInfoResultsThresholds struct {
		LowerNonRec  float64 `json:"lower_non_recoverable,omitempty"`
		LowerCrit    float64 `json:"lower_critical,omitempty"`
		LowerNonCrit float64 `json:"lower_non_critical,omitempty"`
		UpperNonCrit float64 `json:"upper_non_critical,omitempty"`
		UpperCrit    float64 `json:"upper_critical,omitempty"`
		UpperNonRec  float64 `json:"upper_non_recoverable,omitempty"`
	}

	type SystemSensorInfoResults struct {
		Name       string                            `json:"name"`
		Type       string                            `json:"type"`
		Value      float64                           `json:"value"`
		Alarm      bool                              `json:"alarm"`
		Thresholds SystemSensorInfoResultsThresholds `json:"thresholds"`
	}

	type SystemSensorInfo struct {
		Results []SystemSensorInfoResults `json:"results"`
	}

	var res SystemSensorInfo
	if err := c.Get("api/v2/monitor/system/sensor-info", "vdom=root", &res); err != nil {
		log.Printf("Warning: %v", err)
		return nil, false
	}

	m := []prometheus.Metric{}
	for _, r := range res.Results {
		switch meta.VersionMajor {
		case 7:
			alarm := 0.0
			if r.Alarm {
				alarm = 1.0
			}
			m = append(m, prometheus.MustNewConstMetric(sensorAlarm, prometheus.GaugeValue, alarm, r.Name))
			v := reflect.ValueOf(r.Thresholds)

			for i := range v.NumField() {
				if v.Field(i).Float() != 0 {
					m = append(m, prometheus.MustNewConstMetric(sensorThresholds, prometheus.GaugeValue, v.Field(i).Float(), r.Name, v.Type().Field(i).Name))
				}
			}
		}
		switch r.Type {
		case "temperature":
			m = append(m, prometheus.MustNewConstMetric(sensorTemperature, prometheus.GaugeValue, r.Value, r.Name))
		case "fan":
			m = append(m, prometheus.MustNewConstMetric(sensorFan, prometheus.GaugeValue, r.Value, r.Name))
		case "voltage":
			m = append(m, prometheus.MustNewConstMetric(sensorVoltage, prometheus.GaugeValue, r.Value, r.Name))
		}
	}

	return m, true
}
