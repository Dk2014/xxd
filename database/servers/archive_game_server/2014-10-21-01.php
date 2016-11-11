<?php

db_execute($db, "

ALTER TABLE `skill` ADD COLUMN `require_ghost_star` tinyint(4) DEFAULT '0' COMMENT '要求星级，魂侍技能专用';
");

?>
