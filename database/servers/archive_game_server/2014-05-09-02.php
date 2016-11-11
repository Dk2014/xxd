<?php

db_execute($db, "
CREATE TABLE `player_info` (
  `pid` bigint(20) NOT NULL AUTO_INCREMENT,
  `ingot` bigint(20) NOT NULL DEFAULT '0' COMMENT '元宝',
  `coins` bigint(20) NOT NULL DEFAULT '0' COMMENT '铜钱',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家信息表';
  
");
?>
