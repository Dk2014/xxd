<?php
db_execute($db, "
	
DROP TABLE IF EXISTS `ghost_mission`;


CREATE TABLE `ghost_mission` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `town_id` tinyint(4) NOT NULL DEFAULT '0' COMMENT '城镇id',
  `ghost_mission_key` int(11) NOT NULL DEFAULT '0' COMMENT '进入影界需求权值',
  `senior_ghost_rand` smallint(6) NOT NULL DEFAULT '0' COMMENT '抽中高级魂侍的概率',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='魂侍副本表';


");
?>