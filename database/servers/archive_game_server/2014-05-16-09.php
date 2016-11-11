<?php
db_execute($db, "
ALTER TABLE `player_ghost_state` ADD COLUMN `purify_update_time`  bigint(20) NOT NULL DEFAULT '0' COMMENT '净化时间';
");
?>