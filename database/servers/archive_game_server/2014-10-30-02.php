<?php

db_execute($db, "
alter table `battle_pet` add column `parent_pet_id` int(11) NOT NULL DEFAULT '0' COMMENT '父灵宠ID(enemy_role)';

alter table  `battle_pet` add column `health` int(11) NOT NULL DEFAULT '0' COMMENT '生命';
alter table  `battle_pet` add column `attack` int(11) NOT NULL DEFAULT '0' COMMENT '攻击';
alter table  `battle_pet` add column `defence` int(11) NOT NULL DEFAULT '0' COMMENT '防御';
alter table  `battle_pet` add column `speed` int(11) NOT NULL DEFAULT '0' COMMENT '速度';
alter table  `battle_pet` add column `force` int(11) NOT NULL DEFAULT '0' COMMENT '绝招威力';
alter table  `battle_pet` add column `star` tinyint(4) NOT NULL DEFAULT '1' COMMENT '星级 [1,5]';

alter table `battle_pet` drop column `activate_ball_num`;
alter table `battle_pet` drop column `item_battle_pet_id`;

alter table `player_battle_pet` drop column `ball_num`;
alter table `player_battle_pet` drop column `activated`;

alter table `player_battle_pet` add column `star` tinyint(4) NOT NULL DEFAULT '1' COMMENT '星级[1,5]';

alter table `player_battle_pet` add column `parent_pet_id` int(11) NOT NULL DEFAULT '0' COMMENT '父灵宠ID';

");
?>
