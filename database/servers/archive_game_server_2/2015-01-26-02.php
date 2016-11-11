<?php

$this->AddSQL("
CREATE TABLE `teamship` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '数据ID',
  `level` smallint(6) NOT NULL COMMENT '团队单项加成等级',
  `needs_relationship` int(11) NOT NULL COMMENT '所需友情点数',
  `health` int(11) NOT NULL COMMENT '生命项加成',
  `attack` int(11) NOT NULL COMMENT '攻击项加成',
  `defence` int(11) NOT NULL COMMENT '防御项加成',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=241 DEFAULT CHARSET=utf8mb4 COMMENT='团队配合数据';
");

$this->AddSQL("
alter table `mission_level` add `award_relationship` int(11) NOT NULL DEFAULT '0' COMMENT '奖励友情';

alter table `multi_level` add `award_relationship` int(11) NOT NULL DEFAULT '0' COMMENT '奖励友情';
");

