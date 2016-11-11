<?php 
$this->AddSQL("
create table `equipment_resonance` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `require_level` smallint(6) NOT NULL COMMENT '要求等级',
  `health` int(11) NOT NULL COMMENT '生命 - health',
  `attack` int(11) NOT NULL COMMENT '普攻 - attack',
  `defence` int(11) NOT NULL COMMENT '普防 - defence',
  PRIMARY KEY (`id`),
  UNIQUE KEY `require_level` (`require_level`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='装备共鸣';
");
$this->AddSQL("
CREATE TABLE `player_driving_sword_eventmask` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `cloud_id` smallint(6) NOT NULL COMMENT '云层id',
  `events` blob NOT NULL COMMENT '事件比特阵列',
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家云海事件地图';

alter table `player_driving_sword_map` drop `event_mask`, drop `opened_area_count`, drop `hole_count`, drop `teleport_count`;
");


?>
