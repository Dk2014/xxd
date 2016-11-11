<?php
db_execute($db,"

alter table `trader_grid_config` add column  `cost` bigint(20) NOT NULL COMMENT '价格';
alter table `trader_grid` drop column `cost`;
alter table `trader_grid` modify column `money_type` tinyint(4) NOT NULL COMMENT '货币类型';

drop table if exists `trader_extra_talk`;
create table `trader_extra_talk` (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`trader_id` smallint(6) NOT NULL  COMMENT '随机商人ID',
	`time` tinyint(4) NOT NULL COMMENT '点击次数',
	`talk` varchar(200) NOT NULL COMMENT '对话',
	PRIMARY KEY (`id`)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='随机商店商人额外对话';



drop table if exists `player_trader_store_state`;
create table `player_trader_store_state`(
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`trader_id` smallint(6) NOT NULL  COMMENT '随机商人ID',
	`grid_id` int(11) NOT NULL  COMMENT '格子ID',
	`item_id` smallint(6) NOT NULL COMMENT '物品ID',
	`num` smallint(6) NOT NULL COMMENT '物品数量',
	`cost` bigint(20) NOT NULL  COMMENT '价格',
	`stock` tinyint(4) not null default '0' comment '剩余可购买次数',
	PRIMARY KEY (`id`)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='玩家随机商店状态';
drop table if exists `player_trader_refresh_state`;
create table `player_trader_refresh_state`(
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`last_update_time` bigint(20) NOT NULL COMMENT '最近一次*手动*刷新时间',
	`auto_update_time` bigint(20) NOT NULL COMMENT '最近一次*自动*刷新时间',
	`trader_id` smallint(6) NOT NULL  COMMENT '随机商人ID',
	`refresh_num` smallint(6) not null  comment '已刷新次数',
	PRIMARY KEY (`id`)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='玩家随机商店手动刷新次数状态';


"
);
?>
