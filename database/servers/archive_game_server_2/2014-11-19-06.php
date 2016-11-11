<?php
$this->AddSQL("
create table `ghost_upgrade_price` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',	
  `quality` tinyint(4) NOT NULL COMMENT '品质',	
  `cost` int(11) NOT NULL COMMENT '碎片单价',
PRIMARY KEY(`id`),
UNIQUE KEY `quality` (`quality`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='魂侍升级价格';
");

$this-> AddSQL("
	alter table `player_ghost_state` drop column `upgrade_by_ingot_num`;
");

$this->AddSQL("
	alter table `player_ghost_state` drop column `upgrade_by_ingot_time`;
");

?>
