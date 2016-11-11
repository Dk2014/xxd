<?php

db_execute($db, "

	ALTER TABLE `mission` ADD COLUMN `order` tinyint(4) NOT NULL COMMENT '区域开启顺序';
	ALTER TABLE `mission` modify `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '区域ID';
	ALTER TABLE `mission` modify `name` varchar(10) NOT NULL COMMENT '区域名称';
	ALTER TABLE `mission` comment '城镇区域';

	ALTER TABLE `player_mission_level` change `last_lock` `max_lock` int(11) NOT NULL COMMENT '已开启的关卡最大权值';
	
	ALTER TABLE `player_mission` change `last_key` `max_order` tinyint(4) NOT NULL COMMENT '已开启区域的最大序号';

	
");
?>