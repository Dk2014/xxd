<?php
db_execute($db, "
ALTER TABLE `enemy_role` ADD COLUMN `body_size` tinyint(4) NOT NULL DEFAULT '1' COMMENT '怪物体型';

");
?>