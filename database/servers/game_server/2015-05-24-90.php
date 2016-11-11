<?php 
$this->AddSQL("
CREATE TABLE if not exists `player_any_sdk_order` (
	`id` bigint(20) NOT NULL COMMENT 'ID',
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`order_id` bigint(20) NOT NULL COMMENT '内部订单ID',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='anysdk 订单处理纪录';
");
?>
