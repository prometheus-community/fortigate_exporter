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

func TestSystemGlobalLocation(t *testing.T) {
	c := newFakeClient()
	c.prepare("api/v2/cmdb/system/global", "testdata/system-global-location.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !testProbe(probeSystemGlobalLocation, c, r) {
		t.Errorf("probeSystemGlobalLocation() returned non-success")
	}

	em := `
	# HELP fortigate_location_info System geographic location (static metadata)
	# TYPE fortigate_location_info gauge
	fortigate_location_info{latitude="66.543508", longitude="25.8467468"} 1
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}
