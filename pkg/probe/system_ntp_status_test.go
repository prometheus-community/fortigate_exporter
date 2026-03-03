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

func TestSystemNtpStatus_7_4_0(t *testing.T) {
	c := newFakeClient()
	c.prepare("api/v2/monitor/system/ntp/status", "testdata/system-ntp-status.jsonnet")
	meta := &TargetMetadata{
		VersionMajor: 7,
		VersionMinor: 4,
	}
	r := prometheus.NewPedanticRegistry()
	if !testProbeWithMetadata(probeSystemNtpStatus, c, meta, r) {
		t.Errorf("probeSystemNtpStatus() returned non-success")
	}

	em := `# HELP fortigate_system_ntp_delay_seconds NTP round trip delay, in seconds
	# TYPE fortigate_system_ntp_delay_seconds gauge
	fortigate_system_ntp_delay_seconds{ip="127.0.0.1",reachable="true",selected="true",server="HA-TEST",vdom="google",version="1035"} 0.324
	fortigate_system_ntp_delay_seconds{ip="127.0.0.1",reachable="true",selected="true",server="HA-TEST",vdom="vdomtest",version="1035"} 0.324
	fortigate_system_ntp_delay_seconds{ip="127.0.0.2",reachable="false",selected="false",server="HA-CAP",vdom="google",version="1035"} 0.324
	fortigate_system_ntp_delay_seconds{ip="127.0.0.2",reachable="false",selected="false",server="HA-CAP",vdom="vdomtest",version="1035"} 0.324
	fortigate_system_ntp_delay_seconds{ip="127.0.0.3",reachable="true",selected="true",server="HA-CODE",vdom="google",version="1035"} 0.324
	fortigate_system_ntp_delay_seconds{ip="127.0.0.3",reachable="true",selected="true",server="HA-CODE",vdom="vdomtest",version="1035"} 0.324
	# HELP fortigate_system_ntp_dispersion_seconds NTP dispersion to primary clock, in seconds
	# TYPE fortigate_system_ntp_dispersion_seconds gauge
	fortigate_system_ntp_dispersion_seconds{ip="127.0.0.1",reachable="true",selected="true",server="HA-TEST",vdom="google",version="1035"} 0.342
	fortigate_system_ntp_dispersion_seconds{ip="127.0.0.1",reachable="true",selected="true",server="HA-TEST",vdom="vdomtest",version="1035"} 0.342
	fortigate_system_ntp_dispersion_seconds{ip="127.0.0.2",reachable="false",selected="false",server="HA-CAP",vdom="google",version="1035"} 0.342
	fortigate_system_ntp_dispersion_seconds{ip="127.0.0.2",reachable="false",selected="false",server="HA-CAP",vdom="vdomtest",version="1035"} 0.342
	fortigate_system_ntp_dispersion_seconds{ip="127.0.0.3",reachable="true",selected="true",server="HA-CODE",vdom="google",version="1035"} 0.342
	fortigate_system_ntp_dispersion_seconds{ip="127.0.0.3",reachable="true",selected="true",server="HA-CODE",vdom="vdomtest",version="1035"} 0.342
	# HELP fortigate_system_ntp_dispersion_peer_seconds NTP peer dispersion, in seconds
	# TYPE fortigate_system_ntp_dispersion_peer_seconds gauge
	fortigate_system_ntp_dispersion_peer_seconds{ip="127.0.0.1",reachable="true",selected="true",server="HA-TEST",vdom="google",version="1035"} 0.123
	fortigate_system_ntp_dispersion_peer_seconds{ip="127.0.0.1",reachable="true",selected="true",server="HA-TEST",vdom="vdomtest",version="1035"} 0.123
	fortigate_system_ntp_dispersion_peer_seconds{ip="127.0.0.2",reachable="false",selected="false",server="HA-CAP",vdom="google",version="1035"} 0.123
	fortigate_system_ntp_dispersion_peer_seconds{ip="127.0.0.2",reachable="false",selected="false",server="HA-CAP",vdom="vdomtest",version="1035"} 0.123
	fortigate_system_ntp_dispersion_peer_seconds{ip="127.0.0.3",reachable="true",selected="true",server="HA-CODE",vdom="google",version="1035"} 0.123
	fortigate_system_ntp_dispersion_peer_seconds{ip="127.0.0.3",reachable="true",selected="true",server="HA-CODE",vdom="vdomtest",version="1035"} 0.123
	# HELP fortigate_system_ntp_expires_seconds NTP expire time, in seconds
	# TYPE fortigate_system_ntp_expires_seconds gauge
	fortigate_system_ntp_expires_seconds{ip="127.0.0.1",reachable="true",selected="true",server="HA-TEST",vdom="google",version="1035"} 145438
	fortigate_system_ntp_expires_seconds{ip="127.0.0.1",reachable="true",selected="true",server="HA-TEST",vdom="vdomtest",version="1035"} 145438
	fortigate_system_ntp_expires_seconds{ip="127.0.0.2",reachable="false",selected="false",server="HA-CAP",vdom="google",version="1035"} 124438
	fortigate_system_ntp_expires_seconds{ip="127.0.0.2",reachable="false",selected="false",server="HA-CAP",vdom="vdomtest",version="1035"} 124438
	fortigate_system_ntp_expires_seconds{ip="127.0.0.3",reachable="true",selected="true",server="HA-CODE",vdom="google",version="1035"} 245438
	fortigate_system_ntp_expires_seconds{ip="127.0.0.3",reachable="true",selected="true",server="HA-CODE",vdom="vdomtest",version="1035"} 245438
	# HELP fortigate_system_ntp_offset_seconds NTP combined offset, in seconds
	# TYPE fortigate_system_ntp_offset_seconds gauge
	fortigate_system_ntp_offset_seconds{ip="127.0.0.1",reachable="true",selected="true",server="HA-TEST",vdom="google",version="1035"} 5.482
	fortigate_system_ntp_offset_seconds{ip="127.0.0.1",reachable="true",selected="true",server="HA-TEST",vdom="vdomtest",version="1035"} 5.482
	fortigate_system_ntp_offset_seconds{ip="127.0.0.2",reachable="false",selected="false",server="HA-CAP",vdom="google",version="1035"} 5.482
	fortigate_system_ntp_offset_seconds{ip="127.0.0.2",reachable="false",selected="false",server="HA-CAP",vdom="vdomtest",version="1035"} 5.482
	fortigate_system_ntp_offset_seconds{ip="127.0.0.3",reachable="true",selected="true",server="HA-CODE",vdom="google",version="1035"} 5.482
	fortigate_system_ntp_offset_seconds{ip="127.0.0.3",reachable="true",selected="true",server="HA-CODE",vdom="vdomtest",version="1035"} 5.482
	# HELP fortigate_system_ntp_reftime_seconds NTP reftime in epoch seconds
	# TYPE fortigate_system_ntp_reftime_seconds counter
	fortigate_system_ntp_reftime_seconds{ip="127.0.0.1",reachable="true",selected="true",server="HA-TEST",vdom="google",version="1035"} 85742
	fortigate_system_ntp_reftime_seconds{ip="127.0.0.1",reachable="true",selected="true",server="HA-TEST",vdom="vdomtest",version="1035"} 85742
	fortigate_system_ntp_reftime_seconds{ip="127.0.0.2",reachable="false",selected="false",server="HA-CAP",vdom="google",version="1035"} 85742
	fortigate_system_ntp_reftime_seconds{ip="127.0.0.2",reachable="false",selected="false",server="HA-CAP",vdom="vdomtest",version="1035"} 85742
	fortigate_system_ntp_reftime_seconds{ip="127.0.0.3",reachable="true",selected="true",server="HA-CODE",vdom="google",version="1035"} 85742
	fortigate_system_ntp_reftime_seconds{ip="127.0.0.3",reachable="true",selected="true",server="HA-CODE",vdom="vdomtest",version="1035"} 85742
	# HELP fortigate_system_ntp_stratum NTP stratum value
	# TYPE fortigate_system_ntp_stratum gauge
	fortigate_system_ntp_stratum{ip="127.0.0.1",reachable="true",selected="true",server="HA-TEST",vdom="google",version="1035"} 3
	fortigate_system_ntp_stratum{ip="127.0.0.1",reachable="true",selected="true",server="HA-TEST",vdom="vdomtest",version="1035"} 3
	fortigate_system_ntp_stratum{ip="127.0.0.2",reachable="false",selected="false",server="HA-CAP",vdom="google",version="1035"} 3
	fortigate_system_ntp_stratum{ip="127.0.0.2",reachable="false",selected="false",server="HA-CAP",vdom="vdomtest",version="1035"} 3
	fortigate_system_ntp_stratum{ip="127.0.0.3",reachable="true",selected="true",server="HA-CODE",vdom="google",version="1035"} 3
	fortigate_system_ntp_stratum{ip="127.0.0.3",reachable="true",selected="true",server="HA-CODE",vdom="vdomtest",version="1035"} 3
	`

	if err := testutil.CollectAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}

func TestSystemNtpStatus(t *testing.T) {
	c := newFakeClient()
	c.prepare("api/v2/monitor/system/ntp/status", "testdata/system-ntp-status.jsonnet")
	meta := &TargetMetadata{
		VersionMajor: 7,
		VersionMinor: 2,
	}
	r := prometheus.NewPedanticRegistry()
	if !testProbeWithMetadata(probeSystemNtpStatus, c, meta, r) {
		t.Errorf("probeSystemNtpStatus() returned non-success")
	}

	em := ``

	if err := testutil.CollectAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}
