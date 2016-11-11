<?php
db_execute($db, "

CREATE TABLE `player_resource_level` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `pass_time` bigint(20) NOT NULL COMMENT '当前通关时间',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家资源关卡信息';

");
?>
