<?php

$this->AddSQL("
/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table item_type
# ------------------------------------------------------------

DROP TABLE IF EXISTS `item_type`;

CREATE TABLE `item_type` (
  `id` tinyint(4) NOT NULL AUTO_INCREMENT COMMENT '类型ID',
  `name` varchar(10) NOT NULL COMMENT '类型名称',
  `max_num_in_pos` smallint(6) NOT NULL DEFAULT '1' COMMENT '每个位置最多可堆叠的数量',
  `sign` varchar(50) DEFAULT '' COMMENT '类型标志',
  `order` int(11) DEFAULT '0' COMMENT '客户端排序权重',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COMMENT='物品类型';
/*!40000 ALTER TABLE `item_type` DISABLE KEYS */;

INSERT INTO `item_type` (`id`, `name`, `max_num_in_pos`, `sign`, `order`)
VALUES
	(2,'材料',9999,'Stuff',55),
	(3,'武器',1,'Weapon',10),
	(4,'战袍',1,'Clothes',15),
	(5,'靴子',1,'Shoes',20),
	(6,'饰品',1,'Accessories',25),
	(8,'战斗道具',999,'FightProp',65),
	(9,'礼包',999,'Chest',35),
	(10,'资源',999,'Resource',70),
	(11,'灵宠契约球',999,'Pet',80),
	(12,'道具',9999,'Props',60),
	(13,'魂侍碎片',999,'GhostFragment',75),
	(14,'剑心碎片',999,'SwordSoulFragment',85),
	(15,'消耗道具',9999,'CostProp',45),
	(16,'时装',1,'Fashion',50),
	(17,'武功秘籍',1,'Skill',30),
	(18,'喜好品',999,'Preference',90),
	(19,'任务道具',999,'Quest',95),
	(20,'龙珠',999,'DragonBall',40),
	(21,'秘宝',1,'ActReflect',5),
	(22,'活动道具',999,'Event',96);

/*!40000 ALTER TABLE `item_type` ENABLE KEYS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

UPDATE `item` SET `type_id`=22 WHERE `id`=850;
");
