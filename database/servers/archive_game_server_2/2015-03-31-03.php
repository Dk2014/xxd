<?php 


$this->AddSQL("

CREATE TABLE IF NOT EXISTS  `town_star_awards` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `town_id` smallint(6)   NOT NULL COMMENT '城镇ID',
  `totalstar` smallint(6)  NOT NULL DEFAULT '0' COMMENT '通关回合数',
  `ingot` smallint(6) NOT NULL COMMENT '奖励元宝',
  `coins` int(11) NOT NULL COMMENT '奖励铜钱',
  `heart` smallint(6) NOT NULL COMMENT '奖励爱心',
  `item1_id` smallint(6) DEFAULT '0' COMMENT '物品1',
  `item1_num` smallint(6) DEFAULT '0' COMMENT '物品1数量',
  `item2_id` smallint(6) DEFAULT '0' COMMENT '物品2',
  `item2_num` smallint(6) DEFAULT '0' COMMENT '物品2数量',
  `item3_id` smallint(6) DEFAULT '0' COMMENT '物品3',
  `item3_num` smallint(6) DEFAULT '0' COMMENT '物品3数量',
  `item4_id` smallint(6) DEFAULT '0' COMMENT '物品4',
  `item4_num` smallint(6) DEFAULT '0' COMMENT '物品4数量',
  `item5_id` smallint(6) DEFAULT '0' COMMENT '物品5',
  `item5_num` smallint(6) DEFAULT '0' COMMENT '物品5数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=303 DEFAULT CHARSET=utf8mb4 COMMENT='城镇评星奖励';
");

$this->AddSQL("
CREATE TABLE IF NOT EXISTS `player_mission_star_award` (
  `id` bigint(20) NOT NULL COMMENT '主键',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `town_id` smallint(6) NOT NULL COMMENT '城镇ID',
  `award_level` tinyint(4) NOT NULL DEFAULT '1' COMMENT '已经领奖等级',
   PRIMARY KEY (`id`),
   KEY `msta_pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家区域评星领奖记录';
");

?>

