<?php
$this->AddSQL(
"/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table event_vn_dragon_ball_exchange
# ------------------------------------------------------------

DROP TABLE IF EXISTS `event_vn_dragon_ball_exchange`;

CREATE TABLE `event_vn_dragon_ball_exchange` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '兑换物品标识',
  `quality` tinyint(4) NOT NULL DEFAULT '0' COMMENT '品质',
  `name` varchar(100) NOT NULL COMMENT '名称',
  `desc` varchar(200) NOT NULL COMMENT '描述',
  `sign` varchar(30) DEFAULT NULL COMMENT '资源标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='越南龙珠活动兑换列表';

/*!40000 ALTER TABLE `event_vn_dragon_ball_exchange` DISABLE KEYS */;

INSERT INTO `event_vn_dragon_ball_exchange` (`id`, `quality`, `name`, `desc`, `sign`)
VALUES
	(1,2,'大铜钱袋','需神、龙、降、世龙珠','DaTongQianDai'),
	(2,2,'茶叶蛋','需任意一颗龙珠','ChaYeDan'),
	(3,3,'iphone6','需集齐5颗龙珠','iphone6');

/*!40000 ALTER TABLE `event_vn_dragon_ball_exchange` ENABLE KEYS */;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
");