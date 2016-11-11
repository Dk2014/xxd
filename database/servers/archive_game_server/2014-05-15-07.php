<?php
db_execute($db, "
	ALTER TABLE `item` ADD COLUMN `name_prefix` tinyint(4) DEFAULT '0' COMMENT '品质前缀';
");
?>