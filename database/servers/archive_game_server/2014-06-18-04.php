<?php

db_execute($db, "

DROP TABLE IF EXISTS `player_item_recast_state`;
CREATE TABLE `player_item_recast_state` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `player_item_id` bigint(20) NOT NULL COMMENT '玩家装备ID',
  `selected_attr` tinyint(4) NOT NULL COMMENT '选中的属性',
  `attr1_type` tinyint(4) NOT NULL COMMENT '重铸属性1类型',
  `attr1_value` int(11) NOT NULL COMMENT '重铸属性1数值',
  `attr2_type` tinyint(4) NOT NULL COMMENT '重铸属性2类型',
  `attr2_value` int(11) NOT NULL COMMENT '重铸属性2数值',
  `attr3_type` tinyint(4) NOT NULL COMMENT '重铸属性3类型',
  `attr3_value` int(11) NOT NULL COMMENT '重铸属性3数值',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家装备重铸状态';

");
?>
