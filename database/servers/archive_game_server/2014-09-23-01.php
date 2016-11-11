<?php
db_execute($db, "

CREATE TABLE `events_fight_power` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `lock` smallint(6) NOT NULL COMMENT '档位',
  `fight` int(11) NOT NULL COMMENT '战力',
  `ingot` smallint(6) NOT NULL COMMENT '奖励元宝',
  `item1_id` smallint(6) DEFAULT '0' COMMENT '物品1',
  `item1_num` smallint(6) DEFAULT '0' COMMENT '物品1数量',
  `item2_id` smallint(6) DEFAULT '0' COMMENT '物品2',
  `item2_num` smallint(6) DEFAULT '0' COMMENT '物品2数量',
  `item3_id` smallint(6) DEFAULT '0' COMMENT '物品3',
  `item3_num` smallint(6) DEFAULT '0' COMMENT '物品3数量',  
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='战力运营活动';

CREATE TABLE `player_events_fight_power` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `lock` int(11) NOT NULL COMMENT '当前已奖励的战力档位',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家战力运营活动记录';

");
?>
