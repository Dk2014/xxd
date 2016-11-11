<?php
$this->AddSQL("

create table `addition_quest` (
	`id` int(11) NOT NULL AUTO_INCREMENT  COMMENT 'ID',

	`serial_number` int(11) NOT NULL COMMENT '任务系列',
	`award_lock` smallint(6) NOT NULL DEFAULT 0 COMMENT '奖励支线权值',
	`require_serial_number` int(11) NOT NULL COMMENT '前置任务系列',
	`require_lock` smallint(6) NOT NULL DEFAULT 0 COMMENT '前置支线权值',
	`require_level` smallint(6) NOT NULL DEFAULT 0 COMMENT '要求登记',

	`name` varchar(128)  NOT NULL COMMENT '任务名称',
	`description` varchar(1024)  NOT NULL COMMENT '任务描述',
	`showup_main_quest`  smallint(6) NOT NULL COMMENT '出现主线任务',
	`disappear_main_quest`  smallint(6) NOT NULL COMMENT '消失主线任务',
	`publish_npc` int(11) NOT NULL COMMENT '发布NPC',

	`type` tinyint(4) NOT NULL COMMENT '任务类型 1-NPC对话 2-消灭敌人 3-通关关卡 4-收集物品 5-展示物品 6-区域评星 7-招募伙伴',

	`npc_id` int(11) NOT NULL COMMENT '对话类任务NPC',
	`role_id` tinyint(4) NOT NULL DEFAULT '0' COMMENT '任务伙伴ID',
	`npc_role` int(11) NOT NULL DEFAULT '0' COMMENT '任务伙伴NPC',
	`mission_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '任务区域关卡ID',
	`mission_level_id` int(11) NOT NULL DEFAULT '0' COMMENT '任务区域关卡ID',
	`enemy_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '任务击杀怪物ID（客户端用）',
	`mission_enemy_id` int(11) NOT NULL DEFAULT '0' COMMENT '任务关卡怪物组ID',

	`quest_item_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '任务掉落物品（任务关卡掉落）',
	`quest_item_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '任务掉落数量',
	`quest_item_rate` smallint(6) NOT NULL DEFAULT '0' COMMENT '任务掉落概率',

	`require_item_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '任务需求物品类型 1-装备 2-物品',
	`require_item_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '任务需求物品',

	`required_progress` smallint(6) NOT NULL DEFAULT '0' COMMENT '任务要求进度（物品数量通关次数等）',

	`award_item_1` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品1',
	`award_num_1` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励数量1',
	`award_item_2` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品2',
	`award_num_2` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励数量2',
	`award_equip_1` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励装备1',
	`award_equip_num_1` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励装备1',
	`award_exp`  int(11) NOT NULL DEFAULT '0' COMMENT '奖励经验',
	`award_coins`  int(11) NOT NULL DEFAULT '0' COMMENT '奖励铜钱',

	`conversion_reciving_quest` varchar(1024)  NOT NULL COMMENT '领取任务对话',
	`conversion_recived_quest` varchar(1024)  NOT NULL COMMENT '任务中对话',
	`conversion_finish_quest` varchar(1024)  NOT NULL COMMENT '完成任务对话',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='支线任务';

create table `player_addition_quest` (
  	`id` bigint(20) NOT NULL  COMMENT 'ID',
  	`pid` bigint(20) NOT NULL COMMENT '用户ID',
	`serial_number` int(11) NOT NULL COMMENT '任务系列ID',
	`quest_id` int(11) NOT NULL COMMENT '当前任务ID',
	`lock` smallint(6) NOT NULL DEFAULT '0' COMMENT '任务链权值',
	`progress` smallint(6) NOT NULL DEFAULT '0' COMMENT '任务进度',
	`state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0--未完成 1--已完成 2--已奖励 3--已放弃',
KEY `idx_pid` (`pid`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家支线任务';


");
?>
