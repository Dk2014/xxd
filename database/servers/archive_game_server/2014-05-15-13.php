<?php
db_execute($db, "

DROP TABLE IF EXISTS `ingot_ghost_mission`;

CREATE TABLE `ingot_ghost_mission` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `vip_level` tinyint(4) NOT NULL DEFAULT '0' COMMENT '可进入vip等级',
  `ghost_egg_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '魂蛋数量',
  `max_egg_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '一天开启魂蛋数量',
  `egg_num_price` smallint(6) NOT NULL DEFAULT '0' COMMENT '开启魂蛋价格(元宝)',
  `yellow_ghost_rand` smallint(6) NOT NULL DEFAULT '0' COMMENT '金色魂侍概率(万分之)',
  `purple_ghost_rand` smallint(6) NOT NULL DEFAULT '0' COMMENT '紫色魂侍概率(万分之)',
  `orange_ghost_rand` smallint(6) NOT NULL DEFAULT '0' COMMENT '橙色魂侍概率(万分之)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='极暗净土';

");
?>