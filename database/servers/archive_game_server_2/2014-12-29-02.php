<?php

$this->AddSQL("

CREATE TABLE `battle_pet_skill_training` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `pet_id` int(6) NOT NULL COMMENT '灵宠ID(怪物ID)',
  `level` int(11) NOT NULL COMMENT '技能等级',
  `cost_coins` bigint(20) NOT NULL COMMENT '升出本等级所需金币数',
  `power` int(11) NOT NULL COMMENT '技能威力',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='宠物技能升级表';

CREATE TABLE `battle_pet_level_info` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `pet_id` int(6) NOT NULL COMMENT '灵宠ID(怪物ID)',
  `level` int(11) NOT NULL COMMENT '灵宠等级',
  `health` int(11) NOT NULL COMMENT '生命 - health',
  `speed` int(11) NOT NULL COMMENT '速度 - speed',
  `attack` int(11) NOT NULL COMMENT '普攻 - attack',
  `defence` int(11) NOT NULL COMMENT '普防 - defence',
  `sunder_max_value` int(11) NOT NULL COMMENT '护甲值',
  `sunder_min_hurt_rate` int(11) NOT NULL COMMENT '破甲前起始的伤害转换率（百分比）',
  `sunder_end_hurt_rate` int(11) NOT NULL COMMENT '破甲后的伤害转换率（百分比）',
  `sunder_end_defend_rate` int(11) NOT NULL DEFAULT '0' COMMENT '破甲后减防（百分比）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='宠物升级表';

CREATE TABLE `battle_pet_exp` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `level` smallint(6) NOT NULL COMMENT '灵宠等级',
  `exp` bigint(20) NOT NULL COMMENT '升级所需经验',
  `need_soul_num` int(11) NOT NULL COMMENT '所需灵魄数量',
  `min_add_exp` bigint(20) NOT NULL COMMENT '最小经验加值',
  `max_add_exp` bigint(20) NOT NULL COMMENT '最大经验加值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='灵宠等级表';

ALTER TABLE `player_battle_pet`
  ADD `level` int(11) NOT NULL DEFAULT '1' COMMENT '灵宠等级',
  ADD `exp` bigint(6) NOT NULL COMMENT '当前经验',
  ADD `skill_level` int(11) NOT NULL DEFAULT '1' COMMENT '技能等级';

");

