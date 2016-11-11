<?php

db_execute($db, "

DROP TABLE IF EXISTS `mission_level_box`;

CREATE TABLE `mission_level_box` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `mission_level_id` int(11) NOT NULL COMMENT '关卡id',
  `order` tinyint(4) NOT NULL COMMENT '品质顺序',
  `award_type` tinyint(4) NOT NULL COMMENT '奖励类型(0--铜钱;1--道具;2--装备)',
  `award_chance` tinyint(4) NOT NULL COMMENT '奖励概率',
  `award_num` int(11) NOT NULL COMMENT '奖励数量',
  `item_id`	 int(11) NOT NULL DEFAULT '0' COMMENT '物品ID(物品奖励填写)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='区域关卡宝箱';


ALTER TABLE `mission_level` modify `award_key` int(11) NOT NULL COMMENT '奖励钥匙数';
ALTER TABLE `mission_level` modify `award_exp` int(11) NOT NULL COMMENT '奖励经验';

");
?>