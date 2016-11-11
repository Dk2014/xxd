<?php
db_execute($db,"
DROP TABLE IF EXISTS `level_star`;
CREATE TABLE `level_star` (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`level_id` int(11) NOT NULL COMMENT '关卡ID',
	`two_star_score` int(11) NOT NULL COMMENT '两星要求分数',
	`three_star_score` int(11) NOT NULL COMMENT '三星要求分数',
	UNIQUE KEY `level_id` (`level_id`),
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='关卡星级分数表';
"); ?>

