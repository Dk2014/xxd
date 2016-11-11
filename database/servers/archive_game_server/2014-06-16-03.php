<?php

db_execute($db, "

CREATE TABLE `battle_pet` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `quality` tinyint(4) NOT NULL COMMENT '品质',
  `pet_id` int(10) unsigned NOT NULL COMMENT '灵宠ID(enemy_role)',
  `desc` varchar(100) DEFAULT '' COMMENT '灵宠描述',
  `round_attack` tinyint(4) DEFAULT '1' COMMENT '单回合行动次数',
  `cost_power` tinyint(4) NOT NULL COMMENT '召唤时消耗精气',
  `live_round` tinyint(4) NOT NULL COMMENT '召唤后存活回合数',
  `live_pos` tinyint(4) NOT NULL COMMENT '召唤后出现的位置(1-前排；2-后排；3-左侧)',
  `activate_ball_num` tinyint(4) NOT NULL COMMENT '激活需要的契约球数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='灵宠';

CREATE TABLE `level_battle_pet` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `mission_enemy_id` int(10) unsigned NOT NULL COMMENT '关卡怪物组',
  `battle_pet_id` smallint(6) NOT NULL COMMENT '灵宠ID',
  `rate` tinyint(4) NOT NULL COMMENT '出现概率%',
  `live_round` tinyint(4) NOT NULL COMMENT '出现后存活回合数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='关卡灵宠配置';

CREATE TABLE `player_battle_pet` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `battle_pet_id` smallint(6) NOT NULL COMMENT '灵宠ID',
  `ball_num` tinyint(4) NOT NULL COMMENT '已有的灵宠契约球数量',
  `activated` tinyint(4) NOT NULL DEFAULT '0' COMMENT '灵宠是否已激活',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家灵宠数据';


CREATE TABLE `player_battle_pet_config` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `grid1` bigint(20) DEFAULT '-1' COMMENT '灵宠栏位1(-1-未开启)',
  `grid2` bigint(20) DEFAULT '-1' COMMENT '灵宠栏位2(-1-未开启)',
  `grid3` bigint(20) DEFAULT '-1' COMMENT '灵宠栏位3(-1-未开启)',
  `grid4` bigint(20) DEFAULT '-1' COMMENT '灵宠栏位4(-1-未开启)',
  `grid5` bigint(20) DEFAULT '-1' COMMENT '灵宠栏位5(-1-未开启)',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家灵宠配置';

");
?>