<?php

$this->AddSQL("
ALTER TABLE `player_skill` ADD `skill_trnlv` int(11) NOT NULL DEFAULT '1' COMMENT '技能等级';

CREATE TABLE `skill_training_cost` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `level` smallint(6) NOT NULL COMMENT '等级',
  `cost` bigint(20) NOT NULL COMMENT '费用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb4 COMMENT='角色技能训练价格';

DROP TABLE `ghost_skill_force`;
");

