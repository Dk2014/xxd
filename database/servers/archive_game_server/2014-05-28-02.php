<?php
db_execute($db, "

DROP TABLE IF EXISTS `player_heart`;

CREATE TABLE `player_heart` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `value` smallint(6) NOT NULL COMMENT '爱心值',
  `update_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最后更新时间',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='玩家爱心表';

");
?>