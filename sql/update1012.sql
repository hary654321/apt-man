UPDATE `apt`.`probe_info` SET `probe_match_type`='cert', `cert_issuer_o:Let\'s Encrypt' WHERE  `probe_id`=9;
INSERT INTO `probe_info` (`probe_id`, `probe_name`, `probe_group`, `probe_tags`, `probe_protocol`, `probe_match_type`, `probe_send`, `probe_recv`, `probe_desc`, `probe_create_time`, `probe_update_time`, `sys`, `is_deleted`) VALUES (80, 'CobaltStrike-默认证书', '载荷工具', '载荷、工具', 'cert', 'keyword', 'GET /manager/text/list/ HTTP/1.1\\r\\nConnection: close\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7\\r\\nAccept-Language: zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6\\r\\nCache-Control: no-cache\\r\\nPragma: no-cache\\r\\nUpgrade-Insecure-Requests: 1\\r\\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edge/115.0.1901.200\\r\\n\\r\\n', 'cert_issuer_o:cobaltstrike', '', '2023-08-11 18:10:38', '2023-10-12 06:14:55', 1, 0);
