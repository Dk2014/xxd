<?php
$this->AddSQL("
CREATE TABLE `random_award_box` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `mission_level_id` int(11) NOT NULL COMMENT '关卡id',
    `order` tinyint(4) NOT NULL COMMENT '品质顺序',
  `award_type` tinyint(4) NOT NULL COMMENT '奖励类型(0--铜钱;1--道具;2--装备; 3--经验;4--经验倍数; 5--铜钱倍数,6--契约球)',
  `award_chance` tinyint(4) NOT NULL COMMENT '奖励概率',
  `award_num` int(11) NOT NULL COMMENT '奖励数量',
  `item_id` int(11) NOT NULL DEFAULT '0' COMMENT '物品ID(物品奖励填写)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='随机奖品宝箱'; 
");




