<?php
db_execute($db, "
alter table trader_grid_config modify column `goods_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '货物类型0-物品 1-剑心 2-魂侍 3-灵宠契约球 4-爱心 5-铜币 6-元宝 7-体力 8-等级装备';


create table trader_grid_config_equiement(
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`config_id` int(11) NOT NULL COMMENT '配置ID',
	`cost` bigint(20) NOT NULL COMMENT '价格',
	`item_id` smallint(6) NOT NULL COMMENT '物品ID',
	`min_level` smallint(6) NOT NULL COMMENT '等级下限',
	`max_level` smallint(6) NOT NULL COMMENT '等级上限',
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='随机商店格子配置等级装备配置';
");
?>
