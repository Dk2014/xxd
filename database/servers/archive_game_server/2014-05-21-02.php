<?php
db_execute($db, "

DROP TABLE IF EXISTS `role_level_exp`;

CREATE TABLE `role_level_exp` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `level` int(11) NOT NULL COMMENT '等级 - level',
  `exp` bigint(20) NOT NULL COMMENT '升到下一级所需经验',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色等级经验表';

");
?>