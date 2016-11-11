<?php

db_execute($db, "
create table `player_battle_pet_grid` (
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`pet_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '灵宠ID',
	`grid_id` tinyint(4) NOT NULL COMMENT '灵宠格子ID',
	`level` smallint(6) NOT NULL DEFAULT '1' COMMENT '格子等级',
	`exp` int(11) NOT NULL DEFAULT '0' COMMENT '格子经验',
	PRIMARY KEY (`id`),
KEY `idx_pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家灵宠';

");
?>
