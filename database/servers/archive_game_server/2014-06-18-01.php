<?php

db_execute($db, "

CREATE TABLE `heart_draw` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `draw_type` tinyint(4) NOT NULL  COMMENT '抽奖类型（1-大转盘；2-刮刮卡）',
  `daily_num` tinyint(4) NOT NULL DEFAULT '0' COMMENT '每日可抽奖次数',
  `cost_heart` tinyint(4) NOT NULL  COMMENT '每次抽奖消耗爱心数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='爱心抽奖';

CREATE TABLE `heart_draw_award` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `heart_draw_id` smallint(6) NOT NULL COMMENT '爱心抽奖ID',
  `award_type` tinyint(4) NOT NULL  COMMENT '奖品类型（1-铜钱；2-元宝；3-道具）',
  `award_num` smallint(6) NOT NULL COMMENT '奖品数量',
  `item_id` smallint(6) DEFAULT '0' COMMENT '道具奖品ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='爱心抽奖奖品配置';

CREATE TABLE `player_heart_draw` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `draw_type` tinyint(4) NOT NULL  COMMENT '抽奖类型（1-大转盘；2-刮刮卡）',
  `daily_num` tinyint(4) NOT NULL COMMENT '当日已抽次数',
  `draw_time` bigint(20) NOT NULL COMMENT '最近一次抽奖时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家爱心抽奖';

CREATE TABLE `player_heart_draw_wheel_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `award_type` tinyint(4) NOT NULL  COMMENT '奖品类型（1-铜钱；2-元宝；3-道具）',
  `award_num` smallint(6) NOT NULL COMMENT '奖品数量',
  `item_id` smallint(6) DEFAULT '0' COMMENT '道具奖品ID',
  `draw_time` bigint(20) NOT NULL COMMENT '抽奖时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家爱心大转盘抽奖记录';

CREATE TABLE `player_heart_draw_card_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `award_type` tinyint(4) NOT NULL  COMMENT '奖品类型（1-铜钱；2-元宝；3-道具）',
  `award_num` smallint(6) NOT NULL COMMENT '奖品数量',
  `item_id` smallint(6) DEFAULT '0' COMMENT '道具奖品ID',
  `draw_time` bigint(20) NOT NULL COMMENT '抽奖时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家爱心刮刮卡抽奖记录';

");
?>
