<?php
db_execute($db, "

DROP TABLE `events_level_up`;
CREATE TABLE `events_level_up` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `require_level` smallint(6) NOT NULL COMMENT '需要等级',
  `ingot` smallint(6) NOT NULL COMMENT '奖励元宝',
  `item1_id` smallint(6) DEFAULT '0' COMMENT '物品1',
  `item1_num` smallint(6) DEFAULT '0' COMMENT '物品1数量',
  `item2_id` smallint(6) DEFAULT '0' COMMENT '物品2',
  `item2_num` smallint(6) DEFAULT '0' COMMENT '物品2数量',
  `item3_id` smallint(6) DEFAULT '0' COMMENT '物品3',
  `item3_num` smallint(6) DEFAULT '0' COMMENT '物品3数量',  
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色升级运营活动';

");
?>
