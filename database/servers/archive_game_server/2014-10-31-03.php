<?php

db_execute($db, "
CREATE TABLE `battle_pet_catch_rate`(
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`enemy_id` int(11) NOT NULL COMMENT '灵宠ID(怪物ID）',
	`rate` tinyint(4) NOT NULL DEFAULT '0' COMMENT '增加概率',
	PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='灵宠虚弱捕捉概率';
");

?>
