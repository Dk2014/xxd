<?php

db_execute($db, "
DROP TABLE IF EXISTS `player_skill`;
CREATE TABLE `player_skill` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `role_id` tinyint(4) NOT NULL COMMENT '角色ID',
  `skill_id` smallint(6) NOT NULL COMMENT '绝招ID',
  `level` smallint(6) NOT NULL DEFAULT '1' COMMENT '绝招等级',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家角色绝招表';

DROP TABLE IF EXISTS `player_use_skill`;
CREATE TABLE `player_use_skill` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `role_id` tinyint(4) NOT NULL COMMENT '角色ID',
  `skill_id1` smallint(6) NOT NULL DEFAULT '0' COMMENT '招式1',
  `skill_id2` smallint(6) NOT NULL DEFAULT '0' COMMENT '招式2',
  `skill_id3` smallint(6) NOT NULL DEFAULT '0' COMMENT '招式3',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家角色当前使用的绝招表';
");
?>