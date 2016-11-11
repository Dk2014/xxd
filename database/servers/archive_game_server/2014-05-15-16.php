<?php
db_execute($db, "
	ALTER TABLE `player_use_skill` ADD COLUMN `skill_id0` smallint(6) NOT NULL DEFAULT '0' COMMENT '主角默认招式';
");
?>