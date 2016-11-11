<?php

db_execute($db, "
DROP TABLE IF EXISTS `enemy_deploy_form`;

CREATE TABLE `enemy_deploy_form` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL COMMENT '关联此阵法的某表唯一ID',  
  `battle_type` tinyint(4) NOT NULL COMMENT '战场类型',
  `pos1` int(11) NOT NULL DEFAULT '0' COMMENT '位置1上的敌人ID',
  `pos2` int(11) NOT NULL DEFAULT '0' COMMENT '位置2上的敌人ID',
  `pos3` int(11) NOT NULL DEFAULT '0' COMMENT '位置3上的敌人ID',
  `pos4` int(11) NOT NULL DEFAULT '0' COMMENT '位置4上的敌人ID',
  `pos5` int(11) NOT NULL DEFAULT '0' COMMENT '位置5上的敌人ID',
  `pos6` int(11) NOT NULL DEFAULT '0' COMMENT '位置6上的敌人ID',
  `pos7` int(11) NOT NULL DEFAULT '0' COMMENT '位置7上的敌人ID',
  `pos8` int(11) NOT NULL DEFAULT '0' COMMENT '位置8上的敌人ID',
  `pos9` int(11) NOT NULL DEFAULT '0' COMMENT '位置9上的敌人ID',
  `pos10` int(11) NOT NULL DEFAULT '0' COMMENT '位置10上的敌人ID',
  `pos11` int(11) NOT NULL DEFAULT '0' COMMENT '位置11上的敌人ID',
  `pos12` int(11) NOT NULL DEFAULT '0' COMMENT '位置12上的敌人ID',
  `pos13` int(11) NOT NULL DEFAULT '0' COMMENT '位置13上的敌人ID',
  `pos14` int(11) NOT NULL DEFAULT '0' COMMENT '位置14上的敌人ID',
  `pos15` int(11) NOT NULL DEFAULT '0' COMMENT '位置15上的敌人ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='怪物阵法表单';

");
?>