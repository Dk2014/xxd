<?php
db_execute($db,"

ALTER TABLE `player_sword_soul_state`
DROP COLUMN `fragment_num`,
DROP COLUMN `last_is_ingot`,
DROP COLUMN `free_num`,
DROP COLUMN `ingot_num`,
DROP COLUMN `is_first_time`,
DROP COLUMN `protect_num`;

"
);
?>
