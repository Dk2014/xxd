<?php

	db_execute($db, "

    update `player_mission_level_record` set `score` = '0';
    
	");

?>