<?php

$this->AddSQL("
CREATE TABLE `player_team_info` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `relationship` int(11) NOT NULL COMMENT '友情点数',
  `health_lv` smallint(6) NOT NULL COMMENT '生命项等级',
  `attack_lv` smallint(6) NOT NULL COMMENT '攻击项等级',
  `defence_lv` smallint(6) NOT NULL COMMENT '防御项等级',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家队伍相关信息';
");

