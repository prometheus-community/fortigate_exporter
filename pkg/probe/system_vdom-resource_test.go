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

func TestVdomResource(t *testing.T) {
	c := newFakeClient()
	c.prepare("api/v2/monitor/system/vdom-resource", "testdata/vdom-resource.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !testProbe(probeSystemVdomResource, c, r) {
		t.Errorf("probeSystemVdomResource() returned non-success")
	}

	em := `
	# HELP fortigate_vdom_resource_cpu_usage_ratio Current VDOM CPU usage in percentage
	# TYPE fortigate_vdom_resource_cpu_usage_ratio gauge
	fortigate_vdom_resource_cpu_usage_ratio{vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_cpu_usage_ratio{vdom="kalle-test"} 0
	# HELP fortigate_vdom_resource_deletable 1 if VDOM is deletable
	# TYPE fortigate_vdom_resource_deletable gauge
	fortigate_vdom_resource_deletable{vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_deletable{vdom="kalle-test"} 1
	# HELP fortigate_vdom_resource_memory_usage_ratio Current VDOM memory usage in percentage
	# TYPE fortigate_vdom_resource_memory_usage_ratio gauge
	fortigate_vdom_resource_memory_usage_ratio{vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_memory_usage_ratio{vdom="kalle-test"} 0
	# HELP fortigate_vdom_resource_object_current_usage Object Current usage
	# TYPE fortigate_vdom_resource_object_current_usage gauge
	fortigate_vdom_resource_object_current_usage{object="custom-service",vdom="gur-proddmzext"} 87
	fortigate_vdom_resource_object_current_usage{object="custom-service",vdom="kalle-test"} 88
	fortigate_vdom_resource_object_current_usage{object="dialup-tunnel",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_current_usage{object="dialup-tunnel",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_current_usage{object="firewall-address",vdom="gur-proddmzext"} 14
	fortigate_vdom_resource_object_current_usage{object="firewall-address",vdom="kalle-test"} 10
	fortigate_vdom_resource_object_current_usage{object="firewall-addrgrp",vdom="gur-proddmzext"} 1
	fortigate_vdom_resource_object_current_usage{object="firewall-addrgrp",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_current_usage{object="firewall-policy",vdom="gur-proddmzext"} 2
	fortigate_vdom_resource_object_current_usage{object="firewall-policy",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_current_usage{object="ipsec-phase1",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_current_usage{object="ipsec-phase1",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_current_usage{object="ipsec-phase1-interface",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_current_usage{object="ipsec-phase1-interface",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_current_usage{object="ipsec-phase2",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_current_usage{object="ipsec-phase2",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_current_usage{object="ipsec-phase2-interface",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_current_usage{object="ipsec-phase2-interface",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_current_usage{object="log-disk-quota",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_current_usage{object="log-disk-quota",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_current_usage{object="onetime-schedule",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_current_usage{object="onetime-schedule",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_current_usage{object="proxy",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_current_usage{object="proxy",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_current_usage{object="recurring-schedule",vdom="gur-proddmzext"} 2
	fortigate_vdom_resource_object_current_usage{object="recurring-schedule",vdom="kalle-test"} 2
	fortigate_vdom_resource_object_current_usage{object="service-group",vdom="gur-proddmzext"} 4
	fortigate_vdom_resource_object_current_usage{object="service-group",vdom="kalle-test"} 4
	fortigate_vdom_resource_object_current_usage{object="session",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_current_usage{object="session",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_current_usage{object="sslvpn",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_current_usage{object="sslvpn",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_current_usage{object="user",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_current_usage{object="user",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_current_usage{object="user-group",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_current_usage{object="user-group",vdom="kalle-test"} 0
	# HELP fortigate_vdom_resource_object_custom_max Object Custom Max
	# TYPE fortigate_vdom_resource_object_custom_max gauge
	fortigate_vdom_resource_object_custom_max{object="custom-service",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_custom_max{object="custom-service",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_custom_max{object="dialup-tunnel",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_custom_max{object="dialup-tunnel",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_custom_max{object="firewall-address",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_custom_max{object="firewall-address",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_custom_max{object="firewall-addrgrp",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_custom_max{object="firewall-addrgrp",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_custom_max{object="firewall-policy",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_custom_max{object="firewall-policy",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_custom_max{object="ipsec-phase1",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_custom_max{object="ipsec-phase1",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_custom_max{object="ipsec-phase1-interface",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_custom_max{object="ipsec-phase1-interface",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_custom_max{object="ipsec-phase2",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_custom_max{object="ipsec-phase2",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_custom_max{object="ipsec-phase2-interface",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_custom_max{object="ipsec-phase2-interface",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_custom_max{object="log-disk-quota",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_custom_max{object="log-disk-quota",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_custom_max{object="onetime-schedule",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_custom_max{object="onetime-schedule",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_custom_max{object="proxy",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_custom_max{object="proxy",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_custom_max{object="recurring-schedule",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_custom_max{object="recurring-schedule",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_custom_max{object="service-group",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_custom_max{object="service-group",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_custom_max{object="session",vdom="gur-proddmzext"} 200000
	fortigate_vdom_resource_object_custom_max{object="session",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_custom_max{object="sslvpn",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_custom_max{object="sslvpn",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_custom_max{object="user",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_custom_max{object="user",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_custom_max{object="user-group",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_custom_max{object="user-group",vdom="kalle-test"} 0
	# HELP fortigate_vdom_resource_object_custom_max_value Object Maximum custom value
	# TYPE fortigate_vdom_resource_object_custom_max_value gauge
	fortigate_vdom_resource_object_custom_max_value{object="custom-service",vdom="gur-proddmzext"} 10240
	fortigate_vdom_resource_object_custom_max_value{object="custom-service",vdom="kalle-test"} 10240
	fortigate_vdom_resource_object_custom_max_value{object="dialup-tunnel",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_custom_max_value{object="dialup-tunnel",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_custom_max_value{object="firewall-address",vdom="gur-proddmzext"} 198192
	fortigate_vdom_resource_object_custom_max_value{object="firewall-address",vdom="kalle-test"} 198192
	fortigate_vdom_resource_object_custom_max_value{object="firewall-addrgrp",vdom="gur-proddmzext"} 33192
	fortigate_vdom_resource_object_custom_max_value{object="firewall-addrgrp",vdom="kalle-test"} 33192
	fortigate_vdom_resource_object_custom_max_value{object="firewall-policy",vdom="gur-proddmzext"} 201024
	fortigate_vdom_resource_object_custom_max_value{object="firewall-policy",vdom="kalle-test"} 201024
	fortigate_vdom_resource_object_custom_max_value{object="ipsec-phase1",vdom="gur-proddmzext"} 20000
	fortigate_vdom_resource_object_custom_max_value{object="ipsec-phase1",vdom="kalle-test"} 20000
	fortigate_vdom_resource_object_custom_max_value{object="ipsec-phase1-interface",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_custom_max_value{object="ipsec-phase1-interface",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_custom_max_value{object="ipsec-phase2",vdom="gur-proddmzext"} 20000
	fortigate_vdom_resource_object_custom_max_value{object="ipsec-phase2",vdom="kalle-test"} 20000
	fortigate_vdom_resource_object_custom_max_value{object="ipsec-phase2-interface",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_custom_max_value{object="ipsec-phase2-interface",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_custom_max_value{object="log-disk-quota",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_custom_max_value{object="log-disk-quota",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_custom_max_value{object="onetime-schedule",vdom="gur-proddmzext"} 5000
	fortigate_vdom_resource_object_custom_max_value{object="onetime-schedule",vdom="kalle-test"} 5000
	fortigate_vdom_resource_object_custom_max_value{object="proxy",vdom="gur-proddmzext"} 64000
	fortigate_vdom_resource_object_custom_max_value{object="proxy",vdom="kalle-test"} 64000
	fortigate_vdom_resource_object_custom_max_value{object="recurring-schedule",vdom="gur-proddmzext"} 1024
	fortigate_vdom_resource_object_custom_max_value{object="recurring-schedule",vdom="kalle-test"} 1024
	fortigate_vdom_resource_object_custom_max_value{object="service-group",vdom="gur-proddmzext"} 4000
	fortigate_vdom_resource_object_custom_max_value{object="service-group",vdom="kalle-test"} 4000
	fortigate_vdom_resource_object_custom_max_value{object="session",vdom="gur-proddmzext"} 2.2e+07
	fortigate_vdom_resource_object_custom_max_value{object="session",vdom="kalle-test"} 2.2e+07
	fortigate_vdom_resource_object_custom_max_value{object="sslvpn",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_custom_max_value{object="sslvpn",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_custom_max_value{object="user",vdom="gur-proddmzext"} 5000
	fortigate_vdom_resource_object_custom_max_value{object="user",vdom="kalle-test"} 5000
	fortigate_vdom_resource_object_custom_max_value{object="user-group",vdom="gur-proddmzext"} 2000
	fortigate_vdom_resource_object_custom_max_value{object="user-group",vdom="kalle-test"} 2000
	# HELP fortigate_vdom_resource_object_custom_min_value Object Minimum custom value
	# TYPE fortigate_vdom_resource_object_custom_min_value gauge
	fortigate_vdom_resource_object_custom_min_value{object="custom-service",vdom="gur-proddmzext"} 87
	fortigate_vdom_resource_object_custom_min_value{object="custom-service",vdom="kalle-test"} 88
	fortigate_vdom_resource_object_custom_min_value{object="dialup-tunnel",vdom="gur-proddmzext"} 1
	fortigate_vdom_resource_object_custom_min_value{object="dialup-tunnel",vdom="kalle-test"} 1
	fortigate_vdom_resource_object_custom_min_value{object="firewall-address",vdom="gur-proddmzext"} 14
	fortigate_vdom_resource_object_custom_min_value{object="firewall-address",vdom="kalle-test"} 10
	fortigate_vdom_resource_object_custom_min_value{object="firewall-addrgrp",vdom="gur-proddmzext"} 1
	fortigate_vdom_resource_object_custom_min_value{object="firewall-addrgrp",vdom="kalle-test"} 1
	fortigate_vdom_resource_object_custom_min_value{object="firewall-policy",vdom="gur-proddmzext"} 2
	fortigate_vdom_resource_object_custom_min_value{object="firewall-policy",vdom="kalle-test"} 1
	fortigate_vdom_resource_object_custom_min_value{object="ipsec-phase1",vdom="gur-proddmzext"} 1
	fortigate_vdom_resource_object_custom_min_value{object="ipsec-phase1",vdom="kalle-test"} 1
	fortigate_vdom_resource_object_custom_min_value{object="ipsec-phase1-interface",vdom="gur-proddmzext"} 1
	fortigate_vdom_resource_object_custom_min_value{object="ipsec-phase1-interface",vdom="kalle-test"} 1
	fortigate_vdom_resource_object_custom_min_value{object="ipsec-phase2",vdom="gur-proddmzext"} 1
	fortigate_vdom_resource_object_custom_min_value{object="ipsec-phase2",vdom="kalle-test"} 1
	fortigate_vdom_resource_object_custom_min_value{object="ipsec-phase2-interface",vdom="gur-proddmzext"} 1
	fortigate_vdom_resource_object_custom_min_value{object="ipsec-phase2-interface",vdom="kalle-test"} 1
	fortigate_vdom_resource_object_custom_min_value{object="log-disk-quota",vdom="gur-proddmzext"} 100
	fortigate_vdom_resource_object_custom_min_value{object="log-disk-quota",vdom="kalle-test"} 100
	fortigate_vdom_resource_object_custom_min_value{object="onetime-schedule",vdom="gur-proddmzext"} 1
	fortigate_vdom_resource_object_custom_min_value{object="onetime-schedule",vdom="kalle-test"} 1
	fortigate_vdom_resource_object_custom_min_value{object="proxy",vdom="gur-proddmzext"} 1
	fortigate_vdom_resource_object_custom_min_value{object="proxy",vdom="kalle-test"} 1
	fortigate_vdom_resource_object_custom_min_value{object="recurring-schedule",vdom="gur-proddmzext"} 2
	fortigate_vdom_resource_object_custom_min_value{object="recurring-schedule",vdom="kalle-test"} 2
	fortigate_vdom_resource_object_custom_min_value{object="service-group",vdom="gur-proddmzext"} 4
	fortigate_vdom_resource_object_custom_min_value{object="service-group",vdom="kalle-test"} 4
	fortigate_vdom_resource_object_custom_min_value{object="session",vdom="gur-proddmzext"} 1
	fortigate_vdom_resource_object_custom_min_value{object="session",vdom="kalle-test"} 1
	fortigate_vdom_resource_object_custom_min_value{object="sslvpn",vdom="gur-proddmzext"} 1
	fortigate_vdom_resource_object_custom_min_value{object="sslvpn",vdom="kalle-test"} 1
	fortigate_vdom_resource_object_custom_min_value{object="user",vdom="gur-proddmzext"} 1
	fortigate_vdom_resource_object_custom_min_value{object="user",vdom="kalle-test"} 1
	fortigate_vdom_resource_object_custom_min_value{object="user-group",vdom="gur-proddmzext"} 1
	fortigate_vdom_resource_object_custom_min_value{object="user-group",vdom="kalle-test"} 1
	# HELP fortigate_vdom_resource_object_global_max Object Global max
	# TYPE fortigate_vdom_resource_object_global_max gauge
	fortigate_vdom_resource_object_global_max{object="custom-service",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_global_max{object="custom-service",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_global_max{object="dialup-tunnel",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_global_max{object="dialup-tunnel",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_global_max{object="firewall-address",vdom="gur-proddmzext"} 198192
	fortigate_vdom_resource_object_global_max{object="firewall-address",vdom="kalle-test"} 198192
	fortigate_vdom_resource_object_global_max{object="firewall-addrgrp",vdom="gur-proddmzext"} 33192
	fortigate_vdom_resource_object_global_max{object="firewall-addrgrp",vdom="kalle-test"} 33192
	fortigate_vdom_resource_object_global_max{object="firewall-policy",vdom="gur-proddmzext"} 201024
	fortigate_vdom_resource_object_global_max{object="firewall-policy",vdom="kalle-test"} 201024
	fortigate_vdom_resource_object_global_max{object="ipsec-phase1",vdom="gur-proddmzext"} 20000
	fortigate_vdom_resource_object_global_max{object="ipsec-phase1",vdom="kalle-test"} 20000
	fortigate_vdom_resource_object_global_max{object="ipsec-phase1-interface",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_global_max{object="ipsec-phase1-interface",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_global_max{object="ipsec-phase2",vdom="gur-proddmzext"} 20000
	fortigate_vdom_resource_object_global_max{object="ipsec-phase2",vdom="kalle-test"} 20000
	fortigate_vdom_resource_object_global_max{object="ipsec-phase2-interface",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_global_max{object="ipsec-phase2-interface",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_global_max{object="log-disk-quota",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_global_max{object="log-disk-quota",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_global_max{object="onetime-schedule",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_global_max{object="onetime-schedule",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_global_max{object="proxy",vdom="gur-proddmzext"} 64000
	fortigate_vdom_resource_object_global_max{object="proxy",vdom="kalle-test"} 64000
	fortigate_vdom_resource_object_global_max{object="recurring-schedule",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_global_max{object="recurring-schedule",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_global_max{object="service-group",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_global_max{object="service-group",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_global_max{object="session",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_global_max{object="session",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_global_max{object="sslvpn",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_global_max{object="sslvpn",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_global_max{object="user",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_global_max{object="user",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_global_max{object="user-group",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_global_max{object="user-group",vdom="kalle-test"} 0
	# HELP fortigate_vdom_resource_object_guaranteed Object Guaranteed
	# TYPE fortigate_vdom_resource_object_guaranteed gauge
	fortigate_vdom_resource_object_guaranteed{object="custom-service",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed{object="custom-service",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed{object="dialup-tunnel",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed{object="dialup-tunnel",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed{object="firewall-address",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed{object="firewall-address",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed{object="firewall-addrgrp",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed{object="firewall-addrgrp",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed{object="firewall-policy",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed{object="firewall-policy",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed{object="ipsec-phase1",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed{object="ipsec-phase1",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed{object="ipsec-phase1-interface",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed{object="ipsec-phase1-interface",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed{object="ipsec-phase2",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed{object="ipsec-phase2",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed{object="ipsec-phase2-interface",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed{object="ipsec-phase2-interface",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed{object="log-disk-quota",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed{object="log-disk-quota",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed{object="onetime-schedule",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed{object="onetime-schedule",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed{object="proxy",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed{object="proxy",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed{object="recurring-schedule",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed{object="recurring-schedule",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed{object="service-group",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed{object="service-group",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed{object="session",vdom="gur-proddmzext"} 200000
	fortigate_vdom_resource_object_guaranteed{object="session",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed{object="sslvpn",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed{object="sslvpn",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed{object="user",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed{object="user",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed{object="user-group",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed{object="user-group",vdom="kalle-test"} 0
	# HELP fortigate_vdom_resource_object_guaranteed_max_value Object Minimum guaranteed value
	# TYPE fortigate_vdom_resource_object_guaranteed_max_value gauge
	fortigate_vdom_resource_object_guaranteed_max_value{object="custom-service",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="custom-service",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="dialup-tunnel",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="dialup-tunnel",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="firewall-address",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="firewall-address",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="firewall-addrgrp",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="firewall-addrgrp",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="firewall-policy",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="firewall-policy",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="ipsec-phase1",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="ipsec-phase1",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="ipsec-phase1-interface",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="ipsec-phase1-interface",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="ipsec-phase2",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="ipsec-phase2",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="ipsec-phase2-interface",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="ipsec-phase2-interface",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="log-disk-quota",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="log-disk-quota",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="onetime-schedule",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="onetime-schedule",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="proxy",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="proxy",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="recurring-schedule",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="recurring-schedule",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="service-group",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="service-group",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="session",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="session",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="sslvpn",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="sslvpn",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="user",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="user",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="user-group",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed_max_value{object="user-group",vdom="kalle-test"} 0
	# HELP fortigate_vdom_resource_object_guaranteed_min_value Object Maximum guaranteed value
	# TYPE fortigate_vdom_resource_object_guaranteed_min_value gauge
	fortigate_vdom_resource_object_guaranteed_min_value{object="custom-service",vdom="gur-proddmzext"} 10240
	fortigate_vdom_resource_object_guaranteed_min_value{object="custom-service",vdom="kalle-test"} 10240
	fortigate_vdom_resource_object_guaranteed_min_value{object="dialup-tunnel",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed_min_value{object="dialup-tunnel",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed_min_value{object="firewall-address",vdom="gur-proddmzext"} 198094
	fortigate_vdom_resource_object_guaranteed_min_value{object="firewall-address",vdom="kalle-test"} 198090
	fortigate_vdom_resource_object_guaranteed_min_value{object="firewall-addrgrp",vdom="gur-proddmzext"} 33189
	fortigate_vdom_resource_object_guaranteed_min_value{object="firewall-addrgrp",vdom="kalle-test"} 33188
	fortigate_vdom_resource_object_guaranteed_min_value{object="firewall-policy",vdom="gur-proddmzext"} 201009
	fortigate_vdom_resource_object_guaranteed_min_value{object="firewall-policy",vdom="kalle-test"} 201007
	fortigate_vdom_resource_object_guaranteed_min_value{object="ipsec-phase1",vdom="gur-proddmzext"} 20000
	fortigate_vdom_resource_object_guaranteed_min_value{object="ipsec-phase1",vdom="kalle-test"} 20000
	fortigate_vdom_resource_object_guaranteed_min_value{object="ipsec-phase1-interface",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed_min_value{object="ipsec-phase1-interface",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed_min_value{object="ipsec-phase2",vdom="gur-proddmzext"} 20000
	fortigate_vdom_resource_object_guaranteed_min_value{object="ipsec-phase2",vdom="kalle-test"} 20000
	fortigate_vdom_resource_object_guaranteed_min_value{object="ipsec-phase2-interface",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed_min_value{object="ipsec-phase2-interface",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed_min_value{object="log-disk-quota",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed_min_value{object="log-disk-quota",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed_min_value{object="onetime-schedule",vdom="gur-proddmzext"} 5000
	fortigate_vdom_resource_object_guaranteed_min_value{object="onetime-schedule",vdom="kalle-test"} 5000
	fortigate_vdom_resource_object_guaranteed_min_value{object="proxy",vdom="gur-proddmzext"} 64000
	fortigate_vdom_resource_object_guaranteed_min_value{object="proxy",vdom="kalle-test"} 64000
	fortigate_vdom_resource_object_guaranteed_min_value{object="recurring-schedule",vdom="gur-proddmzext"} 1024
	fortigate_vdom_resource_object_guaranteed_min_value{object="recurring-schedule",vdom="kalle-test"} 1024
	fortigate_vdom_resource_object_guaranteed_min_value{object="service-group",vdom="gur-proddmzext"} 4000
	fortigate_vdom_resource_object_guaranteed_min_value{object="service-group",vdom="kalle-test"} 4000
	fortigate_vdom_resource_object_guaranteed_min_value{object="session",vdom="gur-proddmzext"} 200000
	fortigate_vdom_resource_object_guaranteed_min_value{object="session",vdom="kalle-test"} 2.165e+07
	fortigate_vdom_resource_object_guaranteed_min_value{object="sslvpn",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_guaranteed_min_value{object="sslvpn",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_guaranteed_min_value{object="user",vdom="gur-proddmzext"} 5000
	fortigate_vdom_resource_object_guaranteed_min_value{object="user",vdom="kalle-test"} 5000
	fortigate_vdom_resource_object_guaranteed_min_value{object="user-group",vdom="gur-proddmzext"} 2000
	fortigate_vdom_resource_object_guaranteed_min_value{object="user-group",vdom="kalle-test"} 2000
	# HELP fortigate_vdom_resource_object_id Object Resource ID
	# TYPE fortigate_vdom_resource_object_id gauge
	fortigate_vdom_resource_object_id{object="custom-service",vdom="gur-proddmzext"} 9
	fortigate_vdom_resource_object_id{object="custom-service",vdom="kalle-test"} 9
	fortigate_vdom_resource_object_id{object="dialup-tunnel",vdom="gur-proddmzext"} 5
	fortigate_vdom_resource_object_id{object="dialup-tunnel",vdom="kalle-test"} 5
	fortigate_vdom_resource_object_id{object="firewall-address",vdom="gur-proddmzext"} 7
	fortigate_vdom_resource_object_id{object="firewall-address",vdom="kalle-test"} 7
	fortigate_vdom_resource_object_id{object="firewall-addrgrp",vdom="gur-proddmzext"} 8
	fortigate_vdom_resource_object_id{object="firewall-addrgrp",vdom="kalle-test"} 8
	fortigate_vdom_resource_object_id{object="firewall-policy",vdom="gur-proddmzext"} 6
	fortigate_vdom_resource_object_id{object="firewall-policy",vdom="kalle-test"} 6
	fortigate_vdom_resource_object_id{object="ipsec-phase1",vdom="gur-proddmzext"} 1
	fortigate_vdom_resource_object_id{object="ipsec-phase1",vdom="kalle-test"} 1
	fortigate_vdom_resource_object_id{object="ipsec-phase1-interface",vdom="gur-proddmzext"} 3
	fortigate_vdom_resource_object_id{object="ipsec-phase1-interface",vdom="kalle-test"} 3
	fortigate_vdom_resource_object_id{object="ipsec-phase2",vdom="gur-proddmzext"} 2
	fortigate_vdom_resource_object_id{object="ipsec-phase2",vdom="kalle-test"} 2
	fortigate_vdom_resource_object_id{object="ipsec-phase2-interface",vdom="gur-proddmzext"} 4
	fortigate_vdom_resource_object_id{object="ipsec-phase2-interface",vdom="kalle-test"} 4
	fortigate_vdom_resource_object_id{object="log-disk-quota",vdom="gur-proddmzext"} 17
	fortigate_vdom_resource_object_id{object="log-disk-quota",vdom="kalle-test"} 17
	fortigate_vdom_resource_object_id{object="onetime-schedule",vdom="gur-proddmzext"} 11
	fortigate_vdom_resource_object_id{object="onetime-schedule",vdom="kalle-test"} 11
	fortigate_vdom_resource_object_id{object="proxy",vdom="gur-proddmzext"} 16
	fortigate_vdom_resource_object_id{object="proxy",vdom="kalle-test"} 16
	fortigate_vdom_resource_object_id{object="recurring-schedule",vdom="gur-proddmzext"} 12
	fortigate_vdom_resource_object_id{object="recurring-schedule",vdom="kalle-test"} 12
	fortigate_vdom_resource_object_id{object="service-group",vdom="gur-proddmzext"} 10
	fortigate_vdom_resource_object_id{object="service-group",vdom="kalle-test"} 10
	fortigate_vdom_resource_object_id{object="session",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_id{object="session",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_id{object="sslvpn",vdom="gur-proddmzext"} 15
	fortigate_vdom_resource_object_id{object="sslvpn",vdom="kalle-test"} 15
	fortigate_vdom_resource_object_id{object="user",vdom="gur-proddmzext"} 13
	fortigate_vdom_resource_object_id{object="user",vdom="kalle-test"} 13
	fortigate_vdom_resource_object_id{object="user-group",vdom="gur-proddmzext"} 14
	fortigate_vdom_resource_object_id{object="user-group",vdom="kalle-test"} 14
	# HELP fortigate_vdom_resource_object_usage_ratio Object Usage percentage
	# TYPE fortigate_vdom_resource_object_usage_ratio gauge
	fortigate_vdom_resource_object_usage_ratio{object="custom-service",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_usage_ratio{object="custom-service",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_usage_ratio{object="dialup-tunnel",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_usage_ratio{object="dialup-tunnel",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_usage_ratio{object="firewall-address",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_usage_ratio{object="firewall-address",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_usage_ratio{object="firewall-addrgrp",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_usage_ratio{object="firewall-addrgrp",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_usage_ratio{object="firewall-policy",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_usage_ratio{object="firewall-policy",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_usage_ratio{object="ipsec-phase1",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_usage_ratio{object="ipsec-phase1",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_usage_ratio{object="ipsec-phase1-interface",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_usage_ratio{object="ipsec-phase1-interface",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_usage_ratio{object="ipsec-phase2",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_usage_ratio{object="ipsec-phase2",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_usage_ratio{object="ipsec-phase2-interface",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_usage_ratio{object="ipsec-phase2-interface",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_usage_ratio{object="log-disk-quota",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_usage_ratio{object="log-disk-quota",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_usage_ratio{object="onetime-schedule",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_usage_ratio{object="onetime-schedule",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_usage_ratio{object="proxy",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_usage_ratio{object="proxy",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_usage_ratio{object="recurring-schedule",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_usage_ratio{object="recurring-schedule",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_usage_ratio{object="service-group",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_usage_ratio{object="service-group",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_usage_ratio{object="session",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_usage_ratio{object="session",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_usage_ratio{object="sslvpn",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_usage_ratio{object="sslvpn",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_usage_ratio{object="user",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_usage_ratio{object="user",vdom="kalle-test"} 0
	fortigate_vdom_resource_object_usage_ratio{object="user-group",vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_object_usage_ratio{object="user-group",vdom="kalle-test"} 0
	# HELP fortigate_vdom_resource_setup_ratio Current VDOM memory usage in percentage
	# TYPE fortigate_vdom_resource_setup_ratio gauge
	fortigate_vdom_resource_setup_ratio{vdom="gur-proddmzext"} 0
	fortigate_vdom_resource_setup_ratio{vdom="kalle-test"} 0
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}
