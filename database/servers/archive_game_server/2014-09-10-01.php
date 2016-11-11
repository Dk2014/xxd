<?php
db_execute($db, "
alter table  `player_rainbow_level_state_bin` modify column `bin`  MEDIUMBLOB NOT NULL COMMENT '彩虹状态';
")
?>
