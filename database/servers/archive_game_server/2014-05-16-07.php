
<?php
db_execute($db, "

CREATE TABLE `sword_soul` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '剑心ID',
  `type_id` int(11) NOT NULL COMMENT '类型ID',
  `name` varchar(10) NOT NULL COMMENT '剑心名称',
  `desc` varchar(20) NOT NULL COMMENT '剑心描述',
  `quality` tinyint(4) NOT NULL COMMENT '品质',
  `fragment_num` smallint(6) DEFAULT NULL COMMENT '碎片数量',
  `kendo_level` tinyint(4) DEFAULT NULL COMMENT '兑换需要剑道等级',
  `only_exchange` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否只能兑换获得',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='剑心';

CREATE TABLE `sword_soul_exchange` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `new_sword_soul_id` smallint(6) NOT NULL COMMENT '制作出来的剑心ID',
  `sword_soul_id1` smallint(6) NOT NULL COMMENT '需求剑心ID1',
  `sword_soul_level1` tinyint(4) NOT NULL COMMENT '需求剑心等级1',
  `num1` tinyint(4) NOT NULL COMMENT '需求数量1',
  `sword_soul_id2` smallint(6) NOT NULL COMMENT '需求剑心ID2',
  `sword_soul_level2` tinyint(4) NOT NULL COMMENT '需求剑心等级2',
  `num2` tinyint(4) NOT NULL COMMENT '需求数量2',
  `sword_soul_id3` smallint(6) NOT NULL COMMENT '需求剑心ID3',
  `sword_soul_level3` tinyint(4) NOT NULL COMMENT '需求剑心等级3',
  `num3` tinyint(4) NOT NULL COMMENT '需求数量3',
  `sword_soul_id4` smallint(6) NOT NULL COMMENT '需求剑心ID4',
  `sword_soul_level4` tinyint(4) NOT NULL COMMENT '需求剑心等级4',
  `num4` tinyint(4) NOT NULL COMMENT '需求数量4',
  `sword_soul_id5` smallint(6) NOT NULL COMMENT '需求剑心ID5',
  `sword_soul_level5` tinyint(4) NOT NULL COMMENT '需求剑心等级5',
  `num5` tinyint(4) NOT NULL COMMENT '需求数量5',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='剑心兑换信息';

CREATE TABLE `sword_soul_level` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '剑心等级ID',
  `sword_soul_id` smallint(6) NOT NULL COMMENT '剑心ID',
  `level` tinyint(4) NOT NULL COMMENT '等级',
  `value` int(11) NOT NULL COMMENT '属性加值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='剑心等级';

CREATE TABLE `sword_soul_quality` (
  `id` smallint(6) NOT NULL COMMENT '剑心等级ID',
  `name` varchar(10) NOT NULL COMMENT '品质名称',
  `sign` varchar(20) NOT NULL COMMENT '程序标示',
  `init_exp` int(11) DEFAULT NULL COMMENT '初始经验',
  `price` int(11) DEFAULT NULL COMMENT '售价',
  `color` varchar(20) DEFAULT NULL COMMENT '颜色',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='剑心品质';

CREATE TABLE `sword_soul_quality_level` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '剑心等级ID',
  `quality_id` smallint(6) NOT NULL COMMENT '品质名称',
  `level` tinyint(4) NOT NULL COMMENT '等级',
  `exp` int(11) NOT NULL COMMENT '升到这一级所需的经验',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='剑心品质';

CREATE TABLE `sword_soul_type` (
  `id` smallint(6) NOT NULL COMMENT '类型ID',
  `name` varchar(10) NOT NULL COMMENT '类型名称',
  `sign` varchar(20) DEFAULT NULL COMMENT '程序标示',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='剑心类型';

");
?>