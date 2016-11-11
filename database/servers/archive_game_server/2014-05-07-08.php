<?php

db_execute($db, "
DROP TABLE IF EXISTS `player_buddy`;

CREATE TABLE `player_buddy` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '玩家伙伴角色ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `role_id` tinyint(4) NOT NULL COMMENT '角色ID',
  `level` smallint(6) NOT NULL COMMENT '伙伴等级',
  `exp` bigint(6) NOT NULL COMMENT '伙伴经验',
  `is_in_team` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否在队伍中 0不在 1在',
  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家伙伴角色数据';
  
");
?>
