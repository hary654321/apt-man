
CREATE TABLE IF NOT EXISTS `operate` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `uid` char(18) NOT NULL DEFAULT '' COMMENT '操作用户ID',
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '操作用户名',
  `role` int(11) NOT NULL DEFAULT '0' COMMENT '操作用户类型',
  `method` varchar(6) NOT NULL DEFAULT '' COMMENT '操作类型',
  `module` varchar(10) NOT NULL DEFAULT '' COMMENT '操作模块',
  `modulename` varchar(30) NOT NULL DEFAULT '' COMMENT '操作模块名称 例如任务名称',
  `operatetime` int(11) NOT NULL DEFAULT '0' COMMENT '操作时间',
  `description` varchar(200) DEFAULT NULL COMMENT '操作说明，一般用户用户操作未直接改变数据库变化的操作，例如运行任务',
  `columns` mediumtext COMMENT '修改的字段',
  PRIMARY KEY (`id`),
  KEY `idx_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4;
