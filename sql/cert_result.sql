CREATE TABLE `cert_result` (
	`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
	`ip` VARCHAR(256) NOT NULL COLLATE 'utf8_general_ci',
	`port` VARCHAR(256) NOT NULL COLLATE 'utf8_general_ci',
	`probe_name` VARCHAR(256) NOT NULL COLLATE 'utf8_general_ci',
	`cert_base64` TEXT NOT NULL COLLATE 'utf8_general_ci',
	`cert_fingerprint` VARCHAR(256) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`cert_issuer` VARCHAR(256) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`cert_issuer_c` VARCHAR(256) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`cert_issuer_cn` VARCHAR(256) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`cert_subject` VARCHAR(256) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`cert_issuer_o` VARCHAR(256) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`cert_serialno` VARCHAR(256) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`cert_subject_c` VARCHAR(256) NOT NULL DEFAULT '' COLLATE 'utf8_general_ci',
	`cert_subject_cn` LONGTEXT NOT NULL COLLATE 'utf8_general_ci',
	`cert_subject_o` LONGTEXT NOT NULL COLLATE 'utf8_general_ci',
	`valid_from` VARCHAR(128) NOT NULL COLLATE 'utf8_general_ci',
	`valid_to` LONGTEXT NOT NULL COLLATE 'utf8_general_ci',
	`is_deleted` TINYINT(1) NOT NULL DEFAULT '0',
	`matched` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '0未开始匹配1匹配上2未匹配上',
	`create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`run_task_id` VARCHAR(255) NULL DEFAULT NULL COLLATE 'utf8_general_ci',
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `ssl_cert_result_main_task_id_b744052e_fk_ssl_cert_` (`run_task_id`) USING BTREE
)
ENGINE=InnoDB
AUTO_INCREMENT=1
DEFAULT CHARSET=utf8;
;
