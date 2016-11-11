<?php

$this->AddSQL("
    CREATE TABLE IF NOT EXISTS  `mission_star_awards` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
   `mission_id` smallint(6) NOT NULL COMMENT '区域ID',
  `totalstar` smallint(6) NOT NULL DEFAULT '0' COMMENT '通关回合数',
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=utf8mb4 COMMENT='城镇评星奖励';
");

?>
