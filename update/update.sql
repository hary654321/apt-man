ALTER TABLE `probe_result` ADD COLUMN `remark` VARCHAR(255) NOT NULL DEFAULT '' AFTER `response`;
ALTER TABLE `probe_result` ADD COLUMN `dealed` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '1 未处理  2已处理' AFTER `response`;
ALTER TABLE `port_result` ADD COLUMN `remark` VARCHAR(255) NOT NULL DEFAULT '' AFTER `response`;