<?php 

$this->AddSQL("
create table `main_role_cooperation` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT 'ID',
   `buddy_num` tinyint(4) NOT NULL COMMENT '要求伙伴人数',
   `health` int(11) NOT NULL DEFAULT '0' COMMENT '生命 - health',
   `attack` int(11) NOT NULL DEFAULT '0' COMMENT '普攻 - attack',
   `defence` int(11) NOT NULL DEFAULT '0'  COMMENT '普防 - defence',
   `speed` int(11) NOT NULL DEFAULT '0' COMMENT '速度',
   `cultivation` int(11) DEFAULT '0' COMMENT '内力',
  PRIMARY KEY (`id`),
  UNIQUE KEY `buddy_num` (`buddy_num`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='伙伴协力';
");

?>




