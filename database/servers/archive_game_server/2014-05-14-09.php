<?php
db_execute($db, "
	ALTER TABLE `item` ADD COLUMN `sign` varchar(30) DEFAULT NULL COMMENT '资源标识';
");
?>