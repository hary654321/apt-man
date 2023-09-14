
INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'Admin', '/api/v1/plug*', '(GET)|(POST)|(DELETE)|(PUT)', '', '', '');

ALTER TABLE `task` ADD COLUMN `plug` VARCHAR(1000) NOT NULL DEFAULT '' AFTER `probeId`;

ALTER TABLE `probe_result` CHANGE COLUMN `port` `port` INT(10) NOT NULL DEFAULT 0 COLLATE 'utf8_general_ci' AFTER `run_task_id`;


ALTER TABLE `plug_result`	CHANGE COLUMN `type` `plug` VARCHAR(50) NOT NULL DEFAULT '0' AFTER `run_task_id`;

DROP TABLE  `plug`;

CREATE TABLE IF NOT EXISTS `plug` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `filename` varchar(255) NOT NULL DEFAULT '',
  `cmd` varchar(255) NOT NULL DEFAULT '',
  `desc` varchar(255) NOT NULL DEFAULT '',
  `sys` tinyint(1) NOT NULL DEFAULT '0',
  `status` tinyint(1) NOT NULL DEFAULT '0',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


INSERT INTO `plug` (`id`, `name`, `filename`, `cmd`, `desc`, `sys`, `status`,`is_deleted`, `create_time`, `update_time`) VALUES
	(1, 'NMAP', 'nmap', 'nmap -T5   -p {port}  -oX  {res}  {ip}', '内置插件NMAP', 1,3, 0, '2023-09-11 09:08:40', '2023-09-12 07:13:08');

DROP TABLE `probe_group`;

CREATE TABLE IF NOT EXISTS `probe_group` (
  `probe_group_id` int(11) NOT NULL AUTO_INCREMENT,
  `probe_group_name` varchar(255) NOT NULL,
  `probe_group_type` varchar(255) NOT NULL,
  `probe_group_region` varchar(255) NOT NULL,
  `probe_group_desc` longtext NOT NULL,
  `probe_group_create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `probe_group_update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`probe_group_id`),
  UNIQUE KEY `probe_group_name` (`probe_group_name`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;



INSERT INTO `probe_group` ( `probe_group_name`, `probe_group_type`, `probe_group_region`, `probe_group_desc`, `probe_group_create_time`, `probe_group_update_time`, `is_deleted`) VALUES
	('Confucius', 'APT组织', 'IN', '摩罗桫', '2022-06-21 17:52:16', '2023-09-08 03:51:49', 0),
	('Kimsuky', 'APT组织', 'KP', 'Mystery Baby, Baby Coin, Smoke Screen, BabyShark, Cobra Venom', '2022-06-21 17:52:16', '2023-09-08 03:51:23', 0),
	('CIA', 'APT组织', 'US', 'CIA', '2022-06-21 17:52:16', '2023-09-08 03:51:49', 0),
	('APT32', 'APT组织', 'VN', '海莲花（OceanLotus）', '2022-06-21 17:52:16', '2023-09-08 03:51:49', 0),
	('TransparentTribe', 'APT组织', 'PK', 'ProjectM、C-Major', '2022-06-21 17:52:16', '2023-09-08 03:51:49', 0),
	('钓鱼雀', 'APT组织', 'KR', '', '2023-08-11 15:03:44', '2023-09-08 03:51:49', 0),
	('DarkHotel', 'APT组织', 'KR', '', '2022-06-21 17:52:16', '2023-09-08 03:51:49', 0),
	('摩诃草', 'APT组织', 'IN', 'APT-C-09、白象', '2022-06-21 17:52:16', '2023-09-08 03:51:49', 0),
	('HAFNIUM', 'APT组织', 'CN', '', '2022-06-21 17:52:16', '2023-09-08 03:51:49', 0),
	( 'BITTER', 'APT组织', 'IN', '蔓灵花', '2022-06-21 17:52:16', '2023-09-08 03:51:49', 0),
	( 'StrongPity', 'APT组织', 'TR', '', '2022-06-21 17:52:16', '2023-09-08 03:51:49', 0),
	( 'Turla', 'APT组织', 'RU', 'Venomous Bear、Waterbug和Uroboros', '2022-06-21 17:52:16', '2023-09-08 03:51:49', 0),
	( 'Wellmess', 'APT组织', 'RU', 'APT-C-42', '2022-06-21 17:52:16', '2023-09-08 03:51:49', 0),
	( 'APT-C-01', 'APT组织', 'TW', '毒云藤、绿斑', '2022-06-21 17:52:16', '2023-09-08 03:51:49', 0),
	( 'Donot', 'APT组织', 'IN', 'APT-C-35、肚脑虫', '2022-06-21 17:52:16', '2023-09-08 03:51:49', 0),
	( 'PROMETHIUM', 'APT组织', 'TR', 'APT-C-41、蓝色魔眼、StrongPity', '2022-06-21 17:52:16', '2023-09-08 03:51:49', 0),
	( '响尾蛇', 'APT组织', 'IN', '', '2022-06-21 17:52:16', '2023-09-08 03:51:49', 0),
	( 'APT-C-53', 'APT组织', 'RU', 'Gamaredon', '2022-06-21 17:52:16', '2023-09-08 03:51:49', 0),
	( 'MuddyWater', 'APT组织', 'IR', '', '2022-06-21 17:52:16', '2023-09-08 03:51:49', 0),
	( 'APT37', 'APT组织', 'KP', 'Group123', '2022-06-21 17:52:16', '2023-09-08 03:51:49', 0),
	( 'APT-C-12', 'APT组织', 'TW', '蓝宝菇', '2022-06-21 17:52:16', '2023-09-08 03:51:49', 0),
	( 'Lazarus Group', 'APT组织', 'KP', 'T-APT-15', '2022-06-21 17:52:16', '2023-09-08 03:51:49', 0),
	( '载荷工具', 'APT组织', '-', '', '2022-06-21 17:52:16', '2023-09-08 03:51:49', 0),
	( 'APT-C-23', 'APT组织', '中东', '双尾蝎', '2022-06-21 17:52:16', '2023-09-08 03:51:49', 0),
	( 'APT28', 'APT组织', 'RU', 'Pawn Storm, Sofacy Group, Sednit或STRONTIUM，奇幻熊', '2022-06-21 17:52:16', '2023-09-08 03:51:49', 0),
	( 'APT29', 'APT组织', 'RU', '舒适熊', '2022-06-21 17:52:16', '2023-09-08 03:51:49', 0),
	( '内网穿透', '后门', '', '', '2022-06-21 17:52:16', '2023-09-08 03:52:51', 0);


