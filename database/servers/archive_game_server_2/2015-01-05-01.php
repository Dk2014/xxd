<?php

$this->AddSQL("

CREATE TABLE `town_treasure` (
	`id` SMALLINT (6) NOT NULL AUTO_INCREMENT COMMENT '主键'
	,`town_id` SMALLINT (6) NOT NULL COMMENT '城镇ID'
	,`award_ingot` INT (11) NOT NULL DEFAULT '0' COMMENT '奖励元宝'
	,`award_exp` INT (11) NOT NULL DEFAULT '0' COMMENT '奖励经验'
	,`award_coins` BIGINT (20) NOT NULL DEFAULT '0' COMMENT '奖励铜钱'
	,`award_physical` TINYINT (4) NOT NULL DEFAULT '0' COMMENT '奖励体力'
	,`award_item1_id` SMALLINT (6) NOT NULL DEFAULT '0' COMMENT '奖励物品1'
	,`award_item1_num` SMALLINT (6) NOT NULL DEFAULT '0' COMMENT '奖励物品1数量'
	,`award_item2_id` SMALLINT (6) NOT NULL DEFAULT '0' COMMENT '奖励物品2'
	,`award_item2_num` SMALLINT (6) NOT NULL DEFAULT '0' COMMENT '奖励物品2数量'
	,`award_item3_id` SMALLINT (6) NOT NULL DEFAULT '0' COMMENT '奖励物品3'
	,`award_item3_num` SMALLINT (6) NOT NULL DEFAULT '0' COMMENT '奖励物品3数量'
	,`award_item4_id` SMALLINT (6) NOT NULL DEFAULT '0' COMMENT '奖励物品4'
	,`award_item4_num` SMALLINT (6) NOT NULL DEFAULT '0' COMMENT '奖励物品4数量'
	,PRIMARY KEY (`id`)
	);

CREATE TABLE `player_opened_town_treasure` (
	`id` BIGINT (20) NOT NULL COMMENT '主键'
	,`pid` BIGINT (20) NOT NULL COMMENT '玩家ID'
	,`town_id` SMALLINT (6) NOT NULL COMMENT '城镇ID'
	,PRIMARY KEY (`id`)
	);

");

