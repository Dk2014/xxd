<?php
db_execute($db, "
drop table if exists `player_global_mail`;

create table `player_mail_state`(
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '附件邮件ID',
	`latest_global_mail_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '玩家收到的最大全局邮件ID',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家全局邮件记录';
");

?>
