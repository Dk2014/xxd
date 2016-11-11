<?php
db_execute($db,"

CREATE TABLE `player_is_operated` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `operate_value` bigint(20) NOT NULL DEFAULT '0' COMMENT '操作值',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='记录玩家是否第一次操作';

"
);
?>
