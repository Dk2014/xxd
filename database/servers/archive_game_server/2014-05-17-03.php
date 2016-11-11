<?php
db_execute($db, "

ALTER TABLE `player_mission_level_record` ADD COLUMN `last_enter_time`  bigint(20) NOT NULL DEFAULT '0' COMMENT '最后一次进入时间';

");
?>
