<?php
db_execute($db, "
DROP TABLE IF EXISTS `rainbow_level`;
CREATE TABLE `rainbow_level` (
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`segment` smallint(6) NOT NULL COMMENT '段数',
	PRIMARY KEY (`id`),
UNIQUE KEY `segment` (`segment`)
  ) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='极限关卡彩虹桥';

alter table `rainbow_level_award` modify column `award_type` tinyint(4) NOT NULL COMMENT '奖励类型(0--铜钱;1--道具;2--装备 3--经验 4--经验倍数 5--铜钱倍数 6--恢复伙伴技能 7--恢复魂侍技能 8--恢复灵宠状态)';

");
?>
