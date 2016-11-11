<?php
db_execute($db, "

CREATE TABLE `player_physical` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `value` smallint(6) NOT NULL COMMENT '体力值',
  `update_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最后更新时间',
  `buy_count` bigint(20) DEFAULT '0' COMMENT '购买次数',
  `buy_update_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '购买次数更新时间',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家体力表';

");
?>