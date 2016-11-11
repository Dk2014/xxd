<?php

db_execute($db, "

ALTER TABLE `arena_award_box` CHANGE COLUMN `fame` `ingot` int(11) NOT NULL COMMENT '元宝';
ALTER TABLE `player_arena` DROP COLUMN `daily_award_coin` ;


");
?>
