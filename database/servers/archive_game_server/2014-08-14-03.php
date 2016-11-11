<?php
db_execute($db, "
create table `player_global_mail`(
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '附件邮件ID',
	`latest_global_mail_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '玩家收到的最大全局邮件ID',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家全局邮件记录';

create table `global_mail`(
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '公告ID',
	`send_time` bigint(20) NOT NULL COMMENT '发送时间戳',
	`expire_time` bigint(20) NOT NULL COMMENT '创建时间戳',
	`title` varchar(30) NOT NULL COMMENT '标题',
	`content` varchar(1024) NOT NULL COMMENT '内容',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='全局邮件';


create table `global_mail_attachments`(
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '附件邮件ID',
	`mail_id` bigint(20) NOT NULL COMMENT '全局邮件表主键',
	`item_id` smallint(6) NOT NULL COMMENT '附件ID',
	`attachment_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '附件类型',
	`item_num` int(11) NOT NULL DEFAULT '0' COMMENT '数量',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='全局邮件附件';

");

?>
