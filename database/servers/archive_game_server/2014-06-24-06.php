<?php

db_execute($db, "

DROP TABLE IF EXISTS `sword_soul_quality_level`;
CREATE TABLE `sword_soul_quality_level` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '剑心等级ID',
  `quality_id` smallint(6) NOT NULL COMMENT '品质名称',
  `level` tinyint(4) NOT NULL COMMENT '等级',
  `exp` int(11) NOT NULL COMMENT '升到这一级所需的经验',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='剑心品质' AUTO_INCREMENT=52 ;

");
?>
