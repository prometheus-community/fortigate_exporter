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

	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus-community/fortigate_exporter/pkg/http"
)

func probeNetworkDNSLatency(c http.FortiHTTP, _ *TargetMetadata) ([]prometheus.Metric, bool) {
	dnsLatency := prometheus.NewDesc(
		"fortigate_network_dns_latency_seconds",
		"Network dns latency",
		[]string{"service", "ip"}, nil,
	)

	type DNSLatencty struct {
		Service    string  `json:"service"`
		Latency    float64 `json:"latency"`
		LastUpdate float64 `json:"last_update"`
		IP         string  `json:"ip"`
	}

	type DNSLatencyResult struct {
		Results []DNSLatencty `json:"results"`
	}

	var res DNSLatencyResult
	if err := c.Get("api/v2/monitor/network/dns/latency", "", &res); err != nil {
		log.Printf("Warning: %v", err)
		return nil, false
	}
	m := []prometheus.Metric{}
	for _, r := range res.Results {
		m = append(m, prometheus.MustNewConstMetric(dnsLatency, prometheus.GaugeValue, r.Latency*0.001, r.Service, r.IP))
	}

	return m, true
}
