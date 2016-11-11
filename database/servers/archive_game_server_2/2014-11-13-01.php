<?php

$this->AddSQL("
CREATE TABLE IF NOT EXISTS `global_group_buy_status`(
		`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID标识',
		`cid` int(11) NOT NULL COMMENT '外键，指定对应得团购物品记录id',
		`status` int(11) NOT NULL DEFAULT 0 COMMENT '当前团购状态，即购买总数',
		PRIMARY KEY(`id`)
	)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT 'hd服务器记录团购状态信息';
");

$this->AddSQL("
	CREATE TABLE IF NOT EXISTS `events_group_buy`(
		`id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT 'ID标识',
		`item_id` smallint(6) NOT NULL COMMENT '参与团购得物品id',
		`base_price` mediumint(6) NOT NULL COMMENT '团购物品底价',
		`buy_times1` smallint(6) NOT NULL COMMENT '购买次数1',
		`buy_percent1` float(3,2) NOT NULL COMMENT '购买折扣1',
		`buy_times2` smallint(6) NOT NULL COMMENT '购买次数2',
		`buy_percent2` float(3,2) NOT NULL COMMENT '购买折扣2',
		`buy_times3` smallint(6) NOT NULL COMMENT '购买次数3',
		`buy_percent3` float(3,2) NOT NULL COMMENT '购买折扣3',
		`buy_times4` smallint(6) NOT NULL COMMENT '购买次数4',
		`buy_percent4` float(3,2) NOT NULL COMMENT '购买折扣4',
		`buy_times5` smallint(6) NOT NULL COMMENT '购买次数5',
		`buy_percent5` float(3,2) NOT NULL COMMENT '购买折扣5',
		`buy_times6` smallint(6) NOT NULL COMMENT '购买次数6',
		`buy_percent6` float(3,2) NOT NULL COMMENT '购买折扣6',
		PRIMARY KEY(`id`)
		)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '团购内容';
	");
?>