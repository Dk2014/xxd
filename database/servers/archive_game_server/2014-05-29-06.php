<?php
db_execute($db, 
"
DROP TABLE IF EXISTS `item_exchange`;
CREATE TABLE `item_exchange` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `target_item_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '目标物品id',
  `item_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '物品id',
  `item_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '物品数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='物品兑换表';

"
);

?>