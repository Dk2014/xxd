<?php
db_execute($db, "
	ALTER TABLE `item_type` ADD COLUMN `order` int(11) DEFAULT '0' COMMENT '客户端排序权重';
");
?>