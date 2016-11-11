
<?php
db_execute($db, 
	"
		ALTER TABLE `mission_level` ADD COLUMN `flip_horizontal` tinyint(4) NOT NULL COMMENT '水平翻转';

	"
);

?>