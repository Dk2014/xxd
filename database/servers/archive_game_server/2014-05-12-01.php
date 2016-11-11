<?php /* dump file */

	db_execute($db, <<<THESQL1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
DROP TABLE IF EXISTS `enemy_deploy_form`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `enemy_deploy_form` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL COMMENT '关联此阵法的某表唯一ID',
  `battle_type` tinyint(4) NOT NULL COMMENT '战场类型(0--关卡;)',
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
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COMMENT='怪物阵法表单';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `enemy_deploy_form` DISABLE KEYS */;
INSERT INTO `enemy_deploy_form` VALUES (1,93,0,0,3,2,3,0,0,0,0,0,0,0,0,0,0,0),(2,91,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(4,6,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(5,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(6,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(7,36,0,0,24,8,24,0,0,0,25,0,0,0,0,0,0,0),(8,18,0,0,3,4,3,0,0,0,2,0,0,0,0,0,0,0),(9,27,0,0,5,8,5,0,0,0,0,0,0,0,0,0,0,0),(10,45,0,0,9,0,9,0,0,0,11,0,0,0,0,0,0,0),(11,54,0,0,27,0,27,0,0,0,14,0,0,0,0,13,0,0),(12,87,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0);
/*!40000 ALTER TABLE `enemy_deploy_form` ENABLE KEYS */;
DROP TABLE IF EXISTS `enemy_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
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
  `sunder_max_value` int(11) NOT NULL COMMENT '护甲值',
  `sunder_min_hurt_rate` int(11) NOT NULL COMMENT '破甲前起始的伤害转换率（百分比）',
  `sunder_end_hurt_rate` int(11) NOT NULL COMMENT '破甲后的伤害转换率（百分比）',
  `sunder_attack` int(11) NOT NULL COMMENT '攻击破甲值',
  `skill_wait` tinyint(4) NOT NULL COMMENT '绝招蓄力回合',
  `release_num` tinyint(4) NOT NULL COMMENT '释放次数',
  `recover_round_num` tinyint(4) NOT NULL COMMENT '恢复回合数',
  `common_attack_num` tinyint(4) NOT NULL COMMENT '入场普通攻击次数',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COMMENT='敌人角色数据';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `enemy_role` DISABLE KEYS */;
INSERT INTO `enemy_role` VALUES (1,'草妖','Grass',10,300,0,0,50,0,0,0,0,0,0,0,0,0,0,0,50,100,200,1,0,1,2,0),(2,'二货兔','ErHuoTu',5,500,0,0,20,0,0,0,0,0,0,0,0,0,0,0,100,100,200,1,0,1,2,0),(3,'竹筒精','ZhuTongJing',3,50,0,0,5,0,0,0,0,0,0,0,0,0,0,0,50,100,200,1,0,1,2,0),(4,'林精','LinJing',4,100,0,0,10,0,0,0,0,0,0,0,0,0,0,0,50,100,200,1,0,1,2,0),(5,'黑狼','HeiLang',5,200,0,0,25,0,0,0,0,0,0,0,0,0,0,0,50,100,200,1,0,1,2,0),(6,'鬼火','GuiHuo',6,200,0,0,30,0,0,0,0,0,0,0,0,0,0,0,50,100,200,1,0,1,2,0),(7,'灯笼怪','DengLongGuai',7,250,0,0,35,0,0,0,0,0,0,0,0,0,0,0,50,100,200,1,0,1,2,0),(8,'天狼妖','TianLangYao',7,1000,0,0,60,0,0,0,0,0,0,0,0,0,0,0,100,100,200,1,0,1,2,0),(9,'莲藕精','LianOuJing',9,300,0,0,50,0,0,0,0,0,0,0,0,0,0,0,50,100,200,1,0,1,2,0),(10,'迷路的林精','LinJing',10,300,0,0,60,0,0,0,0,0,0,0,0,0,0,0,50,100,200,1,0,1,2,0),(11,'金蟾王','JingChanWang',11,2000,0,0,100,0,0,0,0,0,0,0,0,0,0,0,100,100,200,1,0,1,2,0),(12,'水妖','ShuiYao',14,450,0,0,120,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,1,2,0),(13,'毒蛇','DuShe',12,350,0,0,80,0,0,0,0,0,0,0,0,0,0,0,50,100,200,1,0,1,2,0),(14,'剧毒臭泥','JuDuChouNi',13,2500,0,0,130,0,0,0,0,0,0,0,0,0,0,0,100,100,200,1,0,1,2,0),(15,'火蝎','HuoXie',13,400,0,0,100,0,0,0,0,0,0,0,0,0,0,0,50,100,200,1,0,1,2,0),(16,'燃魁','RanKui',14,400,0,0,110,0,0,0,0,0,0,0,0,0,0,0,50,100,200,1,0,1,2,0),(17,'熔岩虫','RongYanChong',16,300,0,0,150,0,0,0,0,0,0,0,0,0,0,0,50,100,200,1,0,1,2,0),(18,'燃烧的天狼妖','RanShaoDeTianLangYao',15,3500,0,0,160,0,0,0,0,0,0,0,0,0,0,0,100,100,200,1,0,1,2,0),(19,'炎龙','YanLong',17,4000,0,0,200,0,0,0,0,0,0,0,0,0,0,0,100,100,200,1,0,1,2,0),(20,'穿山甲','ChuanShanJia',17,450,0,0,180,0,0,0,0,0,0,0,0,0,0,0,50,100,200,1,0,1,2,0),(21,'矿工','KuangGong',18,450,0,0,190,0,0,0,0,0,0,0,0,0,0,0,50,100,200,1,0,1,2,0),(22,'拳之守卫','QuanZhiShouWei',19,500,0,0,200,0,0,0,0,0,0,0,0,0,0,0,50,100,200,1,0,1,2,0),(23,'古代武圣','GuDaiWuSheng',20,5000,0,0,300,0,0,0,0,0,0,0,0,0,0,0,100,100,200,1,0,1,2,0),(24,'阴影','YinYing',7,250,0,0,35,0,0,0,0,0,0,0,0,0,0,0,50,100,200,1,0,1,2,0),(25,'影魔','YingMo',8,250,0,0,40,0,0,0,0,0,0,0,0,0,0,0,50,100,200,1,0,1,2,0),(26,'妖龙','YaoLong',9,1500,0,0,70,0,0,0,0,0,0,0,0,0,0,0,100,100,200,1,0,1,2,0),(27,'黑翼巨蝠','HeYiJuFu',12,350,0,0,80,0,0,0,0,0,0,0,0,0,0,0,50,100,200,1,0,1,2,0);
/*!40000 ALTER TABLE `enemy_role` ENABLE KEYS */;
DROP TABLE IF EXISTS `item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `item` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '物品ID',
  `type_id` int(11) NOT NULL COMMENT '类型ID',
  `quality` tinyint(4) DEFAULT NULL COMMENT '品质',
  `name` varchar(20) NOT NULL COMMENT '物品名称',
  `level` int(11) DEFAULT NULL COMMENT '需求等级',
  `desc` varchar(100) DEFAULT NULL COMMENT '物品描述',
  `price` int(11) NOT NULL DEFAULT '0' COMMENT '物品售价',
  `can_use` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否可在格子中使用，0不能，1反之',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='物品';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `item` DISABLE KEYS */;
INSERT INTO `item` VALUES (1,2,1,'真气龙珠',1,'可用于提升角色武功境界等级',0,0);
/*!40000 ALTER TABLE `item` ENABLE KEYS */;
DROP TABLE IF EXISTS `item_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `item_type` (
  `id` tinyint(4) NOT NULL AUTO_INCREMENT COMMENT '类型ID',
  `name` varchar(10) NOT NULL COMMENT '类型名称',
  `max_num_in_pos` smallint(6) NOT NULL DEFAULT '1' COMMENT '每个位置最多可堆叠的数量',
  `sign` varchar(50) DEFAULT '' COMMENT '类型标志',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='物品类型';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `item_type` DISABLE KEYS */;
INSERT INTO `item_type` VALUES (1,'特殊材料',999,'SpecialMaterial'),(2,'材料',99,'Stuff');
/*!40000 ALTER TABLE `item_type` ENABLE KEYS */;
DROP TABLE IF EXISTS `mission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `mission` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '区域ID',
  `town_id` smallint(6) NOT NULL COMMENT '城镇ID',
  `keys` int(11) NOT NULL COMMENT '开启钥匙数',
  `name` varchar(10) NOT NULL COMMENT '区域名称',
  `sign` varchar(50) NOT NULL COMMENT '资源标识',
  `order` tinyint(4) NOT NULL COMMENT '区域开启顺序',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='城镇区域';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mission` DISABLE KEYS */;
INSERT INTO `mission` VALUES (1,1,0,'青竹林','QingZhuLin',1),(2,1,6,'黑夜森林','HeiYeShenLin',2),(3,1,6,'莲花峰','LianHuaFeng',3),(4,1,6,'熔岩火山','RongYanHuoShan',4),(5,1,6,'宝藏密室','BaoZhangMiShi',5);
/*!40000 ALTER TABLE `mission` ENABLE KEYS */;
DROP TABLE IF EXISTS `mission_enemy`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
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
  `boss_id` int(11) NOT NULL COMMENT 'boss id',
  `order` tinyint(4) NOT NULL COMMENT '顺序',
  `talk` varchar(200) NOT NULL DEFAULT '' COMMENT '副本对话（怪物旁白）',
  `boss_dir` tinyint(4) NOT NULL COMMENT '怪物朝向(0--左;1--右)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=94 DEFAULT CHARSET=utf8mb4 COMMENT='副本敌人';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mission_enemy` DISABLE KEYS */;
INSERT INTO `mission_enemy` VALUES (1,1,1,520,782,3,100,0,0,0,0,0,0,0,0,0,1,'',0),(2,2,2,520,782,3,100,0,0,0,0,0,0,0,0,0,1,'',0),(4,1,2,840,546,3,100,0,0,0,0,0,0,0,0,0,2,'',0),(5,1,2,1230,660,3,100,0,0,0,0,0,0,0,0,0,3,'',0),(6,2,2,840,546,3,100,0,0,0,0,0,0,0,0,0,2,'',0),(7,2,2,1230,660,3,100,0,0,0,0,0,0,0,0,0,3,'',0),(10,4,3,437,464,3,80,4,20,0,0,0,0,0,0,0,10,'',0),(11,4,3,1022,369,3,80,4,20,0,0,0,0,0,0,0,20,'',0),(12,4,3,1607,359,3,80,4,20,0,0,0,0,0,0,0,30,'',0),(13,5,3,437,464,3,80,4,20,0,0,0,0,0,0,0,10,'',0),(14,5,3,1022,369,3,80,4,20,0,0,0,0,0,0,0,20,'',0),(15,5,3,1607,359,3,80,4,20,0,0,0,0,0,0,0,30,'',0),(16,6,3,959,314,3,80,4,20,0,0,0,0,0,0,0,10,'',0),(17,6,3,938,809,3,80,4,20,0,0,0,0,0,0,0,20,'',0),(18,6,3,1482,518,3,80,4,20,0,0,0,0,0,0,2,30,'',0),(19,7,3,806,966,5,60,6,30,7,0,0,0,0,10,0,1,'',0),(20,7,3,1280,676,5,60,6,30,7,0,0,0,0,10,0,2,'',0),(21,7,3,1724,316,5,60,6,30,7,0,0,0,0,10,0,3,'',0),(22,8,3,806,966,5,70,6,20,7,0,0,0,0,10,0,1,'',0),(23,8,3,1280,676,5,70,6,20,7,0,0,0,0,10,0,2,'',0),(24,8,3,1724,316,5,70,6,20,7,0,0,0,0,10,0,3,'',0),(25,9,3,542,504,5,70,6,20,7,0,0,0,0,10,0,1,'',0),(26,9,3,878,420,5,70,6,20,7,0,0,0,0,10,0,2,'',0),(27,9,3,1202,581,5,70,6,20,7,0,0,0,0,10,8,3,'',0),(28,10,3,760,550,24,60,6,20,7,0,0,0,0,20,0,1,'',0),(29,10,3,1140,900,24,60,6,20,7,0,0,0,0,20,0,2,'',0),(30,10,3,1542,478,24,60,6,20,7,0,0,0,0,20,0,3,'',0),(31,11,3,760,550,24,70,6,20,25,0,0,0,0,10,0,1,'',0),(32,11,3,1140,900,24,70,6,20,25,0,0,0,0,10,0,2,'',0),(33,11,3,1542,478,24,70,6,20,25,0,0,0,0,10,0,3,'',0),(34,12,3,612,447,24,60,7,20,25,0,0,0,0,20,0,1,'',0),(35,12,3,1148,302,24,60,7,20,25,0,0,0,0,20,0,2,'',0),(36,12,3,1050,620,24,60,7,20,25,0,0,0,0,20,8,3,'',0),(37,13,3,876,477,9,60,1,30,10,0,0,0,0,10,0,1,'',0),(38,13,3,1620,894,9,60,1,30,10,0,0,0,0,10,0,2,'',0),(39,13,3,984,1300,9,60,1,30,10,0,0,0,0,10,0,3,'',0),(40,14,3,876,477,9,70,1,20,10,0,0,0,0,10,0,1,'',0),(41,14,3,1620,894,9,70,1,20,10,0,0,0,0,10,0,2,'',0),(42,14,3,984,1300,9,70,1,20,10,0,0,0,0,10,0,3,'',0),(43,15,3,1136,300,9,70,1,20,10,0,0,0,0,10,0,1,'',0),(44,15,3,947,404,9,70,1,20,10,0,0,0,0,10,0,2,'',0),(45,15,3,698,548,9,70,1,20,10,0,0,0,0,10,11,3,'',0),(46,16,3,933,351,27,60,13,40,0,0,0,0,0,0,0,1,'',0),(47,16,3,822,756,27,60,13,40,0,0,0,0,0,0,0,2,'',0),(48,16,3,998,1227,27,60,13,40,0,0,0,0,0,0,0,3,'',0),(49,17,3,933,351,27,50,13,50,0,0,0,0,0,0,0,1,'',0),(50,17,3,822,756,27,50,13,50,0,0,0,0,0,0,0,2,'',0),(51,17,3,998,1227,27,50,13,50,0,0,0,0,0,0,0,3,'',0),(52,18,3,630,375,27,60,13,40,0,0,0,0,0,0,0,1,'',0),(53,18,3,1470,386,27,60,13,40,0,0,0,0,0,0,0,2,'',0),(54,18,3,1265,650,27,60,13,40,0,0,0,0,0,0,14,3,'',0),(55,19,3,422,569,15,50,16,40,17,0,0,0,0,10,0,1,'',0),(56,19,3,870,305,15,50,16,40,17,0,0,0,0,10,0,2,'',0),(57,19,3,1376,530,15,50,16,40,17,0,0,0,0,10,0,3,'',0),(58,20,3,1023,762,15,50,16,40,17,0,0,0,0,10,0,1,'',0),(59,20,3,1740,1218,15,50,16,40,17,0,0,0,0,10,0,2,'',0),(60,20,3,900,1839,15,50,16,40,17,0,0,0,0,10,0,3,'',0),(61,21,3,1527,348,15,50,16,40,17,0,0,0,0,10,0,1,'',0),(62,21,3,464,216,15,50,16,40,17,0,0,0,0,10,0,2,'',0),(63,21,3,860,528,15,50,16,40,17,0,0,0,0,10,19,3,'',0),(64,23,3,1023,762,15,50,16,40,17,0,0,0,0,10,0,1,'',0),(65,23,3,1740,1218,15,50,16,40,17,0,0,0,0,10,0,2,'',0),(66,23,3,900,1839,15,50,16,40,17,0,0,0,0,10,0,3,'',0),(67,22,3,422,569,15,50,16,40,17,0,0,0,0,10,0,1,'',0),(68,22,3,870,305,15,50,16,40,17,0,0,0,0,10,0,2,'',0),(69,22,3,1376,530,15,50,16,40,17,0,0,0,0,10,0,3,'',0),(70,24,3,1527,348,15,50,16,40,17,0,0,0,0,10,0,1,'',0),(71,24,3,464,216,15,50,16,40,17,0,0,0,0,10,0,2,'',0),(72,24,3,860,528,15,50,16,40,17,0,0,0,0,10,19,3,'',0),(73,25,3,606,413,21,70,20,20,22,0,0,0,0,10,0,1,'',0),(74,25,3,1134,327,21,70,20,20,22,0,0,0,0,10,0,2,'',0),(75,25,3,1419,587,21,70,20,20,22,0,0,0,0,10,0,3,'',0),(76,26,3,606,413,21,70,20,20,22,0,0,0,0,10,0,1,'',0),(77,26,3,1134,327,21,70,20,20,22,0,0,0,0,10,0,2,'',0),(78,26,3,1419,587,21,70,20,20,22,0,0,0,0,10,0,3,'',0),(79,28,3,606,413,21,70,20,20,22,0,0,0,0,10,0,1,'',0),(80,28,3,1134,327,21,70,20,20,22,0,0,0,0,10,0,2,'',0),(81,28,3,1419,587,21,70,20,20,22,0,0,0,0,10,0,3,'',0),(82,29,3,606,413,21,70,20,20,22,0,0,0,0,10,0,1,'',0),(83,29,3,1134,327,21,70,20,20,22,0,0,0,0,10,0,2,'',0),(84,29,3,1419,587,21,70,20,20,22,0,0,0,0,10,0,3,'',0),(85,27,3,1152,530,22,80,21,10,20,0,0,0,0,10,0,1,'',0),(86,27,3,722,525,22,80,21,10,20,0,0,0,0,10,0,2,'',0),(87,27,3,542,686,22,80,21,10,20,0,0,0,0,10,23,3,'',0),(88,30,3,1152,530,22,80,21,10,20,0,0,0,0,10,0,1,'',0),(89,30,3,722,525,22,80,21,10,20,0,0,0,0,10,0,2,'',0),(90,30,3,542,686,22,80,21,10,20,0,0,0,0,10,23,3,'',0),(91,3,2,749,647,3,100,0,0,0,0,0,0,0,0,0,1,'',0),(92,3,2,974,221,3,100,0,0,0,0,0,0,0,0,0,2,'',0),(93,3,2,1365,438,3,100,0,0,0,0,0,0,0,0,2,3,'',0);
/*!40000 ALTER TABLE `mission_enemy` ENABLE KEYS */;
DROP TABLE IF EXISTS `mission_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `mission_level` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '区域关卡ID',
  `mission_id` smallint(6) NOT NULL COMMENT '区域ID',
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
  `award_lock` tinyint(4) NOT NULL COMMENT '通关奖励权值',
  `box_dir` tinyint(4) NOT NULL COMMENT '宝箱朝向(0--左;1--右)',
  PRIMARY KEY (`id`),
  KEY `idx_mission_id` (`mission_id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COMMENT='区域关卡配置';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mission_level` DISABLE KEYS */;
INSERT INTO `mission_level` VALUES (1,1,0,'青竹林1',0,-1,6,1398,812,1,100,1114,176,'QingZhuLin','QingZhuCunBaoZang','Music',0,0),(2,1,0,'青竹林2',1,3,12,1398,812,1,200,1114,176,'QingZhuLin','QingZhuCunBaoZang','Music',0,0),(3,1,0,'竹林深处',2,2,20,1612,316,1,400,618,190,'QingZhuCunBaoZang','QingZhuCunBaoZang','Music',0,0),(4,1,0,'溪边小径1',0,-1,6,1792,350,1,150,828,144,'LinJianXiaoDao','LinJianXiaoDao1','Music',0,0),(5,1,0,'溪边小径2',1,3,12,1792,350,1,300,828,144,'LinJianXiaoDao','LinJianXiaoDao1','Music',0,0),(6,1,0,'溪边小径3',2,2,20,1658,536,1,500,428,366,'LinJianXiaoDao1','LinJianXiaoDao1','Music',0,0),(7,2,0,'黑夜森林1',0,-1,6,2048,268,1,200,856,114,'HeiYeShenLin','HeiYeShenLinWar','Music',0,0),(8,2,0,'黑夜森林2',1,3,12,2048,268,1,400,856,114,'HeiYeShenLin','HeiYeShenLinWar','Music',0,0),(9,2,0,'黑夜森林3',2,2,20,1334,668,1,700,388,140,'HeiYeShenLin1','HeiYeShenLinWar','Music',0,0),(10,2,0,'暗影秘境1',0,-1,6,2084,200,1,300,286,152,'LeiMingZe','LeiMingJinDi','Music',0,0),(11,2,0,'暗影秘境2',1,3,12,2084,200,1,600,286,152,'LeiMingZe','LeiMingJinDi','Music',0,0),(12,2,0,'暗影禁地',2,2,20,1224,662,1,1000,124,336,'LeiMingJinDi','LeiMingJinDi','Music',0,0),(13,3,0,'莲花峰1',0,-1,6,741,1439,1,350,369,126,'LianHuaFeng','LianHuaFeng1','Music',0,0),(14,3,0,'莲花峰2',1,3,12,741,1439,1,700,369,126,'LianHuaFeng','LianHuaFeng1','Music',0,0),(15,3,0,'莲花峰顶',2,2,20,584,602,1,1300,122,1486,'LianHuaFeng1','LianHuaFeng1','Music',0,0),(16,3,0,'水溶洞1',0,-1,6,598,1380,1,400,306,150,'ShuiRongDong','ShuiRongDongDi','Music',0,0),(17,3,0,'水溶洞2',1,3,12,598,1380,1,800,306,150,'ShuiRongDong','ShuiRongDongDi','Music',0,0),(18,3,0,'水溶洞底',2,2,20,1400,770,1,1400,362,250,'ShuiRongDongDi','ShuiRongDongDi','Music',0,0),(19,4,0,'南部山道',0,-1,6,1742,764,1,500,814,104,'YanLongHuoShanShanDao','YanLongHuoShanWar','Music',0,0),(20,4,0,'内部山道1',1,3,12,288,1980,1,1000,612,246,'YanLongHuoShanNeiBu','YanLongHuoShanWar','Music',0,0),(21,4,0,'火山内部1',2,2,20,512,558,1,1800,192,1752,'YanLongHuoShanDongXue','YanLongHuoShanWar','Music',0,0),(22,4,0,'北部山道',0,-1,6,1742,764,1,600,814,104,'YanLongHuoShanShanDao','YanLongHuoShanWar','Music',0,0),(23,4,0,'内部山道2',1,3,12,288,1980,1,1200,612,246,'YanLongHuoShanNeiBu','YanLongHuoShanWar','Music',0,0),(24,4,0,'火山内部2',2,2,20,512,558,1,2100,192,1752,'YanLongHuoShanDongXue','YanLongHuoShanWar','Music',0,0),(25,5,0,'密道1',0,-1,6,1582,686,1,700,102,132,'BeiLinMiDao','BeiLinMiDaoWar','Music',0,0),(26,5,0,'密道2',1,3,12,1582,686,1,1400,102,132,'BeiLinMiDao','BeiLinMiDaoWar','Music',0,0),(27,5,0,'宝藏密室1',2,2,20,458,726,1,2400,332,1250,'BaoZhangMiShi','BaoZhangMiShiWar','Music',0,0),(28,5,0,'密道3',0,-1,6,1582,686,1,800,102,132,'BeiLinMiDao','BeiLinMiDaoWar','Music',0,0),(29,5,0,'密道4',1,3,12,1582,686,1,1600,102,134,'BeiLinMiDao','BeiLinMiDaoWar','Music',0,0),(30,5,0,'宝藏密室2',2,2,20,458,726,1,2800,332,1250,'WuShenMiShi','BaoZhangMiShiWar','Music',0,0);
/*!40000 ALTER TABLE `mission_level` ENABLE KEYS */;
DROP TABLE IF EXISTS `mission_level_box`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `mission_level_box` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `mission_level_id` int(11) NOT NULL COMMENT '关卡id',
  `order` tinyint(4) NOT NULL COMMENT '品质顺序',
  `award_type` tinyint(4) NOT NULL COMMENT '奖励类型(0--铜钱;1--道具;2--装备)',
  `award_chance` tinyint(4) NOT NULL COMMENT '奖励概率',
  `award_num` int(11) NOT NULL COMMENT '奖励数量',
  `item_id` int(11) NOT NULL DEFAULT '0' COMMENT '物品ID(物品奖励填写)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='区域关卡宝箱';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mission_level_box` DISABLE KEYS */;
/*!40000 ALTER TABLE `mission_level_box` ENABLE KEYS */;
DROP TABLE IF EXISTS `role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='角色表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `role` DISABLE KEYS */;
INSERT INTO `role` VALUES (1,'义峰','',1,0,0,0,0,0),(2,'昕苒','',1,0,0,0,0,0),(3,'叶开','',2,0,5,6,0,0),(4,'朱媛媛','',2,0,7,8,0,0);
/*!40000 ALTER TABLE `role` ENABLE KEYS */;
DROP TABLE IF EXISTS `role_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
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
) ENGINE=InnoDB AUTO_INCREMENT=77 DEFAULT CHARSET=utf8mb4 COMMENT='角色等级';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `role_level` DISABLE KEYS */;
INSERT INTO `role_level` VALUES (11,1,1,2,120,16,7,2,56,15,6,0,6,0,0,6,2,100,100,200,2,2),(12,1,2,21,140,22,9,4,62,15,6,0,6,0,0,6,2,100,100,200,2,2),(13,1,3,26,160,28,11,6,68,15,6,0,6,0,0,6,2,100,100,200,2,2),(14,1,4,26,180,34,13,8,74,15,6,0,6,0,0,6,2,100,100,200,2,2),(15,1,5,26,200,40,15,10,80,15,6,0,6,0,0,6,2,100,100,200,2,2),(16,1,6,32,220,46,17,12,86,15,6,0,6,0,0,6,2,100,100,200,2,2),(17,1,7,32,240,52,19,14,92,15,6,0,6,0,0,6,2,100,100,200,2,2),(18,1,8,32,260,58,21,16,98,15,6,0,6,0,0,6,2,100,100,200,2,2),(19,1,9,42,280,64,23,18,104,15,6,0,6,0,0,6,2,100,100,200,2,2),(20,1,10,62,300,70,25,20,110,15,6,0,6,0,0,6,2,100,100,200,2,2),(21,1,11,84,340,82,29,24,122,15,6,0,6,0,0,6,2,100,100,200,2,2),(22,1,12,94,380,94,33,28,134,15,6,0,6,0,0,6,2,100,100,200,2,2),(23,1,13,104,420,106,37,32,146,15,6,0,6,0,0,6,2,100,100,200,2,2),(24,1,14,124,460,118,41,36,158,15,6,0,6,0,0,6,2,100,100,200,2,2),(25,1,15,6000,500,130,45,40,170,15,6,0,6,0,0,6,2,100,100,200,2,2),(26,1,16,7000,540,142,49,44,182,15,6,0,6,0,0,6,2,100,100,200,2,2),(27,1,17,10000,580,154,53,48,194,15,6,0,6,0,0,6,2,100,100,200,2,2),(28,1,18,12000,620,166,57,52,206,15,6,0,6,0,0,6,2,100,100,200,2,2),(29,1,19,12000,660,178,61,56,218,15,6,0,6,0,0,6,2,100,100,200,2,2),(30,1,20,15000,700,190,65,60,230,15,6,0,6,0,0,6,2,100,100,200,2,2),(31,1,21,15000,760,208,71,66,248,15,6,0,6,0,0,6,2,125,100,200,2,2),(32,1,22,15000,820,226,77,72,266,15,6,0,6,0,0,6,2,125,100,200,2,2),(33,1,23,15000,880,244,83,78,284,15,6,0,6,0,0,6,2,125,100,200,2,2),(34,1,24,15000,940,262,89,84,302,15,6,0,6,0,0,6,2,125,100,200,2,2),(35,1,25,57000,1000,280,95,90,320,15,6,0,6,0,0,6,2,125,100,200,2,2),(36,1,26,66000,1060,298,101,96,338,15,6,0,6,0,0,6,2,125,100,200,2,2),(37,1,27,77000,1120,316,107,102,356,15,6,0,6,0,0,6,2,125,100,200,2,2),(38,1,28,88000,1180,334,113,108,374,15,6,0,6,0,0,6,2,125,100,200,2,2),(39,1,29,110000,1240,352,119,114,392,15,6,0,6,0,0,6,2,125,100,200,2,2),(40,1,30,130000,1300,370,125,120,410,15,6,0,6,0,0,6,2,125,100,200,2,2),(41,1,31,150000,1370,391,132,127,431,15,6,0,6,0,0,6,2,150,100,200,2,2),(42,1,32,180000,1440,412,139,134,452,15,6,0,6,0,0,6,2,150,100,200,2,2),(43,1,33,220000,1510,433,146,141,473,15,6,0,6,0,0,6,2,150,100,200,2,2),(44,1,34,250000,1580,454,153,148,494,15,6,0,6,0,0,6,2,150,100,200,2,2),(45,1,35,290000,1650,475,160,155,515,15,6,0,6,0,0,6,2,150,100,200,2,2),(46,1,36,420000,1720,496,167,162,536,15,6,0,6,0,0,6,2,150,100,200,2,2),(47,1,37,460000,1790,517,174,169,557,15,6,0,6,0,0,6,2,150,100,200,2,2),(48,1,38,490000,1860,538,181,176,578,15,6,0,6,0,0,6,2,150,100,200,2,2),(49,1,39,530000,1930,559,188,183,599,15,6,0,6,0,0,6,2,150,100,200,2,2),(50,1,40,570000,2000,580,195,190,620,15,6,0,6,0,0,6,2,150,100,200,2,2),(51,1,41,640000,2080,604,203,198,644,15,6,0,6,0,0,6,2,175,100,200,2,2),(52,1,42,820000,2160,628,211,206,668,15,6,0,6,0,0,6,2,175,100,200,2,2),(53,1,43,1030000,2240,652,219,214,692,15,6,0,6,0,0,6,2,175,100,200,2,2),(54,1,44,1100000,2320,676,227,222,716,15,6,0,6,0,0,6,2,175,100,200,2,2),(57,1,23,0,333,333,333,3333,33,0,0,0,0,0,0,0,0,0,0,0,0,0),(64,1,45,1160000,2400,700,235,230,740,15,6,0,6,0,0,6,2,175,100,200,2,2),(65,1,46,1230000,2480,724,243,238,764,15,6,0,6,0,0,6,2,175,100,200,2,2),(66,1,47,1310000,2560,748,251,246,788,15,6,0,6,0,0,6,2,175,100,200,2,2),(67,1,48,1380000,2640,772,259,254,812,15,6,0,6,0,0,6,2,175,100,200,2,2),(68,1,49,1460000,2720,796,267,262,836,15,6,0,6,0,0,6,2,175,100,200,2,2),(69,1,50,1540000,2800,820,275,270,860,15,6,0,6,0,0,6,2,175,100,200,2,2),(70,1,51,1640000,2890,847,284,279,887,15,6,0,6,0,0,6,2,200,100,200,2,2),(71,3,1,0,100,30,10,10,10,6,0,0,6,0,0,0,0,100,100,200,2,2),(72,3,2,0,100,30,10,10,10,6,0,0,6,0,0,0,0,100,100,200,2,2),(73,3,3,0,100,30,10,10,10,6,0,0,6,0,0,0,0,100,100,200,2,2),(74,4,1,20,100,30,10,10,10,6,0,0,6,0,0,0,0,100,100,200,2,2),(75,4,2,20,100,30,10,10,10,6,0,0,6,0,0,0,0,100,100,200,2,2),(76,4,3,20,100,30,10,10,10,6,0,0,6,0,0,0,0,100,100,200,2,2);
/*!40000 ALTER TABLE `role_level` ENABLE KEYS */;
DROP TABLE IF EXISTS `skill`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `skill` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `name` varchar(10) DEFAULT NULL COMMENT '绝招名称',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型',
  `child_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '子类型',
  `sign` varchar(30) DEFAULT NULL COMMENT '资源标识',
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
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COMMENT='绝招表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `skill` DISABLE KEYS */;
INSERT INTO `skill` VALUES (1,'闪击',1,1,'',-2,1,'单体攻击，获得2点精气',1,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 50,\"LevelAttack\": 0,\"Cul2AtkRate\": 10,\"DecPower\": 0,\"IncPower\": 2,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"LevelSunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,1,1,1,0),(2,'裂地斩',1,1,'',-2,5,'横排攻击，对多个敌人产生伤害',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 100,\"LevelAttack\": 0,\"Cul2AtkRate\": 60,\"DecPower\": 2,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"LevelSunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,2,1,0),(3,'甲溃',1,1,'',-2,10,'单体攻击，破坏敌方护甲',1,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 200,\"LevelAttack\": 0,\"Cul2AtkRate\": 60,\"DecPower\": 2,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 50,\"LevelSunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,3,1,0),(4,'必杀技',1,1,'',-2,15,'单体攻击，消耗精气产生大量伤害',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1000,\"LevelAttack\": 0,\"Cul2AtkRate\": 60,\"DecPower\": 6,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 50,\"LevelSunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,1,4,1,0),(5,'刀芒',1,1,'',3,0,'横排攻击，产生大量伤害',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 50,\"LevelAttack\": 0,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"LevelSunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,1,5,1,0),(6,'剑盾',1,3,'',3,0,'防御1回合，自身减免40%的伤害，并吸引攻击',0,0,'{\"TargetMode\": 4,\"AttackMode\": 1,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"LevelAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"LevelSunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [{\"Type\": 18, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"LevelRate\": 0, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 60, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"LevelValue\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0},{\"Type\": 19, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"LevelRate\": 0, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"LevelValue\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,1,6,1,0),(7,'风咒',1,1,'',4,0,'单体攻击',0,0,'{\"TargetMode\": 3,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 50,\"LevelAttack\": 0,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"LevelSunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,1,7,1,0),(8,'治愈之风',1,4,'',4,0,'单体治疗，优先治疗受伤最重的伙伴',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"LevelAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"LevelSunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"LevelRate\": 0, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 500, \"RawValueRate\": 0, \"AttackRate\": 33, \"SkillForceRate\": 0, \"HurtRate\": 0, \"LevelValue\": 20, \"Cul2ValueRate\": 100, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,1,8,1,0),(9,'猛击',5,0,'MengJiM',-1,0,'单体',1,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"LevelAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"LevelSunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,1);
/*!40000 ALTER TABLE `skill` ENABLE KEYS */;
DROP TABLE IF EXISTS `skill_content`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `skill_content` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `skill_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '绝招ID',
  `release_num` int(11) NOT NULL DEFAULT '0' COMMENT '释放次数',
  `recover_round_num` int(11) NOT NULL DEFAULT '0' COMMENT '恢复回合数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='绝招数据表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `skill_content` DISABLE KEYS */;
INSERT INTO `skill_content` VALUES (1,5,2,1),(2,6,1,0),(3,7,2,1),(4,8,3,2);
/*!40000 ALTER TABLE `skill_content` ENABLE KEYS */;
DROP TABLE IF EXISTS `town`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `town` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '城镇ID,-1为集会所',
  `lock` int(11) NOT NULL COMMENT '解锁权值',
  `name` varchar(10) NOT NULL DEFAULT '' COMMENT '城镇名称',
  `sign` varchar(30) NOT NULL DEFAULT '' COMMENT '资源标识',
  `music` varchar(20) NOT NULL COMMENT '音乐资源标识',
  `start_x` int(11) NOT NULL COMMENT '出生点x轴坐标',
  `start_y` int(11) NOT NULL COMMENT '出生点y轴坐标',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='城镇';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `town` DISABLE KEYS */;
INSERT INTO `town` VALUES (1,100000,'青竹村','QingZhuCun','Music',993,569),(2,666666,'天地盟','TianDiMeng','Music',1290,879);
/*!40000 ALTER TABLE `town` ENABLE KEYS */;
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
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
DROP TABLE IF EXISTS `player`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `player` (
  `id` bigint(20) NOT NULL COMMENT '玩家ID',
  `user` varchar(250) NOT NULL COMMENT '平台传递过来的用户标识',
  `nick` varchar(50) NOT NULL COMMENT '玩家昵称',
  `main_role_id` bigint(20) NOT NULL COMMENT '主角ID',
  PRIMARY KEY (`id`),
  KEY `ix_player_sign` (`user`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家基础信息';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_formation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家阵型站位';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `player_info` (
  `pid` bigint(20) NOT NULL AUTO_INCREMENT,
  `ingot` bigint(20) NOT NULL DEFAULT '0' COMMENT '元宝',
  `coins` bigint(20) NOT NULL DEFAULT '0' COMMENT '铜钱',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家信息表';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `player_item` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `item_id` smallint(6) NOT NULL COMMENT '物品ID',
  `num` smallint(6) NOT NULL COMMENT '数量',
  `is_dressed` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否被装备',
  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家物品';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_mission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `player_mission` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `key` int(11) NOT NULL COMMENT '拥有的区域钥匙数',
  `max_order` tinyint(4) NOT NULL COMMENT '已开启区域的最大序号',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家区域数据';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_mission_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `player_mission_level` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `lock` int(11) NOT NULL COMMENT '当前的关卡权值',
  `max_lock` int(11) NOT NULL COMMENT '已开启的关卡最大权值',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家区域关卡数据';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_mission_level_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `player_mission_level_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '记录ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `mission_id` smallint(6) NOT NULL COMMENT '区域ID',
  `mission_level_id` int(11) NOT NULL COMMENT '开启的关卡ID',
  `open_time` bigint(20) NOT NULL COMMENT '关卡开启时间',
  `star` tinyint(4) NOT NULL DEFAULT '0' COMMENT '通关星数',
  `round` tinyint(4) NOT NULL DEFAULT '0' COMMENT '通关回合数',
  `daily_num` tinyint(4) NOT NULL COMMENT '当日已进入关卡的次数',
  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`),
  KEY `idx_mission_id` (`mission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='关卡记录';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_mission_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `player_mission_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '记录ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `town_id` smallint(6) NOT NULL COMMENT '城镇ID',
  `mission_id` smallint(6) NOT NULL COMMENT '开启的区域ID',
  `open_time` bigint(20) NOT NULL COMMENT '开启的区域时间',
  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`),
  KEY `idx_town_id` (`town_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='区域记录';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `player_role` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '玩家角色ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `role_id` tinyint(4) NOT NULL COMMENT '角色模板ID',
  `level` smallint(6) NOT NULL COMMENT '等级',
  `exp` bigint(6) NOT NULL COMMENT '经验',
  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COMMENT='玩家角色数据';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_skill`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `player_skill` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `role_id` tinyint(4) NOT NULL COMMENT '角色ID',
  `skill_id` smallint(6) NOT NULL COMMENT '绝招ID',
  `level` smallint(6) NOT NULL DEFAULT '1' COMMENT '绝招等级',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家角色绝招表';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_town`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `player_town` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `town_id` smallint(6) NOT NULL COMMENT '当前玩家所处的城镇ID',
  `lock` int(11) NOT NULL COMMENT '当前拥有的城镇权值',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家城镇数据';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_use_skill`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `player_use_skill` (
  `id` bigint(20) NOT NULL COMMENT '主键ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `role_id` tinyint(4) NOT NULL COMMENT '角色ID',
  `skill_id1` smallint(6) NOT NULL DEFAULT '0' COMMENT '招式1',
  `skill_id2` smallint(6) NOT NULL DEFAULT '0' COMMENT '招式2',
  `skill_id3` smallint(6) NOT NULL DEFAULT '0' COMMENT '招式3',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家角色当前使用的绝招表';
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