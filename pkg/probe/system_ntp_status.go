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

func probeSystemNtpStatus(c http.FortiHTTP, meta *TargetMetadata) ([]prometheus.Metric, bool) {
	var (
		ntpExpires = prometheus.NewDesc(
			"fortigate_system_ntp_expires",
			"NTP expire time, in seconds",
			[]string{"ip", "server", "reachable", "selected", "version", "vdom"}, nil,
		)
		ntpStratum = prometheus.NewDesc(
			"fortigate_system_ntp_stratum",
			"NTP stratum value",
			[]string{"ip", "server", "reachable", "selected", "version", "vdom"}, nil,
		)
		ntpRefTime = prometheus.NewDesc(
			"fortigate_system_ntp_reftime",
			"NTP reftime in epoch seconds",
			[]string{"ip", "server", "reachable", "selected", "version", "vdom"}, nil,
		)
		ntpOffset = prometheus.NewDesc(
			"fortigate_system_ntp_offset",
			"NTP combined offset, in milliseconds",
			[]string{"ip", "server", "reachable", "selected", "version", "vdom"}, nil,
		)
		ntpDelay = prometheus.NewDesc(
			"fortigate_system_ntp_delay",
			"NTP round trip delay, in milliseconds",
			[]string{"ip", "server", "reachable", "selected", "version", "vdom"}, nil,
		)
		ntpDispersion = prometheus.NewDesc(
			"fortigate_system_ntp_dispersion",
			"NTP dispersion to primary clock, in milliseconds",
			[]string{"ip", "server", "reachable", "selected", "version", "vdom"}, nil,
		)
		ntpPeerDispersion = prometheus.NewDesc(
			"fortigate_system_ntp_dispersion_peer",
			"NTP peer dispersion, in milliseconds",
			[]string{"ip", "server", "reachable", "selected", "version", "vdom"}, nil,
		)
	)

	type SystemNtpStatus struct {
		Ip             string  `json:"ip"`
		Server         string  `json:"server"`
		Reachable      bool    `json:"reachable"`
		Expires        int     `json:"expires"`
		Selected       bool    `json:"selected"`
		Version        int     `json:"version"`
		Stratum        int     `json:"stratum"`
		Reftime        int     `json:"reftime"`
		Offset         float64 `json:"offset"`
		Delay          float64 `json:"delay"`
		Dispersion     float64 `json:"dispersion"`
		PeerDispersion int     `json:"peer_dispersion"`
	}

	type SystemNtpStatusResult struct {
		Results []SystemNtpStatus `json:"results"`
		VDOM    string            `json:"vdom"`
	}

	var result []SystemNtpStatusResult
	if err := c.Get("api/v2/monitor/system/ntp/status","vdom=*",&result); err != nil {
		log.Printf("Error: %v", err)
		return nil, false
	}

	m := []prometheus.Metric{}
	if meta.VersionMajor >=7 && meta.VersionMinor >=4 {
		for _, res := range result {
			for _, r := range res.Results {
					m = append(m, prometheus.MustNewConstMetric(ntpExpires, prometheus.GaugeValue, float64(r.Expires), r.Ip, r.Server, strconv.FormatBool(r.Reachable), strconv.FormatBool(r.Reachable), strconv.Itoa(r.Version), res.VDOM))
					m = append(m, prometheus.MustNewConstMetric(ntpStratum, prometheus.GaugeValue, float64(r.Stratum), r.Ip, r.Server, strconv.FormatBool(r.Reachable), strconv.FormatBool(r.Reachable), strconv.Itoa(r.Version), res.VDOM))
					m = append(m, prometheus.MustNewConstMetric(ntpRefTime, prometheus.CounterValue, float64(r.Reftime), r.Ip, r.Server, strconv.FormatBool(r.Reachable), strconv.FormatBool(r.Reachable), strconv.Itoa(r.Version), res.VDOM))
					m = append(m, prometheus.MustNewConstMetric(ntpOffset, prometheus.GaugeValue, float64(r.Offset), r.Ip, r.Server, strconv.FormatBool(r.Reachable), strconv.FormatBool(r.Reachable), strconv.Itoa(r.Version), res.VDOM))
					m = append(m, prometheus.MustNewConstMetric(ntpDelay, prometheus.GaugeValue, float64(r.Delay), r.Ip, r.Server, strconv.FormatBool(r.Reachable), strconv.FormatBool(r.Reachable), strconv.Itoa(r.Version), res.VDOM))
					m = append(m, prometheus.MustNewConstMetric(ntpDispersion, prometheus.GaugeValue, float64(r.Dispersion), r.Ip, r.Server, strconv.FormatBool(r.Reachable), strconv.FormatBool(r.Reachable), strconv.Itoa(r.Version), res.VDOM))
					m = append(m, prometheus.MustNewConstMetric(ntpPeerDispersion, prometheus.GaugeValue, float64(r.PeerDispersion), r.Ip, r.Server, strconv.FormatBool(r.Reachable), strconv.FormatBool(r.Reachable), strconv.Itoa(r.Version), res.VDOM))
			}
		}
	} else {
		log.Printf("Not implemented in versions under 7.4")
	}
	return m, true
}