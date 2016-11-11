<?php

db_execute($db, "

  ALTER TABLE `skill` ADD `order` bigint(20) NOT NULL DEFAULT '0' COMMENT '排序字段';
");
?>