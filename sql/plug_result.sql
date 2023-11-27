CREATE TABLE `probe_result` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
	`ip` VARCHAR(40) NOT NULL DEFAULT '' COMMENT 'ip' COLLATE 'utf8_general_ci',
	`run_task_id` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '执行ID' COLLATE 'utf8_general_ci',
	`task_id` CHAR(18) NOT NULL DEFAULT '' COMMENT '任务ID' COLLATE 'utf8_general_ci',
	`port` INT(10) NOT NULL DEFAULT '0' COMMENT '端口',
	`probe_name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '规则名称' COLLATE 'utf8_general_ci',
	`cert` TEXT NOT NULL COMMENT '证书' COLLATE 'utf8_general_ci',
	`is_deleted` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '是否删除',
	`matched` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '0未开始匹配1匹配上2未匹配上',
	`response` LONGTEXT NOT NULL COMMENT '响应' COLLATE 'utf8_general_ci',
	`dealed` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '1 未处理  2已处理',
	`remark` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '备注' COLLATE 'utf8_general_ci',
	`create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `probe_scan_result_main_task_id_2f750a41_fk_probe_sca` (`run_task_id`) USING BTREE,
	INDEX `task_id` (`task_id`) USING BTREE,
	INDEX `ip` (`ip`) USING BTREE,
	INDEX `probe_name` (`probe_name`) USING BTREE
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB
AUTO_INCREMENT=1
;
