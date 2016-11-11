<?php

db_execute($db, "

DROP TABLE IF EXISTS `player_item_buyback`;

CREATE TABLE `player_item_buyback` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `back_id1` bigint(20) NOT NULL DEFAULT '0' COMMENT '回购格子1,player_item表主键ID',
  `back_id2` bigint(20) NOT NULL DEFAULT '0' COMMENT '回购格子2,player_item表主键ID',
  `back_id3` bigint(20) NOT NULL DEFAULT '0' COMMENT '回购格子3,player_item表主键ID',
  `back_id4` bigint(20) NOT NULL DEFAULT '0' COMMENT '回购格子4,player_item表主键ID',
  `back_id5` bigint(20) NOT NULL DEFAULT '0' COMMENT '回购格子5,player_item表主键ID',
  `back_id6` bigint(20) NOT NULL DEFAULT '0' COMMENT '回购格子6,player_item表主键ID',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家物品回购表';
  
");
?>