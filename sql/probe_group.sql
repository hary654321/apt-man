BEGIN;


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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;


INSERT INTO `probe_group` (`probe_group_id`, `probe_group_name`, `probe_group_type`, `probe_group_region`, `probe_group_desc`, `probe_group_create_time`, `probe_group_update_time`, `is_deleted`) VALUES
	(1, '响尾蛇', 'APT组织', '东南亚', '', '2023-08-02 08:42:37', '2023-08-02 09:14:10', 0),
	(2, 'BITTER', 'APT组织', '南亚地区', '', '2023-08-02 08:42:37', '2023-08-02 09:16:52', 0),
	(3, 'Lazarus Group', 'APT组织', '朝鲜', '', '2023-08-02 08:42:37', '2023-08-02 09:17:18', 0),
	(4, 'APT32', 'APT组织', '东南亚', '', '2023-08-02 08:42:37', '2023-08-02 09:17:55', 0),
	(5, 'PROMETHIUM', 'APT组织', '中东', '', '2023-08-02 08:42:37', '2023-08-04 09:40:38', 0),
	(6, '内网穿透', '代理', '', '', '2023-08-02 09:47:04', '2023-08-04 09:40:59', 0);


COMMIT;
