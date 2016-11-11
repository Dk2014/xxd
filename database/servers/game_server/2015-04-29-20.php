<?php

$this->AddSQL("
create table player_global_clique_kongfu (
	`id` bigint(20) NOT NULL COMMENT 'ID',
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`kongfu_id` int(11) NOT NULL COMMENT '武功ID',
	`level` smallint(6) NOT NULL COMMENT '武功等级',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家帮派武学';
");


?>
