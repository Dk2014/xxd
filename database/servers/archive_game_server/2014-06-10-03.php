<?php

db_execute($db, "

ALTER TABLE `player_item` ADD COLUMN `appendix_id` bigint(20) DEFAULT '0' COMMENT '附加属性ID';

DROP TABLE IF EXISTS `player_item_appendix`;
CREATE TABLE `player_item_appendix` (
  `id` bigint(20) NOT NULL  COMMENT 'ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `health` int(11) DEFAULT '0' COMMENT '生命',
  `cultivation` int(11) DEFAULT '0' COMMENT '内力',
  `speed` int(11) DEFAULT '0' COMMENT '速度',
  `attack` int(11) DEFAULT '0' COMMENT '攻击',
  `defence` int(11) DEFAULT '0' COMMENT '防御',
  `dodge_level` int(11) DEFAULT '0' COMMENT '闪避',
  `hit_level` int(11) DEFAULT '0' COMMENT '命中',
  `block_level` int(11) DEFAULT '0' COMMENT '格挡',
  `critical_level` int(11) DEFAULT '0' COMMENT '暴击',
  `tenacity_level` int(11) DEFAULT '0' COMMENT '韧性',
  `destroy_level` int(11) DEFAULT '0' COMMENT '破击',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家装备追加属性表';

");
?>