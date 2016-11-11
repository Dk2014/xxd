
<?php
db_execute($db, 
	"
		ALTER TABLE `quest` drop column `mission_id`;
		ALTER TABLE `quest` change column `award_mission_key` `award_mission_level_lock` int(11) NOT NULL DEFAULT '0' COMMENT '奖励关卡权值';

	"
);

?>