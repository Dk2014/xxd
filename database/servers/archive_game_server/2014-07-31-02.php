<?php
db_execute($db, "
CREATE TABLE `platform_friend_award` (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`require_friend_num` int(11) NOT NULL COMMENT '平台好友数',
	`award_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '奖励类型 0-物品 1-剑心 2-魂侍 3-灵宠契约球 4-爱心 5-铜币 6-元宝 7-体力 ',
	`award_id` smallint(6) NOT NULL COMMENT '物品ID',
	`num` smallint(6) NOT NULL COMMENT '物品数量',
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='平台好友奖励';
");
?>
