BEGIN;

CREATE TABLE IF NOT EXISTS `host` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主机ID',
  `hostname` varchar(100) NOT NULL COMMENT '主机名',
  `ip` varchar(20) NOT NULL COMMENT '主机ip',
  `sshPort` int(11) NOT NULL DEFAULT '0' COMMENT 'ssh端口',
  `servicePort` int(11) NOT NULL DEFAULT '0' COMMENT '服务端口',
  `sshUser` varchar(100) NOT NULL COMMENT 'ssh用户',
  `sshPwd` varchar(100) NOT NULL COMMENT 'ssh密码',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT 'ssh 状态',
  `runningTasks` varchar(2000) DEFAULT '' COMMENT '运行的任务',
  `weight` int(11) NOT NULL DEFAULT '100' COMMENT '权重',
  `stop` int(11) NOT NULL DEFAULT '0' COMMENT '主机暂停执行任务',
  `version` varchar(20) NOT NULL COMMENT '版本号',
  `lastUpdateTimeUnix` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `remark` varchar(100) DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `ip` (`ip`),
  UNIQUE KEY `hostname` (`hostname`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

INSERT INTO `host` (`id`, `hostname`, `ip`, `sshPort`, `servicePort`, `sshUser`, `sshPwd`, `status`, `runningTasks`, `weight`, `stop`, `version`, `lastUpdateTimeUnix`, `remark`) VALUES (1, '本机', '127.0.0.1', 22, 61666, 'root', '123456', 2, '', 50, 0, '1.0.0', 0, '');

COMMIT;


