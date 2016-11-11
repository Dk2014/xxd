<?php

db_execute($db, "
drop table if exists `player_push_notify_switch`;

create table `player_push_notify_switch`(
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`options` bigint(20) NOT NULL DEFAULT '0' COMMENT '推送通知开关',
	PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家推送通知开关列表'; 
");
?>
