<?php
db_execute($db, "
alter table fashion drop column fashion_item;
alter table `fashion` drop column `valid_hours`;
create table `fashion_exchange` (
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
	`fashion_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '时装id',
	`item_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '物品id',
	PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='时装兑换表';
");
?>
