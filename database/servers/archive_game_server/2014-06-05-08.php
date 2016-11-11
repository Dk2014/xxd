<?php
db_execute($db, 

"
DROP TABLE IF EXISTS `mission_level_small_box`;

CREATE TABLE `mission_level_small_box` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `mission_level_id` int(11) NOT NULL COMMENT '关卡id',
  `box_x` int(11) NOT NULL COMMENT '宝箱x坐标',
  `box_y` int(11) NOT NULL COMMENT '宝箱y坐标', 
  `award_chance` tinyint(4) NOT NULL COMMENT '奖励概率',


    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='关卡小宝箱';

"
);

?>