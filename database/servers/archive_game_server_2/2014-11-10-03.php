<?php

$this->AddSQL("
CREATE TABLE IF NOT EXISTS `events_richman_club_awards` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `require_vip_level` smallint(6) NOT NULL COMMENT '所需的vip等级',
  `require_vip_count` smallint(6) NOT NULL COMMENT '所需的vip相应人数',
  `award_vip_level1` smallint(6) DEFAULT '0' COMMENT '能领奖的vip等级1',
  `award_vip_item1_id` smallint(6) DEFAULT '0' COMMENT '能领奖的vip的奖励物品1 ID',
  `award_vip_item1_num` smallint(6) DEFAULT '1' COMMENT '能领奖的vip的奖励物品1数量 默认为1',
  `award_vip_level2` smallint(6) DEFAULT '0' COMMENT '能领奖的vip等级2',
  `award_vip_item2_id` smallint(6) DEFAULT '0' COMMENT '能领奖的vip的奖励物品2 ID',
  `award_vip_item2_num` smallint(6) DEFAULT '1' COMMENT '能领奖的vip的奖励物品2数量 默认为1',
  `award_vip_level3` smallint(6) DEFAULT '0' COMMENT '能领奖的vip等级3',
  `award_vip_item3_id` smallint(6) DEFAULT '0' COMMENT '能领奖的vip的奖励物品3 ID',
  `award_vip_item3_num` smallint(6) DEFAULT '1' COMMENT '能领奖的vip的奖励物品3数量 默认为1',
  `award_vip_level4` smallint(6) DEFAULT '0' COMMENT '能领奖的vip等级4',
  `award_vip_item4_id` smallint(6) DEFAULT '0' COMMENT '能领奖的vip的奖励物品4 ID',
  `award_vip_item4_num` smallint(6) DEFAULT '1' COMMENT '能领奖的vip的奖励物品4数量 默认为1',
  `award_vip_level5` smallint(6) DEFAULT '0' COMMENT '能领奖的vip等级5',
  `award_vip_item5_id` smallint(6) DEFAULT '0' COMMENT '能领奖的vip的奖励物品5 ID',
  `award_vip_item5_num` smallint(6) DEFAULT '1' COMMENT '能领奖的vip的奖励物品5数量 默认为1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='土豪俱乐部运营活动';
");

?>