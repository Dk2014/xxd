<?php

db_execute($db, "
DROP TABLE IF EXISTS `player_item`;
CREATE TABLE `player_item` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `item_id` smallint(6) NOT NULL COMMENT '物品ID',
  `pos` smallint(6) unsigned NOT NULL COMMENT '位置',
  `num` smallint(6) NOT NULL COMMENT '数量',
  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家物品';
");
?>