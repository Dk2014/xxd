<?php
db_execute($db, "

alter table `skill` change column `order` `order` smallint(6) NOT NULL DEFAULT '0' COMMENT '排序字段';

");

?>
