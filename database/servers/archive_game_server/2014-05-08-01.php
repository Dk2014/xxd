<?php
db_execute($db, "

DROP TABLE IF EXISTS `role_level`;

CREATE TABLE `role_level` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '角色等级ID',
  `role_id` tinyint(4) NOT NULL COMMENT '角色ID',
  `level` int(11) NOT NULL COMMENT '等级 - level',
  `exp` bigint(20) NOT NULL COMMENT '升到下一级所需经验',
  `health` int(11) NOT NULL COMMENT '生命 - health',
  `attack` int(11) NOT NULL COMMENT '普攻 - attack',
  `defence` int(11) NOT NULL COMMENT '普防 - defence',
  `cultivation` int(11) NOT NULL COMMENT '内力 - cultivation',
  `speed` int(11) NOT NULL COMMENT '速度 - speed',
  `critial` int(11) NOT NULL COMMENT '暴击 - critial',
  `block` int(11) NOT NULL COMMENT '格挡 - block',
  `hit` int(11) NOT NULL COMMENT '命中 - hit',
  `dodge` int(11) NOT NULL COMMENT '闪避 - dodge',
  `critial_hurt` int(11) NOT NULL COMMENT '暴击伤害,必杀 – critial hurt',
  `will` int(11) NOT NULL DEFAULT '0' COMMENT '意志',
  `max_power` int(11) NOT NULL DEFAULT '0' COMMENT '精气上限',
  `init_power` int(11) NOT NULL DEFAULT '0' COMMENT '初始精气',
  `sunder_max_value` int(11) NOT NULL DEFAULT '0' COMMENT '护甲值',
  `sunder_hurt_rate` int(11) NOT NULL DEFAULT '0' COMMENT '破甲前起始的伤害转换率（百分比）',
  `sunder_end_hurt_rate` int(11) NOT NULL DEFAULT '0' COMMENT '破甲后的伤害转换率（百分比）',
  `sunder_round_num` tinyint(4) NOT NULL DEFAULT '0' COMMENT '破甲持续回合数',
  `sunder_dizziness` tinyint(4) NOT NULL DEFAULT '0' COMMENT '破甲后眩晕回合数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='角色等级';

");
?>
