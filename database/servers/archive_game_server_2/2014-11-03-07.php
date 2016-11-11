<?php

$this->AddSQL("
CREATE TABLE `pve_level`(
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`floor` smallint(6) NOT NULL  COMMENT '关卡层数',
	`award_item` smallint(6) NOT NULL DEFAULT '0' COMMENT '首次通关奖励物品ID',
	`award_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '首次通关奖励物品数量',
	`moster_num` smallint(6) NOT NULL COMMENT '怪物数量',
	PRIMARY KEY(`id`),
UNIQUE KEY(`floor`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='灵宠幻境';

CREATE TABLE `player_pve_state` (
	  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
	  `max_passed_floor` smallint(6) NOT NULL DEFAULT '0' COMMENT '已通关最大层数',
	  `max_awarded_floor` smallint(6) NOT NULL DEFAULT '0' COMMENT '已奖励最大层数',
	  `unpassed_floor_enemy_num`  smallint(6) NOT NULL DEFAULT '0' COMMENT '未通关关卡杀敌数',
	  PRIMARY KEY(`pid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家灵宠状态';
");

?>
