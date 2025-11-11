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

	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus-community/fortigate_exporter/pkg/http"
)

type IPPool struct {
	Name      string  `json:"name"`
	IPTotal   int     `json:"natip_total"`
	IPInUse   int     `json:"natip_in_use"`
	Clients   int     `json:"clients"`
	Available float64 `json:"available"`
	Used      int     `json:"used"`
	Total     int     `json:"total"`
	PbaPerIP  int     `json:"pba_per_ip"`
}

type IPPoolResponse struct {
	Results map[string]IPPool `json:"results"`
	VDOM    string            `json:"vdom"`
	Version string            `json:"version"`
}

func probeFirewallIPPool(c http.FortiHTTP, _ *TargetMetadata) ([]prometheus.Metric, bool) {
	mAvailable := prometheus.NewDesc(
		"fortigate_ippool_available_ratio",
		"Percentage available in ippool (0 - 1.0)",
		[]string{"vdom", "name"}, nil,
	)
	mIPUsed := prometheus.NewDesc(
		"fortigate_ippool_used_ips",
		"Ip addresses in use in ippool",
		[]string{"vdom", "name"}, nil,
	)
	mIPTotal := prometheus.NewDesc(
		"fortigate_ippool_total_ips",
		"Ip addresses total in ippool",
		[]string{"vdom", "name"}, nil,
	)
	mClients := prometheus.NewDesc(
		"fortigate_ippool_clients",
		"Amount of clients using ippool",
		[]string{"vdom", "name"}, nil,
	)
	mUsed := prometheus.NewDesc(
		"fortigate_ippool_used_items",
		"Amount of items used in ippool",
		[]string{"vdom", "name"}, nil,
	)
	mTotal := prometheus.NewDesc(
		"fortigate_ippool_total_items",
		"Amount of items total in ippool",
		[]string{"vdom", "name"}, nil,
	)

	mPbaPerIP := prometheus.NewDesc(
		"fortigate_ippool_pba_per_ip",
		"Amount of available port block allocations per ip",
		[]string{"vdom", "name"}, nil,
	)

	var rs []IPPoolResponse

	if err := c.Get("api/v2/monitor/firewall/ippool", "vdom=*", &rs); err != nil {
		log.Printf("Error: %v", err)
		return nil, false
	}

	m := []prometheus.Metric{}

	for _, r := range rs {
		for _, ippool := range r.Results {
			m = append(m, prometheus.MustNewConstMetric(mAvailable, prometheus.GaugeValue, ippool.Available/100, r.VDOM, ippool.Name))
			m = append(m, prometheus.MustNewConstMetric(mIPUsed, prometheus.GaugeValue, float64(ippool.IPInUse), r.VDOM, ippool.Name))
			m = append(m, prometheus.MustNewConstMetric(mIPTotal, prometheus.GaugeValue, float64(ippool.IPTotal), r.VDOM, ippool.Name))
			m = append(m, prometheus.MustNewConstMetric(mClients, prometheus.GaugeValue, float64(ippool.Clients), r.VDOM, ippool.Name))
			m = append(m, prometheus.MustNewConstMetric(mUsed, prometheus.GaugeValue, float64(ippool.Used), r.VDOM, ippool.Name))
			m = append(m, prometheus.MustNewConstMetric(mTotal, prometheus.GaugeValue, float64(ippool.Total), r.VDOM, ippool.Name))
			m = append(m, prometheus.MustNewConstMetric(mPbaPerIP, prometheus.GaugeValue, float64(ippool.PbaPerIP), r.VDOM, ippool.Name))
		}
	}

	return m, true
}
