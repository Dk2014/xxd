<?php
db_execute($db, "
	ALTER TABLE `enemy_role` DROP column `skill_id2`;
	ALTER TABLE `enemy_role` DROP column `skill_force2`;
");
?>