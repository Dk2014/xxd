<?php
	

db_execute($db, "

	ALTER TABLE `player_info` ADD COLUMN `heart_num` tinyint(4) NOT NULL DEFAULT '0' COMMENT '玩家爱心数';
	
");

?>