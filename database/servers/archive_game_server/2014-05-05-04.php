<?php
db_execute($db, "
	ALTER TABLE `mission_enemy` DROP column `x1`;
	ALTER TABLE `mission_enemy` DROP column `y1`;

	ALTER TABLE `mission_enemy` DROP column `x2`;
	ALTER TABLE `mission_enemy` DROP column `y2`;

	ALTER TABLE `mission_enemy` change x0 x int(11) not null;
	ALTER TABLE `mission_enemy` change y0 y int(11) not null;
");
?>