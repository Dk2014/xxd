<?php

$this->AddSQL("
CREATE TABLE `awaken_attr` (
  `id` smallint(6) unsigned NOT NULL AUTO_INCREMENT COMMENT '觉醒属性ID',
  `name` varchar(10) NOT NULL COMMENT '觉醒名称',
  `is_skill` tinyint(4) NOT NULL COMMENT '是否为技能',
  `skill_id` smallint(6) NOT NULL COMMENT '技能ID',
  `type` tinyint(4) NOT NULL COMMENT '属性类型',
  `attr` int(11) NOT NULL COMMENT '加成数值',
  `lights` tinyint(4) NOT NULL COMMENT '希望之光需求量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5098 DEFAULT CHARSET=utf8mb4 COMMENT='觉醒属性';
");

$this->AddSQL("
CREATE TABLE `awaken_graphic` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '觉醒依赖ID',
  `role_id` tinyint(4) unsigned NOT NULL COMMENT '角色ID',
  `impl_id` smallint(6) signed NOT NULL COMMENT '实例ID',
  `attr_id` smallint(6) NOT NULL COMMENT '觉醒属性',
  `dep_impl` smallint(6) signed NOT NULL COMMENT '依赖实例ID',
  `direct` tinyint(4) NOT NULL COMMENT '派生方向，0西，1北，2东，3南',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5098 DEFAULT CHARSET=utf8mb4 COMMENT='觉醒图谱';
");

$this->AddSQL("
CREATE TABLE `player_awaken_graphic` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `role_id` tinyint(4) NOT NULL COMMENT '角色ID',
  `attr_impl` smallint(6) NOT NULL COMMENT '觉醒属性实例ID',
  `level` tinyint(4) NOT NULL DEFAULT '0' COMMENT '觉醒等级',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家觉醒图谱';
");
