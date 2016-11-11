<?php 
db_execute($db, "
create table `battle_pet_soul_exchange` (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
	`quality` tinyint(4) NOT NULL COMMENT '品质',
	`star` tinyint(4) NOT NULL DEFAULT '1' COMMENT '星级 [1,5]',
	`soul_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '获得碎片个数',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='灵宠碎片兑换'; 

");
?>
