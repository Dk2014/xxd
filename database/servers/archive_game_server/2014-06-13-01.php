
<?php
db_execute($db, 
"
    ALTER TABLE `player_ghost_mission` MODIFY COLUMN `mission_id` smallint(6) NOT NULL COMMENT '关卡主键id';
"
);

?>

