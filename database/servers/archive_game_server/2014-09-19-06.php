<?php
db_execute($db, "
CREATE TABLE `fashion` (
	`id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '时装ID',
	`name` varchar(20) NOT NULL COMMENT '时装名称',
	`level` int(11) DEFAULT NULL COMMENT '需求等级',
	`desc` varchar(100) DEFAULT NULL COMMENT '时装描述',
	`source` varchar(30) DEFAULT NULL COMMENT '时装来源',
	`sign` varchar(30) DEFAULT NULL COMMENT '资源标识',
	`valid_hours` int(11) NOT NULL DEFAULT '0' COMMENT '有效小时数 0--永久 其他--有效时间',
	`health` int(11) DEFAULT '0' COMMENT '生命',
	`speed` int(11) DEFAULT '0' COMMENT '速度',
	`cultivation` int(11) DEFAULT '0' COMMENT '内力',
	`attack` int(11) DEFAULT '0' COMMENT '攻击',
	`defence` int(11) DEFAULT '0' COMMENT '防御',
	`fashion_item` smallint(6) NOT NULL COMMENT '关联物品ID',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='时装';

CREATE TABLE `player_fashion_state` (
	  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
	  `update_time` bigint(20) NOT NULL COMMENT '状态更新时间',
	  `dressed_fashion_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '当前装备的时装',
	  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家时装状态';

CREATE TABLE `player_fashion` (
	  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
	  `fashion_id` smallint(6) NOT NULL COMMENT '时装模版ID',
	  `expire_time` bigint(20) NOT NULL COMMENT '过期时间 0--永远有效',
	  PRIMARY KEY (`id`),
KEY `idx_pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家时装';

");
?>
