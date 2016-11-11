<?php

db_execute($db, "
alter table push_notify modify column   `trigger_time` int(11)  NOT NULL DEFAULT '0' COMMENT '触发时间一天内第几秒 [0,86400)';
");
?>
