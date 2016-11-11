<?php
db_execute($db, "
	

CREATE TABLE `player_ghost_equipment` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `role_id` tinyint(4) NOT NULL COMMENT '角色ID',
  `main_ghost_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '主魂侍id,player_ghost主键',
  `assist_ghost_id1` smallint(6) NOT NULL DEFAULT '0' COMMENT '辅魂侍id1,player_ghost主键',
  `assist_ghost_id2` smallint(6) NOT NULL DEFAULT '0' COMMENT '辅魂侍id2,player_ghost主键',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家魂侍装备表';

");

?>