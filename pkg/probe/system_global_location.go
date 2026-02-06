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
	"encoding/json"
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus-community/fortigate_exporter/pkg/http"
)

func probeSystemGlobalLocation(c http.FortiHTTP, _ *TargetMetadata) ([]prometheus.Metric, bool) {

	location := prometheus.NewDesc(
		"fortigate_location_info",
		"System geographic location (static metadata)",
		[]string{"latitude", "longitude"},
		nil,
	)

	type SystemGlobalLocation struct {
		Latitude  string `json:"gui-device-latitude"`
		Longitude string `json:"gui-device-longitude"`
	}

	type systemGlobalLocationResponse struct {
		Results json.RawMessage `json:"results"`
	}

	var resp systemGlobalLocationResponse

	if err := c.Get("api/v2/cmdb/system/global", "vdom=root", &resp); err != nil {
		log.Printf("Error: %v", err)
		return nil, false
	}

	var loc SystemGlobalLocation

	// Try object first
	if err := json.Unmarshal(resp.Results, &loc); err != nil {
		// Fallback to array
		var arr []SystemGlobalLocation
		if err := json.Unmarshal(resp.Results, &arr); err != nil || len(arr) == 0 {
			return nil, true
		}
		loc = arr[0]
	}

	if loc.Latitude == "" || loc.Longitude == "" {
		return nil, true
	}

	m := []prometheus.Metric{
		prometheus.MustNewConstMetric(
			location,
			prometheus.GaugeValue,
			1,
			loc.Latitude,
			loc.Longitude,
		),
	}
	return m, true
}
