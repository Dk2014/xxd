<?php
db_execute($db, "
DROP TABLE IF EXISTS `town_npc_item`;
CREATE TABLE `town_npc_item` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `town_npc_id` int(11) NOT NULL COMMENT '城镇NPC ID',
  `item_id` smallint(6) NOT NULL COMMENT '物品ID',
  `stock` smallint(6) NOT NULL COMMENT '库存',
  `vip` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'vip特供，1表示vip',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='城镇NPC对话';
");
?>