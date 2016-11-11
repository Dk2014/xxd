
<?php
db_execute($db, 
"

CREATE TABLE `player_tower_level` (
   `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `floor` smallint(6) NOT NULL COMMENT '当前层数',
  `battle_state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '当前楼层战斗状态(0--从未打过; 1--失败)',
  `open_time` bigint(20) NOT NULL COMMENT '开启当前层数的时间',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家战斗力';

"

);
?>