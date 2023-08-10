BEGIN;

CREATE TABLE IF NOT EXISTS `probe_info` (
  `probe_id` int(11) NOT NULL AUTO_INCREMENT,
  `probe_name` varchar(255) NOT NULL,
  `probe_group` varchar(50) NOT NULL DEFAULT '0',
  `probe_tags` varchar(255) NOT NULL,
  `probe_protocol` varchar(255) NOT NULL,
  `probe_match_type` varchar(255) NOT NULL,
  `probe_send` longtext NOT NULL,
  `probe_recv` longtext NOT NULL,
  `probe_desc` varchar(255) NOT NULL DEFAULT '',
  `probe_create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `probe_update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`probe_id`),
  UNIQUE KEY `probe_name` (`probe_name`)
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8;


INSERT INTO `probe_info` (`probe_id`, `probe_name`, `probe_group`, `probe_tags`, `probe_protocol`, `probe_match_type`, `probe_send`, `probe_recv`, `probe_desc`, `probe_create_time`, `probe_update_time`, `is_deleted`) VALUES
	(1, 'SideWinder-5path', '响尾蛇', 'C2', 'HTTP', 'keyword', 'GET /1/2/3/4/5 HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 404 Not Found\\r\\nContent-Length: 13\\r\\nConnection: keep-alive\\r\\nContent-Type: text/html; charset=utf-8\\r\\nServer: nginx/1.14.0 (Ubuntu)\\r\\nVary: Accept-Encoding\\r\\n\\r\\n404 Not Found', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(2, 'SideWinder-6path', '响尾蛇', 'C2', 'HTTP', 'keyword', 'GET /1/2/3/4/5/6 HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', '<pre>Cannot GET /1/2/3/4/5/6</pre>', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(3, 'BITTER-imgLog.jpg', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /imgLog.jpg HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(4, 'BITTER-login page', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /LoginPageNew.html HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'index1Auth.php', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(5, 'LazarusAPT-tmp', 'Lazarus Group', 'C2', 'HTTP', 'keyword', 'GET /tmp/ HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(6, 'Bitter-RguhsT-1', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /RguhsT/ HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(7, 'Bitter-ergdfbd-1', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /ergdfbd/ HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(8, 'Bitter-healthne-1', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /healthne/ HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(9, 'Bitter-ourtyaz-1', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /ourtyaz/ HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(10, 'Bitter-ourtyaz-2', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /ourtyaz/dwnack.php HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(11, 'Bitter-ourtyaz-3', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /ourtyaz/qwe.php HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(12, 'Bitter-ourtyaz-4', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /ourtyaz/qwf.php HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(13, 'Bitter-ergdfbd-2', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /ergdfbd/wscspl HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(14, 'Bitter-heathne-3', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /healthne/regdl HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(15, 'Bitter-healthne-2', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /healthne/accept.php HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(16, 'Bitter-RguhsT-2', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /RguhsT/accept.php HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(17, 'Biiter-RguhsT-3', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /RguhsT/ HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 404 Not Found', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(18, 'SideWinder-5path-images', '响尾蛇', 'C2', 'HTTP', 'keyword', 'GET /images/1/2/3/4/5 HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', '<h1>404 Not Found</h1>', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(19, 'Lazarus-pma', 'Lazarus Group', 'C2', 'HTTP', 'keyword', 'GET /phpmyadmin/ HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 403 Forbidden', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(20, 'Bitter-RsdvgiMincSnyYu-1', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /RsdvgiMincSnyYu/ HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK\\r\\n<span>Password', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(21, 'Bitter-PsehestyvuPw-1', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /PsehestyvuPw/ HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK\\r\\n<span>Password', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(22, 'Bitter-tstRsdvgiMincSnyYutsphp-1', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /tstRsdvgiMincSnyYutsphp/ HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK\\r\\n<span>Password', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(23, 'Bitter-tstRsdvgiMincSnyYutspph-1', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /tstRsdvgiMincSnyYutspph/ HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK\\r\\n<span>Password', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(24, 'Bitter-tstRsdvgiMincSnyYutsphp-2', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /tstRsdvgiMincSnyYutsphp/tstPerHyPfilbmiw1.php HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(25, 'Bitter-tstRsdvgiMincSnyYutsphp-3', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /tstRsdvgiMincSnyYutsphp/tstPerHyPfilbmiwts2t.php HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(26, 'Bitter-tstRsdvgiMincSnyYutspph-3', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /tstRsdvgiMincSnyYutspph/tstPerHyPfilbmiwts2t.php HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(27, 'Bitter-tstRsdvgiMincSnyYutspph-2', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /tstRsdvgiMincSnyYutspph/tstPerHyPfilbmiw1.php HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(28, 'Bitter-PsehestyvuPw-2', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /PsehestyvuPw/F1l3estPhPInf1.php HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(29, 'Bitter-PsehestyvuPw-3', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /PsehestyvuPw/F1l3estPhPInf2.php HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(30, 'Bitter-RsdvgiMincSnyYu-2', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /RsdvgiMincSnyYu/PerHyPfilbmiw1.php HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(31, 'Bitter-RsdvgiMincSnyYu-3', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /RsdvgiMincSnyYu/PerHyPfilbmiw2.php HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(32, 'Remy-connect', 'APT32', 'C2', 'TCP', '==', '02', '03', '', '2022-02-07 10:19:00', '2023-08-05 02:03:53', 0),
	(33, 'Remy-heart', 'APT32', 'C2', 'TCP', 're', '0000000000000000', '^30000000', '', '2022-02-07 10:19:00', '2023-08-04 01:14:50', 0),
	(34, 'bittertcp', 'BITTER', 'C2', 'TCP', 'keyword', '9100CA0A62EE49ABCAAECAA7CAA3CAA4CACACA8BCA8ECA87CA83CA84CAE7CA9ACA89CACACA87CAA3CAA9CAB8CAA5CAB9CAA5CAACCABECAEACA9DCAA3CAA4CAAECAA5CABDCAB9CAEACAFDCAEACABACAB8CAA5CAEACACACAFCCAFECAE7CAA8CAA3CABECACACAFBCAF8CAF0CAFACA89CAF0CAF9CAF2CAF0CAFECAFBCAF0CAF2CA80CAF0CAFACAFECACACAF9CAE4CAFACACACA', '', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(35, 'Bitter_autolan.php', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /autolan.php HTTP/1.1\\r\\nCache-Control: max-age=0\\r\\nUser-Agent: Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.26 (KHTML, like Gecko) Chrome/91.0.4718.131 Safari/527.36\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/\\r\\nConnection: close', '200 and blank page', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(36, 'valogin', '响尾蛇', 'C2', 'HTTP', 'keyword', 'GET /v2login/ HTTP/1.1\\r\\nCache-Control: max-age=0\\r\\nSec-Ch-Ua: &quot;Chromium&quot;;v=&quot;92&quot;, &quot; Not A;Brand&quot;;v=&quot;99&quot;, &quot;Google Chrome&quot;;v=&quot;92&quot;\\r\\nSec-Ch-Ua-Mobile: ?0\\r\\nUpgrade-Insecure-Requests: 1\\r\\nUser-Agent: Mozilla/7.0 (Windows NT 6.0; Win32; x32) AppleWebKit/53.36 (KHTML, like Gecko) Chrome/77.0.4495.0 Safari/53.36\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9\\r\\nSec-Fetch-Site: none\\r\\nSec-Fetch-Mode: navigate\\r\\nSec-Fetch-User: ?1\\r\\nSec-Fetch-Dest: document\\r\\nAccept-Encoding: gzip, deflate\\r\\nAccept-Language: en-US,en;q=0.9\\r\\nConnection: close', '200\\r\\nContent-Length: 6802', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(37, 'StrongPity_sy.php', 'PROMETHIUM', 'C2', 'HTTP', 'keyword', 'GET /sy.php HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(38, 'StrongPity_sys.php', 'PROMETHIUM', 'C2', 'HTTP', 'keyword', 'GET /sys.php HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-02-07 10:19:00', '2022-02-07 10:19:00', 0),
	(39, 'termite', '代理', 'termite', 'TCP', 'keyword', '010000000900000000000000ff000000ff0000000000000000000000000000000200000000000000060000000000008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002141fa120380ffff020000000000000000000000000000000000000000000000300000005b000000e0be05edfc7f0000000000000000000000000000000000006e000000770000000000000000000000dfbe05edfc7f00000000000000000000efbe05edfc7f000000000000000000007c0000000000000000000000000000007c0000000380ffff60a78e28307f00009140fa120380ffffd602000000000000b50000000000000060a78e28307f000001000000000000000013400000000000d0c105edfc7f0000000000000000000000000000000000008c875a28307f000000000000000000005001d328307f00001900000000000000967140000000000080bf05edfc7f0000a7c3400000000000-0000000001000000010000001200000054686973206e6f64652069732041646d696e00000000000000000000000000006e000000770000000000000000000000dfbe05edfc7f00000000000000000000efbe05edfc7f000000000000000000007c0000000000000000000000000000007c0000000380ffff60a78e28307f00009140fa120380ffffd602000000000000b50000000000000060a78e28307f000001000000000000000013400000000000d0c105edfc7f0000000000000000000000000000000000008c875a28307f000000000000000000005001d328307f00001900000000000000967140000000000080bf05edfc7f0000a7c3400000000000ffffffff0000000000b8ff9f9552d6e4b0bf05edfc7f00001d6e40000000000050b3af0000000000c0b3af00-01000000', '5468697320436c69656e74204e6f6465', '', '2022-02-07 10:19:00', '2023-08-05 09:08:44', 0),
	(40, 'gortcp', '代理', 'gortcp', 'TCP', '==', '010000000900000000000000ff000000ff0000000000000000000000000000000200000000000000060000000000008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002141fa120380ffff020000000000000000000000000000000000000000000000300000005b000000e0be05edfc7f0000000000000000000000000000000000006e000000770000000000000000000000dfbe05edfc7f00000000000000000000efbe05edfc7f000000000000000000007c0000000000000000000000000000007c0000000380ffff60a78e28307f00009140fa120380ffffd602000000000000b50000000000000060a78e28307f000001000000000000000013400000000000d0c105edfc7f0000000000000000000000000000000000008c875a28307f000000000000000000005001d328307f00001900000000000000967140000000000080bf05edfc7f0000a7c3400000000000-0000000001000000010000001200000054686973206e6f64652069732041646d696e00000000000000000000000000006e000000770000000000000000000000dfbe05edfc7f00000000000000000000efbe05edfc7f000000000000000000007c0000000000000000000000000000007c0000000380ffff60a78e28307f00009140fa120380ffffd602000000000000b50000000000000060a78e28307f000001000000000000000013400000000000d0c105edfc7f0000000000000000000000000000000000008c875a28307f000000000000000000005001d328307f00001900000000000000967140000000000080bf05edfc7f0000a7c3400000000000ffffffff0000000000b8ff9f9552d6e4b0bf05edfc7f00001d6e40000000000050b3af0000000000c0b3af00-01000000', '0700000019534552564552204552524f523a2061757468206572726f720a', '', '2022-02-07 10:19:00', '2023-08-05 07:41:38', 0),
	(41, 'fuso', '代理', 'proxy', 'TCP', 're', '0500000040000000a2-a4000000f09330819f300d06092a864886f70d010101050003818d0030818902818100e42b643814d3b9006fc4fbd6f50c5ace6aaedd2e5ea940ee8d8d1143c9a014d08ad7820c836f7bc355ba96db20f8d4830d52ed8373325e2b398b432e7cac71c4da3613c91a93791c285699fb38f405110ceee5922f2d515fb2af979df6fa324407489d55974338c33f38721d113d5b7dae7843f3b7913c29717ddbbb217db4430203010001', '^0500000040000000a2a4000000f09330819f300d06092a864886f70d010101050003818d0030818902818100', '', '2022-02-07 10:19:00', '2023-08-05 07:41:40', 0),
	(42, 'gost', '代理', 'Tunnel', 'HTTP', 'keyword', 'GET / HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'Proxy-Agent: gost', '', '2022-02-07 10:19:00', '2023-08-05 07:41:41', 0),
	(43, 'frp-tcp', '代理', 'frp', 'TCP', 'keyword', '9a1fdc41ffedd3648a038f23c6e08774-00000000000000010000000b-5096d592ef438333af4cee-000000000000000100000010-94b2af7a914cd002ae67f519288a2fe0000000000000000100000043', '000100020000000100000000', '', '2022-02-07 10:19:00', '2023-08-05 07:57:50', 0),
	(44, 'frp-http', '代理', 'frp', 'HTTP', 'keyword', 'GET / HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'https://github.com/fatedier/frp', '', '2022-02-07 10:19:00', '2023-08-05 07:41:43', 0),
	(45, 'nps-backend', '代理', 'nps', 'HTTP', 'keyword', 'GET /login/index HTTP/1.1\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7\r\nAccept-Encoding: gzip, deflate\r\nAccept-Language: zh-CN,zh;q=0.9,en;q=0.8\r\nCache-Control: no-cache\r\nConnection: keep-alive\r\nCookie: crocodile=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTE3NDgwNDQsImlzcyI6ImNyb2NvZGlsZSIsIlVJRCI6IjYxNTI2MDI4OTgxMzcxMjg5NiIsIlVzZXJOYW1lIjoiencifQ.awkN0Xi69Dj35xtG91zfXCqsCuV_jtf242D0sGhvkkI; sidebarStatus=1; beegosessionID=a5d50b513506f895491f9788a6820347; lang=zh-CN\r\nHost: 192.168.56.132:8080\r\nPragma: no-cache\r\nUpgrade-Insecure-Requests: 1\r\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36', 'window.nps', '', '2022-02-07 10:19:00', '2023-08-05 08:41:20', 0),
	(46, 'nps-http', '代理', 'nps', 'HTTP', 'keyword', 'GET / HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'nps error', '', '2022-02-07 10:19:00', '2023-08-05 08:31:35', 0);


	COMMIT;
	

