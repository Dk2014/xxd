<?php

db_execute($db, "
DROP TABLE IF EXISTS `item_type`;
CREATE TABLE `item_type` (
  `id` tinyint(4) NOT NULL AUTO_INCREMENT COMMENT '类型ID',
  `name` varchar(10) NOT NULL COMMENT '类型名称',
  `max_num_in_pos` smallint(6) NOT NULL DEFAULT '1' COMMENT '每个位置最多可堆叠的数量',
  `sign` varchar(50) DEFAULT '' COMMENT '类型标志',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='物品类型';
");
?>