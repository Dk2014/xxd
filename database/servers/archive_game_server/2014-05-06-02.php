<?php

db_execute($db, "
	ALTER TABLE `player` ADD COLUMN `main_role_id` bigint(20) NOT NULL COMMENT '主角ID';
");
?>