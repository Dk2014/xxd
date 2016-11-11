
<?php
db_execute($db, 
"

DROP TABLE IF EXISTS `equipment_decompose`;
CREATE TABLE `equipment_decompose` (
  `id` tinyint(4) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `level` int(11) NOT NULL COMMENT '等级下限',
  `fragment_num` smallint(6) NOT NULL COMMENT '获得部位碎片数量',
  `crystal_num` smallint(6) NOT NULL COMMENT '获得结晶数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='装备分解获得材料';

DROP TABLE IF EXISTS `equipment_recast`;
CREATE TABLE `equipment_recast` (
  `id` tinyint(4) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `level` int(11) NOT NULL COMMENT '等级下限',
  `quality` tinyint(4) NOT NULL COMMENT '品质',
  `fragment_num` smallint(6) NOT NULL COMMENT '需要部位碎片数量',
  `blue_crystal_num` smallint(6) NOT NULL COMMENT '需要蓝色结晶数量',
  `purple_crystal_num` smallint(6) NOT NULL COMMENT '需要紫色结晶数量',
  `golden_crystal_num` smallint(6) NOT NULL COMMENT '需要金色结晶数量',
  `orange_crystal_num` smallint(6) NOT NULL COMMENT '需要橙色结晶数量',
  `consume_coins` bigint(20) NOT NULL COMMENT '消耗的铜钱',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='装备重铸消耗';

DROP TABLE IF EXISTS `equipment_refine`;
CREATE TABLE `equipment_refine` (
  `id` tinyint(4) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `level` int(11) NOT NULL COMMENT '等级下限',
  `quality` tinyint(4) NOT NULL COMMENT '品质',
  `fragment_num` smallint(6) NOT NULL COMMENT '需要部位碎片数量',
  `blue_crystal_num` smallint(6) NOT NULL COMMENT '需要蓝色结晶数量',
  `purple_crystal_num` smallint(6) NOT NULL COMMENT '需要紫色结晶数量',
  `golden_crystal_num` smallint(6) NOT NULL COMMENT '需要金色结晶数量',
  `orange_crystal_num` smallint(6) NOT NULL COMMENT '需要橙色结晶数量',
  `level1_consume_coins` bigint(20) NOT NULL COMMENT '重铸到1级消耗的铜钱',
  `level2_consume_coins` bigint(20) NOT NULL COMMENT '重铸到2级消耗的铜钱',
  `level3_consume_coins` bigint(20) NOT NULL COMMENT '重铸到3级消耗的铜钱',
  `level4_consume_coins` bigint(20) NOT NULL COMMENT '重铸到4级消耗的铜钱',
  `level5_consume_coins` bigint(20) NOT NULL COMMENT '重铸到5级消耗的铜钱',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='装备精练消耗';

DROP TABLE IF EXISTS `equipment_refine_level`;
CREATE TABLE `equipment_refine_level` (
  `id` tinyint(4) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `quality` tinyint(4) NOT NULL COMMENT '品质',
  `level` int(11) NOT NULL COMMENT '等级下限',
  `probability` tinyint(4) NOT NULL COMMENT '精练成功概率',
  `gain_pct` tinyint(4) NOT NULL COMMENT '增益百分比',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='精练武器对应概率与加成';

ALTER TABLE `player_item_appendix`
ADD COLUMN `refine_level`  tinyint(4) NULL DEFAULT 0 COMMENT '精练等级' AFTER `destroy_level`,
ADD COLUMN `recast_attr`  smallint(6) NULL DEFAULT 0 COMMENT '重铸属性' AFTER `refine_level`;

"
);

?>

