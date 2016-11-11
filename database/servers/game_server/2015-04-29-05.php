<?php

$this->AddSQL("
ALTER TABLE `clique_center_building_level_info` ADD COLUMN `daily_max_coins` bigint(20) NOT NULL DEFAULT 0 COMMENT '每日贡献上限';

-- 帮派建筑物表
CREATE TABLE IF NOT EXISTS `clique_building`(
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '标示ID',
	`name` varchar(30) NOT NULL COMMENT '建筑名称',
	`biaoshi` varchar(30) NOT NULL COMMENT '建筑标识',
	`desc` text COMMENT '建筑描述',
	`order` smallint(4) NOT NULL DEFAULT 0 COMMENT '优先权排序',
	primary key(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='帮派建筑物';

-- 帮派建筑物 钱庄
CREATE TABLE IF NOT EXISTS `clique_building_bank`(
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '标识ID',
	`level` smallint(6) NOT NULL DEFAULT 0 COMMENT '钱庄等级',
	`sliver_coupon_coins` int(11) NOT NULL DEFAULT 0 COMMENT '银劵底价',
	 `sliver_coupon_num` smallint(6) NOT NULL DEFAULT 0 COMMENT '银劵限购',
	 `sliver_coupon_sold` int(11) NOT NULL DEFAULT 0 COMMENT '银劵售价',
	 `gold_coupon_coins` int(11) NOT NULL DEFAULT 0 COMMENT '金劵底价',
	 `gold_coupon_num` smallint(6) NOT NULL DEFAULT 0 COMMENT '金劵限购',
	 `gold_coupon_sold` int(11) NOT NULL DEFAULT 0 COMMENT '金劵售价',
	 `upgrade` bigint(20) NOT NULL DEFAULT 0 COMMENT '钱庄升级所需铜钱',
	 primary key(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='帮派建筑物钱庄';
");
?>