<?php /* dump file */

	db_execute($db, <<<THESQL1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
DROP TABLE IF EXISTS `announcement`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `announcement` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '公告模版ID',
  `sign` varchar(30) DEFAULT NULL COMMENT '唯一标识',
  `type` tinyint(4) NOT NULL COMMENT '0后台公告， 1模块公告, 2活动公告',
  `name` varchar(30) DEFAULT NULL COMMENT '公告名',
  `parameters` varchar(1024) NOT NULL COMMENT '参数',
  `content` varchar(1024) NOT NULL COMMENT '内容',
  `duration` int(11) NOT NULL COMMENT '消息存活时间（秒）',
  `show_cyle` int(11) NOT NULL COMMENT '重复展示时间间隔（秒）',
  PRIMARY KEY (`id`),
  UNIQUE KEY `sign` (`sign`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='公告模版';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `announcement` DISABLE KEYS */;
INSERT INTO `announcement` VALUES (1,'TestAnnounce',0,'测试公告','p1,参数一','欢迎来到仙侠道',12000,60);
/*!40000 ALTER TABLE `announcement` ENABLE KEYS */;
DROP TABLE IF EXISTS `arena_award_box`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `arena_award_box` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `max_rank` int(11) NOT NULL COMMENT '排名',
  `ingot` int(11) NOT NULL COMMENT '元宝',
  `coins` int(11) NOT NULL COMMENT '铜钱',
  `item_id` smallint(6) NOT NULL COMMENT '物品ID',
  `item_num` smallint(6) NOT NULL COMMENT '物品数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='比武场奖励宝箱';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `arena_award_box` DISABLE KEYS */;
/*!40000 ALTER TABLE `arena_award_box` ENABLE KEYS */;
DROP TABLE IF EXISTS `battle_item_config`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `battle_item_config` (
  `id` smallint(6) NOT NULL COMMENT '道具物品ID',
  `target_type` tinyint(4) NOT NULL COMMENT '目标对象',
  `effect_type` tinyint(4) NOT NULL COMMENT '效果类型',
  `config` text COMMENT '产生效果配置',
  `keep` tinyint(4) NOT NULL COMMENT '持续回合',
  `max_override` tinyint(4) NOT NULL COMMENT '最大叠加数',
  `can_use_level` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否可以在关卡中使用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='战斗道具';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `battle_item_config` DISABLE KEYS */;
INSERT INTO `battle_item_config` VALUES (209,3,1,'[{\"mod\":\"1\",\"val\":\"5000\"}]',0,0,1),(210,3,1,'[{\"mod\":\"1\",\"val\":\"20000\"}]',0,0,1),(211,3,1,'[{\"mod\":\"1\",\"val\":\"99999\"}]',0,0,1),(212,6,1,'[{\"mod\":\"3\",\"val\":\"100\"}]',0,0,0),(213,5,1,'[{\"mod\":\"1\",\"val\":\"-10000\"}]',0,0,0),(214,5,1,'[{\"mod\":\"1\",\"val\":\"-30000\"}]',0,0,0),(215,2,3,'[{\"mod\":\"13\",\"val\":\"1000\"}]',2,1,0),(216,1,3,'[{\"mod\":\"0\",\"val\":\"10\"},{\"mod\":\"16\",\"val\":\"60\"}]',1,1,0),(217,2,3,'[{\"mod\":\"7\",\"val\":\"0\"}]',1,1,0),(250,7,2,'[{\"mod\":\"4\",\"val\":\"50\"}]',0,0,0),(251,7,2,'[{\"mod\":\"4\",\"val\":\"100\"}]',0,0,0);
/*!40000 ALTER TABLE `battle_item_config` ENABLE KEYS */;
DROP TABLE IF EXISTS `battle_pet`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `battle_pet` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `quality` tinyint(4) NOT NULL COMMENT '品质',
  `pet_id` int(10) unsigned NOT NULL COMMENT '灵宠ID(enemy_role)',
  `desc` varchar(100) DEFAULT '' COMMENT '灵宠描述',
  `round_attack` tinyint(4) DEFAULT '1' COMMENT '单回合行动次数',
  `cost_power` tinyint(4) NOT NULL COMMENT '召唤时消耗精气',
  `live_round` tinyint(4) NOT NULL COMMENT '召唤后存活回合数',
  `live_pos` tinyint(4) NOT NULL COMMENT '召唤后出现的位置(1-前排；2-后排；3-左侧；4-右侧)',
  `activate_ball_num` tinyint(4) NOT NULL COMMENT '激活需要的契约球数量',
  `skill` smallint(6) NOT NULL COMMENT '灵宠技能',
  `item_battle_pet_id` int(11) NOT NULL COMMENT '灵宠契约球',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='灵宠';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `battle_pet` DISABLE KEYS */;
INSERT INTO `battle_pet` VALUES (1,0,91,'',1,1,3,2,1,108,252),(2,0,92,'可在熔岩火山六捕捉',1,1,3,1,2,36,253),(3,1,93,'可在剑灵密室六捕捉',2,2,2,3,3,35,254);
/*!40000 ALTER TABLE `battle_pet` ENABLE KEYS */;
DROP TABLE IF EXISTS `chest`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `chest` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `type` tinyint(4) NOT NULL COMMENT '类型:1 - 青铜宝箱, 2 - 神龙宝箱',
  `quality` tinyint(4) NOT NULL COMMENT '宝箱品质',
  `probability` tinyint(4) NOT NULL COMMENT '概率（%）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COMMENT='宝箱品质';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `chest` DISABLE KEYS */;
INSERT INTO `chest` VALUES (1,2,1,50),(2,1,1,50),(3,1,2,20),(4,1,3,15),(5,1,4,8),(6,1,5,5),(7,1,6,2),(8,2,2,20),(9,2,3,15),(10,2,4,8),(11,2,5,5),(12,2,6,2);
/*!40000 ALTER TABLE `chest` ENABLE KEYS */;
DROP TABLE IF EXISTS `chest_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `chest_item` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `chest_id` int(11) NOT NULL COMMENT '宝箱id',
  `type` tinyint(4) NOT NULL COMMENT '物品类型',
  `item_id` smallint(6) NOT NULL COMMENT '物品',
  `item_num` int(11) NOT NULL COMMENT '数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb4 COMMENT='宝箱物品';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `chest_item` DISABLE KEYS */;
INSERT INTO `chest_item` VALUES (1,2,3,30,5),(2,2,3,236,5),(3,2,3,238,5),(4,2,3,237,5),(5,2,3,231,5),(6,2,3,230,1),(7,2,3,239,1),(9,3,3,209,2),(10,3,1,244,12000),(11,3,3,235,5),(12,3,3,200,3),(13,4,3,52,5),(14,4,3,210,5),(15,4,1,244,18000),(16,4,3,30,10),(19,5,5,37,3),(21,1,3,30,20),(22,1,3,200,10),(23,1,3,232,10),(24,1,3,241,20),(25,1,2,242,300),(26,8,1,0,25000),(27,8,3,30,80),(28,8,3,52,50),(29,8,3,249,5),(30,8,3,200,5),(34,12,3,212,1),(35,11,3,234,5),(36,11,3,215,3),(37,11,5,1,1),(38,11,5,5,1),(39,9,3,52,30),(40,9,3,210,10),(41,9,5,37,5),(42,9,3,214,5),(43,8,3,235,50),(44,10,5,4,1),(45,10,5,8,1),(46,10,5,10,1),(48,11,3,213,3),(49,6,3,232,5),(51,6,3,215,1),(61,11,5,14,1),(66,10,3,211,5),(67,10,5,17,1),(68,10,5,19,1),(71,12,5,23,1),(72,12,5,11,1),(73,12,5,35,1),(74,7,5,4,1),(75,7,5,8,1),(76,7,5,20,1),(77,7,5,21,1),(78,7,5,25,1),(86,6,5,3,1),(87,6,5,7,1),(88,6,5,15,1),(89,6,5,12,1),(90,6,3,213,1),(92,3,5,2,1),(93,3,5,6,1),(98,1,3,236,80),(99,1,3,250,1),(100,1,3,206,1);
/*!40000 ALTER TABLE `chest_item` ENABLE KEYS */;
DROP TABLE IF EXISTS `coins_exchange`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `coins_exchange` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `unique_key` smallint(6) NOT NULL COMMENT '第几次兑换',
  `ingot` bigint(20) NOT NULL COMMENT '消耗元宝',
  `coins` bigint(20) NOT NULL COMMENT '获得铜币',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_key` (`unique_key`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8mb4 COMMENT='铜币兑换收益表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `coins_exchange` DISABLE KEYS */;
INSERT INTO `coins_exchange` VALUES (1,1,10,20000),(2,2,20,21000),(3,3,20,22000),(4,4,20,23000),(5,5,20,24000),(8,6,40,25000),(9,7,40,26000),(10,8,40,27000),(11,9,40,28000),(12,10,40,29000),(13,11,80,32000),(14,12,80,33000),(15,13,80,34000),(16,14,80,35000),(17,15,80,36000),(18,16,100,38000),(19,17,100,39000),(20,18,100,40000),(21,19,100,41000),(22,20,100,42000),(23,21,150,44000),(24,22,150,45000),(25,23,150,46000),(26,24,150,47000),(27,25,150,48000),(28,26,200,50000),(29,27,200,51000),(30,28,200,52000),(31,29,200,53000),(32,30,200,54000),(33,31,300,56000),(34,32,300,57000),(35,33,300,58000),(36,34,300,59000),(37,35,300,60000),(38,36,400,63000),(39,37,400,64000),(40,38,400,65000),(41,39,400,66000),(42,40,400,67000);
/*!40000 ALTER TABLE `coins_exchange` ENABLE KEYS */;
DROP TABLE IF EXISTS `enemy_boss_script`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `enemy_boss_script` (
  `boss_id` int(10) unsigned NOT NULL COMMENT '怪物ID',
  `config` text COMMENT '脚本',
  PRIMARY KEY (`boss_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='boss配置脚本';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `enemy_boss_script` DISABLE KEYS */;
INSERT INTO `enemy_boss_script` VALUES (2,'[{\"condition\":\"1\",\"val\":\"100\",\"skill_id\":\"36\",\"power\":\"20\"}]'),(8,'[{\"condition\":\"1\",\"val\":\"100\",\"skill_id\":\"9\",\"power\":\"150\"}]'),(11,'[{\"condition\":\"1\",\"val\":\"1200\",\"skill_id\":\"32\",\"power\":\"200\"},{\"condition\":\"1\",\"val\":\"1200\",\"skill_id\":\"33\",\"power\":\"200\"},{\"condition\":\"1\",\"val\":\"600\",\"skill_id\":\"34\",\"power\":\"200\"}]'),(14,'[{\"condition\":\"1\",\"val\":\"2000\",\"skill_id\":\"34\",\"power\":\"220\"}]'),(18,'[{\"condition\":\"1\",\"val\":\"2000\",\"skill_id\":\"30\",\"power\":\"500\"}]'),(19,'[{\"condition\":\"1\",\"val\":\"3000\",\"skill_id\":\"18\",\"power\":\"600\"},{\"condition\":\"1\",\"val\":\"2000\",\"skill_id\":\"18\",\"power\":\"600\"}]'),(23,'[{\"condition\":\"1\",\"val\":\"3500\",\"skill_id\":\"40\",\"power\":\"700\"},{\"condition\":\"1\",\"val\":\"2000\",\"skill_id\":\"40\",\"power\":\"700\"},{\"condition\":\"1\",\"val\":\"1000\",\"skill_id\":\"108\",\"power\":\"10000\"}]'),(25,'[{\"condition\":\"1\",\"val\":\"1500\",\"skill_id\":\"45\",\"power\":\"180\"},{\"condition\":\"1\",\"val\":\"500\",\"skill_id\":\"45\",\"power\":\"180\"}]'),(26,'[{\"condition\":\"1\",\"val\":\"500\",\"skill_id\":\"9\",\"power\":\"40\"}]'),(88,'[{\"condition\":\"1\",\"val\":\"1000\",\"skill_id\":\"42\",\"power\":\"50\"}]'),(89,'[{\"condition\":\"1\",\"val\":\"1000\",\"skill_id\":\"42\",\"power\":\"50\"}]');
/*!40000 ALTER TABLE `enemy_boss_script` ENABLE KEYS */;
DROP TABLE IF EXISTS `enemy_deploy_form`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `enemy_deploy_form` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL COMMENT '关联此阵法的某表唯一ID',
  `battle_type` tinyint(4) NOT NULL COMMENT '战场类型(0--关卡;1--资源关卡;2--极限关卡;3--多人关卡)',
  `pos1` int(11) NOT NULL DEFAULT '0' COMMENT '位置1上的敌人ID',
  `pos2` int(11) NOT NULL DEFAULT '0' COMMENT '位置2上的敌人ID',
  `pos3` int(11) NOT NULL DEFAULT '0' COMMENT '位置3上的敌人ID',
  `pos4` int(11) NOT NULL DEFAULT '0' COMMENT '位置4上的敌人ID',
  `pos5` int(11) NOT NULL DEFAULT '0' COMMENT '位置5上的敌人ID',
  `pos6` int(11) NOT NULL DEFAULT '0' COMMENT '位置6上的敌人ID',
  `pos7` int(11) NOT NULL DEFAULT '0' COMMENT '位置7上的敌人ID',
  `pos8` int(11) NOT NULL DEFAULT '0' COMMENT '位置8上的敌人ID',
  `pos9` int(11) NOT NULL DEFAULT '0' COMMENT '位置9上的敌人ID',
  `pos10` int(11) NOT NULL DEFAULT '0' COMMENT '位置10上的敌人ID',
  `pos11` int(11) NOT NULL DEFAULT '0' COMMENT '位置11上的敌人ID',
  `pos12` int(11) NOT NULL DEFAULT '0' COMMENT '位置12上的敌人ID',
  `pos13` int(11) NOT NULL DEFAULT '0' COMMENT '位置13上的敌人ID',
  `pos14` int(11) NOT NULL DEFAULT '0' COMMENT '位置14上的敌人ID',
  `pos15` int(11) NOT NULL DEFAULT '0' COMMENT '位置15上的敌人ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8mb4 COMMENT='怪物阵法表单';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `enemy_deploy_form` DISABLE KEYS */;
INSERT INTO `enemy_deploy_form` VALUES (1,93,0,0,3,0,3,0,0,0,26,0,0,0,0,0,0,0),(7,36,0,0,0,24,25,24,0,0,0,0,0,0,0,0,0,0),(8,18,0,0,3,0,3,0,0,0,2,0,0,0,0,0,0,0),(9,27,0,0,5,0,5,0,0,0,8,0,0,0,0,0,0,0),(10,45,0,0,9,0,9,0,0,0,11,0,0,0,0,0,0,0),(11,54,0,0,27,0,27,0,0,0,14,0,0,0,0,0,0,0),(12,87,0,0,22,0,22,0,0,0,23,0,0,0,0,0,0,0),(13,63,0,0,16,0,0,16,0,0,0,18,0,0,0,0,0,0),(14,72,0,0,15,0,0,15,0,0,0,19,0,0,0,0,0,0),(15,90,0,0,22,0,22,0,0,22,23,22,0,0,0,0,0,0),(17,1,3,0,0,0,0,0,0,0,26,0,0,0,0,0,0,0),(18,2,3,0,0,8,0,0,0,0,0,0,0,0,0,0,0,0),(19,3,3,0,0,0,25,0,0,0,0,0,0,0,0,0,0,0),(20,4,3,0,0,11,0,0,0,0,0,0,0,0,0,0,0,0),(21,5,3,0,0,0,14,0,0,0,0,0,0,0,0,0,0,0),(24,13,0,0,3,0,3,0,0,0,0,0,0,0,0,0,0,0),(25,14,0,0,0,0,0,0,0,3,0,3,0,0,0,0,0,0),(26,2,0,0,3,0,3,0,0,0,4,0,0,0,0,0,0,0),(27,95,0,0,0,3,0,0,0,4,0,4,0,0,0,0,0,0),(28,170,1,0,0,0,0,0,0,0,2,0,0,0,88,0,88,0),(30,172,1,0,0,0,0,0,0,0,2,0,0,0,88,0,88,0),(32,176,1,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0),(33,174,1,0,0,0,0,0,0,0,26,0,0,0,88,0,88,0),(34,178,1,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0),(36,179,1,0,0,0,88,0,0,0,26,0,0,0,88,0,0,0),(37,180,1,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0),(39,182,1,0,0,0,0,0,0,0,25,0,0,0,88,0,88,0),(41,184,1,0,0,0,0,0,0,0,19,0,0,0,88,0,88,0),(43,186,1,0,0,0,0,0,0,0,11,0,0,0,88,0,88,0),(44,188,1,0,0,0,0,0,0,0,11,0,0,0,88,0,88,0),(45,190,1,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0),(46,192,1,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0),(47,194,1,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0),(48,195,1,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0),(49,197,1,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0),(50,199,1,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0),(51,6,3,0,0,18,0,0,0,0,0,0,0,0,0,0,0,0),(52,7,3,0,0,0,19,0,0,0,0,0,0,0,0,0,0,0),(53,8,3,0,0,0,0,0,0,0,23,0,0,0,0,0,0,0);
/*!40000 ALTER TABLE `enemy_deploy_form` ENABLE KEYS */;
DROP TABLE IF EXISTS `enemy_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
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
  `skill_id` smallint(6) NOT NULL COMMENT '绝招ID',
  `skill_force` int(11) NOT NULL COMMENT '绝招威力',
  `skill2_id` smallint(6) NOT NULL COMMENT '绝招2 ID',
  `skill2_force` int(11) NOT NULL COMMENT '绝招2 威力',
  `sunder_max_value` int(11) NOT NULL COMMENT '护甲值',
  `sunder_min_hurt_rate` int(11) NOT NULL COMMENT '破甲前起始的伤害转换率（百分比）',
  `sunder_end_hurt_rate` int(11) NOT NULL COMMENT '破甲后的伤害转换率（百分比）',
  `sunder_attack` int(11) NOT NULL COMMENT '攻击破甲值',
  `skill_wait` tinyint(4) NOT NULL COMMENT '绝招蓄力回合',
  `release_num` tinyint(4) NOT NULL COMMENT '释放次数',
  `recover_round_num` tinyint(4) NOT NULL COMMENT '恢复回合数',
  `common_attack_num` tinyint(4) NOT NULL COMMENT '入场普通攻击次数',
  `jump_attack` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否跳跃攻击',
  `body_size` tinyint(4) NOT NULL DEFAULT '1' COMMENT '怪物体型',
  `is_boss` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否是boss 0否,1是',
  `scale_size` smallint(5) NOT NULL DEFAULT '100' COMMENT '怪物缩放比%',
  `skill_config` text COMMENT '技能配置',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=95 DEFAULT CHARSET=utf8mb4 COMMENT='敌人角色数据';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `enemy_role` DISABLE KEYS */;
INSERT INTO `enemy_role` VALUES (1,'疯狂草妖','CaoYao',10,2926,0,0,350,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,1,0,144,NULL),(2,'二货兔','ErHuoTu',4,350,0,0,100,0,0,0,0,0,0,0,0,0,-1,20,0,0,100,100,200,1,0,1,2,1,1,2,1,180,'[{\"rand\":\"50\",\"skill_id\":\"35\",\"power\":\"10\"}]'),(3,'疯狂竹筒精','ZhuTongJing',2,89,0,0,104,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,1,0,144,''),(4,'疯狂林精','LinJing',3,123,0,0,123,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,1,0,144,NULL),(5,'疯狂黑狼','HeiLang',6,948,0,0,208,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,1,0,144,''),(6,'疯狂鬼火','GuiHuo',6,948,0,0,208,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,1,0,144,NULL),(7,'疯狂灯笼怪','DengLongGuai',6,948,0,0,208,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,1,0,144,NULL),(8,'天狼妖','TianLangYao',7,3160,0,0,208,0,0,0,0,0,0,0,0,0,10,80,0,0,100,100,200,1,0,1,2,0,1,3,1,144,'[{\"rand\":\"30\",\"skill_id\":\"10\",\"power\":\"50\"},{\"rand\":\"10\",\"skill_id\":\"19\",\"power\":\"50\"},{\"rand\":\"20\",\"skill_id\":\"44\",\"power\":\"50\"}]'),(9,'疯狂莲藕精','LianOuJing',10,2926,0,0,350,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,1,0,144,NULL),(10,'愤怒林精','LinJing',10,2926,0,0,350,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,1,0,144,NULL),(11,'金蟾王','JingChanWang',11,8127,0,0,454,0,0,0,0,0,0,0,0,0,32,100,0,0,100,100,200,1,0,1,2,1,1,3,1,144,'[{\"rand\":\"20\",\"skill_id\":\"36\",\"power\":\"200\"},{\"rand\":\"20\",\"skill_id\":\"31\",\"power\":\"200\"},{\"rand\":\"20\",\"skill_id\":\"43\",\"power\":\"200\"}]'),(12,'测试咆哮厉爪','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,150,1,0,1,2,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"9\",\"power\":\"10\"}]'),(13,'疯狂毒蛇','DuShe',12,3335,0,0,458,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,1,0,144,NULL),(14,'剧毒臭泥','JuDuChouNi',13,9262,0,0,524,0,0,0,0,0,0,0,0,0,32,150,0,0,100,100,200,1,0,1,2,1,0,3,1,144,'[{\"rand\":\"30\",\"skill_id\":\"31\",\"power\":\"220\"},{\"rand\":\"15\",\"skill_id\":\"32\",\"power\":\"220\"},{\"rand\":\"15\",\"skill_id\":\"33\",\"power\":\"220\"}]'),(15,'疯狂火蝎','HuoXie',15,4149,0,0,525,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,1,0,120,NULL),(16,'疯狂燃魁','RanKui',15,4149,0,0,525,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,2,0,120,NULL),(17,'疯狂熔岩虫','RongYanChong',16,4149,0,0,525,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,2,0,130,NULL),(18,'燃魁首领','RanKui',17,10390,0,0,595,0,0,0,0,0,0,0,0,0,15,200,0,0,100,100,200,1,0,1,2,1,0,3,1,180,'[{\"rand\":\"30\",\"skill_id\":\"15\",\"power\":\"450\"},{\"rand\":\"15\",\"skill_id\":\"16\",\"power\":\"500\"},{\"rand\":\"15\",\"skill_id\":\"29\",\"power\":\"500\"}]'),(19,'炎龙','YanLong',17,11525,0,0,665,0,0,0,0,0,0,0,0,0,17,250,0,0,100,100,200,1,0,1,2,1,0,3,1,144,'[{\"rand\":\"30\",\"skill_id\":\"15\",\"power\":\"400\"},{\"rand\":\"15\",\"skill_id\":\"16\",\"power\":\"500\"},{\"rand\":\"15\",\"skill_id\":\"17\",\"power\":\"500\"},{\"rand\":\"10\",\"skill_id\":\"43\",\"power\":\"500\"}]'),(20,'疯狂吸血蝙蝠','HeYiJuFu',19,4961,0,0,644,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,2,0,144,NULL),(21,'疯狂败亡之剑','BaiWangZhiJian',19,4961,0,0,644,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,1,0,144,NULL),(22,'疯狂剑之守卫','JianZhiShouWei',19,4961,0,0,673,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,1,0,120,NULL),(23,'古代剑灵','GuDaiJianLing',20,16536,0,0,806,0,0,0,0,0,0,0,0,0,18,350,0,0,100,100,200,1,0,1,2,0,1,2,1,144,'[{\"rand\":\"20\",\"skill_id\":\"39\",\"power\":\"700\"},{\"rand\":\"20\",\"skill_id\":\"38\",\"power\":\"700\"},{\"rand\":\"20\",\"skill_id\":\"42\",\"power\":\"700\"}]'),(24,'疯狂阴影','YinYing',6,1482,0,0,236,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,2,0,144,''),(25,'妖龙','YaoLong',8,3804,0,0,346,0,0,0,0,0,0,0,0,0,9,90,0,0,50,100,200,1,0,1,2,0,1,3,1,144,'[{\"rand\":\"10\",\"skill_id\":\"20\",\"power\":\"150\"},{\"rand\":\"10\",\"skill_id\":\"25\",\"power\":\"150\"},{\"rand\":\"10\",\"skill_id\":\"21\",\"power\":\"150\"},{\"rand\":\"10\",\"skill_id\":\"24\",\"power\":\"150\"}]'),(26,'刀疤兔','DaoBaTu',5,1300,0,0,151,0,0,0,0,0,0,0,0,0,43,30,0,0,100,100,200,1,0,1,2,1,1,2,1,144,'[{\"rand\":\"30\",\"skill_id\":\"35\",\"power\":\"20\"},{\"rand\":\"20\",\"skill_id\":\"43\",\"power\":\"30\"}]'),(27,'疯狂黑翼巨蝠','HeYiJuFu',12,3335,0,0,458,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,2,0,144,NULL),(31,'竹筒精','ZhuTongJing',1,81,0,0,93,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,1,0,144,NULL),(33,'林精','LinJing',2,111,0,0,105,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,1,0,144,NULL),(34,'黑狼','HeiLang',5,506,0,0,180,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,1,0,144,''),(35,'鬼火','GuiHuo',5,506,0,0,180,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,1,0,144,NULL),(36,'灯笼怪','DengLongGuai',5,506,0,0,180,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,1,0,144,NULL),(37,'阴影','YinYing',5,912,0,0,236,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,2,0,144,NULL),(38,'暴走林精','LinJing',9,1300,0,0,350,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,1,0,144,NULL),(39,'莲藕精','LianOuJing',9,1300,0,0,350,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,1,0,144,''),(40,'草妖','CaoYao',9,1300,0,0,350,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,1,0,144,NULL),(41,'黑翼巨蝠','HeYiJuFu',11,1482,0,0,427,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,2,0,144,NULL),(42,'毒蛇','DuShe',11,1482,0,0,427,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,1,0,144,NULL),(43,'火蝎','HuoXie',14,2075,0,0,466,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,1,0,120,NULL),(44,'燃魁','RanKui',14,2075,0,0,496,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,2,0,120,NULL),(45,'熔岩虫','RongYanChong',14,2075,0,0,466,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,2,0,130,NULL),(46,'败亡之剑','BaiWangZhiJian',18,2481,0,0,584,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,1,0,144,NULL),(47,'吸血蝙蝠','HeYiJuFu',18,2481,0,0,584,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,2,0,144,NULL),(48,'剑之守卫','JianZhiShouWei',18,2481,0,0,614,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,1,0,1,2,0,1,1,0,120,NULL),(49,'测试凶猛撕咬','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"10\",\"power\":\"100\"}]'),(50,'测试冰烈','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"11\",\"power\":\"100\"}]'),(51,'测试火烈','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"15\",\"power\":\"100\"}]'),(52,'测试风烈','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"19\",\"power\":\"100\"}]'),(53,'测试雷烈','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"23\",\"power\":\"100\"}]'),(54,'测试土烈','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"27\",\"power\":\"100\"}]'),(55,'测试毒烈','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"31\",\"power\":\"100\"}]'),(56,'测试多连斩','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"35\",\"power\":\"100\"}]'),(57,'测试力劈华山','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"36\",\"power\":\"100\"}]'),(58,'测试死亡标记','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"41\",\"power\":\"100\"}]'),(59,'测试万箭穿心','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"42\",\"power\":\"100\"}]'),(60,'测试狮吼功','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"43\",\"power\":\"100\"}]'),(61,'测试驱散','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"46\",\"power\":\"100\"}]'),(62,'测试增益','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"109\",\"power\":\"100\"}]'),(63,'测试白莲横江','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"37\",\"power\":\"100\"}]'),(64,'测试横扫千军','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"38\",\"power\":\"100\"}]'),(65,'测试乾坤刀气','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"39\",\"power\":\"100\"}]'),(66,'测试野蛮冲撞','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"44\",\"power\":\"100\"}]'),(67,'测试治疗','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"108\",\"power\":\"100\"}]'),(68,'测试三千洛水剑','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"40\",\"power\":\"100\"}]'),(69,'测试如殂随行','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"0100\",\"skill_id\":\"45\",\"power\":\"100\"}]'),(70,'测试冰烈横向','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"12\",\"power\":\"100\"}]'),(71,'测试火烈横向','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"16\",\"power\":\"100\"}]'),(72,'测试风烈横向','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"20\",\"power\":\"100\"}]'),(73,'测试雷烈横向','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"24\",\"power\":\"100\"}]'),(74,'测试土烈横向','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,2,0,144,'[{\"rand\":\"100\",\"skill_id\":\"28\",\"power\":\"100\"}]'),(75,'测试毒烈横向','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"32\",\"power\":\"100\"}]'),(76,'测试冰烈纵向','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"13\",\"power\":\"100\"}]'),(77,'测试火烈纵向','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"17\",\"power\":\"100\"}]'),(78,'测试风烈纵向','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"21\",\"power\":\"100\"}]'),(79,'测试雷烈纵向','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"25\",\"power\":\"100\"}]'),(80,'测试土烈纵向','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"29\",\"power\":\"100\"}]'),(81,'测试毒烈纵向','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"33\",\"power\":\"100\"}]'),(82,'测试冰烈全体','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"14\",\"power\":\"100\"}]'),(83,'测试火烈全体','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"18\",\"power\":\"100\"}]'),(84,'测试风烈全体','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"22\",\"power\":\"100\"}]'),(85,'测试雷烈全体','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"26\",\"power\":\"100\"}]'),(86,'测试土烈全体','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"30\",\"power\":\"100\"}]'),(87,'测试毒烈全体','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"34\",\"power\":\"100\"}]'),(88,'铜钱怪','TongQianGuai',20,3000,0,0,600,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,1,1,0,144,NULL),(89,'经验怪','JingYanGuai',20,3000,0,0,600,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,1,1,0,144,NULL),(90,'测试破甲','GuDaiJianLing',99,10000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,50,0,0,0,0,0,1,0,144,''),(91,'草精','CaoYao',10,300,0,0,150,0,0,0,0,0,0,0,0,0,0,0,0,0,200,100,200,1,0,0,0,0,0,1,0,120,'[{\"rand\":\"100\",\"skill_id\":\"108\",\"power\":\"1000\"}]'),(92,'火灵','RanKui',20,700,0,0,240,50,0,0,0,0,0,0,0,0,0,0,0,0,200,100,200,1,0,0,0,0,0,1,0,100,'[{\"rand\":\"100\",\"skill_id\":\"36\",\"power\":\"500\"}]'),(93,'剑魄','BaiWangZhiJian',25,1000,0,100,400,100,0,0,0,0,0,0,0,0,0,0,0,0,200,100,200,1,0,0,0,0,0,1,0,120,'[{\"rand\":\"100\",\"skill_id\":\"35\",\"power\":\"500\"}]'),(94,'测试中毒','GuDaiJianLing	',99,10000,100,100,500,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,2,0,144,'[{\"rand\":\"100\",\"skill_id\":\"34\",\"power\":\"0\"}]');
/*!40000 ALTER TABLE `enemy_role` ENABLE KEYS */;
DROP TABLE IF EXISTS `equipment_appendix`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `equipment_appendix` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `level` int(11) NOT NULL COMMENT '等级',
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
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COMMENT='装备追加属性表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `equipment_appendix` DISABLE KEYS */;
INSERT INTO `equipment_appendix` VALUES (1,1,100,20,20,20,10,5,5,5,5,5,5),(2,2,250,100,100,100,50,20,20,20,20,20,20),(3,3,1000,400,400,400,200,40,40,40,40,40,40),(4,4,2000,800,800,800,400,80,80,80,80,80,80),(5,5,3000,1200,1200,1200,600,120,120,120,120,120,120),(6,6,4000,1600,1600,1600,800,160,160,160,160,160,160),(7,7,5000,2000,2000,2000,1000,200,200,200,200,200,200);
/*!40000 ALTER TABLE `equipment_appendix` ENABLE KEYS */;
DROP TABLE IF EXISTS `equipment_decompose`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `equipment_decompose` (
  `id` tinyint(4) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `level` int(11) NOT NULL COMMENT '等级下限',
  `fragment_num` smallint(6) NOT NULL COMMENT '获得部位碎片数量',
  `crystal_num` smallint(6) NOT NULL COMMENT '获得结晶数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COMMENT='装备分解获得材料';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `equipment_decompose` DISABLE KEYS */;
INSERT INTO `equipment_decompose` VALUES (1,1,1,1),(2,20,2,2),(3,40,4,4),(4,60,6,6),(5,80,8,8),(6,100,10,10);
/*!40000 ALTER TABLE `equipment_decompose` ENABLE KEYS */;
DROP TABLE IF EXISTS `equipment_recast`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `equipment_recast` (
  `id` tinyint(4) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `level` int(11) NOT NULL COMMENT '等级下限',
  `quality` tinyint(4) NOT NULL COMMENT '品质',
  `fragment_num` smallint(6) NOT NULL COMMENT '需要部位碎片数量',
  `blue_crystal_num` smallint(6) NOT NULL COMMENT '需要蓝色结晶数量',
  `purple_crystal_num` smallint(6) NOT NULL COMMENT '需要紫色结晶数量',
  `golden_crystal_num` smallint(6) NOT NULL COMMENT '需要金色结晶数量',
  `orange_crystal_num` smallint(6) NOT NULL COMMENT '需要橙色结晶数量',
  `consume_coin` bigint(20) NOT NULL COMMENT '消耗的铜钱',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COMMENT='装备重铸消耗';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `equipment_recast` DISABLE KEYS */;
INSERT INTO `equipment_recast` VALUES (1,1,1,1,1,0,0,0,1000),(2,20,1,2,2,0,0,0,2000),(3,40,1,4,4,0,0,0,4000),(4,60,1,6,6,0,0,0,6000),(5,80,1,8,8,0,0,0,8000),(6,100,1,10,10,0,0,0,10000),(7,1,2,1,1,1,0,0,2000),(8,20,2,2,2,2,0,0,4000),(9,40,2,4,4,4,0,0,8000),(10,60,2,6,6,6,0,0,12000),(11,80,2,8,8,8,0,0,16000),(12,100,2,10,10,10,0,0,20000),(13,1,3,1,1,1,1,0,4000),(14,20,3,2,2,2,2,0,8000),(15,40,3,4,4,4,4,0,16000),(16,60,3,6,6,6,6,0,24000),(17,80,3,8,8,8,8,0,32000),(18,100,3,10,10,10,10,0,40000),(19,1,4,1,1,1,1,1,8000),(20,20,4,2,2,2,2,2,16000),(21,40,4,4,4,4,4,4,32000),(22,60,4,6,6,6,6,6,48000),(23,80,4,8,8,8,8,8,64000),(24,100,4,10,10,10,10,10,80000);
/*!40000 ALTER TABLE `equipment_recast` ENABLE KEYS */;
DROP TABLE IF EXISTS `equipment_refine`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `equipment_refine` (
  `id` tinyint(4) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `level` int(11) NOT NULL COMMENT '等级下限',
  `quality` tinyint(4) NOT NULL COMMENT '品质',
  `fragment_num` smallint(6) NOT NULL COMMENT '需要部位碎片数量',
  `blue_crystal_num` smallint(6) NOT NULL COMMENT '需要蓝色结晶数量',
  `purple_crystal_num` smallint(6) NOT NULL COMMENT '需要紫色结晶数量',
  `golden_crystal_num` smallint(6) NOT NULL COMMENT '需要金色结晶数量',
  `orange_crystal_num` smallint(6) NOT NULL COMMENT '需要橙色结晶数量',
  `level1_consume_coin` bigint(20) NOT NULL COMMENT '精练到1级消耗的铜钱',
  `level2_consume_coin` bigint(20) NOT NULL COMMENT '精练到2级消耗的铜钱',
  `level3_consume_coin` bigint(20) NOT NULL COMMENT '精练到3级消耗的铜钱',
  `level4_consume_coin` bigint(20) NOT NULL COMMENT '精练到4级消耗的铜钱',
  `level5_consume_coin` bigint(20) NOT NULL COMMENT '精练到5级消耗的铜钱',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COMMENT='装备精练消耗';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `equipment_refine` DISABLE KEYS */;
INSERT INTO `equipment_refine` VALUES (1,1,1,1,1,0,0,0,500,1000,1500,2000,2500),(2,20,1,1,1,0,0,0,1000,2000,3000,4000,5000),(3,40,1,2,2,0,0,0,2000,4000,6000,8000,10000),(4,60,1,2,2,0,0,0,3000,6000,9000,12000,15000),(5,80,1,3,3,0,0,0,4000,8000,12000,16000,20000),(6,100,1,3,3,0,0,0,5000,10000,15000,20000,25000),(7,1,2,1,1,1,0,0,1000,2000,3000,4000,5000),(8,20,2,1,1,1,0,0,2000,4000,6000,8000,10000),(9,40,2,2,2,2,0,0,4000,8000,12000,16000,20000),(10,60,2,2,2,2,0,0,6000,12000,18000,24000,30000),(11,80,2,3,3,3,0,0,8000,16000,24000,32000,40000),(12,100,2,3,3,3,0,0,10000,20000,30000,40000,50000),(13,1,3,1,1,1,1,0,2000,4000,6000,8000,10000),(14,20,3,1,1,1,1,0,4000,8000,12000,16000,20000),(15,40,3,2,2,2,2,0,8000,16000,24000,32000,40000),(16,60,3,2,2,2,2,0,12000,24000,36000,48000,60000),(17,80,3,3,3,3,3,0,16000,32000,48000,64000,80000),(18,100,3,3,3,3,3,0,20000,40000,60000,80000,100000),(19,1,4,1,1,1,1,1,4000,8000,12000,16000,20000),(20,20,4,1,1,1,1,1,8000,16000,24000,32000,40000),(21,40,4,2,2,2,2,2,16000,32000,48000,64000,80000),(22,60,4,2,2,2,2,2,24000,48000,72000,96000,120000),(23,80,4,3,3,3,3,3,32000,64000,96000,128000,160000),(24,100,4,3,3,3,3,3,40000,80000,120000,160000,200000);
/*!40000 ALTER TABLE `equipment_refine` ENABLE KEYS */;
DROP TABLE IF EXISTS `equipment_refine_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `equipment_refine_level` (
  `id` tinyint(4) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `level` tinyint(4) NOT NULL COMMENT '精练级别',
  `quality` tinyint(4) NOT NULL COMMENT '品质',
  `probability` tinyint(4) NOT NULL COMMENT '精练成功概率',
  `gain_pct` int(11) NOT NULL COMMENT '增益百分比',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COMMENT='精练武器对应概率与加成';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `equipment_refine_level` DISABLE KEYS */;
INSERT INTO `equipment_refine_level` VALUES (1,1,1,100,20),(2,2,1,80,40),(3,3,1,60,80),(4,4,1,40,120),(5,5,1,20,200),(6,1,2,100,10),(7,2,2,80,20),(8,3,2,60,40),(9,4,2,40,60),(10,5,2,20,100),(11,1,3,100,10),(12,2,3,80,20),(13,3,3,60,40),(14,4,3,40,60),(15,5,3,20,100),(16,1,4,100,10),(17,2,4,80,20),(18,3,4,60,40),(19,4,4,40,60),(20,5,4,20,100);
/*!40000 ALTER TABLE `equipment_refine_level` ENABLE KEYS */;
DROP TABLE IF EXISTS `faq`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `faq` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `order` int(11) NOT NULL COMMENT '顺序',
  `question` varchar(512) NOT NULL COMMENT '问题',
  `answer` varchar(1024) NOT NULL COMMENT '回答',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='问答';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `faq` DISABLE KEYS */;
INSERT INTO `faq` VALUES (1,1,'问：扣款成功但没有收到钻石，怎么办？','答：网络繁忙时，会在三小时内同步数据，如三个小时内还没有收到钻石，请联系客服，'),(2,2,'问：我明明有钻石，为什么钻石显示为0？','答：该情况可能由于服务器繁忙，数据拉取失败造成的。请稍后查看，或者退出游戏，结束程序，然后重新登陆游戏查看。如果还有时没有显示，请联系客服查询。请放心，您的钻石时不会丢失的。'),(3,3,'问：各区之间的钻石关系时怎么样的？','答：我们有四个大区：IOS手Q、IOS微信、android手Q、android微信。这四个大区之间彼此的数据不互通。每个大区下面会有数个小区。同一大区小的小区与小区之间的数据也互不想通。同一小区内的最多只有一个角色。'),(4,4,'问：为什么运营活动的开始和结束时间每个区不一样？','答：运营活动是时间时根据开区之日算起的，不同的区，开区的时间不一样所以活动的时间不一样。');
/*!40000 ALTER TABLE `faq` ENABLE KEYS */;
DROP TABLE IF EXISTS `func`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `func` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(10) NOT NULL COMMENT '功能名称',
  `sign` varchar(30) NOT NULL COMMENT '功能标识',
  `lock` smallint(6) NOT NULL COMMENT '功能权值',
  `level` smallint(6) NOT NULL DEFAULT '0' COMMENT '开启等级',
  `unique_key` bigint(20) NOT NULL DEFAULT '0' COMMENT '唯一权值',
  `need_play` tinyint(4) DEFAULT '1' COMMENT '是否需要播放 0不需要 1需要',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COMMENT='功能权值配置';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `func` DISABLE KEYS */;
INSERT INTO `func` VALUES (3,'魂侍','FUNC_GHOST',1000,0,1,1),(4,'资源关卡','FUNC_RESOURCE_LEVEL',2000,0,2,0),(5,'通天塔','FUNC_TOWER_LEVEL',3000,0,4,0),(6,'多人关卡','FUNC_MULTI_LEVEL',4000,0,8,1),(7,'剑心','FUNC_SWORD_SOUL',1500,0,16,1),(8,'灵宠','FUNC_BATTLE_PET',1300,0,32,1),(9,'神龙宝箱','FUNC_CHEST',1600,0,64,0);
/*!40000 ALTER TABLE `func` ENABLE KEYS */;
DROP TABLE IF EXISTS `ghost`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ghost` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '魂侍名称',
  `sign` varchar(30) NOT NULL DEFAULT '' COMMENT '资源标识',
  `town_id` tinyint(4) NOT NULL DEFAULT '0' COMMENT '城镇id（影界id）',
  `role_id` tinyint(4) NOT NULL DEFAULT '0' COMMENT '可装备角色id',
  `unique_key` smallint(6) NOT NULL DEFAULT '0' COMMENT '每个影界中魂侍的唯一标记',
  `init_star` tinyint(4) NOT NULL DEFAULT '0' COMMENT '初始星级',
  `health` int(11) NOT NULL DEFAULT '0' COMMENT '生命',
  `attack` int(11) NOT NULL DEFAULT '0' COMMENT '攻击',
  `defense` int(11) NOT NULL DEFAULT '0' COMMENT '防御',
  `speed` int(11) NOT NULL DEFAULT '0' COMMENT '速度',
  `desc` varchar(300) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='魂侍';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `ghost` DISABLE KEYS */;
/*!40000 ALTER TABLE `ghost` ENABLE KEYS */;
DROP TABLE IF EXISTS `ghost_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ghost_level` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `level` smallint(6) NOT NULL DEFAULT '0' COMMENT '魂侍等级',
  `exp` bigint(20) NOT NULL DEFAULT '0' COMMENT '升级所需经验',
  `need_fruit_num` int(11) NOT NULL DEFAULT '0' COMMENT '所需影界果实数量',
  `min_add_exp` bigint(20) NOT NULL DEFAULT '0' COMMENT '最小经验加值',
  `max_add_exp` bigint(20) NOT NULL DEFAULT '0' COMMENT '最大经验加值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=81 DEFAULT CHARSET=utf8mb4 COMMENT='魂侍等级表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `ghost_level` DISABLE KEYS */;
INSERT INTO `ghost_level` VALUES (1,1,900,1,240,360),(2,2,900,1,240,360),(3,3,900,1,240,360),(4,4,1200,1,240,360),(5,5,1200,1,240,360),(6,6,1200,1,240,360),(7,7,1500,1,240,360),(8,8,1500,1,240,360),(9,9,1500,1,240,360),(10,10,3000,1,240,360),(11,11,1800,1,240,360),(12,12,1800,1,240,360),(13,13,1800,1,240,360),(14,14,1800,1,240,360),(15,15,1800,1,240,360),(16,16,1800,1,240,360),(17,17,1800,1,240,360),(18,18,1800,1,240,360),(19,19,1800,1,240,360),(20,20,6000,2,480,720),(21,21,3600,2,480,720),(22,22,3600,2,480,720),(23,23,3600,2,480,720),(24,24,3600,2,480,720),(25,25,3600,2,480,720),(26,26,3600,2,480,720),(27,27,3600,2,480,720),(28,29,3600,2,480,720),(29,30,9000,3,780,1020),(30,31,5400,3,780,1020),(31,32,5400,3,780,1020),(32,33,5400,3,780,1020),(33,34,5400,3,780,1020),(34,35,5400,3,780,1020),(35,36,5400,3,780,1020),(36,37,5400,3,780,1020),(37,38,5400,3,780,1020),(38,39,5400,3,780,1020),(39,40,12000,4,1080,1320),(40,41,7200,4,1080,1320),(41,42,7200,4,1080,1320),(42,43,7200,4,1080,1320),(43,44,7200,4,1080,1320),(44,45,7200,4,1080,1320),(45,46,7200,4,1080,1320),(46,47,7200,4,1080,1320),(47,48,7200,4,1080,1320),(48,49,7200,4,1080,1320),(49,50,15000,5,1380,1620),(50,51,9000,5,1380,1620),(51,52,9000,5,1380,1620),(52,53,9000,5,1380,1620),(53,54,9000,5,1380,1620),(54,55,9000,5,1380,1620),(55,56,9000,5,1380,1620),(56,57,9000,5,1380,1620),(57,58,9000,5,1380,1620),(58,59,9000,5,1380,1620),(59,60,18000,6,1680,1920),(60,61,10800,6,1680,1920),(61,62,10800,6,1680,1920),(62,63,10800,6,1680,1920),(63,64,10800,6,1680,1920),(64,65,10800,6,1680,1920),(65,66,10800,6,1680,1920),(66,67,10800,6,1680,1920),(67,68,10800,6,1680,1920),(68,69,10800,6,1680,1920),(69,70,21000,7,1980,2220),(70,71,12600,7,1980,2220),(71,72,12600,7,1980,2220),(72,73,12600,7,1980,2220),(73,74,12600,7,1980,2220),(74,75,12600,7,1980,2220),(75,76,12600,7,1980,2220),(76,77,12600,7,1980,2220),(77,78,12600,7,1980,2220),(78,79,12600,7,1980,2220),(79,80,24000,8,2280,2520),(80,28,3600,2,480,720);
/*!40000 ALTER TABLE `ghost_level` ENABLE KEYS */;
DROP TABLE IF EXISTS `ghost_mission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ghost_mission` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `town_id` tinyint(4) NOT NULL DEFAULT '0' COMMENT '城镇id',
  `ghost_mission_key` int(11) NOT NULL DEFAULT '0' COMMENT '进入影界需求权值',
  `senior_ghost_rand` smallint(6) NOT NULL DEFAULT '0' COMMENT '抽中高级魂侍的概率',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='魂侍副本表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `ghost_mission` DISABLE KEYS */;
INSERT INTO `ghost_mission` VALUES (2,2,100110,200),(3,1,100000,200);
/*!40000 ALTER TABLE `ghost_mission` ENABLE KEYS */;
DROP TABLE IF EXISTS `ghost_passive_skill`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ghost_passive_skill` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `level` smallint(6) NOT NULL DEFAULT '0' COMMENT '等级',
  `name` varchar(200) NOT NULL DEFAULT '' COMMENT '被动技名称',
  `sign` varchar(200) NOT NULL DEFAULT '' COMMENT '图标标识',
  `desc` varchar(200) NOT NULL DEFAULT '' COMMENT '描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COMMENT='魂侍被动技能表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `ghost_passive_skill` DISABLE KEYS */;
INSERT INTO `ghost_passive_skill` VALUES (1,20,'魂侍护盾','HunShiShouHu','当生命少于30%时，自动施放吸收伤害的魂侍护盾，持续3回合。魂侍护盾吸收的伤害等于魂侍生命值的50%，与其他魂侍合计。魂侍护盾每次战斗只能触发1次。'),(2,30,'初始魂力','HunShiFaDongLv','初始魂力5，与其他魂侍合计'),(3,40,'魂侍技能2级','HunShiJiNengErJi','魂侍技能升级为2级'),(4,50,'魂侍技能2级','HunShiShouHuErJi','当生命少于30%时，自动施放吸收伤害的魂侍护盾，持续3回合。魂侍护盾吸收的伤害等于魂侍生命值的100%，与其他魂侍合计。魂侍护盾每次战斗只能触发1次。'),(5,60,'初始魂力2级','HunShiFaDongLvErJi','初始魂力10，与其他魂侍合计'),(6,70,'魂侍技能3级','HunShiJiNengSanJi','魂侍技能升级为3级');
/*!40000 ALTER TABLE `ghost_passive_skill` ENABLE KEYS */;
DROP TABLE IF EXISTS `ghost_star`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ghost_star` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `star` tinyint(4) NOT NULL DEFAULT '0' COMMENT '星级',
  `need_fragment_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '需要碎片数量',
  `growth` smallint(6) NOT NULL DEFAULT '0' COMMENT '成长值',
  `color` tinyint(4) NOT NULL DEFAULT '0' COMMENT '颜色',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COMMENT='魂侍进阶表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `ghost_star` DISABLE KEYS */;
INSERT INTO `ghost_star` VALUES (1,1,10,10,0),(2,2,30,30,1),(3,3,50,50,2),(4,4,80,80,3),(5,5,100,100,3);
/*!40000 ALTER TABLE `ghost_star` ENABLE KEYS */;
DROP TABLE IF EXISTS `ghost_tip`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ghost_tip` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `tip` varchar(512) DEFAULT NULL COMMENT '提示信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='魂侍提示信息';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `ghost_tip` DISABLE KEYS */;
INSERT INTO `ghost_tip` VALUES (1,'战斗中，角色每次行动与受伤都会获得魂力\r\n当魂力满100时，可以召唤魂侍'),(2,'每只魂侍在同个关卡内只能被触发一次');
/*!40000 ALTER TABLE `ghost_tip` ENABLE KEYS */;
DROP TABLE IF EXISTS `heart_draw`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `heart_draw` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `draw_type` tinyint(4) NOT NULL COMMENT '抽奖类型（1-大转盘；2-刮刮卡）',
  `daily_num` tinyint(4) NOT NULL DEFAULT '0' COMMENT '每日可抽奖次数',
  `cost_heart` tinyint(4) NOT NULL COMMENT '每次抽奖消耗爱心数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='爱心抽奖';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `heart_draw` DISABLE KEYS */;
INSERT INTO `heart_draw` VALUES (1,1,10,2);
/*!40000 ALTER TABLE `heart_draw` ENABLE KEYS */;
DROP TABLE IF EXISTS `heart_draw_award`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `heart_draw_award` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `heart_draw_id` smallint(6) NOT NULL COMMENT '爱心抽奖ID',
  `award_type` tinyint(4) NOT NULL COMMENT '奖品类型（1-铜钱；2-元宝；3-道具）',
  `award_num` smallint(6) NOT NULL COMMENT '奖品数量',
  `item_id` smallint(6) DEFAULT '0' COMMENT '道具奖品ID',
  `chance` tinyint(4) NOT NULL COMMENT '抽奖概率%',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='爱心抽奖奖品配置';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `heart_draw_award` DISABLE KEYS */;
INSERT INTO `heart_draw_award` VALUES (1,1,1,500,0,20),(2,1,2,20,0,10),(3,1,3,5,52,15),(4,1,3,5,200,12),(5,1,3,2,52,13),(6,1,3,1,212,3),(7,1,3,1,230,15),(8,1,3,5,30,12);
/*!40000 ALTER TABLE `heart_draw_award` ENABLE KEYS */;
DROP TABLE IF EXISTS `ingot_ghost_mission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ingot_ghost_mission` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `vip_level` tinyint(4) NOT NULL DEFAULT '0' COMMENT '可进入vip等级',
  `max_egg_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '一天开启魂蛋数量',
  `yellow_ghost_rand` smallint(6) NOT NULL DEFAULT '0' COMMENT '金色魂侍概率(万分之)',
  `purple_ghost_rand` smallint(6) NOT NULL DEFAULT '0' COMMENT '紫色魂侍概率(万分之)',
  `orange_ghost_rand` smallint(6) NOT NULL DEFAULT '0' COMMENT '橙色魂侍概率(万分之)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='极暗净土';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `ingot_ghost_mission` DISABLE KEYS */;
INSERT INTO `ingot_ghost_mission` VALUES (1,0,5,200,2500,7300);
/*!40000 ALTER TABLE `ingot_ghost_mission` ENABLE KEYS */;
DROP TABLE IF EXISTS `item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '物品ID',
  `type_id` int(11) NOT NULL COMMENT '类型ID',
  `quality` tinyint(4) DEFAULT NULL COMMENT '品质',
  `name` varchar(20) NOT NULL COMMENT '物品名称',
  `level` int(11) DEFAULT NULL COMMENT '需求等级',
  `desc` varchar(100) DEFAULT NULL COMMENT '物品描述',
  `price` bigint(20) NOT NULL DEFAULT '0' COMMENT '物品售价',
  `sign` varchar(30) DEFAULT NULL COMMENT '资源标识',
  `can_use` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否可在格子中使用，0不能，1反之',
  `panel` tinyint(4) NOT NULL DEFAULT '0' COMMENT '指向面板',
  `func_id` int(11) NOT NULL DEFAULT '0' COMMENT '使用的功能限制',
  `renew_ingot` int(11) NOT NULL DEFAULT '0' COMMENT '续费的元宝价格',
  `use_ingot` int(11) NOT NULL DEFAULT '0' COMMENT '使用的元宝价格',
  `valid_hours` int(11) NOT NULL DEFAULT '0' COMMENT '有效小时数',
  `equip_type_id` int(11) NOT NULL DEFAULT '0' COMMENT '装备等级类型',
  `health` int(11) DEFAULT '0' COMMENT '生命',
  `speed` int(11) DEFAULT '0' COMMENT '速度',
  `attack` int(11) DEFAULT '0' COMMENT '攻击',
  `defence` int(11) DEFAULT '0' COMMENT '防御',
  `equip_role_id` tinyint(4) NOT NULL DEFAULT '0' COMMENT '可装备角色ID',
  `appendix_num` tinyint(4) NOT NULL DEFAULT '0' COMMENT '追加属性数',
  `appendix_level` int(11) NOT NULL DEFAULT '0' COMMENT '追加属性等级',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=255 DEFAULT CHARSET=utf8mb4 COMMENT='物品';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `item` DISABLE KEYS */;
INSERT INTO `item` VALUES (1,2,1,'真气龙珠',1,'可用于提升角色武功境界等级',0,'ZhenQiLongZhu',1,1,0,0,0,0,0,0,0,0,0,0,0,0),(30,1,1,'龙币',1,'特殊货币，可用于购买真气龙珠，破界龙珠等',0,'LongBi',0,0,0,0,0,0,0,0,0,0,0,0,0,0),(31,3,1,'鹰扬剑',1,NULL,100,'YingYangJian',0,0,0,0,0,0,1,0,0,100,0,-1,0,0),(32,3,1,'鹰扬刀',1,NULL,100,'YingYangDao',0,0,0,0,0,0,1,0,0,100,0,3,0,0),(33,3,1,'鹰扬环',1,NULL,100,'YingYangHuan',0,0,0,0,0,0,1,0,0,100,0,4,0,0),(34,3,1,'鹰扬篮',1,NULL,100,'YingYangLan',0,0,0,0,0,0,1,0,0,100,0,5,0,0),(35,3,1,'鹰扬笔',1,NULL,100,'YingYangBi',0,0,0,0,0,0,1,0,0,100,0,6,0,0),(36,4,1,'鹰扬袍',1,NULL,100,'YingYangPao',0,0,0,0,0,0,1,0,0,0,50,0,0,0),(37,5,1,'鹰扬靴',1,NULL,100,'YingYangXue',0,0,0,0,0,0,1,0,50,0,0,0,0,0),(38,6,1,'鹰扬戒',1,NULL,100,'YingYangJie',0,0,0,0,0,0,1,200,0,0,0,0,0,0),(52,2,1,'影界果实',1,'用于培养魂侍',10,'YingJieGuoShi',0,3,0,0,0,0,0,0,0,0,0,0,0,0),(61,3,2,'朱雀剑',20,NULL,500,'ZhuQueJian',0,0,0,0,0,0,2,0,0,400,0,-1,3,2),(62,3,2,'朱雀刀',20,NULL,500,'ZhuQueDao',0,0,0,0,0,0,2,0,0,400,0,3,3,2),(63,3,2,'朱雀环',20,NULL,500,'ZhuQueHuan',0,0,0,0,0,0,2,0,0,400,0,4,3,2),(64,3,2,'朱雀篮',20,NULL,500,'ZhuQueLan',0,0,0,0,0,0,2,0,0,400,0,5,3,2),(65,3,2,'朱雀笔',20,NULL,500,'ZhuQueBi',0,0,0,0,0,0,2,0,0,400,0,6,3,2),(66,4,2,'朱雀袍',20,NULL,500,'ZhuQuePao',0,0,0,0,0,0,2,0,0,0,200,0,3,2),(67,5,2,'朱雀靴',20,NULL,500,'ZhuQueXue',0,0,0,0,0,0,2,0,100,0,0,0,3,2),(68,6,2,'朱雀戒',20,NULL,500,'ZhuQueJie',0,0,0,0,0,0,2,1000,0,0,0,0,3,2),(93,3,2,'白虎剑',50,NULL,1000,'BaiHuJian',0,0,0,0,0,0,3,0,0,1600,0,-1,3,3),(94,3,2,'白虎刀',50,NULL,1000,'BaiHuDao',0,0,0,0,0,0,3,0,0,1600,0,3,3,3),(95,3,2,'白虎环',50,NULL,1000,'BaiHuHuan',0,0,0,0,0,0,3,0,0,1600,0,4,3,3),(96,3,2,'白虎篮',50,NULL,1000,'BaiHuLan',0,0,0,0,0,0,3,0,0,1600,0,5,3,3),(97,3,2,'白虎笔',50,NULL,1000,'BaiHuBi',0,0,0,0,0,0,3,0,0,1600,0,6,3,3),(98,4,2,'白虎袍',50,NULL,1000,'BaiHuPao',0,0,0,0,0,0,3,0,0,0,800,0,3,3),(99,5,2,'白虎靴',50,NULL,1000,'BaiHuXue',0,0,0,0,0,0,3,0,400,0,0,0,3,3),(100,6,2,'白虎戒',50,NULL,1000,'BaiHuJie',0,0,0,0,0,0,3,4000,0,0,0,0,3,3),(125,3,2,'青龙剑',80,NULL,2000,'QingLongJian',0,0,0,0,0,0,5,0,0,3200,0,-1,3,4),(126,3,2,'青龙刀',80,NULL,2000,'QingLongDao',0,0,0,0,0,0,5,0,0,3200,0,3,3,4),(127,3,2,'青龙环',80,NULL,2000,'QingLongHuan',0,0,0,0,0,0,5,0,0,3200,0,4,3,4),(128,3,2,'青龙篮',80,NULL,2000,'QingLongLan',0,0,0,0,0,0,5,0,0,3200,0,5,3,4),(129,3,2,'青龙笔',80,NULL,2000,'QingLongBi',0,0,0,0,0,0,5,0,0,3200,0,6,3,4),(130,4,2,'青龙袍',80,NULL,2000,'QingLongPao',0,0,0,0,0,0,5,0,0,0,1600,0,3,4),(131,5,2,'青龙靴',80,NULL,2000,'QingLongXue',0,0,0,0,0,0,5,0,800,0,0,0,3,4),(132,6,2,'青龙戒',80,NULL,2000,'QingLongJie',0,0,0,0,0,0,5,8000,0,0,0,0,3,4),(157,3,3,'天罡剑',90,NULL,8000,'TianGangJian',0,0,0,0,0,0,6,0,0,8000,0,-1,4,4),(158,3,3,'天罡刀',90,NULL,8000,'TianGangDao',0,0,0,0,0,0,6,0,0,8000,0,3,4,4),(159,3,3,'天罡环',90,NULL,8000,'TianGangHuan',0,0,0,0,0,0,6,0,0,8000,0,4,4,4),(160,3,3,'天罡篮',90,NULL,8000,'TianGangLan',0,0,0,0,0,0,6,0,0,8000,0,5,4,4),(161,3,3,'天罡笔',90,NULL,8000,'TianGangBi',0,0,0,0,0,0,6,0,0,8000,0,6,4,4),(162,4,3,'天罡袍',90,NULL,8000,'TianGangPao',0,0,0,0,0,0,6,0,0,0,4000,0,4,4),(163,5,3,'天罡靴',90,NULL,8000,'TianGangXue',0,0,0,0,0,0,6,0,2000,0,0,0,4,4),(164,6,3,'天罡戒',90,NULL,8000,'TianGangJie',0,0,0,0,0,0,6,20000,0,0,0,0,4,4),(200,1,2,'魂侍碎片',1,'分解魂侍获得的碎片，可用来兑换魂侍',0,'HunShiSuiPian',1,2,0,0,0,0,0,0,0,0,0,0,0,0),(205,2,3,'一星破界龙珠',39,'可用突破金刚境界',0,'YiXingPoJieLongZhu',1,1,0,0,0,0,0,0,0,0,0,0,0,0),(206,2,3,'二星破界龙珠',59,'可用突破道玄境界',0,'ErXingPoJieLongZhu',1,1,0,0,0,0,0,0,0,0,0,0,0,0),(207,2,3,'三星破界龙珠',79,'可用突破万象境界',0,'SanXingPoJieLongZhu',1,1,0,0,0,0,0,0,0,0,0,0,0,0),(208,2,3,'四星破界龙珠',99,'可用突破天人境界',0,'SiXingPoJieLongZhu',1,1,0,0,0,0,0,0,0,0,0,0,0,0),(209,8,1,'止血草',0,'普通草药，恢复5000生命',200,'ZhiXueCao',1,0,0,0,0,0,0,0,0,0,0,0,0,0),(210,8,2,'金创药',NULL,'疗伤药品，恢复20000生命',400,'JinChuangYao',1,0,0,0,0,0,0,0,0,0,0,0,0,0),(211,8,2,'大还丹',NULL,'珍贵丹药，恢复99999生命',800,'DaHuanDan',1,0,0,0,0,0,0,0,0,0,0,0,0,0),(212,8,3,'凤凰羽毛',NULL,'从凤凰身上掉落的羽毛。复活并恢复100%生命',0,'FengHuangYuMao',1,0,0,0,0,0,0,0,0,0,0,0,0,0),(213,8,2,'飞刀',NULL,'对生命值最少的敌方造成10000点伤害',400,'FeiDao',1,0,0,0,0,0,0,0,0,0,0,0,0,0),(214,8,2,'暗影飞刀',NULL,'对生命值最少的敌方造成30000点伤害',500,'AnYingFeiDao',1,0,0,0,0,0,0,0,0,0,0,0,0,0),(215,8,2,'风的种子',NULL,'全体闪避等级增加1000，持续2回合',400,'FengDeZhongZi',1,0,0,0,0,0,0,0,0,0,0,0,0,0),(216,8,2,'火的种子',NULL,'精气增加10，下回合伤害提升60%',400,'HuoDeZhongZi',1,0,0,0,0,0,0,0,0,0,0,0,0,0),(217,8,2,'水的种子',NULL,'清除全体负面状态',400,'ShuiDeZhongZi',1,0,0,0,0,0,0,0,0,0,0,0,0,0),(219,3,3,'倚天剑',60,NULL,4000,'YiTianJian',0,0,0,0,0,0,4,0,0,4000,0,-1,4,3),(220,3,3,'倚天刀',60,NULL,4000,'YiTianDao',0,0,0,0,0,0,4,0,0,4000,0,3,4,3),(221,3,3,'倚天环',60,NULL,4000,'YiTianHuan',0,0,0,0,0,0,4,0,0,4000,0,4,4,3),(222,3,3,'倚天篮',60,NULL,4000,'YiTianLan',0,0,0,0,0,0,4,0,0,4000,0,5,4,3),(223,3,3,'倚天笔',60,NULL,4000,'YiTianBi',0,0,0,0,0,0,4,0,0,4000,0,6,4,3),(224,4,3,'倚天袍',60,NULL,4000,'YiTianPao',0,0,0,0,0,0,4,0,0,0,2000,0,4,3),(225,5,3,'倚天靴',60,NULL,4000,'YiTianXue',0,0,0,0,0,0,4,0,1000,0,0,0,4,3),(226,6,3,'倚天戒',60,NULL,4000,'YiTianJie',0,0,0,0,0,0,4,10000,0,0,0,0,4,3),(227,3,4,'黎明破晓',1,NULL,100000,'LiMingPoXiao',0,0,0,0,0,0,7,0,0,2000,0,-1,5,4),(230,9,1,'1级装备结晶箱',1,'开启后可获得装备结晶',100,'1JiZhuangBeiJieJingXiang',0,0,0,0,0,0,0,0,0,0,0,0,0,0),(231,2,1,'蓝色结晶',1,'分解蓝色品质装备获得，可用于洗炼和重铸装备',10,'LanSeJieJing',0,0,0,0,0,0,0,0,0,0,0,0,0,0),(232,2,2,'紫色结晶',1,'分解紫色品质装备获得，可用于洗炼和重铸装备',10,'ZiSeJieJing',0,0,0,0,0,0,0,0,0,0,0,0,0,0),(233,2,2,'金色结晶',1,'分解金色品质装备获得，可用于洗炼和重铸装备',10,'JinSeJieJing',0,0,0,0,0,0,0,0,0,0,0,0,0,0),(234,2,4,'橙色结晶',1,'分解橙色品质装备获得，可用于洗炼和重铸装备',10,'ChengSeJieJing',0,0,0,0,0,0,0,0,0,0,0,0,0,0),(235,2,2,'武器碎片',1,'分解武器获得，可用于洗炼和重铸武器',0,'WuQiSuiPian',0,0,0,0,0,0,0,0,0,0,0,0,0,0),(236,2,2,'护甲碎片',1,'分解护甲获得，可用于洗炼和重铸护甲',10,'HuJiaSuiPian',0,0,0,0,0,0,0,0,0,0,0,0,0,0),(237,2,2,'长靴碎片',1,'分解长靴获得，可用于洗炼和重铸长靴',10,'ChangXueSuiPian',0,0,0,0,0,0,0,0,0,0,0,0,0,0),(238,2,2,'饰品碎片',1,'分解饰品获得，可用于洗炼和重铸饰品',10,'ShiPinSuiPian',0,0,0,0,0,0,0,0,0,0,0,0,0,0),(239,9,1,'1级装备碎片箱',1,'开启后可获得装备碎片',100,'1JiZhuangBeiSuiPianXiang',0,0,0,0,0,0,0,0,0,0,0,0,0,0),(241,1,3,'光明钥匙',1,'充满了光明之力的钥匙，可开启被阴影笼罩的区域。',0,'GuangMingYaoShi',0,0,0,0,0,0,0,0,0,0,0,0,0,0),(242,10,3,'元宝',0,'',0,'YuanBao',0,0,0,0,0,0,0,0,0,0,0,0,0,0),(243,10,0,'经验',0,'',0,'JingYan',0,0,0,0,0,0,0,0,0,0,0,0,0,0),(244,10,0,'铜钱',0,'',0,'TongQian',0,0,0,0,0,0,0,0,0,0,0,0,0,0),(245,10,1,'经验双倍',0,'',0,'JingYanShuangBei',0,0,0,0,0,0,0,0,0,0,0,0,0,0),(246,10,1,'铜钱双倍',0,'',0,'TongQianShuangBei',0,0,0,0,0,0,0,0,0,0,0,0,0,0),(247,10,0,'爱心',0,'',0,'AiXin',0,0,0,0,0,0,0,0,0,0,0,0,0,0),(248,10,1,'体力',0,'',0,'TiLi',0,0,0,0,0,0,0,0,0,0,0,0,0,0),(249,1,2,'剑心碎片',1,'剑山拔剑获得的碎片，可用来兑换剑心	',0,'JianXinSuiPian',1,4,0,0,0,0,0,0,0,0,0,0,0,0),(250,8,1,'契约球',NULL,'有很大几率能成功捕获灵宠',1500,'QiYueQiu',1,0,0,0,0,0,0,0,0,0,0,0,0,0),(251,8,3,'至尊契约球',NULL,'一定能成功捕获灵宠',5000,'ZhiZunQiYueQiu',1,0,0,0,0,0,0,0,0,0,0,0,0,0),(252,11,1,'契约球(草精)',NULL,'与草精达成召唤契约的球',100,'QiYueQiuCaoJing',0,0,0,0,0,0,0,0,0,0,0,0,0,0),(253,11,1,'契约球(火灵)',NULL,'与火灵达成召唤契约的球',100,'QiYueQiuHuoLing',0,0,0,0,0,0,0,0,0,0,0,0,0,0),(254,11,2,'契约球(剑魄)',NULL,'与剑魄达成召唤契约的球',500,'QiYueQiuJianPo',0,0,0,0,0,0,0,0,0,0,0,0,0,0);
/*!40000 ALTER TABLE `item` ENABLE KEYS */;
DROP TABLE IF EXISTS `item_box_content`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item_box_content` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `item_id` smallint(6) NOT NULL COMMENT '物品宝箱的ID',
  `type` tinyint(4) NOT NULL COMMENT '类型，0铜钱，1元宝，2物品',
  `mode` tinyint(4) NOT NULL COMMENT '随机方式，0直接获得，1概率数量，2概率获得',
  `get_item_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '得到的物品ID',
  `item_id_set` text COMMENT '随机的物品ID集',
  `item_desc` varchar(50) DEFAULT NULL COMMENT '随机物品集的描述',
  `min_num` int(11) NOT NULL DEFAULT '0' COMMENT '最少数量',
  `max_num` int(11) NOT NULL DEFAULT '0' COMMENT '最多数量',
  `probability` tinyint(4) NOT NULL DEFAULT '0' COMMENT '概率',
  PRIMARY KEY (`id`),
  KEY `idx_item_id` (`item_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COMMENT='宝箱内容';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `item_box_content` DISABLE KEYS */;
INSERT INTO `item_box_content` VALUES (1,230,2,0,231,'','',2,0,0),(2,230,2,0,232,'','',1,0,0),(3,239,2,0,235,'','',1,0,0),(4,239,2,0,236,'','',1,0,0),(5,239,2,0,237,'','',1,0,0),(6,239,2,0,238,'','',1,0,0);
/*!40000 ALTER TABLE `item_box_content` ENABLE KEYS */;
DROP TABLE IF EXISTS `item_exchange`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item_exchange` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `target_item_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '目标物品id',
  `item_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '物品id',
  `item_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '物品数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='物品兑换表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `item_exchange` DISABLE KEYS */;
INSERT INTO `item_exchange` VALUES (1,1,30,1),(2,205,30,6),(3,206,30,12),(4,207,30,18),(5,208,30,24);
/*!40000 ALTER TABLE `item_exchange` ENABLE KEYS */;
DROP TABLE IF EXISTS `item_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item_type` (
  `id` tinyint(4) NOT NULL AUTO_INCREMENT COMMENT '类型ID',
  `name` varchar(10) NOT NULL COMMENT '类型名称',
  `max_num_in_pos` smallint(6) NOT NULL DEFAULT '1' COMMENT '每个位置最多可堆叠的数量',
  `sign` varchar(50) DEFAULT '' COMMENT '类型标志',
  `order` int(11) DEFAULT '0' COMMENT '客户端排序权重',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COMMENT='物品类型';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `item_type` DISABLE KEYS */;
INSERT INTO `item_type` VALUES (1,'特殊材料',999,'SpecialMaterial',7),(2,'材料',99,'Stuff',6),(3,'武器',1,'Weapon',1),(4,'战袍',1,'Clothes',2),(5,'靴子',1,'Shoe',3),(6,'饰品',1,'Accessories',4),(7,'功能道具',99,'FuncProp',8),(8,'战斗道具',99,'FightProp',5),(9,'礼包',99,'Chest',9),(10,'资源',999,'Resource',10),(11,'灵宠契约球',99,'Pet',11);
/*!40000 ALTER TABLE `item_type` ENABLE KEYS */;
DROP TABLE IF EXISTS `level_battle_pet`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `level_battle_pet` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `mission_enemy_id` int(10) unsigned NOT NULL COMMENT '关卡怪物组',
  `battle_pet_id` smallint(6) NOT NULL COMMENT '灵宠ID',
  `rate` tinyint(4) NOT NULL COMMENT '出现概率%',
  `live_round` tinyint(4) NOT NULL COMMENT '出现后存活回合数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COMMENT='关卡灵宠配置';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `level_battle_pet` DISABLE KEYS */;
INSERT INTO `level_battle_pet` VALUES (1,64,2,60,3),(2,65,2,60,3),(3,66,2,60,3),(4,82,3,30,2),(5,83,3,30,2),(6,100,3,30,2),(7,58,2,100,3),(8,59,2,60,3),(9,60,2,60,3),(10,76,3,100,2),(11,77,3,30,2),(12,99,3,30,2),(13,40,1,100,3);
/*!40000 ALTER TABLE `level_battle_pet` ENABLE KEYS */;
DROP TABLE IF EXISTS `mail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mail` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '邮件ID',
  `sign` varchar(30) DEFAULT NULL COMMENT '唯一标识',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型',
  `title` varchar(30) NOT NULL COMMENT '标题',
  `parameters` varchar(1024) NOT NULL COMMENT '参数',
  `content` varchar(1024) NOT NULL COMMENT '内容',
  PRIMARY KEY (`id`),
  UNIQUE KEY `sign` (`sign`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COMMENT='系统邮件模板';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mail` DISABLE KEYS */;
INSERT INTO `mail` VALUES (1,'BagFull',0,'背包已满提示','func,功能','您的背包已满，系统已自动将{0}的物品暂存至附件，请及时点击领取。'),(2,'Heart',0,'爱心赠送邮件','who,发送者','您的好友{0}赠送您一颗爱心哦！请及时点击领取。'),(3,'TestMail',0,'测试邮件','p1, 参数1; p2, 参数2','我{0}只是一封测试邮件{1}，虽然偶只是测试用的，但四，你八可以八理我，if你不点我我that会伤心的。so !选我！选我！选我！'),(4,'MultiLevel',0,'多人关卡战斗奖励','name,关卡名称','您在参与的多人关卡{0}中取得了胜利，附件里是您获得的奖励。'),(5,'Welcome',0,'欢迎进入仙侠道','','我们诚挚欢迎您加入仙侠道的武侠世界，可以在这里一圆您的武侠梦！特此送上我们的小礼物，快去发觉他的妙用吧~'),(6,'GhostBagFull',0,'魂侍背包已满提示','func,功能','您的魂侍背包已满，系统已自动将您刚抽取到的{0}个魂侍暂存至附件，请及时点击领取。'),(7,'SwordSoulBagFull',0,'剑心背包已满提示','func,功能','您的剑心背包已满，系统已自动将您刚抽取到的{0}个剑心暂存至附件，请及时点击领取。');
/*!40000 ALTER TABLE `mail` ENABLE KEYS */;
DROP TABLE IF EXISTS `mail_attachments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mail_attachments` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '邮件ID',
  `mail_id` int(11) NOT NULL COMMENT 'mail表主键',
  `item_id` smallint(6) NOT NULL COMMENT '物品',
  `attachment_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '附件类型',
  `item_num` int(11) NOT NULL DEFAULT '0' COMMENT '数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COMMENT='系统邮件附件';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mail_attachments` DISABLE KEYS */;
INSERT INTO `mail_attachments` VALUES (3,3,1,0,1),(4,3,31,0,1),(5,3,238,0,10),(6,2,0,3,1),(7,3,94,0,2),(8,3,128,0,1),(9,5,1,0,10);
/*!40000 ALTER TABLE `mail_attachments` ENABLE KEYS */;
DROP TABLE IF EXISTS `mission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mission` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '区域ID',
  `town_id` smallint(6) NOT NULL COMMENT '城镇ID',
  `keys` int(11) NOT NULL COMMENT '开启钥匙数',
  `name` varchar(10) NOT NULL COMMENT '区域名称',
  `sign` varchar(50) NOT NULL COMMENT '资源标识',
  `order` tinyint(4) NOT NULL COMMENT '区域开启顺序',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COMMENT='城镇区域';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mission` DISABLE KEYS */;
INSERT INTO `mission` VALUES (1,1,0,'青竹林','QingZhuLin',1),(2,1,8,'黑夜森林','HeiYeSenLin',2),(3,1,8,'莲花峰','LianHuaFeng',3),(4,1,8,'熔岩火山','RongYanHuoShan',4),(5,1,10,'剑灵密室','JianLingMiShi',5),(6,3,0,'测试区域','QingZhuLin',1),(7,3,0,'作废','QingZhuLin',2);
/*!40000 ALTER TABLE `mission` ENABLE KEYS */;
DROP TABLE IF EXISTS `mission_enemy`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mission_enemy` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `mission_level_id` int(11) NOT NULL COMMENT '副本关卡id',
  `monster_num` tinyint(4) NOT NULL COMMENT '怪物数量',
  `enter_x` int(11) NOT NULL COMMENT '出生点x坐标',
  `enter_y` int(11) NOT NULL COMMENT '出生点y坐标',
  `monster1_id` int(11) NOT NULL COMMENT '怪物1 ID',
  `monster1_chance` tinyint(4) NOT NULL COMMENT '出现概率',
  `monster2_id` int(11) NOT NULL COMMENT '怪物2 ID',
  `monster2_chance` tinyint(4) NOT NULL COMMENT '出现概率',
  `monster3_id` int(11) NOT NULL COMMENT '怪物3 ID',
  `monster4_id` int(11) NOT NULL COMMENT '怪物4 ID',
  `monster4_chance` tinyint(4) NOT NULL COMMENT '出现概率',
  `monster5_id` int(11) NOT NULL COMMENT '怪物5 ID',
  `monster5_chance` tinyint(4) NOT NULL COMMENT '出现概率',
  `monster3_chance` tinyint(4) NOT NULL COMMENT '出现概率',
  `is_boss` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否是boss 0否,1是',
  `order` tinyint(4) NOT NULL COMMENT '顺序',
  `talk` varchar(200) NOT NULL DEFAULT '' COMMENT '副本对话（怪物旁白）',
  `boss_dir` tinyint(4) NOT NULL COMMENT '怪物朝向(0--左;1--右)',
  `best_round` tinyint(4) NOT NULL COMMENT '最好的通关回合数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=201 DEFAULT CHARSET=utf8mb4 COMMENT='副本敌人';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mission_enemy` DISABLE KEYS */;
INSERT INTO `mission_enemy` VALUES (1,1,2,656,741,31,20,33,80,0,0,0,0,0,0,0,1,'',0,0),(2,2,2,851,565,0,0,0,0,0,0,0,0,0,0,0,1,'',0,0),(10,4,2,605,377,31,100,0,0,0,0,0,0,0,0,0,1,'',0,0),(11,4,2,1217,312,31,100,0,0,0,0,0,0,0,0,0,2,'',0,0),(13,5,2,810,318,0,0,0,20,0,0,0,0,0,0,0,1,'',0,0),(14,5,2,1545,502,0,0,0,20,0,0,0,0,0,0,0,2,'',0,0),(18,6,3,1185,474,0,0,0,0,0,0,0,0,0,0,1,1,'',0,0),(19,7,3,821,922,34,60,35,30,36,0,0,0,0,10,0,1,'',0,0),(20,7,3,1492,350,34,60,35,30,36,0,0,0,0,10,0,2,'',0,0),(22,8,3,959,508,5,70,6,20,7,0,0,0,0,10,0,1,'',0,0),(23,8,3,1537,918,5,70,6,20,7,0,0,0,0,10,0,2,'',0,0),(27,9,3,1149,479,0,0,0,0,0,0,0,0,0,100,1,1,'',0,0),(28,10,3,760,590,37,60,35,20,36,0,0,0,0,20,0,1,'',0,0),(29,10,3,1531,654,37,60,35,20,36,0,0,0,0,20,0,2,'',0,0),(31,11,3,810,669,24,70,6,20,7,0,0,0,0,10,0,1,'',0,0),(32,11,3,1595,556,24,70,6,20,7,0,0,0,0,10,0,2,'',0,0),(36,12,3,900,480,0,0,0,0,0,0,0,0,0,0,1,1,'',0,0),(37,13,3,876,477,38,10,39,60,40,0,0,0,0,30,0,1,'',0,0),(38,13,3,1620,894,38,10,39,60,40,0,0,0,0,30,0,2,'',0,0),(39,13,3,984,1300,38,10,39,60,40,0,0,0,0,30,0,3,'',0,0),(40,14,3,876,477,10,10,9,70,1,0,0,0,0,20,0,1,'',0,0),(41,14,3,1620,894,10,10,9,70,1,0,0,0,0,20,0,2,'',0,0),(42,14,3,984,1300,10,10,9,70,1,0,0,0,0,20,0,3,'',0,0),(45,15,3,796,512,0,0,0,0,0,0,0,0,0,0,1,1,'',1,0),(46,16,3,818,235,41,60,42,40,0,0,0,0,0,0,0,1,'',0,0),(47,16,3,939,873,41,60,42,40,0,0,0,0,0,0,0,2,'',0,0),(49,17,3,1103,354,27,50,13,50,0,0,0,0,0,0,0,1,'',0,0),(50,17,3,1010,998,27,50,13,50,0,0,0,0,0,0,0,2,'',0,0),(54,18,3,1113,487,0,0,0,0,0,0,0,0,0,0,1,1,'',0,0),(55,19,3,422,569,44,40,43,50,45,0,0,0,0,10,0,1,'',0,0),(56,19,3,870,305,44,40,43,50,45,0,0,0,0,10,0,2,'',0,0),(57,19,3,1376,530,44,40,43,50,45,0,0,0,0,10,0,3,'',0,0),(58,20,3,1339,708,16,40,15,50,17,0,0,0,0,10,0,1,'',0,0),(59,20,3,566,1162,16,40,15,50,17,0,0,0,0,10,0,2,'',0,0),(60,20,3,1083,1651,16,40,15,50,17,0,0,0,0,10,0,3,'',0,0),(63,21,3,1034,526,0,0,0,0,0,0,0,0,0,0,1,1,'',0,0),(64,23,3,1023,762,16,40,15,50,17,0,0,0,0,10,0,1,'',0,0),(65,23,3,1740,1218,16,40,15,50,17,0,0,0,0,10,0,2,'',0,0),(66,23,3,900,1839,16,40,15,50,17,0,0,0,0,10,0,3,'',0,0),(67,22,3,422,569,44,40,43,50,45,0,0,0,0,10,0,1,'',0,0),(68,22,3,870,305,44,40,43,50,45,0,0,0,0,10,0,2,'',0,0),(69,22,3,1376,530,44,40,43,50,45,0,0,0,0,10,0,3,'',0,0),(72,24,3,880,510,0,0,0,0,0,0,0,0,0,0,1,1,'',1,0),(73,25,3,753,363,46,70,47,20,48,0,0,0,0,10,0,1,'',0,0),(74,25,3,1592,314,46,70,47,20,48,0,0,0,0,10,0,2,'',0,0),(76,26,3,1547,363,21,70,20,20,22,0,0,0,0,10,0,1,'',0,0),(77,26,3,708,314,21,70,20,20,22,0,0,0,0,10,0,2,'',0,0),(79,28,3,753,363,46,70,47,20,48,0,0,0,0,10,0,1,'',0,0),(80,28,3,1592,314,46,70,47,20,48,0,0,0,0,10,0,2,'',0,0),(82,29,3,1547,363,21,70,20,20,22,0,0,0,0,10,0,1,'',0,0),(83,29,3,708,314,21,70,20,20,22,0,0,0,0,10,0,2,'',0,0),(87,27,3,1302,637,0,0,0,0,0,0,0,0,0,0,1,1,'',0,0),(90,30,3,618,637,0,0,0,0,0,0,0,0,0,0,1,1,'',1,0),(93,3,3,981,470,0,0,0,0,0,0,0,0,0,0,1,1,'',0,0),(94,1,2,1051,668,31,20,33,80,0,0,0,0,0,0,0,2,'',0,0),(95,2,2,1264,741,0,0,0,0,0,0,0,0,0,0,0,2,'',0,0),(96,31,1,0,0,12,100,0,0,0,0,0,0,0,0,0,2,'',0,0),(97,25,3,1019,914,46,70,47,20,48,0,0,0,0,10,0,3,'',0,0),(98,28,3,1019,914,46,70,47,20,48,0,0,0,0,10,0,3,'',0,0),(99,26,3,1281,914,21,70,20,20,22,0,0,0,0,10,0,3,'',0,0),(100,29,3,1281,914,21,70,20,20,22,0,0,0,0,10,0,3,'',0,0),(101,32,2,605,377,31,100,0,0,0,0,0,0,0,0,0,1,'',0,0),(102,32,2,1217,312,31,100,0,0,0,0,0,0,0,0,0,2,'',0,0),(103,33,2,656,741,31,20,33,80,0,0,0,0,0,0,0,1,'',0,0),(104,33,2,1051,668,31,20,33,80,0,0,0,0,0,0,0,2,'',0,0),(105,31,1,0,0,49,100,0,0,0,0,0,0,0,0,0,1,'',0,0),(106,31,1,0,0,50,100,0,0,0,0,0,0,0,0,0,17,'',0,0),(107,31,1,0,0,51,100,0,0,0,0,0,0,0,0,0,21,'',0,0),(108,31,1,0,0,52,100,0,0,0,0,0,0,0,0,0,25,'',0,0),(109,31,1,0,0,53,100,0,0,0,0,0,0,0,0,0,29,'',0,0),(110,31,1,0,0,54,100,0,0,0,0,0,0,0,0,0,33,'',0,0),(111,31,1,0,0,55,100,0,0,0,0,0,0,0,0,0,37,'',0,0),(112,31,1,0,0,56,100,0,0,0,0,0,0,0,0,0,3,'',0,0),(113,31,1,0,0,57,100,0,0,0,0,0,0,0,0,0,4,'',0,0),(114,31,1,0,0,58,100,0,0,0,0,0,0,0,0,0,5,'',0,0),(115,31,1,0,0,59,100,0,0,0,0,0,0,0,0,0,6,'',0,0),(116,31,1,0,0,60,100,0,0,0,0,0,0,0,0,0,7,'',0,0),(117,31,1,0,0,61,100,0,0,0,0,0,0,0,0,0,14,'',0,0),(118,31,1,0,0,62,100,0,0,0,0,0,0,0,0,0,15,'',0,0),(119,31,1,0,0,67,100,0,0,0,0,0,0,0,0,0,16,'',0,0),(120,31,1,0,0,63,100,0,0,0,0,0,0,0,0,0,8,'',0,0),(121,31,1,0,0,64,100,0,0,0,0,0,0,0,0,0,9,'',0,0),(122,31,1,0,0,65,100,0,0,0,0,0,0,0,0,0,10,'',0,0),(123,31,1,0,0,66,100,0,0,0,0,0,0,0,0,0,11,'',0,0),(124,31,1,0,0,68,100,0,0,0,0,0,0,0,0,0,12,'',0,0),(125,31,1,0,0,69,100,0,0,0,0,0,0,0,0,0,13,'',0,0),(126,31,1,0,0,70,100,0,0,0,0,0,0,0,0,0,18,'',0,0),(127,31,1,0,0,76,100,0,0,0,0,0,0,0,0,0,19,'',0,0),(128,31,1,0,0,82,100,0,0,0,0,0,0,0,0,0,20,'',0,0),(129,31,1,0,0,71,100,0,0,0,0,0,0,0,0,0,22,'',0,0),(130,31,1,0,0,77,100,0,0,0,0,0,0,0,0,0,23,'',0,0),(131,31,1,0,0,83,100,0,0,0,0,0,0,0,0,0,24,'',0,0),(132,31,1,0,0,72,100,0,0,0,0,0,0,0,0,0,26,'',0,0),(133,31,1,0,0,78,100,0,0,0,0,0,0,0,0,0,27,'',0,0),(134,31,1,0,0,84,100,0,0,0,0,0,0,0,0,0,28,'',0,0),(135,31,1,0,0,73,100,0,0,0,0,0,0,0,0,0,30,'',0,0),(136,31,1,0,0,79,100,0,0,0,0,0,0,0,0,0,31,'',0,0),(137,31,1,0,0,85,100,0,0,0,0,0,0,0,0,0,32,'',0,0),(138,31,1,0,0,74,100,0,0,0,0,0,0,0,0,0,34,'',0,0),(139,31,1,0,0,80,100,0,0,0,0,0,0,0,0,0,35,'',0,0),(140,31,1,0,0,86,100,0,0,0,0,0,0,0,0,0,36,'',0,0),(141,31,1,0,0,75,100,0,0,0,0,0,0,0,0,0,38,'',0,0),(142,31,1,0,0,81,100,0,0,0,0,0,0,0,0,0,39,'',0,0),(143,31,1,0,0,87,100,0,0,0,0,0,0,0,0,0,40,'',0,0),(144,36,3,821,922,34,60,35,30,36,0,0,0,0,10,0,1,'',0,0),(145,36,3,1492,350,34,60,35,30,36,0,0,0,0,10,0,2,'',0,0),(146,37,3,760,590,37,60,35,20,36,0,0,0,0,20,0,1,'',0,0),(147,37,3,1531,654,37,60,35,20,36,0,0,0,0,20,0,2,'',0,0),(148,38,5,900,450,88,100,0,0,0,0,0,0,0,0,0,1,'',0,0),(149,39,5,900,450,89,100,0,0,0,0,0,0,0,0,0,1,'',0,0),(150,40,3,876,477,38,10,39,60,40,0,0,0,0,30,0,1,'',0,0),(151,40,3,1620,894,38,10,39,60,40,0,0,0,0,30,0,2,'',0,0),(152,40,3,984,1300,38,10,39,60,40,0,0,0,0,30,0,3,'',0,0),(153,41,3,818,235,41,60,42,40,0,0,0,0,0,0,0,1,'',0,0),(154,41,3,939,873,41,60,42,40,0,0,0,0,0,0,0,2,'',0,0),(155,42,3,1023,762,44,40,43,50,45,0,0,0,0,10,0,1,'',0,0),(156,42,3,1740,1218,44,40,43,50,45,0,0,0,0,10,0,2,'',0,0),(157,42,3,900,1839,44,40,43,50,45,0,0,0,0,10,0,3,'',0,0),(158,43,3,1339,708,44,40,43,50,45,0,0,0,0,10,0,1,'',0,0),(159,43,3,566,1162,44,40,43,50,45,0,0,0,0,10,0,2,'',0,0),(160,43,3,1083,1651,44,40,43,50,45,0,0,0,0,10,0,3,'',0,0),(161,44,3,753,363,46,70,47,20,48,0,0,0,0,10,0,1,'',0,0),(162,44,3,1592,314,46,70,47,20,48,0,0,0,0,10,0,2,'',0,0),(163,44,3,1019,914,46,70,47,20,48,0,0,0,0,10,0,3,'',0,0),(164,45,3,753,363,46,70,47,20,48,0,0,0,0,10,0,1,'',0,0),(165,45,3,1592,314,46,70,47,20,48,0,0,0,0,10,0,2,'',0,0),(166,45,3,1019,914,46,70,47,20,48,0,0,0,0,10,0,3,'',0,0),(167,46,1,100,100,90,100,0,0,0,0,0,0,0,0,0,1,'',0,0),(168,51,5,900,450,88,100,0,0,0,0,0,0,0,0,0,1,'',0,0),(169,52,5,900,450,89,100,0,0,0,0,0,0,0,0,0,1,'',0,0),(170,38,3,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,'',0,0),(171,47,5,900,450,88,100,0,0,0,0,0,0,0,0,0,1,'',0,0),(172,47,3,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,'',0,0),(173,49,5,900,450,88,100,0,0,0,0,0,0,0,0,0,1,'',0,0),(174,49,3,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,'',0,0),(175,59,5,900,450,88,100,0,0,0,0,0,0,0,0,0,1,'',0,0),(176,59,1,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,'',0,0),(177,60,5,900,450,89,100,0,0,0,0,0,0,0,0,0,1,'',0,0),(178,60,1,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,'',0,0),(179,51,3,1400,400,0,0,0,0,0,0,0,0,0,0,1,2,'',0,0),(180,39,1,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,'',0,0),(181,53,5,900,450,88,100,0,0,0,0,0,0,0,0,0,1,'',0,0),(182,53,3,1400,400,0,0,0,0,0,0,0,0,0,0,1,2,'',0,0),(183,55,5,900,450,88,100,0,0,0,0,0,0,0,0,0,1,'',0,0),(184,55,3,1400,400,0,0,0,0,0,0,0,0,0,0,1,2,'',0,0),(185,57,5,900,450,88,100,0,0,0,0,0,0,0,0,0,1,'',0,0),(186,57,3,1400,400,0,0,0,0,0,0,0,0,0,0,1,2,'',0,0),(187,61,5,900,450,88,100,0,0,0,0,0,0,0,0,0,1,'',0,0),(188,61,3,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,'',0,0),(189,58,5,900,450,88,100,0,0,0,0,0,0,0,0,0,1,'',0,0),(190,58,1,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,'',0,0),(191,48,5,900,450,89,100,0,0,0,0,0,0,0,0,0,1,'',0,0),(192,48,1,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,'',0,0),(193,50,5,900,450,89,100,0,0,0,0,0,0,0,0,0,1,'',0,0),(194,50,1,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,'',0,0),(195,52,1,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,'',0,0),(196,54,5,900,450,89,100,0,0,0,0,0,0,0,0,0,1,'',0,0),(197,54,1,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,'',0,0),(198,56,5,900,450,89,100,0,0,0,0,0,0,0,0,0,1,'',0,0),(199,56,1,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,'',0,0),(200,46,1,100,100,94,100,0,0,0,0,0,0,0,0,0,2,'',0,0);
/*!40000 ALTER TABLE `mission_enemy` ENABLE KEYS */;
DROP TABLE IF EXISTS `mission_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mission_level` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '区域关卡ID',
  `mission_id` smallint(6) NOT NULL COMMENT '区域ID',
  `parent_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '关联关卡类型(0--无;1--资源关卡;2--通天塔)',
  `parent_id` smallint(6) DEFAULT '0' COMMENT '关联关卡的外键',
  `lock` int(11) NOT NULL COMMENT '关卡开启的权值',
  `name` varchar(10) NOT NULL COMMENT '关卡名称',
  `type` tinyint(4) NOT NULL COMMENT '关卡类型(0--普通;1--精英;2--Boss)',
  `daily_num` tinyint(4) NOT NULL COMMENT '允许每天进入次数,0表示不限制',
  `physical` tinyint(4) NOT NULL COMMENT '每次进入消耗的体力',
  `box_x` int(11) NOT NULL COMMENT '宝箱x坐标',
  `box_y` int(11) NOT NULL COMMENT '宝箱y坐标',
  `award_key` int(11) NOT NULL COMMENT '奖励钥匙数',
  `award_exp` int(11) NOT NULL COMMENT '奖励经验',
  `enter_y` int(11) NOT NULL COMMENT '出生点y坐标',
  `enter_x` int(11) NOT NULL COMMENT '出生点x坐标',
  `sign` varchar(50) NOT NULL COMMENT '资源标识',
  `sign_war` varchar(50) NOT NULL COMMENT '关卡战斗资源标识',
  `music` varchar(20) NOT NULL COMMENT '音乐资源标识',
  `award_lock` int(11) NOT NULL COMMENT '通关奖励权值',
  `box_dir` tinyint(4) NOT NULL COMMENT '宝箱朝向(0--左;1--右)',
  `award_coin` int(11) NOT NULL COMMENT '奖励铜钱',
  `award_box` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否奖励宝箱',
  `flip_horizontal` tinyint(4) NOT NULL COMMENT '水平翻转',
  `sub_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '关卡子类型(0--无;1--铜钱关卡;2--经验关卡)',
  PRIMARY KEY (`id`),
  KEY `idx_mission_id` (`mission_id`)
) ENGINE=InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET=utf8mb4 COMMENT='区域关卡配置';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mission_level` DISABLE KEYS */;
INSERT INTO `mission_level` VALUES (1,1,0,0,100130,'青竹林五',0,-1,6,1398,812,1,10,1114,176,'QingZhuLin5','QingZhuLin8','PingHe',100140,0,100,1,0,0),(2,1,0,0,100140,'青竹林六',1,3,12,1698,1103,1,15,813,509,'QingZhuLin5','QingZhuLin8','PingHe',100145,0,200,1,1,0),(3,1,0,0,100150,'青竹林八',2,2,20,1374,482,1,20,618,190,'QingZhuLin8','QingZhuLin8','PingHe',110100,0,300,1,0,0),(4,1,0,0,100100,'青竹林一',0,-1,6,1769,343,1,10,828,144,'QingZhuLin1','QingZhuLin4','PingHe',100110,0,100,1,0,0),(5,1,0,0,100110,'青竹林二',1,3,12,1746,827,1,15,328,109,'QingZhuLin1','QingZhuLin4','PingHe',100115,0,200,1,1,0),(6,1,0,0,100120,'青竹林四',2,2,20,1528,489,1,20,428,366,'QingZhuLin4','QingZhuLin4','PingHe',100130,0,300,1,0,0),(7,2,0,0,110100,'黑夜森林一',0,-1,6,2048,268,1,15,856,114,'HeiYeSenLin1','HeiYeSenLin1War','YinSen',110110,0,150,1,0,0),(8,2,0,0,110110,'黑夜森林二',1,3,12,2167,866,1,30,242,275,'HeiYeSenLin1','HeiYeSenLin1War','YinSen',110115,0,300,1,1,0),(9,2,0,0,110120,'黑夜森林四',2,2,20,1419,606,1,45,388,140,'HeiYeSenLin4','HeiYeSenLin1War','YinSen',110130,0,450,1,0,0),(10,2,0,0,110130,'暗影秘境一',0,-1,6,2075,245,1,15,286,152,'AnYingMiJing1','AnYingMiJing4','YinSen',110140,0,150,1,0,0),(11,2,0,0,110140,'暗影秘境二',1,3,12,2155,312,1,30,155,171,'AnYingMiJing1','AnYingMiJing4','YinSen',110145,0,300,1,1,0),(12,2,0,0,110150,'暗影秘境四',2,2,20,1189,613,1,45,124,336,'AnYingMiJing4','AnYingMiJing4','YinSen',120100,0,450,1,0,0),(13,3,0,0,120100,'莲花峰一',0,-1,6,712,1460,1,40,369,126,'LianHuaFeng1','LianHuaFeng4','YouQu',120110,1,200,1,0,0),(14,3,0,0,120110,'莲花峰二',1,3,12,712,1460,1,80,369,126,'LianHuaFeng1','LianHuaFeng4','YouQu',120115,1,400,1,0,0),(15,3,0,0,120120,'莲花峰四',2,2,20,586,611,1,120,122,1486,'LianHuaFeng4','LianHuaFeng4','YouQu',120130,1,600,1,0,0),(16,3,0,0,120130,'水溶洞一',0,-1,6,693,1322,1,40,306,150,'ShuiRongDong1','ShuiRongDong4','YinSen',120140,1,200,1,0,0),(17,3,0,0,120140,'水溶洞二',1,3,12,1287,1314,1,80,296,1844,'ShuiRongDong1','ShuiRongDong4','YinSen',120145,0,400,1,1,0),(18,3,0,0,120150,'水溶洞四',2,2,20,1314,722,1,120,362,250,'ShuiRongDong4','ShuiRongDong4','YinSen',130100,0,600,1,0,0),(19,4,0,0,130100,'熔岩火山一',0,-1,6,1742,764,1,125,814,104,'RongYanHuoShan1','RongYanHuoShan2War','YinSen',130110,0,400,1,0,0),(20,4,0,0,130110,'熔岩火山二',1,3,12,1917,1952,1,250,589,2102,'RongYanHuoShan2','RongYanHuoShan2War','YinSen',130115,0,800,1,1,0),(21,4,0,0,130120,'熔岩火山四',2,2,20,1544,513,1,375,189,169,'RongYanHuoShan4','RongYanHuoShan2War','YinSen',130130,0,1200,1,1,0),(22,4,0,0,130130,'熔岩火山五',0,-1,6,1742,764,1,125,814,104,'RongYanHuoShan1','RongYanHuoShan2War','YinSen',130140,0,400,1,0,0),(23,4,0,0,130140,'熔岩火山六',1,3,12,288,1980,1,250,612,246,'RongYanHuoShan2','RongYanHuoShan2War','YinSen',130145,1,800,1,0,0),(24,4,0,0,130150,'熔岩火山八',2,2,20,378,511,1,375,192,1752,'RongYanHuoShan4','RongYanHuoShan2War','YinSen',140100,1,1200,1,0,0),(25,5,0,0,140100,'剑灵密室一',0,-1,6,540,1299,1,280,441,134,'JianLingMiShi1','JianLingMiShi1War','ShenMi',140110,1,600,1,0,0),(26,5,0,0,140110,'剑灵密室二',1,3,12,1760,1299,1,560,441,2166,'JianLingMiShi1','JianLingMiShi1War','ShenMi',140115,0,1200,1,1,0),(27,5,0,0,140120,'剑灵密室四',2,2,20,1410,708,1,840,333,714,'JianLingMiShi4','JianLingMiShi4War','ShenMi',140130,0,1800,1,1,0),(28,5,0,0,140130,'剑灵密室五',0,-1,6,540,1299,1,280,441,134,'JianLingMiShi1','JianLingMiShi1War','ShenMi',140140,1,600,1,0,0),(29,5,0,0,140140,'剑灵密室六',1,3,12,1760,1299,1,560,441,2166,'JianLingMiShi1','JianLingMiShi1War','ShenMi',140145,0,1200,1,1,0),(30,5,0,0,140150,'剑灵密室八',2,2,20,510,708,1,840,333,1206,'JianLingMiShi4','JianLingMiShi4War','ShenMi',200100,1,1800,1,0,0),(31,6,0,0,999996,'绝招关卡',0,-1,0,1398,812,0,0,1114,176,'QingZhuLin1','QingZhuLin4','',999997,0,0,1,0,0),(32,1,0,0,100115,'青竹林三',0,-1,6,1769,343,1,10,828,144,'QingZhuLin1','QingZhuLin4','PingHe',100120,0,100,1,0,0),(33,1,0,0,100145,'青竹林七',0,-1,6,1398,812,1,10,1114,176,'QingZhuLin5','QingZhuLin8','PingHe',100150,0,100,1,0,0),(36,2,0,0,110115,'黑夜森林三',0,-1,6,2048,268,1,15,856,114,'HeiYeSenLin1','HeiYeSenLin1War','YinSen',110120,0,150,1,0,0),(37,2,0,0,110145,'暗影秘境三',0,-1,6,2075,245,1,15,286,152,'AnYingMiJing1','AnYingMiJing4','YinSen',110150,0,150,1,0,0),(38,0,1,2,0,'20级铜钱关卡',0,2,6,1500,500,0,100,500,300,'QingZhuLin8','QingZhuLin8','PingHe',0,0,10000,1,0,1),(39,0,1,2,0,'20级经验关卡',0,2,6,1500,500,0,10000,500,300,'QingZhuLin8','QingZhuLin8','PingHe',0,0,100,1,0,2),(40,3,0,0,120115,'莲花峰三',0,-1,6,712,1460,1,40,369,126,'LianHuaFeng1','LianHuaFeng4','YouQu',120120,1,200,1,0,0),(41,3,0,0,120145,'水溶洞三',0,-1,6,693,1322,1,40,306,150,'ShuiRongDong1','ShuiRongDong4','YinSen',120150,0,200,1,0,0),(42,4,0,0,130115,'熔岩火山三',0,-1,6,288,1980,1,125,612,246,'RongYanHuoShan2','RongYanHuoShan2War','YinSen',130120,1,400,1,0,0),(43,4,0,0,130145,'熔岩火山七',0,-1,6,1917,1952,1,125,589,2102,'RongYanHuoShan2','RongYanHuoShan2War','YinSen',130150,0,400,1,1,0),(44,5,0,0,140115,'剑灵密室三',0,-1,6,540,1299,1,280,441,134,'JianLingMiShi1','JianLingMiShi1War','ShenMi',140120,1,600,1,0,0),(45,5,0,0,140145,'剑灵密室七',0,-1,6,540,1299,1,280,441,134,'JianLingMiShi1','JianLingMiShi1War','ShenMi',140150,1,340,1,0,0),(46,6,0,0,999998,'状态关卡',0,-1,0,1398,812,0,0,1114,176,'QingZhuLin1','QingZhuLin4','',999999,0,0,1,0,0),(47,0,1,3,0,'30级铜钱关卡',0,2,6,1500,500,0,100,500,300,'QingZhuLin8','QingZhuLin8','PingHe',0,0,10000,1,0,1),(48,0,1,3,0,'30级经验关卡',0,2,6,1500,500,0,20000,500,300,'QingZhuLin8','QingZhuLin8','PingHe',0,0,100,1,0,2),(49,0,1,4,0,'40级铜钱关卡',0,2,6,1500,500,0,100,500,300,'QingZhuLin8','QingZhuLin8','PingHe',0,0,10000,1,0,1),(50,0,1,4,0,'40级经验关卡',0,2,6,1500,500,0,50000,500,300,'QingZhuLin8','QingZhuLin8','PingHe',0,0,100,1,0,2),(51,0,1,5,0,'50级铜钱关卡',0,2,6,1500,500,0,100,500,300,'QingZhuLin8','QingZhuLin8','PingHe',0,0,10000,1,0,1),(52,0,1,5,0,'50级经验关卡',0,2,6,1500,500,0,80000,500,300,'QingZhuLin8','QingZhuLin8','PingHe',0,0,100,1,0,2),(53,0,1,6,0,'60级铜钱关卡',0,2,6,1500,500,0,100,500,300,'QingZhuLin8','QingZhuLin8','PingHe',0,0,10000,1,0,1),(54,0,1,6,0,'60级经验关卡',0,2,6,1500,500,0,110000,500,300,'QingZhuLin8','QingZhuLin8','PingHe',0,0,100,1,0,2),(55,0,1,7,0,'70级铜钱关卡',0,2,6,1500,500,0,100,500,300,'QingZhuLin8','QingZhuLin8','PingHe',0,0,10000,1,0,1),(56,0,1,7,0,'70级经验关卡',0,2,6,1500,500,0,140000,500,300,'QingZhuLin8','QingZhuLin8','PingHe',0,0,100,1,0,2),(57,0,1,8,0,'80级铜钱关卡',0,2,6,1500,500,0,100,500,300,'QingZhuLin8','QingZhuLin8','PingHe',0,0,10000,1,0,1),(58,0,1,8,0,'80级经验关卡',0,2,6,1500,500,0,160000,500,300,'QingZhuLin8','QingZhuLin8','PingHe',0,0,100,1,0,2),(59,0,1,9,0,'90级铜钱关卡',0,2,6,1500,500,0,100,500,300,'QingZhuLin8','QingZhuLin8','PingHe',0,0,10000,1,0,1),(60,0,1,9,0,'90级经验关卡',0,2,6,1500,500,0,180000,500,300,'QingZhuLin8','QingZhuLin8','PingHe',0,0,100,1,0,2),(61,0,1,10,0,'100级铜钱关卡',0,2,6,1500,500,0,100,500,300,'QingZhuLin8','QingZhuLin8','PingHe',0,0,10000,1,0,1);
/*!40000 ALTER TABLE `mission_level` ENABLE KEYS */;
DROP TABLE IF EXISTS `mission_level_box`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mission_level_box` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `mission_level_id` int(11) NOT NULL COMMENT '关卡id',
  `order` tinyint(4) NOT NULL COMMENT '品质顺序',
  `award_type` tinyint(4) NOT NULL COMMENT '奖励类型(0--铜钱;1--道具;2--装备)',
  `award_chance` tinyint(4) NOT NULL COMMENT '奖励概率',
  `award_num` int(11) NOT NULL COMMENT '奖励数量',
  `item_id` int(11) NOT NULL DEFAULT '0' COMMENT '物品ID(物品奖励填写)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=336 DEFAULT CHARSET=utf8mb4 COMMENT='区域关卡宝箱';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mission_level_box` DISABLE KEYS */;
INSERT INTO `mission_level_box` VALUES (1,1,1,1,20,1,1),(2,1,2,4,10,2,245),(3,1,3,5,10,2,246),(4,1,4,3,30,12,243),(5,1,5,0,30,100,244),(7,2,1,1,10,2,30),(8,2,2,1,20,2,1),(9,2,3,5,20,2,246),(10,2,5,0,25,200,244),(12,3,1,2,20,1,31),(13,3,2,2,20,1,33),(14,3,3,5,20,2,246),(15,3,4,1,20,1,230),(16,3,5,1,20,1,239),(17,4,1,1,20,1,1),(18,4,2,4,10,2,245),(19,4,3,5,10,2,246),(20,4,4,3,30,10,243),(21,4,5,0,30,100,244),(24,2,4,3,25,24,243),(25,5,1,1,10,1,30),(26,5,2,1,20,2,1),(28,5,3,5,20,2,246),(29,5,4,3,25,20,243),(30,5,5,0,25,200,244),(33,6,1,2,20,1,36),(34,6,3,2,20,1,38),(35,6,2,2,20,1,37),(36,6,4,1,20,1,239),(37,6,5,1,20,1,230),(38,7,1,1,20,2,1),(41,7,2,4,10,2,245),(42,7,3,5,10,2,246),(43,7,4,3,30,25,243),(44,8,1,1,10,2,30),(45,8,2,1,20,2,1),(47,8,3,5,20,2,246),(48,8,4,3,25,55,243),(49,8,5,0,25,400,244),(50,9,1,2,20,1,36),(51,9,2,2,20,1,37),(55,9,3,2,20,1,38),(56,7,5,0,30,200,244),(57,9,4,1,20,1,230),(58,9,5,1,20,1,239),(59,10,1,1,20,2,1),(60,10,2,4,10,2,245),(61,10,3,5,10,2,246),(62,10,4,3,30,33,243),(63,10,5,0,30,200,244),(64,11,1,1,10,4,30),(65,11,2,1,20,4,1),(66,11,3,5,20,2,246),(67,11,4,3,25,61,243),(68,11,5,0,25,400,244),(69,12,1,2,20,1,31),(70,12,2,2,20,1,32),(71,12,3,2,20,1,33),(72,12,4,1,20,1,230),(73,12,5,0,20,1,239),(74,13,1,1,20,2,1),(75,13,2,4,10,2,245),(76,13,3,5,10,2,246),(77,13,4,3,30,70,243),(79,13,5,0,30,300,244),(80,14,1,1,10,4,30),(81,14,2,1,10,4,1),(82,14,3,1,10,2,52),(83,14,4,3,40,140,243),(85,14,5,0,30,600,244),(86,15,1,2,20,1,36),(87,15,2,2,20,1,37),(88,15,3,2,20,1,38),(90,15,4,1,20,1,230),(91,15,5,1,20,1,239),(92,16,1,1,20,2,1),(93,16,2,4,10,2,245),(94,16,3,5,10,2,246),(95,16,4,3,30,84,243),(97,16,5,0,30,300,244),(98,17,1,1,10,4,30),(99,17,2,1,20,4,1),(100,17,3,1,20,2,52),(101,17,4,3,25,166,243),(103,17,5,0,25,600,244),(104,18,1,2,20,1,31),(105,18,2,2,20,1,32),(106,18,3,2,20,1,33),(108,18,4,1,20,1,230),(109,18,5,1,20,1,239),(110,19,1,1,20,3,1),(111,19,2,4,10,2,245),(112,19,3,5,10,2,246),(113,19,4,3,30,230,243),(115,19,5,0,50,400,244),(117,20,1,1,10,6,30),(118,20,2,1,20,6,1),(119,20,3,1,20,3,52),(120,20,4,3,25,460,243),(122,20,5,0,25,800,244),(123,21,1,2,20,1,36),(124,21,2,2,20,1,37),(125,21,3,2,20,1,38),(127,21,4,1,20,1,230),(128,21,5,1,20,1,239),(129,22,1,1,20,3,1),(130,22,2,4,10,2,245),(131,22,3,5,10,2,246),(132,22,4,3,30,258,243),(134,22,5,0,30,400,244),(136,23,1,1,10,6,30),(137,23,2,1,20,6,1),(138,23,3,1,20,3,52),(139,23,4,3,25,528,243),(141,23,5,0,25,800,244),(142,24,1,2,20,1,31),(143,24,2,2,20,1,32),(144,24,3,2,20,1,33),(146,24,4,1,20,1,230),(147,24,5,1,20,1,239),(148,25,1,1,20,3,1),(149,25,2,4,10,2,245),(150,25,3,5,10,2,246),(151,25,4,3,30,540,243),(153,25,5,0,30,500,244),(155,26,1,1,10,6,30),(156,26,2,1,20,6,1),(157,26,3,1,20,4,52),(158,26,4,3,25,540,243),(160,26,5,0,25,500,244),(161,27,1,2,20,1,66),(162,27,2,2,20,1,67),(163,27,3,2,20,1,68),(164,27,4,1,20,1,230),(166,27,5,1,20,1,239),(167,28,1,1,20,3,1),(168,28,2,4,10,2,245),(169,28,3,5,10,2,246),(170,28,4,3,30,572,243),(172,28,5,0,30,500,244),(174,29,1,1,10,6,30),(175,29,2,1,20,6,1),(176,29,3,1,20,4,52),(177,29,4,3,25,1122,243),(179,29,5,0,25,1000,244),(180,30,1,2,20,1,61),(181,30,2,2,20,1,62),(182,30,3,2,20,1,63),(184,30,4,2,20,1,64),(185,30,5,1,20,1,230),(186,31,1,0,20,1,244),(187,31,2,0,20,1,244),(188,31,3,0,20,1,244),(189,31,4,0,20,1,244),(190,31,5,0,20,1,244),(191,32,1,1,20,1,1),(192,32,2,4,10,2,245),(193,32,3,5,10,2,246),(194,32,4,3,30,10,243),(195,32,5,0,30,100,244),(196,33,1,1,20,1,1),(197,33,2,4,10,2,245),(198,33,3,5,10,2,246),(199,33,4,3,30,12,243),(200,33,5,0,30,100,244),(201,36,1,1,20,2,1),(202,36,2,4,10,2,245),(203,36,3,5,10,2,246),(204,36,4,3,30,25,243),(205,36,5,0,30,200,244),(206,37,1,1,20,2,1),(207,37,2,4,10,2,245),(208,37,3,5,10,2,246),(209,37,4,3,30,33,243),(210,37,5,0,30,200,244),(211,38,1,0,50,3000,244),(212,38,2,0,20,6000,244),(213,38,3,0,15,9000,244),(214,38,4,0,10,12000,244),(215,38,5,5,5,2,246),(216,39,1,3,50,1000,243),(217,39,2,3,20,1300,243),(218,39,3,3,15,1500,243),(219,39,4,3,10,1600,243),(220,39,5,4,5,2,245),(221,40,1,1,20,2,1),(222,40,2,4,10,2,245),(223,40,3,5,10,2,246),(224,40,4,3,30,70,243),(225,40,5,0,30,300,244),(226,41,1,1,20,2,1),(227,41,2,4,10,2,245),(228,41,3,5,10,2,246),(229,41,4,3,30,84,243),(230,41,5,0,30,300,244),(231,42,1,1,20,3,1),(232,42,2,4,10,2,245),(233,42,3,5,10,2,246),(234,42,4,3,30,230,243),(235,42,5,0,30,400,244),(236,43,1,1,20,3,1),(237,43,2,4,10,2,245),(238,43,3,5,10,2,246),(239,43,4,3,30,258,243),(240,43,5,0,30,400,244),(241,44,1,1,20,3,1),(242,44,2,4,10,2,245),(243,44,3,5,10,2,246),(244,44,4,3,30,540,243),(245,44,5,0,30,500,244),(246,45,1,1,20,3,1),(247,45,2,4,10,2,245),(248,45,3,5,10,2,246),(249,45,4,3,30,572,243),(250,45,5,0,30,500,244),(251,46,1,0,20,1,0),(252,46,2,0,20,1,0),(253,46,3,0,20,1,0),(254,46,4,0,20,1,0),(255,46,5,0,20,1,0),(256,51,1,0,50,6000,244),(257,51,2,0,20,9000,244),(258,51,3,0,15,12000,244),(259,51,4,0,10,15000,244),(260,51,5,5,5,2,246),(261,52,1,3,50,25400,243),(262,52,2,3,20,32700,243),(263,52,3,3,15,38100,243),(264,52,4,3,10,40000,243),(265,52,5,4,5,2,245),(266,47,1,0,50,4000,244),(267,47,2,0,20,7000,244),(268,47,3,0,15,10000,244),(269,47,4,0,10,13000,244),(270,47,5,5,5,2,246),(271,49,1,0,50,5000,244),(272,49,2,0,20,8000,244),(273,49,3,0,15,11000,244),(274,49,4,0,10,14000,244),(275,49,5,5,5,2,246),(276,53,1,0,50,7000,244),(277,53,2,0,20,10000,244),(278,53,3,0,15,13000,244),(279,53,4,0,10,16000,244),(280,53,5,5,5,2,246),(281,48,1,3,50,4600,243),(282,48,2,3,20,6000,243),(283,48,3,3,15,7000,243),(284,48,4,3,10,7500,243),(285,48,5,4,5,2,245),(286,55,1,0,50,8000,244),(287,55,2,0,20,11000,244),(288,55,3,0,15,14000,0),(289,55,4,0,10,17000,244),(290,55,5,5,5,2,246),(291,50,1,3,50,10800,243),(292,50,2,3,20,14000,243),(293,50,3,3,15,16200,243),(294,50,4,3,10,17000,243),(295,50,5,3,5,2,245),(296,57,1,0,50,9000,244),(297,57,2,0,20,12000,244),(298,57,3,0,15,15000,244),(299,57,4,0,10,18000,244),(300,57,5,5,5,2,246),(301,59,1,0,50,10000,244),(302,59,2,0,20,13000,244),(303,59,3,0,10,16000,244),(304,59,4,0,15,19000,244),(305,59,5,5,5,2,246),(306,54,1,3,50,38100,243),(307,54,2,3,20,49000,243),(308,54,3,3,15,57200,243),(309,54,4,3,10,60000,243),(310,54,5,4,5,2,245),(311,56,1,3,50,45200,243),(312,56,2,3,20,58100,243),(313,56,3,3,15,68000,243),(314,56,4,3,10,71100,243),(315,56,5,4,5,2,245),(316,58,1,3,50,55300,243),(317,58,2,3,20,71100,243),(318,58,3,3,15,83000,243),(319,58,4,3,10,87000,243),(320,58,5,4,5,2,245),(321,60,1,3,50,58300,243),(322,60,2,3,20,75000,243),(323,60,3,3,15,87500,243),(324,60,4,3,10,91600,243),(325,60,5,3,5,2,245),(331,61,1,0,50,11000,244),(332,61,2,0,20,14000,244),(333,61,3,0,15,17000,244),(334,61,4,0,10,20000,244),(335,61,5,0,5,2,246);
/*!40000 ALTER TABLE `mission_level_box` ENABLE KEYS */;
DROP TABLE IF EXISTS `mission_level_small_box`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mission_level_small_box` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `mission_level_id` int(11) NOT NULL COMMENT '关卡id',
  `box_x` int(11) NOT NULL COMMENT '宝箱x坐标',
  `box_y` int(11) NOT NULL COMMENT '宝箱y坐标',
  `probability` tinyint(4) NOT NULL COMMENT '出现概率',
  `box_dir` tinyint(4) NOT NULL COMMENT '宝箱朝向(0--左;1--右)',
  `items_kind` tinyint(4) NOT NULL COMMENT '出现物品有几种',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COMMENT='关卡小宝箱';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mission_level_small_box` DISABLE KEYS */;
INSERT INTO `mission_level_small_box` VALUES (7,7,1498,729,30,0,1),(8,8,1198,411,30,0,1),(9,8,1577,1165,30,0,1),(10,36,786,1102,30,0,1),(11,13,1526,410,30,0,1),(12,14,1526,410,30,0,1),(13,40,1526,410,30,0,1),(14,10,911,1054,30,0,1),(15,37,1367,922,30,0,1),(16,11,519,661,30,1,1),(17,11,1495,970,30,0,1),(18,16,1121,469,30,0,1),(19,41,1557,908,30,0,1),(20,17,1402,1000,30,0,1),(21,17,722,964,30,0,1),(22,42,1684,1645,30,0,1),(23,23,2052,1093,30,0,1),(24,20,1000,1076,30,0,1),(25,43,966,1744,30,1,1),(26,25,510,701,30,0,1),(27,44,1124,378,30,0,1),(28,28,349,700,30,1,1),(29,45,1426,344,30,1,1),(30,26,1113,1024,30,0,1),(31,26,1467,861,30,0,1),(32,29,1113,1024,30,0,1),(33,29,1467,861,30,0,1);
/*!40000 ALTER TABLE `mission_level_small_box` ENABLE KEYS */;
DROP TABLE IF EXISTS `mission_level_small_box_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mission_level_small_box_items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `small_box_id` int(11) NOT NULL COMMENT '小宝箱id',
  `item_id` int(11) NOT NULL COMMENT '物品ID',
  `probability` tinyint(4) NOT NULL COMMENT '出现概率',
  `item_number` int(11) NOT NULL DEFAULT '0' COMMENT '奖励数量',
  `award_type` tinyint(4) NOT NULL COMMENT '奖励类型(0--铜钱;1--道具;2--装备)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=116 DEFAULT CHARSET=utf8mb4 COMMENT='关卡小宝箱';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mission_level_small_box_items` DISABLE KEYS */;
INSERT INTO `mission_level_small_box_items` VALUES (2,1,31,20,0,0),(5,4,244,50,100,0),(6,4,244,25,200,0),(7,4,244,25,300,0),(8,7,244,30,100,0),(9,7,244,30,150,0),(10,7,244,30,200,0),(11,7,244,10,400,0),(12,8,244,30,100,0),(13,8,244,30,150,0),(14,8,244,30,200,0),(15,8,244,10,400,0),(16,9,244,30,100,0),(17,9,244,30,150,0),(18,9,244,30,200,0),(19,9,244,10,400,0),(20,10,244,30,100,0),(21,10,244,30,150,0),(22,10,244,30,200,0),(23,10,244,10,400,0),(24,11,244,30,100,0),(25,11,244,30,150,0),(26,11,244,30,200,0),(27,11,244,10,400,0),(28,12,244,30,100,0),(29,12,244,30,150,0),(30,12,244,30,200,0),(31,12,244,10,400,0),(32,14,244,30,100,0),(33,14,244,30,150,0),(34,14,244,30,200,0),(35,14,244,10,400,0),(36,15,244,30,100,0),(37,15,244,30,150,0),(38,15,244,30,200,0),(39,15,244,10,400,0),(40,16,244,30,100,0),(41,16,244,30,150,0),(42,16,244,30,200,0),(43,16,244,10,400,0),(44,17,244,30,100,0),(45,17,244,30,150,0),(46,17,244,30,200,0),(47,17,244,10,400,0),(48,13,244,30,100,0),(49,13,244,30,150,0),(50,13,244,30,200,0),(51,13,244,10,400,0),(52,18,243,30,100,0),(53,18,243,30,150,0),(54,18,243,30,200,0),(55,18,243,10,400,0),(56,19,244,30,100,0),(57,19,244,30,150,0),(58,19,244,30,200,0),(59,19,244,10,400,0),(60,20,244,30,100,0),(61,20,244,30,150,0),(62,20,244,30,200,0),(63,20,244,10,400,0),(64,21,244,30,100,0),(65,21,244,30,150,0),(66,21,244,30,200,0),(67,21,244,10,400,0),(68,22,244,30,100,0),(69,22,244,30,150,0),(70,22,244,30,200,0),(71,22,244,10,400,0),(72,23,244,30,100,0),(73,23,244,30,150,0),(74,23,244,30,200,0),(75,23,244,10,400,0),(76,24,244,30,100,0),(77,24,244,30,150,0),(78,24,244,30,200,0),(79,24,244,10,400,0),(80,25,244,30,100,0),(81,25,244,30,150,0),(82,25,244,30,200,0),(83,25,244,10,400,0),(84,26,244,30,100,0),(85,26,244,30,150,0),(86,26,244,30,200,0),(87,26,244,10,400,0),(88,27,244,30,100,0),(89,27,244,30,150,0),(90,27,244,30,200,0),(91,27,244,10,400,0),(92,28,244,30,100,0),(93,28,244,30,150,0),(94,28,244,30,200,0),(95,28,244,10,400,0),(96,29,244,30,100,0),(97,29,244,30,150,0),(98,29,244,30,200,0),(99,29,244,10,400,0),(100,33,244,30,100,0),(101,33,244,30,150,0),(102,33,244,30,200,0),(103,33,244,10,400,0),(104,32,244,30,100,0),(105,32,244,30,150,0),(106,32,244,30,200,0),(107,32,244,10,400,0),(108,31,244,30,100,0),(109,31,244,30,150,0),(110,31,244,30,200,0),(111,31,244,10,400,0),(112,30,244,30,100,0),(113,30,244,30,150,0),(114,30,244,30,200,0),(115,30,244,10,400,0);
/*!40000 ALTER TABLE `mission_level_small_box_items` ENABLE KEYS */;
DROP TABLE IF EXISTS `mission_talk`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mission_talk` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '对话ID',
  `enemy_id` int(11) NOT NULL COMMENT '副本敌人ID',
  `content` varchar(1024) DEFAULT '' COMMENT '对话内容',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='副本战场对话';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mission_talk` DISABLE KEYS */;
/*!40000 ALTER TABLE `mission_talk` ENABLE KEYS */;
DROP TABLE IF EXISTS `multi_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `multi_level` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `sign_war` varchar(50) NOT NULL COMMENT '战斗资源标识',
  `music` varchar(20) NOT NULL DEFAULT '' COMMENT '音乐资源标识',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '关卡名称',
  `require_level` smallint(6) NOT NULL COMMENT '主角等级要求',
  `award_exp` bigint(20) NOT NULL DEFAULT '0' COMMENT '奖励经验',
  `award_coin` bigint(20) NOT NULL DEFAULT '0' COMMENT '奖励铜钱',
  `award_item1_id` int(11) NOT NULL DEFAULT '0' COMMENT '奖励物品1 id',
  `award_item1_num` int(11) NOT NULL DEFAULT '0' COMMENT '物品1数量',
  `award_item2_id` int(11) NOT NULL DEFAULT '0' COMMENT '奖励物品2 id',
  `award_item2_num` int(11) NOT NULL DEFAULT '0' COMMENT '物品2数量',
  `award_item3_id` int(11) NOT NULL DEFAULT '0' COMMENT '奖励物品3 id',
  `award_item3_num` int(11) NOT NULL DEFAULT '0' COMMENT '物品3数量',
  `lock` int(11) NOT NULL DEFAULT '0' COMMENT '关卡开启权值',
  `award_lock` int(11) NOT NULL DEFAULT '0' COMMENT '奖励权值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='多人关卡';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `multi_level` DISABLE KEYS */;
INSERT INTO `multi_level` VALUES (1,'QingZhuLin8','','刀疤兔 Lv.30',30,5000,5000,61,1,62,1,63,1,0,100110),(2,'HeiYeSenLin4','','天狼妖 Lv.35',35,5000,5000,61,1,64,1,65,1,100110,100120),(3,'AnYingMiJing4','','妖龙 Lv.40',40,5000,5000,66,1,67,1,68,1,100120,100130),(4,'LianHuaFeng4','','金蟾王 Lv.45',45,5000,5000,215,1,216,1,217,1,100130,100140),(5,'ShuiRongDong4','','剧毒臭泥 Lv.50',50,5000,5000,213,1,211,1,214,1,100140,100150),(6,'RongYanHuoShan4','','燃魁首领 Lv.55',55,5000,5000,216,1,214,1,215,1,100150,100160),(7,'RongYanHuoShan4','','炎龙 Lv.60',60,5000,5000,211,1,215,1,213,1,100160,100170),(8,'JianLingMiShi4','','古代剑灵 Lv.65',65,5000,5000,217,1,211,1,216,1,110170,110180);
/*!40000 ALTER TABLE `multi_level` ENABLE KEYS */;
DROP TABLE IF EXISTS `quest`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `quest` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `order` int(11) NOT NULL COMMENT '任务排序，低的排前',
  `name` varchar(10) NOT NULL COMMENT '任务标题',
  `type` tinyint(4) DEFAULT '0' COMMENT '任务类型',
  `desc` varchar(240) DEFAULT '' COMMENT '简介',
  `require_level` int(11) NOT NULL COMMENT '要求玩家等级',
  `town_id` smallint(6) NOT NULL DEFAULT '-1' COMMENT '城镇ID',
  `town_npc_id` int(11) NOT NULL COMMENT '完成任务所需的城镇NPC ID',
  `mission_level_id` int(11) DEFAULT '0' COMMENT '完成任务所需的关卡ID',
  `enemy_num` int(10) unsigned NOT NULL COMMENT '敌人组数',
  `enemy_id` smallint(6) DEFAULT '0' COMMENT '敌人ID',
  `drama_mode` tinyint(4) DEFAULT '0' COMMENT '剧情模式(1--任务完成播放剧情)',
  `award_exp` int(11) NOT NULL COMMENT '奖励经验',
  `award_coins` bigint(20) NOT NULL DEFAULT '0' COMMENT '奖励铜钱',
  `award_item1_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品1',
  `award_item1_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品1数量',
  `award_item2_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品2',
  `award_item2_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品2数量',
  `award_item3_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品3',
  `award_item3_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品3数量',
  `award_item4_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品4',
  `award_item4_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品4数量',
  `award_func_key` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励功能权值',
  `award_role_id` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '奖励角色ID',
  `award_role_level` tinyint(4) NOT NULL DEFAULT '0' COMMENT '奖励角色等级',
  `award_mission_level_lock` int(11) NOT NULL DEFAULT '0' COMMENT '奖励关卡权值',
  `award_town_key` int(11) NOT NULL DEFAULT '0' COMMENT '奖励城镇权值',
  `award_physical` tinyint(4) NOT NULL DEFAULT '0' COMMENT '奖励体力',
  `auto_mission_level_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '自动进入关卡id',
  `auto_fight` tinyint(4) DEFAULT '0' COMMENT '自动打怪',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8mb4 COMMENT='主线任务';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `quest` DISABLE KEYS */;
INSERT INTO `quest` VALUES (1,10,'侠客降临',0,'大侠降临侠客岛。',1,1,0,0,0,0,1,0,10,0,0,0,0,0,0,0,0,0,0,0,100100,0,0,4,0),(2,20,'探索青竹林一',3,'探索青竹林一。',1,0,0,4,0,0,0,0,10,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(3,30,'探索青竹林二',3,'探索青竹林二。',0,0,0,5,0,0,0,0,10,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(4,40,'兔子也称王',4,'为什么一只兔子成了竹林里的老大。',0,0,0,6,18,2,1,10,100,1,5,0,0,0,0,0,0,0,0,0,0,0,40,0,0),(6,50,'阴影碎片',1,'把阴影碎片带回神龙岛给龙姬看看。',0,1,1,0,0,0,1,0,20,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(7,60,'探索青竹林五',3,'探索青竹林五。',0,0,0,1,0,0,0,0,20,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(8,70,'探索青竹林六',3,'探索青竹林六。',0,0,0,2,0,0,0,0,20,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(9,80,'探索青竹林八',2,'探索青竹林八。',0,0,0,3,0,0,1,0,20,0,0,0,0,0,0,0,0,0,4,5,0,0,0,0,0),(10,90,'搭救伙伴',4,'那只兔子看上去很凶，不管了，救人要紧。',0,0,0,3,93,26,1,20,200,1,5,0,0,0,0,0,0,0,0,0,0,0,40,0,1),(12,100,'返回神龙岛',1,'先回神龙岛再说了。',0,1,1,0,0,0,1,0,30,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(13,110,'探索黑夜森林一',3,'',0,0,0,7,0,0,0,0,30,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(14,120,'探索黑夜森林二',3,'',0,0,0,8,0,0,0,0,30,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(15,130,'探索黑夜森林四',2,'',0,0,0,9,0,0,1,0,30,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(16,140,'伙伴踪迹',4,'',0,0,0,9,27,8,1,20,300,1,10,0,0,0,0,0,0,0,3,7,0,0,40,0,0),(17,150,'魔化元凶',1,'',0,1,1,0,0,0,1,0,40,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(18,160,'探索暗影秘境一',3,'',0,0,0,10,0,0,0,0,40,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(19,170,'探索暗影秘境二',3,'',0,0,0,11,0,0,0,0,40,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(20,180,'探索暗影秘境四',2,'',0,0,0,12,0,0,1,0,40,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(21,190,'影魂龙珠',4,'',0,0,0,12,36,25,1,30,400,1,10,0,0,0,0,0,0,1000,0,0,0,0,40,0,0),(22,200,'莲花峰异变',1,'',0,1,1,0,0,0,1,0,50,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(23,210,'探索莲花峰一',3,'',0,0,0,13,0,0,0,0,50,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(24,220,'探索莲花峰二',3,'',0,0,0,14,0,0,0,0,50,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(25,230,'探索莲花峰四',2,'',0,0,0,15,0,0,1,0,50,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(26,240,'净化金蟾王',4,'',0,0,0,15,45,11,1,50,500,1,10,0,0,0,0,0,0,0,0,0,0,0,40,0,0),(27,250,'探索水溶洞一',3,'',0,0,0,16,0,0,0,0,60,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(28,260,'探索水溶洞二',3,'',0,0,0,17,0,0,0,0,60,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(29,270,'探索水溶洞四',2,'',0,0,0,18,0,0,1,0,60,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(30,280,'清除剧毒臭泥',4,'',0,0,0,18,54,14,1,80,1000,1,10,0,0,0,0,0,0,5000,0,0,0,0,40,0,0),(31,290,'伙伴消息',1,'',0,1,1,0,0,0,1,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(33,300,'探索熔岩火山一',3,'',0,0,0,19,0,0,1,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(34,310,'探索熔岩火山二',3,'',0,0,0,20,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(35,320,'探索熔岩火山四',3,'',0,0,0,21,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(36,330,'探索熔岩火山五',3,'',0,0,0,22,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(37,340,'探索熔岩火山六',3,'',0,0,0,23,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(38,350,'探索熔岩火山八',2,'',0,0,0,24,0,0,1,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(39,360,'炎龙的误会',4,'',0,0,0,24,72,19,1,440,5000,1,10,0,0,0,0,0,0,0,0,0,0,0,40,0,0),(40,370,'探索剑灵密室一',3,'',0,0,0,25,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(41,380,'探索剑灵密室二',3,'',0,0,0,26,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(42,390,'探索剑灵密室四',3,'',0,0,0,27,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(43,400,'探索剑灵密室五',3,'',0,0,0,28,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(44,410,'探索剑灵密室六',3,'',0,0,0,29,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(45,420,'探索剑灵密室八',2,'',0,0,0,30,0,0,1,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(46,430,'古代剑灵',4,'',0,0,0,30,90,23,1,1000,10000,1,10,0,0,0,0,0,0,0,5,20,0,100110,0,0,0),(47,440,'前往天地盟',0,'',0,2,0,0,0,0,1,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(49,35,'探索青竹林四',2,'探索青竹林四。',0,0,0,6,0,0,0,0,10,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(50,33,'探索青竹林三',3,'探索青竹林三。',0,0,0,32,0,0,0,0,10,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(51,75,'探索青竹林七',3,'探索青竹林七。',0,0,0,33,0,0,0,0,20,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(53,125,'探索黑夜森林三',3,'',0,0,0,36,0,0,0,0,30,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(54,175,'探索暗影秘境三',3,'',0,0,0,37,0,0,0,0,40,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(55,225,'探索莲花峰三',3,'',0,0,0,40,0,0,0,0,50,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(56,265,'探索水溶洞三',3,'',0,0,0,41,0,0,0,0,60,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(57,315,'探索熔岩火山三',3,'',0,0,0,42,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(58,345,'探索熔岩火山七',3,'',0,0,0,43,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(59,385,'探索剑灵密室三',3,'',0,0,0,44,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(60,415,'探索剑灵密室七',3,'',0,0,0,45,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0);
/*!40000 ALTER TABLE `quest` ENABLE KEYS */;
DROP TABLE IF EXISTS `resource_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `resource_level` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `max_level` smallint(6) NOT NULL COMMENT '允许开放的等级上限',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COMMENT='资源关卡';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `resource_level` DISABLE KEYS */;
INSERT INTO `resource_level` VALUES (2,29),(3,39),(4,49),(5,59),(6,69),(7,79),(8,89),(9,99),(10,109);
/*!40000 ALTER TABLE `resource_level` ENABLE KEYS */;
DROP TABLE IF EXISTS `role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role` (
  `id` tinyint(4) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(10) NOT NULL COMMENT '角色名称',
  `sign` varchar(20) DEFAULT NULL COMMENT '资源标识',
  `type` tinyint(4) NOT NULL COMMENT '类型：1.主角，2.伙伴',
  `is_special` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否特殊伙伴 0不是 1是',
  `skill_id1` smallint(6) NOT NULL DEFAULT '0' COMMENT '默认绝招1',
  `skill_id2` smallint(6) NOT NULL DEFAULT '0' COMMENT '默认绝招2',
  `buddy_level` smallint(6) NOT NULL DEFAULT '0' COMMENT '伙伴等级',
  `mission_lock` int(11) DEFAULT '0' COMMENT '解锁副本权值',
  `scale` tinyint(4) NOT NULL DEFAULT '100' COMMENT '缩放比',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COMMENT='角色表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `role` DISABLE KEYS */;
INSERT INTO `role` VALUES (1,'义峰','YiFeng',1,0,0,0,0,0,100),(2,'昕苒','XinRan',1,0,0,0,0,0,100),(3,'南宫白夜','NanGongBaiYe',2,0,5,6,0,0,100),(4,'朱媛媛','ZhuYuanYuan',2,0,7,8,0,0,100),(5,'车晓芸','CheXiaoYun',2,0,49,50,0,0,100),(6,'燕无名','YanWuMing',2,0,51,52,0,0,100);
/*!40000 ALTER TABLE `role` ENABLE KEYS */;
DROP TABLE IF EXISTS `role_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role_level` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '角色等级ID',
  `role_id` tinyint(4) NOT NULL COMMENT '角色ID',
  `level` int(11) NOT NULL COMMENT '等级 - level',
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
) ENGINE=InnoDB AUTO_INCREMENT=527 DEFAULT CHARSET=utf8mb4 COMMENT='角色等级';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `role_level` DISABLE KEYS */;
INSERT INTO `role_level` VALUES (11,1,1,200,40,15,10,80,15,6,0,6,0,0,6,2,200,100,200,2,2),(12,1,2,300,70,25,20,110,15,6,0,6,0,0,6,2,200,100,200,2,2),(13,1,3,400,100,35,30,140,15,6,0,6,0,0,6,2,200,100,200,2,2),(14,1,4,500,130,45,40,170,15,6,0,6,0,0,6,2,200,100,200,2,2),(15,1,5,600,160,55,50,200,15,6,0,6,0,0,6,2,200,100,200,2,2),(16,1,6,700,190,65,60,230,15,6,0,6,0,0,6,2,200,100,200,2,2),(17,1,7,800,220,75,70,260,15,6,0,6,0,0,6,2,200,100,200,2,2),(18,1,8,900,250,85,80,290,15,6,0,6,0,0,6,2,100,100,200,2,2),(19,1,9,1000,280,95,90,320,15,6,0,6,0,0,6,2,100,100,200,2,2),(20,1,10,1100,310,105,100,350,15,6,0,6,0,0,6,2,100,100,200,2,2),(21,1,11,1210,343,116,111,383,15,6,0,6,0,0,6,2,100,100,200,2,2),(22,1,12,1320,376,127,122,416,15,6,0,6,0,0,6,2,100,100,200,2,2),(23,1,13,1430,409,138,133,449,15,6,0,6,0,0,6,2,100,100,200,2,2),(24,1,14,1540,442,149,144,482,15,6,0,6,0,0,6,2,100,100,200,2,2),(25,1,15,1650,475,160,155,515,15,6,0,6,0,0,6,2,100,100,200,2,2),(26,1,16,1760,508,171,166,548,15,6,0,6,0,0,6,2,100,100,200,2,2),(27,1,17,1870,541,182,177,581,15,6,0,6,0,0,6,2,100,100,200,2,2),(28,1,18,1980,574,193,188,614,15,6,0,6,0,0,6,2,100,100,200,2,2),(29,1,19,2090,607,204,199,647,15,6,0,6,0,0,6,2,100,100,200,2,2),(30,1,20,2200,640,215,210,680,15,6,0,6,0,0,6,2,100,100,200,2,2),(31,1,21,2320,676,227,222,716,15,6,0,6,0,0,6,2,125,100,200,2,2),(32,1,22,2440,712,239,234,752,15,6,0,6,0,0,6,2,125,100,200,2,2),(33,1,23,2560,748,251,246,788,15,6,0,6,0,0,6,2,125,100,200,2,2),(34,1,24,2680,784,263,258,824,15,6,0,6,0,0,6,2,125,100,200,2,2),(35,1,25,2800,820,275,270,860,15,6,0,6,0,0,6,2,125,100,200,2,2),(36,1,26,2920,856,287,282,896,15,6,0,6,0,0,6,2,125,100,200,2,2),(37,1,27,3040,892,299,294,932,15,6,0,6,0,0,6,2,125,100,200,2,2),(38,1,28,3160,928,311,306,968,15,6,0,6,0,0,6,2,125,100,200,2,2),(39,1,29,3280,964,323,318,1004,15,6,0,6,0,0,6,2,125,100,200,2,2),(40,1,30,3400,1000,335,330,1040,15,6,0,6,0,0,6,2,125,100,200,2,2),(41,1,31,3530,1039,348,343,1079,15,6,0,6,0,0,6,2,150,100,200,2,2),(42,1,32,3660,1078,361,356,1118,15,6,0,6,0,0,6,2,150,100,200,2,2),(43,1,33,3790,1117,374,369,1157,15,6,0,6,0,0,6,2,150,100,200,2,2),(44,1,34,3920,1156,387,382,1196,15,6,0,6,0,0,6,2,150,100,200,2,2),(45,1,35,4050,1195,400,395,1235,15,6,0,6,0,0,6,2,150,100,200,2,2),(46,1,36,4180,1234,413,408,1274,15,6,0,6,0,0,6,2,150,100,200,2,2),(47,1,37,4310,1273,426,421,1313,15,6,0,6,0,0,6,2,150,100,200,2,2),(48,1,38,4440,1312,439,434,1352,15,6,0,6,0,0,6,2,150,100,200,2,2),(49,1,39,4570,1351,452,447,1391,15,6,0,6,0,0,6,2,150,100,200,2,2),(50,1,40,4700,1390,465,460,1430,15,6,0,6,0,0,6,2,150,100,200,2,2),(51,1,41,2080,604,203,198,644,15,6,0,6,0,0,6,2,175,100,200,2,2),(52,1,42,2160,628,211,206,668,15,6,0,6,0,0,6,2,175,100,200,2,2),(53,1,43,2240,652,219,214,692,15,6,0,6,0,0,6,2,175,100,200,2,2),(54,1,44,2320,676,227,222,716,15,6,0,6,0,0,6,2,175,100,200,2,2),(64,1,45,2400,700,235,230,740,15,6,0,6,0,0,6,2,175,100,200,2,2),(65,1,46,2480,724,243,238,764,15,6,0,6,0,0,6,2,175,100,200,2,2),(66,1,47,2560,748,251,246,788,15,6,0,6,0,0,6,2,175,100,200,2,2),(67,1,48,2640,772,259,254,812,15,6,0,6,0,0,6,2,175,100,200,2,2),(68,1,49,2720,796,267,262,836,15,6,0,6,0,0,6,2,175,100,200,2,2),(69,1,50,2800,820,275,270,860,15,6,0,6,0,0,6,2,175,100,200,2,2),(70,1,51,2890,847,284,279,887,15,6,0,6,0,0,6,2,200,100,200,2,2),(71,3,1,800,130,120,20,20,6,6,0,0,0,0,0,0,100,100,200,2,2),(72,3,2,1100,160,140,30,30,6,6,0,0,0,0,0,0,100,100,200,2,2),(73,3,3,1400,190,160,40,40,6,6,0,0,0,0,0,0,100,100,200,2,2),(74,4,1,200,40,20,130,30,6,6,0,6,0,0,0,0,100,100,200,2,2),(75,4,2,300,70,30,160,50,6,6,0,6,0,0,0,0,100,100,200,2,2),(76,4,3,400,100,40,190,70,6,6,0,6,0,0,0,0,100,100,200,2,2),(77,1,52,2980,874,293,288,914,15,6,0,6,0,0,6,2,200,100,200,2,2),(78,1,53,3070,901,302,297,941,15,6,0,6,0,0,6,2,200,100,200,2,2),(79,1,54,3160,928,311,306,968,15,6,0,6,0,0,6,2,200,100,200,2,2),(80,1,55,3250,955,320,315,995,15,6,0,6,0,0,6,2,200,100,200,2,2),(81,1,56,3340,982,329,324,1022,15,6,0,6,0,0,6,2,200,100,200,2,2),(82,1,57,3430,1009,338,333,1049,15,6,0,6,0,0,6,2,200,100,200,2,2),(83,1,58,3520,1036,347,342,1076,15,6,0,6,0,0,6,2,200,100,200,2,2),(84,1,59,3610,1063,356,351,1103,15,6,0,6,0,0,6,2,200,100,200,2,2),(85,1,60,3700,1090,365,360,1130,15,6,0,6,0,0,6,2,200,100,200,2,2),(86,1,61,3800,1120,375,370,1160,15,6,0,6,0,0,6,2,200,100,200,2,2),(87,1,62,3900,1150,385,380,1190,15,6,0,6,0,0,6,2,200,100,200,2,2),(88,1,63,4000,1180,395,390,1220,15,6,0,6,0,0,6,2,200,100,200,2,2),(89,1,64,4100,1210,405,400,1250,15,6,0,6,0,0,6,2,200,100,200,2,2),(90,1,65,4200,1240,415,410,1280,15,6,0,6,0,0,6,2,200,100,200,2,2),(91,1,66,4300,1270,425,420,1310,15,6,0,6,0,0,6,2,200,100,200,2,2),(92,1,67,4400,1300,435,430,1340,15,6,0,6,0,0,6,2,200,100,200,2,2),(93,1,68,4500,1330,445,440,1370,15,6,0,6,0,0,6,2,200,100,200,2,2),(94,1,69,4600,1360,455,450,1400,15,6,0,6,0,0,6,2,200,100,200,2,2),(95,1,70,4700,1390,465,450,1430,15,6,0,6,0,0,6,2,200,100,200,2,2),(96,1,71,4810,1423,476,471,1463,15,6,0,6,0,0,6,2,200,100,200,2,2),(97,1,72,4920,1456,487,482,1496,15,6,0,6,0,0,6,2,200,100,200,2,2),(98,1,73,5030,1489,498,493,1529,15,6,0,6,0,0,6,2,200,100,200,2,2),(99,1,74,5140,1522,509,504,1562,15,6,0,6,0,0,6,2,200,100,200,2,2),(100,1,75,5250,1555,520,515,1595,15,6,0,6,0,0,6,2,200,100,200,2,2),(101,1,76,5360,1588,531,526,1628,15,6,0,6,0,0,6,2,200,100,200,2,2),(102,1,77,5470,1621,542,537,1661,15,6,0,6,0,0,6,2,200,100,200,2,2),(103,1,78,5580,1654,553,548,1694,15,6,0,6,0,0,6,2,200,100,200,2,2),(104,1,79,5690,1687,564,559,1727,15,6,0,6,0,0,6,2,200,100,200,2,2),(105,1,80,5800,1720,575,570,1760,15,6,0,6,0,0,6,2,200,100,200,2,2),(106,1,81,5920,1756,587,582,1796,15,6,0,6,0,0,6,2,200,100,200,2,2),(107,1,82,6040,1792,599,594,1832,15,6,0,6,0,0,6,2,200,100,200,2,2),(108,1,83,6160,1828,611,606,1868,15,6,0,6,0,0,6,2,200,100,200,2,2),(109,1,84,6280,1864,623,618,1904,15,6,0,6,0,0,6,2,200,100,200,2,2),(110,1,85,6400,1900,635,630,1940,15,6,0,6,0,0,6,2,200,100,200,2,2),(111,1,86,6520,1936,647,642,1976,15,6,0,6,0,0,6,2,200,100,200,2,2),(112,1,87,6640,1972,659,654,2012,15,6,0,6,0,0,6,2,200,100,200,2,2),(113,1,88,6760,2008,671,666,2048,15,6,0,6,0,0,6,2,200,100,200,2,2),(114,1,89,6880,2044,683,678,2084,15,6,0,6,0,0,6,2,200,100,200,2,2),(115,1,90,7000,2080,695,690,2120,15,6,0,6,0,0,6,2,200,100,200,2,2),(116,1,91,7130,2119,708,703,2159,15,6,0,6,0,0,6,2,200,100,200,2,2),(117,1,92,7260,2158,721,716,2198,15,6,0,6,0,0,6,2,200,100,200,2,2),(118,1,93,7390,2197,734,729,2237,15,6,0,6,0,0,6,2,200,100,200,2,2),(119,1,94,7520,2236,747,742,2276,15,6,0,6,0,0,6,2,200,100,200,2,2),(120,1,95,7650,2275,760,755,2315,15,6,0,6,0,0,6,2,200,100,200,2,2),(121,1,96,7780,2314,773,768,2354,15,6,0,6,0,0,6,2,200,100,200,2,2),(122,1,97,7910,2353,786,781,2393,15,6,0,6,0,0,6,2,200,100,200,2,2),(123,1,98,8040,2392,799,794,2432,15,6,0,6,0,0,6,2,200,100,200,2,2),(124,1,99,8170,2431,812,807,2471,15,6,0,6,0,0,6,2,200,100,200,2,2),(125,1,100,8300,2470,825,820,2510,15,6,0,6,0,0,6,2,200,100,200,2,2),(126,3,4,1700,220,180,50,50,6,6,0,0,0,0,0,0,100,100,200,2,2),(127,3,5,2000,250,200,60,60,6,6,0,0,0,0,0,0,100,100,200,2,2),(128,3,6,2300,280,220,70,70,6,6,0,0,0,0,0,0,100,100,200,2,2),(129,3,7,2600,310,240,80,80,6,6,0,0,0,0,0,0,100,100,200,2,2),(130,3,8,2900,340,260,90,90,6,6,0,0,0,0,0,0,100,100,200,2,2),(131,3,9,3200,370,280,100,100,6,6,0,0,0,0,0,0,100,100,200,2,2),(132,3,10,3500,400,300,110,110,6,6,0,0,0,0,0,0,100,100,200,2,2),(133,3,11,3830,433,322,121,121,6,6,0,0,0,0,0,0,100,100,200,2,2),(134,3,12,4160,466,344,132,132,6,6,0,0,0,0,0,0,100,100,200,2,2),(135,3,13,4490,499,366,143,143,6,6,0,0,0,0,0,0,100,100,200,2,2),(136,3,14,4820,532,388,154,154,6,6,0,0,0,0,0,0,100,100,200,2,2),(137,3,15,5150,565,410,165,165,6,6,0,0,0,0,0,0,100,100,200,2,2),(138,3,16,5480,598,432,176,176,6,6,0,0,0,0,0,0,100,100,200,2,2),(139,3,17,5810,631,454,187,187,6,6,0,0,0,0,0,0,100,100,200,2,2),(140,3,18,6140,664,476,198,198,6,6,0,0,0,0,0,0,100,100,200,2,2),(141,3,19,6470,697,498,209,209,6,6,0,0,0,0,0,0,100,100,200,2,2),(142,3,20,6800,730,520,220,220,6,6,0,0,0,0,0,0,100,100,200,2,2),(143,3,21,7160,766,544,232,232,6,6,0,0,0,0,0,0,125,100,200,2,2),(144,3,22,7520,802,568,244,244,6,6,0,0,0,0,0,0,125,100,200,2,2),(145,3,23,7880,838,592,256,256,6,6,0,0,0,0,0,0,125,100,200,2,2),(146,3,24,8240,874,616,268,268,6,6,0,0,0,0,0,0,125,100,200,2,2),(147,3,25,8600,910,640,280,280,6,6,0,0,0,0,0,0,125,100,200,2,2),(148,3,26,8960,946,664,292,292,6,6,0,0,0,0,0,0,125,100,200,2,2),(149,3,27,9320,982,688,304,304,6,6,0,0,0,0,0,0,125,100,200,2,2),(150,3,28,9680,1018,712,316,316,6,6,0,0,0,0,0,0,125,100,200,2,2),(151,3,29,10040,1054,736,328,328,6,6,0,0,0,0,0,0,125,100,200,2,2),(152,3,30,10400,1090,760,340,340,6,6,0,0,0,0,0,0,125,100,200,2,2),(153,3,31,10790,1129,786,353,353,6,6,0,0,0,0,0,0,150,100,200,2,2),(154,3,32,11180,1168,812,366,366,6,6,0,0,0,0,0,0,150,100,200,2,2),(155,3,33,11570,1207,838,379,379,6,6,0,0,0,0,0,0,150,100,200,2,2),(156,3,34,11960,1246,864,392,392,6,6,0,0,0,0,0,0,150,100,200,2,2),(157,3,35,12350,1285,890,405,405,6,6,0,0,0,0,0,0,150,100,200,2,2),(158,3,36,12740,1324,916,418,418,6,6,0,0,0,0,0,0,150,100,200,2,2),(159,3,37,13130,1363,942,431,431,6,6,0,0,0,0,0,0,150,100,200,2,2),(160,3,38,13520,1402,968,444,444,6,6,0,0,0,0,0,0,150,100,200,2,2),(161,3,39,13910,1441,994,457,457,6,6,0,0,0,0,0,0,150,100,200,2,2),(162,3,40,14300,1480,1020,470,470,6,6,0,0,0,0,0,0,150,100,200,2,2),(163,3,41,14720,1522,1048,484,484,6,6,0,0,0,0,0,0,175,100,200,2,2),(164,3,42,15140,1564,1076,498,498,6,6,0,0,0,0,0,0,175,100,200,2,2),(165,3,43,15560,1606,1104,512,512,6,6,0,0,0,0,0,0,175,100,200,2,2),(166,3,44,15980,1648,1132,526,526,6,6,0,0,0,0,0,0,175,100,200,2,2),(167,3,45,16400,1690,1160,540,540,6,6,0,0,0,0,0,0,175,100,200,2,2),(168,3,46,16820,1732,1188,554,554,6,6,0,0,0,0,0,0,175,100,200,2,2),(169,3,47,17240,1774,1216,568,568,6,6,0,0,0,0,0,0,175,100,200,2,2),(170,3,48,17660,1816,1244,582,582,6,6,0,0,0,0,0,0,175,100,200,2,2),(171,3,49,18080,1858,1272,596,596,6,6,0,0,0,0,0,0,175,100,200,2,2),(172,3,50,18500,1900,1300,610,610,6,6,0,0,0,0,0,0,175,100,200,2,2),(173,3,51,18950,1945,1330,625,625,6,6,0,0,0,0,0,0,200,100,200,2,2),(174,3,52,19400,1990,1360,640,640,6,6,0,0,0,0,0,0,200,100,200,2,2),(175,3,53,19850,2035,1390,655,655,6,6,0,0,0,0,0,0,200,100,200,2,2),(176,3,54,20300,2080,1420,670,670,6,6,0,0,0,0,0,0,200,100,200,2,2),(177,3,55,20750,2125,1450,685,685,6,6,0,0,0,0,0,0,200,100,200,2,2),(178,3,56,21200,2170,1480,700,700,6,6,0,0,0,0,0,0,200,100,200,2,2),(179,3,57,21650,2215,1510,715,715,6,6,0,0,0,0,0,0,200,100,200,2,2),(180,3,58,22100,2260,1540,730,730,6,6,0,0,0,0,0,0,200,100,200,2,2),(181,3,59,22550,2305,1570,745,745,6,6,0,0,0,0,0,0,200,100,200,2,2),(182,3,60,23000,2350,1600,760,760,6,6,0,0,0,0,0,0,200,100,200,2,2),(183,3,61,23480,2398,1632,776,776,6,6,0,0,0,0,0,0,200,100,200,2,2),(184,3,62,23960,2446,1664,792,792,6,6,0,0,0,0,0,0,200,100,200,2,2),(185,3,63,24440,2494,1696,808,808,6,6,0,0,0,0,0,0,200,100,200,2,2),(186,3,64,24920,2542,1728,824,824,6,6,0,0,0,0,0,0,200,100,200,2,2),(187,3,65,25400,2590,1760,840,840,6,6,0,0,0,0,0,0,200,100,200,2,2),(188,3,66,25880,2638,1792,856,856,6,6,0,0,0,0,0,0,200,100,200,2,2),(189,3,67,26360,2686,1824,872,872,6,6,0,0,0,0,0,0,200,100,200,2,2),(190,3,68,26840,2734,1856,888,888,6,6,0,0,0,0,0,0,200,100,200,2,2),(191,3,69,27320,2782,1888,904,904,6,6,0,0,0,0,0,0,200,100,200,2,2),(192,3,70,27800,2830,1920,920,920,6,6,0,0,0,0,0,0,200,100,200,2,2),(193,3,71,28310,2881,1954,937,937,6,6,0,0,0,0,0,0,200,100,200,2,2),(194,3,72,28820,2932,1988,954,954,6,6,0,0,0,0,0,0,200,100,200,2,2),(195,3,73,29330,2983,2022,971,971,6,6,0,0,0,0,0,0,200,100,200,2,2),(196,3,74,29840,3034,2056,988,988,6,6,0,0,0,0,0,0,200,100,200,2,2),(197,3,75,30350,3085,2090,1005,1005,6,6,0,0,0,0,0,0,200,100,200,2,2),(198,3,76,30860,3136,2124,1022,1022,6,6,0,0,0,0,0,0,200,100,200,2,2),(199,3,77,31370,3187,2158,1039,1039,6,6,0,0,0,0,0,0,200,100,200,2,2),(200,3,78,31880,3238,2192,1056,1056,6,6,0,0,0,0,0,0,200,100,200,2,2),(201,3,79,32390,3289,2226,1073,1073,6,6,0,0,0,0,0,0,200,100,200,2,2),(202,3,80,32900,3340,2260,1090,1090,6,6,0,0,0,0,0,0,200,100,200,2,2),(203,3,81,33440,3394,2296,1108,1108,6,6,0,0,0,0,0,0,200,100,200,2,2),(204,3,82,33980,3448,2332,1126,1126,6,6,0,0,0,0,0,0,200,100,200,2,2),(205,3,83,34520,3502,2368,1144,1144,6,6,0,0,0,0,0,0,200,100,200,2,2),(206,3,84,35060,3556,2404,1162,1162,6,6,0,0,0,0,0,0,200,100,200,2,2),(207,3,85,35600,3610,2440,1180,1180,6,6,0,0,0,0,0,0,200,100,200,2,2),(208,3,86,36140,3664,2476,1198,1198,6,6,0,0,0,0,0,0,200,100,200,2,2),(209,3,87,36680,3718,2512,1216,1216,6,6,0,0,0,0,0,0,200,100,200,2,2),(210,3,88,37220,3772,2548,1234,1234,6,6,0,0,0,0,0,0,200,100,200,2,2),(211,3,89,37760,3826,2584,1252,1252,6,6,0,0,0,0,0,0,200,100,200,2,2),(212,3,90,38300,3880,2620,1270,1270,6,6,0,0,0,0,0,0,200,100,200,2,2),(213,3,91,38870,3937,2658,1289,1289,6,6,0,0,0,0,0,0,200,100,200,2,2),(214,3,92,39440,3994,2696,1308,1308,6,6,0,0,0,0,0,0,200,100,200,2,2),(215,3,93,40010,4051,2734,1327,1327,6,6,0,0,0,0,0,0,200,100,200,2,2),(216,3,94,40580,4108,2772,1346,1346,6,6,0,0,0,0,0,0,200,100,200,2,2),(217,3,95,41150,4165,2810,1365,1365,6,6,0,0,0,0,0,0,200,100,200,2,2),(218,3,96,41720,4222,2848,1384,1384,6,6,0,0,0,0,0,0,200,100,200,2,2),(219,3,97,42290,4279,2886,1403,1403,6,6,0,0,0,0,0,0,200,100,200,2,2),(220,3,98,42860,4336,2924,1422,1422,6,6,0,0,0,0,0,0,200,100,200,2,2),(221,3,99,43430,4393,2962,1441,1441,6,6,0,0,0,0,0,0,200,100,200,2,2),(222,3,100,44000,4450,3000,1460,1460,6,6,0,0,0,0,0,0,200,100,200,2,2),(223,4,4,500,130,50,220,90,6,6,0,6,0,0,0,0,100,100,200,2,2),(224,4,5,600,160,60,250,110,6,6,0,6,0,0,0,0,100,100,200,2,2),(225,4,6,700,190,70,280,130,6,6,0,6,0,0,0,0,100,100,200,2,2),(226,4,7,800,220,80,310,150,6,6,0,6,0,0,0,0,100,100,200,2,2),(227,4,8,900,250,90,340,170,6,6,0,6,0,0,0,0,100,100,200,2,2),(228,4,9,1000,280,100,370,190,6,6,0,6,0,0,0,0,100,100,200,2,2),(229,4,10,1100,310,110,400,210,6,6,0,6,0,0,0,0,100,100,200,2,2),(230,4,11,1210,343,121,433,232,6,6,0,6,0,0,0,0,100,100,200,2,2),(231,4,12,1320,376,132,466,254,6,6,0,6,0,0,0,0,100,100,200,2,2),(232,4,13,1430,409,143,499,276,6,6,0,6,0,0,0,0,100,100,200,2,2),(233,4,14,1540,442,154,532,298,6,6,0,6,0,0,0,0,100,100,200,2,2),(234,4,15,1650,475,165,565,320,6,6,0,6,0,0,0,0,100,100,200,2,2),(235,4,16,1760,508,176,598,342,6,6,0,6,0,0,0,0,100,100,200,2,2),(236,4,17,1870,541,187,631,364,6,6,0,6,0,0,0,0,100,100,200,2,2),(237,4,18,1980,574,198,664,386,6,6,0,6,0,0,0,0,100,100,200,2,2),(238,4,19,2090,607,209,697,408,6,6,0,6,0,0,0,0,100,100,200,2,2),(239,4,20,2200,640,220,730,430,6,6,0,6,0,0,0,0,100,100,200,2,2),(240,4,21,2320,676,232,766,454,6,6,0,6,0,0,0,0,125,100,200,2,2),(241,4,22,2440,712,244,802,478,6,6,0,6,0,0,0,0,125,100,200,2,2),(242,4,23,2560,748,256,838,502,6,6,0,6,0,0,0,0,125,100,200,2,2),(243,4,24,2680,784,268,874,526,6,6,0,6,0,0,0,0,125,100,200,2,2),(244,4,25,2800,820,280,910,550,6,6,0,6,0,0,0,0,125,100,200,2,2),(245,4,26,2920,856,292,946,574,6,6,0,6,0,0,0,0,125,100,200,2,2),(246,4,27,3040,892,304,982,598,6,6,0,6,0,0,0,0,125,100,200,2,2),(247,4,28,3160,928,316,1018,622,6,6,0,6,0,0,0,0,125,100,200,2,2),(248,4,29,3280,964,328,1054,646,6,6,0,6,0,0,0,0,125,100,200,2,2),(249,4,30,3400,1000,340,1090,670,6,6,0,6,0,0,0,0,125,100,200,2,2),(250,4,31,3530,1039,353,1129,696,6,6,0,6,0,0,0,0,150,100,200,2,2),(251,4,32,3660,1078,366,1168,722,6,6,0,6,0,0,0,0,150,100,200,2,2),(252,4,33,3790,1117,379,1207,748,6,6,0,6,0,0,0,0,150,100,200,2,2),(253,4,34,3920,1156,392,1246,774,6,6,0,6,0,0,0,0,150,100,200,2,2),(254,4,35,4050,1195,405,1285,800,6,6,0,6,0,0,0,0,150,100,200,2,2),(255,4,36,4180,1234,418,1324,826,6,6,0,6,0,0,0,0,150,100,200,2,2),(256,4,37,4310,1273,431,1363,852,6,6,0,6,0,0,0,0,150,100,200,2,2),(257,4,38,4440,1312,444,1402,878,6,6,0,6,0,0,0,0,150,100,200,2,2),(258,4,39,4570,1351,457,1441,904,6,6,0,6,0,0,0,0,150,100,200,2,2),(259,4,40,4700,1390,470,1480,930,6,6,0,6,0,0,0,0,150,100,200,2,2),(260,4,41,4840,1432,484,1522,958,6,6,0,6,0,0,0,0,175,100,200,2,2),(261,4,42,4980,1474,498,1564,986,6,6,0,6,0,0,0,0,175,100,200,2,2),(262,4,43,5120,1516,512,1606,1014,6,6,0,6,0,0,0,0,175,100,200,2,2),(263,4,44,5260,1558,526,1648,1042,6,6,0,6,0,0,0,0,175,100,200,2,2),(264,4,45,5400,1600,540,1690,1070,6,6,0,6,0,0,0,0,175,100,200,2,2),(265,4,46,5540,1642,554,1732,1098,6,6,0,6,0,0,0,0,175,100,200,2,2),(266,4,47,5680,1684,568,1774,1126,6,6,0,6,0,0,0,0,175,100,200,2,2),(267,4,48,5820,1726,582,1816,1154,6,6,0,6,0,0,0,0,175,100,200,2,2),(268,4,49,5960,1768,596,1858,1182,6,6,0,6,0,0,0,0,175,100,200,2,2),(269,4,50,6100,1810,610,1900,1210,6,6,0,6,0,0,0,0,175,100,200,2,2),(270,4,51,6250,1855,625,1945,1240,6,6,0,6,0,0,0,0,200,100,200,2,2),(271,4,52,6400,1900,640,1990,1270,6,6,0,6,0,0,0,0,200,100,200,2,2),(272,4,53,6550,1945,655,2035,1300,6,6,0,6,0,0,0,0,200,100,200,2,2),(273,4,54,6700,1990,670,2080,1330,6,6,0,6,0,0,0,0,200,100,200,2,2),(274,4,55,6850,2035,685,2125,1360,6,6,0,6,0,0,0,0,200,100,200,2,2),(275,4,56,7000,2080,700,2170,1390,6,6,0,6,0,0,0,0,200,100,200,2,2),(276,4,57,7150,2125,715,2215,1420,6,6,0,6,0,0,0,0,200,100,200,2,2),(277,4,58,7300,2170,730,2260,1450,6,6,0,6,0,0,0,0,200,100,200,2,2),(278,4,59,7450,2215,745,2305,1480,6,6,0,6,0,0,0,0,200,100,200,2,2),(279,4,60,7600,2260,760,2350,1510,6,6,0,6,0,0,0,0,200,100,200,2,2),(280,4,61,7760,2308,776,2398,1542,6,6,0,6,0,0,0,0,200,100,200,2,2),(281,4,62,7920,2356,792,2446,1574,6,6,0,6,0,0,0,0,200,100,200,2,2),(282,5,1,650,145,120,120,20,6,0,0,6,0,0,0,0,200,100,200,2,2),(283,5,2,800,190,140,140,30,6,0,0,6,0,0,0,0,200,100,200,2,2),(284,4,63,8080,2404,808,2494,1606,6,6,0,6,0,0,0,0,200,100,200,2,2),(285,4,64,8240,2452,824,2542,1638,6,6,0,6,0,0,0,0,200,100,200,2,2),(286,4,65,8400,2500,840,2590,1670,6,6,0,6,0,0,0,0,200,100,200,2,2),(287,4,66,8560,2548,856,2638,1702,6,6,0,6,0,0,0,0,200,100,200,2,2),(288,4,67,8720,2596,872,2686,1734,6,6,0,6,0,0,0,0,200,100,200,2,2),(289,4,68,8880,2644,888,2734,1766,6,6,0,6,0,0,0,0,200,100,200,2,2),(290,4,69,9040,2692,904,2782,1798,6,6,0,6,0,0,0,0,200,100,200,2,2),(291,4,70,9200,2740,920,2830,1830,6,6,0,6,0,0,0,0,200,100,200,2,2),(292,4,71,9370,2791,937,2881,1864,6,6,0,6,0,0,0,0,200,100,200,2,2),(293,5,3,950,235,160,160,40,6,0,0,6,0,0,0,0,200,100,200,2,2),(294,4,72,9540,2842,954,2932,1898,6,6,0,6,0,0,0,0,200,100,200,2,2),(295,4,73,9710,2893,971,2983,1932,6,6,0,6,0,0,0,0,200,100,200,2,2),(296,4,74,9880,2944,988,3034,1966,6,6,0,6,0,0,0,0,200,100,200,2,2),(297,4,75,10050,2995,1005,3085,2000,6,6,0,6,0,0,0,0,200,100,200,2,2),(298,6,1,600,160,110,130,20,3,0,6,0,0,0,0,0,200,100,200,2,2),(299,6,2,700,220,120,160,30,3,0,6,0,0,0,0,0,200,100,200,2,2),(300,6,3,800,280,130,190,40,3,0,6,0,0,0,0,0,200,100,200,2,2),(301,4,76,10220,3046,1022,3136,2034,6,6,0,6,0,0,0,0,200,100,200,2,2),(302,4,77,10390,3097,1039,3187,2068,6,6,0,6,0,0,0,0,200,100,200,2,2),(303,4,78,10560,3148,1056,3238,2102,6,6,0,6,0,0,0,0,200,100,200,2,2),(304,4,79,10730,3199,1073,3289,2136,6,6,0,6,0,0,0,0,200,100,200,2,2),(305,4,80,10900,3250,1090,3340,2170,6,6,0,6,0,0,0,0,200,100,200,2,2),(306,4,81,11080,3304,1108,3394,2206,6,6,0,6,0,0,0,0,200,100,200,2,2),(307,4,82,11260,3358,1126,3448,2242,6,6,0,6,0,0,0,0,200,100,200,2,2),(308,4,83,11440,3412,1144,3502,2278,6,6,0,6,0,0,0,0,200,100,200,2,2),(309,4,84,11620,3466,1162,3556,2314,6,6,0,6,0,0,0,0,200,100,200,2,2),(310,4,85,11800,3520,1180,3610,2350,6,6,0,6,0,0,0,0,200,100,200,2,2),(311,4,86,11980,3574,1198,3664,2386,6,6,0,6,0,0,0,0,200,100,200,2,2),(312,4,87,12160,3628,1216,3718,2422,6,6,0,6,0,0,0,0,200,100,200,2,2),(313,4,88,12340,3682,1234,3772,2458,6,6,0,6,0,0,0,0,200,100,200,2,2),(314,4,89,12520,3736,1252,3826,2494,6,6,0,6,0,0,0,0,200,100,200,2,2),(315,4,90,12700,3790,1270,3880,2530,6,6,0,6,0,0,0,0,200,100,200,2,2),(316,4,91,12890,3847,1289,3937,2568,6,6,0,6,0,0,0,0,200,100,200,2,2),(317,4,92,13080,3904,1308,3994,2606,6,6,0,6,0,0,0,0,200,100,200,2,2),(318,4,93,13270,3961,1327,4051,2644,6,6,0,6,0,0,0,0,200,100,200,2,2),(319,4,94,13460,4018,1346,4108,2682,6,6,0,6,0,0,0,0,200,100,200,2,2),(320,4,95,13650,4075,1365,4165,2720,6,6,0,6,0,0,0,0,200,100,200,2,2),(321,4,96,13840,4132,1384,4222,2758,6,6,0,6,0,0,0,0,200,100,200,2,2),(322,4,97,14030,4189,1403,4279,2796,6,6,0,6,0,0,0,0,200,100,200,2,2),(323,4,98,14220,4246,1422,4336,2834,6,6,0,6,0,0,0,0,200,100,200,2,2),(324,4,99,14410,4303,1441,4393,2872,6,6,0,6,0,0,0,0,200,100,200,2,2),(325,4,100,14600,4360,1460,4450,2910,6,6,0,6,0,0,0,0,200,100,200,2,2),(326,5,4,1100,280,180,180,50,6,0,0,6,0,0,0,0,200,100,200,2,2),(327,5,5,1250,325,200,200,60,6,0,0,6,0,0,0,0,200,100,200,2,2),(328,5,6,1400,370,220,220,70,6,0,0,6,0,0,0,0,200,100,200,2,2),(329,5,7,1550,415,240,240,80,6,0,0,6,0,0,0,0,200,100,200,2,2),(330,5,8,1700,460,260,260,90,6,0,0,6,0,0,0,0,200,100,200,2,2),(331,5,9,1850,505,280,280,100,6,0,0,6,0,0,0,0,200,100,200,2,2),(332,5,10,2000,550,300,300,110,6,0,0,6,0,0,0,0,200,100,200,2,2),(333,5,11,2165,600,322,322,121,6,0,0,6,0,0,0,0,200,100,200,2,2),(334,5,12,2330,650,344,344,132,6,0,0,6,0,0,0,0,200,100,200,2,2),(335,5,13,2495,700,366,366,143,6,0,0,6,0,0,0,0,200,100,200,2,2),(336,5,14,2660,750,388,388,154,6,0,0,6,0,0,0,0,200,100,200,2,2),(337,5,15,2825,800,410,410,165,6,0,0,6,0,0,0,0,200,100,200,2,2),(338,5,16,2990,850,432,432,176,6,0,0,6,0,0,0,0,200,100,200,2,2),(339,5,17,3155,900,454,454,187,6,0,0,6,0,0,0,0,200,100,200,2,2),(340,5,20,3650,1050,520,520,220,6,0,0,6,0,0,0,0,200,100,200,2,2),(341,5,19,3485,1000,498,498,209,6,0,0,6,0,0,0,0,200,100,200,2,2),(342,5,18,3320,950,476,476,198,6,0,0,6,0,0,0,0,200,100,200,2,2),(343,5,21,3830,1104,544,544,232,6,0,0,6,0,0,0,0,250,100,200,2,2),(344,5,22,4010,1158,568,568,244,6,0,0,6,0,0,0,0,250,100,200,2,2),(345,5,23,4190,1212,592,592,256,6,0,0,6,0,0,0,0,250,100,200,2,2),(346,5,24,4370,1266,616,616,268,6,0,0,6,0,0,0,0,250,100,200,2,2),(347,5,25,4550,1320,640,640,280,6,0,0,6,0,0,0,0,250,100,200,2,2),(348,5,26,4730,1374,664,664,292,6,0,0,6,0,0,0,0,250,100,200,2,2),(349,5,27,4910,1428,688,688,304,6,0,0,6,0,0,0,0,250,100,200,2,2),(350,5,28,5090,1482,712,712,316,6,0,0,6,0,0,0,0,250,100,200,2,2),(351,5,29,5270,1536,736,736,328,6,0,0,6,0,0,0,0,250,100,200,2,2),(352,5,30,5450,1590,760,760,340,6,0,0,6,0,0,0,0,250,100,200,2,2),(353,5,31,5645,1649,786,786,353,6,0,0,6,0,0,0,0,300,100,200,2,2),(354,5,32,5840,1708,812,812,366,6,0,0,6,0,0,0,0,300,100,200,2,2),(355,5,33,6035,1767,838,838,379,6,0,0,6,0,0,0,0,300,100,200,2,2),(356,5,34,6230,1826,864,864,392,6,0,0,6,0,0,0,0,300,100,200,2,2),(360,6,4,900,340,140,220,50,3,0,6,0,0,0,0,0,200,100,200,2,2),(361,6,5,1000,400,150,250,60,3,0,6,0,0,0,0,0,200,100,200,2,2),(362,6,6,1100,460,160,280,70,3,0,6,0,0,0,0,0,200,100,200,2,2),(363,6,7,1200,520,170,310,80,3,0,6,0,0,0,0,0,200,100,200,2,2),(364,6,8,1300,580,180,340,90,3,0,6,0,0,0,0,0,200,100,200,2,2),(365,6,9,1400,640,190,370,100,3,0,6,0,0,0,0,0,200,100,200,2,2),(366,6,10,1500,700,200,400,110,3,0,6,0,0,0,0,0,200,100,200,2,2),(367,6,11,1610,766,211,433,121,3,0,6,0,0,0,0,0,200,100,200,2,2),(369,6,12,1720,832,222,466,132,3,0,6,0,0,0,0,0,200,100,200,2,2),(370,6,13,1830,898,233,499,143,3,0,6,0,0,0,0,0,200,100,200,2,2),(371,6,14,1940,964,244,532,154,3,0,6,0,0,0,0,0,200,100,200,2,2),(372,6,15,2050,1030,255,565,165,3,0,6,0,0,0,0,0,200,100,200,2,2),(373,6,16,2160,1096,266,598,176,3,0,6,0,0,0,0,0,200,100,200,2,2),(374,6,17,2270,1162,277,631,187,3,0,6,0,0,0,0,0,200,100,200,2,2),(375,6,18,2380,1228,288,664,198,3,0,6,0,0,0,0,0,200,100,200,2,2),(376,6,19,2490,1294,299,697,209,3,0,6,0,0,0,0,0,200,100,200,2,2),(377,6,20,2600,1360,310,730,220,3,0,6,0,0,0,0,0,200,100,200,2,2),(378,6,21,2720,1432,322,766,232,3,0,6,0,0,0,0,0,250,100,200,2,2),(379,6,22,2840,1504,334,802,244,3,0,6,0,0,0,0,0,250,100,200,2,2),(380,6,23,2960,1576,346,838,256,3,0,6,0,0,0,0,0,250,100,200,2,2),(381,6,24,3080,1648,358,874,268,3,0,6,0,0,0,0,0,250,100,200,2,2),(382,6,25,3200,1720,370,910,280,3,0,6,0,0,0,0,0,250,100,200,2,2),(383,6,26,3320,1792,382,946,292,3,0,6,0,0,0,0,0,250,100,200,2,2),(384,6,27,3440,1864,394,982,304,3,0,6,0,0,0,0,0,250,100,200,2,2),(385,6,28,3560,1936,406,1018,316,3,0,6,0,0,0,0,0,250,100,200,2,2),(386,6,29,3680,2008,418,1054,328,3,0,6,0,0,0,0,0,250,100,200,2,2),(387,6,30,3800,2080,430,1090,340,3,0,6,0,0,0,0,0,250,100,200,2,2),(388,6,31,3930,2158,443,1129,353,3,0,6,0,0,0,0,0,300,100,200,2,2),(389,6,32,4060,2236,456,1168,366,3,0,6,0,0,0,0,0,300,100,200,2,2),(390,6,33,4190,2314,469,1207,379,3,0,6,0,0,0,0,0,300,100,200,2,2),(391,6,34,4320,2392,482,1246,392,3,0,6,0,0,0,0,0,300,100,200,2,2),(392,6,35,4450,2470,495,1285,405,3,0,6,0,0,0,0,0,300,100,200,2,2),(393,6,36,4580,2548,508,1324,418,3,0,6,0,0,0,0,0,300,100,200,2,2),(394,6,37,4710,2626,521,1363,431,3,0,6,0,0,0,0,0,300,100,200,2,2),(395,6,38,4840,2704,534,1402,444,3,0,6,0,0,0,0,0,300,100,200,2,2),(396,6,39,4970,2782,547,1441,457,3,0,6,0,0,0,0,0,300,100,200,2,2),(397,6,40,5100,2860,560,1480,470,3,0,6,0,0,0,0,0,300,100,200,2,2),(398,6,41,5240,2944,574,1522,484,3,0,6,0,0,0,0,0,350,100,200,2,2),(399,6,42,5380,3028,588,1564,498,3,0,6,0,0,0,0,0,350,100,200,2,2),(400,6,43,5520,3112,602,1606,512,3,0,6,0,0,0,0,0,350,100,200,2,2),(401,6,44,5660,3190,616,1648,526,3,0,6,0,0,0,0,0,350,100,200,2,2),(402,6,45,5800,3280,630,1690,540,3,0,6,0,0,0,0,0,350,100,200,2,2),(403,6,46,5940,3364,644,1732,554,3,0,6,0,0,0,0,0,350,100,200,2,2),(404,6,47,6080,3448,658,1774,568,3,0,6,0,0,0,0,0,350,100,200,2,2),(405,6,48,6226,3532,672,1816,582,3,0,6,0,0,0,0,0,350,100,200,2,2),(406,6,49,6360,3616,686,1858,596,3,0,6,0,0,0,0,0,350,100,200,2,2),(407,6,50,6500,3700,700,1900,610,3,0,6,0,0,0,0,0,350,100,200,2,2),(409,6,51,6650,3790,715,1945,625,3,0,6,0,0,0,0,0,400,100,200,2,2),(410,6,52,6800,3880,730,1990,640,3,0,6,0,0,0,0,0,400,100,200,2,2),(411,6,53,6950,3970,745,2035,655,3,0,6,0,0,0,0,0,400,100,200,2,2),(412,6,54,7100,4060,760,2080,670,3,0,6,0,0,0,0,0,400,100,200,2,2),(413,6,55,7250,4150,775,2125,685,3,0,6,0,0,0,0,0,400,100,200,2,2),(414,6,56,7400,4240,790,2170,700,3,0,6,0,0,0,0,0,400,100,200,2,2),(415,6,57,7550,4330,805,2215,715,3,0,6,0,0,0,0,0,400,100,200,2,2),(416,6,58,7700,4420,820,2260,730,3,0,6,0,0,0,0,0,400,100,200,2,2),(417,6,59,7850,4510,835,2305,745,3,0,6,0,0,0,0,0,400,100,200,2,2),(418,6,60,8000,4600,850,2350,760,3,0,6,0,0,0,0,0,400,100,200,2,2),(419,6,61,8160,4696,866,2398,776,3,0,6,0,0,0,0,0,400,100,200,2,2),(420,6,62,8320,4792,882,2446,792,3,0,6,0,0,0,0,0,400,100,200,2,2),(421,6,63,8480,4888,898,2494,808,3,0,6,0,0,0,0,0,400,100,200,2,2),(422,6,64,8640,4984,914,2542,824,3,0,6,0,0,0,0,0,400,100,200,2,2),(423,6,65,8800,5080,930,2590,840,3,0,6,0,0,0,0,0,400,100,200,2,2),(424,6,66,8960,5176,946,2638,856,3,0,6,0,0,0,0,0,400,100,200,2,2),(425,6,67,9120,5272,962,2686,872,3,0,6,0,0,0,0,0,400,100,200,2,2),(426,6,68,9280,5368,978,2734,888,3,0,6,0,0,0,0,0,400,100,200,2,2),(427,6,69,9440,5464,994,2782,904,3,0,6,0,0,0,0,0,400,100,200,2,2),(428,6,70,9600,5560,1010,2830,920,3,0,6,0,0,0,0,0,400,100,200,2,2),(430,6,71,9770,5662,1027,2881,937,3,0,6,0,0,0,0,0,400,100,200,2,2),(431,6,72,9940,5764,1044,2932,954,3,0,6,0,0,0,0,0,400,100,200,2,2),(432,6,73,10110,5866,1061,2983,971,3,0,6,0,0,0,0,0,400,100,200,2,2),(433,6,74,10280,5968,1078,3034,988,3,0,6,0,0,0,0,0,400,100,200,2,2),(434,6,75,10450,6070,1095,3085,1005,3,0,6,0,0,0,0,0,400,100,200,2,2),(435,6,76,10620,6172,1112,3136,1022,3,0,6,0,0,0,0,0,400,100,200,2,2),(436,6,77,10790,6274,1129,3187,1039,3,0,6,0,0,0,0,0,400,100,200,2,2),(437,6,78,10960,6376,1146,3238,1056,3,0,6,0,0,0,0,0,400,100,200,2,2),(438,6,79,11130,6478,1163,3289,1073,3,0,6,0,0,0,0,0,400,100,200,2,2),(439,6,80,11300,6580,1180,3340,1090,3,0,6,0,0,0,0,0,400,100,200,2,2),(441,6,81,11480,6688,1198,3394,1108,3,0,6,0,0,0,0,0,400,100,200,2,2),(442,6,82,11660,6796,1216,3448,1126,3,0,6,0,0,0,0,0,400,100,200,2,2),(443,6,83,11840,6904,1234,3502,1144,3,0,6,0,0,0,0,0,400,100,200,2,2),(444,6,84,12020,7012,1252,3556,1162,3,0,6,0,0,0,0,0,400,100,200,2,2),(445,6,85,12200,7120,1270,3610,1180,3,0,6,0,0,0,0,0,400,100,200,2,2),(446,6,86,12380,7228,1288,3664,1198,3,0,6,0,0,0,0,0,400,100,200,2,2),(447,6,87,12560,7336,1306,3718,1216,3,0,6,0,0,0,0,0,400,100,200,2,2),(448,6,88,12740,7444,1324,3772,1234,3,0,6,0,0,0,0,0,400,100,200,2,2),(449,6,89,12920,7552,1342,3826,1252,3,0,6,0,0,0,0,0,400,100,200,2,2),(450,6,90,13100,7660,1360,3880,1270,3,0,6,0,0,0,0,0,400,100,200,2,2),(451,6,91,13290,7774,1379,3937,1289,3,0,6,0,0,0,0,0,400,100,200,2,2),(452,6,92,13480,7888,1398,3994,1308,3,0,6,0,0,0,0,0,400,100,200,2,2),(453,6,93,13670,8002,1417,4051,1327,3,0,6,0,0,0,0,0,400,100,200,2,2),(454,6,94,13860,8116,1436,4108,1346,3,0,6,0,0,0,0,0,400,100,200,2,2),(455,6,95,14050,8230,1455,4165,1365,3,0,6,0,0,0,0,0,400,100,200,2,2),(456,6,96,14240,8344,1474,4222,1384,3,0,6,0,0,0,0,0,400,100,200,2,2),(457,6,97,14430,8458,1493,4279,1403,3,0,6,0,0,0,0,0,400,100,200,2,2),(458,6,98,14620,8572,1512,4336,1422,3,0,6,0,0,0,0,0,400,100,200,2,2),(459,6,99,14810,8686,1531,4393,1441,3,0,6,0,0,0,0,0,400,100,200,2,2),(460,6,100,15000,8800,1550,4450,1460,3,0,6,0,0,0,0,0,400,100,200,2,2),(461,5,35,6425,1885,890,890,405,6,0,0,6,0,0,0,0,300,100,200,2,2),(462,5,36,6620,1944,916,916,418,6,0,0,6,0,0,0,0,300,100,200,2,2),(463,5,37,6815,2003,942,942,431,6,0,0,6,0,0,0,0,300,100,200,2,2),(464,5,38,7010,2062,968,968,444,6,0,0,6,0,0,0,0,300,100,200,2,2),(465,5,39,7205,2121,994,994,457,6,0,0,6,0,0,0,0,300,100,200,2,2),(466,5,40,7400,2180,1020,1020,470,6,0,0,6,0,0,0,0,300,100,200,2,2),(467,5,41,7610,2243,1048,1048,484,6,0,0,6,0,0,0,0,350,100,200,2,2),(468,5,42,7820,2306,1076,1076,498,6,0,0,6,0,0,0,0,350,100,200,2,2),(469,5,43,8030,2369,1104,1104,512,6,0,0,6,0,0,0,0,350,100,200,2,2),(470,5,44,8240,2432,1132,1132,526,6,0,0,6,0,0,0,0,350,100,200,2,2),(471,5,45,8450,2495,1160,1160,540,6,0,0,6,0,0,0,0,350,100,200,2,2),(472,5,46,8660,2558,1188,1188,554,6,0,0,6,0,0,0,0,350,100,200,2,2),(473,5,47,8870,2621,1216,1216,568,6,0,0,6,0,0,0,0,350,100,200,2,2),(474,5,48,9080,2684,1244,1244,582,6,0,0,6,0,0,0,0,350,100,200,2,2),(475,5,49,9290,2749,1272,1272,596,6,0,0,6,0,0,0,0,350,100,200,2,2),(476,5,50,9500,2810,1300,1300,610,6,0,0,6,0,0,0,0,350,100,200,2,2),(477,5,51,9725,2878,1330,1330,625,6,0,0,6,0,0,0,0,400,100,200,2,2),(478,5,52,9950,2946,1360,1360,640,6,0,0,6,0,0,0,0,400,100,200,2,2),(479,5,53,10175,3014,1390,1390,655,6,0,0,6,0,0,0,0,400,100,200,2,2),(480,5,54,10400,3082,1420,1420,670,6,0,0,6,0,0,0,0,400,100,200,2,2),(481,5,55,10625,3150,1450,1450,685,6,0,0,6,0,0,0,0,400,100,200,2,2),(482,5,56,10850,3218,1480,1480,700,6,0,0,6,0,0,0,0,400,100,200,2,2),(483,5,57,11075,3286,1510,1510,715,6,0,0,6,0,0,0,0,400,100,200,2,2),(484,5,58,11300,3354,1540,1540,730,6,0,0,6,0,0,0,0,400,100,200,2,2),(485,5,59,11525,3422,1570,1570,745,6,0,0,6,0,0,0,0,400,100,200,2,2),(486,5,60,11750,3490,1600,1600,760,6,0,0,6,0,0,0,0,400,100,200,2,2),(487,5,61,11990,3562,1632,1632,776,6,0,0,6,0,0,0,0,400,100,200,2,2),(488,5,62,12230,3634,1664,1664,792,6,0,0,6,0,0,0,0,400,100,200,2,2),(489,5,63,12470,3706,1696,1696,808,6,0,0,6,0,0,0,0,400,100,200,2,2),(490,5,64,12710,3778,1728,1728,824,6,0,0,6,0,0,0,0,400,100,200,2,2),(491,5,65,12950,3850,1760,1760,840,6,0,0,6,0,0,0,0,400,100,200,2,2),(492,5,66,13190,3922,1792,1792,856,6,0,0,6,0,0,0,0,400,100,200,2,2),(493,5,67,13430,3994,1824,1824,872,6,0,0,6,0,0,0,0,400,100,200,2,2),(494,5,68,13670,4066,1856,1856,888,6,0,0,6,0,0,0,0,400,100,200,2,2),(495,5,69,13910,4138,1888,1888,904,6,0,0,6,0,0,0,0,400,100,200,2,2),(496,5,70,14150,4210,1920,1920,920,6,0,0,6,0,0,0,0,400,100,200,2,2),(497,5,71,14405,4287,1954,1954,937,6,0,0,6,0,0,0,0,400,100,200,2,2),(498,5,72,14660,4364,1988,1988,954,6,0,0,6,0,0,0,0,400,100,200,2,2),(499,5,73,14915,4441,2022,2022,971,6,0,0,6,0,0,0,0,400,100,200,2,2),(500,5,74,15170,4518,2056,2056,988,6,0,0,6,0,0,0,0,400,100,200,2,2),(501,5,75,15425,4595,2090,2090,1005,6,0,0,6,0,0,0,0,400,100,200,2,2),(502,5,76,15680,4672,2124,2124,1022,6,0,0,6,0,0,0,0,400,100,200,2,2),(503,5,77,15935,4749,2158,2158,1039,6,0,0,6,0,0,0,0,400,100,200,2,2),(504,5,78,16190,4826,2192,2192,1056,6,0,0,6,0,0,0,0,400,100,200,2,2),(505,5,79,16445,4903,2226,2226,1073,6,0,0,6,0,0,0,0,400,100,200,2,2),(506,5,80,16700,4980,2260,2260,1090,6,0,0,6,0,0,0,0,400,100,200,2,2),(507,5,81,16970,5061,2296,2296,1108,6,0,0,6,0,0,0,0,400,100,200,2,2),(508,5,82,17240,5142,2332,2332,1126,6,0,0,6,0,0,0,0,400,100,200,2,2),(509,5,83,17510,5223,2368,2368,1144,6,0,0,6,0,0,0,0,400,100,200,2,2),(510,5,84,17780,5304,2404,2404,1162,6,0,0,6,0,0,0,0,400,100,200,2,2),(511,5,85,18050,5385,2440,2440,1180,6,0,0,6,0,0,0,0,400,100,200,2,2),(512,5,86,18320,5466,2476,2476,1198,6,0,0,6,0,0,0,0,400,100,200,2,2),(513,5,87,18590,5547,2512,2512,1216,6,0,0,6,0,0,0,0,400,100,200,2,2),(514,5,88,18860,5628,2548,2548,1234,6,0,0,6,0,0,0,0,400,100,200,2,2),(515,5,89,19130,5709,2584,2584,1252,6,0,0,6,0,0,0,0,400,100,200,2,2),(516,5,90,19400,5790,2620,2620,1270,6,0,0,6,0,0,0,0,400,100,200,2,2),(517,5,91,19685,5876,2658,2658,1289,6,0,0,6,0,0,0,0,400,100,200,2,2),(518,5,92,19970,5962,2696,2696,1308,6,0,0,6,0,0,0,0,400,100,200,2,2),(519,5,93,20255,6048,2734,2734,1327,6,0,0,6,0,0,0,0,400,100,200,2,2),(520,5,94,20540,6134,2772,2772,1346,6,0,0,6,0,0,0,0,400,100,200,2,2),(521,5,95,20825,6220,2810,2810,1365,6,0,0,6,0,0,0,0,400,100,200,2,2),(522,5,96,21110,6306,2848,2848,1384,6,0,0,6,0,0,0,0,400,100,200,2,2),(523,5,97,21395,6392,2886,2886,1403,6,0,0,6,0,0,0,0,400,100,200,2,2),(524,5,98,21680,6478,2924,2924,1422,6,0,0,6,0,0,0,0,400,100,200,2,2),(525,5,99,21965,6564,2962,2962,1441,6,0,0,6,0,0,0,0,400,100,200,2,2),(526,5,100,22250,6650,3000,3000,1460,6,0,0,6,0,0,0,0,400,100,200,2,2);
/*!40000 ALTER TABLE `role_level` ENABLE KEYS */;
DROP TABLE IF EXISTS `role_level_exp`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role_level_exp` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `level` int(11) NOT NULL COMMENT '等级 - level',
  `exp` bigint(20) NOT NULL COMMENT '升到下一级所需经验',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb4 COMMENT='角色等级经验表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `role_level_exp` DISABLE KEYS */;
INSERT INTO `role_level_exp` VALUES (1,1,10),(2,2,20),(3,3,30),(4,4,40),(5,5,50),(6,6,60),(7,7,60),(8,8,80),(9,9,90),(10,10,100),(11,11,200),(12,12,300),(13,13,400),(14,14,500),(15,15,600),(16,16,700),(17,17,1000),(18,18,1200),(19,19,1300),(20,20,1500),(21,21,10000),(22,22,11000),(23,23,12000),(24,24,13500),(25,25,15000),(26,26,17000),(27,27,19000),(28,28,21500),(29,29,24000),(30,30,27000),(31,31,35000),(32,32,37000),(33,33,40000),(34,34,44000),(35,35,48000),(36,36,52000),(37,37,56000),(38,38,60000),(39,39,65000),(40,40,70000),(41,41,100000),(42,42,105000),(43,43,111000),(44,44,118000),(45,45,126000),(46,46,135000),(47,47,145000),(48,48,156000),(49,49,168000),(50,50,181000),(51,51,195000),(52,52,210000),(53,53,226000),(54,54,243000),(55,55,261000),(56,56,280000),(57,57,300000),(58,58,321000),(59,59,343000),(60,60,366000),(61,61,450000),(62,62,460000),(63,63,472000),(64,64,485000),(65,65,499000),(66,66,514000),(67,67,530000),(68,68,547000),(69,69,565000),(70,70,584000),(71,71,604000),(72,72,625000),(73,73,647000),(74,74,670000),(75,75,694000),(76,76,719000),(77,77,745000),(78,78,772000),(79,79,800000),(80,80,900000),(81,81,920000),(82,82,930000),(83,83,969000),(84,84,998000),(85,85,1030000),(86,86,1065000),(87,87,1103000),(88,88,1144000),(89,89,1188000),(90,90,1235000),(91,91,1285000),(92,92,1338000),(93,93,1394000),(94,94,1453000),(95,95,1600000),(96,96,1660000),(97,97,1730000),(98,98,1810000),(99,99,1900000),(100,100,2000000);
/*!40000 ALTER TABLE `role_level_exp` ENABLE KEYS */;
DROP TABLE IF EXISTS `role_realm_class`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
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
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='角色境界阶级表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `role_realm_class` DISABLE KEYS */;
INSERT INTO `role_realm_class` VALUES (1,'炼气',1,0,205,7,0,0,0,0,0,0,0,0,0,0,0,0,'初始阶段，进步容易'),(2,'金刚',2,20,206,14,5000,50,0,0,0,0,0,0,0,0,0,10,'金刚境界，武功境界可提升至40级，身体强壮，可抵御一定外力伤害'),(3,'道玄',3,40,207,28,10000,100,2,0,30,0,0,0,0,0,0,10,'道玄境界，武功境界可提升至60级，劲气外放，操控自如'),(4,'万象',4,60,208,56,15000,150,2,0,30,100,100,100,100,100,100,10,'万象境界，武功境界可提升至80级，神念可感应天地万物'),(5,'天人',5,80,208,112,25000,200,4,4,30,100,100,100,100,100,100,20,'天人境界，武功境界可提升至100级，精气神达到完美的极致境界');
/*!40000 ALTER TABLE `role_realm_class` ENABLE KEYS */;
DROP TABLE IF EXISTS `role_realm_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
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
) ENGINE=InnoDB AUTO_INCREMENT=81 DEFAULT CHARSET=utf8mb4 COMMENT='角色境界等级表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `role_realm_level` DISABLE KEYS */;
INSERT INTO `role_realm_level` VALUES (1,1,1500,0,1,100,20,20,20,20),(2,2,1500,0,1,200,40,40,40,40),(3,3,1500,0,1,300,60,60,60,60),(4,4,1500,0,1,500,80,80,80,80),(5,5,3000,0,1,500,100,100,100,100),(6,6,3000,0,1,600,120,120,120,120),(7,7,3000,0,1,700,140,140,140,140),(8,8,3000,0,1,800,160,160,160,160),(9,9,3000,0,1,900,180,180,180,180),(10,10,4500,0,1,1000,200,200,200,200),(11,11,4500,0,1,1100,220,220,220,220),(12,12,4500,0,1,1200,240,240,240,240),(13,13,4500,0,1,1300,260,260,260,260),(14,14,4500,0,1,1400,280,280,280,280),(15,15,4500,0,1,1500,300,300,300,300),(16,16,4500,0,1,1600,320,320,320,320),(17,17,4500,0,1,1700,340,340,340,340),(18,18,4500,0,1,1800,360,360,360,360),(19,19,4500,0,1,1900,380,380,380,380),(20,20,9000,0,2,2000,400,400,400,400),(21,21,9000,0,2,2100,420,420,420,420),(22,22,9000,0,2,2200,440,440,440,440),(23,23,9000,0,2,2300,460,460,460,460),(24,24,9000,0,2,2400,480,480,480,480),(25,25,9000,0,2,2500,500,500,500,500),(26,26,9000,0,2,2600,520,520,520,520),(27,27,9000,0,2,2700,540,540,540,540),(28,28,9000,0,2,2800,560,560,560,560),(29,29,9000,0,2,2900,580,580,580,580),(30,30,13500,0,3,3000,600,600,600,600),(31,31,13500,0,3,3100,620,620,620,620),(32,32,13500,0,3,3200,640,640,640,640),(33,33,13500,0,3,3300,660,660,660,660),(34,34,13500,0,3,3400,680,680,680,680),(35,35,13500,0,3,3500,700,700,700,700),(36,36,13500,0,3,3600,720,720,720,720),(37,37,13500,0,3,3700,740,740,740,740),(38,38,13500,0,3,3800,760,760,760,760),(39,39,13500,0,3,3900,780,780,780,780),(40,40,18000,2,4,4000,800,800,800,800),(41,41,18000,2,4,4100,820,820,820,820),(42,42,18000,2,4,4200,840,840,840,840),(43,43,18000,2,4,4300,860,860,860,860),(44,44,18000,2,4,4400,880,880,880,880),(45,45,18000,2,4,4500,900,900,900,900),(46,46,18000,2,4,4600,920,920,920,920),(47,47,18000,2,4,4700,940,940,940,940),(48,48,18000,2,4,4800,960,960,960,960),(49,49,18000,2,4,4900,980,980,980,980),(50,50,22500,2,5,5000,1000,1000,1000,1000),(51,51,22500,2,5,5100,1020,1020,1020,1020),(52,52,22500,2,5,5200,1040,1040,1040,1040),(53,53,22500,2,5,5300,1060,1060,1060,1060),(54,54,22500,2,5,5400,1080,1080,1080,1080),(55,55,22500,2,5,5500,1100,1100,1100,1100),(56,56,22500,2,5,5600,1120,1120,1120,1120),(57,57,22500,2,5,5700,1140,1140,1140,1140),(58,58,22500,2,5,5800,1160,1160,1160,1160),(59,59,2250,2,5,5900,1180,1180,1180,1180),(60,60,27000,3,6,6000,1200,1200,1200,1200),(61,61,27000,3,6,6100,1220,1220,1220,1220),(62,62,27000,3,6,6200,1240,1240,1240,1240),(63,63,27000,3,6,6300,1260,1260,1260,1260),(64,64,27000,3,6,6400,1280,1280,1280,1280),(65,65,27000,3,6,6500,1300,1300,1300,1300),(66,66,27000,3,6,6600,1320,1320,1320,1320),(67,67,27000,3,6,6700,1340,1340,1340,1340),(68,68,27000,3,6,6800,1360,1360,1360,1360),(69,69,27000,3,6,6900,1380,1380,1380,1380),(70,70,31500,3,7,7000,1400,1400,1400,1400),(71,71,31500,3,7,7100,1420,1420,1420,1420),(72,72,31500,3,7,7200,1440,1440,1440,1440),(73,73,31500,3,7,7300,1460,1460,1460,1460),(74,74,31500,3,7,7400,1480,1480,1480,1480),(75,75,31500,3,7,7500,1500,1500,1500,1500),(76,76,31500,3,7,7600,1520,1520,1520,1520),(77,77,31500,3,7,7700,1540,1540,1540,1540),(78,78,31500,3,7,7800,1560,1560,1560,1560),(79,79,31500,3,7,7900,1580,1580,1580,1580),(80,80,36000,4,8,8000,1600,1600,1600,1600);
/*!40000 ALTER TABLE `role_realm_level` ENABLE KEYS */;
DROP TABLE IF EXISTS `skill`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `skill` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `name` varchar(10) DEFAULT NULL COMMENT '绝招名称',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型',
  `child_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '子类型',
  `sign` varchar(30) DEFAULT NULL COMMENT '资源标识',
  `music_sign` varchar(30) DEFAULT NULL COMMENT '音乐资源标识',
  `role_id` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '角色ID',
  `required_level` int(11) NOT NULL DEFAULT '0' COMMENT '境界等级',
  `info` varchar(50) DEFAULT NULL COMMENT '绝招描述',
  `jump_attack` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否跳跃攻击',
  `display_param` int(11) NOT NULL DEFAULT '0' COMMENT '显示参数',
  `config` text COMMENT '绝招配置',
  `quality` tinyint(4) NOT NULL DEFAULT '0' COMMENT '技能品质',
  `can_add_level` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否可升级绝招(招式使用)',
  `parent_skill_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '父绝招id(招式使用)',
  `skill_level` smallint(6) NOT NULL DEFAULT '0' COMMENT '绝招等级(招式使用)',
  `order` bigint(20) NOT NULL DEFAULT '0' COMMENT '排序字段',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COMMENT='绝招表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `skill` DISABLE KEYS */;
INSERT INTO `skill` VALUES (1,'闪击',1,1,'ShanJi','ShanJi',-2,1,'单体攻击，获得2点精气',1,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 50,\"Cul2AtkRate\": 10,\"DecPower\": 0,\"IncPower\": 2,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,1,1,1,0),(2,'裂地斩',1,1,'LieDiZhan','LieDiZhan',-2,1,'横排攻击，对多个敌人产生伤害',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 100,\"Cul2AtkRate\": 60,\"DecPower\": 2,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,2,1,0),(3,'甲溃',1,1,'JiaKui','JiaKui',-2,5,'单体攻击，破坏敌方护甲',1,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 200,\"Cul2AtkRate\": 60,\"DecPower\": 2,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 50,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,3,1,0),(4,'必杀技',1,1,'BiShaJi','BiShaJi',-2,2,'单体攻击，消耗精气产生大量伤害',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1000,\"Cul2AtkRate\": 60,\"DecPower\": 6,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 50,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,1,4,1,0),(5,'刀芒',1,1,'DaoMang','DaoMang',3,0,'纵向攻击，产生大量伤害',0,0,'{\"TargetMode\": 3,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 50,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,1,5,1,0),(6,'刀盾',1,3,'DaoDun','DaoDun',3,0,'防御1回合，自身减免40%的伤害，并吸引攻击',0,0,'{\"TargetMode\": 4,\"AttackMode\": 1,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [{\"Type\": 18, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 40, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0},{\"Type\": 19, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,1,6,1,0),(7,'风咒',1,1,'FengZhou','FengZhou',4,0,'单体攻击',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 50,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,1,7,1,0),(8,'治愈之风',1,4,'ZhiYuZhiFeng','ZhiYuZhiFeng',4,0,'单体治疗，优先治疗受伤最重的伙伴',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 500, \"RawValueRate\": 0, \"AttackRate\": 33, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 100, \"ValueCountRate\": 0, \"TargetMode\": 2}]}',0,1,8,1,0),(9,'咆哮利爪',5,1,'PaoXiaoLiZhuaM','PaoXiaoLiZhua',-1,0,'全体',0,0,'{\"TargetMode\": 1,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,30),(10,'凶猛撕咬',5,1,'XiongMengSiYaoM','XiongMengSiYao',-1,0,'单体',1,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0),(11,'冰烈',5,1,'BingLieM','BingLie',-1,0,'单体',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0),(12,'冰烈横向',5,1,'BingLieHengXiangM','BingLie',-1,0,'横向',0,0,'{\"TargetMode\": 2,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 5,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,10),(13,'冰烈纵向',5,1,'BingLieZongXiangM','BingLie',-1,0,'纵向',0,0,'{\"TargetMode\": 3,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,20),(14,'冰烈全体',5,1,'BingLieQuanTiM','BingLie',-1,0,'全体',0,0,'{\"TargetMode\": 1,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,30),(15,'火烈',5,1,'HuoLieM','HuoLie',-1,0,'单体',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0),(16,'火烈横向',5,1,'HuoLieHengXiangM','HuoLie',-1,0,'横向',0,0,'{\"TargetMode\": 2,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,10),(17,'火烈纵向',5,1,'HuoLieZongXiangM','HuoLie',-1,0,'纵向',0,0,'{\"TargetMode\": 3,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,20),(18,'火烈全体',5,1,'HuoLieQuanTiM','HuoLie',-1,0,'全体',0,0,'{\"TargetMode\": 1,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,30),(19,'风烈',5,1,'FengLieM','FengLie',-1,0,'单体',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0),(20,'风烈横向',5,1,'FengLieHengXiangM','FengLie',-1,0,'横向',0,0,'{\"TargetMode\": 2,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,10),(21,'风烈纵向',5,1,'FengLieZongXiangM','FengLie',-1,0,'纵向',0,0,'{\"TargetMode\": 3,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,20),(22,'风烈全体',5,1,'FengLieQuanTiM','FengLie',-1,0,'全体',0,0,'{\"TargetMode\": 1,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,30),(23,'雷烈',5,1,'LeiLieM','LeiLie',-1,0,'单体',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0),(24,'雷烈横向',5,1,'LeiLieHengXiangM','LeiLie',-1,0,'横向',0,0,'{\"TargetMode\": 2,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,10),(25,'雷烈纵向',5,1,'LeiLieZongXiangM','LeiLie',-1,0,'纵向',0,0,'{\"TargetMode\": 3,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,20),(26,'雷烈全体',5,1,'LeiLieQuanTiM','LeiLie',-1,0,'全体',0,0,'{\"TargetMode\": 1,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,30),(27,'土烈',5,1,'TuLieM','TuLie',-1,0,'单体',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0),(28,'土烈横向',5,1,'TuLieHengXiangM','TuLie',-1,0,'横向',0,0,'{\"TargetMode\": 2,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,10),(29,'土烈纵向',5,1,'TuLieZongXiangM','TuLie',-1,0,'纵向',0,0,'{\"TargetMode\": 3,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,20),(30,'土烈全体',5,1,'TuLieQuanTiM','TuLie',-1,0,'全体',0,0,'{\"TargetMode\": 1,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,30),(31,'毒烈',5,1,'DuLieM','DuLie',-1,0,'单体',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 6, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 0, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 30, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',0,0,0,0,0),(32,'毒烈横向',5,1,'DuLieHengXiangM','DuLie',-1,0,'横向',0,0,'{\"TargetMode\": 2,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 6, \"Keep\": 3, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 0, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 30, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',0,0,0,0,10),(33,'毒烈纵向',5,1,'DuLieZongXiangM','DuLie',-1,0,'纵向',0,0,'{\"TargetMode\": 3,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 6, \"Keep\": 3, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 0, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 30, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',0,0,0,0,20),(34,'毒烈全体',5,1,'DuLieQuanTiM','DuLie',-1,0,'全体',0,0,'{\"TargetMode\": 1,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 6, \"Keep\": 3, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 0, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 30, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',0,0,0,0,30),(35,'多连斩',5,1,'DuoLianZhanM','DuoLianZhan',-1,0,'单体攻击',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0),(36,'力劈华山',5,1,'LiPiHuaShanM','LiPiHuaShan',-1,0,'单体攻击，对敌人造成猛烈的打击',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0),(37,'白莲横江',5,1,'BaiLianHengJiangM','LieDiZhan',-1,0,'横向',0,0,'{\"TargetMode\": 2,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,10),(38,'横扫千军',5,1,'HengSaoQianJunM','HengSaoQianJun',-1,0,'横向',0,0,'{\"TargetMode\": 2,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,10),(39,'乾坤刀气',5,1,'QianKunDaoQiM','QianKunDaoQi',-1,0,'纵向',0,0,'{\"TargetMode\": 3,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,20),(40,'三千洛水剑',5,1,'SanQianLuoShuiJianM','SanQianLuoShuiJian',-1,0,'全体',0,0,'{\"TargetMode\": 1,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,30),(41,'死亡标记',5,1,'SiWangBiaoJiM','SiWangBiaoJi',-1,0,'单体',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0),(42,'万箭穿心',5,1,'WanJianChuanXinM','WanJianChuanXin',-1,0,'单体',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0),(43,'狮吼功',5,1,'ShiHouGongM','ShiHouGong',-1,0,'单体',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0),(44,'野蛮冲撞',5,1,'YeManChongZhuangM','YeManChongZhuang',-1,0,'纵向',0,0,'{\"TargetMode\": 3,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,20),(45,'如殂随行',5,1,'RuCuSuiXingM','RuCuSuiXing',-1,0,'全体',0,0,'{\"TargetMode\": 1,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,30),(46,'驱散',5,5,'QuSanM','QuSan',-1,0,'自身',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 7, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 0, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,0,0,0,5),(49,'青竹咒',1,1,'QinZhuZhou','QinZhuZhou',5,0,'单体攻击',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 10,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,1,49,1,0),(50,'雨润',1,5,'YuRun','YuRun',5,0,'全体辅助，增加攻击500，持续2回合',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 2, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,1,50,1,0),(51,'墨画影狼',1,1,'MoHuaYingLang','MoHuaYingLang',6,0,'单体攻击',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 500,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,1,51,1,0),(52,'墨画巫雀',1,5,'MoHuaWuQue','MoHuaWuQue',6,0,'单体攻击，同时每次降低敌方600防御，持续1回合',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 500,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 1, \"BaseValue\": 600, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',0,1,52,1,0),(53,'金刚刀芒',1,1,'DaoMang','DaoMang',3,20,'纵向攻击，产生大量伤害',0,0,'{\"TargetMode\": 3,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,1,5,2,0),(54,'道玄刀芒',1,1,'DaoMang','DaoMang',3,40,'纵向攻击，产生大量伤害',0,0,'{\"TargetMode\": 3,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,5,3,0),(55,'万象刀芒',1,1,'DaoMang','DaoMang',3,60,'纵向攻击，产生大量伤害',0,0,'{\"TargetMode\": 3,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 3000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,5,4,0),(56,'天人刀芒',1,1,'DaoMang','DaoMang',3,80,'纵向攻击，产生大量伤害',0,0,'{\"TargetMode\": 3,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 4000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,5,5,0),(57,'金刚刀盾',1,3,'DaoDun','DaoDun',3,20,'防御1回合，自身减免50%的伤害，并吸引攻击。队伍防御增加500，持续1回合',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [{\"Type\": 18, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 50, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0},{\"Type\": 19, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,1,6,2,0),(58,'道玄刀盾',1,3,'DaoDun','DaoDun',3,40,'防御1回合，自身减免60%的伤害，并吸引攻击。队伍防御增加1000，持续1回合',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [{\"Type\": 18, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 60, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0},{\"Type\": 19, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',1,1,6,3,0),(59,'万象刀盾',1,3,'DaoDun','DaoDun',3,60,'防御1回合，自身减免70%的伤害，并吸引攻击。队伍防御增加2000，持续1回合',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [{\"Type\": 18, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 70, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0},{\"Type\": 19, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 2000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',1,1,6,4,0),(60,'天人刀盾',1,3,'DaoDun','DaoDun',3,80,'防御1回合，自身减免80%的伤害，并吸引攻击。队伍防御增加3000，持续1回合',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [{\"Type\": 18, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 80, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0},{\"Type\": 19, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 3000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',1,1,6,5,0),(61,'有限无敌',1,3,'YouXianWuDi','YouXianWuDi',3,30,'防御1回合，自身减免99%的伤害，并吸引攻击，持续1回合',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [{\"Type\": 18, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 99, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0},{\"Type\": 19, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,0,61,1,0),(62,'金刚大风咒',1,1,'DaFengZhou','FengZhou',4,20,'横排攻击',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,7,2,0),(63,'道玄大风咒',1,1,'DaFengZhou','FengZhou',4,40,'横排攻击',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,7,3,0),(64,'万象大风咒',1,1,'DaFengZhou','FengZhou',4,60,'横排攻击',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 3000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,7,4,0),(65,'天人大风咒',1,1,'DaFengZhou','FengZhou',4,80,'横排攻击',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 4000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,7,5,0),(66,'金刚治愈之风',1,4,'ZhiYuZhiFeng','ZhiYuZhiFeng',4,20,'单体治疗，优先治疗受伤最重的伙伴',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 2000, \"RawValueRate\": 0, \"AttackRate\": 30, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 100, \"ValueCountRate\": 0, \"TargetMode\": 2}]}',1,1,8,2,0),(67,'道玄治愈之风',1,4,'ZhiYuZhiFeng','ZhiYuZhiFeng',4,40,'单体治疗，优先治疗受伤最重的伙伴',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 4000, \"RawValueRate\": 0, \"AttackRate\": 30, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 100, \"ValueCountRate\": 0, \"TargetMode\": 2}]}',1,1,8,3,0),(68,'万象治愈之风',1,4,'ZhiYuZhiFeng','ZhiYuZhiFeng',4,60,'单体治疗，优先治疗受伤最重的伙伴',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 0, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 6000, \"RawValueRate\": 0, \"AttackRate\": 30, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 100, \"ValueCountRate\": 0, \"TargetMode\": 2}]}',1,1,8,4,0),(69,'天人治愈之风',1,4,'ZhiYuZhiFeng','ZhiYuZhiFeng',4,80,'单体治疗，优先治疗受伤最重的伙伴',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 8000, \"RawValueRate\": 0, \"AttackRate\": 30, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 100, \"ValueCountRate\": 0, \"TargetMode\": 2}]}',0,1,8,5,0),(70,'承天载物',1,4,'ChengTianZaiWu','ChengTianZaiWu',4,30,'全体治疗，产生大量治疗量',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 500, \"RawValueRate\": 0, \"AttackRate\": 30, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',2,0,70,1,0),(71,'金刚雨润',1,5,'YuRun','YuRun',5,20,'全体辅助，增加攻击1000、防御500，持续2回合',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 2, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 3, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',1,1,50,2,0),(72,'道玄雨润',1,5,'YuRun','YuRun',5,40,'全体辅助，增加攻击1500、防御1000，持续2回合',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 2, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 3, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',1,1,50,3,0),(73,'万象雨润',1,5,'YuRun','YuRun',5,60,'全体辅助，增加攻击2000、防御1500，持续2回合',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 2, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 2000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 3, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',1,1,50,4,0),(74,'天人雨润',1,5,'YuRun','YuRun',5,80,'全体辅助，增加攻击3000、防御2000，持续2回合',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 2, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 3000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 3, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 2000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',1,1,50,5,0),(75,'金刚青竹咒',1,1,'QinZhuZhou','QinZhuZhou',5,20,'单体攻击，附毒500，持续2回合',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 500,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 6, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',1,1,49,2,0),(76,'道玄青竹咒',1,1,'QinZhuZhou','QinZhuZhou',5,40,'单体攻击，附毒1000，持续2回合',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 6, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',1,1,49,3,0),(77,'万象青竹咒',1,1,'QinZhuZhou','QinZhuZhou',5,60,'单体攻击，附毒1500，持续2回合',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1500,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 6, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',1,1,49,4,0),(78,'天人青竹咒',1,1,'QinZhuZhou','QinZhuZhou',5,80,'单体攻击，附毒2000，持续2回合',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 6, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 2000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',1,1,49,5,0),(79,'金刚墨画影狼',1,1,'MoHuaYingLang','MoHuaYingLang',6,20,'横排穿透',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 500,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,51,2,0),(80,'道玄墨画影狼',1,1,'MoHuaYingLang','MoHuaYingLang',6,40,'横排穿透',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,51,3,0),(81,'万象墨画影狼',1,1,'MoHuaYingLang','MoHuaYingLang',6,60,'横排穿透',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1500,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,51,4,0),(82,'天人墨画影狼',1,1,'MoHuaYingLang','MoHuaYingLang',6,80,'横排穿透',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,51,5,0),(83,'金刚墨画巫雀',1,5,'MoHuaWuQue','MoHuaWuQue',6,20,'单体攻击，同时每次降低敌方800防御，持续1回合',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 500,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 1, \"BaseValue\": 800, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',1,1,52,2,0),(84,'道玄墨画巫雀',1,5,'MoHuaWuQue','MoHuaWuQue',6,40,'单体攻击，同时每次降低敌方800防御，持续1回合',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 1, \"BaseValue\": 1000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',1,1,52,3,0),(85,'万象墨画巫雀',1,5,'MoHuaWuQue','MoHuaWuQue',6,60,'单体攻击，同时每次降低敌方1200防御，持续1回合',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1200,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 1, \"BaseValue\": 1200, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',1,1,52,4,0),(86,'天人墨画巫雀',1,5,'MoHuaWuQue','MoHuaWuQue',6,80,'单体攻击，同时每次降低敌方1400防御，持续1回合',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1200,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 1, \"BaseValue\": 1200, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',1,1,52,5,0),(87,'神魔封禁',1,5,'ShenMoFengJin','ShenMoFengJin',6,30,'全体攻击，并80%几率禁魔1回合',0,0,'{\"TargetMode\": 1,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 600,\"Cul2AtkRate\": 30,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 24, \"Keep\": 1, \"Override\": 1, \"Rate\": 80, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',2,0,87,1,0),(88,'圣白莲',1,5,'ShengBaiLian','ShengBaiLian',5,30,'解除负面状态，恢复生命',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 7, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 0, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": , \"RawValueRate\": 100, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',2,0,88,1,0),(89,'断岳',1,1,'DuanYue','LieDiZhan',-2,10,'纵向攻击',0,0,'{\"TargetMode\": 3,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 60,\"Cul2AtkRate\": 60,\"DecPower\": 2,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,0,89,1,0),(90,'金刚闪击',1,1,'ShanJi','ShanJi',-2,20,'单体攻击，获得3点精气',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1000,\"Cul2AtkRate\": 10,\"DecPower\": 0,\"IncPower\": 3,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,1,2,0),(91,'金刚必杀技',1,1,'BiShaJi','BiShaJi',-2,20,'单体攻击，消耗精气产生大量伤害',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1000,\"Cul2AtkRate\": 60,\"DecPower\": 6,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 50,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,4,2,0),(92,'金刚庇护',1,3,'JinGangBiHu','JinGangBiHu',-2,30,'防御1回合，自身减免60%的伤害，并吸引攻击。同时提高全体防御3000，持续1回合',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [{\"Type\": 18, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 60, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 60, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0},{\"Type\": 19, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 3000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',1,0,92,1,0),(93,'道玄闪击',1,1,'ShanJi','ShanJi',-2,40,'单体攻击，获得4点精气',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1000,\"Cul2AtkRate\": 10,\"DecPower\": 0,\"IncPower\": 4,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 5,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,1,3,0),(94,'道玄必杀技',1,1,'BiShaJi','BiShaJi',-2,40,'单体攻击，消耗精气产生大量伤害',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1000,\"Cul2AtkRate\": 60,\"DecPower\": 6,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 50,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,4,3,0),(95,'气疗术',1,4,'QiLiaoShu','ZhiYuZhiFeng',-2,50,'单体治疗，治疗损失的生命最多的成员',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 4000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 2}]}',2,0,0,1,0),(96,'万象闪击',1,1,'ShanJi','ShanJi',-2,60,'单体攻击，获得5点精气',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 3000,\"Cul2AtkRate\": 10,\"DecPower\": 0,\"IncPower\": 5,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,1,4,0),(97,'万象必杀技',1,1,'BiShaJi','BiShaJi',-2,60,'单体攻击，消耗精气产生大量伤害',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1000,\"Cul2AtkRate\": 60,\"DecPower\": 6,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 50,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,4,4,0),(98,'天剑',1,1,'TianJian','TianJian',-2,70,'全体攻击，一剑开天门',0,0,'{\"TargetMode\": 1,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1000,\"Cul2AtkRate\": 60,\"DecPower\": 4,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',2,0,0,1,0),(99,'天人闪击',1,1,'ShanJi','ShanJi',-2,80,'单体攻击，获得6点精气',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 4000,\"Cul2AtkRate\": 10,\"DecPower\": 0,\"IncPower\": 6,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,1,5,0),(100,'天人必杀技',1,1,'BiShaJi','BiShaJi',-2,80,'单体攻击，消耗精气产生大量伤害',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1000,\"Cul2AtkRate\": 60,\"DecPower\": 6,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 50,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,4,5,0),(108,'治疗',5,4,'ZhiLiaoM','ZhiYuZhiFeng',-1,0,'治疗生命最少的队友',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 0, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 100, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 2}]}',0,0,0,0,40),(109,'增益',5,5,'ZengYiM','ZengYi',-1,0,'自身',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,5),(998,'破甲眩晕',0,0,'PoJiaXuanYun',NULL,-1,0,NULL,0,0,NULL,0,0,0,0,0),(999,'t',7,1,NULL,NULL,1,1,'desc',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0);
/*!40000 ALTER TABLE `skill` ENABLE KEYS */;
DROP TABLE IF EXISTS `skill_content`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `skill_content` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `skill_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '绝招ID',
  `release_num` int(11) NOT NULL DEFAULT '0' COMMENT '释放次数',
  `recover_round_num` int(11) NOT NULL DEFAULT '0' COMMENT '恢复回合数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8mb4 COMMENT='绝招数据表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `skill_content` DISABLE KEYS */;
INSERT INTO `skill_content` VALUES (1,5,2,1),(2,6,1,0),(3,7,2,1),(4,8,3,2),(5,49,2,1),(6,50,1,2),(7,51,2,1),(8,52,1,2),(9,53,2,1),(10,54,2,1),(11,55,2,1),(12,56,2,1),(13,57,1,0),(14,58,1,0),(15,59,1,0),(16,60,1,0),(17,61,1,4),(18,62,2,1),(19,63,2,1),(20,64,2,1),(21,65,2,1),(22,66,3,2),(23,67,3,2),(24,68,3,2),(25,69,3,2),(26,70,1,3),(27,71,1,2),(28,72,1,2),(29,73,1,2),(30,74,1,2),(31,75,2,1),(32,76,2,1),(33,77,2,1),(34,78,2,1),(35,79,3,1),(36,80,3,1),(37,81,3,1),(38,82,3,1),(39,83,1,2),(40,84,1,2),(41,85,1,2),(42,86,1,2),(43,87,1,2),(44,88,1,3);
/*!40000 ALTER TABLE `skill_content` ENABLE KEYS */;
DROP TABLE IF EXISTS `sword_soul`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sword_soul` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '剑心ID',
  `type_id` int(11) NOT NULL COMMENT '类型ID',
  `name` varchar(10) NOT NULL COMMENT '剑心名称',
  `sign` varchar(30) NOT NULL DEFAULT '' COMMENT '资源标识',
  `desc` varchar(20) NOT NULL COMMENT '剑心描述',
  `quality` tinyint(4) NOT NULL COMMENT '品质',
  `fragment_num` smallint(6) DEFAULT NULL COMMENT '碎片数量',
  `only_exchange` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否只能兑换获得',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COMMENT='剑心';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `sword_soul` DISABLE KEYS */;
INSERT INTO `sword_soul` VALUES (1,12,'斗铠','DouKai','',5,0,0),(2,1,'白虹','BaiHong','',2,0,0),(3,1,'断水','DuanShui','',3,0,0),(4,1,'龙泉','LongQuan','',4,0,0),(5,1,'巨阙','JuQue','',5,20,0),(6,2,'永用','YongYong','',2,0,0),(7,2,'青干','QingGan','',3,0,0),(8,2,'纯阳','ChunYang','',4,0,0),(9,2,'湛泸','ZhanLu','',5,20,0),(10,3,'惊鲵','JingNi','',4,0,0),(11,3,'赤霄','ChiXiao','',5,20,0),(12,5,'转魄','ZhuanPo','',3,0,0),(13,5,'工布','GongBu','',4,0,0),(14,5,'龙渊','LongYuan','',5,20,0),(15,4,'悬翦','XuanJian','',3,0,0),(16,4,'含光','HanGuang','',4,0,0),(17,4,'承影','ChengYing','',5,20,0),(18,7,'燕支','YanZhi','',3,0,0),(19,7,'灭魂','MieHun','',4,0,0),(20,7,'干将','GanJiang','',5,20,0),(21,11,'紫电','ZiDian','',3,0,0),(22,11,'宵练','XiaoLian','',4,0,0),(23,11,'莫邪','MoXie','',5,20,0),(24,8,'神龟','ShenGui','',3,0,0),(25,8,'北斗','BeiDou','',4,0,0),(26,8,'星河','XingHe','',5,20,0),(27,6,'百步','BaiBu','',3,0,0),(28,6,'鱼肠','YuChang','',4,0,0),(29,6,'春秋','ChunQiu','',5,20,0),(30,10,'磐石','PanShi','',3,0,0),(31,10,'无尘','WuChen','',4,0,0),(32,10,'青梅','QingMei','',5,20,0),(33,9,'却邪','QueXie','',3,0,0),(34,9,'真刚','ZhenGang','',4,0,0),(35,9,'竹马','ZhuMa','',5,20,0),(36,3,'轩辕','XuanYuan','装备后改角色其他剑心等级提升1级',6,0,1),(37,13,'潜龙','QianLong','',1,0,0);
/*!40000 ALTER TABLE `sword_soul` ENABLE KEYS */;
DROP TABLE IF EXISTS `sword_soul_exchange`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='剑心兑换信息';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `sword_soul_exchange` DISABLE KEYS */;
INSERT INTO `sword_soul_exchange` VALUES (1,36,5,3,1,17,3,1,9,3,1,11,3,1,14,3,1);
/*!40000 ALTER TABLE `sword_soul_exchange` ENABLE KEYS */;
DROP TABLE IF EXISTS `sword_soul_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sword_soul_level` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `sword_soul_id` smallint(6) NOT NULL COMMENT '剑心ID',
  `level` tinyint(4) NOT NULL COMMENT '等级',
  `value` int(11) NOT NULL COMMENT '属性加值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=399 DEFAULT CHARSET=utf8mb4 COMMENT='剑心等级';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `sword_soul_level` DISABLE KEYS */;
INSERT INTO `sword_soul_level` VALUES (1,1,1,10),(2,1,2,20),(3,1,3,30),(4,1,4,40),(5,1,5,50),(6,1,6,60),(7,1,7,70),(8,1,8,80),(9,1,9,90),(10,1,10,100),(11,1,11,110),(12,2,1,30),(13,2,2,60),(14,2,3,90),(15,2,4,120),(16,2,5,150),(17,2,6,180),(18,2,7,210),(19,2,8,240),(20,2,9,270),(21,2,10,300),(22,2,11,330),(23,3,1,75),(24,3,2,150),(25,3,3,225),(26,3,4,300),(27,3,5,375),(28,3,6,450),(29,3,7,525),(30,3,8,600),(31,3,9,675),(32,3,10,750),(33,3,11,825),(34,4,1,150),(35,4,2,300),(36,4,3,450),(37,4,4,600),(38,4,5,750),(39,4,6,900),(40,4,7,1050),(41,4,8,1200),(42,4,9,1350),(43,4,10,1500),(44,4,11,1650),(47,5,1,300),(48,5,2,600),(49,5,3,900),(50,5,4,1200),(51,5,5,1500),(52,5,6,1800),(53,5,7,2100),(54,5,8,2400),(55,5,9,2700),(56,5,10,3000),(57,5,11,3300),(58,6,1,25),(59,6,2,50),(60,6,3,75),(61,6,4,100),(62,6,5,125),(63,6,6,150),(64,6,7,175),(65,6,8,200),(66,6,9,225),(67,6,10,250),(68,6,11,275),(69,7,1,50),(70,7,2,100),(71,7,3,150),(72,7,4,200),(73,7,5,250),(74,7,6,300),(75,7,7,350),(76,7,8,400),(77,7,9,450),(78,7,10,500),(79,7,11,550),(80,8,1,100),(81,8,2,200),(82,8,3,300),(83,8,4,400),(84,8,5,500),(85,8,6,600),(86,8,7,700),(87,8,8,800),(88,8,9,900),(89,8,10,1000),(90,8,11,1100),(91,9,1,200),(92,9,2,400),(93,9,3,600),(94,9,4,800),(95,9,5,1000),(96,9,6,1200),(97,9,7,1400),(98,9,8,1600),(99,9,9,1800),(100,9,10,2000),(101,9,11,2200),(102,10,1,500),(103,10,2,1000),(104,10,3,1500),(105,10,4,2000),(106,10,5,2500),(107,10,6,3000),(108,10,7,3500),(109,10,8,4000),(110,10,9,4500),(111,10,10,5000),(112,10,11,5500),(113,11,1,1000),(114,11,2,2000),(115,11,3,3000),(116,11,4,4000),(117,11,5,5000),(118,11,6,6000),(119,11,7,7000),(120,11,8,8000),(121,11,9,9000),(122,11,10,10000),(123,11,11,11000),(124,12,1,50),(125,12,2,100),(126,12,3,150),(127,12,4,200),(128,12,5,250),(129,12,6,300),(130,12,7,350),(131,12,8,400),(132,12,9,450),(133,12,10,500),(134,12,11,550),(135,13,1,100),(136,13,2,200),(137,13,3,300),(138,13,4,400),(139,13,5,500),(140,13,6,600),(141,13,7,700),(142,13,8,800),(143,13,9,900),(144,13,10,1000),(145,13,11,1100),(146,14,1,200),(147,14,2,400),(148,14,3,600),(149,14,4,800),(150,14,5,1000),(151,14,6,1200),(152,14,7,1400),(153,14,8,1600),(154,14,9,1800),(155,14,10,2000),(156,14,11,2200),(157,15,1,50),(158,15,2,100),(159,15,3,150),(160,15,4,200),(161,15,5,250),(162,15,6,300),(163,15,7,350),(164,15,8,400),(165,15,9,450),(166,15,10,500),(167,15,11,550),(168,16,1,100),(169,16,2,200),(170,16,3,300),(171,16,4,400),(172,16,5,500),(173,16,6,600),(174,16,7,700),(175,16,8,800),(176,16,9,900),(177,16,10,1000),(178,16,11,1100),(179,17,1,200),(180,17,2,400),(181,17,3,600),(182,17,4,800),(183,17,5,1000),(184,17,6,1200),(185,17,7,1400),(186,17,8,1600),(187,17,9,1800),(188,17,10,2000),(189,17,11,2200),(190,18,1,5),(191,18,2,10),(192,18,3,15),(193,18,4,20),(194,18,5,25),(195,18,6,30),(196,18,7,35),(197,18,8,40),(198,18,9,45),(199,18,10,50),(200,18,11,55),(201,19,1,10),(202,19,2,20),(203,19,3,30),(204,19,4,40),(205,19,5,50),(206,19,6,60),(207,19,7,70),(208,19,8,80),(209,19,9,90),(210,19,10,100),(211,19,11,110),(212,20,1,20),(213,20,2,40),(214,20,3,60),(215,20,4,80),(216,20,5,100),(217,20,6,120),(218,20,7,140),(219,20,8,160),(220,20,9,180),(221,20,10,200),(222,20,11,220),(223,21,1,5),(224,21,2,10),(225,21,3,15),(226,21,4,20),(227,21,5,25),(228,21,6,30),(229,21,7,35),(230,21,8,40),(231,21,9,45),(232,21,10,50),(233,21,11,55),(234,22,1,10),(235,22,2,20),(236,22,3,30),(237,22,4,40),(238,22,5,50),(239,22,6,60),(240,22,7,70),(241,22,8,80),(242,22,9,90),(243,22,10,100),(244,22,11,110),(245,23,1,20),(246,23,2,40),(247,23,3,60),(248,23,4,80),(249,23,5,100),(250,23,6,120),(251,23,7,140),(252,23,8,160),(253,23,9,180),(254,23,10,200),(255,23,11,220),(256,24,1,5),(257,24,2,10),(258,24,3,15),(259,24,4,20),(260,24,5,25),(261,24,6,30),(262,24,7,35),(263,24,8,40),(264,24,9,45),(265,24,10,50),(266,24,11,55),(267,25,1,10),(268,25,2,20),(269,25,3,30),(270,25,4,40),(271,25,5,50),(272,25,6,60),(273,25,7,70),(274,25,8,80),(275,25,9,90),(276,25,10,100),(277,25,11,110),(278,26,1,20),(279,26,2,40),(280,26,3,60),(281,26,4,80),(282,26,5,100),(283,26,6,120),(284,26,7,140),(285,26,8,160),(286,26,9,180),(287,26,10,200),(288,26,11,220),(289,27,1,5),(290,27,2,10),(291,27,3,15),(292,27,4,20),(293,27,5,25),(294,27,6,30),(295,27,7,35),(296,27,8,40),(297,27,9,45),(298,27,10,50),(299,27,11,55),(300,28,1,10),(301,28,2,20),(302,28,3,30),(303,28,4,40),(304,28,5,50),(305,28,6,60),(306,28,7,70),(307,28,8,80),(308,28,9,90),(309,28,10,100),(310,28,11,110),(311,29,1,20),(312,29,2,40),(313,29,3,60),(314,29,4,80),(315,29,5,100),(316,29,6,120),(317,29,7,140),(318,29,8,160),(319,29,9,180),(320,29,10,200),(321,29,11,220),(322,30,1,5),(323,30,2,10),(324,30,3,15),(325,30,4,20),(326,30,5,25),(327,30,6,30),(328,30,7,35),(329,30,8,40),(330,30,9,45),(331,30,10,50),(332,30,11,55),(333,31,1,10),(334,31,2,20),(335,31,3,30),(336,31,4,40),(337,31,5,50),(338,31,6,60),(339,31,7,70),(340,31,8,80),(341,31,9,90),(342,31,10,100),(343,31,11,110),(344,32,1,20),(345,32,2,40),(346,32,3,60),(347,32,4,80),(348,32,5,100),(349,32,6,120),(350,32,7,140),(351,32,8,160),(352,32,9,180),(353,32,10,200),(354,32,11,220),(355,33,1,5),(356,33,2,10),(357,33,3,15),(358,33,4,20),(359,33,5,25),(360,33,6,30),(361,33,7,35),(362,33,8,40),(363,33,9,45),(364,33,10,50),(365,33,11,55),(366,34,1,10),(367,34,2,20),(368,34,3,30),(369,34,4,40),(370,34,5,50),(371,34,6,60),(372,34,7,70),(373,34,8,80),(374,34,9,90),(375,34,10,100),(376,34,11,110),(377,35,1,20),(378,35,2,40),(379,35,3,60),(380,35,4,80),(381,35,5,100),(382,35,6,120),(383,35,7,140),(384,35,8,160),(385,35,9,180),(386,35,10,200),(387,35,11,220),(388,36,1,2000),(389,36,2,4000),(390,36,3,6000),(391,36,4,8000),(392,36,5,10000),(393,36,6,12000),(394,36,7,14000),(395,36,8,16000),(396,36,9,18000),(397,36,10,20000),(398,37,1,268434455);
/*!40000 ALTER TABLE `sword_soul_level` ENABLE KEYS */;
DROP TABLE IF EXISTS `sword_soul_quality`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sword_soul_quality` (
  `id` smallint(6) NOT NULL COMMENT '剑心等级ID',
  `name` varchar(10) NOT NULL COMMENT '品质名称',
  `sign` varchar(20) NOT NULL COMMENT '程序标示',
  `color` varchar(20) DEFAULT NULL COMMENT '颜色',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='剑心品质';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `sword_soul_quality` DISABLE KEYS */;
INSERT INTO `sword_soul_quality` VALUES (0,'杂物','NULL','0xc5c3b7'),(1,'特殊','SPECIAL','0xdb2e00'),(2,'优良','FINE','0x22ac38'),(3,'精良','EXCELLENT','0x00a0e9'),(4,'传奇','LEGEND','0xc301c3'),(5,'神器','ARTIFACT','0xfff100'),(6,'唯一','ONLY','0xf39700');
/*!40000 ALTER TABLE `sword_soul_quality` ENABLE KEYS */;
DROP TABLE IF EXISTS `sword_soul_quality_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sword_soul_quality_level` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '剑心等级ID',
  `quality_id` smallint(6) NOT NULL COMMENT '品质名称',
  `level` tinyint(4) NOT NULL COMMENT '等级',
  `exp` int(11) NOT NULL COMMENT '升到这一级所需的经验',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=52 DEFAULT CHARSET=utf8mb4 COMMENT='剑心品质';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `sword_soul_quality_level` DISABLE KEYS */;
INSERT INTO `sword_soul_quality_level` VALUES (1,1,1,1000),(2,2,1,50),(3,2,2,125),(4,2,3,250),(5,2,4,500),(6,2,5,1000),(7,2,6,2000),(8,2,7,4000),(9,2,8,5000),(10,2,9,6000),(11,2,10,7000),(12,3,1,200),(13,3,2,250),(14,3,3,500),(15,3,4,1000),(16,3,5,2000),(17,3,6,4000),(18,3,7,5000),(19,3,8,6000),(20,3,9,7000),(21,3,10,8000),(22,4,1,500),(23,4,2,500),(24,4,3,1000),(25,4,4,2000),(26,4,5,4000),(27,4,6,5000),(28,4,7,6000),(29,4,8,7000),(30,4,9,8000),(31,4,10,9000),(32,5,1,1500),(33,5,2,1000),(34,5,3,2000),(35,5,4,4000),(36,5,5,5000),(37,5,6,6000),(38,5,7,7000),(39,5,8,8000),(40,5,9,9000),(41,5,10,10000),(42,6,1,1080),(43,6,2,2000),(44,6,3,4000),(45,6,4,5000),(46,6,5,6000),(47,6,6,7000),(48,6,7,8000),(49,6,8,9000),(50,6,9,10000),(51,6,10,11000);
/*!40000 ALTER TABLE `sword_soul_quality_level` ENABLE KEYS */;
DROP TABLE IF EXISTS `sword_soul_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sword_soul_type` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '类型ID',
  `name` varchar(10) NOT NULL COMMENT '类型名称',
  `sign` varchar(20) DEFAULT NULL COMMENT '程序标示',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COMMENT='剑心类型';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `sword_soul_type` DISABLE KEYS */;
INSERT INTO `sword_soul_type` VALUES (1,'攻击','ATTACK '),(2,'防御','DEFENCE '),(3,'生命','HEALTH '),(4,'速度','SPEED '),(5,'内力','CULTIVATION '),(6,'命中','HIT_LEVEL '),(7,'暴击','CRITICAL_LEVEL'),(8,'格挡','BLOCK_LEVEL '),(9,'破击','DESTROY_LEVEL '),(10,'韧性','TENACITY_LEVEL'),(11,'闪避','DODGE_LEVEL'),(12,'护甲','SUNDER_MAX_VALUE'),(13,'剑心经验','EXP');
/*!40000 ALTER TABLE `sword_soul_type` ENABLE KEYS */;
DROP TABLE IF EXISTS `tower_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tower_level` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `floor` smallint(6) NOT NULL COMMENT '楼层',
  PRIMARY KEY (`id`),
  UNIQUE KEY `floor` (`floor`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='极限关卡通天塔';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `tower_level` DISABLE KEYS */;
INSERT INTO `tower_level` VALUES (1,1);
/*!40000 ALTER TABLE `tower_level` ENABLE KEYS */;
DROP TABLE IF EXISTS `town`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `town` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '城镇ID,-1为集会所',
  `lock` int(11) NOT NULL COMMENT '解锁权值',
  `name` varchar(10) NOT NULL DEFAULT '' COMMENT '城镇名称',
  `sign` varchar(30) NOT NULL DEFAULT '' COMMENT '资源标识',
  `music` varchar(20) NOT NULL COMMENT '音乐资源标识',
  `start_x` int(11) NOT NULL COMMENT '出生点x轴坐标',
  `start_y` int(11) NOT NULL COMMENT '出生点y轴坐标',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='城镇';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `town` DISABLE KEYS */;
INSERT INTO `town` VALUES (1,100000,'神龙岛','ShenLongDao','ZhuTiQu',1589,965),(2,100110,'烛堡','ZhuBao','ZhuTiQu',1369,760),(3,999999,'测试城镇','ShenLongDao','ZhuTiQu',1282,529),(4,100120,'铸剑山庄','ZhuJianShanZhuang','ZhuTiQu',1369,760);
/*!40000 ALTER TABLE `town` ENABLE KEYS */;
DROP TABLE IF EXISTS `town_npc`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `town_npc` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `town_id` smallint(6) NOT NULL COMMENT '城镇ID',
  `name` varchar(10) NOT NULL COMMENT 'NPC名称',
  `sign` varchar(20) NOT NULL COMMENT '资源标识',
  `x` int(11) NOT NULL COMMENT 'x轴坐标',
  `y` int(11) NOT NULL COMMENT 'y轴坐标',
  `direction` varchar(20) DEFAULT NULL COMMENT '朝向',
  `talk` varchar(200) NOT NULL COMMENT '对话',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COMMENT='城镇NPC';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `town_npc` DISABLE KEYS */;
INSERT INTO `town_npc` VALUES (1,1,'龙姬','LongJi',1670,1280,'b','阴影这么多该怎么办！||我感觉体内有股奇怪的力量。'),(2,1,'梦妖','MengYao',1560,1350,'rb','你见过神龙么？||作为龙姬大人的灵宠，日子过的还是挺舒坦的。'),(3,1,'商人','ShangRen',1160,1020,'rb','江湖险恶，常备各种良药才是保命上策。||做点小生意不容易，多来光顾呀！'),(4,1,'梦妖小清','MengYao',1064,345,'lb','我好想离开神龙岛去外面的世界。||如果可以像个大侠一样去冒险……'),(5,1,'梦妖小辉','MengYao',1296,1352,'l','小杰说他是最熟悉神龙岛的人。||你想知道什么？'),(6,1,'梦妖小杰','MengYao',1180,1352,'r','我对神龙岛了如指掌。||传说在火山的内部有一个密室。||我知道一个长生不老的秘密。'),(7,1,'梦妖小立','MengYao',1730,600,'t','我躲在这里谁也找不到我。||玩捉迷藏我最厉害了。||他们找了我两天都没找到我。||今天吃了小苹果，感觉自己萌萌哒~'),(8,1,'梦妖阿剑','MengYao',2030,823,'lb','我也是出过岛的妖！||江湖上可是流传着我的传说。||我在黑夜森林看见了奇怪的黑影。'),(9,1,'梦妖阿樂','MengYao',611,557,'l','宝剑尚未配妥，出门已是江湖。||我有一天梦见了上古的剑灵。');
/*!40000 ALTER TABLE `town_npc` ENABLE KEYS */;
DROP TABLE IF EXISTS `town_npc_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `town_npc_item` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `town_npc_id` int(11) NOT NULL COMMENT '城镇NPC ID',
  `item_id` smallint(6) NOT NULL COMMENT '物品ID',
  `stock` smallint(6) NOT NULL COMMENT '库存',
  `vip` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'vip特供，1表示vip',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='城镇NPC对话';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `town_npc_item` DISABLE KEYS */;
INSERT INTO `town_npc_item` VALUES (1,3,31,9999,0),(2,3,36,9999,0),(3,3,37,9999,0),(4,3,38,9999,0);
/*!40000 ALTER TABLE `town_npc_item` ENABLE KEYS */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;


THESQL1
	);

if ($renew) {
	db_execute($db, <<<THESQL2

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
DROP TABLE IF EXISTS `global_announcement`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `global_announcement` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '公告ID',
  `expire_time` bigint(20) NOT NULL COMMENT '创建时间戳',
  `tpl_id` int(11) NOT NULL COMMENT '公告模版ID',
  `parameters` varchar(1024) NOT NULL COMMENT '模版参数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=78 DEFAULT CHARSET=utf8mb4 COMMENT='公告列表';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `global_arena_rank`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `global_arena_rank` (
  `rank` int(11) NOT NULL COMMENT '排名',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  PRIMARY KEY (`rank`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='全局比武场数据';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player` (
  `id` bigint(20) NOT NULL COMMENT '玩家ID',
  `user` varchar(250) NOT NULL COMMENT '平台传递过来的用户标识',
  `nick` varchar(50) NOT NULL COMMENT '玩家昵称',
  `main_role_id` bigint(20) NOT NULL COMMENT '主角ID',
  PRIMARY KEY (`id`),
  KEY `ix_player_sign` (`user`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家基础信息';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_arena`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_arena` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `daily_num` smallint(6) NOT NULL COMMENT '今日已挑战次数',
  `failed_cd_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '战败CD结束时间',
  `record_read_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '战报阅读时间',
  `win_times` smallint(6) NOT NULL DEFAULT '0' COMMENT '>0 连胜场次; 0 保持不变; -1 下降趋势',
  `daily_award_longbi` int(11) NOT NULL DEFAULT '0' COMMENT '今日获得龙币累计',
  `new_record_count` smallint(6) NOT NULL COMMENT '新战报计数',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家比武场数据';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_arena_rank`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_arena_rank` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `rank` int(11) NOT NULL DEFAULT '0' COMMENT '昨天排名',
  `rank1` int(11) NOT NULL DEFAULT '0' COMMENT '1天前排名',
  `rank2` int(11) NOT NULL DEFAULT '0' COMMENT '2天前排名',
  `rank3` int(11) NOT NULL DEFAULT '0' COMMENT '3天前排名',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家比武场最近排名记录';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_arena_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_arena_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '记录ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `mode` tinyint(4) NOT NULL COMMENT '记录类型，0无数据，1挑战成功，2挑战失败，3被挑战且成功，4被挑战且失败',
  `old_rank` int(11) NOT NULL COMMENT '原排位',
  `new_rank` int(11) NOT NULL COMMENT '新排位',
  `fight_num` int(11) NOT NULL COMMENT '战力',
  `target_pid` bigint(20) NOT NULL COMMENT '对手玩家ID',
  `target_nick` varchar(50) NOT NULL COMMENT '对手昵称',
  `target_old_rank` int(11) NOT NULL COMMENT '对手原排位',
  `target_new_rank` int(11) NOT NULL COMMENT '对手新排位',
  `target_fight_num` int(11) NOT NULL COMMENT '对手战力',
  `record_time` bigint(20) NOT NULL COMMENT '记录时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家比武场记录';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_battle_pet`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_battle_pet` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `battle_pet_id` smallint(6) NOT NULL COMMENT '灵宠ID',
  `ball_num` tinyint(4) NOT NULL COMMENT '已有的灵宠契约球数量',
  `activated` tinyint(4) NOT NULL DEFAULT '0' COMMENT '灵宠是否已激活',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=56 DEFAULT CHARSET=utf8mb4 COMMENT='玩家灵宠数据';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_battle_pet_config`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_battle_pet_config` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `grid1` bigint(20) DEFAULT '-1' COMMENT '灵宠栏位1(-1-未开启)',
  `grid2` bigint(20) DEFAULT '-1' COMMENT '灵宠栏位2(-1-未开启)',
  `grid3` bigint(20) DEFAULT '-1' COMMENT '灵宠栏位3(-1-未开启)',
  `grid4` bigint(20) DEFAULT '-1' COMMENT '灵宠栏位4(-1-未开启)',
  `grid5` bigint(20) DEFAULT '-1' COMMENT '灵宠栏位5(-1-未开启)',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家灵宠配置';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_chest_state`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_chest_state` (
  `pid` bigint(20) NOT NULL COMMENT '玩家id',
  `free_coin_chest_num` int(11) NOT NULL COMMENT '每日免费青铜宝箱数',
  `last_free_coin_chest_at` bigint(20) NOT NULL COMMENT '上次开免费青铜宝箱时间',
  `coin_chest_num` int(11) NOT NULL COMMENT '今天开青铜宝箱次数',
  `coin_chest_ten_num` int(11) NOT NULL COMMENT '今日青铜宝箱十连抽次数',
  `coin_chest_consume` bigint(20) NOT NULL COMMENT '今天开青铜宝箱花费铜钱数',
  `last_coin_chest_at` bigint(20) NOT NULL COMMENT '上次开消费青铜宝箱时间',
  `last_free_ingot_chest_at` bigint(20) NOT NULL COMMENT '上次开免费神龙宝箱时间',
  `ingot_chest_num` int(11) NOT NULL COMMENT '今天开神龙宝箱次数',
  `ingot_chest_ten_num` int(11) NOT NULL COMMENT '今日神龙宝箱十连抽次数',
  `ingot_chest_consume` bigint(20) NOT NULL COMMENT '今天开神龙宝箱花费元宝数',
  `last_ingot_chest_at` bigint(20) NOT NULL COMMENT '上次开消费神龙宝箱时间',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家宝箱状态';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_coins`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_coins` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `buy_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '购买更新时间',
  `daily_count` smallint(6) DEFAULT '0' COMMENT '当天购买次数',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家铜币兑换表';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_equipment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_equipment` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `role_id` tinyint(4) NOT NULL COMMENT '角色ID',
  `weapon_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '武器的player_item表主键ID',
  `clothes_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '战袍的player_item表主键ID',
  `accessories_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '饰品的player_item表主键ID',
  `shoe_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '靴子的player_item表主键ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家装备表';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_fight_num`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_fight_num` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `fight_num` int(11) NOT NULL COMMENT '战力力',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家战斗力';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_formation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_formation` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `pos0` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '站位0',
  `pos1` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '站位1',
  `pos2` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '站位2',
  `tactical_grid` tinyint(4) NOT NULL DEFAULT '0' COMMENT '玩家选中的战术格子',
  `pos3` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '站位3',
  `pos4` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '站位4',
  `pos5` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '站位5',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家阵型站位';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_friend`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_friend` (
  `id` bigint(20) NOT NULL COMMENT '好友关系ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `friend_pid` bigint(20) NOT NULL COMMENT '好友ID',
  `friend_nick` varchar(50) NOT NULL DEFAULT '' COMMENT '玩家昵称',
  `friend_role_id` tinyint(4) NOT NULL COMMENT '好友角色ID',
  `friend_mode` tinyint(4) NOT NULL COMMENT '好友关系:0陌生人,1仅关注,2仅被关注,3互相关注(好友)',
  `last_chat_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最后聊天时间',
  `friend_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '成为好友时间',
  `send_heart_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '上次送爱心时间',
  `block_mode` tinyint(1) NOT NULL DEFAULT '0' COMMENT '黑名单状态:0-否,1-是',
  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家好友列表';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_friend_chat`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_friend_chat` (
  `id` bigint(20) NOT NULL COMMENT '消息ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `friend_pid` bigint(20) NOT NULL COMMENT '对方玩家ID',
  `mode` tinyint(4) NOT NULL COMMENT '1发送，2接收',
  `send_time` bigint(20) NOT NULL COMMENT '发送时间戳',
  `message` varchar(140) NOT NULL COMMENT '消息内容',
  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家聊天记录';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_friend_state`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_friend_state` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `delete_day_count` int(11) NOT NULL DEFAULT '0' COMMENT '每日删除好友数量',
  `delete_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '删除好友时间',
  `exist_offline_chat` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0没有离线消息，1有离线消息',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家好友功能状态数据';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_func_key`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_func_key` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `key` smallint(6) NOT NULL COMMENT '功能权值',
  `played_key` smallint(6) NOT NULL DEFAULT '0' COMMENT '已播放提示的功能',
  `unique_key` bigint(20) NOT NULL DEFAULT '0' COMMENT '已开启功能的唯一权值',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家功能开放表';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_ghost`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_ghost` (
  `id` bigint(20) NOT NULL COMMENT '主键',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `ghost_id` smallint(6) NOT NULL COMMENT '魂侍ID',
  `star` tinyint(4) NOT NULL DEFAULT '0' COMMENT '星级',
  `level` smallint(6) NOT NULL DEFAULT '1' COMMENT '魂侍等级',
  `exp` bigint(20) NOT NULL DEFAULT '0' COMMENT '魂侍经验',
  `pos` smallint(6) NOT NULL COMMENT '位置',
  `is_new` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否新魂侍',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家魂侍表';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_ghost_equipment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_ghost_equipment` (
  `id` bigint(20) NOT NULL COMMENT '主键',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `role_id` tinyint(4) NOT NULL COMMENT '角色ID',
  `ghost_power` tinyint(4) NOT NULL COMMENT '魂力',
  `pos1` bigint(20) NOT NULL COMMENT '装备位置1的魂侍id',
  `pos2` bigint(20) NOT NULL COMMENT '装备位置2的魂侍id',
  `pos3` bigint(20) NOT NULL COMMENT '装备位置3的魂侍id',
  `pos4` bigint(20) NOT NULL COMMENT '装备位置4的魂侍id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家魂侍装备表';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_ghost_mission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_ghost_mission` (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `mission_id` smallint(6) NOT NULL COMMENT '关卡主键id',
  `ghost_unique_key` smallint(6) NOT NULL DEFAULT '0' COMMENT '获得魂侍的信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家魂侍副本表';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_ghost_state`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_ghost_state` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `purify_day_count` bigint(20) DEFAULT '0' COMMENT '每日净化次数',
  `ghost_unique_key` smallint(6) NOT NULL DEFAULT '0' COMMENT '获得金魂的信息',
  `purify_update_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '净化时间',
  `ghost_mission_key` int(11) NOT NULL DEFAULT '0' COMMENT '开启影界最大权值',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家魂侍状态表';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_heart`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_heart` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `value` smallint(6) NOT NULL COMMENT '爱心值',
  `update_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最后更新时间',
  `add_day_count` int(11) NOT NULL DEFAULT '0' COMMENT '每日领取数量',
  `add_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最后领取时间',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家爱心表';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_heart_draw`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_heart_draw` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `draw_type` tinyint(4) NOT NULL COMMENT '抽奖类型（1-大转盘；2-刮刮卡）',
  `daily_num` tinyint(4) NOT NULL COMMENT '当日已抽次数',
  `draw_time` bigint(20) NOT NULL COMMENT '最近一次抽奖时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=547 DEFAULT CHARSET=utf8mb4 COMMENT='玩家爱心抽奖';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_heart_draw_card_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_heart_draw_card_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `award_type` tinyint(4) NOT NULL COMMENT '奖品类型（1-铜钱；2-元宝；3-道具）',
  `award_num` smallint(6) NOT NULL COMMENT '奖品数量',
  `item_id` smallint(6) DEFAULT '0' COMMENT '道具奖品ID',
  `draw_time` bigint(20) NOT NULL COMMENT '抽奖时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家爱心刮刮卡抽奖记录';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_heart_draw_wheel_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_heart_draw_wheel_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `award_type` tinyint(4) NOT NULL COMMENT '奖品类型（1-铜钱；2-元宝；3-道具）',
  `award_num` smallint(6) NOT NULL COMMENT '奖品数量',
  `item_id` smallint(6) DEFAULT '0' COMMENT '道具奖品ID',
  `draw_time` bigint(20) NOT NULL COMMENT '抽奖时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=207 DEFAULT CHARSET=utf8mb4 COMMENT='玩家爱心大转盘抽奖记录';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_info` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `ingot` bigint(20) NOT NULL DEFAULT '0' COMMENT '元宝',
  `coins` bigint(20) NOT NULL DEFAULT '0' COMMENT '铜钱',
  `new_mail_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '新邮件数',
  `last_login_time` bigint(20) NOT NULL COMMENT '上次登录时间',
  `last_offline_time` bigint(20) NOT NULL COMMENT '上次下线时间',
  `total_online_time` bigint(20) NOT NULL COMMENT '总在线时间',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家信息表';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_item` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `item_id` smallint(6) NOT NULL COMMENT '物品ID',
  `num` smallint(6) NOT NULL COMMENT '数量',
  `is_dressed` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否被装备',
  `buy_back_state` tinyint(1) NOT NULL DEFAULT '0' COMMENT '记录物品是否在回购栏',
  `appendix_id` bigint(20) DEFAULT '0' COMMENT '附加属性ID',
  `refine_level` tinyint(4) NOT NULL DEFAULT '0' COMMENT '精练等级',
  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家物品';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_item_appendix`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_item_appendix` (
  `id` bigint(20) NOT NULL COMMENT 'ID',
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
  `recast_attr` tinyint(4) DEFAULT '0' COMMENT '重铸属性',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家装备追加属性表';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_item_buyback`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_item_buyback` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `back_id1` bigint(20) NOT NULL DEFAULT '0' COMMENT '回购格子1,player_item表主键ID',
  `back_id2` bigint(20) NOT NULL DEFAULT '0' COMMENT '回购格子2,player_item表主键ID',
  `back_id3` bigint(20) NOT NULL DEFAULT '0' COMMENT '回购格子3,player_item表主键ID',
  `back_id4` bigint(20) NOT NULL DEFAULT '0' COMMENT '回购格子4,player_item表主键ID',
  `back_id5` bigint(20) NOT NULL DEFAULT '0' COMMENT '回购格子5,player_item表主键ID',
  `back_id6` bigint(20) NOT NULL DEFAULT '0' COMMENT '回购格子6,player_item表主键ID',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家物品回购表';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_item_recast_state`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_item_recast_state` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `player_item_id` bigint(20) NOT NULL COMMENT '玩家装备ID',
  `selected_attr` tinyint(4) NOT NULL COMMENT '选中的属性',
  `attr1_type` tinyint(4) NOT NULL COMMENT '重铸属性1类型',
  `attr1_value` int(11) NOT NULL COMMENT '重铸属性1数值',
  `attr2_type` tinyint(4) NOT NULL COMMENT '重铸属性2类型',
  `attr2_value` int(11) NOT NULL COMMENT '重铸属性2数值',
  `attr3_type` tinyint(4) NOT NULL COMMENT '重铸属性3类型',
  `attr3_value` int(11) NOT NULL COMMENT '重铸属性3数值',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家装备重铸状态';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_mail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_mail` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '玩家邮件ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `mail_id` int(11) NOT NULL COMMENT '邮件模版ID',
  `state` tinyint(4) NOT NULL COMMENT '0未读，1已读',
  `send_time` bigint(20) NOT NULL COMMENT '发送时间戳',
  `parameters` varchar(1024) NOT NULL COMMENT '模版参数',
  `have_attachment` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否有附件',
  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`),
  KEY `send_time` (`send_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1981 DEFAULT CHARSET=utf8mb4 COMMENT='玩家邮件表';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_mail_attachment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_mail_attachment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '玩家邮件附件ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `player_mail_id` bigint(20) NOT NULL COMMENT 'player_mail 主键ID',
  `attachment_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '附件类型',
  `item_id` smallint(6) NOT NULL COMMENT '物品',
  `item_num` bigint(20) NOT NULL DEFAULT '0' COMMENT '数量',
  PRIMARY KEY (`id`),
  KEY `idx_pid_mail` (`pid`)
) ENGINE=InnoDB AUTO_INCREMENT=1970 DEFAULT CHARSET=utf8mb4 COMMENT='玩家邮件附件表';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_mission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_mission` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `key` int(11) NOT NULL COMMENT '拥有的区域钥匙数',
  `max_order` tinyint(4) NOT NULL COMMENT '已开启区域的最大序号',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家区域数据';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_mission_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_mission_level` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `lock` int(11) NOT NULL COMMENT '当前的关卡权值',
  `max_lock` int(11) NOT NULL COMMENT '已开启的关卡最大权值',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家区域关卡数据';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_mission_level_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_mission_level_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '记录ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `mission_id` smallint(6) NOT NULL COMMENT '区域ID',
  `mission_level_id` int(11) NOT NULL COMMENT '开启的关卡ID',
  `open_time` bigint(20) NOT NULL COMMENT '关卡开启时间',
  `score` int(11) NOT NULL DEFAULT '0' COMMENT 'boss战得分',
  `round` tinyint(4) NOT NULL DEFAULT '0' COMMENT '通关回合数',
  `daily_num` tinyint(4) NOT NULL COMMENT '当日已进入关卡的次数',
  `last_enter_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最后一次进入时间',
  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`),
  KEY `idx_mission_id` (`mission_id`)
) ENGINE=InnoDB AUTO_INCREMENT=409 DEFAULT CHARSET=utf8 COMMENT='关卡记录';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_mission_level_state_bin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_mission_level_state_bin` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `bin` blob NOT NULL COMMENT '状态MissionLevelState',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家区域关卡状态保存';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_mission_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_mission_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '记录ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `town_id` smallint(6) NOT NULL COMMENT '城镇ID',
  `mission_id` smallint(6) NOT NULL COMMENT '开启的区域ID',
  `open_time` bigint(20) NOT NULL COMMENT '开启的区域时间',
  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`),
  KEY `idx_town_id` (`town_id`)
) ENGINE=InnoDB AUTO_INCREMENT=372 DEFAULT CHARSET=utf8 COMMENT='区域记录';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_multi_level_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_multi_level_info` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `buddy_role_id` tinyint(4) NOT NULL COMMENT '上阵伙伴角色模板ID',
  `buddy_row` tinyint(4) NOT NULL COMMENT '上阵伙伴所在行（1或2)',
  `tactical_grid` tinyint(4) NOT NULL DEFAULT '0' COMMENT '战术',
  `daily_num` tinyint(4) DEFAULT '0' COMMENT '今日已战斗次数',
  `battle_time` bigint(20) DEFAULT '0' COMMENT '最近一次战斗时间',
  `lock` int(11) NOT NULL DEFAULT '0' COMMENT '关卡开启权值',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家多人关卡信息';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_physical`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_physical` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `value` smallint(6) NOT NULL COMMENT '体力值',
  `update_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最后更新时间',
  `buy_count` bigint(20) DEFAULT '0' COMMENT '购买次数',
  `buy_update_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '购买次数更新时间',
  `daily_count` tinyint(1) DEFAULT '0' COMMENT '当天购买次数',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家体力表';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_quest`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_quest` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `quest_id` smallint(6) NOT NULL COMMENT '当前任务ID',
  `state` tinyint(4) DEFAULT NULL,
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家任务';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_realm`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_realm` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `role_id` tinyint(4) NOT NULL COMMENT '角色ID',
  `realm_level` smallint(6) NOT NULL COMMENT '角色境界等级',
  `realm_exp` bigint(20) NOT NULL COMMENT '角色境界经验',
  `realm_class` smallint(6) NOT NULL COMMENT '角色境界阶级',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家角色境界表';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_resource_level_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_resource_level_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '记录ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `max_level` smallint(6) NOT NULL COMMENT '允许开放的等级上限',
  `level_id` int(11) NOT NULL COMMENT '关卡ID',
  `daily_num` tinyint(4) NOT NULL COMMENT '当日已进入关卡的次数',
  `last_enter_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最后一次进入时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COMMENT='玩家资源关卡记录';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_role` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '玩家角色ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `role_id` tinyint(4) NOT NULL COMMENT '角色模板ID',
  `level` smallint(6) NOT NULL COMMENT '等级',
  `exp` bigint(6) NOT NULL COMMENT '经验',
  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`)
) ENGINE=InnoDB AUTO_INCREMENT=414 DEFAULT CHARSET=utf8 COMMENT='玩家角色数据';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_skill`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_skill` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `role_id` tinyint(4) NOT NULL COMMENT '角色ID',
  `skill_id` smallint(6) NOT NULL COMMENT '绝招ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家角色绝招表';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_sword_soul`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_sword_soul` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '玩家物品ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `pos` smallint(6) NOT NULL COMMENT '位置',
  `sword_soul_id` smallint(6) NOT NULL COMMENT '剑心ID',
  `exp` int(11) NOT NULL COMMENT '当前经验',
  `level` tinyint(4) NOT NULL COMMENT '等级',
  PRIMARY KEY (`id`),
  KEY `ix_player_sword_soul_pid` (`pid`)
) ENGINE=InnoDB AUTO_INCREMENT=1181 DEFAULT CHARSET=utf8 COMMENT='玩家剑心数据';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_sword_soul_equipment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_sword_soul_equipment` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `role_id` tinyint(4) NOT NULL COMMENT '角色ID',
  `is_equipped_xuanyuan` tinyint(4) NOT NULL DEFAULT '0',
  `type_bits` bigint(20) NOT NULL COMMENT '所有装备剑心类型的位标记',
  `pos1` bigint(20) NOT NULL COMMENT '装备位置1的剑心',
  `pos2` bigint(20) NOT NULL COMMENT '装备位置2的剑心',
  `pos3` bigint(20) NOT NULL COMMENT '装备位置3的剑心',
  `pos4` bigint(20) NOT NULL COMMENT '装备位置4的剑心',
  `pos5` bigint(20) NOT NULL COMMENT '装备位置5的剑心',
  `pos6` bigint(20) NOT NULL COMMENT '装备位置6的剑心',
  `pos7` bigint(20) NOT NULL COMMENT '装备位置7的剑心',
  `pos8` bigint(20) NOT NULL COMMENT '装备位置8的剑心',
  `pos9` bigint(20) NOT NULL COMMENT '装备位置9的剑心',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家剑心装备表';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_sword_soul_state`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_sword_soul_state` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `fragment_num` smallint(6) NOT NULL COMMENT '碎片数量',
  `box_state` tinyint(4) NOT NULL COMMENT '开箱子的状态(位操作)',
  `last_is_ingot` tinyint(4) NOT NULL COMMENT '上次是用为元宝开箱',
  `num` tinyint(4) NOT NULL COMMENT '今日次数',
  `free_num` tinyint(4) NOT NULL,
  `ingot_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '当日已元宝拔剑次数',
  `update_time` bigint(20) NOT NULL COMMENT '更新时间',
  `is_first_time` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否是第一次拔剑',
  `protect_num` tinyint(4) NOT NULL DEFAULT '0' COMMENT '概率保护次数',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家拔剑状态';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_tower_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_tower_level` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `floor` smallint(6) NOT NULL COMMENT '当前层数',
  `battle_state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '当前楼层战斗状态(0--从未打过; 1--失败)',
  `open_time` bigint(20) NOT NULL COMMENT '开启当前层数的时间',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家战斗力';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_town`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_town` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `town_id` smallint(6) NOT NULL COMMENT '当前玩家所处的城镇ID',
  `lock` int(11) NOT NULL COMMENT '当前拥有的城镇权值',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家城镇数据';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_use_skill`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_use_skill` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `role_id` tinyint(4) NOT NULL COMMENT '角色ID',
  `skill_id1` smallint(6) NOT NULL DEFAULT '0' COMMENT '招式1',
  `skill_id2` smallint(6) NOT NULL DEFAULT '0' COMMENT '招式2',
  `skill_id3` smallint(6) NOT NULL DEFAULT '0' COMMENT '招式3',
  `skill_id0` smallint(6) NOT NULL DEFAULT '0' COMMENT '主角默认招式',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家角色当前使用的绝招表';
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;


THESQL2
	);
}
?>