BEGIN;

CREATE TABLE `task` (
	`id` CHAR(18) NOT NULL COMMENT 'ID' COLLATE 'utf8mb4_general_ci',
	`name` VARCHAR(30) NOT NULL COMMENT '任务名称' COLLATE 'utf8mb4_general_ci',
	`TaskType` TINYINT(4) NOT NULL DEFAULT '0' COMMENT '任务类型 1:Code 2:HTTP',
	`ip` TEXT NULL DEFAULT NULL COMMENT 'ip' COLLATE 'utf8mb4_general_ci',
	`port` TEXT NULL DEFAULT NULL COMMENT 'port' COLLATE 'utf8mb4_general_ci',
	`run` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '是否自动调度运行',
	`status` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '任务的状态',
	`routePolicy` INT(11) NOT NULL DEFAULT '0' COMMENT '路由策略 1:Random 2:RoundRobin 3:Weight 4:LeastTask',
	`timeout` INT(11) NOT NULL DEFAULT '5' COMMENT '任务超时时间，默认5s',
	`threads` INT(11) NOT NULL DEFAULT '5',
	`parentTaskIds` VARCHAR(380) NOT NULL DEFAULT '' COMMENT '父任务ID，最多20个' COLLATE 'utf8mb4_general_ci',
	`parentRunParallel` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '父任务是否并行运行',
	`childRunParallel` TINYINT(1) NOT NULL DEFAULT '0' COMMENT '子任务是否并行运行',
	`childTaskIds` VARCHAR(380) NOT NULL DEFAULT '' COMMENT '子任务ID，最多20个' COLLATE 'utf8mb4_general_ci',
	`createByID` CHAR(18) NOT NULL DEFAULT '' COMMENT '创建人ID' COLLATE 'utf8mb4_general_ci',
	`hostGroupID` CHAR(18) NOT NULL DEFAULT '' COMMENT '主机组ID' COLLATE 'utf8mb4_general_ci',
	`cronExpr` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '定时任务表达式,共7位 秒、分、时、日、月、周、年' COLLATE 'utf8mb4_general_ci',
	`alarmUserIds` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '报警用户 最多设置10个' COLLATE 'utf8mb4_general_ci',
	`alarmStatus` INT(11) NOT NULL DEFAULT '0' COMMENT '报警策略 1:任务运行结束 2:任务运行失败 3:任务运行成功',
	`priority` INT(11) NOT NULL DEFAULT '0' COMMENT '优先级',
	`remark` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '备注' COLLATE 'utf8mb4_general_ci',
	`probeScanId` VARCHAR(100) NOT NULL DEFAULT '' COLLATE 'utf8mb4_general_ci',
	`probeId` VARCHAR(5000) NOT NULL DEFAULT '' COMMENT '探针ID' COLLATE 'utf8mb4_general_ci',
	`createTime` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '任务创建时间 时间戳(秒)',
	`updateTime` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '任务上次修改时间 时间戳(秒)',
	`isDeleted` TINYINT(1) NOT NULL DEFAULT '0',
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `idx_name` (`name`) USING BTREE,
	INDEX `idx_cbi` (`createByID`) USING BTREE
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8;

INSERT INTO `task` (`id`, `name`, `TaskType`, `ip`, `port`, `run`, `status`, `routePolicy`, `timeout`, `threads`, `parentTaskIds`, `parentRunParallel`, `childRunParallel`, `childTaskIds`, `createByID`, `hostGroupID`, `cronExpr`, `alarmUserIds`, `alarmStatus`, `priority`, `remark`, `probeScanId`, `probeId`, `createTime`, `updateTime`, `isDeleted`) VALUES ('test', '127', 1, '127.0.0.1', 'all', 1, 0, 1, 5, 100, '', 0, 0, '', '1', 'zd', '', '', -1, 0, '', '', '', '2023-07-31 06:56:35', '2023-08-18 09:43:34', 0);

COMMIT;