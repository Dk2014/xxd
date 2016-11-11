<?php
db_execute($db, "

DROP TABLE IF EXISTS `rainbow_level`;
CREATE TABLE `rainbow_level` (
	`id` smallint(6) NOT NULL AUTO_INCREMENT,
	`segment` smallint(6) NOT NULL COMMENT '段数',
	PRIMARY KEY (`id`),
UNIQUE KEY `segment` (`segment`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='极限关卡彩虹桥';

DROP TABLE IF EXISTS `rainbow_level_award`;
CREATE TABLE `rainbow_level_award` (
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
	`mission_level_id` int(11) NOT NULL COMMENT '关卡id',
	`award_type` tinyint(4) NOT NULL COMMENT '奖励类型(0--铜钱;1--道具;2--装备 3--恢复伙伴技能 4--恢复魂侍技能 5--恢复灵宠状态)',
	`award_chance` tinyint(4) NOT NULL COMMENT '奖励概率',
	`award_num` int(11) NOT NULL COMMENT '奖励数量',
	`item_id` int(11) NOT NULL DEFAULT '0' COMMENT '物品ID(物品奖励填写)',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='彩虹关卡宝箱';

ALTER TABLE `mission_level` MODIFY COLUMN `parent_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '关联关卡类型(0-区域关卡;1-资源关卡;2-通天塔;8-难度关卡;9-伙伴关卡;10-灵宠关卡;11-魂侍关卡 12-彩虹关卡)';

ALTER TABLE `mission_level` ADD COLUMN `order` tinyint(4) NOT NULL COMMENT '彩虹关卡顺序';
");
?>
