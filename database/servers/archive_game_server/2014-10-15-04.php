<?php

db_execute($db, "
ALTER TABLE `quest_activity_center` ADD COLUMN `condition_template` varchar(60) COMMENT '领奖条件模版,{val}代表临界值';
ALTER TABLE `quest_activity_center` MODIFY COLUMN `start` bigint(20) DEFAULT 0 COMMENT '活动开始时间戳';
ALTER TABLE `quest_activity_center` MODIFY COLUMN `end` bigint(20) DEFAULT 0 COMMENT '活动结束时间戳';
ALTER TABLE `quest_activity_center` DROP COLUMN `dispose_days`;
ALTER TABLE `quest_activity_center` ADD COLUMN `dispose` bigint(20) COMMENT '活动过期时间戳';
 ");
?>