<?php
db_execute($db, "

DROP TABLE IF EXISTS `ghost`;

	
CREATE TABLE `ghost` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '魂侍名称',
  `sign` varchar(30) NOT NULL DEFAULT '' COMMENT '图标标识',
  `town_id` tinyint(4) NOT NULL DEFAULT '0' COMMENT '城镇id',
  `unique_key` smallint(6) NOT NULL DEFAULT '0' COMMENT '获得魂侍的信息',
  `init_level` smallint(6) NOT NULL DEFAULT '1' COMMENT '魂侍初始等级',
  `init_exp` bigint(20) NOT NULL DEFAULT '0' COMMENT '魂侍初始经验',
  `potential` tinyint(4) DEFAULT '0' COMMENT '魂侍潜力点',
  `hth` int(11) NOT NULL DEFAULT '0' COMMENT '生命',
  `atk` int(11) NOT NULL DEFAULT '0' COMMENT '攻击',
  `def` int(11) NOT NULL DEFAULT '0' COMMENT '防御',
  `spd` int(11) NOT NULL DEFAULT '0' COMMENT '速度',
  `hit_level` int(11) NOT NULL DEFAULT '0' COMMENT '命中等级',
  `dge_level` int(11) NOT NULL DEFAULT '0' COMMENT '闪避等级',
  `cri_level` int(11) NOT NULL DEFAULT '0' COMMENT '暴击等级',
  `blk_level` int(11) NOT NULL DEFAULT '0' COMMENT '格挡等级',
  `desc` varchar(300) DEFAULT NULL COMMENT '描述',
  `quality` tinyint(4) NOT NULL DEFAULT '0' COMMENT '品质',
  `max_hit_level` int(11) NOT NULL DEFAULT '0' COMMENT '最大命中等级',
  `max_dge_level` int(11) NOT NULL DEFAULT '0' COMMENT '最大闪避等级',
  `max_cri_level` int(11) NOT NULL DEFAULT '0' COMMENT '最大暴击等级',
  `max_blk_level` int(11) NOT NULL DEFAULT '0' COMMENT '最大格挡等级',
  `can_mission_get` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否可以在影界获得(1为可以)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='魂侍主表';

");
?>