<?php
db_execute($db, "
ALTER TABLE `player_hard_level_record` DROP COLUMN pass_of_times;
ALTER TABLE `player_hard_level` ADD COLUMN award_lock int(12) default 0 COMMENT '已获得过奖励关卡的最大lock';
ALTER TABLE `player_mission_level` ADD COLUMN award_lock int(12) default 0 COMMENT '已获得过奖励关卡的最大lock';
");