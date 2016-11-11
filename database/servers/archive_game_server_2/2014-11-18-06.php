<?php
$this->AddSQL("
CREATE TABLE `rainbow_buy_cost_config` (
	`id` int(12) NOT NULL AUTO_INCREMENT,
	`times` int(12) NOT NULL COMMENT '购买次数',
	`cost` int(12) NOT NULL COMMENT '购买所需元宝',
	PRIMARY KEY(`id`),
	UNIQUE KEY `times`(`times`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='购买彩虹关卡扫荡次数元宝设置';

alter table `player_rainbow_level` add column  `buy_times` smallint(6) NOT NULL DEFAULT '0' COMMENT '今日购买次数';
");
?>