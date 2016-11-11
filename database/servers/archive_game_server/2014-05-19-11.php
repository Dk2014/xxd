
<?php
db_execute($db, "

DROP TABLE IF EXISTS `player_sword_soul`;

CREATE TABLE `player_sword_soul` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '玩家物品ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `pos` smallint(6) NOT NULL COMMENT '位置',
  `sword_soul_id` smallint(6) NOT NULL COMMENT '剑心ID',
  `exp` int(11) NOT NULL COMMENT '当前经验',
  `level` tinyint(4) NOT NULL COMMENT '等级',
  PRIMARY KEY (`id`),
  KEY `ix_player_sword_soul_pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家剑心数据';

DROP TABLE IF EXISTS `player_sword_soul_state`;

CREATE TABLE `player_sword_soul_state` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `fragment_num` smallint(6) DEFAULT NULL COMMENT '碎片数量',
  `box_state` tinyint(4) NOT NULL COMMENT '开箱子的状态(位操作)',
  `last_is_ingot` tinyint(4) DEFAULT NULL COMMENT '上次是用为元宝开箱',
  `num` tinyint(4) NOT NULL COMMENT '今日次数',
  `ingot_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '当日已元宝拔剑次数',
  `update_time` bigint(20) NOT NULL COMMENT '更新时间',
  `is_first_time` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否是第一次拔剑',
  `protect_num` tinyint(4) DEFAULT '0' COMMENT '概率保护次数',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家拔剑状态';

DROP TABLE IF EXISTS `player_sword_soul_equipped`;

CREATE TABLE `player_sword_soul_equipped` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `pos0` bigint(20) NOT NULL COMMENT '装备位置1的剑心',
  `pos1` bigint(20) NOT NULL COMMENT '装备位置2的剑心',
  `pos2` bigint(20) NOT NULL COMMENT '装备位置3的剑心',
  `pos3` bigint(20) NOT NULL COMMENT '装备位置4的剑心',
  `pos4` bigint(20) NOT NULL COMMENT '装备位置5的剑心',
  `pos5` bigint(20) NOT NULL COMMENT '装备位置6的剑心',
  `pos6` bigint(20) NOT NULL COMMENT '装备位置7的剑心',
  `pos7` bigint(20) NOT NULL COMMENT '装备位置8的剑心',
  `pos8` bigint(20) NOT NULL COMMENT '装备位置9的剑心',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家装备的剑心';

");
?>