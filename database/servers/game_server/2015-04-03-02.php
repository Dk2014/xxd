<?php

$this->AddSQL("
	CREATE TABLE IF NOT EXISTS ghost_relationship(
		`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
		`ghost_a` smallint(6) NOT NULL DEFAULT 0 COMMENT '魂侍A',
		`ghost_a_star` tinyint(4) NOT NULL DEFAULT 0 COMMENT '魂侍A所需星级',
		`ghost_b` smallint(6) NOT NULL DEFAULT 0 COMMENT '魂侍B',
		`ghost_b_star` tinyint(4) NOT NULL DEFAULT 0 COMMENT '魂侍B所需星级',
		PRIMARY KEY(`id`)
	)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='魂侍锁链关系表';
	");

?>