<?php

db_execute($db, "

DROP TABLE IF EXISTS `player_equipment`;

CREATE TABLE `player_equipment` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `role_id` tinyint(4) NOT NULL COMMENT '角色ID',
  `weapon_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '武器的player_item表主键ID',
  `clothes_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '战袍的player_item表主键ID',
  `accessories_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '饰品的player_item表主键ID',
  `shoe_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '靴子的player_item表主键ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家装备表';
  
");
?>