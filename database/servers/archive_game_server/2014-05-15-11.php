<?php
db_execute($db, "

DROP TABLE IF EXISTS `ghost_item_exchange`;

CREATE TABLE `ghost_item_exchange` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `target_ghost` smallint(6) NOT NULL DEFAULT '0' COMMENT '魂侍id',
  `item_id1` smallint(6) NOT NULL DEFAULT '0' COMMENT '物品id',
  `item_num1` smallint(6) NOT NULL DEFAULT '0' COMMENT '物品数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='物品兑换魂侍表';

");
?>