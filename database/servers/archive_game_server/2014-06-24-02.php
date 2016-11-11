<?php

db_execute($db, "

CREATE TABLE `chest_item` (
`id`  int(11) NOT NULL AUTO_INCREMENT COMMENT '主键' ,
`type`  tinyint(4) NOT NULL COMMENT '类型:1 - 青铜宝箱, 2 - 神龙宝箱' ,
`quality`  tinyint(4) NOT NULL COMMENT '品质' ,
`item_id`  smallint(6) NOT NULL COMMENT '物品' ,
`item_num`  int(11) NOT NULL COMMENT '数量'  ,
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8mb4 COMMENT='宝箱物品';

");
?>

