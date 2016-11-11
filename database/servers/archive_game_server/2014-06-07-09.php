<?php

db_execute($db, "

ALTER TABLE `player_multi_level_info` ADD COLUMN `daily_num` tinyint(4) DEFAULT '0' COMMENT '今日已战斗次数';
ALTER TABLE `player_multi_level_info` ADD COLUMN `battle_time` bigint(20) DEFAULT '0' COMMENT '最近一次战斗时间';

");
?>