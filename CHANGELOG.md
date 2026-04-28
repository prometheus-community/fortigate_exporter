## v1.25.0 / 2026-03-15

First release after the move to the prometheus-community

> [!CAUTION]
> [**BREAKINGCHANGE**] Rewrite BGP probes to use Prometheus state sets instead of numeric values for FortiOS 7.6 (#366)
> this one will **REPLACE** the exisisting `fortigate_bgp_neighbor_ipv(4|6)_info` metric to a stateset `fortigate_bgp_neighbor_ipv(4|6)_state` impacting cardinality by +17 per bgp neighbor

> [!IMPORTANT]
> [CHANGE] Rewrite SDN Connector Status probe to use Prometheus state sets (#368)
> this one will **ADD** a stateset `fortigate_system_sdn_connector_state` impacing cardinality by +5 per sdn connector. 
> 
> **We are planning to deprecate `fortigate_system_sdn_connector_status` in a future release**

- [CHANGE] Rename internal version package (#375)
- [CHANGE] Rename internal HTTP client package (#371)
- [CHANGE] Update project license (#309); add missing license description (#376)
- [CHANGE] Transfer project to Prometheus Community; update CI/CD (#312, #314, #341, #344)
- [CHANGE] Update maintainers list and remove archive notice from README (#372)
- [CHANGE] Refactor README for improved structure and clarity (#349)
- [FEATURE] Add System HA Peer probe (#347)
- [FEATURE] Add System Performance Status probe (#336)
- [FEATURE] Add DNS Latency probe (#334)
- [FEATURE] Add VDOM Resources probe (#338)
- [FEATURE] Add Interface Transceivers probe for FortiOS 7.4+ (#332)
- [FEATURE] Add System NTP Status probe (#335)
- [FEATURE] Add Sensor Info probe for FortiOS 7.4+ (#331)
- [FEATURE] Add Central Management Status probe (#337)
- [FEATURE] Add System Global Location probe (#363)
- [FEATURE] Add FortiOS 7.6 support for system status probe (#367)
- [BUGFIX] Fix managed-switch API endpoint for updated FortiOS API (#324)
- [ENHANCEMENT] Enable additional golangci-lint checks (#346)
- [DEPENDENCY] Bump `golang.org/x/crypto` from 0.36.0 to 0.45.0 (#374)
- [DEPENDENCY] Bump `github.com/prometheus/client_golang` (#339)
- [DEPENDENCY] Bump `github.com/google/go-jsonnet` from 0.20.0 to 0.21.0 (#325)
- [DEPENDENCY] Bump GitHub Actions dependencies (#326, #330, #352, #353, #356, #357, #360, #361, #369)

## v1.24.1 / 2024-01-08

Previous release.
