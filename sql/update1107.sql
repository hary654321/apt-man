ALTER TABLE `task` ADD COLUMN `group` VARCHAR(30) NOT NULL DEFAULT '' COMMENT '分组' AFTER `name`;
ALTER TABLE `os` ADD COLUMN `port` TEXT NULL AFTER `os`,ADD COLUMN `create_time` DATETIME NULL AFTER `port`;
ALTER TABLE `probe_info` ADD COLUMN `probe_port` TEXT NULL AFTER `probe_desc`;
UPDATE `probe_info` SET `probe_port`='80,443,8443,50050,8080,8443,8888,81' ;
UPDATE `probe_info` SET `probe_port`='7000' WHERE  `probe_name` like '%frp%';
UPDATE `probe_info` SET `probe_port`='6722' WHERE  `probe_name` like '%fus%';
UPDATE `probe_info` SET `probe_port`='33456' WHERE  `probe_name` like '%gortcp%';
UPDATE `probe_info` SET `probe_port`='8081' WHERE  `probe_name` like '%gost%';
