<?php
db_execute($db, 

"
DROP TABLE IF EXISTS `mission_level_small_box_items`;

CREATE TABLE `mission_level_small_box_items` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `small_box_id` int(11) NOT NULL COMMENT '小宝箱id',
  `item_id` int(11) NOT NULL COMMENT '物品ID',
  `award_chance` tinyint(4) NOT NULL COMMENT '奖励概率',

   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='关卡小宝箱';

"
);

?>