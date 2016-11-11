<?php

db_execute($db, "

ALTER TABLE `enemy_deploy_form` MODIFY COLUMN `battle_type` tinyint(4) NOT NULL COMMENT '战场类型(0--关卡;1--资源关卡;2--极限关卡;3--多人关卡)';

CREATE TABLE `multi_level` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `sign_war` varchar(50) NOT NULL COMMENT '战斗资源标识',
  `music` varchar(20) NOT NULL DEFAULT '' COMMENT '音乐资源标识',

  `require_level` smallint(6) NOT NULL COMMENT '主角等级要求',
  `daily_num` tinyint(4) NOT NULL COMMENT '允许每天进入次数,0表示不限制',

  `award_exp` int(11) NOT NULL DEFAULT '0' COMMENT '奖励经验',
  `award_coin` int(11) NOT NULL DEFAULT '0' COMMENT '奖励铜钱',

  `award_item1_id` int(11) NOT NULL DEFAULT '0' COMMENT '奖励物品1 id',
  `award_item1_num` int(11) NOT NULL DEFAULT '0' COMMENT '物品1数量',

  `award_item2_id` int(11) NOT NULL DEFAULT '0' COMMENT '奖励物品2 id',
  `award_item2_num` int(11) NOT NULL DEFAULT '0' COMMENT '物品2数量',

  `award_item3_id` int(11) NOT NULL DEFAULT '0' COMMENT '奖励物品3 id',
  `award_item3_num` int(11) NOT NULL DEFAULT '0' COMMENT '物品3数量',

  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='多人关卡';


CREATE TABLE `player_multi_level_info` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `buddy_role_id` tinyint(4) NOT NULL COMMENT '上阵伙伴角色模板ID',
  `buddy_row` tinyint(4) NOT NULL COMMENT '上阵伙伴所在行（1或2)',
  `tactical_grid` tinyint(4) NOT NULL DEFAULT '0' COMMENT '战术',

  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家多人关卡信息';


");
?>