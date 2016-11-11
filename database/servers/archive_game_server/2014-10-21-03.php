<?php

db_execute($db, "
create table meditation_exp (
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`require_level` int(11) NOT NULL  COMMENT '要求等级',
	`exp` int(11) NOT NULL COMMENT '15秒奖励经验',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='打坐经验';

");

?>
