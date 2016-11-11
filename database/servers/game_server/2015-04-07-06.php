<?php 


$this->AddSQL("
alter table buddy_cooperation 
  	add column `speed` int(11) NOT NULL DEFAULT '0' COMMENT '速度',
  	add column `cultivation` int(11) DEFAULT '0' COMMENT '内力',
  	add column `sunder` int(11) DEFAULT '0' COMMENT '护甲',
	add column `dodge_level` int(11) DEFAULT '0' COMMENT '闪避等级',
	add column `hit_level` int(11) DEFAULT '0' COMMENT '命中等级',
	add column `block_level` int(11) DEFAULT '0' COMMENT '格挡等级',
	add column `tenacity_level` int(11) DEFAULT '0' COMMENT '韧性等级',
	add column `destroy_level` int(11) DEFAULT '0' COMMENT '破击等级',
	add column `critical_level` int(11) DEFAULT '0' COMMENT '暴击等级',
	add column `critical_hurt_level` int(11) DEFAULT '0' COMMENT '必杀等级',
	add column `ghost_power` int(11) DEFAULT '0' COMMENT '初始魂力';

	");

?>




