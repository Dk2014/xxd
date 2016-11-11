
<?php
db_execute($db, 
	"
		ALTER TABLE `role` ADD COLUMN `scale` tinyint(4) NOT NULL DEFAULT '100' COMMENT '缩放比';

	"
);

?>