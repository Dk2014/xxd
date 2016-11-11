<?php
db_execute($db,"

CREATE TABLE `daily_quest` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `type` tinyint(4) DEFAULT '0' COMMENT '任务类型',
  `name` varchar(10) NOT NULL COMMENT '任务标题',
  `desc` varchar(240) DEFAULT '' COMMENT '简介',
  `require_min_level` int(11) NOT NULL COMMENT '要求玩家最低等级',
  `require_max_level` int(11) NOT NULL COMMENT '要求玩家最高等级',
  `require_open_day` varchar(10) DEFAULT '' COMMENT '开放日',

  `require_count` smallint(6) NOT NULL COMMENT '需要数量',

  `award_exp` int(11) NOT NULL COMMENT '奖励经验',
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
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='每日任务';

CREATE TABLE `player_daily_quest` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '记录ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `quest_id` smallint(6) NOT NULL COMMENT '任务ID',
  `finish_count` smallint(6) NOT NULL DEFAULT '0' COMMENT '完成数量',
  `last_finish_time` bigint(20) NOT NULL COMMENT '最近一次完成的时间',
  `award_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '奖励状态; 0 未奖励； 1已奖励',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家每日任务';

"
);
?>
