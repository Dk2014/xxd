<?php
db_execute($db, "

CREATE TABLE `player_state` (
  `pid` bigint(20) unsigned NOT NULL COMMENT '玩家ID',
  `ban_start_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '被冻结的时间',
  `ban_end_time` bigint(20) NOT NULL DEFAULT '-1' COMMENT '被冻结时长, -1 无冻结；0 永久',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家状态';

");

?>
