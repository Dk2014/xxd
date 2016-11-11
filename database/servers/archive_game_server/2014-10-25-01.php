<?php

db_execute($db, "
alter table `global_announcement` add column `spacing_time` bigint(20) DEFAULT 0 NOT NULL COMMENT '间隔时间';
");

?>
