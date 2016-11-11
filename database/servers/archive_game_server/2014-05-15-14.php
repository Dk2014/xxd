<?php
db_execute($db, "
	
DROP TABLE IF EXISTS `player_ghost_mission`;
CREATE TABLE `player_ghost_mission` (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `mission_id` tinyint(4) NOT NULL COMMENT '关卡主键id',
  `ghost_unique_key` smallint(6) NOT NULL DEFAULT '0' COMMENT '获得魂侍的信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家魂侍副本表';

DROP TABLE IF EXISTS `player_ghost_state`;
CREATE TABLE `player_ghost_state` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `gold_fail_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '金魂失败次数',
  `gold_probability` smallint(6) NOT NULL DEFAULT '0' COMMENT '金魂下次概率',
  `purple_fail_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '紫魂失败次数',
  `purple_probability` smallint(6) NOT NULL DEFAULT '0' COMMENT '紫魂下次概率',
  `purify_day_count` bigint(20) DEFAULT '0' COMMENT '每日净化次数',
  `ingot_egg_buy_day_count` bigint(20) DEFAULT '0' COMMENT '每日购买金蛋次数',
  `ingot_egg_buy_update_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '元宝购买魂蛋的时间',
  `ghost_unique_key` smallint(6) NOT NULL DEFAULT '0' COMMENT '获得金魂的信息',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家魂侍状态表';

");
?>