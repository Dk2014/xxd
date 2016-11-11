<?php

$this->AddSQL("
	CREATE TABLE IF NOT EXISTS `player_clique_contrib`(
		`pid` bigint(20) NOT NULL COMMENT '玩家id',
		`contrib` bigint(20) NOT NULL DEFAULT 0 COMMENT '玩家帮派贡献值',
		`updatetime` bigint(20) NOT NULL DEFAULT 0 COMMENT '更新时间戳',
		PRIMARY KEY(`pid`)
	)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家帮派贡献';
");

?>