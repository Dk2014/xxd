<?php
db_execute($db, "
DROP TABLE IF EXISTS `player_coins`;
CREATE TABLE `player_coins` (
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`buy_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '购买更新时间',
	`daily_count` smallint DEFAULT '0' COMMENT '当天购买次数',
	PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家铜币兑换表';

DROP TABLE IF EXISTS `coins_exchange`;
create table `coins_exchange` (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`unique_key` smallint(6)  not null COMMENT '第几次兑换',
	`ingot` bigint(20) not null  COMMENT '消耗元宝',
	`coins` bigint(20) not null  COMMENT '获得铜币',
	UNIQUE KEY `unique_key` (`unique_key`),
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='铜币兑换收益表';
");
?>
