<?php
db_execute($db, "
	ALTER TABLE `player_ghost_state` DROP COLUMN `gold_fail_num`;
	ALTER TABLE `player_ghost_state` DROP COLUMN `gold_probability`;
	ALTER TABLE `player_ghost_state` DROP COLUMN `purple_fail_num`;
	ALTER TABLE `player_ghost_state` DROP COLUMN `purple_probability`;
	ALTER TABLE `player_ghost_state` DROP COLUMN `ingot_egg_buy_day_count`;
	ALTER TABLE `player_ghost_state` DROP COLUMN `ingot_egg_buy_update_time`;
");
?>