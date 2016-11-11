<?php
db_execute($db, "

CREATE TABLE `events_level_up` (
	`require_level` smallint(6) NOT NULL COMMENT '需要等级',
	`ingot` smallint(6) NOT NULL COMMENT '奖励元宝',
	PRIMARY KEY (`require_level`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色升级运营活动';

");

?>
