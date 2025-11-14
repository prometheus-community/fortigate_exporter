# fortigate_exporter metrics

Global:

 * _Network/Dns/Latency_
   * `fortigate_network_dns_latency_`
 * _System/SensorInfo_
   * `fortigate_sensor_alarm_status`
   * `fortigate_sensor_fan_rpm`
   * `fortigate_sensor_temperature_celsius`
   * `fortigate_sensor_voltage_volts`
   * `fortigate_sensor_thresholds`
 * _System/Status_
   * `fortigate_version_info`
 * _System/Transceivers_
   * `fortigate_interface_transceivers`
 * _System/Time/Clock_
   * `fortigate_time_seconds`
 * _System/Resource/Usage_
   * `fortigate_cpu_usage_ratio`
   * `fortigate_memory_usage_ratio`
   * `fortigate_current_sessions`
 * _System/Performance/status/_
   * `fortigate_system_performance_status_cpu_cores_idle_ratio`
   * `fortigate_system_performance_status_cpu_cores_iowait_ratio`
   * `fortigate_system_performance_status_cpu_cores_nice_ratio`
   * `fortigate_system_performance_status_cpu_cores_system_ratio`
   * `fortigate_system_performance_status_cpu_cores_user_ratio`
   * `fortigate_system_performance_status_cpu_idle_ratio`
   * `fortigate_system_performance_status_cpu_iowait_ratio`
   * `fortigate_system_performance_status_cpu_nice_ratio`
   * `fortigate_system_performance_status_cpu_system_ratio`
   * `fortigate_system_performance_status_cpu_user_ratio`
   * `fortigate_system_performance_status_mem_free_bytes`
   * `fortigate_system_performance_status_mem_freeable_bytes`
   * `fortigate_system_performance_status_mem_bytes_total`
   * `fortigate_system_performance_status_mem_used_bytes`
 * _System/HAChecksums_
   * `fortigate_ha_member_has_role`
 * _System/Ntp/Status_
   * `fortigate_system_ntp_delay_seconds`
   * `fortigate_system_ntp_dispersion_seconds`
   * `fortigate_system_ntp_dispersion_peer_seconds`
   * `fortigate_system_ntp_expires_seconds`
   * `fortigate_system_ntp_offset_seconds`
   * `fortigate_system_ntp_reftime_seconds`
   * `fortigate_system_ntp_stratum`
 * _License/Status_
   * `fortigate_license_vdom_usage`
   * `fortigate_license_vdom_max`
 * _WebUI/State_
   * `fortigate_last_reboot_seconds`
   * `fortigate_last_snapshot_seconds`

Per-VDOM:

 * _System/Resource/Usage/VDOM_
   * `fortigate_vdom_cpu_usage_ratio`
   * `fortigate_vdom_memory_usage_ratio`
   * `fortigate_vdom_current_sessions`
 * _Firewall/Policies_
   * `fortigate_policy_active_sessions`
   * `fortigate_policy_bytes_total`
   * `fortigate_policy_hit_count_total`
   * `fortigate_policy_packets_total`
 * _Firewall/IpPool_
   * `fortigate_ippool_available_ratio`
   * `fortigate_ippool_used_ips`
   * `fortigate_ippool_total_ips`
   * `fortigate_ippool_clients`
   * `fortigate_ippool_used_items`
   * `fortigate_ippool_total_items`
   * `fortigate_ippool_pba_per_ip`
 * _System/Fortimanager/Status_
   * `fortigate_fortimanager_connection_status`
   * `fortigate_fortimanager_registration_status`
 * _System/Interface_
   * `fortigate_interface_link_up`
   * `fortigate_interface_speed_bps`
   * `fortigate_interface_transmit_packets_total`
   * `fortigate_interface_receive_packets_total`
   * `fortigate_interface_transmit_bytes_total`
   * `fortigate_interface_receive_bytes_total`
   * `fortigate_interface_transmit_errors_total`
   * `fortigate_interface_receive_errors_total`
 * _System/Interface/Transceivers_
   * `fortigate_inteface_transceivers_info`
 * _System/SDNConnector_
   * `fortigate_system_sdn_connector_status`
   * `fortigate_system_sdn_connector_last_update_seconds`
 * _/System/CentralManagement/Status_
   * `fortigate_system_central_management_mode`
   * `fortigate_system_central_management_status`
   * `fortigate_system_central_management_registration_status`
 * _System/VDOMResource_
   * `fortigate_vdom_resource_cpu_usage`
   * `fortigate_vdom_resource_memory_usage`
   * `fortigate_vdom_resource_setup_rate`
   * `fortigate_vdom_resource_deletable`
   * `fortigate_vdom_resource_object_id`
   * `fortigate_vdom_resource_object_custom_max`
   * `fortigate_vdom_resource_object_custom_min_value`
   * `fortigate_vdom_resource_object_custom_max_value`
   * `fortigate_vdom_resource_object_guaranted`
   * `fortigate_vdom_resource_object_guaranted_max_value`
   * `fortigate_vdom_resource_object_guaranted_min_value`
   * `fortigate_vdom_resource_object_global_max`
   * `fortigate_vdom_resource_object_current_usage`
   * `fortigate_vdom_resource_object_usage_percentage`
 * _User/Fsso_
   * `fortigate_user_fsso_info`
 * _VPN/Ssl/Connections_
   * `fortigate_vpn_connections`
   * `fortigate_vpn_users`
 * _VPN/Ssl/Stats_
   * `fortigate_vpn_ssl_users`
   * `fortigate_vpn_ssl_tunnels`
   * `fortigate_vpn_ssl_connections`
 * _VPN/IPSec_
   * `fortigate_ipsec_tunnel_receive_bytes_total`
   * `fortigate_ipsec_tunnel_transmit_bytes_total`
   * `fortigate_ipsec_tunnel_up`
 * _Wifi/APStatus_
   * `fortigate_wifi_access_points`
   * `fortigate_wifi_fabric_clients`
   * `fortigate_wifi_fabric_max_allowed_clients`
 * _Log/Fortianalyzer/Status_
   * `fortigate_log_fortianalyzer_registration_info`
   * `fortigate_log_fortianalyzer_logs_received`
 * _Log/Fortianalyzer/Queue_
   * `fortigate_log_fortianalyzer_queue_connections`
   * `fortigate_log_fortianalyzer_queue_logs`
 * _Log/DiskUsage_
   * `fortigate_log_disk_used_bytes`
   * `fortigate_log_disk_total_bytes`

 Per-HA-Member and VDOM:
 * _System/HAStatistics_
   * `fortigate_ha_member_info`
   * `fortigate_ha_member_cpu_usage_ratio`
   * `fortigate_ha_member_memory_usage_ratio`
   * `fortigate_ha_member_network_usage_ratio`
   * `fortigate_ha_member_sessions`
   * `fortigate_ha_member_packets_total`
   * `fortigate_ha_member_virus_events_total`
   * `fortigate_ha_member_bytes_total`
   * `fortigate_ha_member_ips_events_total`

 Per-Link and VDOM:
 * _System/LinkMonitor_
   * `fortigate_link_status`
   * `fortigate_link_latency_seconds`
   * `fortigate_link_latency_jitter_seconds`
   * `fortigate_link_packet_loss_ratio`
   * `fortigate_link_packet_sent_total`
   * `fortigate_link_packet_received_total`
   * `fortigate_link_active_sessions`
   * `fortigate_link_bandwidth_tx_byte_per_second`
   * `fortigate_link_bandwidth_rx_byte_per_second`
   * `fortigate_link_status_change_time_seconds`

 Per-SDWAN and VDOM:
 * _VirtualWAN/HealthCheck_
   * `fortigate_virtual_wan_status`
   * `fortigate_virtual_wan_latency_seconds`
   * `fortigate_virtual_wan_latency_jitter_seconds`
   * `fortigate_virtual_wan_packet_loss_ratio`
   * `fortigate_virtual_wan_packet_sent_total`
   * `fortigate_virtual_wan_packet_received_total`
   * `fortigate_virtual_wan_active_sessions`
   * `fortigate_virtual_wan_bandwidth_tx_byte_per_second`
   * `fortigate_virtual_wan_bandwidth_rx_byte_per_second`
   * `fortigate_virtual_wan_status_change_time_seconds`

 Per-BGP-Neighbor and VDOM:
 * _BGP/Neighbors/IPv4_
   * `fortigate_bgp_neighbor_ipv4_info`
 * _BGP/Neighbors/IPv6_
   * `fortigate_bgp_neighbor_ipv6_info`
 * _BGP/NeighborPaths/IPv4_
   * `fortigate_bgp_neighbor_ipv4_paths`
   * `fortigate_bgp_neighbor_ipv4_best_paths`
 * _BGP/NeighborPaths/IPv6_
   * `fortigate_bgp_neighbor_ipv6_paths`
   * `fortigate_bgp_neighbor_ipv6_best_paths`

 Per-OSPF-Neighbor and VDOM:
 * _OSPF/Neighbors_
   * `fortigate_ospf_neighbor_info`

 Per-VirtualServer and VDOM:
 * _Firewall/LoadBalance_
   * `fortigate_lb_virtual_server_info`

 Per-RealServer for each VirtualServer and VDOM:
 * _Firewall/LoadBalance_
   * `fortigate_lb_real_server_info`
   * `fortigate_lb_real_server_mode`
   * `fortigate_lb_real_server_status`
   * `fortigate_lb_real_server_active_sessions`
   * `fortigate_lb_real_server_rtt_seconds`
   * `fortigate_lb_real_server_processed_bytes_total`

 Per-Certificate:
 * _System/AvailableCertificates_
   * `fortigate_certificate_info`
   * `fortigate_certificate_valid_from_seconds`
   * `fortigate_certificate_valid_to_seconds`
   * `fortigate_certificate_cmdb_references`

Per-VDOM and Wifi-Client:
 * _Wifi/Clients_
   * `fortigate_wifi_client_info`
   * `fortigate_wifi_client_data_rate_bps`
   * `fortigate_wifi_client_bandwidth_rx_bps`
   * `fortigate_wifi_client_bandwidth_tx_bps`
   * `fortigate_wifi_client_signal_strength_dBm`
   * `fortigate_wifi_client_signal_noise_dBm`
   * `fortigate_wifi_client_tx_discard_ratio`
   * `fortigate_wifi_client_tx_retries_ratio`

Per-VDOM and managed access point:
 * _Wifi/ManagedAP_
   * `fortigate_wifi_managed_ap_info`
   * `fortigate_wifi_managed_ap_join_time_seconds`
   * `fortigate_wifi_managed_ap_cpu_usage_ratio`
   * `fortigate_wifi_managed_ap_memory_free_bytes`
   * `fortigate_wifi_managed_ap_memory_bytes_total`

Per-VDOM, managed access point and radio:
 * _Wifi/ManagedAP_
   * `fortigate_wifi_managed_ap_radio_info`
   * `fortigate_wifi_managed_ap_radio_client_count`
   * `fortigate_wifi_managed_ap_radio_operating_tx_power_ratio`
   * `fortigate_wifi_managed_ap_radio_operating_channel_utilization_ratio`
   * `fortigate_wifi_managed_ap_radio_bandwidth_rx_bps`
   * `fortigate_wifi_managed_ap_radio_rx_bytes_total`
   * `fortigate_wifi_managed_ap_radio_tx_bytes_total`
   * `fortigate_wifi_managed_ap_radio_interfering_aps`
   * `fortigate_wifi_managed_ap_radio_tx_power_ratio`
   * `fortigate_wifi_managed_ap_radio_tx_discard_ratio`
   * `fortigate_wifi_managed_ap_radio_tx_retries_ratio`

Per-VDOM, managed access point and interface:
 * _Wifi/ManagedAP_
   * `fortigate_wifi_managed_ap_interface_rx_bytes_total`
   * `fortigate_wifi_managed_ap_interface_tx_bytes_total`
   * `fortigate_wifi_managed_ap_interface_rx_packets_total`
   * `fortigate_wifi_managed_ap_interface_tx_packets_total`
   * `fortigate_wifi_managed_ap_interface_rx_errors_total`
   * `fortigate_wifi_managed_ap_interface_tx_errors_total`
   * `fortigate_wifi_managed_ap_interface_rx_dropped_packets_total`
   * `fortigate_wifi_managed_ap_interface_tx_dropped_packets_total`

Per-VDOM, managed switch and interface:
* _Switch/ManagedSwitch_
  * `fortigate_managed_switch_collisions_total`
  * `fortigate_managed_switch_crc_alignments_total`
  * `fortigate_managed_switch_fragments_total`
  * `fortigate_managed_switch_info`
  * `fortigate_managed_switch_jabbers_total`
  * `fortigate_managed_switch_l3_packets_total`
  * `fortigate_managed_switch_max_poe_budget_watt`
  * `fortigate_managed_switch_port_info`
  * `fortigate_managed_switch_port_power_status`
  * `fortigate_managed_switch_port_power_watt`
  * `fortigate_managed_switch_port_status`
  * `fortigate_managed_switch_rx_bcast_packets_total`
  * `fortigate_managed_switch_rx_bytes_total`
  * `fortigate_managed_switch_rx_drops_total`
  * `fortigate_managed_switch_rx_errors_total`
  * `fortigate_managed_switch_rx_mcast_packets_total`
  * `fortigate_managed_switch_rx_oversize_total`
  * `fortigate_managed_switch_rx_packets_total`
  * `fortigate_managed_switch_rx_ucast_packets_total`
  * `fortigate_managed_switch_tx_bcast_packets_total`
  * `fortigate_managed_switch_tx_bytes_total`
  * `fortigate_managed_switch_tx_drops_total`
  * `fortigate_managed_switch_tx_errors_total`
  * `fortigate_managed_switch_tx_mcast_packets_total`
  * `fortigate_managed_switch_tx_oversize_total`
  * `fortigate_managed_switch_tx_packets_total`
  * `fortigate_managed_switch_tx_ucast_packets_total`
  * `fortigate_managed_switch_under_size_total`
