
<?php
db_execute($db, "

ALTER TABLE `player_arena` CHANGE COLUMN `record_read_time` `battle_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最近一次挑战时间';

");
?>
