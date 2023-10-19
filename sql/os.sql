CREATE TABLE `os` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`ip` CHAR(15) NULL DEFAULT NULL COLLATE 'utf8_general_ci',
	`os` VARCHAR(20) NULL DEFAULT NULL COLLATE 'utf8_general_ci',
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE INDEX `ip_os` (`ip`, `os`) USING BTREE
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB
AUTO_INCREMENT=1
;