<?php
db_execute($db, "

ALTER TABLE `player_town` ADD COLUMN `lock` int(11) NOT NULL COMMENT '当前拥有的城镇权值' after `town_id`;

ALTER TABLE `mission_levels` RENAME TO `mission_level`;
ALTER TABLE `mission_level` ADD COLUMN `lock` int(11) NOT NULL COMMENT '关卡开启的权值' after `mission_id`;
ALTER TABLE `mission_level` ADD COLUMN `sign_war` varchar(50) NOT NULL COMMENT '关卡战斗资源标识' after `sign`;

ALTER TABLE `mission_level` ADD COLUMN `enter_x` int(11) NOT NULL COMMENT '出生点x坐标' after `type`;
ALTER TABLE `mission_level` ADD COLUMN `enter_y` int(11) NOT NULL COMMENT '出生点y坐标' after `type`;

ALTER TABLE `mission_level` ADD COLUMN `daily_num` tinyint(4) NOT NULL COMMENT '允许每天进入次数,0表示不限制' after `type`;
ALTER TABLE `mission_level` ADD COLUMN `physical` tinyint(4) NOT NULL COMMENT '每次进入消耗的体力' after `daily_num`;

ALTER TABLE `mission_level` ADD COLUMN `box_x` int(11) NOT NULL COMMENT '宝箱x坐标' after `physical`;
ALTER TABLE `mission_level` ADD COLUMN `box_y` int(11) NOT NULL COMMENT '宝箱y坐标' after `box_x`;

ALTER TABLE `mission_level` ADD COLUMN `award_key` tinyint(4) NOT NULL COMMENT '奖励钥匙数' after `box_y`;
ALTER TABLE `mission_level` ADD COLUMN `award_exp` tinyint(4) NOT NULL COMMENT '奖励经验' after `award_key`;

ALTER TABLE `mission_level` ADD COLUMN `award_item1_type` tinyint(4) NOT NULL COMMENT '物品1的奖励类型(1--铜钱；2--真气龙珠;3--龙币;4--道具;5--装备)' after `award_exp`;
ALTER TABLE `mission_level` ADD COLUMN `award_item1_id` smallint(6) NOT NULL COMMENT '奖励物品1' after `award_item1_type`;
ALTER TABLE `mission_level` ADD COLUMN `award_item1_chance` tinyint(4) NOT NULL COMMENT '奖励物品1概率' after `award_item1_id`;
ALTER TABLE `mission_level` ADD COLUMN `award_item1_num` tinyint(4) NOT NULL COMMENT '奖励物品1数量' after `award_item1_chance`;

ALTER TABLE `mission_level` ADD COLUMN `award_item2_type` tinyint(4) NOT NULL COMMENT '物品2的奖励类型(1--铜钱；2--真气龙珠;3--龙币;4--道具;5--装备)' after `award_item1_num`;
ALTER TABLE `mission_level` ADD COLUMN `award_item2_id` smallint(6) NOT NULL COMMENT '奖励物品2' after `award_item2_type`;
ALTER TABLE `mission_level` ADD COLUMN `award_item2_chance` tinyint(4) NOT NULL COMMENT '奖励物品2概率' after `award_item2_id`;
ALTER TABLE `mission_level` ADD COLUMN `award_item2_num` tinyint(4) NOT NULL COMMENT '奖励物品2数量' after `award_item2_chance`;

ALTER TABLE `mission_level` ADD COLUMN `award_item3_type` tinyint(4) NOT NULL COMMENT '物品3的奖励类型(1--铜钱；2--真气龙珠;3--龙币;4--道具;5--装备)' after `award_item2_num`;
ALTER TABLE `mission_level` ADD COLUMN `award_item3_id` smallint(6) NOT NULL COMMENT '奖励物品3' after `award_item3_type`;
ALTER TABLE `mission_level` ADD COLUMN `award_item3_chance` tinyint(4) NOT NULL COMMENT '奖励物品3概率' after `award_item3_id`;
ALTER TABLE `mission_level` ADD COLUMN `award_item3_num` tinyint(4) NOT NULL COMMENT '奖励物品3数量' after `award_item3_chance`;

ALTER TABLE `mission_level` ADD COLUMN `award_item4_type` tinyint(4) NOT NULL COMMENT '物品4的奖励类型(1--铜钱；2--真气龙珠;3--龙币;4--道具;5--装备)' after `award_item3_num`;
ALTER TABLE `mission_level` ADD COLUMN `award_item4_id` smallint(6) NOT NULL COMMENT '奖励物品4' after `award_item4_type`;
ALTER TABLE `mission_level` ADD COLUMN `award_item4_chance` tinyint(4) NOT NULL COMMENT '奖励物品4概率' after `award_item4_id`;
ALTER TABLE `mission_level` ADD COLUMN `award_item4_num` tinyint(4) NOT NULL COMMENT '奖励物品4数量' after `award_item4_chance`;

ALTER TABLE `mission_level` ADD COLUMN `award_item5_type` tinyint(4) NOT NULL COMMENT '物品5的奖励类型(1--铜钱；2--真气龙珠;3--龙币;4--道具;5--装备)' after `award_item4_num`;
ALTER TABLE `mission_level` ADD COLUMN `award_item5_id` smallint(6) NOT NULL COMMENT '奖励物品5' after `award_item5_type`;
ALTER TABLE `mission_level` ADD COLUMN `award_item5_chance` tinyint(4) NOT NULL COMMENT '奖励物品5概率' after `award_item5_id`;
ALTER TABLE `mission_level` ADD COLUMN `award_item5_num` tinyint(4) NOT NULL COMMENT '奖励物品5数量' after `award_item5_chance`;


ALTER TABLE `mission_enemy` CHANGE `x` `enter_x` int(11) NOT NULL COMMENT '出生点x坐标';
ALTER TABLE `mission_enemy` CHANGE `y` `enter_y` int(11) NOT NULL COMMENT '出生点y坐标';

ALTER TABLE `mission_enemy` CHANGE `monster2_id` `monster3_id` int(11) NOT NULL COMMENT '怪物3 ID';
ALTER TABLE `mission_enemy` CHANGE `monster2_chance` `monster3_chance` tinyint(4) NOT NULL COMMENT '出现概率';
ALTER TABLE `mission_enemy` CHANGE `monster1_id` `monster2_id` int(11) NOT NULL COMMENT '怪物2 ID';
ALTER TABLE `mission_enemy` CHANGE `monster1_chance` `monster2_chance` tinyint(4) NOT NULL COMMENT '出现概率';
ALTER TABLE `mission_enemy` CHANGE `monster0_id` `monster1_id` int(11) NOT NULL COMMENT '怪物1 ID';
ALTER TABLE `mission_enemy` CHANGE `monster0_chance` `monster1_chance` tinyint(4) NOT NULL COMMENT '出现概率';

ALTER TABLE `mission_enemy` ADD COLUMN `monster4_id` int(11) NOT NULL COMMENT '怪物4 ID' after `monster3_id`;
ALTER TABLE `mission_enemy` ADD COLUMN `monster4_chance` tinyint(4) NOT NULL COMMENT '出现概率' after `monster4_id`;

ALTER TABLE `mission_enemy` ADD COLUMN `monster5_id` int(11) NOT NULL COMMENT '怪物5 ID' after `monster4_chance`;
ALTER TABLE `mission_enemy` ADD COLUMN `monster5_chance` tinyint(4) NOT NULL COMMENT '出现概率' after `monster5_id`;
");
?>
