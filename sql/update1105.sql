ALTER TABLE `task` ADD COLUMN `group` VARCHAR(30) NOT NULL DEFAULT '' COMMENT '分组' AFTER `name`;