<?php
$this->AddSQL("drop table `global_payments_rule`");
$this->AddSQL("
	CREATE TABLE `payments_rule` (
	  `id` bigint(20) NOT NULL COMMENT '自增主键',
	  `rule` varchar(1024) NOT NULL COMMENT '规则',
	  `begin_time` bigint(20) NOT NULL COMMENT '开启时间',
	  `end_time` bigint(20) NOT NULL COMMENT '结束时间',
	  PRIMARY KEY (`id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='充值返利规则';
");
$this->AddSQL("
	insert into `payments_rule` (`id`,`rule`,`begin_time`,`end_time`) values (1,'',0,0);
");
?>