<?php
db_execute($db, "

alter table `announcement` modify column  `type` tinyint(4) NOT NULL COMMENT '0-走马灯公告';

create table opera_announcement (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '公告模版ID',
	`sign` varchar(30) DEFAULT NULL COMMENT '唯一标识',
	`title` varchar(50) DEFAULT NULL COMMENT '公告名',
	`content` varchar(1024) NOT NULL COMMENT '内容',
	`effect_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '生效时间',
	`expire_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '结束时间 0 永远有效',
	PRIMARY KEY (`id`),
UNIQUE KEY `sign` (`sign`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='活动公告';

");
?>
