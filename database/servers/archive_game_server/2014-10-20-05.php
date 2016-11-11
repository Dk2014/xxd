<?php

db_execute($db, "

ALTER TABLE `battle_item_config` ADD COLUMN `cost_power` tinyint(4) DEFAULT '0' COMMENT '消耗精气';
");

?>
