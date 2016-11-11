<?php
db_execute($db, "

DROP TABLE IF EXISTS `player_mail_attachment`;

CREATE TABLE `player_mail_attachment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '玩家邮件附件ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `player_mail_id` bigint(20) NOT NULL COMMENT 'player_mail 主键ID',
  `item_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型',
  `item_id` smallint(6) NOT NULL COMMENT '物品',
  `item_num` int(11) NOT NULL DEFAULT '0' COMMENT '数量',
  PRIMARY KEY (`id`),
  KEY `idx_pid_mail` (`pid`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='玩家邮件附件表';

");
?>