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
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestNetworkDnsLatency(t *testing.T) {
	c := newFakeClient()
	c.prepare("api/v2/monitor/network/dns/latency", "testdata/network-dns-latency.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !testProbe(probeNetworkDNSLatency, c, r) {
		t.Errorf("probeNetworkDNSLatency() returned non-success")
	}

	em := `
	# HELP fortigate_network_dns_latency_seconds Network dns latency
	# TYPE fortigate_network_dns_latency_seconds gauge
	fortigate_network_dns_latency_seconds{ip="8.8.8.8",service="dns_server"} 0.01
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}
