
CREATE TABLE IF NOT EXISTS `log` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '任务名称',
  `taskid` char(18) NOT NULL DEFAULT '' COMMENT '任务ID',
  `runTaskId` varchar(50) NOT NULL DEFAULT '',
  `hostid` char(18) NOT NULL DEFAULT '',
  `starttime` bigint(20) NOT NULL DEFAULT '0' COMMENT '开始时间毫秒',
  `endtime` bigint(20) NOT NULL DEFAULT '0' COMMENT '结束时间 毫秒',
  `totalruntime` int(11) NOT NULL DEFAULT '0' COMMENT '总共运行时间',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '执行结束 1:成功 -1:失败',
  `progress` tinyint(4) NOT NULL DEFAULT '0',
  `taskresps` mediumtext COMMENT '任务日志',
  `triggertype` int(11) NOT NULL DEFAULT '0' COMMENT '触发方式',
  `errcode` int(11) NOT NULL DEFAULT '0' COMMENT '错误返回码',
  `errmsg` text COMMENT '错误信息',
  `errtasktype` int(11) NOT NULL DEFAULT '0' COMMENT '出错任务类型',
  `errtaskid` char(18) NOT NULL DEFAULT '' COMMENT '出错任务ID',
  `errtask` char(30) NOT NULL DEFAULT '' COMMENT '出错任务名称',
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`),
  KEY `idx_s_t` (`starttime`,`taskid`),
  KEY `hostid` (`hostid`),
  KEY `runTaskId` (`runTaskId`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

