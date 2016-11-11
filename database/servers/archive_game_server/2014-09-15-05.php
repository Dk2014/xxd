<?php
db_execute($db, "
alter table `player_login_award_record` add column `update_timestamp` bigint(20) NOT NULL DEFAULT '0' COMMENT '更新时间戳';
")
?>
