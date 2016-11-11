<?php

db_execute($db, "

DROP TABLE IF EXISTS `announcement`;
CREATE TABLE `announcement` (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '公告模版ID',
	`sign` varchar(30) DEFAULT NULL COMMENT '唯一标识',
	`type` tinyint(4) NOT NULL COMMENT '0后台公告， 1模块公告, 2活动公告',
	`name` varchar(30) DEFAULT NULL COMMENT '公告名',
	`parameters` varchar(1024) NOT NULL COMMENT '参数',
	`content` varchar(1024) NOT NULL COMMENT '内容',
	`duration` int(11) NOT NULL COMMENT '消息存活时间（秒）',
	`show_cyle` int(11) NOT NULL COMMENT '重复展示时间间隔（秒）',
	PRIMARY KEY (`id`),
UNIQUE KEY `sign` (`sign`)
	    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='公告模版';

DROP TABLE IF EXISTS `global_announcement`;
CREATE TABLE `global_announcement` (
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '公告ID',
	`expire_time` bigint(20) NOT NULL COMMENT '创建时间戳',
	`tpl_id` int(11) NOT NULL COMMENT '公告模版ID',
	`parameters` varchar(1024) NOT NULL COMMENT '模版参数',
	PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='公告列表';

");
?>
