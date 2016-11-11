<?php
db_execute($db, "
	ALTER TABLE `ghost` CHANGE hth heath int(11) NOT NULL DEFAULT '0' COMMENT '生命';
	ALTER TABLE `ghost` CHANGE atk  attack  int(11) NOT NULL DEFAULT '0' COMMENT '攻击';
	ALTER TABLE `ghost` CHANGE def defense  int(11) NOT NULL DEFAULT '0' COMMENT '防御';
	ALTER TABLE `ghost` CHANGE spd speed  int(11) NOT NULL DEFAULT '0' COMMENT '速度';
	ALTER TABLE `ghost` CHANGE dge_level dodge_level int(11) NOT NULL DEFAULT '0' COMMENT '闪避等级';
	ALTER TABLE `ghost` CHANGE cri_level crit_level  int(11) NOT NULL DEFAULT '0' COMMENT '暴击等级';
	ALTER TABLE `ghost` CHANGE blk_level block_level int(11) NOT NULL DEFAULT '0' COMMENT '格挡等级';
");
?>