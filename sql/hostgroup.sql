
CREATE TABLE IF NOT EXISTS `hostgroup` (
  `id` char(18) NOT NULL COMMENT 'ID',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '主机组名称',
  `remark` varchar(100) NOT NULL DEFAULT '' COMMENT '备注',
  `createByID` char(18) NOT NULL DEFAULT '' COMMENT '创建人ID',
  `hostIDs` text COMMENT '主机ID',
  `createTime` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updateTime` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

