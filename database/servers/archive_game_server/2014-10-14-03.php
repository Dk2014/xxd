<?php

db_execute($db, "
alter table push_notify add column   `name` varchar(30)  NOT NULL  COMMENT '推送通知名称';
");
?>
