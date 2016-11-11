<?php

$this->AddSQL("
CREATE TABLE `player_cornucopia` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `open_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '开启更新时间',
  `daily_count` smallint(6) DEFAULT '0' COMMENT '当天开启次数',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家铜币兑换表';
");

