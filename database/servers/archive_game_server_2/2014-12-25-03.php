<?php
$this->AddSQL("
create table `npc_role` (
	`id` int(11) NOT NULL AUTO_INCREMENT  COMMENT 'ID',
	`name` varchar(10) NOT NULL COMMENT '角色名',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='NPC角色';

alter table town_npc add column `npc_role` int(11) NOT NULL DEFAULT '0' COMMENT '抽象NPC角色';

");
?>
