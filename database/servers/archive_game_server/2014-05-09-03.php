<?php

db_execute($db, "

ALTER TABLE `player_mission_record` modify `mission_id` smallint(6) NOT NULL COMMENT '开启的区域ID';
ALTER TABLE `player_mission_record` modify `open_time` bigint(20) NOT NULL COMMENT '开启的区域时间';
ALTER TABLE `player_mission_level_record` modify `open_time` bigint(20) NOT NULL COMMENT '关卡开启时间';

");
?>