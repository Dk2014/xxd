<?php

db_execute($db, "
create table `battle_pet_grid_exp` (
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`level` smallint(6) NOT NULL  COMMENT '等级',
	`exp` bigint(20) NOT NULL DEFAULT '0' COMMENT '升级到下一级需要的经验',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='灵宠格子经验表';


create table `battle_pet_grid_attribute` (
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`grid_id` tinyint(4) NOT NULL COMMENT '格子ID',
	`level` smallint(6) NOT NULL COMMENT '格子等级',
	`health` int(11) NOT NULL DEFAULT '0' COMMENT '生命',
	`attack` int(11) NOT NULL DEFAULT '0' COMMENT '攻击',
	`defence` int(11) NOT NULL DEFAULT '0' COMMENT '防御',
	`speed` int(11) NOT NULL DEFAULT '0' COMMENT '速度',
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='灵宠格子属性';

");
?>
