<?php

db_execute($db, "
DROP TABLE IF EXISTS `player_roles`;

CREATE TABLE `player_roles` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '玩家角色ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `role_id` tinyint(4) NOT NULL COMMENT '角色模板ID',
  `level` smallint(6) NOT NULL COMMENT '等级',
  `exp` bigint(6) NOT NULL COMMENT '经验',
  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家角色数据';

");
?>