<?php
db_execute($db, "
alter table `global_announcement` add column `send_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '发送时间';

");

?>
