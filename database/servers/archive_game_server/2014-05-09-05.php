<?php

db_execute($db, "
	ALTER TABLE `player_mission_level_record` ADD COLUMN `daily_num` tinyint(4) NOT NULL COMMENT '当日已进入关卡的次数';
");
?>