<?php
db_execute($db, "
	ALTER TABLE `enemy_role` DROP column `aoe_reduce`;
	ALTER TABLE `mission_enemy` ADD COLUMN `monster_num` tinyint(4) NOT NULL COMMENT '怪物数量' after `mission_level_id`;
");
?>