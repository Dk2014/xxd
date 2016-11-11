<?php

db_execute($db, "
ALTER TABLE `player_event_award_record` MODIFY COLUMN `record_bytes` mediumblob DEFAULT NULL COMMENT '奖励领取状态';
");

?>
