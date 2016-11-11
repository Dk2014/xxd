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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='公告模版';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `announcement` DISABLE KEYS */;
INSERT INTO `announcement` VALUES (1,'TestAnnounce',0,'测试公告','p1,参数一','欢迎来到仙侠道{0}',0,0),(2,'TestAnnounce2',0,'测试公告2','','测试用公告登录滚动播放',0,0),(3,'TestAnnounce3',0,'测试公告3','','亲爱的各位玩家，服务器将在10分钟后关闭。我们为对广大玩家对来的不便深表歉意，各个关卡将在服务器关闭前2分钟停止进入。',0,0),(4,'TraderShowupAnnounce',0,'巡游商人出现公告','timing,出现时间;disappear,消失时间','巡游商人将在{0]出现，在{1}离开',0,0);
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
  `item2_id` smallint(6) NOT NULL COMMENT '物品2',
  `item2_num` smallint(6) NOT NULL COMMENT '物品2数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COMMENT='比武场奖励宝箱';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `arena_award_box` DISABLE KEYS */;
INSERT INTO `arena_award_box` VALUES (1,1,550,100000,297,1,272,1),(2,3,500,90000,296,1,272,1),(3,5,400,80000,296,1,272,1),(4,10,300,70000,296,1,272,1),(5,20,250,60000,296,1,272,1),(6,30,200,50000,296,1,272,1),(7,40,180,40000,295,1,272,1),(8,50,150,35000,295,1,272,1),(9,100,100,28000,295,1,239,1),(10,200,80,25000,295,1,239,1),(11,500,50,20000,295,1,239,1),(12,1000,30,15000,295,1,239,1),(13,3000,25,12000,294,1,239,1),(14,5000,20,10000,294,1,239,1),(15,7000,15,8000,294,1,0,0),(16,10000,12,5000,294,1,0,0),(17,20000,8,3000,231,1,0,0),(18,30000,5,1500,231,1,0,0),(19,50000,3,1000,0,0,0,0),(20,300000,0,500,0,0,0,0);
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
INSERT INTO `battle_item_config` VALUES (209,2,1,'[{\"mod\":\"1\",\"val\":\"5000\"}]',0,0,1),(210,2,1,'[{\"mod\":\"1\",\"val\":\"20000\"}]',0,0,1),(211,3,1,'[{\"mod\":\"1\",\"val\":\"99999\"}]',0,0,1),(212,6,1,'[{\"mod\":\"3\",\"val\":\"100\"}]',0,0,0),(213,5,1,'[{\"mod\":\"1\",\"val\":\"-10000\"}]',0,0,0),(214,5,1,'[{\"mod\":\"1\",\"val\":\"-30000\"}]',0,0,0),(215,2,3,'[{\"mod\":\"13\",\"val\":\"1000\"}]',2,1,0),(216,1,3,'[{\"mod\":\"0\",\"val\":\"10\"},{\"mod\":\"16\",\"val\":\"60\"}]',1,1,0),(217,2,3,'[{\"mod\":\"7\",\"val\":\"0\"}]',1,1,0),(250,7,2,'[{\"mod\":\"4\",\"val\":\"50\"}]',0,0,0),(251,7,2,'[{\"mod\":\"4\",\"val\":\"100\"}]',0,0,0);
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
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COMMENT='灵宠';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `battle_pet` DISABLE KEYS */;
INSERT INTO `battle_pet` VALUES (1,0,91,'',1,1,3,2,1,108,252),(2,0,92,'可在熔岩火山六捕捉',1,1,3,1,2,36,253),(3,1,93,'可在剑灵密室六捕捉',2,2,2,3,3,35,254),(4,0,116,'出没区域还是一个秘密',1,2,2,4,5,23,308),(5,1,117,'出没区域还是一个秘密',1,4,3,1,10,21,309),(6,1,118,'出没区域还是一个秘密',1,4,3,1,10,38,310),(7,1,119,'出没区域还是一个秘密',1,3,2,2,10,109,311);
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
) ENGINE=InnoDB AUTO_INCREMENT=186 DEFAULT CHARSET=utf8mb4 COMMENT='宝箱物品';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `chest_item` DISABLE KEYS */;
INSERT INTO `chest_item` VALUES (9,3,3,209,20),(10,3,3,259,1),(11,3,3,235,10),(13,4,3,263,5),(14,4,3,210,5),(15,4,3,306,1),(16,4,3,270,8),(19,5,5,37,1),(21,1,3,270,50),(23,1,3,232,20),(24,1,3,271,3),(26,8,3,305,1),(27,8,3,270,80),(28,8,3,263,50),(29,8,3,251,2),(30,8,7,275,3),(34,12,3,212,1),(36,11,3,215,1),(38,11,3,284,2),(39,9,3,263,50),(40,9,3,210,10),(41,9,5,37,3),(42,9,3,214,5),(43,8,3,235,80),(44,10,5,4,1),(45,10,5,8,1),(46,10,5,10,1),(48,11,3,214,3),(49,6,3,232,5),(51,6,3,210,1),(61,11,3,282,2),(66,10,3,211,5),(67,10,5,17,1),(68,10,5,19,1),(71,12,5,23,1),(72,12,5,11,1),(73,12,5,35,1),(74,7,5,4,1),(75,7,5,8,1),(77,7,5,21,1),(78,7,5,25,1),(86,6,5,3,1),(87,6,5,7,1),(88,6,5,15,1),(89,6,5,12,1),(90,6,3,213,1),(92,3,5,2,1),(93,3,5,6,1),(98,1,3,272,3),(99,1,3,250,5),(100,1,3,231,50),(101,2,3,270,5),(102,2,3,236,10),(103,2,3,238,10),(104,2,3,237,10),(105,2,3,231,10),(106,2,3,230,1),(107,2,3,239,1),(108,1,3,213,5),(109,1,3,256,7),(110,8,3,254,3),(111,9,3,233,7),(113,9,7,277,3),(114,10,3,254,1),(115,10,7,261,7),(116,10,3,233,10),(117,11,7,278,3),(118,11,3,216,1),(119,11,3,217,1),(120,11,3,251,5),(121,12,3,234,1),(122,12,7,276,3),(123,12,4,9,1),(124,12,4,8,1),(125,12,4,7,1),(126,12,5,29,1),(127,12,5,5,1),(128,12,5,32,1),(129,2,5,4,1),(130,2,5,7,1),(131,2,5,13,1),(132,3,7,261,1),(133,4,3,272,1),(134,4,3,271,1),(135,4,7,256,3),(136,4,7,257,3),(137,4,7,259,2),(138,5,3,250,2),(139,5,7,260,3),(140,5,7,259,3),(141,5,3,280,1),(142,6,7,274,1),(143,6,7,273,1),(144,6,7,275,1),(145,7,7,276,1),(146,7,7,277,1),(147,7,7,278,1),(148,7,3,233,3),(149,7,3,213,3),(150,6,3,263,30),(151,2,3,36,1),(152,2,3,31,1),(153,2,3,38,1),(154,2,3,37,1),(155,2,3,32,1),(156,2,3,33,1),(157,-1,3,38,1),(158,-1,3,270,10),(159,-1,3,230,1),(160,-1,5,7,1),(161,-1,3,209,3),(162,-1,3,235,10),(163,-1,3,239,1),(164,-1,3,36,1),(165,-1,3,250,3),(168,-2,3,254,1),(169,-2,3,214,5),(170,-2,5,37,3),(171,-2,3,270,50),(172,-2,3,271,1),(173,-2,3,302,10),(174,-2,3,211,3),(175,-2,3,272,1),(176,-2,4,2,1),(177,-2,5,4,1),(178,-1,3,280,1),(179,11,4,1,1),(180,11,4,4,1),(181,11,4,3,1),(182,7,4,1,1),(183,7,4,2,1),(184,1,3,302,10),(185,6,3,303,1);
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
DROP TABLE IF EXISTS `daily_quest`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `daily_quest` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `type` tinyint(4) DEFAULT '0' COMMENT '任务类型',
  `name` varchar(10) NOT NULL COMMENT '任务标题',
  `desc` varchar(240) DEFAULT '' COMMENT '简介',
  `require_min_level` int(11) NOT NULL COMMENT '要求玩家最低等级',
  `require_max_level` int(11) NOT NULL COMMENT '要求玩家最高等级',
  `require_open_day` varchar(10) DEFAULT '' COMMENT '开放日',
  `require_count` smallint(6) NOT NULL COMMENT '需要数量',
  `award_exp` int(11) NOT NULL COMMENT '奖励经验',
  `award_coins` bigint(20) NOT NULL DEFAULT '0' COMMENT '奖励铜钱',
  `award_physical` tinyint(4) NOT NULL DEFAULT '0' COMMENT '奖励体力',
  `award_item1_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品1',
  `award_item1_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品1数量',
  `award_item2_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品2',
  `award_item2_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品2数量',
  `award_item3_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品3',
  `award_item3_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品3数量',
  `award_item4_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品4',
  `award_item4_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品4数量',
  `level_type` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '关卡类型; -1 无; 0-区域关卡;1-资源关卡;2-通天塔;8-难度关卡;9-伙伴关卡;10-灵宠关卡;11-魂侍关卡',
  `level_sub_type` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '关卡子类型(-1--无;1--铜钱关卡;2--经验关卡)',
  `class` smallint(6) NOT NULL COMMENT '任务类别',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COMMENT='每日任务';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `daily_quest` DISABLE KEYS */;
INSERT INTO `daily_quest` VALUES (1,0,'击杀强敌（10级）','少侠一走，魔物又出来作怪了，劳烦少侠辛苦一趟。',10,19,'',2,100,0,0,0,0,0,0,0,0,0,0,0,0,0),(2,0,'勇闯深渊（10级）','深渊的魔物死亡后又重生了，该想个办法才行。',15,19,'',1,100,0,0,0,0,0,0,0,0,0,0,8,-1,4),(3,0,'探秘彩虹桥（20级）','传说在彩虹的尽头是神龙的大宝藏。',25,29,'',5,3000,0,0,0,0,0,0,0,0,0,0,2,-1,3),(4,0,'魂侍试炼（20级）','强大的魂侍将会帮助你逆转厄运。',20,29,'1 4 0',2,3000,0,0,0,0,0,0,0,0,0,0,11,-1,7),(5,0,'灵宠试炼（20级）','灵宠可不只是一个小小的辅助，在关键时刻它将成为你最重要的依靠。',20,29,'2 5 0',2,3000,0,0,0,0,0,0,0,0,0,0,10,-1,6),(6,0,'伙伴试炼（20级）','一个人的强大并不是真的强大，和伙伴一起成长才能让自己成为真正的大侠。',20,29,'3 6 0',2,3000,0,0,0,0,0,0,0,0,0,0,9,-1,5),(7,0,'财源滚滚（10级）','在洞天福地之中有一岛屿，岛上有取之不尽的金银珠宝，没有人知道它从何而来。',15,19,'1 3 5 0',1,100,0,0,0,0,0,0,0,0,0,0,0,1,1),(8,0,'突飞猛进（10级）','在洞天福地之中有一神殿，聚天地之灵气，传说曾是武神修炼之地。',15,19,'2 4 6 0',1,100,0,0,0,0,0,0,0,0,0,0,0,2,2),(9,0,'击杀强敌（20级）','少侠一走，魔物又出来作怪了，劳烦少侠辛苦一趟。',20,29,'',2,1000,0,0,0,0,0,0,0,0,0,0,0,0,0),(10,0,'勇闯深渊（20级）','深渊的魔物死亡后又重生了，该想个办法才行。',20,29,'',1,1000,0,0,0,0,0,0,0,0,0,0,8,-1,4),(11,0,'财源滚滚（20级）','在洞天福地之中有一岛屿，岛上有取之不尽的金银珠宝，没有人知道它从何而来。',20,29,'1 3 5 0',1,2000,0,0,0,0,0,0,0,0,0,0,1,1,1),(12,0,'突飞猛进（20级）','在洞天福地之中有一神殿，聚天地之灵气，传说曾是武神修炼之地。',20,29,'2 4 6 0',1,2000,0,0,0,0,0,0,0,0,0,0,1,2,2);
/*!40000 ALTER TABLE `daily_quest` ENABLE KEYS */;
DROP TABLE IF EXISTS `daily_sign_in_award`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `daily_sign_in_award` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `type` tinyint(4) NOT NULL COMMENT '签到类型 0--首次登录奖励 1--全局循环奖励',
  `award_type` tinyint(4) NOT NULL COMMENT '奖励类型',
  `award_id` smallint(6) NOT NULL COMMENT '奖励物品ID',
  `num` int(11) NOT NULL COMMENT '奖励数量',
  `vip_double` tinyint(4) NOT NULL DEFAULT '0' COMMENT 'vip用户获得双倍奖励',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=71 DEFAULT CHARSET=utf8mb4 COMMENT='每日签到奖励配置';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `daily_sign_in_award` DISABLE KEYS */;
INSERT INTO `daily_sign_in_award` VALUES (1,0,4,0,100,0),(2,0,1,263,50,0),(3,0,4,0,100,0),(4,0,1,254,5,0),(5,0,4,0,100,0),(6,0,1,254,5,0),(7,0,4,0,100,0),(8,0,1,284,5,0),(9,0,4,0,100,0),(10,0,1,284,5,0),(11,0,4,0,100,0),(12,0,1,274,10,0),(13,0,4,0,100,0),(14,0,1,274,10,0),(43,1,3,0,30000,0),(44,1,1,270,5,1),(45,1,4,0,100,0),(46,1,1,261,3,1),(47,1,1,263,5,0),(48,1,1,231,10,1),(49,1,4,0,100,0),(50,1,3,0,30000,0),(51,1,1,270,5,1),(52,1,4,0,100,0),(53,1,1,261,3,1),(54,1,1,263,5,0),(55,1,1,231,10,1),(56,1,4,0,100,0),(57,1,3,0,30000,0),(58,1,1,270,5,1),(59,1,4,0,100,0),(60,1,1,261,3,1),(61,1,1,263,5,0),(62,1,1,231,10,1),(63,1,4,0,100,0),(64,1,3,0,30000,0),(65,1,1,270,5,1),(66,1,4,0,100,0),(67,1,1,261,3,1),(68,1,1,263,5,0),(69,1,1,231,10,1),(70,1,4,0,100,0);
/*!40000 ALTER TABLE `daily_sign_in_award` ENABLE KEYS */;
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
INSERT INTO `enemy_boss_script` VALUES (2,'[{\"condition\":\"1\",\"val\":\"250\",\"skill_id\":\"1035\",\"power\":\"20\"}]'),(8,'[{\"condition\":\"1\",\"val\":\"2000\",\"skill_id\":\"9\",\"power\":\"120\"},{\"condition\":\"1\",\"val\":\"4000\",\"skill_id\":\"25\",\"power\":\"100\"},{\"condition\":\"1\",\"val\":\"1500\",\"skill_id\":\"21\",\"power\":\"100\"}]'),(11,'[{\"condition\":\"1\",\"val\":\"8000\",\"skill_id\":\"12\",\"power\":\"150\"},{\"condition\":\"1\",\"val\":\"6000\",\"skill_id\":\"13\",\"power\":\"150\"},{\"condition\":\"1\",\"val\":\"2000\",\"skill_id\":\"14\",\"power\":\"150\"},{\"condition\":\"1\",\"val\":\"4000\",\"skill_id\":\"33\",\"power\":\"150\"}]'),(14,'[{\"condition\":\"1\",\"val\":\"12000\",\"skill_id\":\"33\",\"power\":\"200\"},{\"condition\":\"1\",\"val\":\"10000\",\"skill_id\":\"28\",\"power\":\"200\"},{\"condition\":\"1\",\"val\":\"8000\",\"skill_id\":\"32\",\"power\":\"200\"},{\"condition\":\"1\",\"val\":\"5000\",\"skill_id\":\"34\",\"power\":\"200\"}]'),(18,'[{\"condition\":\"1\",\"val\":\"20000\",\"skill_id\":\"37\",\"power\":\"300\"},{\"condition\":\"1\",\"val\":\"15000\",\"skill_id\":\"29\",\"power\":\"300\"},{\"condition\":\"1\",\"val\":\"10000\",\"skill_id\":\"30\",\"power\":\"300\"},{\"condition\":\"1\",\"val\":\"5000\",\"skill_id\":\"30\",\"power\":\"300\"}]'),(19,'[{\"condition\":\"1\",\"val\":\"22000\",\"skill_id\":\"17\",\"power\":\"360\"},{\"condition\":\"1\",\"val\":\"17000\",\"skill_id\":\"16\",\"power\":\"360\"},{\"condition\":\"1\",\"val\":\"12000\",\"skill_id\":\"18\",\"power\":\"360\"},{\"condition\":\"1\",\"val\":\"7000\",\"skill_id\":\"18\",\"power\":\"360\"}]'),(23,'[{\"condition\":\"1\",\"val\":\"15000\",\"skill_id\":\"40\",\"power\":\"400\"},{\"condition\":\"1\",\"val\":\"20000\",\"skill_id\":\"40\",\"power\":\"400\"},{\"condition\":\"1\",\"val\":\"10000\",\"skill_id\":\"108\",\"power\":\"10000\"},{\"condition\":\"1\",\"val\":\"25000\",\"skill_id\":\"40\",\"power\":\"400\"}]'),(25,'[{\"condition\":\"1\",\"val\":\"4000\",\"skill_id\":\"45\",\"power\":\"120\"},{\"condition\":\"1\",\"val\":\"2000\",\"skill_id\":\"45\",\"power\":\"120\"},{\"condition\":\"1\",\"val\":\"5000\",\"skill_id\":\"26\",\"power\":\"120\"},{\"condition\":\"1\",\"val\":\"1000\",\"skill_id\":\"22\",\"power\":\"120\"}]'),(26,'[{\"condition\":\"1\",\"val\":\"500\",\"skill_id\":\"44\",\"power\":\"50\"}]'),(88,'[{\"condition\":\"1\",\"val\":\"1000\",\"skill_id\":\"42\",\"power\":\"50\"}]'),(89,'[{\"condition\":\"1\",\"val\":\"1000\",\"skill_id\":\"42\",\"power\":\"50\"}]'),(96,'[{\"condition\":\"1\",\"val\":\"6000\",\"skill_id\":\"26\",\"power\":\"70\"},{\"condition\":\"1\",\"val\":\"4000\",\"skill_id\":\"9\",\"power\":\"70\"},{\"condition\":\"1\",\"val\":\"2000\",\"skill_id\":\"26\",\"power\":\"70\"}]'),(100,'[{\"condition\":\"1\",\"val\":\"35000\",\"skill_id\":\"40\",\"power\":\"2000\"},{\"condition\":\"1\",\"val\":\"20000\",\"skill_id\":\"40\",\"power\":\"2000\"},{\"condition\":\"1\",\"val\":\"30000\",\"skill_id\":\"108\",\"power\":\"10000\"},{\"condition\":\"1\",\"val\":\"10000\",\"skill_id\":\"40\",\"power\":\"2000\"}]'),(101,'[{\"condition\":\"1\",\"val\":\"10000\",\"skill_id\":\"108\",\"power\":\"10000\"},{\"condition\":\"1\",\"val\":\"5000\",\"skill_id\":\"40\",\"power\":\"380\"}]'),(104,'[{\"condition\":\"1\",\"val\":\"20000\",\"skill_id\":\"45\",\"power\":\"600\"},{\"condition\":\"1\",\"val\":\"15000\",\"skill_id\":\"22\",\"power\":\"600\"},{\"condition\":\"1\",\"val\":\"10000\",\"skill_id\":\"108\",\"power\":\"5000\"},{\"condition\":\"1\",\"val\":\"5000\",\"skill_id\":\"45\",\"power\":\"600\"}]'),(105,'[{\"condition\":\"1\",\"val\":\"25000\",\"skill_id\":\"33\",\"power\":\"1200\"},{\"condition\":\"1\",\"val\":\"20000\",\"skill_id\":\"32\",\"power\":\"1200\"},{\"condition\":\"1\",\"val\":\"15000\",\"skill_id\":\"34\",\"power\":\"1200\"},{\"condition\":\"1\",\"val\":\"7000\",\"skill_id\":\"34\",\"power\":\"1200\"}]'),(106,'[{\"condition\":\"1\",\"val\":\"30000\",\"skill_id\":\"16\",\"power\":\"1600\"},{\"condition\":\"1\",\"val\":\"23000\",\"skill_id\":\"28\",\"power\":\"1600\"},{\"condition\":\"1\",\"val\":\"14000\",\"skill_id\":\"18\",\"power\":\"1600\"},{\"condition\":\"1\",\"val\":\"10000\",\"skill_id\":\"30\",\"power\":\"1600\"}]'),(107,'[{\"condition\":\"1\",\"val\":\"40000\",\"skill_id\":\"40\",\"power\":\"2000\"},{\"condition\":\"1\",\"val\":\"30000\",\"skill_id\":\"40\",\"power\":\"2000\"},{\"condition\":\"1\",\"val\":\"20000\",\"skill_id\":\"40\",\"power\":\"2000\"},{\"condition\":\"1\",\"val\":\"10000\",\"skill_id\":\"40\",\"power\":\"2000\"}]'),(110,'[{\"condition\":\"1\",\"val\":\"60000\",\"skill_id\":\"45\",\"power\":\"1000\"},{\"condition\":\"1\",\"val\":\"40000\",\"skill_id\":\"22\",\"power\":\"1000\"},{\"condition\":\"1\",\"val\":\"20000\",\"skill_id\":\"108\",\"power\":\"20000\"},{\"condition\":\"1\",\"val\":\"10000\",\"skill_id\":\"45\",\"power\":\"1000\"}]'),(111,'[{\"condition\":\"1\",\"val\":\"150\",\"skill_id\":\"108\",\"power\":\"200\"}]'),(112,'[{\"condition\":\"1\",\"val\":\"1500\",\"skill_id\":\"39\",\"power\":\"50\"},{\"condition\":\"1\",\"val\":\"1000\",\"skill_id\":\"39\",\"power\":\"50\"},{\"condition\":\"1\",\"val\":\"500\",\"skill_id\":\"39\",\"power\":\"50\"}]'),(115,'[{\"condition\":\"1\",\"val\":\"5000\",\"skill_id\":\"26\",\"power\":\"100\"},{\"condition\":\"1\",\"val\":\"4000\",\"skill_id\":\"45\",\"power\":\"100\"},{\"condition\":\"1\",\"val\":\"2000\",\"skill_id\":\"45\",\"power\":\"100\"},{\"condition\":\"1\",\"val\":\"1000\",\"skill_id\":\"22\",\"power\":\"100\"}]');
/*!40000 ALTER TABLE `enemy_boss_script` ENABLE KEYS */;
DROP TABLE IF EXISTS `enemy_deploy_form`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `enemy_deploy_form` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL COMMENT '关联此阵法的某表唯一ID',
  `battle_type` tinyint(4) NOT NULL COMMENT '战场类型(0-区域关卡;1-资源关卡;2-极限关卡;3-多人关卡;8-难度关卡;9-伙伴关卡;10-灵宠关卡;11-魂侍关卡)',
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
) ENGINE=InnoDB AUTO_INCREMENT=78 DEFAULT CHARSET=utf8mb4 COMMENT='怪物阵法表单';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `enemy_deploy_form` DISABLE KEYS */;
INSERT INTO `enemy_deploy_form` VALUES (1,93,0,0,0,0,0,0,0,0,111,0,0,0,0,0,0,0),(7,36,0,0,0,0,115,0,0,0,0,0,0,0,0,0,0,0),(9,27,0,0,0,0,0,0,0,0,112,0,0,0,0,0,0,0),(10,45,0,0,0,39,0,39,0,0,0,11,0,0,0,0,0,0),(11,54,0,0,0,13,0,13,0,0,0,14,0,0,0,0,0,0),(12,87,0,0,21,21,21,0,0,0,101,0,0,0,0,0,0,0),(13,63,0,0,0,0,0,16,0,0,0,18,0,0,0,16,0,0),(14,72,0,0,0,0,15,15,0,0,0,19,0,0,0,0,0,0),(15,90,0,0,22,0,22,0,0,0,23,0,0,0,22,0,22,0),(17,1,3,0,0,0,0,0,0,0,26,0,0,0,0,0,0,0),(18,2,3,0,0,8,0,0,0,0,0,0,0,0,0,0,0,0),(19,3,3,0,0,0,25,0,0,0,0,0,0,0,0,0,0,0),(20,4,3,0,0,11,0,0,0,0,0,0,0,0,0,0,0,0),(21,5,3,0,0,0,14,0,0,0,0,0,0,0,0,0,0,0),(28,170,1,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0),(30,172,1,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0),(32,176,1,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0),(33,174,1,0,0,0,0,0,0,0,26,0,0,0,88,0,88,0),(34,178,1,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0),(36,179,1,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0),(37,180,1,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0),(39,182,1,0,0,0,0,0,0,0,25,0,0,0,88,0,88,0),(41,184,1,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0),(43,186,1,0,0,0,0,0,0,0,11,0,0,0,88,0,88,0),(44,188,1,0,0,0,0,0,0,0,11,0,0,0,88,0,88,0),(45,190,1,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0),(46,192,1,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0),(47,194,1,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0),(48,195,1,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0),(49,197,1,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0),(50,199,1,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0),(51,6,3,0,0,18,0,0,0,0,0,0,0,0,0,0,0,0),(52,7,3,0,0,0,19,0,0,0,0,0,0,0,0,0,0,0),(53,8,3,0,0,0,0,0,0,0,23,0,0,0,0,0,0,0),(54,201,8,0,0,0,0,0,0,0,96,0,0,0,0,0,0,0),(55,202,8,0,0,0,104,0,0,0,0,0,0,0,0,0,0,0),(56,203,8,0,0,0,0,0,0,0,105,0,0,0,0,0,0,0),(57,204,8,0,0,0,0,0,0,0,0,106,0,0,0,0,0,0),(58,205,8,0,0,0,0,0,0,0,100,0,0,0,0,0,0,0),(60,11,0,0,31,0,31,0,0,0,0,0,0,0,0,0,0,0),(61,10,0,0,0,31,0,0,0,0,31,0,0,0,0,0,0,0),(63,213,9,0,0,0,0,0,0,0,107,0,0,0,0,0,0,0),(64,216,10,0,108,0,108,0,0,0,11,0,0,0,0,0,0,0),(65,218,11,0,0,109,110,109,0,0,0,0,0,0,0,0,0,0),(66,221,9,0,0,0,0,0,0,0,107,0,0,0,0,0,0,0),(67,224,9,0,0,0,0,0,0,0,107,0,0,0,0,0,0,0),(68,227,9,0,0,0,0,0,0,0,107,0,0,0,0,0,0,0),(69,230,10,0,108,0,108,0,0,0,11,0,0,0,0,0,0,0),(70,233,10,0,108,0,108,0,0,0,11,0,0,0,0,0,0,0),(71,236,10,0,108,0,108,0,0,0,11,0,0,0,0,0,0,0),(72,238,11,0,0,109,110,109,0,0,0,0,0,0,0,0,0,0),(73,240,11,0,0,109,110,109,0,0,0,0,0,0,0,0,0,0),(74,242,11,0,0,109,110,109,0,0,0,0,0,0,0,0,0,0),(75,244,0,0,0,111,0,0,0,0,0,0,0,0,0,0,0,0),(76,245,0,0,0,0,0,0,0,0,113,0,0,0,0,0,0,0),(77,246,0,0,0,0,0,0,0,0,114,0,0,0,0,0,0,0);
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
  `skill_id` smallint(6) DEFAULT '0' COMMENT '绝招ID',
  `skill_force` int(11) DEFAULT '0' COMMENT '绝招威力',
  `skill2_id` smallint(6) DEFAULT '0' COMMENT '绝招2 ID',
  `skill2_force` int(11) DEFAULT '0' COMMENT '绝招2 威力',
  `sunder_max_value` int(11) NOT NULL COMMENT '护甲值',
  `sunder_min_hurt_rate` int(11) NOT NULL COMMENT '破甲前起始的伤害转换率（百分比）',
  `sunder_end_hurt_rate` int(11) NOT NULL COMMENT '破甲后的伤害转换率（百分比）',
  `sunder_attack` int(11) NOT NULL COMMENT '攻击破甲值',
  `skill_wait` tinyint(4) DEFAULT '0' COMMENT '绝招蓄力回合',
  `release_num` tinyint(4) DEFAULT '0' COMMENT '释放次数',
  `recover_round_num` tinyint(4) DEFAULT '0' COMMENT '恢复回合数',
  `common_attack_num` tinyint(4) DEFAULT '0' COMMENT '入场普通攻击次数',
  `jump_attack` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否跳跃攻击',
  `body_size` tinyint(4) NOT NULL DEFAULT '1' COMMENT '怪物体型',
  `is_boss` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否是boss 0否,1是',
  `scale_size` smallint(5) NOT NULL DEFAULT '100' COMMENT '怪物缩放比%',
  `skill_config` text COMMENT '技能配置',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=120 DEFAULT CHARSET=utf8mb4 COMMENT='敌人角色数据';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `enemy_role` DISABLE KEYS */;
INSERT INTO `enemy_role` VALUES (1,'疯狂草妖','CaoYao',10,2926,0,0,350,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,1,0,144,NULL),(2,'二货兔','ErHuoTu',4,300,0,0,80,0,0,0,0,0,0,0,0,0,-1,20,0,0,100,100,200,5,0,1,2,1,1,2,1,180,'[{\"rand\":\"50\",\"skill_id\":\"35\",\"power\":\"10\"}]'),(3,'疯狂竹筒精','ZhuTongJing',2,140,0,0,100,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,1,0,144,''),(4,'林精','LinJing',2,140,0,0,100,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,1,0,144,NULL),(5,'疯狂黑狼','HeiLang',6,2000,0,0,300,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,1,0,144,''),(6,'疯狂鬼火','GuiHuo',6,1000,0,0,220,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,1,0,144,NULL),(7,'疯狂灯笼怪','DengLongGuai',6,1000,0,0,220,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,1,0,144,NULL),(8,'天狼妖','TianLangYao',7,4800,0,0,350,0,0,0,0,0,0,0,0,0,10,80,0,0,100,100,200,5,0,1,2,0,1,3,1,144,'[{\"rand\":\"30\",\"skill_id\":\"10\",\"power\":\"70\"},{\"rand\":\"10\",\"skill_id\":\"19\",\"power\":\"70\"},{\"rand\":\"20\",\"skill_id\":\"44\",\"power\":\"70\"}]'),(9,'疯狂莲藕精','LianOuJing',10,2926,0,0,400,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,1,0,144,NULL),(10,'愤怒林精','LinJing',10,2926,0,0,350,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,1,0,144,NULL),(11,'金蟾王','JingChanWang',11,10000,0,0,360,0,0,0,0,0,0,0,0,0,32,100,0,0,100,100,200,5,0,1,2,1,1,3,1,144,'[{\"rand\":\"20\",\"skill_id\":\"36\",\"power\":\"110\"},{\"rand\":\"20\",\"skill_id\":\"44\",\"power\":\"110\"},{\"rand\":\"20\",\"skill_id\":\"43\",\"power\":\"110\"}]'),(12,'测试咆哮厉爪','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,150,1,0,1,2,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"9\",\"power\":\"10\"}]'),(13,'疯狂毒蛇','DuShe',12,2000,0,0,420,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,1,0,144,NULL),(14,'剧毒臭泥','JuDuChouNi',13,15000,0,0,450,0,0,0,0,0,0,0,0,0,32,150,0,0,100,100,200,5,0,1,2,1,0,3,1,144,'[{\"rand\":\"10\",\"skill_id\":\"23\",\"power\":\"160\"},{\"rand\":\"15\",\"skill_id\":\"9\",\"power\":\"160\"},{\"rand\":\"25\",\"skill_id\":\"36\",\"power\":\"160\"}]'),(15,'疯狂火蝎','HuoXie',15,5000,0,0,660,200,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,1,0,120,NULL),(16,'疯狂燃魁','RanKui',15,5000,0,0,660,200,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,2,0,120,NULL),(17,'疯狂熔岩虫','RongYanChong',16,5000,0,0,660,200,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,2,0,130,NULL),(18,'燃魁首领','RanKui',17,24000,0,0,690,200,0,0,0,0,0,0,0,0,15,200,0,0,150,100,200,5,0,1,2,1,0,3,1,180,'[{\"rand\":\"20\",\"skill_id\":\"41\",\"power\":\"250\"},{\"rand\":\"25\",\"skill_id\":\"15\",\"power\":\"250\"},{\"rand\":\"15\",\"skill_id\":\"36\",\"power\":\"250\"}]'),(19,'炎龙','YanLong',17,27000,0,0,810,200,0,0,0,0,0,0,0,0,17,250,0,0,150,100,200,5,0,1,2,1,0,3,1,144,'[{\"rand\":\"20\",\"skill_id\":\"27\",\"power\":\"310\"},{\"rand\":\"15\",\"skill_id\":\"28\",\"power\":\"310\"},{\"rand\":\"15\",\"skill_id\":\"44\",\"power\":\"310\"},{\"rand\":\"10\",\"skill_id\":\"43\",\"power\":\"310\"}]'),(20,'疯狂吸血蝙蝠','HeYiJuFu',19,5500,0,0,810,300,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,2,0,144,NULL),(21,'疯狂败亡之剑','BaiWangZhiJian',19,5500,0,0,810,300,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,1,0,144,NULL),(22,'疯狂剑之守卫','JianZhiShouWei',19,5500,0,0,810,300,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,1,0,120,NULL),(23,'古代剑灵','GuDaiJianLing',20,32000,0,0,1000,300,0,0,0,0,0,0,0,0,18,350,0,0,200,100,200,5,0,1,2,0,1,2,1,144,'[{\"rand\":\"30\",\"skill_id\":\"39\",\"power\":\"330\"},{\"rand\":\"30\",\"skill_id\":\"38\",\"power\":\"330\"},{\"rand\":\"30\",\"skill_id\":\"42\",\"power\":\"330\"}]'),(24,'阴影','YinYing',6,1000,0,0,220,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,2,0,144,''),(25,'妖龙','YaoLong',8,6000,0,0,450,0,0,0,0,0,0,0,0,0,9,90,0,0,50,100,200,5,0,1,2,0,1,3,1,144,'[{\"rand\":\"8\",\"skill_id\":\"20\",\"power\":\"90\"},{\"rand\":\"10\",\"skill_id\":\"25\",\"power\":\"90\"},{\"rand\":\"8\",\"skill_id\":\"21\",\"power\":\"90\"},{\"rand\":\"10\",\"skill_id\":\"24\",\"power\":\"90\"}]'),(26,'刀疤兔','DaoBaTu',5,1000,0,0,80,0,0,50,0,0,0,0,0,0,43,30,0,0,100,100,200,5,0,1,2,1,1,2,1,144,'[{\"rand\":\"20\",\"skill_id\":\"35\",\"power\":\"10\"},{\"rand\":\"30\",\"skill_id\":\"43\",\"power\":\"10\"},{\"rand\":\"10\",\"skill_id\":\"10\",\"power\":\"10\"}]'),(27,'疯狂黑翼巨蝠','HeYiJuFu',12,2000,0,0,420,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,2,0,144,NULL),(31,'竹筒精','ZhuTongJing',1,50,0,0,80,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,1,0,144,NULL),(33,'神秘鬼火','GuiHuo',25,5000,0,0,1000,500,0,0,0,0,0,0,0,0,-1,0,0,0,100,100,200,5,0,1,2,0,1,1,0,144,NULL),(34,'黑狼','HeiLang',5,400,0,0,150,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,1,0,144,''),(35,'鬼火','GuiHuo',5,400,0,0,150,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,1,0,144,NULL),(36,'暴走灯笼怪','DengLongGuai',5,400,0,0,150,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,1,0,144,NULL),(37,'阴影2','YinYing',5,912,0,0,236,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,2,0,144,NULL),(38,'暴走林精','LinJing',9,1300,0,0,300,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,1,0,144,NULL),(39,'呆呆莲藕精','LianOuJing',9,1300,0,0,300,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,1,0,144,''),(40,'草妖','CaoYao',9,1300,0,0,300,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,1,0,144,NULL),(41,'黑翼巨蝠','HeYiJuFu',11,1500,0,0,320,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,2,0,144,NULL),(42,'毒蛇','DuShe',11,1500,0,0,320,0,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,1,0,144,NULL),(43,'火蝎','HuoXie',14,2700,0,0,580,100,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,1,0,120,NULL),(44,'燃魁','RanKui',14,2700,0,0,580,100,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,2,0,120,NULL),(45,'熔岩虫','RongYanChong',14,2700,0,0,580,100,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,2,0,130,NULL),(46,'败亡之剑','BaiWangZhiJian',18,3000,0,0,740,200,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,1,0,144,NULL),(47,'吸血蝙蝠','HeYiJuFu',18,3000,0,0,740,200,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,2,0,144,NULL),(48,'剑之守卫','JianZhiShouWei',18,3000,0,0,740,200,0,0,0,0,0,0,0,0,-1,0,0,0,50,100,200,5,0,1,2,0,1,1,0,120,NULL),(49,'测试凶猛撕咬','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"10\",\"power\":\"100\"}]'),(50,'测试冰烈','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"11\",\"power\":\"100\"}]'),(51,'测试火烈','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"15\",\"power\":\"100\"}]'),(52,'测试风烈','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"19\",\"power\":\"100\"}]'),(53,'测试雷烈','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"23\",\"power\":\"100\"}]'),(54,'测试土烈','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"27\",\"power\":\"100\"}]'),(55,'测试毒烈','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"31\",\"power\":\"100\"}]'),(56,'测试多连斩','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"35\",\"power\":\"100\"}]'),(57,'测试力劈华山','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"36\",\"power\":\"100\"}]'),(58,'测试死亡标记','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"41\",\"power\":\"100\"}]'),(59,'测试万箭穿心','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"42\",\"power\":\"100\"}]'),(60,'测试狮吼功','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"43\",\"power\":\"100\"}]'),(61,'测试驱散','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"46\",\"power\":\"100\"}]'),(62,'测试增益','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"109\",\"power\":\"100\"}]'),(63,'测试白莲横江','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"37\",\"power\":\"100\"}]'),(64,'测试横扫千军','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"38\",\"power\":\"100\"}]'),(65,'测试乾坤刀气','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"39\",\"power\":\"100\"}]'),(66,'测试野蛮冲撞','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"44\",\"power\":\"100\"}]'),(67,'测试治疗','GuDaiJianLing',99,1000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"108\",\"power\":\"100\"}]'),(68,'测试三千洛水剑','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"40\",\"power\":\"100\"}]'),(69,'测试如殂随行','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"0100\",\"skill_id\":\"45\",\"power\":\"100\"}]'),(70,'测试冰烈横向','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"12\",\"power\":\"100\"}]'),(71,'测试火烈横向','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"16\",\"power\":\"100\"}]'),(72,'测试风烈横向','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"20\",\"power\":\"100\"}]'),(73,'测试雷烈横向','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"24\",\"power\":\"100\"}]'),(74,'测试土烈横向','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,2,0,144,'[{\"rand\":\"100\",\"skill_id\":\"28\",\"power\":\"100\"}]'),(75,'测试毒烈横向','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"32\",\"power\":\"100\"}]'),(76,'测试冰烈纵向','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"13\",\"power\":\"100\"}]'),(77,'测试火烈纵向','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"17\",\"power\":\"100\"}]'),(78,'测试风烈纵向','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"21\",\"power\":\"100\"}]'),(79,'测试雷烈纵向','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"25\",\"power\":\"100\"}]'),(80,'测试土烈纵向','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"29\",\"power\":\"100\"}]'),(81,'测试毒烈纵向','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"33\",\"power\":\"100\"}]'),(82,'测试冰烈全体','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"14\",\"power\":\"100\"}]'),(83,'测试火烈全体','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"18\",\"power\":\"100\"}]'),(84,'测试风烈全体','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"22\",\"power\":\"100\"}]'),(85,'测试雷烈全体','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"26\",\"power\":\"100\"}]'),(86,'测试土烈全体','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"30\",\"power\":\"100\"}]'),(87,'测试毒烈全体','GuDaiJianLing',99,100,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"34\",\"power\":\"100\"}]'),(88,'铜钱怪','TongQianGuai',20,3000,0,0,600,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,1,1,0,144,NULL),(89,'经验怪','JingYanGuai',20,3000,0,0,600,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,1,1,0,144,NULL),(90,'测试破甲','GuDaiJianLing',99,10000,100,1000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,50,0,0,0,0,0,1,0,144,''),(91,'莲藕精','LianOuJing',10,1000,0,0,300,400,0,0,0,0,0,0,0,0,0,0,0,0,200,100,200,5,0,0,0,0,0,1,0,120,'[{\"rand\":\"100\",\"skill_id\":\"108\",\"power\":\"1000\"}]'),(92,'火灵','RanKui',20,3000,0,0,600,600,0,0,0,0,0,0,0,0,0,0,0,0,200,100,200,5,0,0,0,0,1,1,0,100,'[{\"rand\":\"100\",\"skill_id\":\"36\",\"power\":\"500\"}]'),(93,'剑魄','BaiWangZhiJian',25,5000,0,100,800,500,0,0,0,0,0,0,0,0,0,0,0,0,200,100,200,5,0,0,0,0,0,1,0,120,'[{\"rand\":\"100\",\"skill_id\":\"35\",\"power\":\"500\"}]'),(94,'测试中毒','GuDaiJianLing',99,10000,100,100,500,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"34\",\"power\":\"0\"}]'),(95,'勇气之卫','YongQiZhiWei',99,1000,1,1,100,1,1,1,1,1,1,1,1,1,0,0,0,0,50,100,150,1,0,0,0,0,0,3,0,100,NULL),(96,'噩梦刀疤兔','DaoBaTu',8,8000,0,0,500,0,0,0,0,0,0,0,0,0,0,0,0,0,300,100,150,20,0,0,0,0,1,2,1,144,'[{\"rand\":\"10\",\"skill_id\":\"35\",\"power\":\"50\"},{\"rand\":\"30\",\"skill_id\":\"43\",\"power\":\"50\"},{\"rand\":\"20\",\"skill_id\":\"1035\",\"power\":\"50\"}]'),(97,'盗贼','DaoZei',3,140,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,200,5,0,0,0,0,1,1,0,144,NULL),(98,'测试魂力','GuDaiJianLing',99,10000,50,50,50,0,0,0,0,0,0,0,0,0,0,0,0,0,100,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"19\",\"power\":\"0\"}]'),(100,'噩梦古代剑灵','GuDaiJianLing',24,40000,0,0,1300,400,0,0,0,0,0,0,0,0,0,0,0,0,300,100,150,20,0,0,0,0,1,2,1,144,'[{\"rand\":\"30\",\"skill_id\":\"39\",\"power\":\"1000\"},{\"rand\":\"30\",\"skill_id\":\"38\",\"power\":\"1000\"},{\"rand\":\"30\",\"skill_id\":\"42\",\"power\":\"1000\"}]'),(101,'守剑之魂','JianZhiShouWei',19,28000,0,0,900,300,0,0,0,0,0,0,0,0,0,0,0,0,150,100,200,5,0,0,0,0,1,2,1,140,'[{\"rand\":\"10\",\"skill_id\":\"39\",\"power\":\"320\"},{\"rand\":\"10\",\"skill_id\":\"38\",\"power\":\"320\"},{\"rand\":\"30\",\"skill_id\":\"35\",\"power\":\"320\"}]'),(102,'测试死亡阻击','GuDaiJianLing',99,100,100,10000,100,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,0,144,'[{\"rand\":\"100\",\"skill_id\":\"1035\",\"power\":\"10\"}]'),(103,'刀疤兔小弟','ErHuoTu',4,140,0,0,80,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,200,5,0,0,0,0,1,1,0,120,'[{\"rand\":\"30\",\"skill_id\":\"35\",\"power\":\"10\"}]'),(104,'噩梦妖龙','YaoLong',12,25000,0,0,700,0,0,0,0,0,0,0,0,0,0,0,0,0,300,100,150,20,0,0,0,0,0,3,1,144,'[{\"rand\":\"30\",\"skill_id\":\"9\",\"power\":\"300\"},{\"rand\":\"10\",\"skill_id\":\"24\",\"power\":\"300\"},{\"rand\":\"10\",\"skill_id\":\"25\",\"power\":\"300\"}]'),(105,'噩梦剧毒臭泥','JuDuChouNi',16,30000,0,0,900,0,0,0,0,0,0,0,0,0,0,0,0,0,300,100,150,20,0,0,0,0,0,3,1,144,'[{\"rand\":\"30\",\"skill_id\":\"27\",\"power\":\"600\"},{\"rand\":\"10\",\"skill_id\":\"38\",\"power\":\"600\"},{\"rand\":\"20\",\"skill_id\":\"36\",\"power\":\"600\"}]'),(106,'噩梦炎龙','YanLong',20,35000,0,0,1100,300,0,0,0,0,0,0,0,0,0,0,0,0,300,100,150,20,0,0,0,0,0,3,1,144,'[{\"rand\":\"20\",\"skill_id\":\"41\",\"power\":\"800\"},{\"rand\":\"20\",\"skill_id\":\"1035\",\"power\":\"800\"},{\"rand\":\"10\",\"skill_id\":\"37\",\"power\":\"800\"}]'),(107,'狂暴古代剑灵','GuDaiJianLing',25,50000,0,0,1500,500,0,0,0,0,0,0,0,0,0,0,0,0,100,100,200,5,0,0,0,0,0,3,1,144,'[{\"rand\":\"30\",\"skill_id\":\"39\",\"power\":\"1200\"},{\"rand\":\"30\",\"skill_id\":\"38\",\"power\":\"1200\"},{\"rand\":\"30\",\"skill_id\":\"42\",\"power\":\"1200\"}]'),(108,'迷茫草精','CaoYao',25,1500,0,0,800,0,0,0,0,0,0,0,0,0,0,0,0,0,50,100,200,5,0,0,0,0,1,1,0,144,'[{\"rand\":\"20\",\"skill_id\":\"108\",\"power\":\"1000\"}]'),(109,'神秘灯笼怪','DengLongGuai',25,8000,0,0,1000,500,0,0,0,0,0,0,0,0,0,0,0,0,100,100,200,5,0,0,0,0,1,1,0,144,NULL),(110,'狂暴妖龙','YaoLong',25,80000,0,0,1500,500,0,0,0,0,0,0,0,0,0,0,0,0,200,100,200,5,0,0,0,0,0,3,1,144,'[{\"rand\":\"30\",\"skill_id\":\"9\",\"power\":\"500\"},{\"rand\":\"10\",\"skill_id\":\"24\",\"power\":\"500\"},{\"rand\":\"10\",\"skill_id\":\"25\",\"power\":\"500\"}]'),(111,'朱媛媛','ZhuYuanYuan',2,300,0,0,80,0,0,0,0,0,0,0,0,0,0,0,0,0,100,100,200,5,0,0,0,0,1,1,1,130,'[{\"rand\":\"70\",\"skill_id\":\"19\",\"power\":\"10\"}]'),(112,'袁铭志','YuanMingZhi',5,2000,0,0,190,0,0,0,0,0,0,0,0,0,0,0,0,0,100,100,200,5,0,0,0,0,1,1,1,130,'[{\"rand\":\"60\",\"skill_id\":\"35\",\"power\":\"30\"}]'),(113,'影龙姬','YingLongJi',50,100000,0,0,10000,0,0,0,0,0,0,0,0,0,0,0,0,0,200,100,200,1,0,0,0,0,0,1,1,100,NULL),(114,'奸奇','JianQi',50,1000,0,0,1000,1,0,0,0,0,0,0,0,0,0,0,0,0,50,100,150,1,0,0,0,0,0,1,1,100,NULL),(115,'影龙','YingLong',8,6000,0,0,300,0,0,0,0,0,0,0,0,0,0,0,0,0,100,100,200,5,0,0,0,0,0,3,1,144,'[{\"rand\":\"8\",\"skill_id\":\"20\",\"power\":\"70\"},{\"rand\":\"8\",\"skill_id\":\"21\",\"power\":\"70\"},{\"rand\":\"10\",\"skill_id\":\"24\",\"power\":\"70\"},{\"rand\":\"10\",\"skill_id\":\"25\",\"power\":\"70\"}]'),(116,'灯笼怪','DengLongGuai',25,5000,0,800,1200,500,0,0,0,0,0,0,0,0,0,0,0,0,200,100,200,5,0,0,0,0,1,1,0,120,'[{\"rand\":\"100\",\"skill_id\":\"23\",\"power\":\"500\"}]'),(117,'画妖','HuaYao',35,8000,0,1500,2200,600,0,50,0,0,0,0,0,0,0,0,0,0,200,100,150,5,0,0,0,0,1,1,0,120,'[{\"rand\":\"100\",\"skill_id\":\"21\",\"power\":\"1000\"}]'),(118,'魔笔','MoBi',35,10000,0,1800,1800,800,0,50,0,0,0,0,0,0,0,0,0,0,200,100,150,5,0,0,0,0,1,1,0,120,'[{\"rand\":\"100\",\"skill_id\":\"38\",\"power\":\"1000\"}]'),(119,'梦妖','MengYao',40,10000,0,2000,2000,1000,0,50,0,0,0,0,0,0,0,0,0,0,200,100,200,5,0,0,0,0,0,1,0,120,'[{\"rand\":\"100\",\"skill_id\":\"109\",\"power\":\"1000\"}]');
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
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='装备追加属性表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `equipment_appendix` DISABLE KEYS */;
INSERT INTO `equipment_appendix` VALUES (1,1,100,20,20,20,10,5,5,5,5,5,5),(2,2,250,100,100,100,50,20,20,20,20,20,20),(3,3,1000,400,400,400,200,40,40,40,40,40,40),(4,4,2000,800,800,800,400,80,80,80,80,80,80),(5,5,3000,1200,1200,1200,600,120,120,120,120,120,120),(6,6,4000,1600,1600,1600,800,160,160,160,160,160,160),(7,7,5000,2000,2000,2000,1000,200,200,200,200,200,200),(8,-227,5000,2000,2000,2000,1000,0,0,0,0,0,0);
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
INSERT INTO `equipment_recast` VALUES (1,1,1,1,1,0,0,0,10000),(2,20,1,2,2,0,0,0,20000),(3,40,1,4,4,0,0,0,40000),(4,60,1,6,6,0,0,0,60000),(5,80,1,8,8,0,0,0,80000),(6,100,1,10,10,0,0,0,100000),(7,1,2,1,1,1,0,0,20000),(8,20,2,2,2,2,0,0,40000),(9,40,2,4,4,4,0,0,80000),(10,60,2,6,6,6,0,0,120000),(11,80,2,8,8,8,0,0,160000),(12,100,2,10,10,10,0,0,200000),(13,1,3,1,1,1,1,0,40000),(14,20,3,2,2,2,2,0,80000),(15,40,3,4,4,4,4,0,160000),(16,60,3,6,6,6,6,0,240000),(17,80,3,8,8,8,8,0,320000),(18,100,3,10,10,10,10,0,400000),(19,1,4,1,1,1,1,1,80000),(20,20,4,2,2,2,2,2,160000),(21,40,4,4,4,4,4,4,320000),(22,60,4,6,6,6,6,6,480000),(23,80,4,8,8,8,8,8,640000),(24,100,4,10,10,10,10,10,800000);
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
INSERT INTO `equipment_refine_level` VALUES (1,1,1,100,50),(2,2,1,70,100),(3,3,1,50,150),(4,4,1,30,250),(5,5,1,10,400),(6,1,2,90,50),(7,2,2,70,100),(8,3,2,50,150),(9,4,2,30,250),(10,5,2,10,400),(11,1,3,90,50),(12,2,3,70,100),(13,3,3,50,150),(14,4,3,30,250),(15,5,3,10,400),(16,1,4,90,50),(17,2,4,70,100),(18,3,4,50,150),(19,4,4,30,250),(20,5,4,10,400);
/*!40000 ALTER TABLE `equipment_refine_level` ENABLE KEYS */;
DROP TABLE IF EXISTS `extend_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `extend_level` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `max_level` smallint(6) NOT NULL COMMENT '允许开放的等级上限',
  `level_type` tinyint(4) NOT NULL COMMENT '关卡类型(1-资源关卡;9-伙伴关卡;10-灵宠关卡;11-魂侍关卡)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COMMENT='活动关卡';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `extend_level` DISABLE KEYS */;
INSERT INTO `extend_level` VALUES (2,15,1),(3,30,1),(5,50,1),(7,70,1),(8,20,9),(9,30,9),(10,50,9),(11,70,9),(12,20,10),(13,20,11),(14,30,11),(15,50,11),(16,70,11),(17,30,10),(18,50,10),(19,70,10);
/*!40000 ALTER TABLE `extend_level` ENABLE KEYS */;
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
  `need_play` tinyint(4) DEFAULT '0' COMMENT '是否需要播放 0不需要 1需要',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COMMENT='功能权值配置';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `func` DISABLE KEYS */;
INSERT INTO `func` VALUES (3,'魂侍','FUNC_GHOST',1000,0,1,1),(4,'资源关卡','FUNC_RESOURCE_LEVEL',2000,0,2,0),(5,'通天塔','FUNC_TOWER_LEVEL',3000,0,4,0),(6,'多人关卡','FUNC_MULTI_LEVEL',4000,0,8,0),(7,'剑心','FUNC_SWORD_SOUL',1500,0,16,1),(8,'灵宠','FUNC_BATTLE_PET',1300,0,32,1),(9,'神龙宝箱','FUNC_CHEST',1600,0,64,0),(10,'比武场','FUNC_ARENA',900,0,128,0),(11,'侠之试炼','FUNC_XIA_ZHI_SHI_LIAN',6000,0,256,0);
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
  `fragment_id` smallint(6) NOT NULL COMMENT '对应碎片物品id',
  `unique_key` smallint(6) NOT NULL DEFAULT '0' COMMENT '每个影界中魂侍的唯一标记',
  `init_star` tinyint(4) NOT NULL DEFAULT '0' COMMENT '初始星级',
  `health` int(11) NOT NULL DEFAULT '0' COMMENT '生命',
  `attack` int(11) NOT NULL DEFAULT '0' COMMENT '攻击',
  `defence` int(11) NOT NULL DEFAULT '0' COMMENT '防御',
  `speed` int(11) NOT NULL DEFAULT '0' COMMENT '速度',
  `desc` varchar(300) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COMMENT='魂侍';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `ghost` DISABLE KEYS */;
INSERT INTO `ghost` VALUES (1,'勇气之卫','YongQiZhiWei',1,-1,256,1,1,400,150,30,50,'人型，六臂金刚的造型，天界狂战士。'),(2,'天盾卫士','TianDunWeiShi',1,3,257,2,1,800,270,150,50,'天界持盾卫士。'),(3,'绮梦花妖','QiMengHuaYao',1,4,258,4,1,1000,220,90,50,'花、少女、粉嫩。'),(4,'将臣','JiangChen',2,-1,259,0,1,1000,450,150,50,'一将功成万骨枯。'),(5,'铳斗士','ChongDouShi',2,6,260,0,1,1200,300,90,50,'古代的西部枪客。'),(6,'人鱼公主','RenYuGongZhu',2,5,261,0,1,1600,410,130,50,'大海的精灵。'),(7,'武圣','WuSheng',0,3,273,0,3,2000,400,200,100,'上古圣灵之一，与阴影的战斗永不休止。'),(8,'剑灵','GuDaiJianLing',0,-1,274,0,3,2000,500,150,100,'上古圣灵之一，与阴影的战斗永不休止。'),(9,'飞羽','FeiYu',0,4,275,0,3,2000,800,0,100,'上古圣灵之一，与阴影的战斗永不休止。'),(10,'阿修罗','AXiuLuo',0,-1,276,0,3,4000,1200,400,200,'霸者横栏。'),(11,'洛神','LuoShen',0,5,277,0,3,3000,600,200,100,'女神的化身。'),(12,'木偶戏子','MuOXiZi',0,6,278,0,3,2000,300,300,100,'操纵木偶，捕猎阴影。');
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
INSERT INTO `ghost_passive_skill` VALUES (1,20,'魂侍护盾','HunShiShouHu','当生命少于30%时，自动施放吸收伤害的魂侍护盾，持续3回合。魂侍护盾吸收的伤害等于魂侍生命值的50%，与其他魂侍合计。魂侍护盾每次战斗只能触发1次。'),(2,30,'初始魂力','HunShiFaDongLv','初始魂力5，与其他魂侍合计'),(3,40,'魂侍技能2级','HunShiJiNengErJi','魂侍技能升级为2级'),(4,50,'魂侍护盾2级','HunShiShouHuErJi','当生命少于30%时，自动施放吸收伤害的魂侍护盾，持续3回合。魂侍护盾吸收的伤害等于魂侍生命值的100%，与其他魂侍合计。魂侍护盾每次战斗只能触发1次。'),(5,60,'初始魂力2级','HunShiFaDongLvErJi','初始魂力10，与其他魂侍合计'),(6,70,'魂侍技能3级','HunShiJiNengSanJi','魂侍技能升级为3级');
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
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COMMENT='魂侍进阶表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `ghost_star` DISABLE KEYS */;
INSERT INTO `ghost_star` VALUES (1,1,10,10,0),(2,2,30,30,1),(3,3,50,50,2),(4,4,80,80,3),(5,5,100,100,3),(7,6,0,0,3);
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
INSERT INTO `ghost_tip` VALUES (1,'战斗中，角色每次行动与受伤都会获得魂力\n当魂力满100时，可以召唤魂侍'),(2,'每只魂侍在同个关卡内只能被触发一次');
/*!40000 ALTER TABLE `ghost_tip` ENABLE KEYS */;
DROP TABLE IF EXISTS `ghost_umbra`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ghost_umbra` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `town_id` tinyint(4) NOT NULL DEFAULT '0' COMMENT '城镇id',
  `key` int(11) NOT NULL DEFAULT '0' COMMENT '进入影界需求权值',
  `ghost_probability` smallint(6) NOT NULL DEFAULT '0' COMMENT '抽中魂侍的概率',
  `fragment_probability` smallint(6) NOT NULL DEFAULT '0' COMMENT '抽中碎片的概率',
  `fragment_min_num` smallint(6) NOT NULL COMMENT '碎片随机最小值',
  `fragment_max_num` smallint(6) NOT NULL COMMENT '碎片随机最大值',
  `fruit_probability` smallint(6) NOT NULL DEFAULT '0' COMMENT '抽中果实的概率',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='魂侍影界表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `ghost_umbra` DISABLE KEYS */;
INSERT INTO `ghost_umbra` VALUES (1,1,100000,2,28,1,2,70),(2,2,100110,2,28,1,2,70);
/*!40000 ALTER TABLE `ghost_umbra` ENABLE KEYS */;
DROP TABLE IF EXISTS `ghost_umbra_vip`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ghost_umbra_vip` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `vip_level` tinyint(4) NOT NULL DEFAULT '0' COMMENT 'vip 等级',
  `refresh_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '刷新次数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='魂侍影界VIP表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `ghost_umbra_vip` DISABLE KEYS */;
INSERT INTO `ghost_umbra_vip` VALUES (1,4,1),(2,5,2),(3,6,3),(4,7,4),(5,8,5);
/*!40000 ALTER TABLE `ghost_umbra_vip` ENABLE KEYS */;
DROP TABLE IF EXISTS `ghost_umbra_vip_price`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ghost_umbra_vip_price` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `refresh_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '刷新次数',
  `ingot` bigint(20) NOT NULL DEFAULT '0' COMMENT '元宝',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='魂侍影界VIP价格表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `ghost_umbra_vip_price` DISABLE KEYS */;
INSERT INTO `ghost_umbra_vip_price` VALUES (1,1,20),(2,2,40),(3,3,60),(4,4,80),(5,5,100);
/*!40000 ALTER TABLE `ghost_umbra_vip_price` ENABLE KEYS */;
DROP TABLE IF EXISTS `hard_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `hard_level` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `mission_level_lock` int(11) NOT NULL COMMENT '区域关卡功能权值',
  `desc` varchar(100) NOT NULL COMMENT '关卡描述',
  `town_id` smallint(6) NOT NULL COMMENT '城镇ID',
  `hard_level_lock` int(11) NOT NULL DEFAULT '0' COMMENT '难度关卡权值',
  `award_hard_level_lock` int(11) NOT NULL COMMENT '难度关卡奖励权值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='难度关卡';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `hard_level` DISABLE KEYS */;
INSERT INTO `hard_level` VALUES (1,110100,'刀疤兔的来头果然不简单，而且在深渊中又变的更加强大了。',1,0,100100),(2,120100,'来自异界的妖龙，其实力深不可测，只是它为何要躲躲藏藏。',1,100100,100110),(3,130100,'原来剧毒臭泥是吸收了各种腐化的气息后从地底逃窜到人间的。',1,100110,100120),(4,140100,'虽然炎龙尚未被阴影吞噬，但是在腐烂的地底，炎龙早已不是曾经的炎龙。',1,100120,100130),(5,150100,'上古的剑灵恐怕做梦也不会想到，自己的肉身会被魔物所占据。',1,100130,100140);
/*!40000 ALTER TABLE `hard_level` ENABLE KEYS */;
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
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COMMENT='爱心抽奖奖品配置';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `heart_draw_award` DISABLE KEYS */;
INSERT INTO `heart_draw_award` VALUES (9,1,1,1000,0,20),(10,1,3,1,250,10),(11,1,3,5,263,20),(12,1,3,5,270,12),(13,1,3,1,230,15),(14,1,3,1,212,3),(15,1,3,1,239,15),(16,1,2,20,0,5);
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
  `music_sign` varchar(30) DEFAULT NULL COMMENT '音乐资源',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=312 DEFAULT CHARSET=utf8mb4 COMMENT='物品';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `item` DISABLE KEYS */;
INSERT INTO `item` VALUES (31,3,1,'鹰扬剑',1,NULL,100,'YingYangJian',0,0,0,0,0,0,1,0,0,100,0,-1,0,0,NULL),(32,3,1,'鹰扬刀',1,NULL,100,'YingYangDao',0,0,0,0,0,0,1,0,0,100,0,3,0,0,NULL),(33,3,1,'鹰扬环',1,NULL,100,'YingYangHuan',0,0,0,0,0,0,1,0,0,100,0,4,0,0,NULL),(34,3,1,'鹰扬篮',1,NULL,100,'YingYangLan',0,0,0,0,0,0,1,0,0,100,0,5,0,0,NULL),(35,3,1,'鹰扬笔',1,NULL,100,'YingYangBi',0,0,0,0,0,0,1,0,0,100,0,6,0,0,NULL),(36,4,1,'鹰扬袍',1,NULL,100,'YingYangPao',0,0,0,0,0,0,1,0,0,0,50,0,0,0,NULL),(37,5,1,'鹰扬靴',1,NULL,100,'YingYangXue',0,0,0,0,0,0,1,0,50,0,0,0,0,0,NULL),(38,6,1,'鹰扬戒',1,NULL,100,'YingYangJie',0,0,0,0,0,0,1,200,0,0,0,0,0,0,NULL),(61,3,2,'朱雀剑',20,NULL,500,'ZhuQueJian',0,0,0,0,0,0,2,0,0,400,0,-1,3,2,NULL),(62,3,2,'朱雀刀',20,NULL,500,'ZhuQueDao',0,0,0,0,0,0,2,0,0,400,0,3,3,2,NULL),(63,3,2,'朱雀环',20,NULL,500,'ZhuQueHuan',0,0,0,0,0,0,2,0,0,400,0,4,3,2,NULL),(64,3,2,'朱雀篮',20,NULL,500,'ZhuQueLan',0,0,0,0,0,0,2,0,0,400,0,5,3,2,NULL),(65,3,2,'朱雀笔',20,NULL,500,'ZhuQueBi',0,0,0,0,0,0,2,0,0,400,0,6,3,2,NULL),(66,4,2,'朱雀袍',20,NULL,500,'ZhuQuePao',0,0,0,0,0,0,2,0,0,0,200,0,3,2,NULL),(67,5,2,'朱雀靴',20,NULL,500,'ZhuQueXue',0,0,0,0,0,0,2,0,100,0,0,0,3,2,NULL),(68,6,2,'朱雀戒',20,NULL,500,'ZhuQueJie',0,0,0,0,0,0,2,1000,0,0,0,0,3,2,NULL),(93,3,2,'白虎剑',50,NULL,1000,'BaiHuJian',0,0,0,0,0,0,3,0,0,1600,0,-1,3,3,NULL),(94,3,2,'白虎刀',50,NULL,1000,'BaiHuDao',0,0,0,0,0,0,3,0,0,1600,0,3,3,3,NULL),(95,3,2,'白虎环',50,NULL,1000,'BaiHuHuan',0,0,0,0,0,0,3,0,0,1600,0,4,3,3,NULL),(96,3,2,'白虎篮',50,NULL,1000,'BaiHuLan',0,0,0,0,0,0,3,0,0,1600,0,5,3,3,NULL),(97,3,2,'白虎笔',50,NULL,1000,'BaiHuBi',0,0,0,0,0,0,3,0,0,1600,0,6,3,3,NULL),(98,4,2,'白虎袍',50,NULL,1000,'BaiHuPao',0,0,0,0,0,0,3,0,0,0,800,0,3,3,NULL),(99,5,2,'白虎靴',50,NULL,1000,'BaiHuXue',0,0,0,0,0,0,3,0,400,0,0,0,3,3,NULL),(100,6,2,'白虎戒',50,NULL,1000,'BaiHuJie',0,0,0,0,0,0,3,4000,0,0,0,0,3,3,NULL),(125,3,2,'青龙剑',80,NULL,2000,'QingLongJian',0,0,0,0,0,0,5,0,0,3200,0,-1,3,4,NULL),(126,3,2,'青龙刀',80,NULL,2000,'QingLongDao',0,0,0,0,0,0,5,0,0,3200,0,3,3,4,NULL),(127,3,2,'青龙环',80,NULL,2000,'QingLongHuan',0,0,0,0,0,0,5,0,0,3200,0,4,3,4,NULL),(128,3,2,'青龙篮',80,NULL,2000,'QingLongLan',0,0,0,0,0,0,5,0,0,3200,0,5,3,4,NULL),(129,3,2,'青龙笔',80,NULL,2000,'QingLongBi',0,0,0,0,0,0,5,0,0,3200,0,6,3,4,NULL),(130,4,2,'青龙袍',80,NULL,2000,'QingLongPao',0,0,0,0,0,0,5,0,0,0,1600,0,3,4,NULL),(131,5,2,'青龙靴',80,NULL,2000,'QingLongXue',0,0,0,0,0,0,5,0,800,0,0,0,3,4,NULL),(132,6,2,'青龙戒',80,NULL,2000,'QingLongJie',0,0,0,0,0,0,5,8000,0,0,0,0,3,4,NULL),(157,3,3,'天罡剑',90,NULL,8000,'TianGangJian',0,0,0,0,0,0,6,0,0,8000,0,-1,4,4,NULL),(158,3,3,'天罡刀',90,NULL,8000,'TianGangDao',0,0,0,0,0,0,6,0,0,8000,0,3,4,4,NULL),(159,3,3,'天罡环',90,NULL,8000,'TianGangHuan',0,0,0,0,0,0,6,0,0,8000,0,4,4,4,NULL),(160,3,3,'天罡篮',90,NULL,8000,'TianGangLan',0,0,0,0,0,0,6,0,0,8000,0,5,4,4,NULL),(161,3,3,'天罡笔',90,NULL,8000,'TianGangBi',0,0,0,0,0,0,6,0,0,8000,0,6,4,4,NULL),(162,4,3,'天罡袍',90,NULL,8000,'TianGangPao',0,0,0,0,0,0,6,0,0,0,4000,0,4,4,NULL),(163,5,3,'天罡靴',90,NULL,8000,'TianGangXue',0,0,0,0,0,0,6,0,2000,0,0,0,4,4,NULL),(164,6,3,'天罡戒',90,NULL,8000,'TianGangJie',0,0,0,0,0,0,6,20000,0,0,0,0,4,4,NULL),(209,8,1,'止血草',0,'普通草药，全体恢复5000生命',2000,'ZhiXueCao',1,0,0,0,0,0,0,0,0,0,0,0,0,0,'ZhiLiao'),(210,8,2,'金创药',NULL,'疗伤药品，全体恢复20000生命',15000,'JinChuangYao',1,0,0,0,0,0,0,0,0,0,0,0,0,0,'ZhiLiao'),(211,8,2,'大还丹',NULL,'珍贵丹药，恢复99999生命',8000,'DaHuanDan',1,0,0,0,0,0,0,0,0,0,0,0,0,0,'ZhiLiao'),(212,8,3,'凤凰羽毛',NULL,'从凤凰身上掉落的羽毛。复活并恢复100%生命',0,'FengHuangYuMao',1,0,0,0,0,0,0,0,0,0,0,0,0,0,'FengHuangYuMao'),(213,8,2,'飞刀',NULL,'对生命值最少的敌方造成10000点伤害',3000,'FeiDao',1,0,0,0,0,0,0,0,0,0,0,0,0,0,'FeiDao'),(214,8,2,'暗影飞刀',NULL,'对生命值最少的敌方造成30000点伤害',1000,'AnYingFeiDao',1,0,0,0,0,0,0,0,0,0,0,0,0,0,'FeiDao'),(215,8,2,'风的种子',NULL,'全体闪避等级增加1000，持续2回合',18000,'FengDeZhongZi',1,0,0,0,0,0,0,0,0,0,0,0,0,0,'TianJian'),(216,8,2,'火的种子',NULL,'精气增加10，下回合伤害提升60%',18000,'HuoDeZhongZi',1,0,0,0,0,0,0,0,0,0,0,0,0,0,'HuoDeZhongZi'),(217,8,2,'水的种子',NULL,'清除全体负面状态',18000,'ShuiDeZhongZi',1,0,0,0,0,0,0,0,0,0,0,0,0,0,'HuiChun'),(219,3,3,'倚天剑',60,NULL,4000,'YiTianJian',0,0,0,0,0,0,4,0,0,4000,0,-1,4,3,NULL),(220,3,3,'倚天刀',60,NULL,4000,'YiTianDao',0,0,0,0,0,0,4,0,0,4000,0,3,4,3,NULL),(221,3,3,'倚天环',60,NULL,4000,'YiTianHuan',0,0,0,0,0,0,4,0,0,4000,0,4,4,3,NULL),(222,3,3,'倚天篮',60,NULL,4000,'YiTianLan',0,0,0,0,0,0,4,0,0,4000,0,5,4,3,NULL),(223,3,3,'倚天笔',60,NULL,4000,'YiTianBi',0,0,0,0,0,0,4,0,0,4000,0,6,4,3,NULL),(224,4,3,'倚天袍',60,NULL,4000,'YiTianPao',0,0,0,0,0,0,4,0,0,0,2000,0,4,3,NULL),(225,5,3,'倚天靴',60,NULL,4000,'YiTianXue',0,0,0,0,0,0,4,0,1000,0,0,0,4,3,NULL),(226,6,3,'倚天戒',60,NULL,4000,'YiTianJie',0,0,0,0,0,0,4,10000,0,0,0,0,4,3,NULL),(227,3,4,'黎明破晓',1,NULL,100000,'LiMingPoXiao',0,0,0,0,0,0,7,0,0,2000,0,-1,5,4,NULL),(230,9,1,'1级装备结晶箱',1,'开启后可获得装备结晶',100,'1JiZhuangBeiJieJingXiang',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(231,2,1,'蓝色结晶',1,'分解蓝色品质装备获得，可用于洗炼和重铸装备',10,'LanSeJieJing',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(232,2,2,'紫色结晶',1,'分解紫色品质装备获得，可用于洗炼和重铸装备',10,'ZiSeJieJing',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(233,2,2,'金色结晶',1,'分解金色品质装备获得，可用于洗炼和重铸装备',10,'JinSeJieJing',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(234,2,4,'橙色结晶',1,'分解橙色品质装备获得，可用于洗炼和重铸装备',10,'ChengSeJieJing',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(235,2,2,'武器碎片',1,'分解武器获得，可用于洗炼和重铸武器',0,'WuQiSuiPian',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(236,2,2,'护甲碎片',1,'分解护甲获得，可用于洗炼和重铸护甲',10,'HuJiaSuiPian',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(237,2,2,'长靴碎片',1,'分解长靴获得，可用于洗炼和重铸长靴',10,'ChangXueSuiPian',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(238,2,2,'饰品碎片',1,'分解饰品获得，可用于洗炼和重铸饰品',10,'ShiPinSuiPian',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(239,9,1,'1级装备碎片箱',1,'开启后可获得装备碎片',100,'1JiZhuangBeiSuiPianXiang',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(242,10,3,'元宝',0,'仙侠道内最常用的流通货币。',0,'YuanBao',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(243,10,0,'经验',0,'人物升级所需要的加成。',0,'JingYan',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(244,10,0,'铜钱',0,'仙侠道内最基础的流通货币。',0,'TongQian',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(245,10,1,'经验双倍',0,'获得的经验双倍。',0,'JingYanShuangBei',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(246,10,1,'铜钱双倍',0,'获得的铜钱双倍。',0,'TongQianShuangBei',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(247,10,0,'爱心',0,'好友间友情的证明。',0,'AiXin',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(248,10,1,'体力',0,'挑战关卡时必须消耗的能量。',0,'TiLi',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(250,8,1,'契约球',NULL,'有很大几率能成功捕获灵宠',5000,'QiYueQiu',1,0,0,0,0,0,0,0,0,0,0,0,0,0,'QiYueQiuZhuaBu'),(251,8,3,'至尊契约球',NULL,'一定能成功捕获灵宠',50000,'ZhiZunQiYueQiu',1,0,0,0,0,0,0,0,0,0,0,0,0,0,'QiYueQiuZhuaBu'),(252,11,1,'契约球(莲藕精)',NULL,'与草精达成召唤契约的球',100,'QiYueQiuLianOuJing',0,5,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(253,11,1,'契约球(火灵)',NULL,'与火灵达成召唤契约的球',100,'QiYueQiuHuoLing',0,5,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(254,11,2,'契约球(剑魄)',NULL,'与剑魄达成召唤契约的球',500,'QiYueQiuJianPo',0,5,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(256,13,1,'魂侍碎片(勇气之卫)',NULL,'收集10个魂侍碎片，可以召唤勇气之卫。同时也是勇气之卫进阶的必备材料。',0,'YongQiZhiWei',1,6,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(257,13,1,'魂侍碎片(天盾卫士)',NULL,'收集10个魂侍碎片，可以召唤天盾卫士。同时也是天盾卫士进阶的必备材料。',0,'TianDunWeiShi',1,6,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(258,13,1,'魂侍碎片(绮梦花妖)',NULL,'收集10个魂侍碎片，可以召唤绮梦花妖。同时也是绮梦花妖进阶的必备材料。',0,'QiMengHuaYao',1,6,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(259,13,1,'魂侍碎片(将臣)',NULL,'收集10个魂侍碎片，可以召唤将臣。同时也是将臣进阶的必备材料。',0,'JiangChen',1,6,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(260,13,1,'魂侍碎片(铳斗士)',NULL,'收集10个魂侍碎片，可以召唤铳斗士。同时也是铳斗士进阶的必备材料。',0,'ChongDouShi',1,6,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(261,13,1,'魂侍碎片(人鱼公主)',NULL,'收集10个魂侍碎片，可以召唤人鱼公主。同时也是人鱼公主进阶的必备材料。',0,'RenYuGongZhu',1,6,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(262,12,3,'光明钥匙',NULL,'充满了光明之力的钥匙，可开启被阴影笼罩的区域',0,'GuangMingYaoShi',1,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(263,12,1,'影界果实',NULL,'用于培养魂侍',10,'YingJieGuoShi',1,3,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(264,12,1,'真气龙珠',NULL,'可用于提升角色武功境界等级',0,'ZhenQiLongZhu',1,1,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(265,12,3,'一星破界龙珠',NULL,'可用突破金刚境界',0,'YiXingPoJieLongZhu',1,1,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(266,12,3,'二星破界龙珠',NULL,'可用突破道玄境界',0,'ErXingPoJieLongZhu',1,1,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(267,12,3,'三星破界龙珠',NULL,'可用突破万象境界',0,'SanXingPoJieLongZhu',1,1,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(268,12,3,'四星破界龙珠',NULL,'可用突破天人境界',0,'SiXingPoJieLongZhu',1,1,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(270,12,1,'龙币',NULL,'特殊货币，可用于购买真气龙珠，破界龙珠等',0,'LongBi',1,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(271,9,2,'2级装备结晶箱',1,'开启后可获得装备结晶',100,'1JiZhuangBeiJieJingXiang',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(272,9,2,'2级装备碎片箱',1,'开启后可获得装备碎片',100,'1JiZhuangBeiSuiPianXiang',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(273,13,1,'魂侍碎片(武圣)',NULL,'收集50个魂侍碎片，可以召唤武圣。同时也是武圣进阶的必备材料。',0,'WuSheng',1,6,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(274,13,1,'魂侍碎片(剑灵)',NULL,'收集50个魂侍碎片，可以召唤剑灵。同时也是剑灵进阶的必备材料。',0,'GuDaiJianLing',1,6,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(275,13,1,'魂侍碎片(飞羽)',NULL,'收集50个魂侍碎片，可以召唤飞羽。同时也是飞羽进阶的必备材料。',0,'FeiYu',1,6,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(276,13,1,'魂侍碎片(阿修罗)',NULL,'收集50个魂侍碎片，可以召唤阿修罗。同时也是阿修罗进阶的必备材料。',0,'AXiuLuo',1,6,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(277,13,1,'魂侍碎片(洛神)',NULL,'收集50个魂侍碎片，可以召唤洛神。同时也是洛神进阶的必备材料。',0,'LuoShen',1,6,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(278,13,1,'魂侍碎片(木偶戏子)',NULL,'收集50个魂侍碎片，可以召唤木偶戏子。同时也是木偶戏子进阶的必备材料。',0,'MuOXiZi',1,6,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(282,14,3,'剑心碎片(巨阙)',NULL,'收集10个剑心碎片，可以合成巨阙剑心。提升攻击。',0,'JuQue',1,4,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(283,14,3,'剑心碎片(湛泸)',NULL,'收集10个剑心碎片，可以合成湛泸剑心。提升防御。',0,'ZhanLu',1,4,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(284,14,3,'剑心碎片(赤霄)',NULL,'收集10个剑心碎片，可以合成赤霄剑心。提升生命。',0,'ChiXiao',1,4,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(285,14,3,'剑心碎片(龙渊)',NULL,'收集10个剑心碎片，可以合成龙渊剑心。提升内力。',0,'LongYuan',1,4,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(286,14,3,'剑心碎片(承影)',NULL,'收集10个剑心碎片，可以合成承影剑心。提升速度。',0,'ChengYing',1,4,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(287,14,3,'剑心碎片(干将)',NULL,'收集10个剑心碎片，可以合成干将剑心。提升暴击。',0,'GanJiang',1,4,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(288,14,3,'剑心碎片(莫邪)',NULL,'收集10个剑心碎片，可以合成莫邪剑心。提升闪避。',0,'MoXie',1,4,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(289,14,3,'剑心碎片(星河)',NULL,'收集10个剑心碎片，可以合成星河剑心。提升格挡。',0,'XingHe',1,4,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(290,14,3,'剑心碎片(春秋)',NULL,'收集10个剑心碎片，可以合成春秋剑心。提升生命。',0,'ChunQiu',1,4,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(291,14,3,'剑心碎片(青梅)',NULL,'收集10个剑心碎片，可以合成青梅剑心。提升防御。',0,'QingMei',1,4,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(292,14,3,'剑心碎片(竹马)',NULL,'收集10个剑心碎片，可以合成竹马剑心。提升破击。',0,'ZhuMa',1,4,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(294,9,1,'比武场青铜宝箱',1,'开启后获得比武场奖励',0,'BiWuChangQingTongBaoXiang',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(295,9,2,'比武场白银宝箱',1,'开启后获得比武场奖励',0,'BiWuChangBaiYinBaoXiang',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(296,9,3,'比武场黄金宝箱',1,'开启后获得比武场奖励',0,'BiWuChangHuangJinBaoXiang',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(297,9,4,'比武场传奇宝箱',1,'开启后获得比武场奖励',0,'BiWuChangChuanQiBaoXiang',0,0,7,0,0,0,0,0,0,0,0,0,0,0,NULL),(300,9,2,'真气龙珠袋',1,'内涵20个真气龙珠的袋子',0,'ZhenQiLongZhuDai',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(301,9,3,'真气龙珠宝箱',1,'内涵100个真气龙珠的宝箱',0,'ZhenQiLongZhuBaoXiang',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(302,15,2,'武功心得',NULL,'代代相传的武功秘籍，可以增加10000点伙伴经验',0,'WuGongXinDe',1,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(303,15,2,'茶叶蛋',NULL,'仙灵权贵旅途居家必备闲食，可以恢复20点体力',20,'ChaYeDan',1,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(304,3,3,'BUG剑',1,NULL,100,'YingYangJian',0,0,0,0,0,0,1,100000,10000,10000,10000,0,5,4,NULL),(305,9,2,'珠宝盒',1,'开启获得大量铜钱',0,'ZhuBaoHe',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(306,9,1,'铜钱袋',1,'开启后获得铜钱',0,'TongQianDai',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(307,9,3,'潜龙剑心礼包',1,'开启后获得20个潜龙剑心',0,'QianLongJianXin',0,0,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(308,11,1,'契约球(灯笼怪)',NULL,'与灯笼怪达成召唤契约的球',500,'QiYueQiuDengLongGuai',0,5,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(309,11,2,'契约球(画妖)',NULL,'与画妖达成召唤契约的球',1000,'QiYueQiuHuaYao',0,5,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(310,11,2,'契约球(魔笔)',NULL,'与魔笔达成召唤契约的球',1000,'QiYueQiuMoBi',0,5,0,0,0,0,0,0,0,0,0,0,0,0,NULL),(311,11,2,'契约球(梦妖)',NULL,'与梦妖达成召唤契约的球',1000,'QiYueQiuMengYao',0,5,0,0,0,0,0,0,0,0,0,0,0,0,NULL);
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
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COMMENT='宝箱内容';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `item_box_content` DISABLE KEYS */;
INSERT INTO `item_box_content` VALUES (1,230,2,0,231,'','',2,0,0),(2,230,2,0,232,'','',1,0,0),(3,239,2,0,235,'','',1,0,0),(4,239,2,0,236,'','',1,0,0),(5,239,2,0,237,'','',1,0,0),(6,239,2,0,238,'','',1,0,0),(7,271,2,0,231,'','',5,0,0),(8,271,2,0,232,'','',1,0,0),(9,272,2,0,235,'','',5,0,0),(10,272,2,0,236,'','',5,0,0),(11,272,2,0,237,'','',5,0,0),(12,272,2,0,238,'','',5,0,0),(13,294,2,0,270,'','',2,0,0),(14,294,2,0,231,'','',1,0,0),(15,295,2,0,270,'','',3,0,0),(16,295,2,0,231,'','',2,0,0),(17,295,2,0,232,'','',1,0,0),(18,296,2,0,270,'','',5,0,0),(19,296,2,0,232,'','',5,0,100),(20,296,2,0,233,'','',3,0,0),(21,297,2,0,270,'','',10,0,0),(22,297,2,0,232,'','',10,0,0),(23,297,2,0,233,'','',5,0,0),(24,300,2,0,270,'','',20,0,0),(25,301,2,0,270,'','',100,0,0),(26,305,0,0,0,'','',180000,220000,0),(27,306,0,0,0,'','',12000,18000,0);
/*!40000 ALTER TABLE `item_box_content` ENABLE KEYS */;
DROP TABLE IF EXISTS `item_costprops`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item_costprops` (
  `item_id` smallint(6) NOT NULL COMMENT '道具ID',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '消耗类型； 0 - 经验； 1 - 体力',
  `value` int(11) NOT NULL DEFAULT '0' COMMENT '值',
  PRIMARY KEY (`item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='消耗道具';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `item_costprops` DISABLE KEYS */;
INSERT INTO `item_costprops` VALUES (302,0,10000),(303,1,20);
/*!40000 ALTER TABLE `item_costprops` ENABLE KEYS */;
DROP TABLE IF EXISTS `item_exchange`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item_exchange` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `target_item_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '目标物品id',
  `item_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '物品id',
  `item_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '物品数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COMMENT='物品兑换表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `item_exchange` DISABLE KEYS */;
INSERT INTO `item_exchange` VALUES (1,264,270,1),(2,265,270,6),(3,266,270,12),(4,267,270,18),(5,268,270,24),(6,300,270,40),(7,301,270,100);
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
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COMMENT='物品类型';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `item_type` DISABLE KEYS */;
INSERT INTO `item_type` VALUES (2,'材料',9999,'Stuff',7),(3,'武器',1,'Weapon',1),(4,'战袍',1,'Clothes',2),(5,'靴子',1,'Shoes',3),(6,'饰品',1,'Accessories',4),(8,'战斗道具',99,'FightProp',5),(9,'礼包',99,'Chest',11),(10,'资源',999,'Resource',12),(11,'灵宠契约球',999,'Pet',10),(12,'道具',9999,'Props',6),(13,'魂侍碎片',999,'GhostFragment',8),(14,'剑心碎片',999,'SwordSoulFragment',9),(15,'消耗道具',9999,'CostProp',13);
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
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COMMENT='关卡灵宠配置';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `level_battle_pet` DISABLE KEYS */;
INSERT INTO `level_battle_pet` VALUES (1,64,2,60,3),(2,65,2,60,3),(3,66,2,60,3),(4,82,3,30,2),(5,83,3,30,2),(6,100,3,30,2),(7,58,2,100,3),(10,76,3,100,2),(11,77,3,30,2),(12,99,3,30,2);
/*!40000 ALTER TABLE `level_battle_pet` ENABLE KEYS */;
DROP TABLE IF EXISTS `level_star`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `level_star` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `level_id` int(11) NOT NULL COMMENT '关卡ID',
  `two_star_score` int(11) NOT NULL COMMENT '两星要求分数',
  `three_star_score` int(11) NOT NULL COMMENT '三星要求分数',
  PRIMARY KEY (`id`),
  UNIQUE KEY `level_id` (`level_id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COMMENT='关卡星级分数表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `level_star` DISABLE KEYS */;
INSERT INTO `level_star` VALUES (1,62,500,1000),(2,63,1000,2000),(3,64,1500,2500),(4,65,2000,3500),(5,66,2500,4000),(7,3,300,600),(8,9,600,1100),(9,12,1000,2000),(10,15,1400,2800),(11,18,1500,3000),(12,21,1800,3600),(13,24,1850,3700),(14,27,2500,5200),(15,30,3000,7200);
/*!40000 ALTER TABLE `level_star` ENABLE KEYS */;
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
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COMMENT='系统邮件模板';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mail` DISABLE KEYS */;
INSERT INTO `mail` VALUES (1,'BagFull',0,'背包已满提示','func,功能','您的背包已满，系统已自动将{0}的物品暂存至附件，请及时点击领取。'),(2,'Heart',0,'爱心赠送邮件','who,发送者','您的好友{0}赠送您一颗爱心哦！请及时点击领取。'),(3,'TestMail',0,'测试邮件','p1, 参数1; p2, 参数2','我{0}只是一封测试邮件{1}，虽然偶只是测试用的，但四，你八可以八理我，if你不点我我that会伤心的。so !选我！选我！选我！'),(4,'MultiLevel',0,'多人关卡战斗奖励','name,关卡名称','您在参与的多人关卡{0}中取得了胜利，附件里是您获得的奖励。'),(5,'Welcome',0,'欢迎进入仙侠道','','我们诚挚欢迎您加入仙侠道的武侠世界，可以在这里一圆您的武侠梦！特此送上我们的小礼物，快去发觉他的妙用吧~'),(6,'GhostBagFull',0,'魂侍背包已满提示','func,功能','您的魂侍背包已满，系统已自动将您刚抽取到的{0}个魂侍暂存至附件，请及时点击领取。'),(7,'SwordSoulBagFull',0,'剑心背包已满提示','func,功能','您的剑心背包已满，系统已自动将您刚抽取到的{0}个剑心暂存至附件，请及时点击领取。'),(8,'Recharge',0,'充值成功','time1,时间; num,充值元宝数量','恭喜您在{0}成功充值{1}元宝，仙侠道团体感谢您对我们游戏的支持。'),(9,'PurchaseTips',0,'购买道具成功提示','source,来源;item_name,道具名称;func,功能;','大侠您刚刚在{0}得到了{1}，此道具可在{2}功能开放后在此系统背包中查看。');
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
INSERT INTO `mail_attachments` VALUES (4,3,31,0,1),(5,3,238,0,10),(6,2,0,3,1),(7,3,34,0,2),(8,3,63,0,1),(9,5,264,0,10);
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
INSERT INTO `mission` VALUES (1,1,0,'青竹林','QingZhuLin',1),(2,1,2,'黑夜森林','HeiYeSenLin',2),(3,1,4,'莲花峰','LianHuaFeng',3),(4,1,5,'熔岩火山','RongYanHuoShan',4),(5,1,10,'剑灵密室','JianLingMiShi',5),(6,3,0,'测试区域','QingZhuLin',1),(7,2,8,'血雾岭','XueWuLing',1);
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
  `boss_dir` tinyint(4) NOT NULL COMMENT '怪物朝向(0--左;1--右)',
  `best_round` tinyint(4) NOT NULL COMMENT '最好的通关回合数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=247 DEFAULT CHARSET=utf8mb4 COMMENT='副本敌人';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mission_enemy` DISABLE KEYS */;
INSERT INTO `mission_enemy` VALUES (10,4,2,605,377,0,0,0,0,0,0,0,0,0,0,0,1,0,0),(11,4,2,1217,312,0,0,0,0,0,0,0,0,0,0,0,2,0,0),(19,7,3,821,922,34,40,35,30,36,0,0,0,0,30,0,1,0,0),(20,7,3,1492,350,34,40,35,30,36,0,0,0,0,30,0,2,0,0),(27,9,3,1149,479,0,0,0,0,0,0,0,0,0,100,1,1,0,0),(31,11,3,679,571,24,50,6,25,7,0,0,0,0,25,0,1,0,0),(32,11,3,1543,518,24,50,6,25,7,0,0,0,0,25,0,2,0,0),(36,12,3,900,480,0,0,0,0,0,0,0,0,0,0,1,1,0,0),(37,13,3,876,477,38,20,39,40,40,0,0,0,0,40,0,1,0,0),(38,13,3,1620,894,38,20,39,40,40,0,0,0,0,40,0,2,0,0),(39,13,3,984,1300,38,20,39,40,40,0,0,0,0,40,0,3,0,0),(45,15,3,796,512,0,0,0,0,0,0,0,0,0,0,1,1,1,0),(46,16,3,818,235,41,50,42,50,0,0,0,0,0,0,0,1,0,0),(47,16,3,939,873,41,50,42,50,0,0,0,0,0,0,0,2,0,0),(49,17,3,1103,354,27,50,13,50,0,0,0,0,0,0,0,1,0,0),(50,17,3,1010,998,27,50,13,50,0,0,0,0,0,0,0,2,0,0),(54,18,3,1113,487,0,0,0,0,0,0,0,0,0,0,1,1,0,0),(55,19,3,422,569,44,20,43,50,45,0,0,0,0,30,0,1,0,0),(56,19,3,870,305,44,20,43,50,45,0,0,0,0,30,0,2,0,0),(57,19,3,1376,530,44,20,43,50,45,0,0,0,0,30,0,3,0,0),(58,20,3,1339,708,16,20,15,50,17,0,0,0,0,30,0,1,0,0),(59,20,3,566,1162,16,20,15,50,17,0,0,0,0,30,0,2,0,0),(60,20,3,1083,1651,16,20,15,50,17,0,0,0,0,30,0,3,0,0),(63,21,3,1034,526,0,0,0,0,0,0,0,0,0,0,1,1,0,0),(64,23,3,1023,762,16,50,15,25,17,0,0,0,0,25,0,1,0,0),(65,23,3,1740,1218,16,50,15,25,17,0,0,0,0,25,0,2,0,0),(66,23,3,900,1839,16,50,15,25,17,0,0,0,0,25,0,3,0,0),(67,22,3,422,569,44,50,43,25,45,0,0,0,0,25,0,1,0,0),(68,22,3,870,305,44,50,43,25,45,0,0,0,0,25,0,2,0,0),(69,22,3,1376,530,44,50,43,25,45,0,0,0,0,25,0,3,0,0),(72,24,3,880,510,0,0,0,0,0,0,0,0,0,0,1,1,1,0),(73,25,3,753,363,46,60,47,30,48,0,0,0,0,10,0,1,0,0),(74,25,3,1592,314,46,60,47,30,48,0,0,0,0,10,0,2,0,0),(76,26,3,1547,363,21,60,20,30,22,0,0,0,0,10,0,1,0,0),(77,26,3,708,314,21,60,20,30,22,0,0,0,0,10,0,2,0,0),(79,28,3,753,363,46,40,47,20,48,0,0,0,0,40,0,1,0,0),(80,28,3,1592,314,46,40,47,20,48,0,0,0,0,40,0,2,0,0),(82,29,3,1547,363,21,40,20,20,22,0,0,0,0,40,0,1,0,0),(83,29,3,708,314,21,40,20,20,22,0,0,0,0,40,0,2,0,0),(87,27,3,1302,637,0,0,0,0,0,0,0,0,0,0,1,1,0,0),(90,30,3,618,637,0,0,0,0,0,0,0,0,0,0,1,1,1,0),(93,3,3,981,470,0,0,0,0,0,0,0,0,0,0,1,1,0,0),(96,31,1,0,0,12,100,0,0,0,0,0,0,0,0,0,2,0,0),(97,25,3,1019,914,46,60,47,30,48,0,0,0,0,10,0,3,0,0),(98,28,3,1019,914,46,40,47,20,48,0,0,0,0,40,0,3,0,0),(99,26,3,1281,914,21,60,20,30,22,0,0,0,0,10,0,3,0,0),(100,29,3,1281,914,21,40,20,20,22,0,0,0,0,40,0,3,0,0),(105,31,1,0,0,49,100,0,0,0,0,0,0,0,0,0,1,0,0),(106,31,1,0,0,50,100,0,0,0,0,0,0,0,0,0,17,0,0),(107,31,1,0,0,51,100,0,0,0,0,0,0,0,0,0,21,0,0),(108,31,1,0,0,52,100,0,0,0,0,0,0,0,0,0,25,0,0),(109,31,1,0,0,53,100,0,0,0,0,0,0,0,0,0,29,0,0),(110,31,1,0,0,54,100,0,0,0,0,0,0,0,0,0,33,0,0),(111,31,1,0,0,55,100,0,0,0,0,0,0,0,0,0,37,0,0),(112,31,1,0,0,56,100,0,0,0,0,0,0,0,0,0,3,0,0),(113,31,1,0,0,57,100,0,0,0,0,0,0,0,0,0,4,0,0),(114,31,1,0,0,58,100,0,0,0,0,0,0,0,0,0,5,0,0),(115,31,1,0,0,59,100,0,0,0,0,0,0,0,0,0,6,0,0),(116,31,1,0,0,60,100,0,0,0,0,0,0,0,0,0,7,0,0),(117,31,1,0,0,61,100,0,0,0,0,0,0,0,0,0,14,0,0),(118,31,1,0,0,62,100,0,0,0,0,0,0,0,0,0,15,0,0),(119,31,1,0,0,67,100,0,0,0,0,0,0,0,0,0,16,0,0),(120,31,1,0,0,63,100,0,0,0,0,0,0,0,0,0,8,0,0),(121,31,1,0,0,64,100,0,0,0,0,0,0,0,0,0,9,0,0),(122,31,1,0,0,65,100,0,0,0,0,0,0,0,0,0,10,0,0),(123,31,1,0,0,66,100,0,0,0,0,0,0,0,0,0,11,0,0),(124,31,1,0,0,68,100,0,0,0,0,0,0,0,0,0,12,0,0),(125,31,1,0,0,69,100,0,0,0,0,0,0,0,0,0,13,0,0),(126,31,1,0,0,70,100,0,0,0,0,0,0,0,0,0,18,0,0),(127,31,1,0,0,76,100,0,0,0,0,0,0,0,0,0,19,0,0),(128,31,1,0,0,82,100,0,0,0,0,0,0,0,0,0,20,0,0),(129,31,1,0,0,71,100,0,0,0,0,0,0,0,0,0,22,0,0),(130,31,1,0,0,77,100,0,0,0,0,0,0,0,0,0,23,0,0),(131,31,1,0,0,83,100,0,0,0,0,0,0,0,0,0,24,0,0),(132,31,1,0,0,72,100,0,0,0,0,0,0,0,0,0,26,0,0),(133,31,1,0,0,78,100,0,0,0,0,0,0,0,0,0,27,0,0),(134,31,1,0,0,84,100,0,0,0,0,0,0,0,0,0,28,0,0),(135,31,1,0,0,73,100,0,0,0,0,0,0,0,0,0,30,0,0),(136,31,1,0,0,79,100,0,0,0,0,0,0,0,0,0,31,0,0),(137,31,1,0,0,85,100,0,0,0,0,0,0,0,0,0,32,0,0),(138,31,1,0,0,74,100,0,0,0,0,0,0,0,0,0,34,0,0),(139,31,1,0,0,80,100,0,0,0,0,0,0,0,0,0,35,0,0),(140,31,1,0,0,86,100,0,0,0,0,0,0,0,0,0,36,0,0),(141,31,1,0,0,75,100,0,0,0,0,0,0,0,0,0,38,0,0),(142,31,1,0,0,81,100,0,0,0,0,0,0,0,0,0,39,0,0),(143,31,1,0,0,87,100,0,0,0,0,0,0,0,0,0,40,0,0),(148,38,5,900,450,88,100,0,0,0,0,0,0,0,0,0,1,0,0),(149,39,5,900,450,89,100,0,0,0,0,0,0,0,0,0,1,0,0),(155,42,3,1023,762,44,20,43,50,45,0,0,0,0,30,0,1,0,0),(156,42,3,1740,1218,44,20,43,50,45,0,0,0,0,30,0,2,0,0),(157,42,3,900,1839,44,20,43,50,45,0,0,0,0,30,0,3,0,0),(158,43,3,1339,708,44,50,43,25,45,0,0,0,0,25,0,1,0,0),(159,43,3,566,1162,44,50,43,25,45,0,0,0,0,25,0,2,0,0),(160,43,3,1083,1651,44,50,43,25,45,0,0,0,0,25,0,3,0,0),(161,44,3,753,363,46,60,47,30,48,0,0,0,0,10,0,1,0,0),(162,44,3,1592,314,46,60,47,30,48,0,0,0,0,10,0,2,0,0),(163,44,3,1019,914,46,60,47,30,48,0,0,0,0,10,0,3,0,0),(164,45,3,753,363,46,40,47,20,48,0,0,0,0,40,0,1,0,0),(165,45,3,1592,314,46,40,47,20,48,0,0,0,0,40,0,2,0,0),(166,45,3,1019,914,46,40,47,20,48,0,0,0,0,40,0,3,0,0),(167,46,1,100,100,90,100,0,0,0,0,0,0,0,0,0,1,0,0),(168,51,5,900,450,88,100,0,0,0,0,0,0,0,0,0,1,0,0),(169,52,5,900,450,89,100,0,0,0,0,0,0,0,0,0,1,0,0),(170,38,3,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,0,0),(171,47,5,900,450,88,100,0,0,0,0,0,0,0,0,0,1,0,0),(172,47,3,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,0,0),(173,49,5,900,450,88,100,0,0,0,0,0,0,0,0,0,1,0,0),(174,49,3,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,0,0),(175,59,5,900,450,88,100,0,0,0,0,0,0,0,0,0,1,0,0),(176,59,1,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,0,0),(177,60,5,900,450,89,100,0,0,0,0,0,0,0,0,0,1,0,0),(178,60,1,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,0,0),(179,51,3,1400,400,0,0,0,0,0,0,0,0,0,0,1,2,0,0),(180,39,1,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,0,0),(181,53,5,900,450,88,100,0,0,0,0,0,0,0,0,0,1,0,0),(182,53,3,1400,400,0,0,0,0,0,0,0,0,0,0,1,2,0,0),(183,55,5,900,450,88,100,0,0,0,0,0,0,0,0,0,1,0,0),(184,55,3,1400,400,0,0,0,0,0,0,0,0,0,0,1,2,0,0),(185,57,5,900,450,88,100,0,0,0,0,0,0,0,0,0,1,0,0),(186,57,3,1400,400,0,0,0,0,0,0,0,0,0,0,1,2,0,0),(187,61,5,900,450,88,100,0,0,0,0,0,0,0,0,0,1,0,0),(188,61,3,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,0,0),(189,58,5,900,450,88,100,0,0,0,0,0,0,0,0,0,1,0,0),(190,58,1,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,0,0),(191,48,5,900,450,89,100,0,0,0,0,0,0,0,0,0,1,0,0),(192,48,1,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,0,0),(193,50,5,900,450,89,100,0,0,0,0,0,0,0,0,0,1,0,0),(194,50,1,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,0,0),(195,52,1,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,0,0),(196,54,5,900,450,89,100,0,0,0,0,0,0,0,0,0,1,0,0),(197,54,1,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,0,0),(198,56,5,900,450,89,100,0,0,0,0,0,0,0,0,0,1,0,0),(199,56,1,1300,400,0,0,0,0,0,0,0,0,0,0,1,2,0,0),(200,46,1,100,100,94,100,0,0,0,0,0,0,0,0,0,2,0,0),(201,62,1,959,549,0,0,0,0,0,0,0,0,0,0,1,1,0,0),(202,63,1,959,549,0,0,0,0,0,0,0,0,0,0,1,1,0,0),(203,64,1,959,549,0,0,0,0,0,0,0,0,0,0,1,1,0,0),(204,65,3,959,549,0,0,0,0,0,0,0,0,0,0,1,1,0,0),(205,66,1,959,549,0,0,0,0,0,0,0,0,0,0,1,1,0,0),(206,67,1,100,100,98,100,0,0,0,0,0,0,0,0,0,1,0,0),(210,31,1,0,0,102,100,0,0,0,0,0,0,0,0,0,41,0,0),(211,70,4,753,363,33,100,0,0,0,0,0,0,0,0,0,1,0,0),(212,70,4,1592,314,33,100,0,0,0,0,0,0,0,0,0,2,0,0),(213,70,1,1173,885,0,0,0,0,0,0,0,0,0,0,1,3,1,0),(214,71,2,876,477,108,100,0,0,0,0,0,0,0,0,0,1,0,0),(215,71,2,1620,894,108,100,0,0,0,0,0,0,0,0,0,2,0,0),(216,71,1,960,1288,0,0,0,0,0,0,0,0,0,0,1,3,1,0),(217,72,4,679,571,33,50,109,50,0,0,0,0,0,0,0,1,0,0),(218,72,3,1543,518,0,0,0,0,0,0,0,0,0,0,1,2,0,0),(219,73,4,753,363,33,100,0,0,0,0,0,0,0,0,0,1,0,0),(220,73,4,1592,314,33,100,0,0,0,0,0,0,0,0,0,2,0,0),(221,73,1,1193,885,0,0,0,0,0,0,0,0,0,0,1,3,1,0),(222,74,4,753,363,33,100,0,0,0,0,0,0,0,0,0,1,0,0),(223,74,4,1592,314,33,100,0,0,0,0,0,0,0,0,0,2,0,0),(224,74,1,1173,885,0,0,0,0,0,0,0,0,0,0,1,3,1,0),(225,75,4,753,363,33,100,0,0,0,0,0,0,0,0,0,1,0,0),(226,75,4,1592,314,33,100,0,0,0,0,0,0,0,0,0,2,0,0),(227,75,1,1173,885,0,0,0,0,0,0,0,0,0,0,1,3,1,0),(228,79,2,876,477,108,100,0,0,0,0,0,0,0,0,0,1,0,0),(229,79,2,1620,894,108,100,0,0,0,0,0,0,0,0,0,2,0,0),(230,79,1,960,1288,0,0,0,0,0,0,0,0,0,0,1,3,1,0),(231,80,2,876,477,108,100,0,0,0,0,0,0,0,0,0,1,0,0),(232,80,2,1620,894,108,100,0,0,0,0,0,0,0,0,0,2,0,0),(233,80,1,960,1288,0,0,0,0,0,0,0,0,0,0,1,3,1,0),(234,81,2,876,477,108,100,0,0,0,0,0,0,0,0,0,1,0,0),(235,81,2,1620,894,108,100,0,0,0,0,0,0,0,0,0,2,0,0),(236,81,1,960,1288,0,0,0,0,0,0,0,0,0,0,1,3,1,0),(237,76,4,679,571,33,50,109,50,0,0,0,0,0,0,0,1,0,0),(238,76,3,1543,518,0,0,0,0,0,0,0,0,0,0,1,2,0,0),(239,77,4,679,571,33,50,109,50,0,0,0,0,0,0,0,1,0,0),(240,77,3,1543,518,0,0,0,0,0,0,0,0,0,0,1,2,0,0),(241,78,4,679,571,33,50,109,50,0,0,0,0,0,0,0,1,0,0),(242,78,3,1543,518,0,0,0,0,0,0,0,0,0,0,1,2,0,0),(243,46,15,100,100,31,100,0,0,0,0,0,0,0,0,0,3,0,0),(244,82,1,0,0,0,0,0,0,0,0,0,0,0,0,1,1,0,0),(245,82,1,0,0,0,0,0,0,0,0,0,0,0,0,1,2,0,0),(246,82,1,0,0,0,0,0,0,0,0,0,0,0,0,1,3,0,0);
/*!40000 ALTER TABLE `mission_enemy` ENABLE KEYS */;
DROP TABLE IF EXISTS `mission_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mission_level` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '区域关卡ID',
  `mission_id` smallint(6) NOT NULL COMMENT '区域ID',
  `parent_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '关联关卡类型(0-区域关卡;1-资源关卡;2-通天塔;8-难度关卡;9-伙伴关卡;10-灵宠关卡;11-魂侍关卡)',
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
) ENGINE=InnoDB AUTO_INCREMENT=91 DEFAULT CHARSET=utf8mb4 COMMENT='区域关卡配置';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mission_level` DISABLE KEYS */;
INSERT INTO `mission_level` VALUES (3,1,0,0,100110,'青竹林二',2,2,20,1150,540,1,20,618,190,'QingZhuLin2','QingZhuLin2','PingHe',110100,0,300,1,0,0),(4,1,0,0,100100,'青竹林一',0,-1,6,1769,343,1,10,828,144,'QingZhuLin1','QingZhuLin2','PingHe',100110,0,100,1,0,0),(7,2,0,0,110100,'黑夜森林一',0,-1,6,2048,268,1,15,856,114,'HeiYeSenLin1','HeiYeSenLin1War','YinSen',110110,0,150,1,0,0),(9,2,0,0,110110,'黑夜森林二',2,2,20,1419,606,1,45,388,140,'HeiYeSenLin2','HeiYeSenLin1War','YinSen',110120,0,450,1,0,0),(11,2,0,0,110120,'暗影秘境一',1,3,12,2072,249,1,30,270,203,'AnYingMiJing1','AnYingMiJing2','YinSen',110130,0,300,1,0,0),(12,2,0,0,110130,'暗影秘境二',2,2,20,1189,613,1,45,124,336,'AnYingMiJing2','AnYingMiJing2','YinSen',0,0,450,1,0,0),(13,3,0,0,120100,'莲花峰一',0,-1,6,712,1460,1,40,369,126,'LianHuaFeng1','LianHuaFeng2','YouQu',120110,1,200,1,0,0),(15,3,0,0,120110,'莲花峰二',2,2,20,586,611,1,120,122,1486,'LianHuaFeng2','LianHuaFeng2','YouQu',120120,1,600,1,0,0),(16,3,0,0,120120,'水溶洞一',0,-1,6,693,1322,1,40,306,150,'ShuiRongDong1','ShuiRongDong3','YinSen',120130,1,200,1,0,0),(17,3,0,0,120130,'水溶洞二',1,3,12,1287,1314,1,80,296,1844,'ShuiRongDong1','ShuiRongDong3','YinSen',120140,0,400,1,1,0),(18,3,0,0,120140,'水溶洞三',2,2,20,1314,722,1,120,362,250,'ShuiRongDong3','ShuiRongDong3','YinSen',130100,0,600,1,0,0),(19,4,0,0,130100,'熔岩火山一',0,-1,6,1742,764,1,125,814,104,'RongYanHuoShan1','RongYanHuoShan2War','YinSen',130110,0,400,1,0,0),(20,4,0,0,130110,'熔岩火山二',1,3,12,1917,1952,1,250,589,2102,'RongYanHuoShan2','RongYanHuoShan2War','YinSen',130115,0,800,1,1,0),(21,4,0,0,130120,'熔岩火山四',2,2,20,1544,513,1,375,189,169,'RongYanHuoShan4','RongYanHuoShan2War','YinSen',130130,0,1200,1,1,0),(22,4,0,0,130130,'熔岩火山五',0,-1,6,1742,764,1,125,814,104,'RongYanHuoShan1','RongYanHuoShan2War','YinSen',130140,0,400,1,0,0),(23,4,0,0,130140,'熔岩火山六',1,3,12,288,1980,1,250,612,246,'RongYanHuoShan2','RongYanHuoShan2War','YinSen',130145,1,800,1,0,0),(24,4,0,0,130150,'熔岩火山八',2,2,20,378,511,1,375,192,1752,'RongYanHuoShan4','RongYanHuoShan2War','YinSen',140100,1,1200,1,0,0),(25,5,0,0,140100,'剑灵密室一',0,-1,6,540,1299,1,280,441,134,'JianLingMiShi1','JianLingMiShi1War','ShenMi',140110,1,600,1,0,0),(26,5,0,0,140110,'剑灵密室二',1,3,12,1760,1299,1,560,441,2166,'JianLingMiShi1','JianLingMiShi1War','ShenMi',140115,0,1200,1,1,0),(27,5,0,0,140120,'剑灵密室四',2,2,20,1410,708,1,840,333,714,'JianLingMiShi4','JianLingMiShi4War','ShenMi',140130,0,1800,1,1,0),(28,5,0,0,140130,'剑灵密室五',0,-1,6,540,1299,1,280,441,134,'JianLingMiShi1','JianLingMiShi1War','ShenMi',140140,1,600,1,0,0),(29,5,0,0,140140,'剑灵密室六',1,3,12,1760,1299,1,560,441,2166,'JianLingMiShi1','JianLingMiShi1War','ShenMi',140145,0,1200,1,1,0),(30,5,0,0,140150,'剑灵密室八',2,2,20,510,708,1,840,333,1206,'JianLingMiShi4','JianLingMiShi4War','ShenMi',150100,1,1800,1,0,0),(31,6,0,0,999996,'绝招关卡',0,-1,0,1398,812,0,0,1114,176,'QingZhuLin1','QingZhuLin4','',999997,0,0,1,0,0),(38,0,1,2,0,'15级铜钱关卡',0,2,6,1500,500,0,100,500,300,'QingZhuLin2','QingZhuLin2','PingHe',0,0,10000,1,0,1),(39,0,1,2,0,'15级经验关卡',0,2,6,1500,500,0,10000,500,300,'QingZhuLin2','QingZhuLin2','PingHe',0,0,100,1,0,2),(42,4,0,0,130115,'熔岩火山三',0,-1,6,288,1980,1,125,612,246,'RongYanHuoShan2','RongYanHuoShan2War','YinSen',130120,1,400,1,0,0),(43,4,0,0,130145,'熔岩火山七',0,-1,6,1917,1952,1,125,589,2102,'RongYanHuoShan2','RongYanHuoShan2War','YinSen',130150,0,400,1,1,0),(44,5,0,0,140115,'剑灵密室三',0,-1,6,540,1299,1,280,441,134,'JianLingMiShi1','JianLingMiShi1War','ShenMi',140120,1,600,1,0,0),(45,5,0,0,140145,'剑灵密室七',0,-1,6,540,1299,1,280,441,134,'JianLingMiShi1','JianLingMiShi1War','ShenMi',140150,1,340,1,0,0),(46,6,0,0,999998,'状态关卡',0,-1,0,1398,812,0,0,1114,176,'QingZhuLin1','QingZhuLin4','',999999,0,0,1,0,0),(47,0,1,3,0,'30级铜钱关卡',0,2,6,1500,500,0,100,500,300,'QingZhuLin2','QingZhuLin2','PingHe',0,0,10000,1,0,1),(48,0,1,3,0,'30级经验关卡',0,2,6,1500,500,0,20000,500,300,'QingZhuLin2','QingZhuLin2','PingHe',0,0,100,1,0,2),(51,0,1,5,0,'50级铜钱关卡',0,2,6,1500,500,0,100,500,300,'QingZhuLin2','QingZhuLin2','PingHe',0,0,10000,1,0,1),(52,0,1,5,0,'50级经验关卡',0,2,6,1500,500,0,80000,500,300,'QingZhuLin2','QingZhuLin2','PingHe',0,0,100,1,0,2),(55,0,1,7,0,'70级铜钱关卡',0,2,6,1500,500,0,100,500,300,'QingZhuLin2','QingZhuLin2','PingHe',0,0,10000,1,0,1),(56,0,1,7,0,'70级经验关卡',0,2,6,1500,500,0,140000,500,300,'QingZhuLin2','QingZhuLin2','PingHe',0,0,100,1,0,2),(62,0,8,1,0,'噩梦刀疤兔',2,1,20,1410,652,1,20,759,304,'EMengDaoBaTu','EMengDaoBaTu','PingHe',0,0,300,1,0,0),(63,0,8,2,0,'噩梦妖龙',2,1,20,1410,652,1,90,759,304,'EMengDaoBaTu','EMengDaoBaTu','YinSen',0,0,900,1,0,0),(64,0,8,3,0,'噩梦剧毒臭泥',2,1,20,1410,652,1,120,759,304,'EMengDaoBaTu','EMengDaoBaTu','YinSen',0,0,600,1,0,0),(65,0,8,4,0,'噩梦炎龙',2,1,20,1410,652,1,375,759,304,'EMengDaoBaTu','EMengDaoBaTu','YinSen',0,0,1200,1,0,0),(66,0,8,5,0,'噩梦古代剑灵',2,1,20,1410,652,1,840,759,304,'EMengDaoBaTu','EMengDaoBaTu','ShenMi',0,0,1800,1,0,0),(67,6,0,0,1000000,'魂侍关卡',0,-1,0,1398,812,0,0,1114,176,'QingZhuLin1','QingZhuLin4','',1000001,0,0,1,0,0),(70,0,9,8,0,'20级伙伴试炼',0,5,6,540,1299,0,1000,441,134,'JianLingMiShi1','JianLingMiShi1War','ShenMi',0,1,1000,1,0,0),(71,0,10,12,0,'20级灵宠试炼',0,5,6,712,1460,0,1000,369,126,'LianHuaFeng1','LianHuaFeng2','YouQu',0,1,1000,1,0,0),(72,0,11,13,0,'20级魂侍试炼',0,5,6,2072,249,0,1000,270,203,'AnYingMiJing1','AnYingMiJing2','YinSen',0,0,1000,1,0,0),(73,0,9,9,0,'30级伙伴试炼',0,5,6,540,1299,0,1000,441,134,'JianLingMiShi1','JianLingMiShi1War','ShenMi',0,1,1000,1,0,0),(74,0,9,10,0,'50级伙伴试炼',0,5,6,540,1299,0,1000,441,134,'JianLingMiShi1','JianLingMiShi1War','ShenMi',0,1,1000,1,0,0),(75,0,9,11,0,'70级伙伴试炼',0,5,6,540,1299,0,1000,441,134,'JianLingMiShi1','JianLingMiShi1War','ShenMi',0,1,1000,1,0,0),(76,0,11,14,0,'30级魂侍关卡',0,5,6,2072,249,0,1000,270,203,'AnYingMiJing1','AnYingMiJing2','YinSen',0,0,1000,1,0,0),(77,0,11,15,0,'50级魂侍试炼',0,5,6,2072,249,0,1000,270,203,'AnYingMiJing1','AnYingMiJing2','YinSen',0,0,1000,1,0,0),(78,0,11,16,0,'70级魂侍试炼',0,5,6,2072,249,0,1000,270,203,'AnYingMiJing1','AnYingMiJing2','YinSen',0,0,1000,1,0,0),(79,0,10,17,0,'30级灵宠试炼',0,5,6,712,1460,0,1000,369,126,'LianHuaFeng1','LianHuaFeng2','YouQu',0,1,1000,1,0,0),(80,0,10,18,0,'50级灵宠试炼',0,5,6,712,1460,0,1000,369,126,'LianHuaFeng1','LianHuaFeng2','YouQu',0,1,1000,1,0,0),(81,0,10,19,0,'70级灵宠试炼',0,5,6,712,1460,0,1000,369,126,'LianHuaFeng1','LianHuaFeng2','YouQu',0,1,1000,1,0,0),(82,6,0,0,999999,'怪物测试',0,-1,0,0,0,0,0,0,0,'QingZhuLin1','QingZhuLin2','Music',1000000,0,0,1,0,0),(83,7,0,0,150100,'血雾岭一',0,-1,6,0,0,1,620,0,0,'XueWuLing1','XueWuLing1War','YinSen',150110,0,1500,1,0,0),(84,7,0,0,150110,'血雾岭二',1,3,12,0,0,1,1240,0,0,'XueWuLing1','XueWuLing1War','YinSen',150120,0,3000,1,0,0),(85,7,0,0,150120,'血雾岭三',0,-1,6,0,0,1,620,0,0,'XueWuLing1','XueWuLing1War','YinSen',150130,0,1500,1,0,0),(86,7,0,0,150130,'血雾岭四',2,2,20,0,0,1,1860,0,0,'XueWuLing4','XueWuLing1War','YinSen',150140,0,4500,1,0,0),(87,7,0,0,150140,'血雾地牢一',0,-1,6,0,0,1,620,0,0,'XueWuDiLao1','XueWuDiLao1War','YinSen',150150,0,1500,1,0,0),(88,7,0,0,150150,'血雾地牢二',1,3,12,0,0,1,1240,0,0,'XueWuDiLao1','XueWuDiLao1War','YinSen',150160,0,3000,1,0,0),(89,7,0,0,150160,'血雾地牢三',0,-1,6,0,0,1,620,0,0,'XueWuDiLao1','XueWuDiLao1War','YinSen',150170,0,1500,1,0,0),(90,7,0,0,150170,'血雾地牢四',2,2,20,0,0,1,1860,0,0,'XueWuDiLao4','XueWuDiLao1War','YinSen',160100,0,4500,1,0,0);
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
) ENGINE=InnoDB AUTO_INCREMENT=431 DEFAULT CHARSET=utf8mb4 COMMENT='区域关卡宝箱';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mission_level_box` DISABLE KEYS */;
INSERT INTO `mission_level_box` VALUES (12,3,1,1,20,1,239),(13,3,2,1,20,1,230),(14,3,3,2,20,2,36),(15,3,4,2,20,1,37),(16,3,5,2,20,1,38),(17,4,1,1,20,1,264),(18,4,2,4,10,2,245),(19,4,3,5,10,2,246),(20,4,4,3,30,10,243),(21,4,5,0,30,100,244),(38,7,1,1,20,2,264),(41,7,2,4,10,2,245),(42,7,3,5,10,2,246),(43,7,4,3,30,25,243),(50,9,1,2,20,1,36),(51,9,2,2,20,1,37),(55,9,3,2,20,1,38),(56,7,5,0,30,200,244),(57,9,4,1,20,1,230),(58,9,5,1,20,1,239),(64,11,1,1,10,4,270),(65,11,2,1,20,4,264),(66,11,3,5,20,2,246),(67,11,4,3,25,61,243),(68,11,5,0,25,400,244),(69,12,1,2,20,1,31),(70,12,2,2,20,1,32),(71,12,3,2,20,1,33),(72,12,4,1,20,1,230),(73,12,5,0,20,1,239),(74,13,1,1,20,2,264),(75,13,2,4,10,2,245),(76,13,3,5,10,2,246),(77,13,4,3,30,70,243),(79,13,5,0,30,300,244),(86,15,1,2,20,1,36),(87,15,2,2,20,1,37),(88,15,3,2,20,1,38),(90,15,4,1,20,1,230),(91,15,5,1,20,1,239),(92,16,1,1,20,2,270),(93,16,2,4,10,2,245),(94,16,3,5,10,2,246),(95,16,4,3,30,84,243),(97,16,5,0,30,300,244),(98,17,1,1,10,4,270),(99,17,2,1,20,4,264),(100,17,3,1,20,2,263),(101,17,4,3,25,166,243),(103,17,5,0,25,600,244),(104,18,1,2,20,1,31),(105,18,2,2,20,1,32),(106,18,3,2,20,1,33),(108,18,4,1,20,1,230),(109,18,5,1,20,1,239),(110,19,1,1,20,3,270),(111,19,2,4,10,2,245),(112,19,3,5,10,2,246),(113,19,4,3,30,230,243),(115,19,5,0,50,400,244),(117,20,1,1,10,6,270),(118,20,2,1,20,6,264),(119,20,3,1,20,3,263),(120,20,4,3,25,460,243),(122,20,5,0,25,800,244),(123,21,1,2,20,1,36),(124,21,2,2,20,1,37),(125,21,3,2,20,1,38),(127,21,4,1,20,1,230),(128,21,5,1,20,1,239),(129,22,1,1,20,3,270),(130,22,2,4,10,2,245),(131,22,3,5,10,2,246),(132,22,4,3,30,258,243),(134,22,5,0,30,400,244),(136,23,1,1,10,6,270),(137,23,2,1,20,6,264),(138,23,3,1,20,3,263),(139,23,4,3,25,528,243),(141,23,5,0,25,800,244),(142,24,1,2,20,1,31),(143,24,2,2,20,1,32),(144,24,3,2,20,1,33),(146,24,4,1,20,1,230),(147,24,5,1,20,1,239),(148,25,1,1,20,3,264),(149,25,2,4,10,2,245),(150,25,3,5,10,2,246),(151,25,4,3,30,540,243),(153,25,5,0,30,500,244),(155,26,1,1,10,6,270),(156,26,2,1,20,6,264),(157,26,3,1,20,4,263),(158,26,4,3,25,540,243),(160,26,5,0,25,500,244),(161,27,1,2,20,1,66),(162,27,2,2,20,1,67),(163,27,3,2,20,1,68),(164,27,4,1,20,1,230),(166,27,5,1,20,1,239),(167,28,1,1,20,3,264),(168,28,2,4,10,2,245),(169,28,3,5,10,2,246),(170,28,4,3,30,572,243),(172,28,5,0,30,500,244),(174,29,1,1,10,6,270),(175,29,2,1,20,6,264),(176,29,3,1,20,4,263),(177,29,4,3,25,1122,243),(179,29,5,0,25,1000,244),(180,30,1,2,20,1,61),(181,30,2,2,20,1,62),(182,30,3,2,20,1,63),(184,30,4,2,20,1,64),(185,30,5,1,20,1,230),(186,31,1,0,20,1,244),(187,31,2,0,20,1,244),(188,31,3,0,20,1,244),(189,31,4,0,20,1,244),(190,31,5,0,20,1,244),(211,38,1,0,50,3000,244),(212,38,2,0,20,6000,244),(213,38,3,0,15,9000,244),(214,38,4,0,10,12000,244),(215,38,5,5,5,2,246),(216,39,1,3,50,1000,243),(217,39,2,3,20,1300,243),(218,39,3,3,15,1500,243),(219,39,4,3,10,1600,243),(220,39,5,4,5,2,245),(231,42,1,1,20,3,264),(232,42,2,4,10,2,245),(233,42,3,5,10,2,246),(234,42,4,3,30,230,243),(235,42,5,0,30,400,244),(236,43,1,1,20,3,270),(237,43,2,4,10,2,245),(238,43,3,5,10,2,246),(239,43,4,3,30,258,243),(240,43,5,0,30,400,244),(241,44,1,1,20,3,264),(242,44,2,4,10,2,245),(243,44,3,5,10,2,246),(244,44,4,3,30,540,243),(245,44,5,0,30,500,244),(246,45,1,1,20,3,264),(247,45,2,4,10,2,245),(248,45,3,5,10,2,246),(249,45,4,3,30,572,243),(250,45,5,0,30,500,244),(251,46,1,0,20,1,0),(252,46,2,0,20,1,0),(253,46,3,0,20,1,0),(254,46,4,0,20,1,0),(255,46,5,0,20,1,0),(256,51,1,0,50,6000,244),(257,51,2,0,20,9000,244),(258,51,3,0,15,12000,244),(259,51,4,0,10,15000,244),(260,51,5,5,5,2,246),(261,52,1,3,50,25400,243),(262,52,2,3,20,32700,243),(263,52,3,3,15,38100,243),(264,52,4,3,10,40000,243),(265,52,5,4,5,2,245),(266,47,1,0,50,4000,244),(267,47,2,0,20,7000,244),(268,47,3,0,15,10000,244),(269,47,4,0,10,13000,244),(270,47,5,5,5,2,246),(271,49,1,0,50,5000,244),(272,49,2,0,20,8000,244),(273,49,3,0,15,11000,244),(274,49,4,0,10,14000,244),(275,49,5,5,5,2,246),(276,53,1,0,50,7000,244),(277,53,2,0,20,10000,244),(278,53,3,0,15,13000,244),(279,53,4,0,10,16000,244),(280,53,5,5,5,2,246),(281,48,1,3,50,4600,243),(282,48,2,3,20,6000,243),(283,48,3,3,15,7000,243),(284,48,4,3,10,7500,243),(285,48,5,4,5,2,245),(286,55,1,0,50,8000,244),(287,55,2,0,20,11000,244),(288,55,3,0,15,14000,0),(289,55,4,0,10,17000,244),(290,55,5,5,5,2,246),(291,50,1,3,50,10800,243),(292,50,2,3,20,14000,243),(293,50,3,3,15,16200,243),(294,50,4,3,10,17000,243),(295,50,5,3,5,2,245),(296,57,1,0,50,9000,244),(297,57,2,0,20,12000,244),(298,57,3,0,15,15000,244),(299,57,4,0,10,18000,244),(300,57,5,5,5,2,246),(301,59,1,0,50,10000,244),(302,59,2,0,20,13000,244),(303,59,3,0,10,16000,244),(304,59,4,0,15,19000,244),(305,59,5,5,5,2,246),(306,54,1,3,50,38100,243),(307,54,2,3,20,49000,243),(308,54,3,3,15,57200,243),(309,54,4,3,10,60000,243),(310,54,5,4,5,2,245),(311,56,1,3,50,45200,243),(312,56,2,3,20,58100,243),(313,56,3,3,15,68000,243),(314,56,4,3,10,71100,243),(315,56,5,4,5,2,245),(316,58,1,3,50,55300,243),(317,58,2,3,20,71100,243),(318,58,3,3,15,83000,243),(319,58,4,3,10,87000,243),(320,58,5,4,5,2,245),(321,60,1,3,50,58300,243),(322,60,2,3,20,75000,243),(323,60,3,3,15,87500,243),(324,60,4,3,10,91600,243),(325,60,5,3,5,2,245),(331,61,1,0,50,11000,244),(332,61,2,0,20,14000,244),(333,61,3,0,15,17000,244),(334,61,4,0,10,20000,244),(335,61,5,0,5,2,246),(336,62,1,2,20,1,61),(337,62,2,2,20,1,62),(338,62,3,2,20,1,63),(339,62,4,1,20,1,271),(340,62,5,1,20,1,272),(341,63,1,2,20,1,66),(342,63,2,2,20,1,67),(343,63,3,2,20,1,68),(344,63,4,1,20,1,271),(345,63,5,1,20,1,272),(346,64,1,2,20,1,61),(347,64,2,2,20,1,64),(348,64,3,2,20,1,65),(349,64,4,1,20,1,271),(350,64,5,1,20,1,272),(351,65,1,2,20,1,66),(352,65,2,2,20,1,67),(353,65,3,2,20,1,68),(354,65,4,1,20,1,271),(355,65,5,1,20,1,272),(356,66,1,2,20,1,61),(357,66,2,2,20,1,62),(358,66,3,2,20,1,63),(359,66,4,1,20,1,271),(360,66,5,1,20,1,272),(371,70,1,1,30,1,302),(372,70,2,1,10,3,302),(373,70,3,1,20,1,256),(374,70,4,1,20,1,257),(375,70,5,1,20,1,258),(376,71,1,1,50,1,250),(377,71,2,1,20,2,250),(378,71,3,1,10,1,252),(379,71,4,1,10,1,253),(380,71,5,1,10,1,254),(381,72,1,1,30,5,263),(382,72,2,1,10,10,263),(383,72,3,1,20,1,259),(384,72,4,1,20,1,260),(385,72,5,1,20,1,261),(386,73,1,1,30,1,302),(387,73,2,1,10,3,302),(388,73,3,1,20,2,256),(389,73,4,1,20,2,257),(390,73,5,1,20,2,258),(391,74,1,1,30,1,302),(392,74,2,1,10,3,302),(393,74,3,1,20,2,256),(394,74,4,1,20,2,257),(395,74,5,1,20,2,258),(396,75,1,1,30,1,302),(397,75,2,1,10,3,302),(398,75,3,1,20,2,256),(399,75,4,1,20,2,257),(400,75,5,1,20,2,258),(401,79,1,1,50,1,250),(402,79,2,1,20,2,250),(403,79,3,1,10,1,252),(404,79,4,1,10,1,253),(405,79,5,1,10,1,254),(406,80,1,1,50,1,250),(407,80,2,1,20,2,250),(408,80,3,1,10,1,252),(409,80,4,1,10,1,253),(410,80,5,1,10,1,254),(411,81,1,1,50,1,250),(412,81,2,1,20,2,250),(413,81,3,1,10,1,252),(414,81,4,1,10,1,253),(415,81,5,1,10,1,254),(416,76,1,1,30,5,263),(417,76,2,1,10,10,263),(418,76,3,1,20,1,259),(419,76,4,1,20,1,260),(420,76,5,1,20,1,261),(421,77,1,1,30,5,263),(422,77,2,1,10,10,263),(423,77,3,1,20,1,259),(424,77,4,1,20,1,260),(425,77,5,1,20,1,261),(426,78,1,1,30,5,263),(427,78,2,1,10,10,263),(428,78,3,1,20,1,259),(429,78,4,1,20,1,260),(430,78,5,1,20,1,261);
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
INSERT INTO `mission_level_small_box` VALUES (7,7,1498,729,30,0,1),(11,13,1526,410,30,0,1),(16,11,926,564,30,1,1),(17,11,1373,956,30,0,1),(18,16,1121,469,30,0,1),(20,17,1402,1000,30,0,1),(21,17,722,964,30,0,1),(22,42,1684,1645,30,0,1),(23,23,2052,1093,30,0,1),(24,20,1000,1076,30,0,1),(25,43,966,1744,30,1,1),(26,25,510,701,30,0,1),(27,44,1124,378,30,0,1),(28,28,349,700,30,1,1),(29,45,1426,344,30,1,1),(30,26,1113,1024,30,0,1),(31,26,1467,861,30,0,1),(32,29,1113,1024,30,0,1),(33,29,1467,861,30,0,1);
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
INSERT INTO `mission_level_small_box_items` VALUES (2,1,31,20,0,0),(5,4,244,50,100,0),(6,4,244,25,200,0),(7,4,244,25,300,0),(8,7,244,30,100,0),(9,7,244,30,150,0),(10,7,244,30,200,0),(11,7,244,10,400,0),(24,11,244,30,100,0),(25,11,244,30,150,0),(26,11,244,30,200,0),(27,11,244,10,400,0),(40,16,244,30,100,0),(41,16,244,30,150,0),(42,16,244,30,200,0),(43,16,244,10,400,0),(44,17,244,30,100,0),(45,17,244,30,150,0),(46,17,244,30,200,0),(47,17,244,10,400,0),(52,18,244,30,100,0),(53,18,244,30,150,0),(54,18,244,30,200,0),(55,18,244,10,400,0),(60,20,244,30,100,0),(61,20,244,30,150,0),(62,20,244,30,200,0),(63,20,244,10,400,0),(64,21,244,30,100,0),(65,21,244,30,150,0),(66,21,244,30,200,0),(67,21,244,10,400,0),(68,22,244,30,100,0),(69,22,244,30,150,0),(70,22,244,30,200,0),(71,22,244,10,400,0),(72,23,244,30,100,0),(73,23,244,30,150,0),(74,23,244,30,200,0),(75,23,244,10,400,0),(76,24,244,30,100,0),(77,24,244,30,150,0),(78,24,244,30,200,0),(79,24,244,10,400,0),(80,25,244,30,100,0),(81,25,244,30,150,0),(82,25,244,30,200,0),(83,25,244,10,400,0),(84,26,244,30,100,0),(85,26,244,30,150,0),(86,26,244,30,200,0),(87,26,244,10,400,0),(88,27,244,30,100,0),(89,27,244,30,150,0),(90,27,244,30,200,0),(91,27,244,10,400,0),(92,28,244,30,100,0),(93,28,244,30,150,0),(94,28,244,30,200,0),(95,28,244,10,400,0),(96,29,244,30,100,0),(97,29,244,30,150,0),(98,29,244,30,200,0),(99,29,244,10,400,0),(100,33,244,30,100,0),(101,33,244,30,150,0),(102,33,244,30,200,0),(103,33,244,10,400,0),(104,32,244,30,100,0),(105,32,244,30,150,0),(106,32,244,30,200,0),(107,32,244,10,400,0),(108,31,244,30,100,0),(109,31,244,30,150,0),(110,31,244,30,200,0),(111,31,244,10,400,0),(112,30,244,30,100,0),(113,30,244,30,150,0),(114,30,244,30,200,0),(115,30,244,10,400,0);
/*!40000 ALTER TABLE `mission_level_small_box_items` ENABLE KEYS */;
DROP TABLE IF EXISTS `mission_talk`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mission_talk` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '对话ID',
  `enemy_id` int(11) NOT NULL COMMENT '副本敌人ID',
  `content` varchar(1024) DEFAULT '' COMMENT '对话内容',
  `town_id` smallint(6) NOT NULL COMMENT '城镇ID',
  `quest_id` smallint(6) NOT NULL COMMENT '任务',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COMMENT='副本战场对话';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `mission_talk` DISABLE KEYS */;
INSERT INTO `mission_talk` VALUES (3,93,'MP=媛媛别闹了，你怎么了\r\nMB=不开心，超生气的\r\nMP=冷静，你被阴影影响了\r\nMB=要你多管\r\n\r\nWP=媛媛，你冷静一点！\r\nWB=我…烦死了！\r\nWP=你被阴影影响了！\r\nWB=要你多管！',1,10),(5,27,'MP=媛媛，注意辅助我\r\nMZ=嗯\r\n\r\nWP=媛媛，注意辅助我\r\nWP=嗯',1,16),(6,45,'MB=让我生气后果很严重\r\nMB=小的们，给我上\r\nMP=小心，金蟾王已经被阴影控制了\r\n\r\nWB=让我生气后果很严重\r\nWB=小的们，给我上\r\nWP=小心，金蟾王已经被阴影控制了',1,26),(7,54,'MZ=烂泥都能成精，这世界没道理了\r\nMB=你看起来最好吃\r\nMB=我一定留到最后再吃\r\nMZ=做梦\r\n\r\nWZ=烂泥都能成精，这世界没道理了\r\nWB=小姑娘，你看起来最好吃\r\nWB=我一定留到最后再吃\r\nWZ=做梦',1,30),(8,63,'MB=小的们，把他们烤熟了带回去\r\nMY=大白天的打灯笼上厕所\r\nMB=什么意思？\r\nMY=找死啊\r\n\r\nWB=小的们，把他们烤熟了带回去\r\nWY=大白天的打灯笼上厕所\r\nWB=什么意思？\r\nWY=找死啊',1,61),(9,72,'MB=我要把你们烧成灰烬\r\nMP=变成灰烬前能问个问题吗？\r\nMB=问吧，可怜虫\r\nMP=最近有人来过这里吗？\r\nMB=没有，你们是第一个\r\nMP=好吧，那我们可以走吗？\r\nMB=不可以，受死吧\r\n\r\nWB=我要把你们烧成灰烬\r\nWP=变成灰烬前能问个问题吗？\r\nWB=问吧，可怜虫\r\nWP=最近有人来过这里吗？\r\nWB=没有，你们是第一个\r\nWP=好吧，那我们可以走吗？\r\nWB=不可以，受死吧',1,39),(10,87,'MB=想过留剑\r\nMP=没门\r\n\r\nWB=想过留剑\r\nWP=没门',1,63),(11,90,'MB=年轻人，用你手中的剑\r\nMB=证明你的剑道吧\r\nMP=得罪了，前辈\r\n\r\nWB=年轻人，用你手中的剑\r\nWB=证明你的剑道吧\r\nWP=得罪了，前辈',1,46);
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
INSERT INTO `multi_level` VALUES (1,'QingZhuLin8','','刀疤兔 Lv.30',1,5000,5000,61,1,62,1,63,1,0,100110),(2,'HeiYeSenLin4','','天狼妖 Lv.35',35,5000,5000,61,1,64,1,65,1,100110,100120),(3,'AnYingMiJing4','','妖龙 Lv.40',40,5000,5000,66,1,67,1,68,1,100120,100130),(4,'LianHuaFeng4','','金蟾王 Lv.45',45,5000,5000,215,1,216,1,217,1,100130,100140),(5,'ShuiRongDong4','','剧毒臭泥 Lv.50',50,5000,5000,213,1,211,1,214,1,100140,100150),(6,'RongYanHuoShan4','','燃魁首领 Lv.55',55,5000,5000,216,1,214,1,215,1,100150,100160),(7,'RongYanHuoShan4','','炎龙 Lv.60',60,5000,5000,211,1,215,1,213,1,100160,100170),(8,'JianLingMiShi4','','古代剑灵 Lv.65',65,5000,5000,217,1,211,1,216,1,110170,110180);
/*!40000 ALTER TABLE `multi_level` ENABLE KEYS */;
DROP TABLE IF EXISTS `npc_talk`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `npc_talk` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '记录ID',
  `npc_id` int(11) NOT NULL COMMENT 'NPC ID',
  `town_id` smallint(6) NOT NULL COMMENT '相关城镇',
  `type` tinyint(4) NOT NULL COMMENT '对话类型 1--首次对话； 2--人物对话',
  `quest_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '关联主线任务',
  `conversion` varchar(1024) NOT NULL COMMENT '对话内容',
  `award_item_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品ID',
  `award_item_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8mb4 COMMENT='城镇NPC对话';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `npc_talk` DISABLE KEYS */;
INSERT INTO `npc_talk` VALUES (2,2,1,0,0,'N=我捡到了这个，送你\r\nP=谢谢，我就收下啦',264,1),(3,1,1,1,2,'P=仙女姐姐，青竹林怎么去呢？\r\nN=点右上角的冒险按钮就可以看到了！\r\nP=谢谢，我知道了。\r\nN=路上小心！',0,0),(10,1,1,1,9,'P=仙女姐姐，青竹林怎么去呢？\r\nN=点右上角的冒险按钮就可以看到了！\r\nP=谢谢，我知道了。\r\nN=路上小心！',0,0),(11,1,1,1,13,'N＝黑夜森林有恶狼出没，梦妖们都被吓坏了。',0,0),(12,1,1,1,15,'N＝听梦妖说，好像有伙伴出现在森林。',0,0),(13,1,1,1,16,'N＝阴影附身后的伙伴会变的比平时更强           P＝勇者无畏，打过才知道。',0,0),(14,1,1,1,19,'N＝阴影的气息更加强烈了，请小心                 P＝谢谢，我会注意的',0,0),(15,1,1,1,20,'N＝这里有一股气息十分的厉害，请小心！',0,0),(16,1,1,1,23,'N＝莲花峰十分的美丽，但是请小心阴影的袭击',0,0),(17,1,1,1,25,'N＝莲花峰的金蟾王很善良，但是千万别说它肥哦。',0,0),(18,1,1,1,26,'N＝要击杀附体的阴影只要打败被附体的人就可以哦',0,0),(19,1,1,1,27,'N＝据梦妖们说在水溶洞附近能闻到阵阵恶臭',0,0),(20,1,1,1,28,'N＝山洞中越来越臭了，不舒服时记得回来休息下',0,0),(21,1,1,1,29,'N＝在山洞的深处有一股阴影之源，请小心',0,0),(22,1,1,1,30,'N=臭泥除里非常臭以外还带有剧毒哦',0,0),(23,1,1,1,31,'N＝梦妖们正在很努力的寻找其余的伙伴们',0,0),(24,1,1,1,33,'N＝出去的梦妖们打听到火山附近有伙伴的行踪',0,0),(25,1,1,1,34,'N＝火山中有些怪物是很有灵性的',0,0),(26,1,1,1,35,'N＝请小心燃魁首领',0,0),(27,1,1,1,61,'N＝请小心燃魁首领',0,0),(28,1,1,1,62,'N＝火山的深处有一只凶暴的炎龙，请小心',0,0),(29,1,1,1,38,'N＝火山的深处有一只凶暴的炎龙，请小心',0,0),(30,1,1,1,39,'N＝凶暴的炎龙讨厌凡人进入它的领地，所以请小心！',0,0),(31,1,1,1,40,'N＝密室很暗，请小心守护者',0,0),(32,1,1,1,41,'N＝在密室中拥有很多充满灵性的飞剑',0,0),(33,1,1,1,59,'N＝密室中的飞剑都是上古流传下来的',0,0),(34,1,1,1,42,'N＝请小心守护密室的剑魂',0,0),(35,1,1,1,63,'N＝上古的剑魂是非常厉害的',0,0),(36,1,1,1,43,'N＝野外很危险，请小心！',0,0),(37,1,1,1,44,'N＝可以和密室中的上古飞剑签订契约',0,0),(38,1,1,1,60,'N＝可以和密室中的上古飞剑签订契约',0,0),(39,1,1,1,45,'N＝密室的深处感觉到一道剑气',0,0),(40,1,1,1,46,'N＝古代剑灵很厉害，但是能获得很好的奖励',0,0);
/*!40000 ALTER TABLE `npc_talk` ENABLE KEYS */;
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
) ENGINE=InnoDB AUTO_INCREMENT=89 DEFAULT CHARSET=utf8mb4 COMMENT='主线任务';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `quest` DISABLE KEYS */;
INSERT INTO `quest` VALUES (1,10,'侠客降临',0,'降临侠客岛。',0,1,0,0,0,0,1,0,10,0,0,0,0,0,0,0,0,0,0,0,100100,0,0,0,0),(2,20,'探索青竹林一',2,'探索青竹林一。',0,0,0,4,0,0,1,0,10,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(9,30,'探索青竹林二',2,'探索青竹林二。',0,0,0,3,0,0,1,0,20,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(10,40,'阴影附身',4,'媛媛被阴影侵蚀了，快去救她。',0,0,0,3,93,111,1,30,2000,0,0,0,0,0,0,0,0,0,4,2,0,0,30,0,1),(12,100,'返回神龙岛',0,'先回神龙岛再说了。',0,1,0,0,0,0,1,0,30,264,1,0,0,0,0,0,0,900,0,0,0,0,0,0,0),(13,110,'探索黑夜森林一',2,'探索黑夜森林一。',0,0,0,7,0,0,1,0,30,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(15,130,'探索黑夜森林二',2,'探索黑夜森林二。',0,0,0,9,0,0,1,0,30,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(16,140,'伙伴阴影',4,'袁铭志也会被阴影侵蚀？不可能吧！',0,0,0,9,27,112,1,50,3000,264,10,0,0,0,0,0,0,0,3,5,0,0,30,0,1),(17,150,'阴影谜团',0,'把好消息告诉仙女姐姐吧！',0,1,0,0,0,0,1,0,500,31,1,231,1,235,1,0,0,0,0,0,0,0,0,0,0),(19,170,'探索暗影秘境一',2,'探索暗影秘境一。',0,0,0,11,0,0,1,0,40,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(20,180,'探索暗影秘境二',2,'探索暗影秘境二。',0,0,0,12,0,0,1,0,40,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(21,190,'魂力爆发',4,'影龙为什么又复活了。',0,0,0,12,36,25,1,60,4000,264,10,0,0,0,0,0,0,1000,0,0,0,0,30,0,1),(22,200,'询问魂侍',0,'魂侍到底是什么东西。',0,1,0,0,0,0,1,0,50,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(23,210,'探索莲花峰一',2,'探索莲花峰一。',0,0,0,13,0,0,1,0,50,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(25,230,'探索莲花峰二',2,'探索莲花峰二。',0,0,0,15,0,0,1,0,50,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(26,240,'净化金蟾王',4,'帮金蟾王清除阴影。',0,0,0,15,45,11,1,70,5000,264,10,0,0,0,0,0,0,1300,0,0,0,0,30,0,1),(27,250,'探索水溶洞一',2,'探索水溶洞一。',0,0,0,16,0,0,1,0,60,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(28,260,'探索水溶洞二',2,'探索水溶洞二。',0,0,0,17,0,0,1,0,60,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(29,270,'探索水溶洞三',2,'探索水溶洞三。',0,0,0,18,0,0,1,0,60,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(30,280,'清除剧毒臭泥',4,'这堆东西太臭了，赶紧清理了吧！',0,0,0,18,54,14,1,80,10000,264,10,0,0,0,0,0,0,0,0,0,0,0,40,0,1),(31,290,'伙伴消息',0,'回去看看仙女姐姐有没有什么新消息。',0,1,0,0,0,0,1,0,100,0,0,0,0,0,0,0,0,2000,0,0,0,0,0,0,0),(33,300,'探索熔岩火山一',2,'探索熔岩火山一。',0,0,0,19,0,0,1,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(34,310,'探索熔岩火山二',2,'探索熔岩火山二。',0,0,0,20,0,0,1,0,100,251,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(35,320,'探索熔岩火山四',2,'探索熔岩火山四。',0,0,0,21,0,0,1,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(36,330,'探索熔岩火山五',3,'探索熔岩火山五。',0,0,0,22,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(37,340,'探索熔岩火山六',2,'探索熔岩火山六。',0,0,0,23,0,0,1,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(38,350,'探索熔岩火山八',2,'探索熔岩火山八。',0,0,0,24,0,0,1,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(39,360,'暴怒的炎龙',4,'这炎龙似乎误会我们了，没办法解释了。',0,0,0,24,72,19,1,440,30000,264,10,0,0,0,0,0,0,0,0,0,0,0,100,0,1),(40,370,'探索剑灵密室一',2,'探索剑灵密室一。',0,0,0,25,0,0,1,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(41,380,'探索剑灵密室二',2,'探索剑灵密室二。',0,0,0,26,0,0,1,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(42,390,'探索剑灵密室四',2,'探索剑灵密室四。',0,0,0,27,0,0,1,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(43,400,'探索剑灵密室五',3,'探索剑灵密室五。',0,0,0,28,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(44,410,'探索剑灵密室六',2,'探索剑灵密室六。',0,0,0,29,0,0,1,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(45,420,'探索剑灵密室八',2,'探索剑灵密室八。',0,0,0,30,0,0,1,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(46,430,'密室古剑',4,'接受古代剑灵的考验吧！',0,0,0,30,90,23,1,1000,50000,264,10,0,0,0,0,0,0,1500,0,0,0,0,0,0,1),(47,440,'前往烛堡',0,'前往烛堡寻找龙姬。',0,2,0,0,0,0,1,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(57,315,'探索熔岩火山三',3,'探索熔岩火山三。',0,0,0,42,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(58,345,'探索熔岩火山七',3,'探索熔岩火山七。',0,0,0,43,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(59,385,'探索剑灵密室三',3,'探索剑灵密室三。',0,0,0,44,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(60,415,'探索剑灵密室七',3,'探索剑灵密室七。',0,0,0,45,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(61,324,'打败燃魁首领',4,'好暴躁的燃魁首领。',0,0,0,21,63,18,1,200,20000,264,10,0,0,0,0,0,0,0,0,0,0,0,0,0,1),(62,328,'火山深处',0,'太热了，先回神龙岛修整一下。',0,1,0,0,0,0,1,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(63,395,'守剑之魂',4,'大侠手中的剑不是谁让放下就放下的。',0,0,0,27,87,101,1,600,40000,264,10,0,0,0,0,0,0,0,0,0,0,0,0,0,1),(64,435,'龙姬被掳',0,'糟糕，好像是中了调虎离山之计了。',0,1,0,0,0,0,1,0,100,0,0,0,0,0,0,0,0,0,0,0,150100,100110,0,0,0),(65,204,'媛媛的担心',1,'媛媛看起来有些担心的样子。',0,1,11,0,0,0,1,0,50,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(66,206,'找袁铭志喝酒',1,'袁铭志又在喝酒了。',0,1,12,0,0,0,1,0,50,0,0,0,0,0,0,0,0,0,0,0,120100,0,0,0,0),(67,208,'救援莲花峰',1,'梦妖急急忙忙的，不知道出了什么事。',0,1,1,0,0,0,0,0,50,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(68,25,'通关青竹林一',3,'探索青竹林一。',0,0,0,4,0,0,0,0,10,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(69,45,'通关青竹林二',3,'探索青竹林二。',0,0,0,3,0,0,0,0,20,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(70,115,'通关黑夜森林一',3,'探索黑夜森林一。',0,0,0,7,0,0,0,0,30,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(71,145,'通关黑夜森林二',3,'探索黑夜森林二。',0,0,0,9,0,0,0,0,30,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(72,170,'通关暗影秘境一',3,'探索暗影秘境一。',0,0,0,11,0,0,0,0,40,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(73,195,'通关暗影秘境二',3,'探索暗影秘境二。',0,0,0,12,0,0,0,0,40,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(74,215,'通关莲花峰一',3,'探索莲花峰一。',0,0,0,13,0,0,0,0,50,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(75,245,'通关莲花峰二',3,'探索莲花峰二。',0,0,0,15,0,0,0,0,60,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(76,255,'通关水溶洞一',3,'探索水溶洞一。',0,0,0,16,0,0,0,0,60,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(77,265,'通关水溶洞二',3,'探索水溶洞二。',0,0,0,17,0,0,0,0,60,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(78,285,'通关水溶洞三',3,'探索水溶洞三。',0,0,0,18,0,0,0,0,60,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(79,305,'通关熔岩火山一',3,'探索熔岩火山一。',0,0,0,19,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(80,312,'通关熔岩火山二',3,'探索熔岩火山二。',0,0,0,20,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(81,326,'通关熔岩火山四',3,'探索熔岩火山四。',0,0,0,21,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(82,342,'通关熔岩火山六',3,'探索熔岩火山六。',0,0,0,23,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(83,365,'通关熔岩火山八',3,'探索熔岩火山八。',0,0,0,24,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(84,375,'通关剑灵密室一',3,'探索剑灵密室一。',0,0,0,25,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(85,382,'通关剑灵密室二',3,'探索剑灵密室二。',0,0,0,26,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(86,398,'通关剑灵密室四',3,'探索剑灵密室四。',0,0,0,27,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(87,412,'通关剑灵密室六',3,'探索剑灵密室六。',0,0,0,29,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0),(88,432,'通关剑灵密室八',3,'探索剑灵密室八。',0,0,0,30,0,0,0,0,100,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0);
/*!40000 ALTER TABLE `quest` ENABLE KEYS */;
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
INSERT INTO `role` VALUES (1,'义峰','YiFeng',1,0,2,0,0,0,100),(2,'昕苒','XinRan',1,0,2,0,0,0,100),(3,'袁铭志','YuanMingZhi',2,0,1037,6,0,0,100),(4,'朱媛媛','ZhuYuanYuan',2,0,7,8,0,0,100),(5,'车晓芸','CheXiaoYun',2,0,49,50,0,0,100),(6,'燕无名','YanWuMing',2,0,51,52,0,0,100);
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
INSERT INTO `role_level` VALUES (11,1,1,200,40,15,10,80,15,6,0,6,0,0,10,2,200,100,200,2,2),(12,1,2,300,70,25,20,110,15,6,0,6,0,0,10,2,200,100,200,2,2),(13,1,3,400,100,35,30,140,15,6,0,6,0,0,10,2,200,100,200,2,2),(14,1,4,500,130,45,40,170,15,6,0,6,0,0,10,2,200,100,200,2,2),(15,1,5,600,160,55,50,200,15,6,0,6,0,0,10,2,200,100,200,2,2),(16,1,6,700,190,65,60,230,15,6,0,6,0,0,10,2,200,100,200,2,2),(17,1,7,800,220,75,70,260,15,6,0,6,0,0,10,2,200,100,200,2,2),(18,1,8,900,250,85,80,290,15,6,0,6,0,0,10,2,200,100,200,2,2),(19,1,9,1000,280,95,90,320,15,6,0,6,0,0,10,2,200,100,200,2,2),(20,1,10,1100,310,105,100,350,15,6,0,6,0,0,10,2,200,100,200,2,2),(21,1,11,1210,343,116,111,383,15,6,0,6,0,0,10,2,200,100,200,2,2),(22,1,12,1320,376,127,122,416,15,6,0,6,0,0,10,2,200,100,200,2,2),(23,1,13,1430,409,138,133,449,15,6,0,6,0,0,10,2,200,100,200,2,2),(24,1,14,1540,442,149,144,482,15,6,0,6,0,0,10,2,200,100,200,2,2),(25,1,15,1650,475,160,155,515,15,6,0,6,0,0,10,2,200,100,200,2,2),(26,1,16,1760,508,171,166,548,15,6,0,6,0,0,10,2,200,100,200,2,2),(27,1,17,1870,541,182,177,581,15,6,0,6,0,0,10,2,200,100,200,2,2),(28,1,18,1980,574,193,188,614,15,6,0,6,0,0,10,2,200,100,200,2,2),(29,1,19,2090,607,204,199,647,15,6,0,6,0,0,10,2,200,100,200,2,2),(30,1,20,2200,640,215,210,680,15,6,0,6,0,0,10,2,200,100,200,2,2),(31,1,21,2320,676,227,222,716,15,6,0,6,0,0,10,2,250,100,200,2,2),(32,1,22,2440,712,239,234,752,15,6,0,6,0,0,10,2,250,100,200,2,2),(33,1,23,2560,748,251,246,788,15,6,0,6,0,0,10,2,250,100,200,2,2),(34,1,24,2680,784,263,258,824,15,6,0,6,0,0,10,2,250,100,200,2,2),(35,1,25,2800,820,275,270,860,15,6,0,6,0,0,10,2,250,100,200,2,2),(36,1,26,2920,856,287,282,896,15,6,0,6,0,0,10,2,250,100,200,2,2),(37,1,27,3040,892,299,294,932,15,6,0,6,0,0,10,2,250,100,200,2,2),(38,1,28,3160,928,311,306,968,15,6,0,6,0,0,10,2,250,100,200,2,2),(39,1,29,3280,964,323,318,1004,15,6,0,6,0,0,10,2,250,100,200,2,2),(40,1,30,3400,1000,335,330,1040,15,6,0,6,0,0,10,2,250,100,200,2,2),(41,1,31,3530,1039,348,343,1079,15,6,0,6,0,0,10,2,300,100,200,2,2),(42,1,32,3660,1078,361,356,1118,15,6,0,6,0,0,10,2,300,100,200,2,2),(43,1,33,3790,1117,374,369,1157,15,6,0,6,0,0,10,2,300,100,200,2,2),(44,1,34,3920,1156,387,382,1196,15,6,0,6,0,0,10,2,300,100,200,2,2),(45,1,35,4050,1195,400,395,1235,15,6,0,6,0,0,10,2,300,300,200,2,2),(46,1,36,4180,1234,413,408,1274,15,6,0,6,0,0,10,2,300,300,200,2,2),(47,1,37,4310,1273,426,421,1313,15,6,0,6,0,0,10,2,300,300,200,2,2),(48,1,38,4440,1312,439,434,1352,15,6,0,6,0,0,10,2,300,300,200,2,2),(49,1,39,4570,1351,452,447,1391,15,6,0,6,0,0,10,2,300,300,200,2,2),(50,1,40,4700,1390,465,460,1430,15,6,0,6,0,0,10,2,300,300,200,2,2),(51,1,41,4840,1432,479,474,1472,15,6,0,6,0,0,10,2,350,300,200,2,2),(52,1,42,4980,1474,493,488,1514,15,6,0,6,0,0,10,2,350,300,200,2,2),(53,1,43,5120,1516,507,502,1556,15,6,0,6,0,0,10,2,350,100,200,2,2),(54,1,44,5260,1558,521,516,1598,15,6,0,6,0,0,10,2,350,100,200,2,2),(64,1,45,5400,1600,535,530,1640,15,6,0,6,0,0,10,2,350,100,200,2,2),(65,1,46,5540,1642,549,544,1682,15,6,0,6,0,0,10,2,350,100,200,2,2),(66,1,47,5680,1684,563,558,1724,15,6,0,6,0,0,10,2,350,100,200,2,2),(67,1,48,5820,1726,577,572,1766,15,6,0,6,0,0,10,2,350,100,200,2,2),(68,1,49,5960,1768,591,586,1808,15,6,0,6,0,0,10,2,350,100,200,2,2),(69,1,50,6100,1810,605,600,1850,15,6,0,6,0,0,10,2,350,100,200,2,2),(70,1,51,6250,1855,620,615,1895,15,6,0,6,0,0,10,2,400,100,200,2,2),(71,3,1,800,130,120,20,20,6,6,0,0,0,0,0,0,200,100,200,2,2),(72,3,2,1100,160,140,30,30,6,6,0,0,0,0,0,0,200,100,200,2,2),(73,3,3,1400,190,160,40,40,6,6,0,0,0,0,0,0,200,100,200,2,2),(74,4,1,200,40,20,130,30,6,6,0,6,0,0,0,0,200,100,200,2,2),(75,4,2,300,70,30,160,50,6,6,0,6,0,0,0,0,200,100,200,2,2),(76,4,3,400,100,40,190,70,6,6,0,6,0,0,0,0,200,100,200,2,2),(77,1,52,6400,1900,635,630,1940,15,6,0,6,0,0,10,2,400,100,200,2,2),(78,1,53,6550,1945,650,645,1985,15,6,0,6,0,0,10,2,400,100,200,2,2),(79,1,54,6700,1990,665,660,2030,15,6,0,6,0,0,10,2,400,100,200,2,2),(80,1,55,6850,2035,680,675,2075,15,6,0,6,0,0,10,2,400,100,200,2,2),(81,1,56,7000,2080,695,690,2120,15,6,0,6,0,0,10,2,400,100,200,2,2),(82,1,57,7150,2125,710,705,2165,15,6,0,6,0,0,10,2,400,100,200,2,2),(83,1,58,7300,2170,725,720,2210,15,6,0,6,0,0,10,2,400,100,200,2,2),(84,1,59,7450,2215,740,735,2255,15,6,0,6,0,0,10,2,400,100,200,2,2),(85,1,60,7600,2260,755,750,2300,15,6,0,6,0,0,10,2,400,100,200,2,2),(86,1,61,7760,2308,771,766,2348,15,6,0,6,0,0,10,2,400,100,200,2,2),(87,1,62,7920,2356,787,782,2396,15,6,0,6,0,0,10,2,400,100,200,2,2),(88,1,63,8080,2404,803,798,2444,15,6,0,6,0,0,10,2,400,100,200,2,2),(89,1,64,8240,2452,819,814,2492,15,6,0,6,0,0,10,2,400,100,200,2,2),(90,1,65,8400,2500,835,830,2540,15,6,0,6,0,0,10,2,400,100,200,2,2),(91,1,66,8560,2548,851,846,2588,15,6,0,6,0,0,10,2,400,100,200,2,2),(92,1,67,8720,2596,867,862,2636,15,6,0,6,0,0,10,2,400,100,200,2,2),(93,1,68,8880,2644,883,878,2684,15,6,0,6,0,0,10,2,400,100,200,2,2),(94,1,69,9040,2692,899,894,2732,15,6,0,6,0,0,10,2,400,100,200,2,2),(95,1,70,9200,2740,915,910,2780,15,6,0,6,0,0,10,2,400,100,200,2,2),(96,1,71,9370,2791,932,927,2831,15,6,0,6,0,0,10,2,400,100,200,2,2),(97,1,72,9540,2842,949,944,2882,15,6,0,6,0,0,10,2,400,100,200,2,2),(98,1,73,9710,2893,966,961,2933,15,6,0,6,0,0,10,2,400,100,200,2,2),(99,1,74,9880,2944,983,978,2984,15,6,0,6,0,0,10,2,400,100,200,2,2),(100,1,75,10050,2995,1000,995,3035,15,6,0,6,0,0,10,2,400,100,200,2,2),(101,1,76,10220,3046,1017,1012,3086,15,6,0,6,0,0,10,2,400,100,200,2,2),(102,1,77,10390,3097,1034,1029,3137,15,6,0,6,0,0,10,2,400,100,200,2,2),(103,1,78,10560,3148,1051,1046,3188,15,6,0,6,0,0,10,2,400,100,200,2,2),(104,1,79,10730,3199,1068,1063,3239,15,6,0,6,0,0,10,2,400,100,200,2,2),(105,1,80,10900,3250,1085,1080,3290,15,6,0,6,0,0,10,2,400,100,200,2,2),(106,1,81,11080,3304,1103,1098,3344,15,6,0,6,0,0,10,2,400,100,200,2,2),(107,1,82,11260,3358,1121,1116,3398,15,6,0,6,0,0,10,2,400,100,200,2,2),(108,1,83,11440,3412,1139,1134,3452,15,6,0,6,0,0,10,2,400,100,200,2,2),(109,1,84,11620,3466,1157,1152,3506,15,6,0,6,0,0,10,2,400,100,200,2,2),(110,1,85,11800,3520,1175,1170,3560,15,6,0,6,0,0,10,2,400,100,200,2,2),(111,1,86,11980,3574,1193,1188,3614,15,6,0,6,0,0,10,2,400,100,200,2,2),(112,1,87,12160,3628,1211,1206,3668,15,6,0,6,0,0,10,2,400,100,200,2,2),(113,1,88,12340,3682,1229,1224,3722,15,6,0,6,0,0,10,2,400,100,200,2,2),(114,1,89,12520,3736,1247,1242,3776,15,6,0,6,0,0,10,2,400,100,200,2,2),(115,1,90,12700,3790,1265,1260,3830,15,6,0,6,0,0,10,2,400,100,200,2,2),(116,1,91,12890,3847,1284,1279,3887,15,6,0,6,0,0,10,2,400,100,200,2,2),(117,1,92,13080,3904,1303,1298,3944,15,6,0,6,0,0,10,2,400,100,200,2,2),(118,1,93,13270,3961,1322,1317,4001,15,6,0,6,0,0,10,2,400,100,200,2,2),(119,1,94,13460,4018,1341,1336,4058,15,6,0,6,0,0,10,2,400,100,200,2,2),(120,1,95,13650,4075,1360,1355,4115,15,6,0,6,0,0,10,2,400,100,200,2,2),(121,1,96,13840,4132,1379,1374,4172,15,6,0,6,0,0,10,2,400,100,200,2,2),(122,1,97,14030,4189,1398,1393,4229,15,6,0,6,0,0,10,2,400,100,200,2,2),(123,1,98,14220,4246,1417,1412,4286,15,6,0,6,0,0,10,2,400,100,200,2,2),(124,1,99,14410,4303,1436,1431,4343,15,6,0,6,0,0,10,2,400,100,200,2,2),(125,1,100,14600,4360,1455,1450,4400,15,6,0,6,0,0,10,2,400,100,200,2,2),(126,3,4,1700,220,180,50,50,6,6,0,0,0,0,0,0,200,100,200,2,2),(127,3,5,2000,250,200,60,60,6,6,0,0,0,0,0,0,200,100,200,2,2),(128,3,6,2300,280,220,70,70,6,6,0,0,0,0,0,0,200,100,200,2,2),(129,3,7,2600,310,240,80,80,6,6,0,0,0,0,0,0,200,100,200,2,2),(130,3,8,2900,340,260,90,90,6,6,0,0,0,0,0,0,200,100,200,2,2),(131,3,9,3200,370,280,100,100,6,6,0,0,0,0,0,0,200,100,200,2,2),(132,3,10,3500,400,300,110,110,6,6,0,0,0,0,0,0,200,100,200,2,2),(133,3,11,3830,433,322,121,121,6,6,0,0,0,0,0,0,200,100,200,2,2),(134,3,12,4160,466,344,132,132,6,6,0,0,0,0,0,0,200,100,200,2,2),(135,3,13,4490,499,366,143,143,6,6,0,0,0,0,0,0,200,100,200,2,2),(136,3,14,4820,532,388,154,154,6,6,0,0,0,0,0,0,200,100,200,2,2),(137,3,15,5150,565,410,165,165,6,6,0,0,0,0,0,0,200,100,200,2,2),(138,3,16,5480,598,432,176,176,6,6,0,0,0,0,0,0,200,100,200,2,2),(139,3,17,5810,631,454,187,187,6,6,0,0,0,0,0,0,200,100,200,2,2),(140,3,18,6140,664,476,198,198,6,6,0,0,0,0,0,0,200,100,200,2,2),(141,3,19,6470,697,498,209,209,6,6,0,0,0,0,0,0,200,100,200,2,2),(142,3,20,6800,730,520,220,220,6,6,0,0,0,0,0,0,200,100,200,2,2),(143,3,21,7160,766,544,232,232,6,6,0,0,0,0,0,0,250,100,200,2,2),(144,3,22,7520,802,568,244,244,6,6,0,0,0,0,0,0,250,100,200,2,2),(145,3,23,7880,838,592,256,256,6,6,0,0,0,0,0,0,250,100,200,2,2),(146,3,24,8240,874,616,268,268,6,6,0,0,0,0,0,0,250,100,200,2,2),(147,3,25,8600,910,640,280,280,6,6,0,0,0,0,0,0,250,100,200,2,2),(148,3,26,8960,946,664,292,292,6,6,0,0,0,0,0,0,250,100,200,2,2),(149,3,27,9320,982,688,304,304,6,6,0,0,0,0,0,0,250,100,200,2,2),(150,3,28,9680,1018,712,316,316,6,6,0,0,0,0,0,0,250,100,200,2,2),(151,3,29,10040,1054,736,328,328,6,6,0,0,0,0,0,0,250,100,200,2,2),(152,3,30,10400,1090,760,340,340,6,6,0,0,0,0,0,0,250,100,200,2,2),(153,3,31,10790,1129,786,353,353,6,6,0,0,0,0,0,0,300,100,200,2,2),(154,3,32,11180,1168,812,366,366,6,6,0,0,0,0,0,0,300,100,200,2,2),(155,3,33,11570,1207,838,379,379,6,6,0,0,0,0,0,0,300,100,200,2,2),(156,3,34,11960,1246,864,392,392,6,6,0,0,0,0,0,0,300,100,200,2,2),(157,3,35,12350,1285,890,405,405,6,6,0,0,0,0,0,0,300,100,200,2,2),(158,3,36,12740,1324,916,418,418,6,6,0,0,0,0,0,0,300,100,200,2,2),(159,3,37,13130,1363,942,431,431,6,6,0,0,0,0,0,0,300,100,200,2,2),(160,3,38,13520,1402,968,444,444,6,6,0,0,0,0,0,0,300,100,200,2,2),(161,3,39,13910,1441,994,457,457,6,6,0,0,0,0,0,0,300,100,200,2,2),(162,3,40,14300,1480,1020,470,470,6,6,0,0,0,0,0,0,300,100,200,2,2),(163,3,41,14720,1522,1048,484,484,6,6,0,0,0,0,0,0,350,100,200,2,2),(164,3,42,15140,1564,1076,498,498,6,6,0,0,0,0,0,0,350,100,200,2,2),(165,3,43,15560,1606,1104,512,512,6,6,0,0,0,0,0,0,350,100,200,2,2),(166,3,44,15980,1648,1132,526,526,6,6,0,0,0,0,0,0,350,100,200,2,2),(167,3,45,16400,1690,1160,540,540,6,6,0,0,0,0,0,0,350,100,200,2,2),(168,3,46,16820,1732,1188,554,554,6,6,0,0,0,0,0,0,350,100,200,2,2),(169,3,47,17240,1774,1216,568,568,6,6,0,0,0,0,0,0,350,100,200,2,2),(170,3,48,17660,1816,1244,582,582,6,6,0,0,0,0,0,0,350,100,200,2,2),(171,3,49,18080,1858,1272,596,596,6,6,0,0,0,0,0,0,350,100,200,2,2),(172,3,50,18500,1900,1300,610,610,6,6,0,0,0,0,0,0,350,100,200,2,2),(173,3,51,18950,1945,1330,625,625,6,6,0,0,0,0,0,0,400,100,200,2,2),(174,3,52,19400,1990,1360,640,640,6,6,0,0,0,0,0,0,400,100,200,2,2),(175,3,53,19850,2035,1390,655,655,6,6,0,0,0,0,0,0,400,100,200,2,2),(176,3,54,20300,2080,1420,670,670,6,6,0,0,0,0,0,0,400,100,200,2,2),(177,3,55,20750,2125,1450,685,685,6,6,0,0,0,0,0,0,400,100,200,2,2),(178,3,56,21200,2170,1480,700,700,6,6,0,0,0,0,0,0,400,100,200,2,2),(179,3,57,21650,2215,1510,715,715,6,6,0,0,0,0,0,0,400,100,200,2,2),(180,3,58,22100,2260,1540,730,730,6,6,0,0,0,0,0,0,400,100,200,2,2),(181,3,59,22550,2305,1570,745,745,6,6,0,0,0,0,0,0,400,100,200,2,2),(182,3,60,23000,2350,1600,760,760,6,6,0,0,0,0,0,0,400,100,200,2,2),(183,3,61,23480,2398,1632,776,776,6,6,0,0,0,0,0,0,400,100,200,2,2),(184,3,62,23960,2446,1664,792,792,6,6,0,0,0,0,0,0,400,100,200,2,2),(185,3,63,24440,2494,1696,808,808,6,6,0,0,0,0,0,0,400,100,200,2,2),(186,3,64,24920,2542,1728,824,824,6,6,0,0,0,0,0,0,400,100,200,2,2),(187,3,65,25400,2590,1760,840,840,6,6,0,0,0,0,0,0,400,100,200,2,2),(188,3,66,25880,2638,1792,856,856,6,6,0,0,0,0,0,0,400,100,200,2,2),(189,3,67,26360,2686,1824,872,872,6,6,0,0,0,0,0,0,400,100,200,2,2),(190,3,68,26840,2734,1856,888,888,6,6,0,0,0,0,0,0,400,100,200,2,2),(191,3,69,27320,2782,1888,904,904,6,6,0,0,0,0,0,0,400,100,200,2,2),(192,3,70,27800,2830,1920,920,920,6,6,0,0,0,0,0,0,400,100,200,2,2),(193,3,71,28310,2881,1954,937,937,6,6,0,0,0,0,0,0,400,100,200,2,2),(194,3,72,28820,2932,1988,954,954,6,6,0,0,0,0,0,0,400,100,200,2,2),(195,3,73,29330,2983,2022,971,971,6,6,0,0,0,0,0,0,400,100,200,2,2),(196,3,74,29840,3034,2056,988,988,6,6,0,0,0,0,0,0,400,100,200,2,2),(197,3,75,30350,3085,2090,1005,1005,6,6,0,0,0,0,0,0,400,100,200,2,2),(198,3,76,30860,3136,2124,1022,1022,6,6,0,0,0,0,0,0,400,100,200,2,2),(199,3,77,31370,3187,2158,1039,1039,6,6,0,0,0,0,0,0,400,100,200,2,2),(200,3,78,31880,3238,2192,1056,1056,6,6,0,0,0,0,0,0,400,100,200,2,2),(201,3,79,32390,3289,2226,1073,1073,6,6,0,0,0,0,0,0,400,100,200,2,2),(202,3,80,32900,3340,2260,1090,1090,6,6,0,0,0,0,0,0,400,100,200,2,2),(203,3,81,33440,3394,2296,1108,1108,6,6,0,0,0,0,0,0,400,100,200,2,2),(204,3,82,33980,3448,2332,1126,1126,6,6,0,0,0,0,0,0,400,100,200,2,2),(205,3,83,34520,3502,2368,1144,1144,6,6,0,0,0,0,0,0,400,100,200,2,2),(206,3,84,35060,3556,2404,1162,1162,6,6,0,0,0,0,0,0,400,100,200,2,2),(207,3,85,35600,3610,2440,1180,1180,6,6,0,0,0,0,0,0,400,100,200,2,2),(208,3,86,36140,3664,2476,1198,1198,6,6,0,0,0,0,0,0,400,100,200,2,2),(209,3,87,36680,3718,2512,1216,1216,6,6,0,0,0,0,0,0,400,100,200,2,2),(210,3,88,37220,3772,2548,1234,1234,6,6,0,0,0,0,0,0,400,100,200,2,2),(211,3,89,37760,3826,2584,1252,1252,6,6,0,0,0,0,0,0,400,100,200,2,2),(212,3,90,38300,3880,2620,1270,1270,6,6,0,0,0,0,0,0,400,100,200,2,2),(213,3,91,38870,3937,2658,1289,1289,6,6,0,0,0,0,0,0,400,100,200,2,2),(214,3,92,39440,3994,2696,1308,1308,6,6,0,0,0,0,0,0,400,100,200,2,2),(215,3,93,40010,4051,2734,1327,1327,6,6,0,0,0,0,0,0,400,100,200,2,2),(216,3,94,40580,4108,2772,1346,1346,6,6,0,0,0,0,0,0,400,100,200,2,2),(217,3,95,41150,4165,2810,1365,1365,6,6,0,0,0,0,0,0,400,100,200,2,2),(218,3,96,41720,4222,2848,1384,1384,6,6,0,0,0,0,0,0,400,100,200,2,2),(219,3,97,42290,4279,2886,1403,1403,6,6,0,0,0,0,0,0,400,100,200,2,2),(220,3,98,42860,4336,2924,1422,1422,6,6,0,0,0,0,0,0,400,100,200,2,2),(221,3,99,43430,4393,2962,1441,1441,6,6,0,0,0,0,0,0,400,100,200,2,2),(222,3,100,44000,4450,3000,1460,1460,6,6,0,0,0,0,0,0,400,100,200,2,2),(223,4,4,500,130,50,220,90,6,6,0,6,0,0,0,0,200,100,200,2,2),(224,4,5,600,160,60,250,110,6,6,0,6,0,0,0,0,200,100,200,2,2),(225,4,6,700,190,70,280,130,6,6,0,6,0,0,0,0,200,100,200,2,2),(226,4,7,800,220,80,310,150,6,6,0,6,0,0,0,0,200,100,200,2,2),(227,4,8,900,250,90,340,170,6,6,0,6,0,0,0,0,200,100,200,2,2),(228,4,9,1000,280,100,370,190,6,6,0,6,0,0,0,0,200,100,200,2,2),(229,4,10,1100,310,110,400,210,6,6,0,6,0,0,0,0,200,100,200,2,2),(230,4,11,1210,343,121,433,232,6,6,0,6,0,0,0,0,200,100,200,2,2),(231,4,12,1320,376,132,466,254,6,6,0,6,0,0,0,0,200,100,200,2,2),(232,4,13,1430,409,143,499,276,6,6,0,6,0,0,0,0,200,100,200,2,2),(233,4,14,1540,442,154,532,298,6,6,0,6,0,0,0,0,200,100,200,2,2),(234,4,15,1650,475,165,565,320,6,6,0,6,0,0,0,0,200,100,200,2,2),(235,4,16,1760,508,176,598,342,6,6,0,6,0,0,0,0,200,100,200,2,2),(236,4,17,1870,541,187,631,364,6,6,0,6,0,0,0,0,200,100,200,2,2),(237,4,18,1980,574,198,664,386,6,6,0,6,0,0,0,0,200,100,200,2,2),(238,4,19,2090,607,209,697,408,6,6,0,6,0,0,0,0,200,100,200,2,2),(239,4,20,2200,640,220,730,430,6,6,0,6,0,0,0,0,200,100,200,2,2),(240,4,21,2320,676,232,766,454,6,6,0,6,0,0,0,0,250,100,200,2,2),(241,4,22,2440,712,244,802,478,6,6,0,6,0,0,0,0,250,100,200,2,2),(242,4,23,2560,748,256,838,502,6,6,0,6,0,0,0,0,250,100,200,2,2),(243,4,24,2680,784,268,874,526,6,6,0,6,0,0,0,0,250,100,200,2,2),(244,4,25,2800,820,280,910,550,6,6,0,6,0,0,0,0,250,100,200,2,2),(245,4,26,2920,856,292,946,574,6,6,0,6,0,0,0,0,250,100,200,2,2),(246,4,27,3040,892,304,982,598,6,6,0,6,0,0,0,0,250,100,200,2,2),(247,4,28,3160,928,316,1018,622,6,6,0,6,0,0,0,0,250,100,200,2,2),(248,4,29,3280,964,328,1054,646,6,6,0,6,0,0,0,0,250,100,200,2,2),(249,4,30,3400,1000,340,1090,670,6,6,0,6,0,0,0,0,250,100,200,2,2),(250,4,31,3530,1039,353,1129,696,6,6,0,6,0,0,0,0,300,100,200,2,2),(251,4,32,3660,1078,366,1168,722,6,6,0,6,0,0,0,0,300,100,200,2,2),(252,4,33,3790,1117,379,1207,748,6,6,0,6,0,0,0,0,300,100,200,2,2),(253,4,34,3920,1156,392,1246,774,6,6,0,6,0,0,0,0,300,100,200,2,2),(254,4,35,4050,1195,405,1285,800,6,6,0,6,0,0,0,0,300,100,200,2,2),(255,4,36,4180,1234,418,1324,826,6,6,0,6,0,0,0,0,300,100,200,2,2),(256,4,37,4310,1273,431,1363,852,6,6,0,6,0,0,0,0,300,100,200,2,2),(257,4,38,4440,1312,444,1402,878,6,6,0,6,0,0,0,0,300,100,200,2,2),(258,4,39,4570,1351,457,1441,904,6,6,0,6,0,0,0,0,300,100,200,2,2),(259,4,40,4700,1390,470,1480,930,6,6,0,6,0,0,0,0,300,100,200,2,2),(260,4,41,4840,1432,484,1522,958,6,6,0,6,0,0,0,0,350,100,200,2,2),(261,4,42,4980,1474,498,1564,986,6,6,0,6,0,0,0,0,350,100,200,2,2),(262,4,43,5120,1516,512,1606,1014,6,6,0,6,0,0,0,0,350,100,200,2,2),(263,4,44,5260,1558,526,1648,1042,6,6,0,6,0,0,0,0,350,100,200,2,2),(264,4,45,5400,1600,540,1690,1070,6,6,0,6,0,0,0,0,350,100,200,2,2),(265,4,46,5540,1642,554,1732,1098,6,6,0,6,0,0,0,0,350,100,200,2,2),(266,4,47,5680,1684,568,1774,1126,6,6,0,6,0,0,0,0,350,100,200,2,2),(267,4,48,5820,1726,582,1816,1154,6,6,0,6,0,0,0,0,350,100,200,2,2),(268,4,49,5960,1768,596,1858,1182,6,6,0,6,0,0,0,0,350,100,200,2,2),(269,4,50,6100,1810,610,1900,1210,6,6,0,6,0,0,0,0,350,100,200,2,2),(270,4,51,6250,1855,625,1945,1240,6,6,0,6,0,0,0,0,400,100,200,2,2),(271,4,52,6400,1900,640,1990,1270,6,6,0,6,0,0,0,0,400,100,200,2,2),(272,4,53,6550,1945,655,2035,1300,6,6,0,6,0,0,0,0,400,100,200,2,2),(273,4,54,6700,1990,670,2080,1330,6,6,0,6,0,0,0,0,400,100,200,2,2),(274,4,55,6850,2035,685,2125,1360,6,6,0,6,0,0,0,0,400,100,200,2,2),(275,4,56,7000,2080,700,2170,1390,6,6,0,6,0,0,0,0,400,100,200,2,2),(276,4,57,7150,2125,715,2215,1420,6,6,0,6,0,0,0,0,400,100,200,2,2),(277,4,58,7300,2170,730,2260,1450,6,6,0,6,0,0,0,0,400,100,200,2,2),(278,4,59,7450,2215,745,2305,1480,6,6,0,6,0,0,0,0,400,100,200,2,2),(279,4,60,7600,2260,760,2350,1510,6,6,0,6,0,0,0,0,400,100,200,2,2),(280,4,61,7760,2308,776,2398,1542,6,6,0,6,0,0,0,0,400,100,200,2,2),(281,4,62,7920,2356,792,2446,1574,6,6,0,6,0,0,0,0,400,100,200,2,2),(282,5,1,650,145,120,120,20,6,0,0,6,0,0,0,0,200,100,200,2,2),(283,5,2,800,190,140,140,30,6,0,0,6,0,0,0,0,200,100,200,2,2),(284,4,63,8080,2404,808,2494,1606,6,6,0,6,0,0,0,0,400,100,200,2,2),(285,4,64,8240,2452,824,2542,1638,6,6,0,6,0,0,0,0,400,100,200,2,2),(286,4,65,8400,2500,840,2590,1670,6,6,0,6,0,0,0,0,400,100,200,2,2),(287,4,66,8560,2548,856,2638,1702,6,6,0,6,0,0,0,0,400,100,200,2,2),(288,4,67,8720,2596,872,2686,1734,6,6,0,6,0,0,0,0,400,100,200,2,2),(289,4,68,8880,2644,888,2734,1766,6,6,0,6,0,0,0,0,400,100,200,2,2),(290,4,69,9040,2692,904,2782,1798,6,6,0,6,0,0,0,0,400,100,200,2,2),(291,4,70,9200,2740,920,2830,1830,6,6,0,6,0,0,0,0,400,100,200,2,2),(292,4,71,9370,2791,937,2881,1864,6,6,0,6,0,0,0,0,400,100,200,2,2),(293,5,3,950,235,160,160,40,6,0,0,6,0,0,0,0,200,100,200,2,2),(294,4,72,9540,2842,954,2932,1898,6,6,0,6,0,0,0,0,400,100,200,2,2),(295,4,73,9710,2893,971,2983,1932,6,6,0,6,0,0,0,0,400,100,200,2,2),(296,4,74,9880,2944,988,3034,1966,6,6,0,6,0,0,0,0,400,100,200,2,2),(297,4,75,10050,2995,1005,3085,2000,6,6,0,6,0,0,0,0,400,100,200,2,2),(298,6,1,600,160,110,130,20,3,0,6,0,0,0,0,0,200,100,200,2,2),(299,6,2,700,220,120,160,30,3,0,6,0,0,0,0,0,200,100,200,2,2),(300,6,3,800,280,130,190,40,3,0,6,0,0,0,0,0,200,100,200,2,2),(301,4,76,10220,3046,1022,3136,2034,6,6,0,6,0,0,0,0,400,100,200,2,2),(302,4,77,10390,3097,1039,3187,2068,6,6,0,6,0,0,0,0,400,100,200,2,2),(303,4,78,10560,3148,1056,3238,2102,6,6,0,6,0,0,0,0,400,100,200,2,2),(304,4,79,10730,3199,1073,3289,2136,6,6,0,6,0,0,0,0,400,100,200,2,2),(305,4,80,10900,3250,1090,3340,2170,6,6,0,6,0,0,0,0,400,100,200,2,2),(306,4,81,11080,3304,1108,3394,2206,6,6,0,6,0,0,0,0,400,100,200,2,2),(307,4,82,11260,3358,1126,3448,2242,6,6,0,6,0,0,0,0,400,100,200,2,2),(308,4,83,11440,3412,1144,3502,2278,6,6,0,6,0,0,0,0,400,100,200,2,2),(309,4,84,11620,3466,1162,3556,2314,6,6,0,6,0,0,0,0,400,100,200,2,2),(310,4,85,11800,3520,1180,3610,2350,6,6,0,6,0,0,0,0,400,100,200,2,2),(311,4,86,11980,3574,1198,3664,2386,6,6,0,6,0,0,0,0,400,100,200,2,2),(312,4,87,12160,3628,1216,3718,2422,6,6,0,6,0,0,0,0,400,100,200,2,2),(313,4,88,12340,3682,1234,3772,2458,6,6,0,6,0,0,0,0,400,100,200,2,2),(314,4,89,12520,3736,1252,3826,2494,6,6,0,6,0,0,0,0,400,100,200,2,2),(315,4,90,12700,3790,1270,3880,2530,6,6,0,6,0,0,0,0,400,100,200,2,2),(316,4,91,12890,3847,1289,3937,2568,6,6,0,6,0,0,0,0,400,100,200,2,2),(317,4,92,13080,3904,1308,3994,2606,6,6,0,6,0,0,0,0,400,100,200,2,2),(318,4,93,13270,3961,1327,4051,2644,6,6,0,6,0,0,0,0,400,100,200,2,2),(319,4,94,13460,4018,1346,4108,2682,6,6,0,6,0,0,0,0,400,100,200,2,2),(320,4,95,13650,4075,1365,4165,2720,6,6,0,6,0,0,0,0,400,100,200,2,2),(321,4,96,13840,4132,1384,4222,2758,6,6,0,6,0,0,0,0,400,100,200,2,2),(322,4,97,14030,4189,1403,4279,2796,6,6,0,6,0,0,0,0,200,100,200,2,2),(323,4,98,14220,4246,1422,4336,2834,6,6,0,6,0,0,0,0,400,100,200,2,2),(324,4,99,14410,4303,1441,4393,2872,6,6,0,6,0,0,0,0,400,100,200,2,2),(325,4,100,14600,4360,1460,4450,2910,6,6,0,6,0,0,0,0,400,100,200,2,2),(326,5,4,1100,280,180,180,50,6,0,0,6,0,0,0,0,200,100,200,2,2),(327,5,5,1250,325,200,200,60,6,0,0,6,0,0,0,0,200,100,200,2,2),(328,5,6,1400,370,220,220,70,6,0,0,6,0,0,0,0,200,100,200,2,2),(329,5,7,1550,415,240,240,80,6,0,0,6,0,0,0,0,200,100,200,2,2),(330,5,8,1700,460,260,260,90,6,0,0,6,0,0,0,0,200,100,200,2,2),(331,5,9,1850,505,280,280,100,6,0,0,6,0,0,0,0,200,100,200,2,2),(332,5,10,2000,550,300,300,110,6,0,0,6,0,0,0,0,200,100,200,2,2),(333,5,11,2165,600,322,322,121,6,0,0,6,0,0,0,0,200,100,200,2,2),(334,5,12,2330,650,344,344,132,6,0,0,6,0,0,0,0,200,100,200,2,2),(335,5,13,2495,700,366,366,143,6,0,0,6,0,0,0,0,200,100,200,2,2),(336,5,14,2660,750,388,388,154,6,0,0,6,0,0,0,0,200,100,200,2,2),(337,5,15,2825,800,410,410,165,6,0,0,6,0,0,0,0,200,100,200,2,2),(338,5,16,2990,850,432,432,176,6,0,0,6,0,0,0,0,200,100,200,2,2),(339,5,17,3155,900,454,454,187,6,0,0,6,0,0,0,0,200,100,200,2,2),(340,5,20,3650,1050,520,520,220,6,0,0,6,0,0,0,0,200,100,200,2,2),(341,5,19,3485,1000,498,498,209,6,0,0,6,0,0,0,0,200,100,200,2,2),(342,5,18,3320,950,476,476,198,6,0,0,6,0,0,0,0,200,100,200,2,2),(343,5,21,3830,1104,544,544,232,6,0,0,6,0,0,0,0,250,100,200,2,2),(344,5,22,4010,1158,568,568,244,6,0,0,6,0,0,0,0,250,100,200,2,2),(345,5,23,4190,1212,592,592,256,6,0,0,6,0,0,0,0,250,100,200,2,2),(346,5,24,4370,1266,616,616,268,6,0,0,6,0,0,0,0,250,100,200,2,2),(347,5,25,4550,1320,640,640,280,6,0,0,6,0,0,0,0,250,100,200,2,2),(348,5,26,4730,1374,664,664,292,6,0,0,6,0,0,0,0,250,100,200,2,2),(349,5,27,4910,1428,688,688,304,6,0,0,6,0,0,0,0,250,100,200,2,2),(350,5,28,5090,1482,712,712,316,6,0,0,6,0,0,0,0,250,100,200,2,2),(351,5,29,5270,1536,736,736,328,6,0,0,6,0,0,0,0,250,100,200,2,2),(352,5,30,5450,1590,760,760,340,6,0,0,6,0,0,0,0,250,100,200,2,2),(353,5,31,5645,1649,786,786,353,6,0,0,6,0,0,0,0,300,100,200,2,2),(354,5,32,5840,1708,812,812,366,6,0,0,6,0,0,0,0,300,100,200,2,2),(355,5,33,6035,1767,838,838,379,6,0,0,6,0,0,0,0,300,100,200,2,2),(356,5,34,6230,1826,864,864,392,6,0,0,6,0,0,0,0,300,100,200,2,2),(360,6,4,900,340,140,220,50,3,0,6,0,0,0,0,0,200,100,200,2,2),(361,6,5,1000,400,150,250,60,3,0,6,0,0,0,0,0,200,100,200,2,2),(362,6,6,1100,460,160,280,70,3,0,6,0,0,0,0,0,200,100,200,2,2),(363,6,7,1200,520,170,310,80,3,0,6,0,0,0,0,0,200,100,200,2,2),(364,6,8,1300,580,180,340,90,3,0,6,0,0,0,0,0,200,100,200,2,2),(365,6,9,1400,640,190,370,100,3,0,6,0,0,0,0,0,200,100,200,2,2),(366,6,10,1500,700,200,400,110,3,0,6,0,0,0,0,0,200,100,200,2,2),(367,6,11,1610,766,211,433,121,3,0,6,0,0,0,0,0,200,100,200,2,2),(369,6,12,1720,832,222,466,132,3,0,6,0,0,0,0,0,200,100,200,2,2),(370,6,13,1830,898,233,499,143,3,0,6,0,0,0,0,0,200,100,200,2,2),(371,6,14,1940,964,244,532,154,3,0,6,0,0,0,0,0,200,100,200,2,2),(372,6,15,2050,1030,255,565,165,3,0,6,0,0,0,0,0,200,100,200,2,2),(373,6,16,2160,1096,266,598,176,3,0,6,0,0,0,0,0,200,100,200,2,2),(374,6,17,2270,1162,277,631,187,3,0,6,0,0,0,0,0,200,100,200,2,2),(375,6,18,2380,1228,288,664,198,3,0,6,0,0,0,0,0,200,100,200,2,2),(376,6,19,2490,1294,299,697,209,3,0,6,0,0,0,0,0,200,100,200,2,2),(377,6,20,2600,1360,310,730,220,3,0,6,0,0,0,0,0,200,100,200,2,2),(378,6,21,2720,1432,322,766,232,3,0,6,0,0,0,0,0,250,100,200,2,2),(379,6,22,2840,1504,334,802,244,3,0,6,0,0,0,0,0,250,100,200,2,2),(380,6,23,2960,1576,346,838,256,3,0,6,0,0,0,0,0,250,100,200,2,2),(381,6,24,3080,1648,358,874,268,3,0,6,0,0,0,0,0,250,100,200,2,2),(382,6,25,3200,1720,370,910,280,3,0,6,0,0,0,0,0,250,100,200,2,2),(383,6,26,3320,1792,382,946,292,3,0,6,0,0,0,0,0,250,100,200,2,2),(384,6,27,3440,1864,394,982,304,3,0,6,0,0,0,0,0,250,100,200,2,2),(385,6,28,3560,1936,406,1018,316,3,0,6,0,0,0,0,0,250,100,200,2,2),(386,6,29,3680,2008,418,1054,328,3,0,6,0,0,0,0,0,250,100,200,2,2),(387,6,30,3800,2080,430,1090,340,3,0,6,0,0,0,0,0,250,100,200,2,2),(388,6,31,3930,2158,443,1129,353,3,0,6,0,0,0,0,0,300,100,200,2,2),(389,6,32,4060,2236,456,1168,366,3,0,6,0,0,0,0,0,300,100,200,2,2),(390,6,33,4190,2314,469,1207,379,3,0,6,0,0,0,0,0,300,100,200,2,2),(391,6,34,4320,2392,482,1246,392,3,0,6,0,0,0,0,0,300,100,200,2,2),(392,6,35,4450,2470,495,1285,405,3,0,6,0,0,0,0,0,300,100,200,2,2),(393,6,36,4580,2548,508,1324,418,3,0,6,0,0,0,0,0,300,100,200,2,2),(394,6,37,4710,2626,521,1363,431,3,0,6,0,0,0,0,0,300,100,200,2,2),(395,6,38,4840,2704,534,1402,444,3,0,6,0,0,0,0,0,300,100,200,2,2),(396,6,39,4970,2782,547,1441,457,3,0,6,0,0,0,0,0,300,100,200,2,2),(397,6,40,5100,2860,560,1480,470,3,0,6,0,0,0,0,0,300,100,200,2,2),(398,6,41,5240,2944,574,1522,484,3,0,6,0,0,0,0,0,350,100,200,2,2),(399,6,42,5380,3028,588,1564,498,3,0,6,0,0,0,0,0,350,100,200,2,2),(400,6,43,5520,3112,602,1606,512,3,0,6,0,0,0,0,0,350,100,200,2,2),(401,6,44,5660,3190,616,1648,526,3,0,6,0,0,0,0,0,350,100,200,2,2),(402,6,45,5800,3280,630,1690,540,3,0,6,0,0,0,0,0,350,100,200,2,2),(403,6,46,5940,3364,644,1732,554,3,0,6,0,0,0,0,0,350,100,200,2,2),(404,6,47,6080,3448,658,1774,568,3,0,6,0,0,0,0,0,350,100,200,2,2),(405,6,48,6226,3532,672,1816,582,3,0,6,0,0,0,0,0,350,100,200,2,2),(406,6,49,6360,3616,686,1858,596,3,0,6,0,0,0,0,0,350,100,200,2,2),(407,6,50,6500,3700,700,1900,610,3,0,6,0,0,0,0,0,350,100,200,2,2),(409,6,51,6650,3790,715,1945,625,3,0,6,0,0,0,0,0,400,100,200,2,2),(410,6,52,6800,3880,730,1990,640,3,0,6,0,0,0,0,0,400,100,200,2,2),(411,6,53,6950,3970,745,2035,655,3,0,6,0,0,0,0,0,400,100,200,2,2),(412,6,54,7100,4060,760,2080,670,3,0,6,0,0,0,0,0,400,100,200,2,2),(413,6,55,7250,4150,775,2125,685,3,0,6,0,0,0,0,0,400,100,200,2,2),(414,6,56,7400,4240,790,2170,700,3,0,6,0,0,0,0,0,400,100,200,2,2),(415,6,57,7550,4330,805,2215,715,3,0,6,0,0,0,0,0,400,100,200,2,2),(416,6,58,7700,4420,820,2260,730,3,0,6,0,0,0,0,0,400,100,200,2,2),(417,6,59,7850,4510,835,2305,745,3,0,6,0,0,0,0,0,400,100,200,2,2),(418,6,60,8000,4600,850,2350,760,3,0,6,0,0,0,0,0,400,100,200,2,2),(419,6,61,8160,4696,866,2398,776,3,0,6,0,0,0,0,0,400,100,200,2,2),(420,6,62,8320,4792,882,2446,792,3,0,6,0,0,0,0,0,400,100,200,2,2),(421,6,63,8480,4888,898,2494,808,3,0,6,0,0,0,0,0,400,100,200,2,2),(422,6,64,8640,4984,914,2542,824,3,0,6,0,0,0,0,0,400,100,200,2,2),(423,6,65,8800,5080,930,2590,840,3,0,6,0,0,0,0,0,400,100,200,2,2),(424,6,66,8960,5176,946,2638,856,3,0,6,0,0,0,0,0,400,100,200,2,2),(425,6,67,9120,5272,962,2686,872,3,0,6,0,0,0,0,0,400,100,200,2,2),(426,6,68,9280,5368,978,2734,888,3,0,6,0,0,0,0,0,400,100,200,2,2),(427,6,69,9440,5464,994,2782,904,3,0,6,0,0,0,0,0,400,100,200,2,2),(428,6,70,9600,5560,1010,2830,920,3,0,6,0,0,0,0,0,400,100,200,2,2),(430,6,71,9770,5662,1027,2881,937,3,0,6,0,0,0,0,0,400,100,200,2,2),(431,6,72,9940,5764,1044,2932,954,3,0,6,0,0,0,0,0,400,100,200,2,2),(432,6,73,10110,5866,1061,2983,971,3,0,6,0,0,0,0,0,400,100,200,2,2),(433,6,74,10280,5968,1078,3034,988,3,0,6,0,0,0,0,0,400,100,200,2,2),(434,6,75,10450,6070,1095,3085,1005,3,0,6,0,0,0,0,0,400,100,200,2,2),(435,6,76,10620,6172,1112,3136,1022,3,0,6,0,0,0,0,0,400,100,200,2,2),(436,6,77,10790,6274,1129,3187,1039,3,0,6,0,0,0,0,0,400,100,200,2,2),(437,6,78,10960,6376,1146,3238,1056,3,0,6,0,0,0,0,0,400,100,200,2,2),(438,6,79,11130,6478,1163,3289,1073,3,0,6,0,0,0,0,0,400,100,200,2,2),(439,6,80,11300,6580,1180,3340,1090,3,0,6,0,0,0,0,0,400,100,200,2,2),(441,6,81,11480,6688,1198,3394,1108,3,0,6,0,0,0,0,0,400,100,200,2,2),(442,6,82,11660,6796,1216,3448,1126,3,0,6,0,0,0,0,0,400,100,200,2,2),(443,6,83,11840,6904,1234,3502,1144,3,0,6,0,0,0,0,0,400,100,200,2,2),(444,6,84,12020,7012,1252,3556,1162,3,0,6,0,0,0,0,0,400,100,200,2,2),(445,6,85,12200,7120,1270,3610,1180,3,0,6,0,0,0,0,0,400,100,200,2,2),(446,6,86,12380,7228,1288,3664,1198,3,0,6,0,0,0,0,0,400,100,200,2,2),(447,6,87,12560,7336,1306,3718,1216,3,0,6,0,0,0,0,0,400,100,200,2,2),(448,6,88,12740,7444,1324,3772,1234,3,0,6,0,0,0,0,0,400,100,200,2,2),(449,6,89,12920,7552,1342,3826,1252,3,0,6,0,0,0,0,0,400,100,200,2,2),(450,6,90,13100,7660,1360,3880,1270,3,0,6,0,0,0,0,0,400,100,200,2,2),(451,6,91,13290,7774,1379,3937,1289,3,0,6,0,0,0,0,0,400,100,200,2,2),(452,6,92,13480,7888,1398,3994,1308,3,0,6,0,0,0,0,0,400,100,200,2,2),(453,6,93,13670,8002,1417,4051,1327,3,0,6,0,0,0,0,0,400,100,200,2,2),(454,6,94,13860,8116,1436,4108,1346,3,0,6,0,0,0,0,0,400,100,200,2,2),(455,6,95,14050,8230,1455,4165,1365,3,0,6,0,0,0,0,0,400,100,200,2,2),(456,6,96,14240,8344,1474,4222,1384,3,0,6,0,0,0,0,0,400,100,200,2,2),(457,6,97,14430,8458,1493,4279,1403,3,0,6,0,0,0,0,0,400,100,200,2,2),(458,6,98,14620,8572,1512,4336,1422,3,0,6,0,0,0,0,0,400,100,200,2,2),(459,6,99,14810,8686,1531,4393,1441,3,0,6,0,0,0,0,0,400,100,200,2,2),(460,6,100,15000,8800,1550,4450,1460,3,0,6,0,0,0,0,0,400,100,200,2,2),(461,5,35,6425,1885,890,890,405,6,0,0,6,0,0,0,0,300,100,200,2,2),(462,5,36,6620,1944,916,916,418,6,0,0,6,0,0,0,0,300,100,200,2,2),(463,5,37,6815,2003,942,942,431,6,0,0,6,0,0,0,0,300,100,200,2,2),(464,5,38,7010,2062,968,968,444,6,0,0,6,0,0,0,0,300,100,200,2,2),(465,5,39,7205,2121,994,994,457,6,0,0,6,0,0,0,0,300,100,200,2,2),(466,5,40,7400,2180,1020,1020,470,6,0,0,6,0,0,0,0,300,100,200,2,2),(467,5,41,7610,2243,1048,1048,484,6,0,0,6,0,0,0,0,350,100,200,2,2),(468,5,42,7820,2306,1076,1076,498,6,0,0,6,0,0,0,0,350,100,200,2,2),(469,5,43,8030,2369,1104,1104,512,6,0,0,6,0,0,0,0,350,100,200,2,2),(470,5,44,8240,2432,1132,1132,526,6,0,0,6,0,0,0,0,350,100,200,2,2),(471,5,45,8450,2495,1160,1160,540,6,0,0,6,0,0,0,0,350,100,200,2,2),(472,5,46,8660,2558,1188,1188,554,6,0,0,6,0,0,0,0,350,100,200,2,2),(473,5,47,8870,2621,1216,1216,568,6,0,0,6,0,0,0,0,350,100,200,2,2),(474,5,48,9080,2684,1244,1244,582,6,0,0,6,0,0,0,0,350,100,200,2,2),(475,5,49,9290,2749,1272,1272,596,6,0,0,6,0,0,0,0,350,100,200,2,2),(476,5,50,9500,2810,1300,1300,610,6,0,0,6,0,0,0,0,350,100,200,2,2),(477,5,51,9725,2878,1330,1330,625,6,0,0,6,0,0,0,0,400,100,200,2,2),(478,5,52,9950,2946,1360,1360,640,6,0,0,6,0,0,0,0,400,100,200,2,2),(479,5,53,10175,3014,1390,1390,655,6,0,0,6,0,0,0,0,400,100,200,2,2),(480,5,54,10400,3082,1420,1420,670,6,0,0,6,0,0,0,0,400,100,200,2,2),(481,5,55,10625,3150,1450,1450,685,6,0,0,6,0,0,0,0,400,100,200,2,2),(482,5,56,10850,3218,1480,1480,700,6,0,0,6,0,0,0,0,400,100,200,2,2),(483,5,57,11075,3286,1510,1510,715,6,0,0,6,0,0,0,0,400,100,200,2,2),(484,5,58,11300,3354,1540,1540,730,6,0,0,6,0,0,0,0,400,100,200,2,2),(485,5,59,11525,3422,1570,1570,745,6,0,0,6,0,0,0,0,400,100,200,2,2),(486,5,60,11750,3490,1600,1600,760,6,0,0,6,0,0,0,0,400,100,200,2,2),(487,5,61,11990,3562,1632,1632,776,6,0,0,6,0,0,0,0,400,100,200,2,2),(488,5,62,12230,3634,1664,1664,792,6,0,0,6,0,0,0,0,400,100,200,2,2),(489,5,63,12470,3706,1696,1696,808,6,0,0,6,0,0,0,0,400,100,200,2,2),(490,5,64,12710,3778,1728,1728,824,6,0,0,6,0,0,0,0,400,100,200,2,2),(491,5,65,12950,3850,1760,1760,840,6,0,0,6,0,0,0,0,400,100,200,2,2),(492,5,66,13190,3922,1792,1792,856,6,0,0,6,0,0,0,0,400,100,200,2,2),(493,5,67,13430,3994,1824,1824,872,6,0,0,6,0,0,0,0,400,100,200,2,2),(494,5,68,13670,4066,1856,1856,888,6,0,0,6,0,0,0,0,400,100,200,2,2),(495,5,69,13910,4138,1888,1888,904,6,0,0,6,0,0,0,0,400,100,200,2,2),(496,5,70,14150,4210,1920,1920,920,6,0,0,6,0,0,0,0,400,100,200,2,2),(497,5,71,14405,4287,1954,1954,937,6,0,0,6,0,0,0,0,400,100,200,2,2),(498,5,72,14660,4364,1988,1988,954,6,0,0,6,0,0,0,0,400,100,200,2,2),(499,5,73,14915,4441,2022,2022,971,6,0,0,6,0,0,0,0,400,100,200,2,2),(500,5,74,15170,4518,2056,2056,988,6,0,0,6,0,0,0,0,400,100,200,2,2),(501,5,75,15425,4595,2090,2090,1005,6,0,0,6,0,0,0,0,400,100,200,2,2),(502,5,76,15680,4672,2124,2124,1022,6,0,0,6,0,0,0,0,400,100,200,2,2),(503,5,77,15935,4749,2158,2158,1039,6,0,0,6,0,0,0,0,400,100,200,2,2),(504,5,78,16190,4826,2192,2192,1056,6,0,0,6,0,0,0,0,400,100,200,2,2),(505,5,79,16445,4903,2226,2226,1073,6,0,0,6,0,0,0,0,400,100,200,2,2),(506,5,80,16700,4980,2260,2260,1090,6,0,0,6,0,0,0,0,400,100,200,2,2),(507,5,81,16970,5061,2296,2296,1108,6,0,0,6,0,0,0,0,400,100,200,2,2),(508,5,82,17240,5142,2332,2332,1126,6,0,0,6,0,0,0,0,400,100,200,2,2),(509,5,83,17510,5223,2368,2368,1144,6,0,0,6,0,0,0,0,400,100,200,2,2),(510,5,84,17780,5304,2404,2404,1162,6,0,0,6,0,0,0,0,400,100,200,2,2),(511,5,85,18050,5385,2440,2440,1180,6,0,0,6,0,0,0,0,400,100,200,2,2),(512,5,86,18320,5466,2476,2476,1198,6,0,0,6,0,0,0,0,400,100,200,2,2),(513,5,87,18590,5547,2512,2512,1216,6,0,0,6,0,0,0,0,400,100,200,2,2),(514,5,88,18860,5628,2548,2548,1234,6,0,0,6,0,0,0,0,400,100,200,2,2),(515,5,89,19130,5709,2584,2584,1252,6,0,0,6,0,0,0,0,400,100,200,2,2),(516,5,90,19400,5790,2620,2620,1270,6,0,0,6,0,0,0,0,400,100,200,2,2),(517,5,91,19685,5876,2658,2658,1289,6,0,0,6,0,0,0,0,400,100,200,2,2),(518,5,92,19970,5962,2696,2696,1308,6,0,0,6,0,0,0,0,400,100,200,2,2),(519,5,93,20255,6048,2734,2734,1327,6,0,0,6,0,0,0,0,400,100,200,2,2),(520,5,94,20540,6134,2772,2772,1346,6,0,0,6,0,0,0,0,400,100,200,2,2),(521,5,95,20825,6220,2810,2810,1365,6,0,0,6,0,0,0,0,400,100,200,2,2),(522,5,96,21110,6306,2848,2848,1384,6,0,0,6,0,0,0,0,400,100,200,2,2),(523,5,97,21395,6392,2886,2886,1403,6,0,0,6,0,0,0,0,400,100,200,2,2),(524,5,98,21680,6478,2924,2924,1422,6,0,0,6,0,0,0,0,400,100,200,2,2),(525,5,99,21965,6564,2962,2962,1441,6,0,0,6,0,0,0,0,400,100,200,2,2),(526,5,100,22250,6650,3000,3000,1460,6,0,0,6,0,0,0,0,400,100,200,2,2);
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
INSERT INTO `role_realm_class` VALUES (1,'炼气',1,20,265,7,0,0,0,0,0,0,0,0,0,0,0,0,'初始阶段，进步容易'),(2,'金刚',2,40,266,14,5000,50,4,4,0,0,0,0,0,0,0,10,'金刚境界，武功境界可提升至40级，身体强壮，可抵御一定外力伤害'),(3,'道玄',3,60,267,28,10000,100,6,4,30,0,0,0,0,0,0,10,'道玄境界，武功境界可提升至60级，劲气外放，操控自如'),(4,'万象',4,80,268,56,15000,150,8,4,30,100,100,100,100,100,100,10,'万象境界，武功境界可提升至80级，神念可感应天地万物'),(5,'天人',5,100,268,112,25000,200,10,8,30,100,100,100,100,100,100,20,'天人境界，武功境界可提升至100级，精气神达到完美的极致境界');
/*!40000 ALTER TABLE `role_realm_class` ENABLE KEYS */;
DROP TABLE IF EXISTS `role_realm_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role_realm_level` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `realm_level` smallint(6) NOT NULL DEFAULT '0' COMMENT '境界等级',
  `exp` bigint(20) NOT NULL DEFAULT '0' COMMENT '升级所需经验',
  `need_realm_class` smallint(6) NOT NULL DEFAULT '0' COMMENT '升级所需阶级',
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
INSERT INTO `role_realm_level` VALUES (1,1,1500,1,1,100,20,20,20,20),(2,2,1500,1,1,200,40,40,40,40),(3,3,1500,1,1,300,60,60,60,60),(4,4,1500,1,1,400,80,80,80,80),(5,5,3000,1,1,500,100,100,100,100),(6,6,3000,1,1,600,120,120,120,120),(7,7,3000,1,1,700,140,140,140,140),(8,8,3000,1,1,800,160,160,160,160),(9,9,3000,1,1,900,180,180,180,180),(10,10,4500,1,1,1000,200,200,200,200),(11,11,4500,1,1,1100,220,220,220,220),(12,12,4500,1,1,1200,240,240,240,240),(13,13,4500,1,1,1300,260,260,260,260),(14,14,4500,1,1,1400,280,280,280,280),(15,15,4500,1,1,1500,300,300,300,300),(16,16,4500,1,1,1600,320,320,320,320),(17,17,4500,1,1,1700,340,340,340,340),(18,18,4500,1,1,1800,360,360,360,360),(19,19,4500,1,1,1900,380,380,380,380),(20,20,9000,2,2,2000,400,400,400,400),(21,21,9000,2,2,2100,420,420,420,420),(22,22,9000,2,2,2200,440,440,440,440),(23,23,9000,2,2,2300,460,460,460,460),(24,24,9000,2,2,2400,480,480,480,480),(25,25,9000,2,2,2500,500,500,500,500),(26,26,9000,2,2,2600,520,520,520,520),(27,27,9000,2,2,2700,540,540,540,540),(28,28,9000,2,2,2800,560,560,560,560),(29,29,9000,2,2,2900,580,580,580,580),(30,30,13500,2,3,3000,600,600,600,600),(31,31,13500,2,3,3100,620,620,620,620),(32,32,13500,2,3,3200,640,640,640,640),(33,33,13500,2,3,3300,660,660,660,660),(34,34,13500,2,3,3400,680,680,680,680),(35,35,13500,2,3,3500,700,700,700,700),(36,36,13500,2,3,3600,720,720,720,720),(37,37,13500,2,3,3700,740,740,740,740),(38,38,13500,2,3,3800,760,760,760,760),(39,39,13500,2,3,3900,780,780,780,780),(40,40,18000,3,4,4000,800,800,800,800),(41,41,18000,3,4,4100,820,820,820,820),(42,42,18000,3,4,4200,840,840,840,840),(43,43,18000,3,4,4300,860,860,860,860),(44,44,18000,3,4,4400,880,880,880,880),(45,45,18000,3,4,4500,900,900,900,900),(46,46,18000,3,4,4600,920,920,920,920),(47,47,18000,3,4,4700,940,940,940,940),(48,48,18000,3,4,4800,960,960,960,960),(49,49,18000,3,4,4900,980,980,980,980),(50,50,22500,3,5,5000,1000,1000,1000,1000),(51,51,22500,3,5,5100,1020,1020,1020,1020),(52,52,22500,3,5,5200,1040,1040,1040,1040),(53,53,22500,3,5,5300,1060,1060,1060,1060),(54,54,22500,3,5,5400,1080,1080,1080,1080),(55,55,22500,3,5,5500,1100,1100,1100,1100),(56,56,22500,3,5,5600,1120,1120,1120,1120),(57,57,22500,3,5,5700,1140,1140,1140,1140),(58,58,22500,3,5,5800,1160,1160,1160,1160),(59,59,2250,3,5,5900,1180,1180,1180,1180),(60,60,27000,4,6,6000,1200,1200,1200,1200),(61,61,27000,4,6,6100,1220,1220,1220,1220),(62,62,27000,4,6,6200,1240,1240,1240,1240),(63,63,27000,4,6,6300,1260,1260,1260,1260),(64,64,27000,4,6,6400,1280,1280,1280,1280),(65,65,27000,4,6,6500,1300,1300,1300,1300),(66,66,27000,4,6,6600,1320,1320,1320,1320),(67,67,27000,4,6,6700,1340,1340,1340,1340),(68,68,27000,4,6,6800,1360,1360,1360,1360),(69,69,27000,4,6,6900,1380,1380,1380,1380),(70,70,31500,4,7,7000,1400,1400,1400,1400),(71,71,31500,4,7,7100,1420,1420,1420,1420),(72,72,31500,4,7,7200,1440,1440,1440,1440),(73,73,31500,4,7,7300,1460,1460,1460,1460),(74,74,31500,4,7,7400,1480,1480,1480,1480),(75,75,31500,4,7,7500,1500,1500,1500,1500),(76,76,31500,4,7,7600,1520,1520,1520,1520),(77,77,31500,4,7,7700,1540,1540,1540,1540),(78,78,31500,4,7,7800,1560,1560,1560,1560),(79,79,31500,4,7,7900,1580,1580,1580,1580),(80,80,36000,4,8,8000,1600,1600,1600,1600);
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
  `target` tinyint(4) NOT NULL DEFAULT '1' COMMENT '攻击目标（客户端展示用）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1042 DEFAULT CHARSET=utf8mb4 COMMENT='绝招表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `skill` DISABLE KEYS */;
INSERT INTO `skill` VALUES (1,'闪击',1,1,'ShanJi','ShanJi',-2,1,'单体攻击，获得2点精气',1,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 10,\"DecPower\": 0,\"IncPower\": 2,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,1,1,1,0,1),(2,'裂地斩',1,1,'LieDiZhan','LieDiZhan',-2,1,'横排攻击，对多个敌人产生伤害',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 100,\"Cul2AtkRate\": 60,\"DecPower\": 2,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,2,1,0,3),(3,'甲溃',1,1,'JiaKui','JiaKui',-2,5,'单体攻击，破坏敌方护甲，获得1点精气',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 10,\"DecPower\": 0,\"IncPower\": 1,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 50,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,3,1,0,1),(4,'星爆光离',1,1,'XingBaoGuangLi','XingBaoGuangLi',-2,2,'单体攻击，消耗精气产生大量伤害',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 60,\"DecPower\": 6,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 50,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,1,4,1,0,1),(5,'刀芒',1,1,'DaoMang','DaoMang',3,10,'纵向攻击，产生大量伤害',0,0,'{\"TargetMode\": 3,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 50,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,1,5,1,0,4),(6,'刀盾',1,3,'DaoDun','DaoDun',3,1,'防御1回合，自身减免40%的伤害，并吸引攻击',0,0,'{\"TargetMode\": 4,\"AttackMode\": 1,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [{\"Type\": 18, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 40, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0},{\"Type\": 19, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,1,6,1,0,1),(7,'风咒',1,1,'FengZhou','FengZhou',4,1,'单体攻击',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 50,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,1,7,1,0,1),(8,'治愈之风',1,4,'ZhiYuZhiFeng','ZhiLiao',4,1,'单体治疗，优先治疗受伤最重的伙伴',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 500, \"RawValueRate\": 0, \"AttackRate\": 33, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 100, \"ValueCountRate\": 0, \"TargetMode\": 2}]}',0,1,8,1,0,1),(9,'咆哮利爪',5,1,'PaoXiaoLiZhuaM','PaoXiaoLiZhua',-1,0,'全体',0,0,'{\"TargetMode\": 1,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,30,2),(10,'凶猛撕咬',5,1,'XiongMengSiYaoM','XiongMengSiYao',-1,0,'单体',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0,1),(11,'冰烈',5,1,'BingLieM','BingLie',-1,0,'单体',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0,1),(12,'冰烈横向',5,1,'BingLieHengXiangM','BingLieQuanTi',-1,0,'横向',0,0,'{\"TargetMode\": 2,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 5,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,10,3),(13,'冰烈纵向',5,1,'BingLieZongXiangM','BingLieQuanTi',-1,0,'纵向',0,0,'{\"TargetMode\": 3,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,20,4),(14,'冰烈全体',5,1,'BingLieQuanTiM','BingLieQuanTi',-1,0,'全体',0,0,'{\"TargetMode\": 1,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,30,2),(15,'火烈',5,1,'HuoLieM','HuoLie',-1,0,'单体',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0,1),(16,'火烈横向',5,1,'HuoLieHengXiangM','HuoLieQuanTi',-1,0,'横向',0,0,'{\"TargetMode\": 2,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,10,3),(17,'火烈纵向',5,1,'HuoLieZongXiangM','HuoLieQuanTi',-1,0,'纵向',0,0,'{\"TargetMode\": 3,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,20,4),(18,'火烈全体',5,1,'HuoLieQuanTiM','HuoLieQuanTi',-1,0,'全体',0,0,'{\"TargetMode\": 1,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,30,2),(19,'风烈',5,1,'FengLieM','FengLie',-1,0,'单体',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0,1),(20,'风烈横向',5,1,'FengLieHengXiangM','FengLieQuanTi',-1,0,'横向',0,0,'{\"TargetMode\": 2,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,10,3),(21,'风烈纵向',5,1,'FengLieZongXiangM','FengLieQuanTi',-1,0,'纵向',0,0,'{\"TargetMode\": 3,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,20,4),(22,'风烈全体',5,1,'FengLieQuanTiM','FengLieQuanTi',-1,0,'全体',0,0,'{\"TargetMode\": 1,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,30,2),(23,'雷烈',5,1,'LeiLieM','LeiLie',-1,0,'单体',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0,1),(24,'雷烈横向',5,1,'LeiLieHengXiangM','LeiLieQuanTi',-1,0,'横向',0,0,'{\"TargetMode\": 2,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,10,3),(25,'雷烈纵向',5,1,'LeiLieZongXiangM','LeiLieQuanTi',-1,0,'纵向',0,0,'{\"TargetMode\": 3,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,20,4),(26,'雷烈全体',5,1,'LeiLieQuanTiM','LeiLieQuanTi',-1,0,'全体',0,0,'{\"TargetMode\": 1,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,30,2),(27,'土烈',5,1,'TuLieM','TuLie',-1,0,'单体',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0,1),(28,'土烈横向',5,1,'TuLieHengXiangM','TuLieQuanTi',-1,0,'横向',0,0,'{\"TargetMode\": 2,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,10,3),(29,'土烈纵向',5,1,'TuLieZongXiangM','TuLieQuanTi',-1,0,'纵向',0,0,'{\"TargetMode\": 3,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,20,4),(30,'土烈全体',5,1,'TuLieQuanTiM','TuLieQuanTi',-1,0,'全体',0,0,'{\"TargetMode\": 1,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,30,2),(31,'毒烈',5,1,'DuLieM','DuLie',-1,0,'单体',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 6, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 0, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 30, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',0,0,0,0,0,1),(32,'毒烈横向',5,1,'DuLieHengXiangM','DuLieQuanTi',-1,0,'横向',0,0,'{\"TargetMode\": 2,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 6, \"Keep\": 3, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 0, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 30, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',0,0,0,0,10,3),(33,'毒烈纵向',5,1,'DuLieZongXiangM','DuLieQuanTi',-1,0,'纵向',0,0,'{\"TargetMode\": 3,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 6, \"Keep\": 3, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 0, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 30, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',0,0,0,0,20,4),(34,'毒烈全体',5,1,'DuLieQuanTiM','DuLieQuanTi',-1,0,'全体',0,0,'{\"TargetMode\": 1,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 6, \"Keep\": 3, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 0, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 30, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',0,0,0,0,30,2),(35,'多连斩',5,1,'DuoLianZhanM','DuoLianZhan',-1,0,'单体攻击',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0,1),(36,'力劈华山',5,1,'LiPiHuaShanM','LiPiHuaShan',-1,0,'单体攻击，对敌人造成猛烈的打击',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0,1),(37,'白莲横江',5,1,'BaiLianHengJiangM','LieDiZhan',-1,0,'横向',0,0,'{\"TargetMode\": 2,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,10,3),(38,'横扫千军',5,1,'HengSaoQianJunM','HengSaoQianJun',-1,0,'横向',0,0,'{\"TargetMode\": 2,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,10,3),(39,'乾坤刀气',5,1,'QianKunDaoQiM','QianKunDaoQi',-1,0,'纵向',0,0,'{\"TargetMode\": 3,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,20,4),(40,'三千洛水剑',5,1,'SanQianLuoShuiJianM','SanQianLuoShuiJian',-1,0,'全体',0,0,'{\"TargetMode\": 1,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,30,2),(41,'死亡标记',5,1,'SiWangBiaoJiM','SiWangBiaoJi',-1,0,'单体',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0,1),(42,'万箭穿心',5,1,'WanJianChuanXinM','WanJianChuanXin',-1,0,'单体',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0,1),(43,'狮吼功',5,1,'ShiHouGongM','ShiHouGong',-1,0,'单体',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0,1),(44,'野蛮冲撞',5,1,'YeManChongZhuangM','YeManChongZhuang',-1,0,'纵向',0,0,'{\"TargetMode\": 3,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,20,4),(45,'如殂随行',5,1,'RuCuSuiXingM','RuCuSuiXing',-1,0,'全体',0,0,'{\"TargetMode\": 1,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,30,2),(46,'驱散',5,5,'QuSanM','QuSan',-1,0,'自身',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [{\"Type\": 7, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 0, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,5,1),(49,'青竹咒',1,1,'QinZhuZhou','QinZhuZhou',5,1,'单体攻击',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 10,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,1,49,1,0,1),(50,'雨润',1,5,'YuRun','YuRun',5,1,'全体辅助，增加攻击500，持续2回合',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 2, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,1,50,1,0,1),(51,'墨画影狼',1,1,'MoHuaYingLang','MoHuaYingLang',6,1,'单体攻击',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 500,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,1,51,1,0,1),(52,'墨画巫雀',1,5,'MoHuaWuQue','MoHuaWuQue',6,1,'单体攻击，同时每次降低敌方600防御，持续1回合',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 500,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 1, \"BaseValue\": 600, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',0,1,52,1,0,1),(53,'金刚刀芒',1,1,'DaoMang','DaoMang',3,20,'纵向攻击，产生大量伤害',0,0,'{\"TargetMode\": 3,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,1,5,2,0,4),(54,'道玄刀芒',1,1,'DaoMang','DaoMang',3,40,'纵向攻击，产生大量伤害',0,0,'{\"TargetMode\": 3,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,5,3,0,4),(55,'万象刀芒',1,1,'DaoMang','DaoMang',3,60,'纵向攻击，产生大量伤害',0,0,'{\"TargetMode\": 3,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 3000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,5,4,0,4),(56,'天人刀芒',1,1,'DaoMang','DaoMang',3,80,'纵向攻击，产生大量伤害',0,0,'{\"TargetMode\": 3,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 4000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,5,5,0,4),(57,'金刚刀盾',1,3,'DaoDun','DaoDun',3,20,'防御1回合，自身减免50%的伤害，并吸引攻击。队伍防御增加500，持续1回合',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [{\"Type\": 18, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 50, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0},{\"Type\": 19, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,1,6,2,0,1),(58,'道玄刀盾',1,3,'DaoDun','DaoDun',3,40,'防御1回合，自身减免60%的伤害，并吸引攻击。队伍防御增加1000，持续1回合',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [{\"Type\": 18, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 60, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0},{\"Type\": 19, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',1,1,6,3,0,1),(59,'万象刀盾',1,3,'DaoDun','DaoDun',3,60,'防御1回合，自身减免70%的伤害，并吸引攻击。队伍防御增加2000，持续1回合',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [{\"Type\": 18, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 70, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0},{\"Type\": 19, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 2000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',1,1,6,4,0,1),(60,'天人刀盾',1,3,'DaoDun','DaoDun',3,80,'防御1回合，自身减免80%的伤害，并吸引攻击。队伍防御增加3000，持续1回合',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [{\"Type\": 18, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 80, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0},{\"Type\": 19, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 3000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',1,1,6,5,0,1),(61,'有限无敌',1,3,'YouXianWuDi','TongQiangTieBi',3,30,'防御1回合，自身减免99%的伤害，并吸引攻击，持续1回合',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [{\"Type\": 18, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 99, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0},{\"Type\": 19, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,0,61,1,0,1),(62,'金刚大风咒',1,1,'DaFengZhou','DaFengZhou',4,20,'横排攻击',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,1036,2,0,3),(63,'道玄大风咒',1,1,'DaFengZhou','DaFengZhou',4,40,'横排攻击',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,1036,3,0,3),(64,'万象大风咒',1,1,'DaFengZhou','DaFengZhou',4,60,'横排攻击',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 3000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,1036,4,0,3),(65,'天人大风咒',1,1,'DaFengZhou','DaFengZhou',4,80,'横排攻击',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 4000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,1036,5,0,3),(66,'金刚治愈之风',1,4,'ZhiYuZhiFeng','ZhiLiao',4,20,'单体治疗，优先治疗受伤最重的伙伴',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 2000, \"RawValueRate\": 0, \"AttackRate\": 30, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 100, \"ValueCountRate\": 0, \"TargetMode\": 2}]}',1,1,8,2,0,1),(67,'道玄治愈之风',1,4,'ZhiYuZhiFeng','ZhiLiao',4,40,'单体治疗，优先治疗受伤最重的伙伴',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 4000, \"RawValueRate\": 0, \"AttackRate\": 30, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 100, \"ValueCountRate\": 0, \"TargetMode\": 2}]}',1,1,8,3,0,1),(68,'万象治愈之风',1,4,'ZhiYuZhiFeng','ZhiLiao',4,60,'单体治疗，优先治疗受伤最重的伙伴',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 0, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 6000, \"RawValueRate\": 0, \"AttackRate\": 30, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 100, \"ValueCountRate\": 0, \"TargetMode\": 2}]}',1,1,8,4,0,1),(69,'天人治愈之风',1,4,'ZhiYuZhiFeng','ZhiLiao',4,80,'单体治疗，优先治疗受伤最重的伙伴',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 8000, \"RawValueRate\": 0, \"AttackRate\": 30, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 100, \"ValueCountRate\": 0, \"TargetMode\": 2}]}',0,1,8,5,0,1),(70,'承天载物',1,4,'ChengTianZaiWu','ChengTianZaiWu',4,30,'全体治疗，产生大量治疗量',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 500, \"RawValueRate\": 0, \"AttackRate\": 30, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',2,0,70,1,0,2),(71,'金刚雨润',1,5,'YuRun','YuRun',5,20,'全体辅助，增加攻击1000、防御500，持续2回合',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 2, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 3, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',1,1,50,2,0,1),(72,'道玄雨润',1,5,'YuRun','YuRun',5,40,'全体辅助，增加攻击1500、防御1000，持续2回合',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 2, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 3, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',1,1,50,3,0,1),(73,'万象雨润',1,5,'YuRun','YuRun',5,60,'全体辅助，增加攻击2000、防御1500，持续2回合',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 2, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 2000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 3, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',1,1,50,4,0,1),(74,'天人雨润',1,5,'YuRun','YuRun',5,80,'全体辅助，增加攻击3000、防御2000，持续2回合',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 2, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 3000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 3, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 2000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',1,1,50,5,0,1),(75,'金刚青竹咒',1,1,'QinZhuZhou','JinGangQinZhuZhou',5,20,'单体攻击，附毒500，持续2回合',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 500,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 6, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',1,1,49,2,0,1),(76,'道玄青竹咒',1,1,'QinZhuZhou','JinGangQinZhuZhou',5,40,'单体攻击，附毒1000，持续2回合',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 6, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',1,1,49,3,0,1),(77,'万象青竹咒',1,1,'QinZhuZhou','JinGangQinZhuZhou',5,60,'单体攻击，附毒1500，持续2回合',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1500,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 6, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',1,1,49,4,0,1),(78,'天人青竹咒',1,1,'QinZhuZhou','JinGangQinZhuZhou',5,80,'单体攻击，附毒2000，持续2回合',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 6, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 2000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',1,1,49,5,0,1),(79,'金刚墨画狼群',1,1,'MoHuaYingLang','MoHuaYingLang',6,20,'横排穿透',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 500,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,51,2,0,3),(80,'道玄墨画狼群',1,1,'MoHuaYingLang','MoHuaYingLang',6,40,'横排穿透',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,51,3,0,3),(81,'万象墨画狼群',1,1,'MoHuaYingLang','MoHuaYingLang',6,60,'横排穿透',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1500,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,51,4,0,3),(82,'天人墨画狼群',1,1,'MoHuaYingLang','MoHuaYingLang',6,80,'横排穿透',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,51,5,0,3),(83,'金刚墨画巫雀',1,5,'MoHuaWuQue','MoHuaWuQue',6,20,'单体攻击，同时每次降低敌方800防御，持续1回合',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 500,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 1, \"BaseValue\": 800, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',1,1,52,2,0,1),(84,'道玄墨画巫雀',1,5,'MoHuaWuQue','MoHuaWuQue',6,40,'单体攻击，同时每次降低敌方1000防御，持续1回合',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 1, \"BaseValue\": 1000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',1,1,52,3,0,1),(85,'万象墨画巫雀',1,5,'MoHuaWuQue','MoHuaWuQue',6,60,'单体攻击，同时每次降低敌方1200防御，持续1回合',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1200,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 1, \"BaseValue\": 1200, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',1,1,52,4,0,1),(86,'天人墨画巫雀',1,5,'MoHuaWuQue','MoHuaWuQue',6,80,'单体攻击，同时每次降低敌方1400防御，持续1回合',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1200,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 1, \"BaseValue\": 1400, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',1,1,52,5,0,1),(87,'神魔封禁',1,5,'ShenMoFengJin','ShenMoFengJin',6,30,'全体攻击，并80%几率禁魔1回合',0,0,'{\"TargetMode\": 1,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 600,\"Cul2AtkRate\": 30,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 24, \"Keep\": 1, \"Override\": 1, \"Rate\": 80, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',2,0,87,1,0,2),(88,'圣白莲',1,5,'ShengBaiLian','ShengBaiLian',5,30,'解除负面状态，恢复生命',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 7, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 5000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',2,0,88,1,0,2),(89,'断岳',1,1,'DuanYue','DuanYue',-2,10,'纵向攻击',0,0,'{\"TargetMode\": 3,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 100,\"Cul2AtkRate\": 60,\"DecPower\": 2,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,0,89,1,0,4),(90,'金刚闪击',1,1,'ShanJi','ShanJi',-2,20,'单体攻击，获得3点精气',1,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1000,\"Cul2AtkRate\": 10,\"DecPower\": 0,\"IncPower\": 3,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,1,2,0,1),(91,'金刚星爆光离',1,1,'XingBaoGuangLi','XingBaoGuangLi',-2,20,'单体攻击，消耗精气产生大量伤害',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 5000,\"Cul2AtkRate\": 60,\"DecPower\": 6,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 50,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,4,2,0,1),(92,'金刚庇护',1,3,'JinGangBiHu','JinGangBiHu',-2,30,'防御1回合，自身减免60%的伤害，并吸引攻击。同时提高全体防御3000，持续1回合',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [{\"Type\": 18, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 60, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 60, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0},{\"Type\": 19, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 3000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',1,0,92,1,0,1),(93,'道玄闪击',1,1,'ShanJi','ShanJi',-2,40,'单体攻击，获得4点精气',1,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 10,\"DecPower\": 0,\"IncPower\": 4,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,1,3,0,1),(94,'道玄星爆光离',1,1,'XingBaoGuangLi','XingBaoGuangLi',-2,40,'单体攻击，消耗精气产生大量伤害',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 10000,\"Cul2AtkRate\": 60,\"DecPower\": 6,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 50,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,4,3,0,1),(95,'气疗术',1,4,'QiLiaoShu','ZhiLiao',-2,50,'单体治疗，治疗损失的生命最多的成员',0,0,'{\"TargetMode\": 4,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 1,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 4000, \"RawValueRate\": 0, \"AttackRate\": 10, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 60, \"ValueCountRate\": 0, \"TargetMode\": 2}]}',2,0,95,1,0,1),(96,'万象闪击',1,1,'ShanJi','ShanJi',-2,60,'单体攻击，获得5点精气',1,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 3000,\"Cul2AtkRate\": 10,\"DecPower\": 0,\"IncPower\": 5,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,1,4,0,1),(97,'万象星爆光离',1,1,'XingBaoGuangLi','XingBaoGuangLi',-2,60,'单体攻击，消耗精气产生大量伤害',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 15000,\"Cul2AtkRate\": 60,\"DecPower\": 6,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 50,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,4,4,0,1),(98,'天剑',1,1,'TianJian','TianJian',-2,70,'全体攻击，一剑开天门',0,0,'{\"TargetMode\": 1,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1000,\"Cul2AtkRate\": 60,\"DecPower\": 4,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',2,0,98,1,0,2),(99,'天人闪击',1,1,'ShanJi','ShanJi',-2,80,'单体攻击，获得6点精气',1,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 4000,\"Cul2AtkRate\": 10,\"DecPower\": 0,\"IncPower\": 6,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,1,5,0,1),(100,'天人星爆光离',1,1,'XingBaoGuangLi','XingBaoGuangLi',-2,80,'单体攻击，消耗精气产生大量伤害',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 20000,\"Cul2AtkRate\": 60,\"DecPower\": 6,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 50,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',1,1,4,5,0,1),(108,'治疗',5,4,'ZhiLiaoM','ZhiLiao',-1,0,'治疗生命最少的队友',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 0, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 100, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 2}]}',0,0,0,0,40,1),(109,'增益',5,5,'ZengYiM','ZengYi',-1,0,'自身',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 2, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 0, \"RawValueRate\": 5, \"AttackRate\": 0, \"SkillForceRate\": 100, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,0,0,0,5,1),(998,'破甲眩晕',0,0,'PoJiaXuanYun',NULL,-1,0,NULL,0,0,NULL,0,0,0,0,0,1),(999,'英勇',7,1,'YingYong','DuanYue',1,1,'纵向攻击，连续攻击2次',0,0,'{\"TargetMode\": 3,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1500,\"Cul2AtkRate\": 120,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 10,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 2,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0,4),(1000,'铜墙铁壁',7,3,'TongQiangTieBi','TongQiangTieBi',2,1,'护甲恢复50，减免50%伤害，并且防御增加500，持续1回合。',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 22, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 50, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 9, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 50, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,0,0,0,0,2),(1001,'回春',7,4,'HuiChun','HuiChun',3,1,'全体治疗血量增加3000，并去除全体异常状态。',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 3000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 7, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 0, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,0,0,0,0,2),(1002,'霸刀',7,1,'BaDao','BaDao',4,1,'高输出全体攻击。',0,0,'{\"TargetMode\": 1,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1500,\"Cul2AtkRate\": 120,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 50,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0,2),(1003,'轰击',7,5,'HongJi','HongJi',5,1,'全体攻击，降低敌方防御500，持续2回合。',0,0,'{\"TargetMode\": 1,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1500,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 3, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 1, \"BaseValue\": 500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',0,0,0,0,0,2),(1004,'人鱼之歌',7,5,'RenYuZhiGe','RenYuZhiGe',6,1,'全体辅助，全体攻击提升1000，降低敌方500防御，持续1回合。',0,0,'{\"TargetMode\": 1,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 1, \"BaseValue\": 500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": [{\"Type\": 2, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,0,0,0,0,2),(1005,'英勇2级',7,1,'YingYong','DuanYue',1,40,'纵向攻击，连续攻击2次',0,0,'{\"TargetMode\": 3,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 120,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 2,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0,4),(1006,'英勇3级',7,1,'YingYong','DuanYue',1,70,'纵向攻击，连续攻击2次',0,0,'{\"TargetMode\": 3,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2500,\"Cul2AtkRate\": 120,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 2,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0,4),(1007,'铜墙铁壁2级',7,3,'TongQiangTieBi','TongQiangTieBi',2,40,'护甲恢复100，减免60%伤害，并且防御增加800，持续1回合。',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 800, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 22, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 100, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 9, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 60, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,0,0,0,0,2),(1008,'铜墙铁壁3级',7,3,'TongQiangTieBi','TongQiangTieBi',2,70,'护甲恢复150，减免80%伤害，并且防御增加1000，持续1回合。',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 22, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 150, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 9, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 80, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,0,0,0,0,2),(1009,'回春2级',7,4,'HuiChun','HuiChun',3,40,'全体治疗血量增加3500，并去除全体异常状态。',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 3500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 7, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 0, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,0,0,0,0,2),(1010,'回春3级',7,4,'HuiChun','HuiChun',3,70,'全体治疗血量增加4000，并去除全体异常状态。',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 4000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 7, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 0, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,0,0,0,0,2),(1011,'霸刀2级',7,1,'BaDao','BaDao',4,40,'高输出全体攻击。',0,0,'{\"TargetMode\": 1,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 120,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 50,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0,2),(1012,'霸刀3级',7,1,'BaDao','BaDao',4,70,'高输出全体攻击。',0,0,'{\"TargetMode\": 1,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2500,\"Cul2AtkRate\": 120,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 50,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0,2),(1013,'轰击2级',7,5,'HongJi','HongJi',5,40,'全体攻击，降低敌方防御800，持续2回合。',0,0,'{\"TargetMode\": 1,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1500,\"Cul2AtkRate\": 120,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 3, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 1, \"BaseValue\": 800, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',0,0,0,0,0,2),(1014,'轰击3级',7,1,'HongJi','HongJi',5,70,'全体攻击，降低敌方防御1000，持续2回合。',0,0,'{\"TargetMode\": 1,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 1500,\"Cul2AtkRate\": 120,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 3, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 1, \"BaseValue\": 1000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',0,0,0,0,0,2),(1015,'人鱼之歌2级',7,5,'RenYuZhiGe','RenYuZhiGe',6,40,'全体辅助，全体攻击提升1500，降低敌方800防御，持续1回合。',0,0,'{\"TargetMode\": 1,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 1, \"BaseValue\": 800, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": [{\"Type\": 2, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,0,0,0,0,2),(1016,'人鱼之歌3级',7,5,'RenYuZhiGe','RenYuZhiGe',6,70,'全体辅助，全体攻击提升2000，降低敌方1000防御，持续1回合。',0,0,'{\"TargetMode\": 1,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 1, \"BaseValue\": 1000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": [{\"Type\": 2, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 2000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,0,0,0,0,2),(1017,'化劲',7,1,'HuoLieQuanTiM','HuaJing',7,1,'全体攻击，防御增加1000，持续2回合。',0,0,'{\"TargetMode\": 1,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 120,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 10,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 3, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,0,0,0,0,2),(1018,'化劲2级',7,1,'HuoLieQuanTiM','HuaJing',7,40,'全体攻击，防御增加1500，持续2回合。',0,0,'{\"TargetMode\": 1,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 120,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 10,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 3, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,0,0,0,0,2),(1019,'化劲3级',7,1,'HuoLieQuanTiM','HuaJing',7,70,'全体攻击，防御增加2000，持续2回合。',0,0,'{\"TargetMode\": 1,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 120,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 3, \"Keep\": 2, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 2000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,0,0,0,0,2),(1020,'致胜',7,1,'HengSaoQianJunM','HengSaoQianJun',8,1,'高输出横向攻击。',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 120,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 50,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 2,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0,3),(1021,'致胜2级',7,1,'HengSaoQianJunM','HengSaoQianJun',8,40,'高输出横向攻击。',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 120,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 100,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 2,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0,3),(1022,'致胜3级',7,1,'HengSaoQianJunM','HengSaoQianJun',8,70,'高输出横向攻击。',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 120,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 150,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 2,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0,3),(1023,'穿透',7,1,'WanJianChuanXinM','WanJianChuanXin',9,1,'横向攻击，暴击，全体治疗血量增加2000。',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 100,\"ReduceDefend\": 0,\"SunderAttack\": 10,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 2000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,0,0,0,0,3),(1024,'穿透2级',7,1,'WanJianChuanXinM','WanJianChuanXin',9,40,'横向攻击，暴击，全体治疗血量增加2500。',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 100,\"ReduceDefend\": 0,\"SunderAttack\": 10,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 2500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,0,0,0,0,3),(1025,'穿透3级',7,1,'WanJianChuanXinM','WanJianChuanXin',9,70,'横向攻击，暴击，全体治疗血量增加3000。',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 100,\"ReduceDefend\": 0,\"SunderAttack\": 10,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 3000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,0,0,0,0,3),(1026,'阿修罗之怒',7,1,'TuLieQuanTiM','TuLieQuanTi',10,1,'全体攻击，给敌方造成眩晕，防御降低1000，持续1回合。',0,0,'{\"TargetMode\": 1,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 3000,\"Cul2AtkRate\": 120,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 50,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 1, \"BaseValue\": 1000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0},{\"Type\": 5, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 0, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',0,0,0,0,0,2),(1027,'阿修罗之怒2级',7,1,'TuLieQuanTiM','TuLieQuanTi',10,40,'全体攻击，给敌方造成眩晕，防御降低1500，持续1回合。',0,0,'{\"TargetMode\": 1,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 120,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 50,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 1, \"BaseValue\": 1500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0},{\"Type\": 5, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 0, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',0,0,0,0,0,2),(1028,'阿修罗之怒3级',7,1,'TuLieQuanTiM','TuLieQuanTi',10,70,'全体攻击，给敌方造成眩晕，防御降低2000，持续1回合。',0,0,'{\"TargetMode\": 1,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 120,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 50,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 1, \"BaseValue\": 2000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0},{\"Type\": 5, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 0, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',0,0,0,0,0,2),(1029,'命运连锁',7,1,'PaoXiaoLiZhuaM','PaoXiaoLiZhua',12,1,'全体攻击，降低防御1000，禁魔1回合。',0,0,'{\"TargetMode\": 1,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 120,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 10,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 1, \"BaseValue\": 1000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0},{\"Type\": 24, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',0,0,0,0,0,2),(1030,'命运连锁2级',7,1,'PaoXiaoLiZhuaM','PaoXiaoLiZhua',12,40,'全体攻击，降低防御1500，禁魔1回合。',0,0,'{\"TargetMode\": 1,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 120,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 10,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 1, \"BaseValue\": 1500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0},{\"Type\": 24, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',0,0,0,0,0,2),(1031,'命运连锁3级',7,1,'PaoXiaoLiZhuaM','PaoXiaoLiZhua',12,70,'全体攻击，降低防御2000，禁魔1回合。',0,0,'{\"TargetMode\": 1,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 2000,\"Cul2AtkRate\": 120,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 1, \"BaseValue\": 2000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0},{\"Type\": 24, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0}],\"BuddyBuffs\": []}',0,0,0,0,0,2),(1032,'洛神',7,5,'SanQianLuoShuiJianM','SanQianLuoShuiJian',11,1,'全体辅助，我方全体治疗血量增加3000，攻击提升1000，防御提升1000，持续1回合。',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 3000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 33, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 2, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,0,0,0,0,2),(1033,'洛神2级',7,5,'SanQianLuoShuiJianM','SanQianLuoShuiJian',11,40,'全体辅助，我方全体治疗血量增加3500，攻击提升1500，防御提升1500，持续1回合。',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 3500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 33, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 2, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 1500, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,0,0,0,0,2),(1034,'洛神3级',7,5,'SanQianLuoShuiJianM','SanQianLuoShuiJian',11,70,'全体辅助，我方全体治疗血量增加4000，攻击提升2000，防御提升2000，持续1回合。',0,0,'{\"TargetMode\": 4,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": [{\"Type\": 4, \"Keep\": 0, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 4000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 33, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 2, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 2000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0},{\"Type\": 3, \"Keep\": 1, \"Override\": 1, \"Rate\": 100, \"CountRate\": 0, \"BuffSign\": 0, \"BaseValue\": 2000, \"RawValueRate\": 0, \"AttackRate\": 0, \"SkillForceRate\": 0, \"HurtRate\": 0, \"Cul2ValueRate\": 0, \"ValueCountRate\": 0, \"TargetMode\": 0}]}',0,0,0,0,0,2),(1035,'死亡阻击',5,1,'SiWangZuJiM','LiPiHuaShan',-1,0,'单体',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,0,1),(1036,'大风咒',1,1,'DaFengZhou','DaFengZhou',4,10,'横排攻击',0,0,'{\"TargetMode\": 2,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 200,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,1036,1,0,3),(1037,'多连斩',1,1,'DuoLianZhanM','DuoLianZhan',3,1,'单体攻击',0,0,'{\"TargetMode\": 0,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 50,\"Cul2AtkRate\": 60,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 5,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,1037,1,0,1),(1038,'紫电刀芒',5,1,'ZiDianDaoMangB','DaoMang',-1,0,'单体',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 25,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,1,1),(1039,'斩杀',5,1,'ZhanShaB','ZhanSha',-1,0,'单体',0,0,'{\"TargetMode\": 0,\"AttackMode\": 0,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 0,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 50,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,0,0,1,1),(1041,'聚气',0,5,'JuQi','JuQi',-2,0,'聚气凝神，获得精气。',0,0,'{\"TargetMode\": 4,\"AttackMode\": 2,\"KillSelfRate\": 1,\"DefaultAttack\": 0,\"Cul2AtkRate\": 0,\"DecPower\": 0,\"IncPower\": 2,\"HurtAdd\": 0,\"HurtAddRate\": 0,\"CureAdd\": 0,\"CureAddRate\": 0,\"Critial\": 0,\"ReduceDefend\": 0,\"SunderAttack\": 0,\"MustHit\": false,\"GhostOverrideBuddyBuff\": false,\"GhostOverrideSelfBuff\": false,\"GhostOverrideTargetBuff\": false,\"AttackNum\": 1,\"SelfBuffs\": [],\"TargetBuffs\": [],\"BuddyBuffs\": []}',0,0,1041,1,0,1);
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
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8mb4 COMMENT='绝招数据表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `skill_content` DISABLE KEYS */;
INSERT INTO `skill_content` VALUES (1,5,2,1),(2,6,1,0),(3,7,2,1),(4,8,3,2),(5,49,2,1),(6,50,1,2),(7,51,2,1),(8,52,1,2),(9,53,2,1),(10,54,2,1),(11,55,2,1),(12,56,2,1),(13,57,1,0),(14,58,1,0),(15,59,1,0),(16,60,1,0),(17,61,1,4),(18,62,2,1),(19,63,2,1),(20,64,2,1),(21,65,2,1),(22,66,3,2),(23,67,3,2),(24,68,3,2),(25,69,3,2),(26,70,1,3),(27,71,1,2),(28,72,1,2),(29,73,1,2),(30,74,1,2),(31,75,2,1),(32,76,2,1),(33,77,2,1),(34,78,2,1),(35,79,3,1),(36,80,3,1),(37,81,3,1),(38,82,3,1),(39,83,1,2),(40,84,1,2),(41,85,1,2),(42,86,1,2),(43,87,1,2),(44,88,1,3),(45,1036,2,1),(46,1037,2,1);
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
  `fragment_id` smallint(6) DEFAULT NULL COMMENT '兑换需要的碎片物品id',
  `only_exchange` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否只能兑换获得',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COMMENT='剑心';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `sword_soul` DISABLE KEYS */;
INSERT INTO `sword_soul` VALUES (2,1,'白虹','BaiHong','',2,0,0,0),(3,1,'断水','DuanShui','',3,0,0,0),(4,1,'龙泉','LongQuan','',4,0,0,0),(5,1,'巨阙','JuQue','',5,10,282,0),(6,2,'永用','YongYong','',2,0,0,0),(7,2,'青干','QingGan','',3,0,0,0),(8,2,'纯阳','ChunYang','',4,0,0,0),(9,2,'湛泸','ZhanLu','',5,10,283,0),(10,3,'惊鲵','JingNi','',4,0,0,0),(11,3,'赤霄','ChiXiao','',5,10,284,0),(12,5,'转魄','ZhuanPo','',3,0,0,0),(13,5,'工布','GongBu','',4,0,0,0),(14,5,'龙渊','LongYuan','',5,10,285,0),(15,4,'悬翦','XuanJian','',3,0,0,0),(16,4,'含光','HanGuang','',4,0,0,0),(17,4,'承影','ChengYing','',5,10,286,0),(18,7,'燕支','YanZhi','',3,0,0,0),(19,7,'灭魂','MieHun','',4,0,0,0),(20,7,'干将','GanJiang','',5,10,287,0),(21,11,'紫电','ZiDian','',3,0,0,0),(22,11,'宵练','XiaoLian','',4,0,0,0),(23,11,'莫邪','MoXie','',5,10,288,0),(24,8,'神龟','ShenGui','',3,0,0,0),(25,8,'北斗','BeiDou','',4,0,0,0),(26,8,'星河','XingHe','',5,10,289,0),(27,6,'百步','BaiBu','',3,0,0,0),(28,6,'鱼肠','YuChang','',4,0,0,0),(29,6,'春秋','ChunQiu','',5,10,290,0),(31,10,'无尘','WuChen','',4,0,0,0),(32,10,'青梅','QingMei','',5,10,291,0),(34,9,'真刚','ZhenGang','',4,0,0,0),(35,9,'竹马','ZhuMa','',5,10,292,0),(37,13,'潜龙','QianLong','',1,0,0,0);
/*!40000 ALTER TABLE `sword_soul` ENABLE KEYS */;
DROP TABLE IF EXISTS `sword_soul_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sword_soul_level` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `sword_soul_id` smallint(6) NOT NULL COMMENT '剑心ID',
  `level` tinyint(4) NOT NULL COMMENT '等级',
  `value` int(11) NOT NULL COMMENT '属性加值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=696 DEFAULT CHARSET=utf8mb4 COMMENT='剑心等级';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `sword_soul_level` DISABLE KEYS */;
INSERT INTO `sword_soul_level` VALUES (12,2,1,30),(13,2,2,60),(14,2,3,90),(15,2,4,120),(16,2,5,150),(17,2,6,210),(18,2,7,270),(19,2,8,330),(20,2,9,390),(21,2,10,450),(22,2,11,525),(23,3,1,75),(24,3,2,150),(25,3,3,225),(26,3,4,300),(27,3,5,375),(28,3,6,490),(29,3,7,605),(30,3,8,720),(31,3,9,835),(32,3,10,950),(33,3,11,1100),(34,4,1,150),(35,4,2,300),(36,4,3,450),(37,4,4,600),(38,4,5,750),(39,4,6,975),(40,4,7,1200),(41,4,8,1425),(42,4,9,1650),(43,4,10,1875),(44,4,11,2175),(47,5,1,300),(48,5,2,600),(49,5,3,900),(50,5,4,1200),(51,5,5,1500),(52,5,6,1950),(53,5,7,2400),(54,5,8,2850),(55,5,9,3300),(56,5,10,3750),(57,5,11,4350),(58,6,1,25),(59,6,2,50),(60,6,3,75),(61,6,4,100),(62,6,5,125),(63,6,6,165),(64,6,7,205),(65,6,8,245),(66,6,9,285),(67,6,10,325),(68,6,11,375),(69,7,1,50),(70,7,2,100),(71,7,3,150),(72,7,4,200),(73,7,5,250),(74,7,6,325),(75,7,7,400),(76,7,8,475),(77,7,9,550),(78,7,10,625),(79,7,11,725),(80,8,1,100),(81,8,2,200),(82,8,3,300),(83,8,4,400),(84,8,5,500),(85,8,6,650),(86,8,7,800),(87,8,8,950),(88,8,9,1100),(89,8,10,1250),(90,8,11,1450),(91,9,1,200),(92,9,2,400),(93,9,3,600),(94,9,4,800),(95,9,5,1000),(96,9,6,1300),(97,9,7,1600),(98,9,8,1900),(99,9,9,2200),(100,9,10,2500),(101,9,11,2900),(102,10,1,500),(103,10,2,1000),(104,10,3,1500),(105,10,4,2000),(106,10,5,2500),(107,10,6,3250),(108,10,7,4000),(109,10,8,4750),(110,10,9,5500),(111,10,10,6250),(112,10,11,7250),(113,11,1,1000),(114,11,2,2000),(115,11,3,3000),(116,11,4,4000),(117,11,5,5000),(118,11,6,6500),(119,11,7,8000),(120,11,8,9500),(121,11,9,11000),(122,11,10,12500),(123,11,11,14500),(124,12,1,50),(125,12,2,100),(126,12,3,150),(127,12,4,200),(128,12,5,250),(129,12,6,325),(130,12,7,400),(131,12,8,475),(132,12,9,550),(133,12,10,625),(134,12,11,725),(135,13,1,100),(136,13,2,200),(137,13,3,300),(138,13,4,400),(139,13,5,500),(140,13,6,650),(141,13,7,800),(142,13,8,950),(143,13,9,1100),(144,13,10,1250),(145,13,11,1450),(146,14,1,200),(147,14,2,400),(148,14,3,600),(149,14,4,800),(150,14,5,1000),(151,14,6,1300),(152,14,7,1600),(153,14,8,1900),(154,14,9,2200),(155,14,10,2500),(156,14,11,2900),(157,15,1,50),(158,15,2,100),(159,15,3,150),(160,15,4,200),(161,15,5,250),(162,15,6,325),(163,15,7,400),(164,15,8,475),(165,15,9,550),(166,15,10,625),(167,15,11,725),(168,16,1,100),(169,16,2,200),(170,16,3,300),(171,16,4,400),(172,16,5,500),(173,16,6,650),(174,16,7,800),(175,16,8,950),(176,16,9,1100),(177,16,10,1250),(178,16,11,1450),(179,17,1,200),(180,17,2,400),(181,17,3,600),(182,17,4,800),(183,17,5,1000),(184,17,6,1300),(185,17,7,1600),(186,17,8,1900),(187,17,9,2200),(188,17,10,2500),(189,17,11,2900),(190,18,1,5),(191,18,2,10),(192,18,3,15),(193,18,4,20),(194,18,5,25),(195,18,6,35),(196,18,7,45),(197,18,8,55),(198,18,9,65),(199,18,10,75),(200,18,11,85),(201,19,1,10),(202,19,2,20),(203,19,3,30),(204,19,4,40),(205,19,5,50),(206,19,6,65),(207,19,7,80),(208,19,8,95),(209,19,9,110),(210,19,10,125),(211,19,11,145),(212,20,1,20),(213,20,2,40),(214,20,3,60),(215,20,4,80),(216,20,5,100),(217,20,6,130),(218,20,7,160),(219,20,8,190),(220,20,9,220),(221,20,10,250),(222,20,11,290),(223,21,1,5),(224,21,2,10),(225,21,3,15),(226,21,4,20),(227,21,5,25),(228,21,6,35),(229,21,7,45),(230,21,8,55),(231,21,9,65),(232,21,10,75),(233,21,11,85),(234,22,1,10),(235,22,2,20),(236,22,3,30),(237,22,4,40),(238,22,5,50),(239,22,6,65),(240,22,7,80),(241,22,8,95),(242,22,9,110),(243,22,10,125),(244,22,11,145),(245,23,1,20),(246,23,2,40),(247,23,3,60),(248,23,4,80),(249,23,5,100),(250,23,6,130),(251,23,7,160),(252,23,8,190),(253,23,9,220),(254,23,10,250),(255,23,11,290),(256,24,1,5),(257,24,2,10),(258,24,3,15),(259,24,4,20),(260,24,5,25),(261,24,6,35),(262,24,7,45),(263,24,8,55),(264,24,9,65),(265,24,10,75),(266,24,11,85),(267,25,1,10),(268,25,2,20),(269,25,3,30),(270,25,4,40),(271,25,5,50),(272,25,6,65),(273,25,7,80),(274,25,8,95),(275,25,9,110),(276,25,10,125),(277,25,11,145),(278,26,1,20),(279,26,2,40),(280,26,3,60),(281,26,4,80),(282,26,5,100),(283,26,6,130),(284,26,7,160),(285,26,8,190),(286,26,9,220),(287,26,10,250),(288,26,11,290),(289,27,1,5),(290,27,2,10),(291,27,3,15),(292,27,4,20),(293,27,5,25),(294,27,6,35),(295,27,7,45),(296,27,8,55),(297,27,9,65),(298,27,10,75),(299,27,11,85),(300,28,1,10),(301,28,2,20),(302,28,3,30),(303,28,4,40),(304,28,5,50),(305,28,6,65),(306,28,7,80),(307,28,8,95),(308,28,9,110),(309,28,10,125),(310,28,11,145),(311,29,1,20),(312,29,2,40),(313,29,3,60),(314,29,4,80),(315,29,5,100),(316,29,6,130),(317,29,7,160),(318,29,8,190),(319,29,9,220),(320,29,10,250),(321,29,11,290),(333,31,1,10),(334,31,2,20),(335,31,3,30),(336,31,4,40),(337,31,5,50),(338,31,6,65),(339,31,7,80),(340,31,8,95),(341,31,9,110),(342,31,10,125),(343,31,11,145),(344,32,1,20),(345,32,2,40),(346,32,3,60),(347,32,4,80),(348,32,5,100),(349,32,6,130),(350,32,7,160),(351,32,8,190),(352,32,9,220),(353,32,10,250),(354,32,11,290),(366,34,1,10),(367,34,2,20),(368,34,3,30),(369,34,4,40),(370,34,5,50),(371,34,6,65),(372,34,7,80),(373,34,8,95),(374,34,9,110),(375,34,10,125),(376,34,11,145),(377,35,1,20),(378,35,2,40),(379,35,3,60),(380,35,4,80),(381,35,5,100),(382,35,6,130),(383,35,7,160),(384,35,8,190),(385,35,9,220),(386,35,10,250),(387,35,11,290),(398,37,1,1000),(408,2,12,600),(409,2,13,675),(410,2,14,750),(411,2,15,825),(412,2,16,920),(413,2,17,1015),(414,2,18,1110),(415,2,19,1205),(416,2,20,1300),(417,3,12,1250),(418,3,13,1400),(419,3,14,1550),(420,3,15,1700),(421,3,16,1890),(422,3,17,2080),(423,3,18,2270),(424,3,19,2460),(425,3,20,2650),(426,4,12,2475),(427,4,13,2775),(428,4,14,3075),(429,4,15,3375),(430,4,16,3750),(431,4,17,4125),(432,4,18,4500),(433,4,19,4875),(434,4,20,5250),(435,5,12,4950),(436,5,13,5550),(437,5,14,6150),(438,5,15,6700),(439,5,16,7450),(440,5,17,8200),(441,5,18,8950),(442,5,19,9700),(443,5,20,10450),(444,6,12,425),(445,6,13,475),(446,6,14,525),(447,6,15,575),(448,6,16,640),(449,6,17,705),(450,6,18,770),(451,6,19,835),(452,6,20,900),(453,7,12,825),(454,7,13,925),(455,7,14,1025),(456,7,15,1125),(457,7,16,1250),(458,7,17,1375),(459,7,18,1500),(460,7,19,1625),(461,7,20,1750),(462,8,12,1650),(463,8,13,1850),(464,8,14,2050),(465,8,15,2250),(466,8,16,2500),(467,8,17,2750),(468,8,18,3000),(469,8,19,3250),(470,8,20,3500),(471,9,12,3300),(472,9,13,3700),(473,9,14,4100),(474,9,15,4500),(475,9,16,5000),(476,9,17,5500),(477,9,18,6000),(478,9,19,6500),(479,9,20,7000),(480,10,12,8250),(481,10,13,9250),(482,10,14,10250),(483,10,15,11250),(484,10,16,12500),(485,10,17,13750),(486,10,18,15000),(487,10,19,16250),(488,10,20,17500),(489,11,12,16500),(490,11,13,18500),(491,11,14,20500),(492,11,15,22500),(493,11,16,25000),(494,11,17,27500),(495,11,18,30000),(496,11,19,32500),(497,11,20,35000),(498,12,12,825),(499,12,13,925),(500,12,14,1025),(501,12,15,1125),(502,12,16,1250),(503,12,17,1375),(504,12,18,1500),(505,12,19,1625),(506,12,20,1750),(507,13,12,1650),(508,13,13,1850),(509,13,14,2050),(510,13,15,2250),(511,13,16,2500),(512,13,17,2750),(513,13,18,3000),(514,13,19,3250),(515,13,20,3500),(516,14,12,3300),(517,14,13,3700),(518,14,14,4100),(519,14,15,4500),(520,14,16,5000),(521,14,17,5500),(522,14,18,6000),(523,14,19,6500),(524,14,20,7000),(525,15,12,825),(526,15,13,925),(527,15,14,1025),(528,15,15,1125),(529,15,16,1250),(530,15,17,1375),(531,15,18,1500),(532,15,19,1625),(533,15,20,1750),(534,16,12,1650),(535,16,13,1850),(536,16,14,2050),(537,16,15,2250),(538,16,16,2500),(539,16,17,2750),(540,16,18,3000),(541,16,19,3250),(542,16,20,3500),(543,17,12,3300),(544,17,13,3700),(545,17,14,4100),(546,17,15,4500),(547,17,16,5000),(548,17,17,5500),(549,17,18,6000),(550,17,19,6500),(551,17,20,7000),(552,18,12,95),(553,18,13,105),(554,18,14,115),(555,18,15,125),(556,18,16,140),(557,18,17,155),(558,18,18,170),(559,18,19,185),(560,18,20,200),(561,19,12,165),(562,19,13,185),(563,19,14,205),(564,19,15,225),(565,19,16,250),(566,19,17,275),(567,19,18,300),(568,19,19,325),(569,19,20,350),(570,20,12,330),(571,20,13,370),(572,20,14,410),(573,20,15,450),(574,20,16,500),(575,20,17,550),(576,20,18,600),(577,20,19,650),(578,20,20,700),(579,21,12,95),(580,21,13,105),(581,21,14,115),(582,21,15,125),(583,21,16,140),(584,21,17,155),(585,21,18,170),(586,21,19,185),(587,21,20,200),(588,22,12,165),(589,22,13,185),(590,22,14,205),(591,22,15,225),(592,22,16,250),(593,22,17,275),(594,22,18,300),(595,22,19,325),(596,22,20,350),(597,23,12,330),(598,23,13,370),(599,23,14,410),(600,23,15,450),(601,23,16,500),(602,23,17,550),(603,23,18,600),(604,23,19,650),(605,23,20,700),(606,24,12,95),(607,24,13,105),(608,24,14,115),(609,24,15,125),(610,24,16,140),(611,24,17,155),(612,24,18,170),(613,24,19,185),(614,24,20,200),(615,25,12,165),(616,25,13,185),(617,25,14,205),(618,25,15,225),(619,25,16,250),(620,25,17,275),(621,25,18,300),(622,25,19,325),(623,25,20,350),(624,26,12,330),(625,26,13,370),(626,26,14,410),(627,26,15,450),(628,26,16,500),(629,26,17,550),(630,26,18,600),(631,26,19,650),(632,26,20,700),(633,27,12,95),(634,27,13,105),(635,27,14,115),(636,27,15,125),(637,27,16,140),(638,27,17,155),(639,27,18,170),(640,27,19,185),(641,27,20,200),(642,28,12,165),(643,28,13,185),(644,28,14,205),(645,28,15,225),(646,28,16,250),(647,28,17,275),(648,28,18,300),(649,28,19,325),(650,28,20,350),(651,29,12,330),(652,29,13,370),(653,29,14,410),(654,29,15,450),(655,29,16,500),(656,29,17,550),(657,29,18,600),(658,29,19,650),(659,29,20,700),(660,31,12,165),(661,31,13,185),(662,31,14,205),(663,31,15,225),(664,31,16,250),(665,31,17,275),(666,31,18,300),(667,31,19,325),(668,31,20,350),(669,32,12,330),(670,32,13,370),(671,32,14,410),(672,32,15,450),(673,32,16,500),(674,32,17,550),(675,32,18,600),(676,32,19,650),(677,32,20,700),(678,34,12,165),(679,34,13,185),(680,34,14,205),(681,34,15,225),(682,34,16,250),(683,34,17,275),(684,34,18,300),(685,34,19,325),(686,34,20,350),(687,35,12,330),(688,35,13,370),(689,35,14,410),(690,35,15,450),(691,35,16,500),(692,35,17,550),(693,35,18,600),(694,35,19,650),(695,35,20,700);
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
INSERT INTO `sword_soul_quality` VALUES (0,'杂物','NULL','0xc5c3b7'),(1,'特殊','SPECIAL','0xfff100'),(2,'优良','FINE','0x22ac38'),(3,'精良','EXCELLENT','0x00a0e9'),(4,'传奇','LEGEND','0xc301c3'),(5,'神器','ARTIFACT','0xfff100'),(6,'唯一','ONLY','0xf39700');
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
) ENGINE=InnoDB AUTO_INCREMENT=102 DEFAULT CHARSET=utf8mb4 COMMENT='剑心品质';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `sword_soul_quality_level` DISABLE KEYS */;
INSERT INTO `sword_soul_quality_level` VALUES (1,1,1,1000),(2,2,1,100),(3,2,2,100),(4,2,3,250),(5,2,4,500),(6,2,5,1000),(7,2,6,2000),(8,2,7,4000),(9,2,8,5000),(10,2,9,6000),(11,2,10,7000),(12,3,1,200),(13,3,2,250),(14,3,3,500),(15,3,4,1000),(16,3,5,2000),(17,3,6,4000),(18,3,7,5000),(19,3,8,6000),(20,3,9,7000),(21,3,10,8000),(22,4,1,500),(23,4,2,500),(24,4,3,1000),(25,4,4,2000),(26,4,5,4000),(27,4,6,5000),(28,4,7,6000),(29,4,8,7000),(30,4,9,8000),(31,4,10,9000),(32,5,1,1500),(33,5,2,1000),(34,5,3,2000),(35,5,4,4000),(36,5,5,5000),(37,5,6,6000),(38,5,7,7000),(39,5,8,8000),(40,5,9,9000),(41,5,10,10000),(42,6,1,3000),(43,6,2,2000),(44,6,3,4000),(45,6,4,5000),(46,6,5,6000),(47,6,6,7000),(48,6,7,8000),(49,6,8,9000),(50,6,9,10000),(51,6,10,11000),(52,2,11,8000),(53,2,12,9000),(54,2,13,10000),(55,2,14,11000),(56,2,15,12000),(57,2,16,13000),(58,2,17,14000),(59,2,18,15000),(60,2,19,16000),(61,2,20,17000),(62,3,11,9000),(63,3,12,10000),(64,3,13,11000),(65,3,14,12000),(66,3,15,13000),(67,3,16,14000),(68,3,17,15000),(69,3,18,16000),(70,3,19,17000),(71,3,20,18000),(72,4,11,10000),(73,4,12,11000),(74,4,13,12000),(75,4,14,13000),(76,4,15,14000),(77,4,16,15000),(78,4,17,16000),(79,4,18,17000),(80,4,19,18000),(81,4,20,19000),(82,5,11,11000),(83,5,12,12000),(84,5,13,13000),(85,5,14,14000),(86,5,15,15000),(87,5,16,16000),(88,5,17,17000),(89,5,18,18000),(90,5,19,19000),(91,5,20,20000),(92,6,11,12000),(93,6,12,13000),(94,6,13,14000),(95,6,14,15000),(96,6,15,16000),(97,6,16,17000),(98,6,17,18000),(99,6,18,19000),(100,6,19,20000),(101,6,20,21000);
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
INSERT INTO `town` VALUES (1,100000,'神龙岛','ShenLongDao','ZhuTiQu',1589,960),(2,100110,'烛堡','ZhuBao','ZhuTiQu',1369,760),(3,999999,'测试城镇','ShenLongDao','ZhuTiQu',1282,529),(4,100120,'铸剑山庄','ZhuJianShanZhuang','ZhuTiQu',1369,760);
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
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COMMENT='城镇NPC';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `town_npc` DISABLE KEYS */;
INSERT INTO `town_npc` VALUES (1,1,'龙姬','LongJi',1670,1280,'b','阴影这么多该怎么办！||我感觉体内有股奇怪的力量。'),(2,1,'梦妖','MengYao',1560,1350,'rb','你见过神龙么？||没人知道我们梦妖一族从何而来'),(3,1,'商人','ShangRen',1167,1023,'rb','江湖险恶，常备良药才是保命上策||做点小生意不容易，多来光顾'),(4,1,'梦妖小清','MengYao',1064,345,'lb','我好想离开神龙岛去外面的世界||如果可以像个大侠一样去冒险……'),(5,1,'梦妖小辉','MengYao',1296,1352,'l','小杰说他是最熟悉神龙岛的人。||你想知道什么？'),(6,1,'梦妖小杰','MengYao',1180,1352,'r','我对神龙岛了如指掌。||传说在火山的内部有一个密室。||我知道一个长生不老的秘密。'),(7,1,'梦妖小立','MengYao',1730,600,'t','我躲在这里谁也找不到我。||玩捉迷藏我最厉害了。||他们找了我两天都没找到我。||今天吃了小苹果，感觉自己萌萌哒'),(8,1,'梦妖阿剑','MengYao',2030,823,'lb','我也是出过岛的妖！||江湖上可是流传着我的传说。||我在黑夜森林看见了奇怪的黑影。'),(9,1,'梦妖阿樂','MengYao',611,557,'l','宝剑尚未配妥，出门已是江湖。||我有一天梦见了上古的剑灵。'),(10,2,'商人','ShangRen',1636,1000,'b','江湖险恶，常备各种良药才是保命上策。||做点小生意不容易，多来光顾呀！'),(11,1,'朱媛媛','ZhuYuanYuan',1402,1263,'rb','今天去哪里玩？'),(12,1,'袁铭志','YuanMingZhi',1050,759,'r','去给我找点酒。');
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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COMMENT='城镇NPC对话';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `town_npc_item` DISABLE KEYS */;
INSERT INTO `town_npc_item` VALUES (2,3,209,-1,0),(3,3,250,-1,0),(5,3,264,-1,0),(6,3,265,-1,0);
/*!40000 ALTER TABLE `town_npc_item` ENABLE KEYS */;
DROP TABLE IF EXISTS `trader`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `trader` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(10) NOT NULL COMMENT 'NPC名称',
  `sign` varchar(20) NOT NULL COMMENT '资源标识',
  `talk` varchar(200) NOT NULL COMMENT '对话',
  `sold_out_talk` varchar(200) NOT NULL COMMENT '售罄对话',
  `deal_talk` varchar(200) NOT NULL COMMENT '成交对话',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='随机商店商人';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `trader` DISABLE KEYS */;
INSERT INTO `trader` VALUES (1,'瀛海集市','YingHaiJiShi','我是神秘商人荧惑，\\n请多关照。！||你是来看货还是来看我的，\\n我要保持神秘。||我会在每天的12点和18点换上新的货物。||你在点哪里？','这批货物已经售空了。\\n大侠...你还要吗？||看来这些商品很受欢迎。','谢谢惠顾，\\n大侠很有眼光。||看来这个商品很受欢迎。'),(2,'巡游商人','XunYouShangRen','我带来了一点新东西，\\n希望你喜欢。||我是真真，\\n我喜欢探索世界。||小天是我的好伙伴哦～\\n我要永远和小天在一起。||我走过了世界各地，\\n我知道什么是好东西～','我很高兴你喜欢我的东西。||我下次再来的时候不要错过哦。','成交！||你的选择非常正确！'),(3,'黑市老大','HeiShiLaoDa','一手交钱，一手交货。||我叫挺爷，我可以给你你想要的东西。||晚上21点再来看看新东西，\\n不要向任何人提起我。||我的宠物经记住你的气味。||时间就是金钱。','只要价钱合适，我可以马上给你搞来一批新货。||这个宝贝炙手可热。','你很有眼光。||一手交钱，一手交货。');
/*!40000 ALTER TABLE `trader` ENABLE KEYS */;
DROP TABLE IF EXISTS `trader_extra_talk`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `trader_extra_talk` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `trader_id` smallint(6) NOT NULL COMMENT '随机商人ID',
  `time` tinyint(4) NOT NULL COMMENT '点击次数',
  `talk` varchar(200) NOT NULL COMMENT '对话',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='随机商店商人额外对话';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `trader_extra_talk` DISABLE KEYS */;
INSERT INTO `trader_extra_talk` VALUES (1,2,20,'不要再点我了！小天会生气了的！'),(2,2,1,'小天和真真是好朋友。'),(3,2,5,'小天会保护真真探索世界的。'),(4,2,10,'为什么要一直点我！');
/*!40000 ALTER TABLE `trader_extra_talk` ENABLE KEYS */;
DROP TABLE IF EXISTS `trader_grid`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `trader_grid` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `trader_id` smallint(6) NOT NULL COMMENT '随机商人ID',
  `money_type` tinyint(4) NOT NULL COMMENT '货币类型',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COMMENT='随机商店货物配置';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `trader_grid` DISABLE KEYS */;
INSERT INTO `trader_grid` VALUES (1,1,1),(2,1,2),(3,1,3),(4,1,0),(5,1,0),(6,3,1),(7,2,0),(8,1,0),(9,3,1),(10,3,3),(11,3,1),(12,3,2),(13,3,1),(14,3,1),(15,3,0),(17,3,1),(18,3,1),(19,3,1),(20,3,1),(21,2,1),(22,2,2),(23,2,3),(24,2,1),(25,2,1),(26,2,1),(27,2,0);
/*!40000 ALTER TABLE `trader_grid` ENABLE KEYS */;
DROP TABLE IF EXISTS `trader_grid_config`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `trader_grid_config` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `grid_id` int(11) NOT NULL COMMENT '格子ID',
  `item_id` smallint(6) NOT NULL COMMENT '物品ID',
  `num` smallint(6) NOT NULL COMMENT '物品数量',
  `probability` tinyint(4) NOT NULL COMMENT '出现概率（％）',
  `cost` bigint(20) NOT NULL COMMENT '价格',
  `goods_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '货物类型0-物品 1-爱心 2-剑心 3-魂侍',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=137 DEFAULT CHARSET=utf8mb4 COMMENT='随机商店货物配置';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `trader_grid_config` DISABLE KEYS */;
INSERT INTO `trader_grid_config` VALUES (1,1,263,3,10,40,0),(2,1,251,1,15,80,0),(3,1,259,1,5,40,0),(4,1,260,1,5,40,0),(5,1,261,1,5,40,0),(13,1,256,1,20,40,0),(14,1,257,1,20,40,0),(15,1,258,1,20,40,0),(16,2,302,1,50,10,0),(17,2,210,1,50,30,0),(18,3,214,1,50,50,0),(19,4,239,1,25,8000,0),(20,4,235,1,20,1500,0),(21,4,236,1,15,1500,0),(22,4,36,1,15,2000,0),(23,4,38,1,15,2000,0),(24,5,231,5,25,5000,0),(25,5,237,1,20,1000,0),(26,5,238,1,15,1000,0),(27,5,232,1,15,5000,0),(28,5,253,1,15,10000,0),(29,5,254,1,10,15000,0),(30,6,217,1,20,40,0),(31,7,219,1,50,100000,0),(32,8,230,1,25,8000,0),(33,8,33,1,20,2000,0),(34,8,32,1,15,2000,0),(35,8,31,1,15,2000,0),(36,8,34,1,15,2000,0),(37,8,215,1,10,12000,0),(38,6,215,1,20,40,0),(39,6,216,1,20,40,0),(40,6,210,1,15,50,0),(41,6,211,1,20,120,0),(42,6,212,1,5,280,0),(44,9,282,2,15,288,0),(45,9,283,2,15,288,0),(46,9,284,2,15,288,0),(47,9,286,2,10,288,0),(48,9,287,2,15,288,0),(49,9,288,2,10,288,0),(50,9,290,2,5,288,0),(51,9,292,2,5,288,0),(52,9,291,2,5,288,0),(53,9,289,2,5,288,0),(54,10,125,1,10,200,0),(55,10,126,1,10,200,0),(56,10,127,1,10,200,0),(57,10,128,1,10,200,0),(58,10,130,1,10,180,0),(59,10,131,1,10,180,0),(60,10,132,1,10,180,0),(61,11,273,3,20,240,0),(62,11,274,3,20,240,0),(63,11,275,3,20,240,0),(64,11,276,3,10,300,0),(65,11,277,3,15,240,0),(66,11,278,3,15,240,0),(67,13,92,3,50,288,3),(68,13,93,3,50,140,3),(69,15,271,1,50,15000,0),(70,15,272,1,50,15000,0),(71,12,302,10,100,100,0),(72,14,263,30,50,180,0),(73,14,233,5,50,200,0),(74,10,227,1,2,1200,0),(75,10,162,1,10,250,0),(76,10,163,1,18,250,0),(77,16,5,1,5,1400,1),(78,16,14,1,5,1400,1),(79,16,35,1,5,1400,1),(80,16,20,1,5,1400,1),(81,16,32,1,5,1400,1),(82,16,23,1,5,1400,1),(83,16,29,1,5,1400,1),(84,16,26,1,5,1400,1),(85,16,17,1,5,1400,1),(86,16,4,1,10,80,1),(87,16,7,1,10,80,1),(88,16,9,1,5,1400,1),(89,16,8,1,10,80,1),(90,16,11,1,5,1400,1),(91,16,10,1,10,80,1),(92,16,37,20,10,2000,1),(93,16,13,1,10,80,1),(94,17,37,20,8,2000,1),(95,17,5,1,5,1400,1),(96,17,9,1,5,1400,1),(97,17,11,1,5,1400,1),(98,17,17,1,5,1400,1),(99,17,14,1,5,1400,1),(100,17,20,1,5,1400,1),(101,17,35,1,5,1400,1),(102,17,32,1,5,1400,1),(103,17,29,1,5,1400,1),(104,17,26,1,5,1400,1),(105,17,23,1,5,1400,1),(106,17,34,1,5,80,1),(107,17,31,1,8,80,1),(108,17,28,1,8,80,1),(109,17,25,1,8,80,1),(110,17,10,1,8,80,1),(111,18,7,1,5,4000,2),(112,18,8,1,8,4000,2),(113,18,9,1,8,4000,2),(114,18,4,1,25,400,2),(115,18,6,1,25,400,2),(116,18,5,1,25,400,2),(117,18,12,1,4,4000,2),(118,19,251,5,50,400,0),(119,19,263,10,50,60,0),(120,20,302,10,95,200,0),(121,20,303,1,5,40,0),(122,21,276,50,50,5000,0),(123,21,277,50,50,4000,0),(124,4,254,1,10,12000,0),(125,7,224,1,50,80000,0),(126,22,263,5,50,10,0),(127,22,251,1,50,10,0),(128,23,234,1,100,400,0),(129,24,37,20,100,2000,1),(130,25,233,10,100,800,0),(131,26,289,10,25,1400,0),(132,26,290,10,25,1400,0),(133,26,291,10,25,1400,0),(134,26,292,10,25,1400,0),(135,27,303,1,100,50000,0),(136,3,0,5,50,20,4);
/*!40000 ALTER TABLE `trader_grid_config` ENABLE KEYS */;
DROP TABLE IF EXISTS `trader_position`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `trader_position` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `trader_id` smallint(6) NOT NULL COMMENT '随机商人ID',
  `town_id` smallint(6) NOT NULL COMMENT '城镇ID',
  `x` int(11) NOT NULL COMMENT 'x轴坐标',
  `y` int(11) NOT NULL COMMENT 'y轴坐标',
  `direction` varchar(20) DEFAULT NULL COMMENT '朝向',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='随机商店商人坐标';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `trader_position` DISABLE KEYS */;
/*!40000 ALTER TABLE `trader_position` ENABLE KEYS */;
DROP TABLE IF EXISTS `trader_refresh_price`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `trader_refresh_price` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `trader_id` smallint(6) NOT NULL COMMENT '随机商人ID',
  `time` smallint(6) NOT NULL COMMENT '点击次数',
  `price` bigint(20) NOT NULL COMMENT '价格',
  PRIMARY KEY (`id`),
  UNIQUE KEY `time` (`trader_id`,`time`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COMMENT='随机商店商人额外对话';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `trader_refresh_price` DISABLE KEYS */;
INSERT INTO `trader_refresh_price` VALUES (1,1,1,40),(2,1,2,40),(3,1,3,80),(4,1,4,80),(5,1,5,120),(6,1,6,120),(7,1,7,120),(8,1,8,160),(9,1,9,160),(10,1,10,200),(11,1,11,200),(12,1,12,200),(13,1,13,400),(14,3,1,100),(15,3,2,100),(16,3,3,200),(17,3,4,200),(18,3,5,200),(19,3,6,400),(20,3,7,400),(21,3,8,400),(22,3,9,800),(23,3,10,800);
/*!40000 ALTER TABLE `trader_refresh_price` ENABLE KEYS */;
DROP TABLE IF EXISTS `vip_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `vip_level` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `level` smallint(6) NOT NULL COMMENT 'VIP等级',
  `ingot` bigint(20) NOT NULL COMMENT '累计充值元宝要求',
  PRIMARY KEY (`id`),
  UNIQUE KEY `level` (`level`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COMMENT='VIP等级数值';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `vip_level` DISABLE KEYS */;
INSERT INTO `vip_level` VALUES (1,0,0),(2,1,10),(3,2,100),(4,3,300),(5,4,500),(6,5,1000),(7,6,2000),(8,7,3000),(9,8,5000),(10,9,7000),(11,10,10000),(12,11,20000),(13,12,50000),(14,13,70000),(15,14,100000),(16,15,150000);
/*!40000 ALTER TABLE `vip_level` ENABLE KEYS */;
DROP TABLE IF EXISTS `vip_privilege`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `vip_privilege` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(20) NOT NULL COMMENT '特权名称',
  `sign` varchar(20) NOT NULL COMMENT '唯一标识',
  `tip` varchar(200) NOT NULL COMMENT '特权描述',
  `order` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `sign` (`sign`),
  UNIQUE KEY `order` (`order`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4 COMMENT='玩家VIP特权表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `vip_privilege` DISABLE KEYS */;
INSERT INTO `vip_privilege` VALUES (1,'爱心福利','AiXinFuLi','每天将通过邮件形式赠送爱心',1),(3,'扫荡特权','SaoDangTeQuan','开启普通、难度关卡扫荡功能',5),(8,'购买铜钱','GouMaiTongQian','每天给予额外的铜钱购买次数',2),(15,'购买体力','GouMaiTiLi','每天给予额外的体力购买次数',3),(23,'比武场特权','BiWuChangTeQuan','可使用元宝清除比武场冷却时间',4),(24,'爱心关卡','AiXinGuanKa','可每天额外参与1次爱心关卡',8),(26,'影界刷新','YinJieShuaXin','每天可刷新影界',9),(27,'爱心抽奖','AiXinChouJiang','提高每天爱心抽奖次数上限',6),(28,'批量兑换','PiLiangDuiHuan','开元通宝可使用批量兑换',10),(29,'彩虹特权','CaiHongTeQuan','每天可额外参加1次彩虹关卡',12),(30,'瀛海集市','YingHaiJiShi','永久开启瀛海集市刷新功能',7),(32,'黑市老大','HeiShiLaoDa','永久开启黑市老大刷新次数',11);
/*!40000 ALTER TABLE `vip_privilege` ENABLE KEYS */;
DROP TABLE IF EXISTS `vip_privilege_config`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `vip_privilege_config` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `privilege_id` int(11) NOT NULL COMMENT '特权ID',
  `level` smallint(6) NOT NULL COMMENT 'VIP等级',
  `times` smallint(6) NOT NULL DEFAULT '0' COMMENT '特权次数',
  `unit` varchar(4) NOT NULL DEFAULT '' COMMENT '特权次数单位',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_privilege_each_level` (`privilege_id`,`level`)
) ENGINE=InnoDB AUTO_INCREMENT=174 DEFAULT CHARSET=utf8mb4 COMMENT='玩家VIP特权表配置表';
/*!40101 SET character_set_client = @saved_cs_client */;

/*!40000 ALTER TABLE `vip_privilege_config` DISABLE KEYS */;
INSERT INTO `vip_privilege_config` VALUES (1,1,1,5,'个'),(2,15,1,2,'次'),(3,8,1,2,'次'),(4,23,1,0,''),(5,3,1,0,''),(6,1,2,6,'个'),(7,15,2,3,'次'),(8,8,2,3,'次'),(9,27,2,11,'次'),(10,23,2,0,''),(11,3,2,0,''),(12,1,3,7,'个'),(13,15,3,4,'次'),(14,8,3,5,'次'),(15,27,3,12,'次'),(16,23,3,0,''),(17,3,3,0,''),(30,1,4,8,'个'),(31,15,4,5,'次'),(32,8,4,10,'次'),(33,27,4,13,'次'),(34,23,4,0,''),(35,3,4,0,''),(36,24,4,1,'次'),(37,1,5,9,'个'),(38,15,5,6,'次'),(39,8,5,15,'次'),(40,27,5,14,'次'),(41,23,5,0,''),(42,3,5,0,''),(43,24,5,1,'次'),(44,26,5,2,'次'),(53,1,6,10,'个'),(54,15,6,7,'次'),(55,8,6,20,'次'),(56,27,6,15,'次'),(57,23,6,0,''),(58,3,6,0,''),(59,24,6,1,'次'),(60,26,6,3,'次'),(69,1,7,11,'个'),(70,15,7,8,'次'),(71,8,7,25,'次'),(72,27,7,16,'次'),(73,23,7,0,''),(74,3,7,0,''),(75,24,7,1,'次'),(76,26,7,4,'次'),(77,1,8,12,'个'),(78,15,8,9,'次'),(79,8,8,30,'次'),(80,27,8,17,'次'),(81,23,8,0,''),(82,3,8,0,''),(83,24,8,1,'次'),(84,26,8,5,'次'),(85,28,8,0,''),(86,1,9,13,'个'),(87,15,9,10,'次'),(88,8,9,35,'次'),(89,27,9,18,'次'),(90,23,9,0,''),(91,3,9,0,''),(92,24,9,1,'次'),(93,26,9,5,'次'),(94,28,9,0,''),(95,1,10,14,'个'),(96,15,10,15,'次'),(97,8,10,40,'次'),(98,27,10,19,'次'),(99,23,10,0,''),(100,3,10,0,''),(101,24,10,1,'次'),(102,26,10,5,'次'),(103,28,10,0,''),(104,1,11,15,'个'),(105,15,11,20,'次'),(106,8,11,45,'次'),(107,27,11,20,'次'),(108,23,11,0,''),(109,3,11,0,''),(110,24,11,1,'次'),(111,26,11,7,'次'),(112,28,11,0,''),(113,29,11,1,'次'),(114,1,12,15,'个'),(115,15,12,25,'次'),(116,8,12,50,'次'),(117,27,12,20,'次'),(118,23,12,0,''),(119,3,12,0,''),(120,24,12,1,'次'),(121,26,12,8,'次'),(122,28,12,0,''),(123,29,12,1,'次'),(124,1,13,15,'个'),(125,15,13,30,'次'),(126,8,13,60,'次'),(127,27,13,20,'次'),(128,23,13,0,''),(129,3,13,0,''),(130,24,13,1,'次'),(131,26,13,9,'次'),(132,28,13,0,''),(133,29,13,1,'次'),(134,1,14,15,'个'),(135,15,14,35,'次'),(136,8,14,70,'次'),(137,27,14,20,'次'),(138,23,14,0,''),(139,3,14,0,''),(140,24,14,1,'次'),(141,26,14,10,'次'),(142,28,14,0,''),(143,29,14,1,'次'),(144,1,15,15,'个'),(145,15,15,40,'次'),(146,8,15,80,'次'),(147,27,15,20,'次'),(148,23,15,0,''),(149,3,15,0,''),(150,24,15,1,'次'),(151,26,15,11,'次'),(152,28,15,0,''),(153,29,15,1,'次'),(154,26,4,1,'次'),(155,30,3,1,'次'),(156,30,4,2,'次'),(157,30,5,3,'次'),(158,30,6,4,'次'),(159,30,7,5,'次'),(160,30,8,6,'次'),(161,30,9,7,'次'),(162,30,10,8,'次'),(163,30,11,9,'次'),(164,30,12,10,'次'),(165,30,13,11,'次'),(166,30,14,12,'次'),(167,30,15,13,'次'),(168,32,10,5,'次'),(169,32,11,6,'次'),(170,32,12,7,'次'),(171,32,13,8,'次'),(172,32,14,9,'次'),(173,32,15,10,'次');
/*!40000 ALTER TABLE `vip_privilege_config` ENABLE KEYS */;
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
) ENGINE=InnoDB AUTO_INCREMENT=8589934624 DEFAULT CHARSET=utf8mb4 COMMENT='公告列表';
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
  `battle_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最近一次挑战时间',
  `win_times` smallint(6) NOT NULL DEFAULT '0' COMMENT '>0 连胜场次; 0 保持不变; -1 下降趋势',
  `daily_award_longbi` int(11) NOT NULL DEFAULT '0' COMMENT '今日获得龙币累计',
  `new_record_count` smallint(6) NOT NULL COMMENT '新战报计数',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家比武场数据';
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
) ENGINE=InnoDB AUTO_INCREMENT=47244640261 DEFAULT CHARSET=utf8mb4 COMMENT='玩家比武场记录';
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
) ENGINE=InnoDB AUTO_INCREMENT=47244640290 DEFAULT CHARSET=utf8mb4 COMMENT='玩家灵宠数据';
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
  `is_first_coin_ten` tinyint(4) NOT NULL COMMENT '是否第一次青龙宝箱十连抽',
  `coin_chest_consume` bigint(20) NOT NULL COMMENT '今天开青铜宝箱花费铜钱数',
  `last_coin_chest_at` bigint(20) NOT NULL COMMENT '上次开消费青铜宝箱时间',
  `last_free_ingot_chest_at` bigint(20) NOT NULL COMMENT '上次开免费神龙宝箱时间',
  `ingot_chest_num` int(11) NOT NULL COMMENT '今天开神龙宝箱次数',
  `ingot_chest_ten_num` int(11) NOT NULL COMMENT '今日神龙宝箱十连抽次数',
  `is_first_ingot_ten` tinyint(4) NOT NULL COMMENT '是否第一次神龙宝箱十连抽',
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
  `batch_bought` smallint(6) NOT NULL DEFAULT '0' COMMENT '玩家批量购买铜币次数',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家铜币兑换表';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_daily_quest`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_daily_quest` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '记录ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `quest_id` smallint(6) NOT NULL COMMENT '任务ID',
  `finish_count` smallint(6) NOT NULL DEFAULT '0' COMMENT '完成数量',
  `last_update_time` bigint(20) NOT NULL COMMENT '最近一次更新时间',
  `award_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '奖励状态; -1 未奖励；0可领取; 1已奖励',
  `class` smallint(6) NOT NULL COMMENT '每日任务类别',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=47244640293 DEFAULT CHARSET=utf8mb4 COMMENT='玩家每日任务';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_daily_sign_in_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_daily_sign_in_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `sign_in_time` bigint(20) DEFAULT NULL COMMENT '签到时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=47244640375 DEFAULT CHARSET=utf8mb4 COMMENT='玩家每日签到记录';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_daily_sign_in_state`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_daily_sign_in_state` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `update_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最新签到时间',
  `record` smallint(6) NOT NULL DEFAULT '0' COMMENT '签到记录',
  `signed_today` tinyint(4) DEFAULT NULL COMMENT '今天是否已签到',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家最近七日每日签到状态';
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
DROP TABLE IF EXISTS `player_extend_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_extend_level` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `coin_pass_time` bigint(20) NOT NULL COMMENT '铜钱关卡通关时间',
  `exp_pass_time` bigint(20) NOT NULL COMMENT '经验关卡通关时间',
  `ghost_pass_time` bigint(20) NOT NULL COMMENT '魂侍关卡通关时间',
  `pet_pass_time` bigint(20) NOT NULL COMMENT '灵宠关卡通关时间',
  `buddy_pass_time` bigint(20) NOT NULL COMMENT '伙伴关卡通关时间',
  `coin_daily_num` tinyint(4) NOT NULL COMMENT '经验关卡今日进入次数',
  `exp_daily_num` tinyint(4) NOT NULL COMMENT '铜钱关卡今日进入次数',
  `buddy_daily_num` tinyint(4) NOT NULL COMMENT '伙伴关卡今日进入次数',
  `pet_daily_num` tinyint(4) NOT NULL COMMENT '灵宠关卡今日进入次数',
  `ghost_daily_num` tinyint(4) NOT NULL COMMENT '魂侍关卡今日进入次数',
  `rand_buddy_role_id` tinyint(4) NOT NULL COMMENT '随机的伙伴角色ID',
  `buddy_pos` tinyint(4) NOT NULL COMMENT '随机的伙伴角色位置',
  `buddy_tactical` tinyint(4) NOT NULL COMMENT '伙伴关卡队伍战术',
  `role_pos` tinyint(4) NOT NULL COMMENT '主角站位',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家活动关卡状态';
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
  `ghost_power` int(11) NOT NULL COMMENT '魂力',
  `pos1` bigint(20) NOT NULL COMMENT '装备位置1的魂侍id',
  `pos2` bigint(20) NOT NULL COMMENT '装备位置2的魂侍id',
  `pos3` bigint(20) NOT NULL COMMENT '装备位置3的魂侍id',
  `pos4` bigint(20) NOT NULL COMMENT '装备位置4的魂侍id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家魂侍装备表';
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
DROP TABLE IF EXISTS `player_ghost_umbra`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_ghost_umbra` (
  `id` bigint(20) NOT NULL COMMENT '主键id',
  `pid` bigint(20) NOT NULL COMMENT '玩家id',
  `umbra_id` smallint(6) NOT NULL COMMENT '影界id',
  `num` smallint(6) NOT NULL COMMENT '今日剩余次数',
  `last_draw_at` bigint(20) NOT NULL COMMENT '上次开启时间',
  `refresh_num` smallint(6) NOT NULL COMMENT '今日刷新次数',
  `last_refresh_at` bigint(20) NOT NULL COMMENT '上次刷新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家魂侍副本表';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_global_arena_rank`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_global_arena_rank` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `rank` int(11) NOT NULL DEFAULT '0' COMMENT '昨天排名',
  `rank1` int(11) NOT NULL DEFAULT '0' COMMENT '1天前排名',
  `rank2` int(11) NOT NULL DEFAULT '0' COMMENT '2天前排名',
  `rank3` int(11) NOT NULL DEFAULT '0' COMMENT '3天前排名',
  `time` bigint(20) NOT NULL COMMENT '宝箱刷新时间',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家比武场最近排名记录';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_global_friend`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_global_friend` (
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
DROP TABLE IF EXISTS `player_global_friend_chat`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_global_friend_chat` (
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
DROP TABLE IF EXISTS `player_global_friend_state`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_global_friend_state` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `delete_day_count` int(11) NOT NULL DEFAULT '0' COMMENT '每日删除好友数量',
  `delete_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '删除好友时间',
  `exist_offline_chat` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0没有离线消息，1有离线消息',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家好友功能状态数据';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_hard_level`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_hard_level` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `lock` int(11) DEFAULT '0' COMMENT '难度关卡功能权值',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='难度关卡功能权值';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_hard_level_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_hard_level_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '记录ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `level_id` int(11) NOT NULL COMMENT '开启的关卡ID',
  `open_time` bigint(20) NOT NULL COMMENT '关卡开启时间',
  `score` int(11) NOT NULL DEFAULT '0' COMMENT '得分',
  `round` tinyint(4) NOT NULL DEFAULT '0' COMMENT '通关回合数',
  `daily_num` tinyint(4) NOT NULL COMMENT '当日已进入关卡的次数',
  `last_enter_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最后一次进入时间',
  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`)
) ENGINE=InnoDB AUTO_INCREMENT=47244640261 DEFAULT CHARSET=utf8 COMMENT='难度关卡记录';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_hard_level_state`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_hard_level_state` (
  `id` bigint(20) NOT NULL COMMENT 'ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `mission_level_id` int(11) NOT NULL COMMENT '难度关卡ID',
  `battle_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最后一次战败时间',
  `pos1` int(11) NOT NULL DEFAULT '0' COMMENT '位置1上的敌人生命',
  `pos2` int(11) NOT NULL DEFAULT '0' COMMENT '位置2上的敌人生命',
  `pos3` int(11) NOT NULL DEFAULT '0' COMMENT '位置3上的敌人生命',
  `pos4` int(11) NOT NULL DEFAULT '0' COMMENT '位置4上的敌人生命',
  `pos5` int(11) NOT NULL DEFAULT '0' COMMENT '位置5上的敌人生命',
  `pos6` int(11) NOT NULL DEFAULT '0' COMMENT '位置6上的敌人生命',
  `pos7` int(11) NOT NULL DEFAULT '0' COMMENT '位置7上的敌人生命',
  `pos8` int(11) NOT NULL DEFAULT '0' COMMENT '位置8上的敌人生命',
  `pos9` int(11) NOT NULL DEFAULT '0' COMMENT '位置9上的敌人生命',
  `pos10` int(11) NOT NULL DEFAULT '0' COMMENT '位置10上的敌人生命',
  `pos11` int(11) NOT NULL DEFAULT '0' COMMENT '位置11上的敌人生命',
  `pos12` int(11) NOT NULL DEFAULT '0' COMMENT '位置12上的敌人生命',
  `pos13` int(11) NOT NULL DEFAULT '0' COMMENT '位置13上的敌人生命',
  `pos14` int(11) NOT NULL DEFAULT '0' COMMENT '位置14上的敌人生命',
  `pos15` int(11) NOT NULL DEFAULT '0' COMMENT '位置15上的敌人生命',
  `round` int(11) NOT NULL DEFAULT '0' COMMENT '累积回合数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='难度关卡失败状态记录';
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
) ENGINE=InnoDB AUTO_INCREMENT=47244640407 DEFAULT CHARSET=utf8mb4 COMMENT='玩家爱心抽奖';
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
) ENGINE=InnoDB AUTO_INCREMENT=42949672964 DEFAULT CHARSET=utf8mb4 COMMENT='玩家爱心大转盘抽奖记录';
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
  `init_global_srv` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否在互动服已初始化. 0 - 没有',
  `first_login_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '玩家注册时间',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家信息表';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_is_operated`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_is_operated` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `operate_value` bigint(20) NOT NULL DEFAULT '0' COMMENT '操作值',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='记录玩家是否第一次操作';
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
) ENGINE=InnoDB AUTO_INCREMENT=47244640337 DEFAULT CHARSET=utf8mb4 COMMENT='玩家邮件表';
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
) ENGINE=InnoDB AUTO_INCREMENT=47244640346 DEFAULT CHARSET=utf8mb4 COMMENT='玩家邮件附件表';
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
) ENGINE=InnoDB AUTO_INCREMENT=47244640354 DEFAULT CHARSET=utf8 COMMENT='关卡记录';
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
) ENGINE=InnoDB AUTO_INCREMENT=47244640358 DEFAULT CHARSET=utf8 COMMENT='区域记录';
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
DROP TABLE IF EXISTS `player_npc_talk_record`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_npc_talk_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '记录ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `npc_id` int(11) NOT NULL COMMENT 'NPC ID',
  `town_id` smallint(6) NOT NULL COMMENT '相关城镇',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=47244640260 DEFAULT CHARSET=utf8mb4 COMMENT='玩家与NPC首次对话记录';
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
) ENGINE=InnoDB AUTO_INCREMENT=47244640376 DEFAULT CHARSET=utf8 COMMENT='玩家角色数据';
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
) ENGINE=InnoDB AUTO_INCREMENT=47244640304 DEFAULT CHARSET=utf8 COMMENT='玩家剑心数据';
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
  `box_state` tinyint(4) NOT NULL COMMENT '开箱子的状态(位操作)',
  `num` smallint(6) NOT NULL COMMENT '当前可拔剑次数',
  `update_time` bigint(20) NOT NULL COMMENT '更新时间',
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
  `at_pos_x` smallint(6) NOT NULL COMMENT '当前城镇的X轴位置',
  `at_pos_y` smallint(6) NOT NULL COMMENT '当前城镇的y轴位置',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家城镇数据';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_trader_refresh_state`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_trader_refresh_state` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `last_update_time` bigint(20) NOT NULL COMMENT '最近一次*手动*刷新时间',
  `auto_update_time` bigint(20) NOT NULL COMMENT '最近一次*自动*刷新时间',
  `trader_id` smallint(6) NOT NULL COMMENT '随机商人ID',
  `refresh_num` smallint(6) NOT NULL COMMENT '已刷新次数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=47244640263 DEFAULT CHARSET=utf8mb4 COMMENT='玩家随机商店手动刷新次数状态';
/*!40101 SET character_set_client = @saved_cs_client */;
DROP TABLE IF EXISTS `player_trader_store_state`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_trader_store_state` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `trader_id` smallint(6) NOT NULL COMMENT '随机商人ID',
  `grid_id` int(11) NOT NULL COMMENT '格子ID',
  `item_id` smallint(6) NOT NULL COMMENT '物品ID',
  `num` smallint(6) NOT NULL COMMENT '物品数量',
  `cost` bigint(20) NOT NULL COMMENT '价格',
  `stock` tinyint(4) NOT NULL DEFAULT '0' COMMENT '剩余可购买次数',
  `goods_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '货物类型0-物品 1-爱心 2-剑心 3-魂侍',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=47244640303 DEFAULT CHARSET=utf8mb4 COMMENT='玩家随机商店状态';
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
DROP TABLE IF EXISTS `player_vip`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_vip` (
  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
  `ingot` bigint(20) NOT NULL DEFAULT '0' COMMENT '累计充值元宝数',
  `level` smallint(6) NOT NULL DEFAULT '0' COMMENT 'VIP等级',
  `card_id` varchar(50) NOT NULL COMMENT 'VIP卡编号',
  PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家VIP卡信息';
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