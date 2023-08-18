ALTER TABLE `probe_result` ADD COLUMN `remark` VARCHAR(255) NOT NULL DEFAULT '' AFTER `response`;
ALTER TABLE `probe_result` ADD COLUMN `dealed` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '1 未处理  2已处理' AFTER `response`;
ALTER TABLE `port_result` ADD COLUMN `remark` VARCHAR(255) NOT NULL DEFAULT '' AFTER `response`;
ALTER TABLE `plug_result` CHANGE COLUMN `res` `res` LONGTEXT NOT NULL COLLATE 'utf8_general_ci' AFTER `type`;
ALTER TABLE `task` CHANGE COLUMN `status` `status` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '任务的状态' AFTER `run`;