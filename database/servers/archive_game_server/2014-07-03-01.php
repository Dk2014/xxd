
<?php
db_execute($db, "

DROP TABLE IF EXISTS  `ghost_umbra`;
CREATE TABLE `ghost_umbra` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `town_id` tinyint(4) NOT NULL DEFAULT '0' COMMENT '城镇id',
  `key` int(11) NOT NULL DEFAULT '0' COMMENT '进入影界需求权值',
  `ghost_probability` smallint(6) NOT NULL DEFAULT '0' COMMENT '抽中魂侍的概率',
  `fragment_probability` smallint(6) NOT NULL DEFAULT '0' COMMENT '抽中碎片的概率',
  `fruit_probability` smallint(6) NOT NULL DEFAULT '0' COMMENT '抽中果实的概率',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='魂侍影界表';


DROP TABLE IF EXISTS `player_ghost_umbra`;
CREATE TABLE `player_ghost_umbra` (
  `id` bigint(20) NOT NULL COMMENT '主键id',
  `pid` bigint(20) NOT NULL COMMENT '玩家id',
  `umbra_id` smallint(6) NOT NULL COMMENT '影界id',
  `num` smallint(6) NOT NULL COMMENT '今日剩余次数',
  `last_draw_at` bigint(20) NOT NULL COMMENT '上次开启时间',
  `refresh_num` smallint(6) NOT NULL COMMENT '今日刷新次数',
  `last_refresh_at` bigint(20) NOT NULL COMMENT '上次刷新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家魂侍副本表';


DROP TABLE IF EXISTS  `ghost_umbra_vip`;
CREATE TABLE `ghost_umbra_vip` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `vip_level` tinyint(4) NOT NULL DEFAULT '0' COMMENT 'vip 等级',
  `refresh_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '刷新次数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='魂侍影界VIP表';

DROP TABLE IF EXISTS  `ghost_umbra_vip_price`;
CREATE TABLE `ghost_umbra_vip_price` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `refresh_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '刷新次数',
  `ingot` bigint(20) NOT NULL DEFAULT '0' COMMENT '元宝',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='魂侍影界VIP价格表';

");
?>
