<?php

db_execute($db, "

DROP TABLE `chest_item`;
CREATE TABLE `chest_item` (
`id`  int(11) NOT NULL AUTO_INCREMENT COMMENT '主键' ,
`chest_id`  int(11) NOT NULL COMMENT '宝箱id' ,
`item_id`  smallint(6) NOT NULL COMMENT '物品' ,
`item_num`  int(11) NOT NULL COMMENT '数量'  ,
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8mb4 COMMENT='宝箱物品';

");
?>

