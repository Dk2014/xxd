<?php

db_execute($db, "
DROP TABLE IF EXISTS `item`;
CREATE TABLE `item` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '物品ID',
  `type_id` int(11) NOT NULL COMMENT '类型ID',
  `quality` tinyint(4) DEFAULT NULL COMMENT '品质',
  `name` varchar(20) NOT NULL COMMENT '物品名称',
  `level` int(11) DEFAULT NULL COMMENT '需求等级',
  `desc` varchar(100) DEFAULT NULL COMMENT '物品描述',
  `price` int(11) NOT NULL DEFAULT '0' COMMENT '物品售价',
  `can_use` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否可在格子中使用，0不能，1反之',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='物品';
");
?>