
<?php
db_execute($db, 
	"
	ALTER TABLE `enemy_boss_script` DROP COLUMN `id`;
	ALTER TABLE `enemy_boss_script` add primary key(`boss_id`);
	"
);

?>