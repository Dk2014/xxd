<?php

$this->AddSQL("

alter table `item` add `act_id` int(11) NOT NULL DEFAULT '0' COMMENT '使用触发功能';

CREATE TABLE `item_reflect_config` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `item_id` smallint(6) NOT NULL COMMENT '物品宝箱的ID',
  `award_coin_min` int(11) NOT NULL DEFAULT '0' COMMENT '最少奖励数量',
  `award_coin_max` int(11) NOT NULL DEFAULT '0' COMMENT '最多奖励数量',
  PRIMARY KEY (`id`),
  KEY `idx_item_id` (`item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='宝箱内容';

INSERT INTO `item` (`id`, `type_id`, `quality`, `name`, `level`, `desc`, `price`, `sign`, `can_use`, `panel`, `func_id`, `renew_ingot`, `use_ingot`, `valid_hours`, `equip_type_id`, `health`, `speed`, `attack`, `defence`, `show_mode`, `equip_role_id`, `appendix_num`, `appendix_level`, `music_sign`, `can_batch`, `refine_base`, `show_origin`, `act_id`)
VALUES
	(681,21,3,'聚宝盆',NULL,'每天可使用一次，每次可获得随机数量的铜钱，感谢各位仙尊对仙侠道的支持！',0,'JuBaoPen',1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL,0,0,0,1);

INSERT INTO `item_reflect_config` (`id`, `item_id`, `award_coin_min`, `award_coin_max`)
VALUES
	(81,681,1000,10000);

INSERT INTO `item_type` (`id`, `name`, `max_num_in_pos`, `sign`, `order`)
VALUES
	(21,'秘宝',1,'ActReflect',19);

");

