<?php
db_execute($db, "
drop table if exists `purchase_limit`;
CREATE TABLE `purchase_limit` (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
	`item_id` smallint(6) NOT NULL COMMENT '物品ID',
	`num` smallint(6) NOT NULL COMMENT '购买次数限制',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='物品购买次数限制';
");

?>
