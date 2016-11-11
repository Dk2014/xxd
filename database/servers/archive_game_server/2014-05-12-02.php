<?php
db_execute($db, "
	ALTER TABLE `player_skill` DROP COLUMN `level`;
");
?>