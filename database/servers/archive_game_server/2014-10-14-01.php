<?php

db_execute($db, "
create table `push_notify` (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '公告模版ID',
	`sign` varchar(30) DEFAULT NULL COMMENT '唯一标识',
	`trigger_time` smallint(6) NOT NULL DEFAULT '0' COMMENT '触发时间一天内第几秒 [0,86400)',
	`content` varchar(1024) NOT NULL COMMENT '内容',
	PRIMARY KEY (`id`),
UNIQUE KEY `sign` (`sign`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='推送';

");
?>
