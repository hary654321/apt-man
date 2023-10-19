CREATE TABLE `probe_result` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
	`ip` VARCHAR(40) NOT NULL COLLATE 'utf8_general_ci',
	`port` INT(10) NOT NULL DEFAULT '0',
	`probe_name` VARCHAR(255) NOT NULL COLLATE 'utf8_general_ci',
	`cert` TEXT NOT NULL COLLATE 'utf8_general_ci',
	`response` LONGTEXT NOT NULL COLLATE 'utf8_general_ci',
	`create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`is_deleted` TINYINT(1) NOT NULL DEFAULT '0',
	`matched` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '0未开始匹配1匹配上2未匹配上',
	`dealed` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '1 未处理  2已处理',
	`remark` VARCHAR(255) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`run_task_id` VARCHAR(255) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `probe_scan_result_main_task_id_2f750a41_fk_probe_sca` (`run_task_id`) USING BTREE
)
ENGINE=InnoDB
AUTO_INCREMENT=1
DEFAULT CHARSET=utf8;

