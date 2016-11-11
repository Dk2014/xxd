<?php

$this->AddSQL("
CREATE TABLE IF NOT EXISTS `player_wegames_uid` (
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`wegames_platform_uid` varchar(50) NOT NULL DEFAULT '' COMMENT 'wegames 平台UID',
	PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家wegames 平台uid';
");

?>
