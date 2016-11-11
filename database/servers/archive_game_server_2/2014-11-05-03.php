<?php

$this->AddSQL("

CREATE TABLE `vip_levelup_gift` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `vip_level` smallint(6) NOT NULL COMMENT '要求vip等级',
  `item_id` smallint(6) NOT NULL COMMENT '物品ID',
  `item_num` smallint(6) NOT NULL COMMENT '物品数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='vip升级奖励';

");

?>