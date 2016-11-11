
<?php
db_execute($db, 
	"
		ALTER TABLE `enemy_role` modify `scale_size` smallint(5) NOT NULL DEFAULT '100' COMMENT '怪物缩放比%';

	"
);

?>