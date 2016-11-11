<?php 
$this->AddSQL("
CREATE TABLE product_info (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
	`app_store_product_id` varchar(32) NOT NULL COMMENT '产品ID',
	`google_play_product_id` varchar(32) NOT NULL COMMENT '产品ID',
	`price` int(11) NOT NULL COMMENT '价格',
	`ingot` int(11) NOT NULL COMMENT '获得元宝',
	`unit` varchar(4) NOT NULL COMMENT '货币单位',
	`is_month_card` tinyint(4) NOT NULL DEFAULT '0' COMMENT '月卡',
	`first_presentation` int(11) NOT NULL COMMENT '首充赠送元宝',
	`regular_presentation` int(11) NOT NULL COMMENT '常规赠送元宝',
PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='内购产品信息';
"
);

$this->AddSQL("
	CREATE TABLE `player_iap_order` (
		`id` bigint(20) NOT NULL COMMENT 'ID',
		`pid` bigint(20) NOT NULL COMMENT '玩家ID',
		`source` tinyint(4) NOT NULL COMMENT '订单来源 0--appstore 1--google play 3--wegames 平台',
		`order_id` bigint(20) NOT NULL COMMENT '内部订单ID',
		`platform_order_id` bigint(20) NOT NULL COMMENT '渠道订单ID',
		PRIMARY KEY (`id`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='anysdk 订单处理纪录(source+order_id组成键)';
");
?>
