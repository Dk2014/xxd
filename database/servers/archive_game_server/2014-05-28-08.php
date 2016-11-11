
<?php
db_execute($db, 
	"
		ALTER TABLE `player_heart` ADD COLUMN `add_day_count` int(11) NOT NULL DEFAULT '0' COMMENT '每日领取数量';

	"
);

?>