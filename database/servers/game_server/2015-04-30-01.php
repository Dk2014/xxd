<?php

//$this->AddSQL("
//	CREATE TABLE IF NOT EXISTS `player_global_clique_contrib`(
//		`pid` bigint(20) NOT NULL COMMENT '玩家id',
//		`clique_id` bigint(20) NOT NULL COMMENT '玩家帮派ID',
//		`contrib` bigint(20) NOT NULL DEFAULT 0 COMMENT '玩家帮派贡献值',
//		PRIMARY KEY(`pid`)
//	)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家帮派贡献';
//");

$this->AddSQL("drop table player_clique_contrib");
//$this->AddSQL("
//	alter table  player_global_clique_info drop column `contrib`;
//	alter table  player_global_clique_info drop column `contrib_clear_time`;
//");

?>
