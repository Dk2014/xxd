
<?php
db_execute($db, 
"
alter table `player_ghost_equipment` change column `main_ghost_id` `main_ghost_id` bigint(20) NOT NULL COMMENT '主魂侍id,player_ghost主键';
alter table `player_ghost_equipment` change column `assist_ghost_id1` `assist_ghost_id1` bigint(20) NOT NULL COMMENT '主魂侍id,player_ghost主键';
alter table `player_ghost_equipment` change column `assist_ghost_id2` `assist_ghost_id2` bigint(20) NOT NULL COMMENT '主魂侍id,player_ghost主键';
"
);

?>

