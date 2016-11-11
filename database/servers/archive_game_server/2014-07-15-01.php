<?php
db_execute($db,"

alter table `player_resource_level` RENAME TO `player_extend_level`;
alter table `player_extend_level` comment '玩家活动关卡状态';

alter table `player_extend_level` add column `buddy_pass_time` bigint(20) NOT NULL COMMENT '伙伴关卡通关时间' after `exp_pass_time`;
alter table `player_extend_level` add column `pet_pass_time` bigint(20) NOT NULL COMMENT '灵宠关卡通关时间' after `exp_pass_time`;
alter table `player_extend_level` add column `ghost_pass_time` bigint(20) NOT NULL COMMENT '魂侍关卡通关时间' after `exp_pass_time`;

alter table `player_extend_level` add column `buddy_daily_num` tinyint(4) NOT NULL COMMENT '伙伴关卡今日进入次数';
alter table `player_extend_level` add column `pet_daily_num` tinyint(4) NOT NULL COMMENT '灵宠关卡今日进入次数';
alter table `player_extend_level` add column `ghost_daily_num` tinyint(4) NOT NULL COMMENT '魂侍关卡今日进入次数';


alter table `resource_level` RENAME TO `extend_level`;
alter table `extend_level` comment '活动关卡';
alter table `extend_level` add column `level_type` tinyint(4) NOT NULL COMMENT '关卡类型(1-资源关卡;9-伙伴关卡;10-灵宠关卡;11-魂侍关卡)';

alter table `mission_level` change column `parent_type` `parent_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '关联关卡类型(0-区域关卡;1-资源关卡;2-通天塔;8-难度关卡;9-伙伴关卡;10-灵宠关卡;11-魂侍关卡)';
alter table `enemy_deploy_form` change column `battle_type` `battle_type` tinyint(4) NOT NULL COMMENT '战场类型(0-区域关卡;1-资源关卡;2-极限关卡;3-多人关卡;8-难度关卡;9-伙伴关卡;10-灵宠关卡;11-魂侍关卡)';

"
);
?>
