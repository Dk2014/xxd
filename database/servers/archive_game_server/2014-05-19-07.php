
<?php
db_execute($db, 
	"
		ALTER TABLE `enemy_role` ADD COLUMN `scale_size` tinyint(4) NOT NULL DEFAULT '1' COMMENT '怪物缩放比%';

	"
);
?>