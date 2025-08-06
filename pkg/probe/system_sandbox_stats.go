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

func probeSystemSandboxStats (c http.FortiHTTP, meta *TargetMetadata) ([]prometheus.Metric, bool) {
	var (
		numberDetected = prometheus.NewDesc(
			"fortigate_sandbox_stats_detected",
			"Number of detected files",
			[]string{}, nil,
		)
		numberClean = prometheus.NewDesc(
			"fortigate_sandbox_stats_clean",
			"Number of clean files",
			[]string{}, nil,
		)
		numberRiskLow = prometheus.NewDesc(
			"fortigate_sandbox_stats_risk_low",
			"Number of low risk files detected",
			[]string{}, nil,
		)
		numberRiskMedium = prometheus.NewDesc(
			"fortigate_sandbox_stats_risk_medium",
			"Number of medium risk files detected",
			[]string{}, nil,
		)
		numberRiskHigh = prometheus.NewDesc(
			"fortigate_sandbox_stats_risk_high",
			"Number of high risk files detected",
			[]string{}, nil,
		)
		numberSubmitted = prometheus.NewDesc(
			"fortigate_sandbox_stats_submitted",
			"Number of submitted files",
			[]string{}, nil,
		)
	)

	type SystemSandboxStats struct {
		Detected  float64
		Clean     float64
		Low       float64 `json:"risk_low"`
		Medium    float64 `json:"risk_med"`
		High      float64 `json:"risk_high"`
		Submitted float64 
	}

	type SystemSandboxStatsResult struct {
		Results []SystemSandboxStats `json:"results"`
	}

	var res SystemSandboxStatsResult
	if err := c.Get("api/v2/monitor/system/sandbox/stats","", &res); err != nil {
		log.Printf("Warning: %v", err)
		return nil, false
	}
	var m = []prometheus.Metric{}

	for _, r := range res.Results {
		m = append(m, prometheus.MustNewConstMetric(numberDetected, prometheus.CounterValue, r.Detected))
		m = append(m, prometheus.MustNewConstMetric(numberClean, prometheus.CounterValue, r.Clean))
		m = append(m, prometheus.MustNewConstMetric(numberRiskLow, prometheus.CounterValue, r.Low))
		m = append(m, prometheus.MustNewConstMetric(numberRiskMedium, prometheus.CounterValue, r.Medium))
		m = append(m, prometheus.MustNewConstMetric(numberRiskHigh, prometheus.CounterValue, r.High))
		m = append(m, prometheus.MustNewConstMetric(numberSubmitted, prometheus.CounterValue, r.Submitted))
	}
	return m, true
}