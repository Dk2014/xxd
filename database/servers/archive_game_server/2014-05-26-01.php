<?php
db_execute($db, "

DROP TABLE IF EXISTS `player_mail`;

CREATE TABLE `player_mail` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '玩家邮件ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `mail_id` int(11) NOT NULL COMMENT '邮件模版ID',
  `state` tinyint(4) NOT NULL COMMENT '0未读，1已读',
  `send_time` bigint(20) NOT NULL COMMENT '发送时间戳',
  `parameters` varchar(1024) NOT NULL COMMENT '模版参数',
  `hava_attachment` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否有附件',
  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`),
  KEY `send_time` (`send_time`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='玩家邮件表';

");
?>