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
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestSystemHaPeerOld(t *testing.T) {
	c := newFakeClient()
	c.prepare("api/v2/monitor/system/ha-peer", "testdata/system-ha-peer.jsonnet")
	r := prometheus.NewPedanticRegistry()
	meta := &TargetMetadata{
		VersionMajor: 7,
		VersionMinor: 0,
	}
	if !testProbeWithMetadata(probeSystemHaPeer, c, meta, r) {
		t.Errorf("probeSystemHaPeer() returned non-success")
	}

	em := `
	# HELP fortigate_ha_peer_info Information about the ha peer.
	# TYPE fortigate_ha_peer_info gauge
	fortigate_ha_peer_info{hostname="None",priority="false",serial="None",vcluster="0"} -1
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}

func TestSystemHaPeerAfter74(t *testing.T) {
	c := newFakeClient()
	c.prepare("api/v2/monitor/system/ha-peer", "testdata/system-ha-peer.jsonnet")
	r := prometheus.NewPedanticRegistry()
	meta := &TargetMetadata{
		VersionMajor: 7,
		VersionMinor: 4,
	}
	if !testProbeWithMetadata(probeSystemHaPeer, c, meta, r) {
		t.Errorf("probeSystemHaPeer() returned non-success")
	}

	em := `
	# HELP fortigate_ha_peer_info Information about the ha peer.
	# TYPE fortigate_ha_peer_info gauge
	fortigate_ha_peer_info{hostname="member-name-1",priority="200",serial="FGT61E4QXXXXXXXX1",vcluster="0"} 1
	fortigate_ha_peer_info{hostname="member-name-2",priority="100",serial="FGT61E4QXXXXXXXX2",vcluster="0"} 1
	# HELP fortigate_ha_peer_master True when the peer device is the HA master.
	# TYPE fortigate_ha_peer_master gauge
	fortigate_ha_peer_master{hostname="member-name-1",vcluster="0"} 0
	fortigate_ha_peer_master{hostname="member-name-2",vcluster="0"} 1
	# HELP fortigate_ha_peer_primary True when the peer device is the HA primary.
	# TYPE fortigate_ha_peer_primary gauge
	fortigate_ha_peer_primary{hostname="member-name-1",vcluster="0"} 1
	fortigate_ha_peer_primary{hostname="member-name-2",vcluster="0"} 0
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}
