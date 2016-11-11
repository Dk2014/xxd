<?php
db_execute($db, "
	ALTER TABLE `mission` DROP column `type`;
	ALTER TABLE `mission_levels` DROP column `fog_mode`;
	ALTER TABLE `mission_levels` ADD COLUMN `type` tinyint(4) NOT NULL COMMENT '关卡类型' after `name`;
");
?>