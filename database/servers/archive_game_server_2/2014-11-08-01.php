<?php

$this->AddSQL("
CREATE TABLE `arena_rank_gap` (
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`rank` int(11) NOT NULL COMMENT '排名条件',
	`gap` smallint(16) NOT NULL DEFAULT '0' COMMENT '排名间隔',
	PRIMARY KEY (`id`),
UNIQUE KEY `rank` (`rank`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='比武场挑战排名间隔';
");

?>
