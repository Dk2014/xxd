<?php
db_execute($db, "

DROP TABLE IF EXISTS `ghost_passive_skill`;

CREATE TABLE `ghost_passive_skill` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `quality` tinyint(4) NOT NULL DEFAULT '0' COMMENT '品质',
  `level` smallint(6) NOT NULL DEFAULT '0' COMMENT '等级',
  `name` varchar(200) NOT NULL DEFAULT '' COMMENT '被动技名称',
  `sign` varchar(200) NOT NULL DEFAULT '' COMMENT '图标标识',
  `desc` varchar(200) NOT NULL DEFAULT '' COMMENT '描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='魂侍被动技能表';

");
?>