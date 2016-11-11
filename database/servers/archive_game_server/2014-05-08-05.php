<?php

db_execute($db, "
DROP TABLE IF EXISTS `player_formation`;

CREATE TABLE `player_formation` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `pos0` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '站位0',
  `pos1` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '站位1',
  `pos2` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '站位2',
  `tactical_grid` tinyint(4) NOT NULL DEFAULT '0' COMMENT '玩家选中的战术格子',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家阵型站位';

");
?>
