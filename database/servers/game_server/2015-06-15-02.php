<?php
$this->AddSQL("
	CREATE TABLE `global_payments_rule` (
	  `id` bigint(20) NOT NULL COMMENT '自增主键',
	  `rule` varchar(1024) NOT NULL COMMENT '规则',
	  `begin_time` bigint(20) NOT NULL COMMENT '开启时间',
	  `end_time` bigint(20) NOT NULL COMMENT '结束时间',
	  PRIMARY KEY (`id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='充值返利规则';

	ALTER TABLE `player_any_sdk_order` ADD COLUMN `present` bigint(20) NOT NULL DEFAULT '0' COMMENT '赠送元宝元宝';
");
?>