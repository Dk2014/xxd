<?php
db_execute($db, "

DROP TABLE IF EXISTS `mail`;

CREATE TABLE `mail` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '邮件ID',
  `sign` varchar(30) DEFAULT NULL COMMENT '唯一标识',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型',
  `title` varchar(30) NOT NULL COMMENT '标题',
  `parameters` varchar(1024) NOT NULL COMMENT '参数',
  `content` varchar(1024) NOT NULL COMMENT '内容',
  PRIMARY KEY (`id`),
  UNIQUE KEY `sign` (`sign`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='系统邮件模板';



DROP TABLE IF EXISTS `mail_attachments`;

CREATE TABLE `mail_attachments` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '邮件ID',
  `mail_id` int(11) NOT NULL COMMENT 'mail表主键',
  `item_id` smallint(6) NOT NULL COMMENT '物品',
  `item_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型',
  `item_num` int(11) NOT NULL DEFAULT '0' COMMENT '数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='系统邮件附件';

");
?>