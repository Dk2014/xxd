<?php

db_execute($db, "

ALTER TABLE `player_mission_record` ADD COLUMN `town_id` smallint(6) NOT NULL COMMENT '城镇ID' after `pid`;
ALTER TABLE `player_mission_level_record` ADD COLUMN `mission_id` smallint(6) NOT NULL COMMENT '区域ID' after `pid`;

ALTER TABLE `player_mission_record` ADD INDEX idx_town_id (town_id);
ALTER TABLE `player_mission_level_record` ADD INDEX idx_mission_id (mission_id);
");
?>