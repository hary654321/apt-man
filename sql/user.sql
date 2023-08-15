
CREATE TABLE IF NOT EXISTS `user` (
  `id` char(18) NOT NULL COMMENT '用户ID',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '用户名',
  `hashpassword` varchar(100) NOT NULL DEFAULT '' COMMENT '加密密码',
  `role` int(1) NOT NULL DEFAULT '0' COMMENT '用户类型 1:普通用户 2:管理员 3:访客',
  `forbid` int(1) NOT NULL DEFAULT '0' COMMENT '用户是否可以登陆 ',
  `remark` varchar(100) NOT NULL DEFAULT '' COMMENT '备注',
  `email` varchar(30) NOT NULL DEFAULT '' COMMENT '邮箱',
  `dingphone` char(11) NOT NULL DEFAULT '' COMMENT '钉钉绑定的手机号',
  `telegram` varchar(20) NOT NULL DEFAULT '' COMMENT 'TelegramBot ID',
  `wechat` varchar(20) NOT NULL DEFAULT '' COMMENT '企业微信ID',
  `createTime` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updateTime` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
