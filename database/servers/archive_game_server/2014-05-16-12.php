<?php
db_execute($db, "
	ALTER TABLE `player_ghost` drop column `dodge_level`;
	ALTER TABLE `player_ghost` drop column `crit_level`;
	ALTER TABLE `player_ghost` drop column `block_level`;
	ALTER TABLE `player_ghost` drop column `hit_level`;

");
?>