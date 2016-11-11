<?php
db_execute($db, 
"
DROP TABLE IF EXISTS `player_fight_num`;
CREATE TABLE `player_fight_num` (
   `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `fight_num` int(11) NOT NULL COMMENT '战力力',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家战斗力';

"
);

?>