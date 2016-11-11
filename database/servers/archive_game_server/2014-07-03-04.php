<?php
db_execute($db, "

DROP TABLE IF EXISTS `vip_level`;
create table `vip_level` (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`level` smallint(6)  NOT NULL COMMENT 'VIP等级',
	`ingot` bigint(20) NOT NULL  COMMENT '累计充值元宝要求',
	UNIQUE KEY `level` (`level`),
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='VIP等级数值';


DROP TABLE IF EXISTS `player_vip`;
CREATE TABLE `player_vip` (
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`ingot` bigint(20) NOT NULL DEFAULT '0' COMMENT '累计充值元宝数',
	`level` smallint(6)  NOT NULL DEFAULT 0 COMMENT 'VIP等级',
	`card_id` varchar(50) NOT NULL COMMENT 'VIP卡编号',
	PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家VIP卡信息';

DROP TABLE IF EXISTS `vip_privilege`;
create table `vip_privilege` (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`name` varchar(20) NOT NULL COMMENT '特权名称',
	`sign` varchar(20) NOT NULL COMMENT '唯一标识',
	`tip` varchar(200) NOT NULL COMMENT '特权描述',
	`times` smallint(6) NOT NULL DEFAULT '0' COMMENT '特权次数',
	`require_vip_level` smallint(6) NOT NULL DEFAULT '0' COMMENT '要求VIP等级',
	UNIQUE Key `sign` (`sign`),
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家VIP特权表';

");
?>

