<?php

db_execute($db, "

	ALTER TABLE `mission_level` ADD COLUMN `award_lock` tinyint(4) NOT NULL COMMENT '通关奖励权值';
	
");
?>