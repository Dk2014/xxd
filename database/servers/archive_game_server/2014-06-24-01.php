<?php

db_execute($db, "

CREATE TABLE `chest_quality` (
`id`  int(11) NOT NULL AUTO_INCREMENT COMMENT '主键' ,
`type`  tinyint(4) NOT NULL COMMENT '类型:1 - 青铜宝箱, 2 - 神龙宝箱' ,
`quality`  tinyint(4) NOT NULL COMMENT '宝箱品质' ,
`probability`  tinyint(4) NOT NULL COMMENT '概率（%）' ,
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8mb4 COMMENT='宝箱品质';

");
?>

