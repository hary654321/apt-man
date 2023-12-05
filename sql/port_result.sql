CREATE TABLE `port_result` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '自动ID',
	`run_task_id` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '任务ID' COLLATE 'utf8_general_ci',
	`task_id` CHAR(18) NOT NULL DEFAULT '' COMMENT '执行ID' COLLATE 'utf8_general_ci',
	`ip` VARCHAR(40) NOT NULL DEFAULT '' COMMENT 'ip' COLLATE 'utf8_general_ci',
	`port` VARCHAR(5) NOT NULL COMMENT 'post' COLLATE 'utf8_general_ci',
	`response` TEXT NOT NULL COMMENT '响应' COLLATE 'utf8_general_ci',
	`remark` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '标注' COLLATE 'utf8_general_ci',
	`type` VARCHAR(40) NOT NULL DEFAULT '' COMMENT '类型' COLLATE 'utf8_general_ci',
	`service` VARCHAR(40) NOT NULL DEFAULT '' COMMENT '服务' COLLATE 'utf8_general_ci',
	`version` VARCHAR(40) NOT NULL DEFAULT '' COMMENT '版本' COLLATE 'utf8_general_ci',
	`product_name` VARCHAR(40) NOT NULL DEFAULT '' COMMENT '产品名称' COLLATE 'utf8_general_ci',
	`os` VARCHAR(40) NOT NULL DEFAULT '' COMMENT '操作系统' COLLATE 'utf8_general_ci',
	`create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
	`is_deleted` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '是否删除',
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `port_scan_result_port_tr_ip_33f818d0` (`ip`) USING BTREE,
	INDEX `port_scan_result_port_tr_create_time_7a565a50` (`create_time`) USING BTREE,
	INDEX `port_scan_result_main_task_id_61908dd9_fk_port_scan` (`run_task_id`) USING BTREE,
	INDEX `task_id` (`task_id`) USING BTREE
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB
AUTO_INCREMENT=2011
;
