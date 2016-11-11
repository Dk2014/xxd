<?php

db_execute($db, "
DROP TABLE IF EXISTS `player_town`;

CREATE TABLE `player_town` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `town_id` smallint(6) NOT NULL COMMENT '当前玩家所处的城镇ID',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家城镇数据';

");
?>