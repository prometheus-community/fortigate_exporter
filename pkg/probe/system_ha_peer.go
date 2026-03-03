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
	"strconv"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus-community/fortigate_exporter/pkg/http"
)

func probeSystemHaPeer(c http.FortiHTTP, meta *TargetMetadata) ([]prometheus.Metric, bool) {
	info := prometheus.NewDesc(
		"fortigate_ha_peer_info",
		"Information about the ha peer.",
		[]string{"serial", "vcluster", "hostname", "priority"}, nil,
	)

	primary := prometheus.NewDesc(
		"fortigate_ha_peer_primary",
		"True when the peer device is the HA primary.",
		[]string{"vcluster", "hostname"}, nil,
	)

	master := prometheus.NewDesc(
		"fortigate_ha_peer_master",
		"True when the peer device is the HA master.",
		[]string{"vcluster", "hostname"}, nil,
	)

	type SystemHaPeer struct {
		Serial   string  `json:"serial_no"`
		Vcluster int64   `json:"vcluster_id"`
		Priority float64 `json:"priority"`
		Hostname string  `json:"hostname"`
		Master   bool    `json:"master"`
		Primary  bool    `json:"primary"`
	}

	type SystemHaPeerResult struct {
		Result []SystemHaPeer `json:"results"`
	}

	var res SystemHaPeerResult
	if err := c.Get("api/v2/monitor/system/ha-peer", "", &res); err != nil {
		log.Printf("Warning: %v", err)
		return nil, false
	}
	m := []prometheus.Metric{}
	for _, r := range res.Result {
		if meta.VersionMajor >= 7 && meta.VersionMinor >= 4 {
			m = append(m, prometheus.MustNewConstMetric(info, prometheus.GaugeValue, 1, r.Serial, strconv.FormatInt(r.Vcluster, 10), r.Hostname, strconv.FormatFloat(r.Priority, 'f', -1, 64)))
			if r.Primary {
				m = append(m, prometheus.MustNewConstMetric(primary, prometheus.GaugeValue, 1, strconv.FormatInt(r.Vcluster, 10), r.Hostname))
			} else {
				m = append(m, prometheus.MustNewConstMetric(primary, prometheus.GaugeValue, 0, strconv.FormatInt(r.Vcluster, 10), r.Hostname))
			}
			if meta.VersionMinor == 4 {
				if r.Master {
					m = append(m, prometheus.MustNewConstMetric(master, prometheus.GaugeValue, 1, strconv.FormatInt(r.Vcluster, 10), r.Hostname))
				} else {
					m = append(m, prometheus.MustNewConstMetric(master, prometheus.GaugeValue, 0, strconv.FormatInt(r.Vcluster, 10), r.Hostname))
				}
			}
		} else {
			m = append(m, prometheus.MustNewConstMetric(info, prometheus.GaugeValue, -1, "None", "0", "None", "false"))
			break
		}
	}
	return m, true
}
