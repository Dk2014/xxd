<?php

$this->AddSQL("
CREATE TABLE `player_mail_lg` (
	`id` bigint(20) NOT NULL COMMENT 'ID',
	`pmid` bigint(20) NOT NULL COMMENT '玩家邮件ID',
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`mail_id` int(11) NOT NULL COMMENT '邮件模版ID',
	`state` tinyint(4) NOT NULL COMMENT '0未读，1已读',
	`send_time` bigint(20) NOT NULL COMMENT '发送时间戳',
	`parameters` varchar(1024) NOT NULL COMMENT '模版参数',
	`have_attachment` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否有附件',
	`title` varchar(30) NOT NULL DEFAULT '' COMMENT '标题',
	`content` varchar(1024) NOT NULL DEFAULT '' COMMENT '内容',
	PRIMARY KEY (`id`),
KEY `idx_pid` (`pid`),
KEY `send_time` (`send_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家已删除邮件表';

CREATE TABLE `player_mail_attachment_lg` (
	`id` bigint(20) NOT NULL COMMENT '玩家邮件附件ID',
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`player_mail_id` bigint(20) NOT NULL COMMENT 'player_mail 主键ID',
	`attachment_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '附件类型',
	`item_id` smallint(6) NOT NULL COMMENT '物品',
	`item_num` bigint(20) NOT NULL DEFAULT '0' COMMENT '数量',
	`take_timestamp` bigint(20) NOT NULL DEFAULT '0' COMMENT '附件领取时间',
	PRIMARY KEY (`id`),
KEY `idx_pid_mail` (`pid`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家领取邮件附件表';
");


?>
