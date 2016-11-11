<?php

db_execute($db, "

DROP TABLE IF EXISTS `sword_soul_type`;
CREATE TABLE `sword_soul_type` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '类型ID',
  `name` varchar(10) NOT NULL COMMENT '类型名称',
  `sign` varchar(20) DEFAULT NULL COMMENT '程序标示',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='剑心类型' AUTO_INCREMENT=14 ;

");
?>
