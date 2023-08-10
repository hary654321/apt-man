CREATE TABLE `match_result` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
	`match_task_id` VARCHAR(255) NOT NULL COLLATE 'utf8_general_ci',
	`run_task_id` VARCHAR(255) NOT NULL COLLATE 'utf8_general_ci',
	`match_desc` VARCHAR(255) NULL DEFAULT NULL COMMENT '备注' COLLATE 'utf8_general_ci',
	`match_source` VARCHAR(255) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`match_ip` VARCHAR(40) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`match_port` VARCHAR(40) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`match_group` VARCHAR(255) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`match_type` VARCHAR(255) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`match_region` VARCHAR(255) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`match_tags` VARCHAR(255) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`match_probe_name` VARCHAR(255) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`match_cert_fingerprint` VARCHAR(256) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`match_cert_subject` VARCHAR(256) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`match_cert_issuer` VARCHAR(256) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`match_cert_dns_names` VARCHAR(256) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`match_cert_valid_from` VARCHAR(128) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`match_cert_valid_to` VARCHAR(128) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`match_cert_base64` LONGTEXT NOT NULL COLLATE 'utf8_general_ci',
	`match_create_time` DATETIME(6) NOT NULL,
	`match_update_time` DATETIME(6) NOT NULL,
	PRIMARY KEY (`id`) USING BTREE
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB
AUTO_INCREMENT=1
;
