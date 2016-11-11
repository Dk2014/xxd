<?php
$this->AddSQL("
CREATE TABLE IF NOT EXISTS `clique_building_kongfu`(
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`building` int(11) NOT NULL  COMMENT '所属建筑等级',
	`level` smallint(6) NOT NULL DEFAULT 0 COMMENT '等级',
	 `upgrade_coins` bigint(20) NOT NULL DEFAULT 0 COMMENT '升级所需铜钱',
	 primary key(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='帮派武学建筑';
");

$this->AddSQL("
CREATE TABLE IF NOT EXISTS `clique_kongfu`(
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '标识ID',
	`building` int(11) NOT NULL  COMMENT '所属建筑等级',
	`name` varchar(20) NOT NULL DEFAULT '' COMMENT '武功名称',
	`require_building_level` smallint(6) NOT NULL  COMMENT '要求建筑等级',
	primary key(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='帮派武功'; 
");

$this->AddSQL("
CREATE TABLE IF NOT EXISTS `clique_kongfu_level`(
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
	`kongfu_id` int(11) NOT NULL  COMMENT '心法ID',
	`consume_contrib` bigint(20) NOT NULL  DEFAULT '0' COMMENT '消耗帮派贡献',
	`level` smallint(6) NOT NULL DEFAULT 0 COMMENT '等级',
	`attrib_type` tinyint(4) NOT NULL  COMMENT '属性类型',
	`attrib_value` int(11) NOT NULL  DEFAULT 0 COMMENT '属性值',
	primary key(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='帮派武功等级信息'; 
");
?>
