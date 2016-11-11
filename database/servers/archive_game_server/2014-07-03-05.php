
<?php
db_execute($db, "

ALTER TABLE `player_ghost_equipment`
MODIFY COLUMN `ghost_power`  int(11) NOT NULL COMMENT '魂力' AFTER `role_id`;

");
?>
