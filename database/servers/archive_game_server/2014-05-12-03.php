<?php

db_execute($db, "

DROP TABLE IF EXISTS `player_info`;

CREATE TABLE `player_info` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `ingot` bigint(20) NOT NULL DEFAULT '0' COMMENT '元宝',
  `coins` bigint(20) NOT NULL DEFAULT '0' COMMENT '铜钱',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家信息表';
  
");
?>
