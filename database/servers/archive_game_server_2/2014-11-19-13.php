<?php
$this->AddSQL("
CREATE TABLE IF NOT EXISTS `events_buy_partner` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `patner_id` smallint(6) NOT NULL COMMENT '伙伴ID',
  `buddy_level` smallint(6) NOT NULL DEFAULT '1' COMMENT '伙伴等级',
  `cost` bigint(20) NOT NULL COMMENT '价格',
  `skill_id1` smallint(6) NOT NULL COMMENT '招式名称1',
  `skill_id2` smallint(6) NOT NULL COMMENT '招式名称2',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='限时购买伙伴';
");