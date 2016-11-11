<?php /* dump file */

	db_execute($db, <<<THESQL1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table enemy_role
# ------------------------------------------------------------

DROP TABLE IF EXISTS `enemy_role`;

CREATE TABLE `enemy_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `name` varchar(10) NOT NULL COMMENT '角色名称',
  `sign` varchar(20) NOT NULL DEFAULT '' COMMENT '资源标识',
  `level` int(11) NOT NULL COMMENT '等级 - level',
  `health` int(11) NOT NULL COMMENT '生命 - health',
  `cultivation` int(11) NOT NULL COMMENT '内力 - cultivation',
  `speed` int(11) NOT NULL COMMENT '速度 - speed',
  `attack` int(11) NOT NULL COMMENT '普攻 - attack',
  `defence` int(11) NOT NULL COMMENT '普防 - defence',
  `dodge` int(11) NOT NULL COMMENT '闪避 - dodge',
  `hit` int(11) NOT NULL COMMENT '命中 - hit',
  `block` int(11) NOT NULL COMMENT '格挡 - block',
  `critial` int(11) NOT NULL COMMENT '暴击 - critial',
  `toughness` int(11) NOT NULL COMMENT '韧性',
  `destroy` int(11) NOT NULL COMMENT '破击',
  `critial_hurt` int(11) NOT NULL COMMENT '必杀 – critial hurt',
  `will` int(11) NOT NULL COMMENT '意志',
  `aoe_reduce` int(11) NOT NULL COMMENT '范围免伤',
  `skill_id` smallint(6) NOT NULL COMMENT '绝招ID',
  `skill_force` int(11) NOT NULL COMMENT '绝招威力',
  `sunder_max_value` int(11) NOT NULL COMMENT '护甲值',
  `sunder_min_hurt_rate` int(11) NOT NULL COMMENT '破甲前起始的伤害转换率（百分比）',
  `sunder_end_hurt_rate` int(11) NOT NULL COMMENT '破甲后的伤害转换率（百分比）',
  `sunder_attack` int(11) NOT NULL COMMENT '攻击破甲值',
  `skill_id2` smallint(6) NOT NULL COMMENT '第二绝招ID',
  `skill_force2` int(11) NOT NULL COMMENT '第二绝招威力',
  `skill_wait` tinyint(4) NOT NULL COMMENT '绝招蓄力回合',
  `release_num` tinyint(4) NOT NULL COMMENT '释放次数',
  `recover_round_num` tinyint(4) NOT NULL COMMENT '恢复回合数',
  `common_attack_num` tinyint(4) NOT NULL COMMENT '入场普通攻击次数',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='敌人角色数据';



# Dump of table mission
# ------------------------------------------------------------

DROP TABLE IF EXISTS `mission`;

CREATE TABLE `mission` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '副本ID',
  `town_id` smallint(6) NOT NULL COMMENT '城镇ID',
  `type` tinyint(4) NOT NULL COMMENT '副本类型',
  `keys` int(11) NOT NULL COMMENT '开启钥匙数',
  `name` varchar(10) NOT NULL DEFAULT '' COMMENT '副本名称',
  `sign` varchar(50) NOT NULL COMMENT '资源标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='副本';



# Dump of table mission_enemy
# ------------------------------------------------------------

DROP TABLE IF EXISTS `mission_enemy`;

CREATE TABLE `mission_enemy` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `mission_level_id` int(11) NOT NULL COMMENT '副本关卡id',
  `x0` int(11) NOT NULL COMMENT 'x0轴坐标',
  `y0` int(11) NOT NULL COMMENT 'y0轴坐标',
  `monster0_id` int(11) NOT NULL COMMENT '怪物1 ID',
  `monster0_chance` tinyint(4) NOT NULL COMMENT '怪物出现概率',
  `x1` int(11) NOT NULL COMMENT 'x1轴坐标',
  `y1` int(11) NOT NULL COMMENT 'y1轴坐标',
  `monster1_id` int(11) NOT NULL COMMENT '怪物2 ID',
  `monster1_chance` tinyint(4) NOT NULL COMMENT '怪物出现概率',
  `x2` int(11) NOT NULL COMMENT 'x2轴坐标',
  `y2` int(11) NOT NULL COMMENT 'y2轴坐标',
  `monster2_id` int(11) NOT NULL COMMENT '怪物2 ID',
  `monster2_chance` tinyint(4) NOT NULL COMMENT '怪物出现概率',
  `boss_id` int(11) NOT NULL COMMENT 'boss id',
  `order` tinyint(4) NOT NULL COMMENT '顺序',
  `talk` varchar(200) NOT NULL DEFAULT '' COMMENT '副本对话（怪物旁白）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='副本敌人';



# Dump of table mission_levels
# ------------------------------------------------------------

DROP TABLE IF EXISTS `mission_levels`;

CREATE TABLE `mission_levels` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '副本地图ID',
  `mission_id` smallint(6) NOT NULL COMMENT '副本ID',
  `name` varchar(10) NOT NULL COMMENT '关卡名称',
  `sign` varchar(50) NOT NULL COMMENT '资源标识',
  `fog_mode` tinyint(4) NOT NULL COMMENT '迷雾模式',
  `music` varchar(20) NOT NULL COMMENT '音乐资源标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='副本地图';



# Dump of table town
# ------------------------------------------------------------

DROP TABLE IF EXISTS `town`;

CREATE TABLE `town` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '城镇ID,-1为集会所',
  `lock` int(11) NOT NULL COMMENT '解锁权值',
  `name` varchar(10) NOT NULL DEFAULT '' COMMENT '城镇名称',
  `sign` varchar(30) NOT NULL DEFAULT '' COMMENT '资源标识',
  `music` varchar(20) NOT NULL COMMENT '音乐资源标识',
  `start_x` int(11) NOT NULL COMMENT '出生点x轴坐标',
  `start_y` int(11) NOT NULL COMMENT '出生点y轴坐标',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='城镇';




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

THESQL1
);

?>