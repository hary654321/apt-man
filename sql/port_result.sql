CREATE TABLE `port_result` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
	`run_task_id` VARCHAR(255) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`ip` VARCHAR(40) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`port` VARCHAR(5) NOT NULL COLLATE 'utf8_general_ci',
	`hex` TEXT NOT NULL COLLATE 'utf8_general_ci',
	`response` TEXT NOT NULL COLLATE 'utf8_general_ci',
	`type` VARCHAR(40) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`service` VARCHAR(40) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`version` VARCHAR(40) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`product_name` VARCHAR(40) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`os` VARCHAR(40) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`probe_name` VARCHAR(40) NOT NULL COLLATE 'utf8_general_ci',
	`create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`is_deleted` TINYINT(1) NOT NULL DEFAULT '0',
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `port_scan_result_port_tr_ip_33f818d0` (`ip`) USING BTREE,
	INDEX `port_scan_result_port_tr_create_time_7a565a50` (`create_time`) USING BTREE,
	INDEX `port_scan_result_main_task_id_61908dd9_fk_port_scan` (`run_task_id`) USING BTREE
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB
AUTO_INCREMENT=1
;
