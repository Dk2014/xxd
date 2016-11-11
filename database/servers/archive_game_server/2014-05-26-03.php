<?php
db_execute($db, 
	"
		ALTER TABLE `item` ADD COLUMN `valid_hours` int(11) NOT NULL DEFAULT '0' COMMENT '有效小时数' AFTER `can_use`;
		ALTER TABLE `item` ADD COLUMN `use_ingot` int(11) NOT NULL DEFAULT '0' COMMENT '使用的元宝价格' AFTER `can_use`;
		ALTER TABLE `item` ADD COLUMN `renew_ingot` int(11) NOT NULL DEFAULT '0' COMMENT '续费的元宝价格' AFTER `can_use`;
		ALTER TABLE `item` ADD COLUMN `func_id` int(11) NOT NULL DEFAULT '0' COMMENT '使用的功能限制' AFTER `can_use`;

		DROP TABLE IF EXISTS `item_box_content`;
		CREATE TABLE `item_box_content` (
  		`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  		`item_id` smallint(6) NOT NULL COMMENT '物品宝箱的ID',
  		`type` tinyint(4) NOT NULL COMMENT '类型，0铜钱，1元宝，2物品',
  		`mode` tinyint(4) NOT NULL COMMENT '随机方式，0直接获得，1概率数量，2概率获得',
 		 `get_item_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '得到的物品ID',
  		`item_id_set` text COMMENT '随机的物品ID集',
  		`item_desc` varchar(50) DEFAULT NULL COMMENT '随机物品集的描述',
  		`min_num` int(11) NOT NULL DEFAULT '0' COMMENT '最少数量',
 		 `max_num` int(11) NOT NULL DEFAULT '0' COMMENT '最多数量',
  		`probability` tinyint(4) NOT NULL DEFAULT '0' COMMENT '概率',
 		 PRIMARY KEY (`id`),
  		 KEY `idx_item_id` (`item_id`)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='宝箱内容';
	"
);

?>