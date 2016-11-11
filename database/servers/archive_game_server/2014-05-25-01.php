<?php
db_execute($db, "

CREATE TABLE `player_mission_level_state_bin` (
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`bin` blob NOT NULL COMMENT '状态MissionLevelState',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家区域关卡状态保存';

");
?>