<?php
$this->AddSQL("
create table `battle_grid_upgrade_price` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',	
  `times` int(11) NOT NULL COMMENT '升级次数',	
  `cost` int(11) NOT NULL COMMENT '一个碎片单价',
PRIMARY KEY(`id`),
UNIQUE KEY `times` (`times`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='灵宠格子元宝升级价格梯度';
");

$this->AddSQL("
CREATE TABLE `player_battle_pet_state` (
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`upgrade_by_ingot_num` int(11) NOT NULL DEFAULT '0' COMMENT '今日使用元宝升级次数',
	`upgrade_by_ingot_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最近一次使用元宝升级时间',
PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家灵宠状态';
");


?>
