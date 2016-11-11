<?php 


$this->AddSQL("
create table `buddy_cooperation` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(20) NOT NULL COMMENT '组合名',
  `desc` varchar(100) NOT NULL COMMENT '描述',
  `require_friendship_level` smallint(6) NOT NULL COMMENT '要求羁绊等级',
  `role_id1` tinyint(4) NOT NULL COMMENT '伙伴1',
  `role_id2` tinyint(4) NOT NULL COMMENT '伙伴2',
  `health` int(11) NOT NULL DEFAULT '0' COMMENT '生命 - health',
  `attack` int(11) NOT NULL DEFAULT '0' COMMENT '普攻 - attack',
  `defence` int(11) NOT NULL DEFAULT '0'  COMMENT '普防 - defence',
  `skill_level` int(11) NOT NULL DEFAULT '0' COMMENT '技能熟练度',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='伙伴协力';

alter table player_role add column `coop_role_id` tinyint(4) NOT NULL DEFAULT '0' COMMENT '协力伙伴';
");

?>




