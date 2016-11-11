<?php

// 数值后台
$this->AddSQL("

CREATE TABLE `driving_sword` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `level` smallint(6) NOT NULL COMMENT '层级',
  `width` tinyint(4) NOT NULL COMMENT '地图宽',
  `height` tinyint(4) NOT NULL COMMENT '地图高',
  `born_x` tinyint(4) unsigned NOT NULL COMMENT '出生地坐标x',
  `born_y` tinyint(4) unsigned NOT NULL COMMENT '出生地坐标y',
  `hole_x` tinyint(4) unsigned NOT NULL COMMENT '传送阵坐标x',
  `hole_y` tinyint(4) unsigned NOT NULL COMMENT '传送阵坐标y',
  `obstacle_count` tinyint(4) NOT NULL COMMENT '障碍总数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='云海御剑';

CREATE TABLE `driving_sword_teleport` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cloud_id` smallint(6) NOT NULL COMMENT '云海id',
  `event_id` tinyint(4) NOT NULL COMMENT '传送点事件id',
  `dest_event_id` tinyint(4) NOT NULL COMMENT '目标传送点事件id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='云海御剑传送点';

CREATE TABLE `driving_sword_exploring` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cloud_id` smallint(6) NOT NULL COMMENT '云海id',
  `event_id` tinyint(4) NOT NULL COMMENT '探险山id',
  `sign` varchar(50) NOT NULL COMMENT '资源标识',
  `award_item1` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励品1',
  `award_item2` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励品2',
  `award_item3` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励品3',
  `award_num1` int(11) NOT NULL DEFAULT '0' COMMENT '奖励品1数量',
  `award_num2` int(11) NOT NULL DEFAULT '0' COMMENT '奖励品2数量',
  `award_num3` int(11) NOT NULL DEFAULT '0' COMMENT '奖励品3数量',
  `award_coin_num` int(11) NOT NULL DEFAULT '0' COMMENT '奖励铜币数量',
  `garrison_item` smallint(6) NOT NULL DEFAULT '0' COMMENT '驻守奖励品',
  `garrison_num` int(11) NOT NULL DEFAULT '0' COMMENT '驻守奖励品数量',
  `garrison_coin_num` int(11) NOT NULL DEFAULT '0' COMMENT '驻守奖励铜币数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='云海御剑探险类';

CREATE TABLE `driving_sword_visiting` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cloud_id` smallint(6) NOT NULL COMMENT '云海id',
  `event_id` tinyint(4) NOT NULL COMMENT '拜访类事件id',
  `sign` varchar(50) NOT NULL COMMENT '资源标识',
  `award_item1` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励品1',
  `award_item2` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励品2',
  `award_item3` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励品3',
  `award_num1` int(11) NOT NULL DEFAULT '0' COMMENT '奖励品1数量',
  `award_num2` int(11) NOT NULL DEFAULT '0' COMMENT '奖励品2数量',
  `award_num3` int(11) NOT NULL DEFAULT '0' COMMENT '奖励品3数量',
  `award_coin_num` int(11) NOT NULL DEFAULT '0' COMMENT '奖励铜币数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='云海御剑拜访类';

CREATE TABLE `driving_sword_treasure` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cloud_id` smallint(6) NOT NULL COMMENT '云海id',
  `event_id` tinyint(4) NOT NULL COMMENT '宝箱类事件id',
  `sign` varchar(50) NOT NULL COMMENT '资源标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='云海御剑宝藏类';

CREATE TABLE `driving_sword_treasure_content` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `treasure_id` int(11) NOT NULL COMMENT '云海宝箱id',
  `order` tinyint(4) NOT NULL COMMENT '奖励顺序',
  `award_item` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励品',
  `award_num` int(11) NOT NULL DEFAULT '0' COMMENT '奖励品数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='云海御剑宝藏类奖励明细';

CREATE TABLE `driving_sword_buy_cost_config` (
  `id` int(12) NOT NULL AUTO_INCREMENT,
  `times` int(12) NOT NULL COMMENT '购买次数',
  `cost` int(12) NOT NULL COMMENT '购买所需元宝',
  PRIMARY KEY (`id`),
  UNIQUE KEY `times` (`times`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='购买云海御剑行动点次数元宝设置';

");

//玩家表
$this->AddSQL("

CREATE TABLE `player_driving_sword_info` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `current_cloud` smallint(6) NOT NULL COMMENT '当前所在云层',
  `highest_cloud` smallint(6) NOT NULL COMMENT '最高开启云层',
  `current_x` tinyint(4) unsigned NOT NULL COMMENT '当前坐标x',
  `current_y` tinyint(4) unsigned NOT NULL COMMENT '当前坐标y',
  `allowed_action` smallint(6) NOT NULL COMMENT '行动点',
  `action_refresh_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '行动点刷新时间',
  `action_buy_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '行动点购买更新时间',
  `daily_action_bought` tinyint(4) DEFAULT '0' COMMENT '行动点当天购买次数',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家御剑基本数据';

CREATE TABLE `player_driving_sword_map` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `cloud_id` smallint(6) NOT NULL COMMENT '云层id',
  `shadows` blob NOT NULL COMMENT '阴影比特阵列',
  `hole_count` tinyint(4) NOT NULL DEFAULT '0' COMMENT '传送阵总数',
  `teleport_count` tinyint(4) NOT NULL DEFAULT '0' COMMENT '传送点总数',
  `obstacle_count` tinyint(4) NOT NULL DEFAULT '0' COMMENT '障碍总数',
  `exploring_count` tinyint(4) NOT NULL DEFAULT '0' COMMENT '探险总数',
  `visiting_count` tinyint(4) NOT NULL DEFAULT '0' COMMENT '拜访总数',
  `treasure_count` tinyint(4) NOT NULL DEFAULT '0' COMMENT '宝藏总数',
  `opened_area_count` smallint(6) NOT NULL DEFAULT '0' COMMENT '开启区域总数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家云海地图';

CREATE TABLE `player_driving_sword_event` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `cloud_id` smallint(6) NOT NULL COMMENT '云层id',
  `x` tinyint(4) NOT NULL COMMENT '事件坐标x',
  `y` tinyint(4) NOT NULL COMMENT '事件坐标y',
  `event_type` tinyint(4) NOT NULL COMMENT '事件类型',
  `data_id` tinyint(4) NOT NULL COMMENT '事件模版数据id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家云海事件列表';

CREATE TABLE `player_driving_sword_event_exploring` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `status` tinyint(4) NOT NULL COMMENT '事件状态',
  `garrison_start` bigint(20) NOT NULL COMMENT '驻守开始时间',
  `garrison_time` bigint(20) NOT NULL COMMENT '已驻守时间',
  `award_time` bigint(20) NOT NULL COMMENT '已领奖时间',
  `role_id` tinyint(4) NOT NULL COMMENT '派驻角色id',
  `cloud_id` smallint(6) NOT NULL COMMENT '云层id',
  `x` tinyint(4) NOT NULL COMMENT '事件坐标x',
  `y` tinyint(4) NOT NULL COMMENT '事件坐标y',
  `data_id` tinyint(4) NOT NULL COMMENT '事件模版数据id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家云海探险事件信息';

CREATE TABLE `player_driving_sword_event_visiting` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '事件状态,1代表已经通过还没领奖，2代表已领奖',
  `target_pid` bigint(20) NOT NULL COMMENT '拜访的玩家',
  `target_side_state` blob COMMENT '拜访玩家战斗状态记录',
  `cloud_id` smallint(6) NOT NULL COMMENT '云层id',
  `x` tinyint(4) NOT NULL COMMENT '事件坐标x',
  `y` tinyint(4) NOT NULL COMMENT '事件坐标y',
  `data_id` tinyint(4) NOT NULL COMMENT '事件模版数据id',
  `target_status` text COMMENT '拜访玩家相关信息记录',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家云海拜访事件信息';

CREATE TABLE `player_driving_sword_event_treasure` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `progress` tinyint(4) NOT NULL COMMENT '开箱进度',
  `cloud_id` smallint(6) NOT NULL COMMENT '云层id',
  `x` tinyint(4) NOT NULL COMMENT '事件坐标x',
  `y` tinyint(4) NOT NULL COMMENT '事件坐标y',
  `data_id` tinyint(4) NOT NULL COMMENT '事件模版数据id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家云海宝藏事件信息';

");

