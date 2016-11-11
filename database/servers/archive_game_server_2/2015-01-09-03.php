<?php
$this->AddSQL(
"DROP TABLE IF EXISTS `event_first_recharge_daily`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `event_first_recharge_daily` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `require_day` smallint(6) NOT NULL COMMENT '索引天数',
  `ingot` smallint(6) NOT NULL COMMENT '奖励元宝',
  `coins` int(11) NOT NULL COMMENT '奖励铜钱',
  `heart` smallint(6) NOT NULL COMMENT '奖励爱心',
  `item1_id` smallint(6) DEFAULT '0' COMMENT '物品1',
  `item1_num` smallint(6) DEFAULT '0' COMMENT '物品1数量',
  `item2_id` smallint(6) DEFAULT '0' COMMENT '物品2',
  `item2_num` smallint(6) DEFAULT '0' COMMENT '物品2数量',
  `item3_id` smallint(6) DEFAULT '0' COMMENT '物品3',
  `item3_num` smallint(6) DEFAULT '0' COMMENT '物品3数量',
  `item4_id` smallint(6) DEFAULT '0' COMMENT '物品4',
  `item4_num` smallint(6) DEFAULT '0' COMMENT '物品4数量',
  `item5_id` smallint(6) DEFAULT '0' COMMENT '物品5',
  `item5_num` smallint(6) DEFAULT '0' COMMENT '物品5数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='每日首充奖励';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `event_first_recharge_daily`
--

/*!40000 ALTER TABLE `event_first_recharge_daily` DISABLE KEYS */;
INSERT INTO `event_first_recharge_daily` VALUES (1,1,0,0,0,335,3,0,0,0,0,0,0,0,0),(2,2,0,0,0,425,3,0,0,0,0,0,0,0,0),(3,3,0,0,0,274,3,0,0,0,0,0,0,0,0);
/*!40000 ALTER TABLE `event_first_recharge_daily` ENABLE KEYS */;

--
-- Table structure for table `event_multiply_config`
--

DROP TABLE IF EXISTS `event_multiply_config`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `event_multiply_config` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `condition` int(11) DEFAULT NULL COMMENT '加成的事件id',
  `times` float(4,2) DEFAULT NULL COMMENT '加成的倍数',
  PRIMARY KEY (`id`),
  UNIQUE KEY `condition` (`condition`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='翻倍活动配置';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `event_multiply_config`
--

/*!40000 ALTER TABLE `event_multiply_config` DISABLE KEYS */;
INSERT INTO `event_multiply_config` VALUES (1,1,2.00),(2,2,2.00),(3,3,1.05),(4,4,1.05);
/*!40000 ALTER TABLE `event_multiply_config` ENABLE KEYS */;

--
-- Table structure for table `events_arena_rank_awards`
--

DROP TABLE IF EXISTS `events_arena_rank_awards`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `events_arena_rank_awards` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `require_arena_rank` smallint(6) NOT NULL COMMENT '需要排名',
  `ingot` smallint(6) NOT NULL COMMENT '奖励元宝',
  `coins` int(11) NOT NULL COMMENT '奖励铜钱',
  `item1_id` smallint(6) DEFAULT '0' COMMENT '物品1',
  `item1_num` smallint(6) DEFAULT '0' COMMENT '物品1数量',
  `item2_id` smallint(6) DEFAULT '0' COMMENT '物品2',
  `item2_num` smallint(6) DEFAULT '0' COMMENT '物品2数量',
  `item3_id` smallint(6) DEFAULT '0' COMMENT '物品3',
  `item3_num` smallint(6) DEFAULT '0' COMMENT '物品3数量',
  `item4_id` smallint(6) DEFAULT '0' COMMENT '物品4',
  `item4_num` smallint(6) DEFAULT '0' COMMENT '物品4数量',
  `item5_id` smallint(6) DEFAULT '0' COMMENT '物品5',
  `item5_num` smallint(6) DEFAULT '0' COMMENT '物品5数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COMMENT='比武场排名活动运营活动';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events_arena_rank_awards`
--

/*!40000 ALTER TABLE `events_arena_rank_awards` DISABLE KEYS */;
INSERT INTO `events_arena_rank_awards` VALUES (1,1,2000,0,0,0,0,0,0,0,0,0,0,0),(3,2,1000,0,0,0,0,0,0,0,0,0,0,0),(4,3,750,0,0,0,0,0,0,0,0,0,0,0),(11,10,500,0,0,0,0,0,0,0,0,0,0,0),(14,50,300,0,0,0,0,0,0,0,0,0,0,0),(15,100,200,0,0,0,0,0,0,0,0,0,0,0),(16,500,100,0,0,0,0,0,0,0,0,0,0,0);
/*!40000 ALTER TABLE `events_arena_rank_awards` ENABLE KEYS */;

--
-- Table structure for table `events_buy_partner`
--

DROP TABLE IF EXISTS `events_buy_partner`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `events_buy_partner` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `patner_id` smallint(6) NOT NULL COMMENT '伙伴ID',
  `buddy_level` smallint(6) NOT NULL DEFAULT '1' COMMENT '伙伴等级',
  `cost` bigint(20) NOT NULL COMMENT '价格',
  `skill_id1` smallint(6) NOT NULL COMMENT '招式名称1',
  `skill_id2` smallint(6) NOT NULL COMMENT '招式名称2',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='限时购买伙伴';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events_buy_partner`
--

/*!40000 ALTER TABLE `events_buy_partner` DISABLE KEYS */;
INSERT INTO `events_buy_partner` VALUES (1,7,40,3000,1300,1301);
/*!40000 ALTER TABLE `events_buy_partner` ENABLE KEYS */;

--
-- Table structure for table `events_dinner_awards`
--

DROP TABLE IF EXISTS `events_dinner_awards`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `events_dinner_awards` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `require_day` smallint(6) NOT NULL COMMENT '需要天数',
  `ingot` smallint(6) NOT NULL COMMENT '奖励元宝',
  `coins` int(11) NOT NULL COMMENT '奖励铜钱',
  `item1_id` smallint(6) DEFAULT '0' COMMENT '物品1',
  `item1_num` smallint(6) DEFAULT '0' COMMENT '物品1数量',
  `item2_id` smallint(6) DEFAULT '0' COMMENT '物品2',
  `item2_num` smallint(6) DEFAULT '0' COMMENT '物品2数量',
  `item3_id` smallint(6) DEFAULT '0' COMMENT '物品3',
  `item3_num` smallint(6) DEFAULT '0' COMMENT '物品3数量',
  `item4_id` smallint(6) DEFAULT '0' COMMENT '物品4',
  `item4_num` smallint(6) DEFAULT '0' COMMENT '物品4数量',
  `item5_id` smallint(6) DEFAULT '0' COMMENT '物品5',
  `item5_num` smallint(6) DEFAULT '0' COMMENT '物品5数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COMMENT='午餐活动运营活动';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events_dinner_awards`
--

/*!40000 ALTER TABLE `events_dinner_awards` DISABLE KEYS */;
INSERT INTO `events_dinner_awards` VALUES (1,1,0,0,303,2,0,0,0,0,0,0,0,0),(2,2,0,0,303,2,0,0,0,0,0,0,0,0),(3,3,0,0,303,2,0,0,0,0,0,0,0,0),(4,4,0,0,303,2,0,0,0,0,0,0,0,0),(5,5,0,0,303,2,0,0,0,0,0,0,0,0),(6,6,0,0,303,4,0,0,0,0,0,0,0,0),(7,7,0,0,303,4,0,0,0,0,0,0,0,0);
/*!40000 ALTER TABLE `events_dinner_awards` ENABLE KEYS */;

--
-- Table structure for table `events_fight_power`
--

DROP TABLE IF EXISTS `events_fight_power`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `events_fight_power` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `lock` smallint(6) NOT NULL COMMENT '档位',
  `fight` int(11) NOT NULL COMMENT '战力',
  `ingot` smallint(6) NOT NULL COMMENT '奖励元宝',
  `item1_id` smallint(6) DEFAULT '0' COMMENT '物品1',
  `item1_num` smallint(6) DEFAULT '0' COMMENT '物品1数量',
  `item2_id` smallint(6) DEFAULT '0' COMMENT '物品2',
  `item2_num` smallint(6) DEFAULT '0' COMMENT '物品2数量',
  `item3_id` smallint(6) DEFAULT '0' COMMENT '物品3',
  `item3_num` smallint(6) DEFAULT '0' COMMENT '物品3数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COMMENT='战力运营活动';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events_fight_power`
--

/*!40000 ALTER TABLE `events_fight_power` DISABLE KEYS */;
INSERT INTO `events_fight_power` VALUES (3,1,20000,100,0,0,0,0,0,0),(6,2,50000,200,0,0,0,0,0,0),(7,3,60000,300,0,0,0,0,0,0),(8,4,100000,400,0,0,0,0,0,0);
/*!40000 ALTER TABLE `events_fight_power` ENABLE KEYS */;

--
-- Table structure for table `events_group_buy`
--

DROP TABLE IF EXISTS `events_group_buy`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `events_group_buy` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT 'ID标识',
  `item_id` smallint(6) NOT NULL COMMENT '参与团购得物品id',
  `base_price` smallint(6) NOT NULL COMMENT '团购物品低价',
  `buy_times1` smallint(6) NOT NULL COMMENT '购买次数1',
  `buy_percent1` float(3,2) NOT NULL COMMENT '购买折扣1',
  `buy_times2` smallint(6) NOT NULL COMMENT '购买次数2',
  `buy_percent2` float(3,2) NOT NULL COMMENT '购买折扣2',
  `buy_times3` smallint(6) NOT NULL COMMENT '购买次数3',
  `buy_percent3` float(3,2) NOT NULL COMMENT '购买折扣3',
  `buy_times4` smallint(6) NOT NULL COMMENT '购买次数4',
  `buy_percent4` float(3,2) NOT NULL COMMENT '购买折扣4',
  `buy_times5` smallint(6) NOT NULL COMMENT '购买次数5',
  `buy_percent5` float(3,2) NOT NULL COMMENT '购买折扣5',
  `buy_times6` smallint(6) NOT NULL COMMENT '购买次数6',
  `buy_percent6` float(3,2) NOT NULL COMMENT '购买折扣6',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='团购内容';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events_group_buy`
--

/*!40000 ALTER TABLE `events_group_buy` DISABLE KEYS */;
INSERT INTO `events_group_buy` VALUES (1,342,2000,0,0.75,1,0.70,2,0.65,3,0.60,4,0.55,5,0.50);
/*!40000 ALTER TABLE `events_group_buy` ENABLE KEYS */;

--
-- Table structure for table `events_level_up`
--

DROP TABLE IF EXISTS `events_level_up`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `events_level_up` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `require_level` smallint(6) NOT NULL COMMENT '需要等级',
  `ingot` smallint(6) NOT NULL COMMENT '奖励元宝',
  `item1_id` smallint(6) DEFAULT '0' COMMENT '物品1',
  `item1_num` smallint(6) DEFAULT '0' COMMENT '物品1数量',
  `item2_id` smallint(6) DEFAULT '0' COMMENT '物品2',
  `item2_num` smallint(6) DEFAULT '0' COMMENT '物品2数量',
  `item3_id` smallint(6) DEFAULT '0' COMMENT '物品3',
  `item3_num` smallint(6) DEFAULT '0' COMMENT '物品3数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COMMENT='角色升级运营活动';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events_level_up`
--

/*!40000 ALTER TABLE `events_level_up` DISABLE KEYS */;
INSERT INTO `events_level_up` VALUES (3,15,20,0,0,0,0,0,0),(4,20,30,0,0,0,0,0,0),(5,25,50,0,0,0,0,0,0),(7,35,50,0,0,0,0,0,0),(8,40,100,0,0,0,0,0,0),(9,45,100,0,0,0,0,0,0),(10,50,100,0,0,0,0,0,0),(11,55,100,0,0,0,0,0,0),(13,30,50,0,0,0,0,0,0),(14,60,100,0,0,0,0,0,0);
/*!40000 ALTER TABLE `events_level_up` ENABLE KEYS */;

--
-- Table structure for table `events_month_card_awards`
--

DROP TABLE IF EXISTS `events_month_card_awards`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `events_month_card_awards` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `require_month_card` smallint(6) NOT NULL COMMENT '需要月卡',
  `ingot` smallint(6) NOT NULL COMMENT '奖励元宝',
  `coins` int(11) NOT NULL COMMENT '奖励铜钱',
  `item1_id` smallint(6) DEFAULT '0' COMMENT '物品1',
  `item1_num` smallint(6) DEFAULT '0' COMMENT '物品1数量',
  `item2_id` smallint(6) DEFAULT '0' COMMENT '物品2',
  `item2_num` smallint(6) DEFAULT '0' COMMENT '物品2数量',
  `item3_id` smallint(6) DEFAULT '0' COMMENT '物品3',
  `item3_num` smallint(6) DEFAULT '0' COMMENT '物品3数量',
  `item4_id` smallint(6) DEFAULT '0' COMMENT '物品4',
  `item4_num` smallint(6) DEFAULT '0' COMMENT '物品4数量',
  `item5_id` smallint(6) DEFAULT '0' COMMENT '物品5',
  `item5_num` smallint(6) DEFAULT '0' COMMENT '物品5数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='月卡活动运营活动';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events_month_card_awards`
--

/*!40000 ALTER TABLE `events_month_card_awards` DISABLE KEYS */;
INSERT INTO `events_month_card_awards` VALUES (1,1,100,0,0,0,0,0,0,0,0,0,0,0);
/*!40000 ALTER TABLE `events_month_card_awards` ENABLE KEYS */;

--
-- Table structure for table `events_physical_awards`
--

DROP TABLE IF EXISTS `events_physical_awards`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `events_physical_awards` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `require_physical` smallint(6) NOT NULL COMMENT '需要活跃度',
  `ingot` smallint(6) NOT NULL COMMENT '奖励元宝',
  `coins` int(11) NOT NULL COMMENT '奖励铜钱',
  `item1_id` smallint(6) DEFAULT '0' COMMENT '物品1',
  `item1_num` smallint(6) DEFAULT '0' COMMENT '物品1数量',
  `item2_id` smallint(6) DEFAULT '0' COMMENT '物品2',
  `item2_num` smallint(6) DEFAULT '0' COMMENT '物品2数量',
  `item3_id` smallint(6) DEFAULT '0' COMMENT '物品3',
  `item3_num` smallint(6) DEFAULT '0' COMMENT '物品3数量',
  `item4_id` smallint(6) DEFAULT '0' COMMENT '物品4',
  `item4_num` smallint(6) DEFAULT '0' COMMENT '物品4数量',
  `item5_id` smallint(6) DEFAULT '0' COMMENT '物品5',
  `item5_num` smallint(6) DEFAULT '0' COMMENT '物品5数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COMMENT='活跃度活动运营活动';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events_physical_awards`
--

/*!40000 ALTER TABLE `events_physical_awards` DISABLE KEYS */;
INSERT INTO `events_physical_awards` VALUES (1,100,0,1000,0,0,0,0,0,0,0,0,0,0),(2,200,5,5000,0,0,0,0,0,0,0,0,0,0),(3,300,5,5000,0,0,0,0,0,0,0,0,0,0),(4,400,5,5000,0,0,0,0,0,0,0,0,0,0),(5,500,10,10000,0,0,0,0,0,0,0,0,0,0),(6,700,10,10000,0,0,0,0,0,0,0,0,0,0),(7,900,15,15000,0,0,0,0,0,0,0,0,0,0),(8,1200,15,15000,0,0,0,0,0,0,0,0,0,0),(9,1500,20,20000,0,0,0,0,0,0,0,0,0,0);
/*!40000 ALTER TABLE `events_physical_awards` ENABLE KEYS */;

--
-- Table structure for table `events_qqvip_gift_awards`
--

DROP TABLE IF EXISTS `events_qqvip_gift_awards`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `events_qqvip_gift_awards` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `require_login_days` smallint(6) NOT NULL COMMENT '需要连续登录天数',
  `ingot` smallint(6) NOT NULL COMMENT '奖励元宝',
  `coins` int(11) NOT NULL COMMENT '奖励铜钱',
  `item1_id` smallint(6) DEFAULT '0' COMMENT '物品1',
  `item1_num` smallint(6) DEFAULT '0' COMMENT '物品1数量',
  `item2_id` smallint(6) DEFAULT '0' COMMENT '物品2',
  `item2_num` smallint(6) DEFAULT '0' COMMENT '物品2数量',
  `item3_id` smallint(6) DEFAULT '0' COMMENT '物品3',
  `item3_num` smallint(6) DEFAULT '0' COMMENT '物品3数量',
  `item4_id` smallint(6) DEFAULT '0' COMMENT '物品4',
  `item4_num` smallint(6) DEFAULT '0' COMMENT '物品4数量',
  `item5_id` smallint(6) DEFAULT '0' COMMENT '物品5',
  `item5_num` smallint(6) DEFAULT '0' COMMENT '物品5数量',
  `title` varchar(30) DEFAULT NULL COMMENT '奖励名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='QQ特权赠送物品运营活动';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events_qqvip_gift_awards`
--

/*!40000 ALTER TABLE `events_qqvip_gift_awards` DISABLE KEYS */;
INSERT INTO `events_qqvip_gift_awards` VALUES (1,1,0,0,232,2,306,1,0,0,0,0,0,0,'初级连续登录奖励'),(2,3,0,0,232,3,306,2,0,0,0,0,0,0,'中级连续登录奖励'),(3,7,0,0,232,5,306,3,0,0,0,0,0,0,'高级连续登录奖励');
/*!40000 ALTER TABLE `events_qqvip_gift_awards` ENABLE KEYS */;

--
-- Table structure for table `events_richman_club_awards`
--

DROP TABLE IF EXISTS `events_richman_club_awards`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `events_richman_club_awards` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `require_vip_level` smallint(6) NOT NULL COMMENT '所需的vip等级',
  `require_vip_count` smallint(6) NOT NULL COMMENT '所需的vip相应人数',
  `award_vip_level1` smallint(6) DEFAULT '0' COMMENT '能领奖的vip等级1',
  `award_vip_item1_id` smallint(6) DEFAULT '0' COMMENT '能领奖的vip的奖励物品1 ID',
  `award_vip_item1_num` smallint(6) DEFAULT '1' COMMENT '能领奖的vip的奖励物品1数量 默认为1',
  `award_vip_level2` smallint(6) DEFAULT '0' COMMENT '能领奖的vip等级2',
  `award_vip_item2_id` smallint(6) DEFAULT '0' COMMENT '能领奖的vip的奖励物品2 ID',
  `award_vip_item2_num` smallint(6) DEFAULT '1' COMMENT '能领奖的vip的奖励物品2数量 默认为1',
  `award_vip_level3` smallint(6) DEFAULT '0' COMMENT '能领奖的vip等级3',
  `award_vip_item3_id` smallint(6) DEFAULT '0' COMMENT '能领奖的vip的奖励物品3 ID',
  `award_vip_item3_num` smallint(6) DEFAULT '1' COMMENT '能领奖的vip的奖励物品3数量 默认为1',
  `award_vip_level4` smallint(6) DEFAULT '0' COMMENT '能领奖的vip等级4',
  `award_vip_item4_id` smallint(6) DEFAULT '0' COMMENT '能领奖的vip的奖励物品4 ID',
  `award_vip_item4_num` smallint(6) DEFAULT '1' COMMENT '能领奖的vip的奖励物品4数量 默认为1',
  `award_vip_level5` smallint(6) DEFAULT '0' COMMENT '能领奖的vip等级5',
  `award_vip_item5_id` smallint(6) DEFAULT '0' COMMENT '能领奖的vip的奖励物品5 ID',
  `award_vip_item5_num` smallint(6) DEFAULT '1' COMMENT '能领奖的vip的奖励物品5数量 默认为1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='土豪俱乐部运营活动';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events_richman_club_awards`
--

/*!40000 ALTER TABLE `events_richman_club_awards` DISABLE KEYS */;
INSERT INTO `events_richman_club_awards` VALUES (1,6,10,1,404,1,0,0,0,0,0,0,0,0,0,0,0,0),(2,7,5,1,405,1,2,406,1,0,0,0,0,0,0,0,0,0),(3,8,3,1,407,1,2,408,1,3,409,1,0,0,0,0,0,0);
/*!40000 ALTER TABLE `events_richman_club_awards` ENABLE KEYS */;

--
-- Table structure for table `events_seven_day_awards`
--

DROP TABLE IF EXISTS `events_seven_day_awards`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `events_seven_day_awards` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `require_day` smallint(6) NOT NULL COMMENT '需要天数',
  `ingot` smallint(6) NOT NULL COMMENT '奖励元宝',
  `coins` int(11) NOT NULL COMMENT '奖励铜钱',
  `item1_id` smallint(6) DEFAULT '0' COMMENT '物品1',
  `item1_num` smallint(6) DEFAULT '0' COMMENT '物品1数量',
  `item2_id` smallint(6) DEFAULT '0' COMMENT '物品2',
  `item2_num` smallint(6) DEFAULT '0' COMMENT '物品2数量',
  `item3_id` smallint(6) DEFAULT '0' COMMENT '物品3',
  `item3_num` smallint(6) DEFAULT '0' COMMENT '物品3数量',
  `item4_id` smallint(6) DEFAULT '0' COMMENT '物品4',
  `item4_num` smallint(6) DEFAULT '0' COMMENT '物品4数量',
  `item5_id` smallint(6) DEFAULT '0' COMMENT '物品5',
  `item5_num` smallint(6) DEFAULT '0' COMMENT '物品5数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='新手七天乐运营活动';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events_seven_day_awards`
--

/*!40000 ALTER TABLE `events_seven_day_awards` DISABLE KEYS */;
INSERT INTO `events_seven_day_awards` VALUES (1,1,100,0,0,0,0,0,0,0,0,0,0,0),(2,2,0,0,305,2,0,0,0,0,0,0,0,0),(3,3,0,0,341,1,0,0,0,0,0,0,0,0),(5,4,128,0,263,50,0,0,0,0,0,0,0,0),(6,5,0,0,426,1,0,0,0,0,0,0,0,0),(7,6,168,0,355,20,0,0,0,0,0,0,0,0),(8,7,0,0,432,20,0,0,0,0,0,0,0,0);
/*!40000 ALTER TABLE `events_seven_day_awards` ENABLE KEYS */;

--
-- Table structure for table `events_share_awards`
--

DROP TABLE IF EXISTS `events_share_awards`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `events_share_awards` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `require_times` smallint(6) NOT NULL COMMENT '需要分享次数',
  `ingot` smallint(6) NOT NULL COMMENT '奖励元宝',
  `coins` int(11) NOT NULL COMMENT '奖励铜钱',
  `item1_id` smallint(6) DEFAULT '0' COMMENT '物品1',
  `item1_num` smallint(6) DEFAULT '0' COMMENT '物品1数量',
  `item2_id` smallint(6) DEFAULT '0' COMMENT '物品2',
  `item2_num` smallint(6) DEFAULT '0' COMMENT '物品2数量',
  `item3_id` smallint(6) DEFAULT '0' COMMENT '物品3',
  `item3_num` smallint(6) DEFAULT '0' COMMENT '物品3数量',
  `item4_id` smallint(6) DEFAULT '0' COMMENT '物品4',
  `item4_num` smallint(6) DEFAULT '0' COMMENT '物品4数量',
  `item5_id` smallint(6) DEFAULT '0' COMMENT '物品5',
  `item5_num` smallint(6) DEFAULT '0' COMMENT '物品5数量',
  `heart` smallint(6) DEFAULT '0' COMMENT '爱心',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='分享送好礼运营活动';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events_share_awards`
--

/*!40000 ALTER TABLE `events_share_awards` DISABLE KEYS */;
INSERT INTO `events_share_awards` VALUES (1,1,50,0,0,0,0,0,0,0,0,0,0,0,10),(2,5,50,0,0,0,0,0,0,0,0,0,0,0,10),(3,3,50,0,0,0,0,0,0,0,0,0,0,0,10),(4,7,50,0,0,0,0,0,0,0,0,0,0,0,10);
/*!40000 ALTER TABLE `events_share_awards` ENABLE KEYS */;

--
-- Table structure for table `events_supper_awards`
--

DROP TABLE IF EXISTS `events_supper_awards`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `events_supper_awards` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `require_day` smallint(6) NOT NULL COMMENT '需要天数',
  `ingot` smallint(6) NOT NULL COMMENT '奖励元宝',
  `coins` int(11) NOT NULL COMMENT '奖励铜钱',
  `item1_id` smallint(6) DEFAULT '0' COMMENT '物品1',
  `item1_num` smallint(6) DEFAULT '0' COMMENT '物品1数量',
  `item2_id` smallint(6) DEFAULT '0' COMMENT '物品2',
  `item2_num` smallint(6) DEFAULT '0' COMMENT '物品2数量',
  `item3_id` smallint(6) DEFAULT '0' COMMENT '物品3',
  `item3_num` smallint(6) DEFAULT '0' COMMENT '物品3数量',
  `item4_id` smallint(6) DEFAULT '0' COMMENT '物品4',
  `item4_num` smallint(6) DEFAULT '0' COMMENT '物品4数量',
  `item5_id` smallint(6) DEFAULT '0' COMMENT '物品5',
  `item5_num` smallint(6) DEFAULT '0' COMMENT '物品5数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COMMENT='晚餐活动运营活动';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events_supper_awards`
--

/*!40000 ALTER TABLE `events_supper_awards` DISABLE KEYS */;
INSERT INTO `events_supper_awards` VALUES (1,1,0,0,303,2,0,0,0,0,0,0,0,0),(2,2,0,0,303,2,0,0,0,0,0,0,0,0),(3,3,0,0,303,2,0,0,0,0,0,0,0,0),(4,4,0,0,303,2,0,0,0,0,0,0,0,0),(5,5,0,0,303,2,0,0,0,0,0,0,0,0),(6,6,0,0,303,4,0,0,0,0,0,0,0,0),(7,7,0,0,303,4,0,0,0,0,0,0,0,0);
/*!40000 ALTER TABLE `events_supper_awards` ENABLE KEYS */;

--
-- Table structure for table `events_ten_draw_awards`
--

DROP TABLE IF EXISTS `events_ten_draw_awards`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `events_ten_draw_awards` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `require_times` smallint(6) NOT NULL COMMENT '需要次数',
  `ingot` smallint(6) NOT NULL COMMENT '奖励元宝',
  `coins` int(11) NOT NULL COMMENT '奖励铜钱',
  `heart` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励爱心',
  `item1_id` smallint(6) DEFAULT '0' COMMENT '物品1',
  `item1_num` smallint(6) DEFAULT '0' COMMENT '物品1数量',
  `item2_id` smallint(6) DEFAULT '0' COMMENT '物品2',
  `item2_num` smallint(6) DEFAULT '0' COMMENT '物品2数量',
  `item3_id` smallint(6) DEFAULT '0' COMMENT '物品3',
  `item3_num` smallint(6) DEFAULT '0' COMMENT '物品3数量',
  `item4_id` smallint(6) DEFAULT '0' COMMENT '物品4',
  `item4_num` smallint(6) DEFAULT '0' COMMENT '物品4数量',
  `item5_id` smallint(6) DEFAULT '0' COMMENT '物品5',
  `item5_num` smallint(6) DEFAULT '0' COMMENT '物品5数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='十连抽运营活动';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events_ten_draw_awards`
--

/*!40000 ALTER TABLE `events_ten_draw_awards` DISABLE KEYS */;
INSERT INTO `events_ten_draw_awards` VALUES (1,1,0,46888,0,425,1,0,0,0,0,0,0,0,0),(2,3,0,88888,0,425,2,0,0,0,0,0,0,0,0),(3,5,0,188888,0,425,3,0,0,0,0,0,0,0,0),(4,7,0,268888,0,425,4,0,0,0,0,0,0,0,0),(5,10,0,368888,0,425,10,0,0,0,0,0,0,0,0);
/*!40000 ALTER TABLE `events_ten_draw_awards` ENABLE KEYS */;

--
-- Table structure for table `events_total_consume`
--

DROP TABLE IF EXISTS `events_total_consume`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `events_total_consume` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `require_cost` smallint(6) NOT NULL COMMENT '需要消耗',
  `ingot` smallint(6) NOT NULL COMMENT '奖励元宝',
  `coins` int(11) NOT NULL COMMENT '奖励铜钱',
  `heart` smallint(6) NOT NULL COMMENT '奖励爱心',
  `item1_id` smallint(6) DEFAULT '0' COMMENT '物品1',
  `item1_num` smallint(6) DEFAULT '0' COMMENT '物品1数量',
  `item2_id` smallint(6) DEFAULT '0' COMMENT '物品2',
  `item2_num` smallint(6) DEFAULT '0' COMMENT '物品2数量',
  `item3_id` smallint(6) DEFAULT '0' COMMENT '物品3',
  `item3_num` smallint(6) DEFAULT '0' COMMENT '物品3数量',
  `item4_id` smallint(6) DEFAULT '0' COMMENT '物品4',
  `item4_num` smallint(6) DEFAULT '0' COMMENT '物品4数量',
  `item5_id` smallint(6) DEFAULT '0' COMMENT '物品5',
  `item5_num` smallint(6) DEFAULT '0' COMMENT '物品5数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='元宝累计消耗奖励运营活动';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events_total_consume`
--

/*!40000 ALTER TABLE `events_total_consume` DISABLE KEYS */;
INSERT INTO `events_total_consume` VALUES (1,200,0,0,0,417,1,0,0,0,0,0,0,0,0);
/*!40000 ALTER TABLE `events_total_consume` ENABLE KEYS */;

--
-- Table structure for table `events_total_login`
--

DROP TABLE IF EXISTS `events_total_login`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `events_total_login` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `require_login_days` smallint(6) NOT NULL COMMENT '需要连续登录天数',
  `ingot` smallint(6) NOT NULL COMMENT '奖励元宝',
  `coins` int(11) NOT NULL COMMENT '奖励铜钱',
  `heart` smallint(6) NOT NULL COMMENT '奖励爱心',
  `item1_id` smallint(6) DEFAULT '0' COMMENT '物品1',
  `item1_num` smallint(6) DEFAULT '0' COMMENT '物品1数量',
  `item2_id` smallint(6) DEFAULT '0' COMMENT '物品2',
  `item2_num` smallint(6) DEFAULT '0' COMMENT '物品2数量',
  `item3_id` smallint(6) DEFAULT '0' COMMENT '物品3',
  `item3_num` smallint(6) DEFAULT '0' COMMENT '物品3数量',
  `item4_id` smallint(6) DEFAULT '0' COMMENT '物品4',
  `item4_num` smallint(6) DEFAULT '0' COMMENT '物品4数量',
  `item5_id` smallint(6) DEFAULT '0' COMMENT '物品5',
  `item5_num` smallint(6) DEFAULT '0' COMMENT '物品5数量',
  `title` varchar(30) DEFAULT NULL COMMENT '奖励名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COMMENT='累计登陆奖励运营活动';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events_total_login`
--

/*!40000 ALTER TABLE `events_total_login` DISABLE KEYS */;
INSERT INTO `events_total_login` VALUES (4,1,20,0,0,306,1,263,5,0,0,0,0,0,0,'初级连续登录奖励'),(5,4,35,0,0,306,3,263,7,0,0,0,0,0,0,'中级连续登录奖励'),(6,7,50,0,0,306,5,263,10,0,0,0,0,0,0,'高级连续登录奖励');
/*!40000 ALTER TABLE `events_total_login` ENABLE KEYS */;

--
-- Table structure for table `events_total_recharge_awards`
--

DROP TABLE IF EXISTS `events_total_recharge_awards`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `events_total_recharge_awards` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `require_total` smallint(6) NOT NULL COMMENT '需要额度',
  `ingot` smallint(6) NOT NULL COMMENT '奖励元宝',
  `coins` int(11) NOT NULL COMMENT '奖励铜钱',
  `heart` smallint(6) NOT NULL COMMENT '奖励爱心',
  `item1_id` smallint(6) DEFAULT '0' COMMENT '物品1',
  `item1_num` smallint(6) DEFAULT '0' COMMENT '物品1数量',
  `item2_id` smallint(6) DEFAULT '0' COMMENT '物品2',
  `item2_num` smallint(6) DEFAULT '0' COMMENT '物品2数量',
  `item3_id` smallint(6) DEFAULT '0' COMMENT '物品3',
  `item3_num` smallint(6) DEFAULT '0' COMMENT '物品3数量',
  `item4_id` smallint(6) DEFAULT '0' COMMENT '物品4',
  `item4_num` smallint(6) DEFAULT '0' COMMENT '物品4数量',
  `item5_id` smallint(6) DEFAULT '0' COMMENT '物品5',
  `item5_num` smallint(6) DEFAULT '0' COMMENT '物品5数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='累计充值运营活动';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events_total_recharge_awards`
--

/*!40000 ALTER TABLE `events_total_recharge_awards` DISABLE KEYS */;
INSERT INTO `events_total_recharge_awards` VALUES (1,500,0,0,0,432,10,0,0,0,0,0,0,0,0),(2,1000,0,0,0,432,10,0,0,0,0,0,0,0,0),(3,2000,0,0,0,432,10,0,0,0,0,0,0,0,0),(4,5000,0,0,0,432,10,0,0,0,0,0,0,0,0),(5,8888,0,0,0,432,10,0,0,0,0,0,0,0,0);
/*!40000 ALTER TABLE `events_total_recharge_awards` ENABLE KEYS */;

--
-- Table structure for table `events_vip_club_awards`
--

DROP TABLE IF EXISTS `events_vip_club_awards`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `events_vip_club_awards` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `require_vip_count` smallint(6) NOT NULL COMMENT '需要VIP人数',
  `ingot` smallint(6) NOT NULL COMMENT '奖励元宝',
  `coins` int(11) NOT NULL COMMENT '奖励铜钱',
  `item1_id` smallint(6) DEFAULT '0' COMMENT '物品1',
  `item1_num` smallint(6) DEFAULT '0' COMMENT '物品1数量',
  `item2_id` smallint(6) DEFAULT '0' COMMENT '物品2',
  `item2_num` smallint(6) DEFAULT '0' COMMENT '物品2数量',
  `item3_id` smallint(6) DEFAULT '0' COMMENT '物品3',
  `item3_num` smallint(6) DEFAULT '0' COMMENT '物品3数量',
  `item4_id` smallint(6) DEFAULT '0' COMMENT '物品4',
  `item4_num` smallint(6) DEFAULT '0' COMMENT '物品4数量',
  `item5_id` smallint(6) DEFAULT '0' COMMENT '物品5',
  `item5_num` smallint(6) DEFAULT '0' COMMENT '物品5数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='VIP俱乐部运营活动';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events_vip_club_awards`
--

/*!40000 ALTER TABLE `events_vip_club_awards` DISABLE KEYS */;
INSERT INTO `events_vip_club_awards` VALUES (1,20,0,0,263,5,231,5,0,0,0,0,0,0),(2,50,0,0,263,10,231,10,0,0,0,0,0,0),(3,100,0,0,263,20,231,20,0,0,0,0,0,0);
/*!40000 ALTER TABLE `events_vip_club_awards` ENABLE KEYS */;

DROP TABLE IF EXISTS `quest_activity_center`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `quest_activity_center` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `relative` smallint(6) NOT NULL COMMENT '关联的活动',
  `weight` int(8) NOT NULL DEFAULT '0' COMMENT '活动权值',
  `name` varchar(60) NOT NULL COMMENT '活动名称(列表左侧)',
  `title` varchar(100) NOT NULL COMMENT '活动标题(列表右侧)',
  `content` text NOT NULL COMMENT '活动描述',
  `start` bigint(20) DEFAULT '0' COMMENT '活动开始时间戳',
  `end` bigint(20) DEFAULT '0' COMMENT '活动结束时间戳',
  `is_go` tinyint(4) DEFAULT NULL COMMENT '是否前往',
  `tag` tinyint(4) DEFAULT NULL COMMENT '活动标签(1:最新,2:限时,3:推荐)',
  `is_mail` tinyint(4) DEFAULT '0' COMMENT '活动结束是否补发奖励',
  `condition_template` varchar(60) DEFAULT NULL COMMENT '领奖条件模版,{val}代表临界值',
  `dispose` bigint(20) DEFAULT NULL COMMENT '活动过期时间戳',
  `sign` varchar(40) DEFAULT NULL COMMENT '活动标识',
  `mail_title` varchar(60) DEFAULT NULL COMMENT '补发奖励邮件标题',
  `mail_content` text COMMENT '补发奖励邮件内容,{val}对应权值',
  `is_relative` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否为相对时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COMMENT='任务活动中心';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `quest_activity_center`
--

/*!40000 ALTER TABLE `quest_activity_center` DISABLE KEYS */;
INSERT INTO `quest_activity_center` VALUES (1,3,20,'累计登入','累计登入','活动期间，玩家累积登入7天，每日可领取各种极品奖励 \r\n活动奖励：\r\n天数           奖励内容\r\n第一天       200元宝\r\n第二天       珠宝盒（约50W铜钱）\r\n第三天       时装及紫色结晶（50个）\r\n第四天       金色魂侍“洛神”\r\n第五天       极品灵宠“雷兽”\r\n第六天       金色剑心“赤霄”（邮件发放）\r\n第七天       潜龙剑心10个（邮件发放）',1414512000,1414598400,1,1,0,'',1414598400,'EVENT_LOGIN_AWARD','','',0),(2,1,4,'等级活动','等级活动','活动期间，玩家每提升5级即可领取对应等级奖励 。\r\n活动奖励：\r\n15级       20元宝\r\n20级       30元宝\r\n25级       50元宝\r\n30级       50元宝\r\n35级       50元宝\r\n40级       100元宝\r\n45级       100元宝\r\n50级       100元宝\r\n55级       100元宝\r\n60级       100元宝',0,2591999,0,2,0,'等级{val}',2591999,'EVENT_LEVEL_AWARD','','',1),(3,2,5,'战力活动','战力活动','活动期间，玩家达到指定战力可领取对应战力奖励 。\r\n活动奖励：\r\n战力20000       100元宝\r\n战力50000       200元宝\r\n战力75000       300元宝\r\n战力100000     400元宝\r\n战力60000       588元宝',0,2591999,0,2,0,'战力{val}',2591999,'EVENT_STRONG_AWARD','','',1),(5,4,11,'首充元宝双倍返利','首充元宝双倍返利','活动期间，首次单笔充值金额满足每个档位对应金额（60元宝除外），即赠送等额元宝。\r\n 活动奖励：\r\n首充300元宝             赠送300元宝\r\n首充980元宝             赠送980元宝\r\n首充1980元宝           赠送1980元宝\r\n首充3280元宝           赠送3280元宝\r\n首充6480元宝           赠送6480元宝',0,0,1,3,0,'',0,'EVENT_RECHARGE_AWARD','','',0),(6,5,9,'最强王者','限时比武大会，谁是最强王者！','新服开启3天内，玩家进行比武场排名竞技，活动结束阶段根据玩家最终排名所在区间，给予对应活动奖励！届时玩家可在活动结束阶段（3天）前往该界面领取奖励。\r\n第1名                2000元宝\r\n第2名                1000元宝\r\n第3名                750元宝 \r\n第4～10名         500元宝\r\n第11～50名       300元宝\r\n第51～100名     200元宝\r\n第101～500名   100元宝',0,259200,1,2,0,'第{val}名',518399,'EVENT_ARENA_RANK_AWARDS','','',1),(7,6,100,'通关奖励翻倍','通关奖励翻倍！','玩家通过（普通关卡、精英关卡、BOSS关卡）中所获铜钱&经验奖励翻倍 \r\n活动时间：2014年11月10日-2014年11月23日',1414512000,1414598400,2,0,0,'',1414598400,'EVENT_MULTIPY_CONFIG','','',0),(8,7,1,'消耗体力换奖励','消耗体力换奖励','玩家每日消耗体力数累计达到活动奖励条件时，即可领取丰厚奖励！',0,1209599,0,4,0,'消耗体力{val}',1209599,'EVENT_PHYSICAL_AWARDS','','',1),(9,8,2,'月卡返利','月卡返利','单笔充值购买300元宝及以上，可自动激活月卡返利活动。\r\n激活后30天内，每天可领取100元宝奖励。',0,0,0,3,0,'',0,'EVENT_MONTH_CARD_AWARDS','','',0),(10,9,0,'豪华午餐','豪华午餐','每日 12：00 到 14：00 领取豪华午餐！吃饱喝足才有力气打怪哦！',0,0,0,4,0,'',0,'EVENT_DINNER_AWARDS','','',0),(11,10,0,'豪华晚餐','豪华晚餐','每日 18：00 到 20：00可领取豪华晚餐！为大侠补充体力～继续闯荡您的仙侠世界！',0,0,0,4,0,'',0,'EVENT_SUPPER_AWARDS','','',0),(12,11,15,'QQ特权奖励','QQ特权奖励','<font color=#FFD167 >特权一、通关奖励加成！</font>\r\nQQ会员玩家通过普通、精英及BOSS关卡时所获得经验和铜钱额外加成10%\r\n<font color=#FFD167 >特权二、每日登录奖励！</font>\r\nQQ会员玩家每日登录游戏，即可在该活动界面中领取当日福利奖励！连续登录更有额外奖励领取',0,0,0,0,0,'',0,'EVENT_QQVIP_GIFT_AWARDS','','',0),(13,12,17,'QQ特权奖励加成','QQ特权奖励加成','',0,0,0,0,0,'',0,'EVENT_QQVIP_ADDITION','','',0),(14,13,7,'仙尊俱乐部','仙尊俱乐部','当本服仙尊人数满足条件时，全体仙尊玩家均能享丰厚奖励！\r\n充值任意金额即能成为仙尊。',0,604799,0,3,0,'',604799,'EVENT_VIP_CLUB','','',1),(15,14,8,'土豪俱乐部','土豪俱乐部','当本服特定仙尊人数达标时，满足条件的仙尊玩家均能享受丰厚奖励！',0,604799,0,3,0,'',604799,'EVENT_RICHMAN_CLUB','','',1),(16,15,12,'公测元宝返还','公测元宝返还','内测参加充值返利活动以及冲级活动的大侠可以前往仙侠道手游官方网站领取活动奖励哦！',0,604799,2,3,0,'',604799,'EVENT_LEVEL_RECHARGE','','',1),(17,16,6,'首充奖励','首充奖励','玩家首次完成任意金额充值即可获得豪华首充奖励！\r\n活动奖励：\r\n紫色魂侍礼包（绮梦花妖）\r\n橙色结晶X2\r\n\r\n珠宝盒X1\r\n影界果实X5\r\n\r\n（奖励将以礼包形式通过邮件发放，请注意查收）',0,0,1,3,0,'',0,'EVENT_FIRST_RECHARGE','','',0),(19,17,3,'七日新手礼','七日新手礼','活动期间，玩家每日登录可领取当日对应的极品奖励。\r\n 活动奖励：\r\n第一天       100元宝\r\n第二天       珠宝盒X2（约25W铜钱）\r\n第三天       限时时装\r\n第四天       128元宝及影界果实（50个）\r\n第五天       紫色灵宠礼包（三阶魔笔）\r\n第六天       168元宝及灵魄（20个）\r\n第七天       喜好品（秘籍残页）x20\r\n（如当日未登录游戏即无法领取本日所对应的奖励）',0,0,1,4,0,'',0,'EVENT_SEVEN_DAY_AWARDS','','',0),(20,18,10,'江湖侠客乐分享','江湖侠客乐分享','大侠在仙侠道世界中成功分享的次数累计达到指定次数，可获得相应的奖励。',0,0,0,3,0,'',0,'EVENT_SHARE_AWARDS','','',0),(21,19,16,'团购活动','时装','活动期间，认购数量越多价格越优惠！若物品最终价格低于您认购时的花费，将退还差价。',1415548800,1415548801,0,2,0,'',1415548801,'EVENT_GROUP_BUY','','',0),(23,20,0,'累计充值狂欢','累计充值狂欢','元旦期间，玩家累计充值元宝（不计算赠送元宝）达到对应要求，即可获得<FONT COLOR=\"#FFFF00\">极品金色剑心碎片</FONT>',0,0,0,2,0,'',0,'EVENT_TOTAL_RECHARGE','','',0),(24,22,0,'邀请胧月','邀请胧月','初始等级[level]级，东瀛女忍者，性情阴晴不定，花费[ingot]即可邀请胧月加入队伍(赠送专属紫色武器及魂侍将通过邮件发放）',0,604800,0,3,0,'',604800,'EVENT_BUY_PARTNER','','',1),(25,23,0,'冬季十连抽大返利','冬季十连抽大返利','活动期间，玩家累计十连抽宝箱（神魂宝箱及神兽宝箱）达到对应次数，即可领取即可领取胧月专属金色魂侍！',0,0,0,2,0,'',0,'EVENT_TEN_DRAW','','',0),(26,21,0,'冬季消费大返利','冬季消费大返利','活动期间，玩家通过任意方式每消耗200元宝即可领取一个随机百宝箱，开启宝箱有几率获得极品魂侍武灵。',0,0,0,2,0,'',0,'EVENT_TOTAL_CONSUME','','',0),(28,24,0,'冬季每日充值礼','冬季每日充值礼','活动期间，玩家每日首次完成任意金额充值，即可领取对“限定魂侍碎片”奖励。\r\n（若充值完成但当日未领取奖励，之后无法领取该日奖励）',1419264000,1419523199,0,2,0,'',1419523199,'EVENT_FIRST_RECHARGE_DAILY','','',0),(29,25,0,'圣诞连续登录奖励','圣诞连续登录奖励','圣诞期间，根据玩家连续登录天数，每日可领取对应不同连续登录天数的连续登录奖励。\r\n连续登录4天及以上为中级登录奖励\r\n连续登录7天及以上为高级登录奖励\r\n（若玩家一日未登录游戏，连续登录天数将会重置）',1419264000,1420473599,0,2,0,'',1420473599,'EVENT_TOTAL_LOGIN','','',0);
/*!40000 ALTER TABLE `quest_activity_center` ENABLE KEYS */;
");
?>