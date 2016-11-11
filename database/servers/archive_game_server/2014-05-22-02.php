
<?php
db_execute($db, 
	"
		ALTER TABLE `mission_level` ADD COLUMN `award_coin` int(11) NOT NULL COMMENT '奖励铜钱';

	"
);

?>