ALTER TABLE `probe_result` CHANGE COLUMN `ip` `ip` CHAR(15) NOT NULL DEFAULT '' COMMENT 'ip' COLLATE 'utf8_general_ci' AFTER `id`;
ALTER TABLE `task` ADD INDEX `group` (`group`);