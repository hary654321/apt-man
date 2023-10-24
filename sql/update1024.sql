ALTER TABLE `port_result` ADD COLUMN `task_id` CHAR(18) NOT NULL DEFAULT '' AFTER `run_task_id`;
ALTER TABLE `probe_result` ADD COLUMN `task_id` CHAR(18) NOT NULL DEFAULT '' AFTER `run_task_id`;
ALTER TABLE `plug_result` ADD COLUMN `task_id` CHAR(18) NOT NULL DEFAULT '' AFTER `run_task_id`;
ALTER TABLE `port_result`	ADD INDEX `task_id` (`task_id`);
ALTER TABLE `plug_result`	ADD INDEX `task_id` (`task_id`);
ALTER TABLE `probe_result`	ADD INDEX `task_id` (`task_id`);
