<?php

$this->AddSQL("
	CREATE TABLE IF NOT EXISTS `player_money_tree`(
		`pid` bigint(20) NOT NULL COMMENT '玩家id',
		`total` int(11) NOT NULL DEFAULT 0 COMMENT '摇钱树铜钱总额',
		`waved_times_total` tinyint(4) NOT NULL DEFAULT 0 COMMENT '总的摇树次数',
		`waved_times` tinyint(4) NOT NULL DEFAULT 0 COMMENT '当天已经摇过的次数',
		`last_waved_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '上次摇树的时间戳',
		PRIMARY KEY(`pid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家摇钱树记录';
");

?>