<?php

db_execute($db, "
	ALTER TABLE `mission_level` comment '区域关卡配置';
	ALTER TABLE `mission_level` ADD COLUMN `box_dir` tinyint(4) NOT NULL COMMENT '宝箱朝向(0--左;1--右)';
	ALTER TABLE `mission_enemy` ADD COLUMN `boss_dir` tinyint(4) NOT NULL COMMENT '怪物朝向(0--左;1--右)';
");
?>