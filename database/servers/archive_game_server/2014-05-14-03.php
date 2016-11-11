<?php
db_execute($db, "
	ALTER TABLE `item` ADD COLUMN `equip_type_id` int(11) NOT NULL DEFAULT '0' COMMENT '装备等级类型';
	ALTER TABLE `item` ADD COLUMN `health` int(11) DEFAULT '0' COMMENT '生命';
	ALTER TABLE `item` ADD COLUMN `cultivation` int(11) DEFAULT '0' COMMENT '内力';
	ALTER TABLE `item` ADD COLUMN `speed` int(11) DEFAULT '0' COMMENT '速度';
	ALTER TABLE `item` ADD COLUMN `attack` int(11) DEFAULT '0' COMMENT '攻击';
	ALTER TABLE `item` ADD COLUMN `defence` int(11) DEFAULT '0' COMMENT '防御';
	ALTER TABLE `item` ADD COLUMN `dodge_level` int(11) DEFAULT '0' COMMENT '闪避';
	ALTER TABLE `item` ADD COLUMN `hit_level` int(11) DEFAULT '0' COMMENT '命中';
	ALTER TABLE `item` ADD COLUMN `block_level` int(11) DEFAULT '0' COMMENT '格挡';
	ALTER TABLE `item` ADD COLUMN `critical_level` int(11) DEFAULT '0' COMMENT '暴击';
	ALTER TABLE `item` ADD COLUMN `critical_hurt_level` int(11) DEFAULT '0' COMMENT '必杀';
");
?>