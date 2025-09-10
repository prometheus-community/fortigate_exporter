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

	"github.com/prometheus-community/fortigate_exporter/pkg/http"
	"github.com/prometheus/client_golang/prometheus"
)

func probeNetworkDnsLatency(c http.FortiHTTP, meta *TargetMetadata) ([]prometheus.Metric, bool) {
	var (
		dnsLatency = prometheus.NewDesc(
			"fortigate_network_dns_latency",
			"Network dns latency",
			[]string{"service", "ip"}, nil,
		)
		dnsLastUpdate = prometheus.NewDesc(
			"fortigate_network_dns_latest_update",
			"Network dns last update",
			[]string{"service", "ip"}, nil,
		)
	)

	type DnsLatencty struct {
		Service    string  `json:"service"`
		Latency    float64 `json:"latency"`
		LastUpdate float64 `json:"last_update"`
		Ip         string  `json:"ip"`
	}

	type DnsLatencyResult struct {
		Results []DnsLatencty `json:"results"`
	}

	var res DnsLatencyResult
	if err := c.Get("api/v2/monitor/network/dns/latency", "", &res); err != nil {
		log.Printf("Warning: %v", err)
		return nil, false
	}
	m := []prometheus.Metric{}
	for _, r := range res.Results {
		m = append(m, prometheus.MustNewConstMetric(dnsLatency, prometheus.GaugeValue, r.Latency, r.Service, r.Ip))
		m = append(m, prometheus.MustNewConstMetric(dnsLastUpdate, prometheus.GaugeValue, r.LastUpdate, r.Service, r.Ip))
	}

	return m, true
}