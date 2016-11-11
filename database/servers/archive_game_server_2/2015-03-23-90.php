<?php 
$this->AddSQL("
CREATE TABLE `world_channel_message` (
	`id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '消息模版ID',
	`desc` varchar(30) NOT NULL COMMENT '描述',
	`sign` varchar(30) DEFAULT NULL COMMENT '唯一标识',
	`parameters` varchar(1024) NOT NULL COMMENT '参数',
	`content` varchar(1024) NOT NULL COMMENT '内容',
	PRIMARY KEY (`id`),
UNIQUE KEY `sign` (`sign`)
  ) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='世界频道消息模版模板';
");
?>
