<?php
db_execute($db, "
create table `player_mail_state`(
	`pid` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '附件邮件ID',
	`latest_global_mail_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '玩家收到的最大全局邮件ID',
	PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家全局邮件记录';

create table `global_mail_pending_queue`(
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`send_time` bigint(20) NOT NULL COMMENT '发送时间戳',
	`expire_time` bigint(20) NOT NULL COMMENT '创建时间戳',
	`title` varchar(30) NOT NULL COMMENT '标题',
	`content` varchar(1024) NOT NULL COMMENT '内容',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='全局邮件待发送队列';

alter table `global_mail` add column `attach_ref` bigint(20) NOT NULL COMMENT '附件查询索引对应global_mail_attachments.attach_ref';

alter table `global_mail_attachments` change column `mail_id` `attach_ref` bigint(20) NOT NULL COMMENT '全局邮件表主键';



");

?>
