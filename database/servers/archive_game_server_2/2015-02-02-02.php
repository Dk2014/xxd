<?php

$this->AddSQL("
create table `item_dragon_ball_config` (
	`id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  	`source_item` smallint(6) NOT NULL COMMENT '消耗物品ID',
  	`item_id` smallint(6) NOT NULL COMMENT '目标物品ID',
  	`item_num` smallint(6) NOT NULL COMMENT '目标物品数量',
  	`rate` smallint(6) NOT NULL COMMENT '目标物品概率',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='物品兑换配置表';

alter table `equipment_decompose` 
	add column  `dragon_ball`  smallint(6) NOT NULL DEFAULT '0' COMMENT '获得龙珠',
	add column `dragon_ball_num`  smallint(6) NOT NULL DEFAULT '0' COMMENT '获得龙珠数量';
");

