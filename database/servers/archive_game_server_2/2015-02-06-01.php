<?php

$this->AddSQL("

CREATE TABLE `shaded_mission` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '影之间隙ID',
  `mission_level_id` int(11) NOT NULL COMMENT '入口区域关卡ID',
  `name` varchar(10) NOT NULL COMMENT '影之间隙名称',  

  `enter_y` int(11) NOT NULL COMMENT '进入点y坐标',
  `enter_x` int(11) NOT NULL COMMENT '进入点x坐标',
  `exit_y` int(11) NOT NULL COMMENT '退出传送点y坐标',
  `exit_x` int(11) NOT NULL COMMENT '退出传送点x坐标',
  
  `mission_link_y` int(11) NOT NULL COMMENT '区域关卡入口y',
  `mission_link_x` int(11) NOT NULL COMMENT '区域关卡入口x',
  `mission_back_y` int(11) NOT NULL COMMENT '区域关卡返回y',
  `mission_back_x` int(11) NOT NULL COMMENT '区域关卡返回x',

  `sign` varchar(50) NOT NULL COMMENT '资源标识',
  `flip_horizontal` tinyint(4) NOT NULL COMMENT '水平翻转',
  `sign_war` varchar(50) NOT NULL COMMENT '关卡战斗资源标识',
  `music` varchar(20) NOT NULL COMMENT '音乐资源标识',

  `box_x` int(11) NOT NULL COMMENT '宝箱x坐标',
  `box_y` int(11) NOT NULL COMMENT '宝箱y坐标',
  `box_dir` tinyint(4) NOT NULL COMMENT '宝箱朝向(0--左;1--右)',

  `award_exp` int(11) NOT NULL COMMENT '奖励经验',
  `award_coin` int(11) NOT NULL COMMENT '奖励铜钱',
  `award_item1` smallint(6) NOT NULL DEFAULT '0' COMMENT '关卡奖励物品1',
  `award_item1_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '关卡奖励物品1数量',
  `award_item2` smallint(6) NOT NULL DEFAULT '0' COMMENT '关卡奖励物品2',
  `award_item2_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '关卡奖励物品2数量',
  `award_item3` smallint(6) NOT NULL DEFAULT '0' COMMENT '关卡奖励物品3',
  `award_item3_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '关卡奖励物品3数量',
  PRIMARY KEY (`id`),
  KEY `idx_mission_level_id` (`mission_level_id`)
) ENGINE=InnoDB AUTO_INCREMENT=450 DEFAULT CHARSET=utf8mb4 COMMENT='影之间隙配置';

ALTER TABLE `mission_enemy` ADD `shaded_mission_id` int(11) NOT NULL DEFAULT '0' COMMENT '所属影之间隙ID，0为关卡怪';

CREATE TABLE `player_shaded_mission_record` (
  `id` bigint(20) NOT NULL COMMENT '记录ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `waiting_shadow_id` int(11) NOT NULL COMMENT '等待清剿的影之间隙ID',
  `mission_level_id` int(11) NOT NULL COMMENT '所属关卡ID',
  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`),
  KEY `idx_mission_level_id` (`mission_level_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='影之间隙记录';

");

