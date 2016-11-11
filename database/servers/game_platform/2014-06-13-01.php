<?php

db_execute($db, "
	ALTER TABLE `account` ADD COLUMN `role_level` smallint(6) NOT NULL DEFAULT '1' COMMENT '角色等级';
");
?>