<?php

db_execute($db, "
	ALTER TABLE `server` ADD COLUMN `ip` varchar(20) NOT NULL COMMENT '游戏服IP';
	ALTER TABLE `server` ADD COLUMN `port` varchar(10) NOT NULL COMMENT '游戏服端口';
");
?>