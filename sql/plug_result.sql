CREATE TABLE `plug_result` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
	`run_task_id` VARCHAR(255) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`task_id` CHAR(18) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`plug` VARCHAR(50) NOT NULL DEFAULT '0' COLLATE 'utf8_general_ci',
	`res` LONGTEXT NOT NULL COLLATE 'utf8_general_ci',
	`create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `port_scan_result_port_tr_create_time_7a565a50` (`create_time`) USING BTREE,
	INDEX `port_scan_result_main_task_id_61908dd9_fk_port_scan` (`run_task_id`) USING BTREE,
	INDEX `port_scan_result_port_tr_ip_33f818d0` (`plug`) USING BTREE,
	INDEX `task_id` (`task_id`) USING BTREE
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB
AUTO_INCREMENT=1
;
