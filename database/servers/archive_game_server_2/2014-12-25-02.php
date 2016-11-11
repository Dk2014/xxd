<?php
$this->AddSQL("
create table `recruit_buddy` (
	`id` int(11) NOT NULL AUTO_INCREMENT  COMMENT 'ID',
	`role_id` tinyint(4) NOT NULL COMMENT '角色ID',
	`init_level` smallint(6) NOT NULL COMMENT '初始等级',
	`description` varchar(1024) DEFAULT '' NOT NULL COMMENT '描述',
  	`favourite_item` smallint(6) NOT NULL COMMENT '喜好品ID',
 	`favourite_count` smallint(6) NOT NULL COMMENT '喜好品需求量',
	`quest_id`  smallint(6) NOT NULL COMMENT '开启任务',
	`related_npc` int(11) NOT NULL COMMENT '关联NPC',
UNIQUE KEY (`role_id`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='招募伙伴';

create table `extend_quest` (
	`id` int(11) NOT NULL AUTO_INCREMENT  COMMENT 'ID',
	`description` varchar(1024) DEFAULT '' NOT NULL COMMENT '任务描述',
	`type` tinyint(4) NOT NULL COMMENT '任务类型 1--通关区域评星 2--连续登录 3--元宝购买',
	`required_quest`  smallint(6) NOT NULL COMMENT '前置主线任务',
	`related_npc` int(11) NOT NULL COMMENT '关联NPC',
	`related_mission` smallint(6) NOT NULL DEFAULT '0' COMMENT '关联主线区域',
	`required_progress` smallint(6) NOT NULL DEFAULT '0' COMMENT '要求进度',
	`award_item_1` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品1',
	`award_num_1` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励数量1',
	`award_item_2` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品2',
	`award_num_2` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励数量2',
	`award_item_3` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品3',
	`award_num_3` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励数量3',
	`award_exp`  int(11) NOT NULL DEFAULT '0' COMMENT '奖励经验',
	`award_coins`  int(11) NOT NULL DEFAULT '0' COMMENT '奖励铜钱',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='招募伙伴';

create table `player_extend_quest` (
  	`id` bigint(20) NOT NULL  COMMENT 'ID',
  	`pid` bigint(20) NOT NULL COMMENT '用户ID',
	`quest_id` int(11) NOT NULL COMMENT '任务ID',
	`progress` smallint(6) NOT NULL DEFAULT '0' COMMENT '任务进度',
	`state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0--未完成 1--已完成 2--已奖励',
KEY `idx_pid` (`pid`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='伙伴任务';


");
?>
