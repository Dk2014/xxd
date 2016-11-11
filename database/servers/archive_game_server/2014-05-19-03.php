<?php
db_execute($db, "

	CREATE TABLE `player_func_key` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `key` smallint(6) NOT NULL COMMENT '功能权值',
  `played_key` smallint(6) NOT NULL DEFAULT '0' COMMENT '已播放提示的功能',
  `unique_key` bigint(20) NOT NULL DEFAULT '0' COMMENT '已开启功能的唯一权值',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家功能开放表';

");
?>