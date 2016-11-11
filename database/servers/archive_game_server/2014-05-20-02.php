
<?php
db_execute($db, 
	"
		ALTER TABLE `ingot_ghost_mission` DROP COLUMN `ghost_egg_num`;
		ALTER TABLE `ingot_ghost_mission` DROP COLUMN `egg_num_price`;
	"
);

?>