UPDATE `probe_info` SET `probe_port`='8082' WHERE  `probe_recv`='https://github.com/fatedier/frp';
INSERT INTO `probe_group` (`probe_group_name`, `probe_group_type`, `probe_group_region`, `probe_group_desc`, `probe_group_create_time`, `probe_group_update_time`, `is_deleted`) VALUES ('nc-后门', '后门', '-', '-', '2023-11-07 20:29:11', '2023-11-20 09:37:49', '0');
INSERT INTO `probe_info` (`probe_name`, `probe_group`, `probe_tags`, `probe_protocol`, `probe_match_type`, `probe_send`, `probe_recv`, `probe_port`, `probe_create_time`, `probe_update_time`, `sys`) VALUES ('nc-shell', 'nc-后门', 'shell--bash', 'TCP', 're', '7077640a', '^2f','6666', '2022-02-07 10:19:00', '2023-11-07 10:09:16', '1');
