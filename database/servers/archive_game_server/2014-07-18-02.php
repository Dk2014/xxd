<?php
db_execute($db,"
drop table if exists `trader_reflesh_price`;
CREATE TABLE `trader_reflesh_price` (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`time` tinyint(4) NOT NULL COMMENT '刷新次数',
	`price` int(11) NOT NULL COMMENT '刷新价格',
	PRIMARY KEY (`id`),
UNIQUE KEY `time` (`time`)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='随机商店刷新价格';

drop table if exists `trader`;
CREATE TABLE `trader` (
	`id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`name` varchar(10) NOT NULL COMMENT 'NPC名称',
	`sign` varchar(20) NOT NULL COMMENT '资源标识',
	`talk` varchar(200) NOT NULL COMMENT '对话',
	`sold_out_talk` varchar(200) NOT NULL COMMENT '售罄对话',
	`deal_talk` varchar(200) NOT NULL COMMENT '成交对话',
	PRIMARY KEY (`id`)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='随机商店商人';

drop table if exists `trader_position`;
create table `trader_position` (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`trader_id` smallint(6) NOT NULL  COMMENT '随机商人ID',
	`town_id` smallint(6) NOT NULL COMMENT '城镇ID',
	`x` int(11) NOT NULL COMMENT 'x轴坐标',
	`y` int(11) NOT NULL COMMENT 'y轴坐标',
	`direction` varchar(20) DEFAULT NULL COMMENT '朝向',
	PRIMARY KEY (`id`)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='随机商店商人坐标';

drop table if exists `trader_grid`;
create table `trader_grid` (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`trader_id` smallint(6) NOT NULL  COMMENT '随机商人ID',
	`cost` bigint(20) NOT NULL COMMENT '价格',
	`money_type` smallint(4) NOT NULL COMMENT '货币类型',
	`stock` smallint(4) NOT NULL COMMENT '库存数量',
	PRIMARY KEY (`id`)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='随机商店货物配置';

drop table if exists `trader_grid_config`;
create table `trader_grid_config` (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`grid_id` int(11) NOT NULL  COMMENT '格子ID',
	`item_id` smallint(6) NOT NULL COMMENT '物品ID',
	`num` smallint(6) NOT NULL COMMENT '物品数量',
	`probability` tinyint(4) NOT NULL COMMENT '出现概率（％）',
	PRIMARY KEY (`id`)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='随机商店货物配置';


drop table if exists `player_trader_store_state`;
create table `player_trader_store_state`(
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`last_update_time` bigint(20) NOT NULL COMMENT '最近一次刷新时间',
	`trader_id` smallint(6) NOT NULL  COMMENT '随机商人ID',
	`item_id` smallint(6) NOT NULL COMMENT '物品ID',
	`num` smallint(6) NOT NULL COMMENT '物品数量',
	PRIMARY KEY (`pid`)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='玩家随机商店状态';



"
);
?>
