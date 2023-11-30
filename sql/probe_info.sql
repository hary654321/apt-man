CREATE TABLE `probe_info` (
	`probe_id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
	`probe_name` VARCHAR(255) NOT NULL COMMENT '规则名称' COLLATE 'utf8_general_ci',
	`probe_group` VARCHAR(50) NOT NULL DEFAULT '0' COMMENT '规则分组' COLLATE 'utf8_general_ci',
	`probe_tags` VARCHAR(255) NOT NULL COMMENT '规则标签' COLLATE 'utf8_general_ci',
	`probe_protocol` VARCHAR(255) NOT NULL COMMENT '规则协议' COLLATE 'utf8_general_ci',
	`probe_match_type` VARCHAR(255) NOT NULL COMMENT '匹配类型' COLLATE 'utf8_general_ci',
	`probe_send` LONGTEXT NOT NULL COMMENT '规则荷载' COLLATE 'utf8_general_ci',
	`probe_recv` LONGTEXT NOT NULL COMMENT '结果匹配' COLLATE 'utf8_general_ci',
	`probe_desc` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '规则描述' COLLATE 'utf8_general_ci',
	`probe_port` TEXT NULL DEFAULT NULL COLLATE 'utf8_general_ci',
	`probe_create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	`probe_update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
	`sys` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '是否是系统规则',
	`is_deleted` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '是否删除',
	PRIMARY KEY (`probe_id`) USING BTREE,
	INDEX `probe_name` (`probe_name`) USING BTREE,
	INDEX `probe_group` (`probe_group`) USING BTREE
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB
AUTO_INCREMENT=1
;

