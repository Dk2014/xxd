
<?php
db_execute($db, 
	"
		ALTER TABLE `player_friend` ADD COLUMN `send_heart_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '上次送爱心时间';

	"
);

?>