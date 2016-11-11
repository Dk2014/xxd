<?php

	db_execute($db, "

    ALTER TABLE `player_mission_level_record` change `star` `score` int(11) NOT NULL DEFAULT '0' COMMENT 'boss战得分';

	");
  
?>