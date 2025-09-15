package probe

import (
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestSystemInterfaceTransceivers(t *testing.T) {
	c := newFakeClient()
	c.prepare("api/v2/monitor/system/interface/transceivers", "testdata/interface-transceivers.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !testProbe(probeSystemInterfaceTransceivers, c, r) {
		t.Errorf("probeSystemInterfaceTransceivers() returned non-success")
	}

	em := `
	# HELP fortigate_inteface_transceivers Interface transceivers information
	# TYPE fortigate_inteface_transceivers gauge
	fortigate_inteface_transceivers{interface="ha1",type="SFP/SFP+/SFP28",vendor="FORTINET",vendorpartnumber="FTL",vendorserialnumber="U00000"} 1
	fortigate_inteface_transceivers{interface="ha2",type="SFP/SFP+/SFP28",vendor="FORTINET",vendorpartnumber="FTL",vendorserialnumber="U00000"} 1
	fortigate_inteface_transceivers{interface="port33",type="QSFP/QSFP+",vendor="FORTINET",vendorpartnumber="FTL",vendorserialnumber="U00000"} 1
	fortigate_inteface_transceivers{interface="port34",type="QSFP/QSFP+",vendor="FORTINET",vendorpartnumber="FTL",vendorserialnumber="U00000"} 1
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}
