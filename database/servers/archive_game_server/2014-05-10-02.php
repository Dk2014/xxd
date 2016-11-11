<?php

db_execute($db, "

	ALTER TABLE `mission_level` ADD INDEX idx_mission_id (mission_id);
");
?>