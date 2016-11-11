<?php
db_execute($db, "
ALTER TABLE `enemy_role` ADD COLUMN `jump_attack` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否跳跃攻击';
");
?>