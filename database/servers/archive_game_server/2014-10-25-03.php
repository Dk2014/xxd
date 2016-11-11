<?php

db_execute($db, "
ALTER TABLE `player_rainbow_level_state_bin` MODIFY COLUMN `bin` mediumblob DEFAULT NULL COMMENT '彩虹状态';
");

?>
