
<?php
db_execute($db, 
	"
		ALTER TABLE `player_heart` ADD COLUMN `add_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最后领取时间';

	"
);

?>