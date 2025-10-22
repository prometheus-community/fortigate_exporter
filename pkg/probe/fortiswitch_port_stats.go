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
	"fmt"
	"github.com/prometheus-community/fortigate_exporter/pkg/http"
	"github.com/prometheus/client_golang/prometheus"
	"log"
)

func probeSwitchPortStats(c http.FortiHTTP, meta *TargetMetadata) ([]prometheus.Metric, bool) {
	var (
		mSwitchStatus = prometheus.NewDesc(
			"fortiswitch_status",
			`Whether the switch is connected or not
# fgt_peer_intf_name = FortiLink interface that the FortiSwitch is connected to
# connection_from = FortiSwitch port that is connected to the FortiLink interface`,
			[]string{"vdom", "name", "serial_number", "fgt_peer_intf_name", "connection_from", "state"}, nil,
		)
		mPortStatus = prometheus.NewDesc(
			"fortiswitch_port_status",
			"Whether the switch port is up or not",
			[]string{"vdom", "name", "serial_number", "interface", "vlan", "duplex"}, nil,
		)
		mPortSpeed = prometheus.NewDesc(
			"fortiswitch_port_speed_bps",
			"Speed negotiated on the interface in bits/s",
			[]string{"vdom", "name", "serial_number", "interface", "vlan", "duplex"}, nil,
		)
		mPortDuplex = prometheus.NewDesc(
			"fortiswitch_port_duplex_status",
			"Duplex status of the FortiSwitch port",
			[]string{"vdom", "name", "serial_number", "interface", "vlan", "duplex"}, nil,
		)
		mTxPkts = prometheus.NewDesc(
			"fortiswitch_port_transmit_packets_total",
			"Number of packets transmitted on the interface",
			[]string{"vdom", "name", "serial_number", "interface"}, nil,
		)
		mTxB = prometheus.NewDesc(
			"fortiswitch_port_transmit_bytes_total",
			"Number of bytes transmitted on the interface",
			[]string{"vdom", "name", "serial_number", "interface"}, nil,
		)
		mTxUcast = prometheus.NewDesc(
			"fortiswitch_port_transmit_unicast_packets_total",
			"Number of unicast packets transmitted on the interface",
			[]string{"vdom", "name", "serial_number", "interface"}, nil,
		)
		mTxMcast = prometheus.NewDesc(
			"fortiswitch_port_transmit_multicast_packets_total",
			"Number of multicast packets transmitted on the interface",
			[]string{"vdom", "name", "serial_number", "interface"}, nil,
		)
		mTxBcast = prometheus.NewDesc(
			"fortiswitch_port_transmit_broadcast_packets_total",
			"Number of broadcast packets transmitted on the interface",
			[]string{"vdom", "name", "serial_number", "interface"}, nil,
		)
		mTxErr = prometheus.NewDesc(
			"fortiswitch_port_transmit_errors_total",
			"Number of transmission errors detected on the interface",
			[]string{"vdom", "name", "serial_number", "interface"}, nil,
		)
		mTxDrops = prometheus.NewDesc(
			"fortiswitch_port_transmit_drops_total",
			"Number of dropped packets detected during transmission on the interface",
			[]string{"vdom", "name", "serial_number", "interface"}, nil,
		)
		mTxOverS = prometheus.NewDesc(
			"fortiswitch_port_transmit_oversized_packets_total",
			"Number of oversized packets transmitted on the interface",
			[]string{"vdom", "name", "serial_number", "interface"}, nil,
		)
		mRxPkts = prometheus.NewDesc(
			"fortiswitch_port_receive_packets_total",
			"Number of packets received on the interface",
			[]string{"vdom", "name", "serial_number", "interface"}, nil,
		)
		mRxB = prometheus.NewDesc(
			"fortiswitch_port_receive_bytes_total",
			"Number of bytes received on the interface",
			[]string{"vdom", "name", "serial_number", "interface"}, nil,
		)
		mRxUcast = prometheus.NewDesc(
			"fortiswitch_port_receive_unicast_packets_total",
			"Number of unicast packets received on the interface",
			[]string{"vdom", "name", "serial_number", "interface"}, nil,
		)
		mRxMcast = prometheus.NewDesc(
			"fortiswitch_port_receive_multicast_packets_total",
			"Number of multicast packets received on the interface",
			[]string{"vdom", "name", "serial_number", "interface"}, nil,
		)
		mRxBcast = prometheus.NewDesc(
			"fortiswitch_port_receive_broadcast_packets_total",
			"Number of broadcast packets received on the interface",
			[]string{"vdom", "name", "serial_number", "interface"}, nil,
		)
		mRxErr = prometheus.NewDesc(
			"fortiswitch_port_receive_errors_total",
			"Number of transmission errors detected on the interface",
			[]string{"vdom", "name", "serial_number", "interface"}, nil,
		)
		mRxDrops = prometheus.NewDesc(
			"fortiswitch_port_receive_drops_total",
			"Number of dropped packets detected during transmission on the interface",
			[]string{"vdom", "name", "serial_number", "interface"}, nil,
		)
		mRxOverS = prometheus.NewDesc(
			"fortiswitch_port_receive_oversized_packets_total",
			"Number of oversized packets received on the interface",
			[]string{"vdom", "name", "serial_number", "interface"}, nil,
		)
	)

	type portStats struct {
		TxPackets  float64 `json:"tx-packets"`
		TxBytes    float64 `json:"tx-bytes"`
		TxErrors   float64 `json:"tx-errors"`
		TxMcast    float64 `json:"tx-mcast"`
		TxUcast    float64 `json:"tx-ucast"`
		TxBcast    float64 `json:"tx-bcast"`
		TxDrops    float64 `json:"tx-drops"`
		TxOversize float64 `json:"tx-oversize"`
		RxPackets  float64 `json:"rx-packets"`
		RxBytes    float64 `json:"rx-bytes"`
		RxErrors   float64 `json:"rx-errors"`
		RxMcast    float64 `json:"rx-mcast"`
		RxUcast    float64 `json:"rx-ucast"`
		RxBcast    float64 `json:"rx-bcast"`
		RxDrops    float64 `json:"rx-drops"`
		RxOversize float64 `json:"rx-oversize"`
	}
	type portsInfo struct {
		Interface string
		Status    string
		Duplex    string
		Speed     float64
		Vlan      string
	}
	type swResult struct {
		Name            string
		Serial          string
		FgPeerIntfName  string `json:"fgt_peer_intf_name"`
		Status          string
		State           string
		Connecting_from string `json:"connecting_from"`
		Vdom            string
		Ports           []portsInfo
		PortStats       map[string]portStats `json:"port_stats"`
	}

	type swPortStats struct {
		Serial string               `json:"serial"`
		Ports  map[string]portStats `json:"ports"`
	}

	type swResponse struct {
		Results []swResult `json:"results"`
	}

	type swPortsNew struct {
		Results []swPortStats `json:"results"`
	}

	var apiPath string
	var portStatsPath string

	if meta.VersionMajor > 7 || (meta.VersionMajor == 7 && meta.VersionMinor >= 0) {
		apiPath = "api/v2/monitor/switch-controller/managed-switch/status"
		portStatsPath = "api/v2/monitor/switch-controller/managed-switch/port-stats"
	} else {
		apiPath = "api/v2/monitor/switch-controller/managed-switch"
		portStatsPath = ""
	}

	var r swResponse

	if portStatsPath == "" {
		// Old API format (FortiOS < 7.0.1)
		if err := c.Get(apiPath, "port_stats=true", &r); err != nil {
			log.Printf("Error: %v", err)
			return nil, false
		}

	} else {
		// New API format (FortiOS 7.0.1+)
		var rPorts swPortsNew

		// Fetch switch status (to get name, vdom, interfaces)
		if err := c.Get(apiPath, "", &r); err != nil {
			log.Printf("Error fetching switch status: %v", err)
			return nil, false
		}

		// Fetch port stats (to get traffic metrics)
		if err := c.Get(portStatsPath, "", &rPorts); err != nil {
			log.Printf("Error fetching switch port stats: %v", err)
			return nil, false
		}

		// Merge `/port-stats` into `/status` response
		for i, sw := range r.Results {
			matched := false
			swr := &r.Results[i]

			// Ensure PortStats is always initialized to prevent nil pointer errors
			if swr.PortStats == nil {
				swr.PortStats = make(map[string]portStats)
			}

			// Match serial numbers and add port stats
			for _, ps := range rPorts.Results {
				if ps.Serial == sw.Serial {
					swr.PortStats = ps.Ports
					matched = true
					break
				}
			}

			// Log if no port stats were found for a switch
			if !matched {
				log.Printf("Warning: No port stats found for switch %s (serial: %s)", swr.Name, swr.Serial)
			}
		}
	}

	m := []prometheus.Metric{}

	for _, swr := range r.Results {
		swStatus := 0.0
		if swr.Status == "Connected" {
			swStatus = 1.0
		}

		var swState float64
		if swr.State == "Authorized" {
			swState = 1.0
		} else if swr.State == "DeAuthorized" {
			swState = 0.0
		} else if swr.State == "Discovered" {
			swState = 2.0
		}

		m = append(m, prometheus.MustNewConstMetric(mSwitchStatus, prometheus.GaugeValue, swStatus, swr.Vdom, swr.Name, swr.Serial, swr.FgPeerIntfName, swr.Connecting_from, fmt.Sprintf("%0.f", swState)))

		for _, pi := range swr.Ports {
			pStatus := 0.0
			if pi.Status == "up" {
				pStatus = 1.0
			}

			pDuplex := 0.0
			if pi.Duplex == "full" {
				pDuplex = 1.0
			} else if pi.Duplex == "half" {
				pDuplex = 0.5
			}
			m = append(m, prometheus.MustNewConstMetric(mPortStatus, prometheus.GaugeValue, pStatus, swr.Vdom, swr.Name, swr.Serial, pi.Interface, pi.Vlan, pi.Duplex))
			m = append(m, prometheus.MustNewConstMetric(mPortSpeed, prometheus.GaugeValue, pi.Speed*1000*1000, swr.Vdom, swr.Name, swr.Serial, pi.Interface, pi.Vlan, pi.Duplex))
			m = append(m, prometheus.MustNewConstMetric(mPortDuplex, prometheus.GaugeValue, pDuplex, swr.Vdom, swr.Name, swr.Serial, pi.Interface, pi.Vlan, pi.Duplex))

		}

		for port, ps := range swr.PortStats {
			m = append(m, prometheus.MustNewConstMetric(mTxPkts, prometheus.CounterValue, ps.TxPackets, swr.Vdom, swr.Name, swr.Serial, port))
			m = append(m, prometheus.MustNewConstMetric(mTxB, prometheus.CounterValue, ps.TxBytes, swr.Vdom, swr.Name, swr.Serial, port))
			m = append(m, prometheus.MustNewConstMetric(mTxUcast, prometheus.CounterValue, ps.TxUcast, swr.Vdom, swr.Name, swr.Serial, port))
			m = append(m, prometheus.MustNewConstMetric(mTxBcast, prometheus.CounterValue, ps.TxBcast, swr.Vdom, swr.Name, swr.Serial, port))
			m = append(m, prometheus.MustNewConstMetric(mTxMcast, prometheus.CounterValue, ps.TxMcast, swr.Vdom, swr.Name, swr.Serial, port))
			m = append(m, prometheus.MustNewConstMetric(mTxErr, prometheus.CounterValue, ps.TxErrors, swr.Vdom, swr.Name, swr.Serial, port))
			m = append(m, prometheus.MustNewConstMetric(mTxDrops, prometheus.CounterValue, ps.TxDrops, swr.Vdom, swr.Name, swr.Serial, port))
			m = append(m, prometheus.MustNewConstMetric(mTxOverS, prometheus.CounterValue, ps.TxOversize, swr.Vdom, swr.Name, swr.Serial, port))
			m = append(m, prometheus.MustNewConstMetric(mRxPkts, prometheus.CounterValue, ps.RxPackets, swr.Vdom, swr.Name, swr.Serial, port))
			m = append(m, prometheus.MustNewConstMetric(mRxB, prometheus.CounterValue, ps.RxBytes, swr.Vdom, swr.Name, swr.Serial, port))
			m = append(m, prometheus.MustNewConstMetric(mRxUcast, prometheus.CounterValue, ps.RxUcast, swr.Vdom, swr.Name, swr.Serial, port))
			m = append(m, prometheus.MustNewConstMetric(mRxBcast, prometheus.CounterValue, ps.RxBcast, swr.Vdom, swr.Name, swr.Serial, port))
			m = append(m, prometheus.MustNewConstMetric(mRxMcast, prometheus.CounterValue, ps.RxMcast, swr.Vdom, swr.Name, swr.Serial, port))
			m = append(m, prometheus.MustNewConstMetric(mRxErr, prometheus.CounterValue, ps.RxErrors, swr.Vdom, swr.Name, swr.Serial, port))
			m = append(m, prometheus.MustNewConstMetric(mRxDrops, prometheus.CounterValue, ps.RxDrops, swr.Vdom, swr.Name, swr.Serial, port))
			m = append(m, prometheus.MustNewConstMetric(mRxOverS, prometheus.CounterValue, ps.RxOversize, swr.Vdom, swr.Name, swr.Serial, port))
		}
	}

	return m, true
}
