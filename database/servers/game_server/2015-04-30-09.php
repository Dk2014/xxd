<?php
$this->AddSQL("
create table `player_clique_kongfu_attrib` (
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`health` int(11) NOT NULL DEFAULT '0' comment '生命',
	`attack` int(11) NOT NULL DEFAULT '0' comment '攻击',
	`defence` int(11) NOT NULL DEFAULT '0' comment '防御',
	PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家帮派武功属性加成';
");
?>
