<?php

db_execute($db, "
ALTER TABLE `item` ADD COLUMN `can_batch` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否可以批量使用,0:非，1:是';
");
?>
