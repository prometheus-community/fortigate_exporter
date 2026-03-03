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

func probeSystemVdomResource(c http.FortiHTTP, _ *TargetMetadata) ([]prometheus.Metric, bool) {
	vdomDesc := make(map[string]*prometheus.Desc)
	vdomDesc["cpu"] = prometheus.NewDesc(
		"fortigate_vdom_resource_cpu_usage_ratio",
		"Current VDOM CPU usage in percentage",
		[]string{"vdom"}, nil,
	)
	vdomDesc["memory"] = prometheus.NewDesc(
		"fortigate_vdom_resource_memory_usage_ratio",
		"Current VDOM memory usage in percentage",
		[]string{"vdom"}, nil,
	)
	vdomDesc["setup_rate"] = prometheus.NewDesc(
		"fortigate_vdom_resource_setup_ratio",
		"Current VDOM memory usage in percentage",
		[]string{"vdom"}, nil,
	)
	vdomDesc["is_deletable"] = prometheus.NewDesc(
		"fortigate_vdom_resource_deletable",
		"1 if VDOM is deletable",
		[]string{"vdom"}, nil,
	)
	vdomDesc["id"] = prometheus.NewDesc(
		"fortigate_vdom_resource_object_id",
		"Object Resource ID",
		[]string{"vdom", "object"}, nil,
	)
	vdomDesc["custom_max"] = prometheus.NewDesc(
		"fortigate_vdom_resource_object_custom_max",
		"Object Custom Max",
		[]string{"vdom", "object"}, nil,
	)
	vdomDesc["min_custom_value"] = prometheus.NewDesc(
		"fortigate_vdom_resource_object_custom_min_value",
		"Object Minimum custom value",
		[]string{"vdom", "object"}, nil,
	)
	vdomDesc["max_custom_value"] = prometheus.NewDesc(
		"fortigate_vdom_resource_object_custom_max_value",
		"Object Maximum custom value",
		[]string{"vdom", "object"}, nil,
	)
	vdomDesc["guaranteed"] = prometheus.NewDesc(
		"fortigate_vdom_resource_object_guaranteed",
		"Object Guaranteed",
		[]string{"vdom", "object"}, nil,
	)
	vdomDesc["min_guaranteed_value"] = prometheus.NewDesc(
		"fortigate_vdom_resource_object_guaranteed_max_value",
		"Object Minimum guaranteed value",
		[]string{"vdom", "object"}, nil,
	)
	vdomDesc["max_guaranteed_value"] = prometheus.NewDesc(
		"fortigate_vdom_resource_object_guaranteed_min_value",
		"Object Maximum guaranteed value",
		[]string{"vdom", "object"}, nil,
	)
	vdomDesc["global_max"] = prometheus.NewDesc(
		"fortigate_vdom_resource_object_global_max",
		"Object Global max",
		[]string{"vdom", "object"}, nil,
	)
	vdomDesc["current_usage"] = prometheus.NewDesc(
		"fortigate_vdom_resource_object_current_usage",
		"Object Current usage",
		[]string{"vdom", "object"}, nil,
	)
	vdomDesc["usage_percent"] = prometheus.NewDesc(
		"fortigate_vdom_resource_object_usage_ratio",
		"Object Usage percentage",
		[]string{"vdom", "object"}, nil,
	)

	type VDOMResourceResult struct {
		Result any    `json:"results"`
		Vdom   string `json:"vdom"`
	}

	var res []VDOMResourceResult
	if err := c.Get("api/v2/monitor/system/vdom-resource", "vdom=*", &res); err != nil {
		log.Printf("Error: %v", err)
		return nil, false
	}

	m := []prometheus.Metric{}
	for _, result := range res {
		for k, elem := range result.Result.(map[string]any) {
			switch k {
			case "cpu", "memory", "setup_rate":
				m = append(m, prometheus.MustNewConstMetric(vdomDesc[k], prometheus.GaugeValue, elem.(float64), result.Vdom))
			case "is_deletable":
				if elem.(bool) {
					m = append(m, prometheus.MustNewConstMetric(vdomDesc[k], prometheus.GaugeValue, 1, result.Vdom))
				} else {
					m = append(m, prometheus.MustNewConstMetric(vdomDesc[k], prometheus.GaugeValue, 0, result.Vdom))
				}
			case "session",
				"ipsec-phase1",
				"ipsec-phase2",
				"ipsec-phase1-interface",
				"ipsec-phase2-interface",
				"dialup-tunnel",
				"firewall-policy",
				"firewall-address",
				"firewall-addrgrp",
				"custom-service",
				"service-group",
				"onetime-schedule",
				"recurring-schedule",
				"user",
				"user-group",
				"sslvpn",
				"proxy",
				"log-disk-quota":
				for val, e := range elem.(map[string]any) {
					m = append(m, prometheus.MustNewConstMetric(vdomDesc[val], prometheus.GaugeValue, e.(float64), result.Vdom, k))
				}
			default:
				log.Printf("Missing handler for: %s", k)
			}
		}
	}
	return m, true
}
