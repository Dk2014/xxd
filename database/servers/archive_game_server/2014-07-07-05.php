<?php
db_execute($db, "

alter table `player_global_arena_rank` add column `time`  bigint(20) NOT NULL COMMENT '宝箱刷新时间';

");
?>
