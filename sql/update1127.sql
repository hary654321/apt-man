ALTER TABLE `probe_info` ADD INDEX `probe_name` (`probe_name`), ADD INDEX `probe_group` (`probe_group`);
ALTER TABLE `probe_result`ADD INDEX `ip` (`ip`),ADD INDEX `probe_name` (`probe_name`);