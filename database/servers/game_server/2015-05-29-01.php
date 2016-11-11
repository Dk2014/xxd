<?php

$this->AddSQL("
create table  `player_global_chat_state` (
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`ban_until` bigint(20) NOT NULL DEFAULT '0' COMMENT '<=0 没有禁言 >0禁言',
	PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='禁言状态';
");


?>
