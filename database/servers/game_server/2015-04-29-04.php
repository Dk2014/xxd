<?php
$this->AddSQL("
alter table `global_clique` add column `auto_audit` tinyint(4) not null default '0' comment '自动审核';
alter table `global_clique` add column `auto_audit_level` smallint(6) not null default '0' comment '自动审核等级下限';
");

$this->AddSQL("
CREATE TABLE `clique_center_building_level_info` (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`level` smallint(6) NOT NULL COMMENT '等级',
	`max_member` smallint(6) NOT NULL COMMENT '成员数量',
	`levelup_coins` bigint(20) NOT NULL COMMENT '升级消耗金钱',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
");

?>
