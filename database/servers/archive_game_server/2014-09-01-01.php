<?php
db_execute($db, "
	ALTER TABLE `quest` drop column `auto_mission_level_id`;
	ALTER TABLE `quest` add column `show_black_curtain` tinyint(1) not null default '0' comment '显示黑幕 1--显示 0--不显示';
");
?>
