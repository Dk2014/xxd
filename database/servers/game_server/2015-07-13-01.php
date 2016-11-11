<?php

$this->AddSQL("

CREATE TABLE IF NOT EXISTS `enemy_passive_skill`(
	`id` int(10) unsigned NOT NULL COMMENT '怪物ID',
	`config` text COMMENT '技能配置',
	PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='怪物属性加成';
");
?>
