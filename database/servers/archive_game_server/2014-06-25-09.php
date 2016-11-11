<?php

db_execute($db, "
ALTER TABLE `battle_pet` ADD COLUMN `skill`  smallint(6) NOT NULL COMMENT '灵宠技能';
");
?>
