<?php

db_execute($db, "
alter table `battle_pet_catch_rate` add column `health_rate` tinyint(4) NOT NULL COMMENT '生效血量百分比';
");

?>
