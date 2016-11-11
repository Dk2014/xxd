<?php

db_execute($db, "
create table player_push_notify_switch(
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`notification_id` int(11) NOT NULL COMMENT '推送通知主键',
	PRIMARY KEY (`id`),
KEY `idx_pid` (`pid`)
 ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家推送通知开关列表';

 ");
?>
