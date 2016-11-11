<?php
db_execute($db, 
	"
  DROP TABLE IF EXISTS `role_realm_level`;
	CREATE TABLE `role_realm_level` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `realm_level` smallint(6) NOT NULL DEFAULT '0' COMMENT '境界等级',
  `exp` bigint(20) NOT NULL DEFAULT '0' COMMENT '升级所需经验',
  `need_realm_class` bigint(20) NOT NULL DEFAULT '0' COMMENT '升级所需阶级',
  `item_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '道具数量',
  `add_health` int(11) NOT NULL DEFAULT '0' COMMENT '增加生命',
  `add_attack` int(11) NOT NULL DEFAULT '0' COMMENT '增加攻击',
  `add_defence` int(11) NOT NULL DEFAULT '0' COMMENT '增加防御',
  `add_cultivation` int(11) NOT NULL DEFAULT '0' COMMENT '增加内力',
  `add_speed` int(11) NOT NULL DEFAULT '0' COMMENT '增加速度',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色境界等级表';
  
  DROP TABLE IF EXISTS `role_realm_class`;
  CREATE TABLE `role_realm_class` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '阶级名称',
  `realm_class` smallint(6) NOT NULL DEFAULT '0' COMMENT '境界阶级',
  `need_realm_level` smallint(6) NOT NULL DEFAULT '0' COMMENT '升级所需境界等级',
  `item_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '道具id',
  `item_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '道具数量',
  `add_hth` int(11) NOT NULL DEFAULT '0' COMMENT '生命加成',
  `add_sunder_value` int(11) NOT NULL DEFAULT '0' COMMENT '护甲加成',
  `add_power` int(11) NOT NULL DEFAULT '0' COMMENT '初始精气加成',
  `add_max_power` int(11) NOT NULL DEFAULT '0' COMMENT '精气上限加成',
  `add_aoe_reduce` int(11) NOT NULL DEFAULT '0' COMMENT '范围免伤加成',
  `add_critial_level` int(11) NOT NULL DEFAULT '0' COMMENT '暴击等级加成',
  `add_dodge_level` int(11) NOT NULL DEFAULT '0' COMMENT '闪避等级加成',
  `add_hit_level` int(11) NOT NULL DEFAULT '0' COMMENT '命中等级加成',
  `add_block_level` int(11) NOT NULL DEFAULT '0' COMMENT '格挡等级加成',
  `add_tenacity_level` int(11) NOT NULL DEFAULT '0' COMMENT '韧性等级加成',
  `add_destroy_level` int(11) NOT NULL DEFAULT '0' COMMENT '破击等级加成',
  `add_sunder_min_hurt_rate` int(11) NOT NULL DEFAULT '0' COMMENT '破甲前免伤',
  `desc` varchar(500) NOT NULL DEFAULT '' COMMENT '简介',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色境界阶级表';

  DROP TABLE IF EXISTS `role_realm_skill`;
  CREATE TABLE `role_realm_skill` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `role_id` tinyint(4) NOT NULL COMMENT '角色ID',
  `realm_class_id` bigint(20) NOT NULL COMMENT '境界阶级ID',
  `skill_id1` smallint(6) NOT NULL COMMENT '境界技能1',
  `skill_id2` smallint(6) NOT NULL COMMENT '境界技能2',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色境界技能表';

	"
);
?>