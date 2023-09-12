
CREATE TABLE IF NOT EXISTS `plug` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `filename` varchar(255) NOT NULL DEFAULT '',
  `cmd` varchar(255) NOT NULL DEFAULT '',
  `desc` varchar(255) NOT NULL DEFAULT '',
  `sys` tinyint(1) NOT NULL DEFAULT '0',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;


INSERT INTO `plug` (`id`, `name`, `filename`, `cmd`, `desc`, `sys`, `is_deleted`, `create_time`, `update_time`) VALUES
	(1, 'NMAP', 'nmap', 'nmap -T5   -p {port}  -oX  res.xml  {ip}', '内置插件NMAP', 1, 0, '2023-09-11 09:08:40', '2023-09-12 07:13:08');