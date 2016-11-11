<?php
db_execute($db, "

CREATE TABLE `player_ghost` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `ghost_id` smallint(6) NOT NULL COMMENT '魂侍ID',
  `pos` smallint(6) NOT NULL COMMENT '位置',
  `level` smallint(6) NOT NULL DEFAULT '1' COMMENT '魂侍等级',
  `exp` bigint(20) NOT NULL DEFAULT '0' COMMENT '魂侍经验',
  `is_new` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否新魂侍',
  `dge_level` int(11) NOT NULL DEFAULT '0' COMMENT '闪避等级',
  `cri_level` int(11) NOT NULL DEFAULT '0' COMMENT '暴击等级',
  `blk_level` int(11) NOT NULL DEFAULT '0' COMMENT '格挡等级',
  `hit_level` int(11) NOT NULL DEFAULT '0' COMMENT '命中等级',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家魂侍表';


");
?>