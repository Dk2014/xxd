<?php
db_execute($db, "

DROP TABLE `ghost_tip`;
CREATE TABLE `ghost_tip` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `tip` varchar(512) DEFAULT NULL COMMENT '提示信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='魂侍提示信息';



");
?>
