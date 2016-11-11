<?php
db_execute($db, "

drop table if exists `player_send_heart_record`;

create table `player_send_heart_record` (
	`id` bigint(20) NOT NULL COMMENT 'ID',
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`friend_pid` bigint(20) NOT NULL COMMENT '好友ID',
	`send_heart_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '上次送爱心时间',
	PRIMARY KEY (`id`),
KEY `idx_pid` (`pid`)

 ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家好友列表';
");
?>


