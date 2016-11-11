<?php
db_execute($db, "
  CREATE TABLE `town_npc` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `town_id` smallint(6) NOT NULL COMMENT '城镇ID',
  `name` varchar(10) NOT NULL COMMENT 'NPC名称',
  `sign` varchar(20) NOT NULL COMMENT '资源标识',
  `x` int(11) NOT NULL COMMENT 'x轴坐标',
  `y` int(11) NOT NULL COMMENT 'y轴坐标',
  `direction` varchar(20) DEFAULT NULL COMMENT '朝向',
  `talk` varchar(200) NOT NULL COMMENT '对话',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='城镇NPC';
");
?>