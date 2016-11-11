<?php

$this->AddSQL("

    CREATE TABLE IF NOT EXISTS `player_global_clique_daily_quest` (
  `id` bigint(20) NOT NULL COMMENT '记录ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `quest_id` smallint(6) NOT NULL COMMENT '任务ID',
  `finish_count` smallint(6) NOT NULL DEFAULT '0' COMMENT '完成数量',
  `last_update_time` bigint(20) NOT NULL COMMENT '最近一次更新时间',
  `award_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '奖励状态; 0 未奖励； 1已奖励',
  `class` smallint(6) NOT NULL COMMENT '每日任务类别',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='玩家每日帮派任务';
");

$this->AddSQL("
CREATE TABLE IF NOT EXISTS `player_global_clique_building_quest` (
  `id` bigint(20) NOT NULL COMMENT '记录ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `quest_id` smallint(6) NOT NULL COMMENT '任务ID',
  `award_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '奖励状态; 0 未奖励； 1已奖励',
  `building_type` smallint(6) NOT NULL COMMENT '建筑类别',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家帮派建筑任务';

");



$this->AddSQL("
CREATE TABLE IF NOT EXISTS  `clique_daily_quest` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `name` varchar(50) DEFAULT NULL,
  `desc` varchar(240) DEFAULT '' COMMENT '简介',
  `require_min_level` int(11) NOT NULL COMMENT '要求玩家最低等级',
  `require_max_level` int(11) NOT NULL COMMENT '要求玩家最高等级',
  `require_open_day` varchar(10) DEFAULT '' COMMENT '开放日',
  `require_count` smallint(6) NOT NULL COMMENT '需要数量',
  `award_exp` int(11) NOT NULL COMMENT '奖励经验',
  `award_clique_contri` bigint(20) NOT NULL DEFAULT '0' COMMENT '奖励帮贡',
  `award_coins` bigint(20) NOT NULL DEFAULT '0' COMMENT '奖励铜钱',
  `award_physical` tinyint(4) NOT NULL DEFAULT '0' COMMENT '奖励体力',
  `award_item1_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品1',
  `award_item1_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品1数量',
  `award_item2_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品2',
  `award_item2_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品2数量',
  `award_item3_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品3',
  `award_item3_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品3数量',
  `award_item4_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品4',
  `award_item4_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品4数量',
  `level_type` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '关卡类型; -1 无; 0-区域关卡;1-资源关卡;2-通天塔;8-难度关卡;9-伙伴关卡;10-灵宠关卡;11-魂侍关卡',
  `level_sub_type` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '关卡子类型(-1--无;1--铜钱关卡;2--经验关卡)',
  `class` smallint(6) NOT NULL COMMENT '任务类别',
  `order` int(11) DEFAULT '0' COMMENT '显示优先级',
  `award_ingot` int(11) NOT NULL DEFAULT '0' COMMENT '奖励元宝',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='帮派每日任务';
");

$this->AddSQL("
CREATE TABLE IF NOT EXISTS  `clique_building_quest` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `name` varchar(60) DEFAULT NULL,
  `desc` varchar(240) DEFAULT '' COMMENT '简介',
  `require_building_level` int(11) NOT NULL COMMENT '要求建筑等级',
  `award_exp` int(11) NOT NULL COMMENT '奖励经验',
  `award_clique_contri` bigint(20) NOT NULL DEFAULT '0' COMMENT '奖励帮贡',
  `award_coins` bigint(20) NOT NULL DEFAULT '0' COMMENT '奖励铜钱',
  `award_physical` tinyint(4) NOT NULL DEFAULT '0' COMMENT '奖励体力',
  `award_item1_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品1',
  `award_item1_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品1数量',
  `award_item2_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品2',
  `award_item2_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品2数量',
  `award_item3_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品3',
  `award_item3_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品3数量',
  `award_item4_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品4',
  `award_item4_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品4数量',
  `class` smallint(6) NOT NULL COMMENT '任务类别',
  `order` int(11) DEFAULT '0' COMMENT '显示优先级',
  `award_ingot` int(11) NOT NULL DEFAULT '0' COMMENT '奖励元宝',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='帮派帮派建筑任务';
");

?>


