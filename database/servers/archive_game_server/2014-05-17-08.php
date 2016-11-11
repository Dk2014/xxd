<?php
db_execute($db, "
ALTER TABLE `player_ghost_state` ADD COLUMN `ghost_mission_key` int(11) NOT NULL DEFAULT '0' COMMENT '开启影界最大权值';
");
?>