
<?php
db_execute($db, 
	"
	ALTER TABLE `enemy_role` ADD COLUMN `skill2_id` smallint(6) NOT NULL COMMENT '绝招2 ID' after `skill_force`;
	ALTER TABLE `enemy_role` ADD COLUMN `skill2_force` int(11) NOT NULL COMMENT '绝招2 威力' after `skill2_id`;


	CREATE TABLE `enemy_boss_script` (
		`id` int(11) NOT NULL AUTO_INCREMENT,
		`boss_id` int(10) unsigned NOT NULL COMMENT '怪物ID',
		`config` text COMMENT '脚本',
		PRIMARY KEY (`id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='boss配置脚本';

	"
);

?>