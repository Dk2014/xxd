<?php 
db_execute($db, "
alter table player_battle_pet_grid modify column   `exp` bigint(20) NOT NULL DEFAULT '0' COMMENT '格子经验';
");
?>
