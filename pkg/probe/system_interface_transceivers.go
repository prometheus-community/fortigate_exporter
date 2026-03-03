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

func probeSystemInterfaceTransceivers(c http.FortiHTTP, _ *TargetMetadata) ([]prometheus.Metric, bool) {
	interfaceTransceivers := prometheus.NewDesc(
		"fortigate_inteface_transceivers_info",
		"Interface transceivers information",
		[]string{"interface", "type", "vendor", "vendorpartnumber", "vendorserialnumber"}, nil,
	)

	type SystemInterfaceTransceiversResult struct {
		Interface          string `json:"interface"`
		Type               string `json:"type"`
		Vendor             string `json:"vendor"`
		VendorPartNumber   string `json:"vendor_part_number"`
		VendorSerialNumber string `json:"vendor_serial_number"`
	}

	type SystemInterfaceTransceivers struct {
		Results []SystemInterfaceTransceiversResult `json:"results"`
	}

	var res SystemInterfaceTransceivers
	if err := c.Get("api/v2/monitor/system/interface/transceivers", "scope=global", &res); err != nil {
		log.Printf("Warning: %v", err)
		return nil, false
	}

	m := []prometheus.Metric{}
	for _, r := range res.Results {
		m = append(m, prometheus.MustNewConstMetric(interfaceTransceivers, prometheus.GaugeValue, 1.0, r.Interface, r.Type, r.Vendor, r.VendorPartNumber, r.VendorSerialNumber))
	}

	return m, true
}
