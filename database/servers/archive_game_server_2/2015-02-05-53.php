<?php
$this->AddSQL("
# ************************************************************
# Sequel Pro SQL dump
# Version 4096
#
# http://www.sequelpro.com/
# http://code.google.com/p/sequel-pro/
#
# Host: 42.120.22.64 (MySQL 5.5.18.1-log)
# Database: xxd_dev20140919
# Generation Time: 2015-02-09 05:31:51 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table hard_level
# ------------------------------------------------------------

DROP TABLE IF EXISTS `hard_level`;

CREATE TABLE `hard_level` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `mission_level_lock` int(11) NOT NULL COMMENT '区域关卡功能权值',
  `desc` varchar(100) NOT NULL COMMENT '关卡描述',
  `town_id` smallint(6) NOT NULL COMMENT '城镇ID',
  `hard_level_lock` int(11) NOT NULL DEFAULT '0' COMMENT '难度关卡权值',
  `award_hard_level_lock` int(11) NOT NULL COMMENT '难度关卡奖励权值',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COMMENT='难度关卡';

/*!40000 ALTER TABLE `hard_level` DISABLE KEYS */;

INSERT INTO `hard_level` (`id`, `mission_level_lock`, `desc`, `town_id`, `hard_level_lock`, `award_hard_level_lock`)
VALUES
	(1,100110,'竹林里的竹子都成精了。',1,0,100100),
	(2,110130,'拥有摄人魂魄之能的灯笼怪。',1,100100,100110),
	(3,120140,'原来剧毒臭泥是吸收了各种腐化的气息后从地底逃窜到人间的。',1,100110,100120),
	(4,130150,'虽然炎龙尚未被阴影吞噬，但是在腐烂的地底，炎龙早已不是曾经的炎龙。',1,100120,100130),
	(5,140150,'上古的剑灵恐怕做梦也不会想到，自己的肉身会被魔物所占据。',1,100130,100140),
	(6,150170,'想不到一个小小的牢头也是大有来头的。',2,100140,100150),
	(7,160170,'野猪王充满了戾气与怨恨，被阴影侵蚀堕入深渊。',2,100150,100160),
	(8,170170,'深渊中的火麒麟被控制后脾气变的更加暴躁。',2,100160,100170),
	(9,180170,'为了除魔而进入深渊的地藏王也被阴影所侵蚀了。',2,100170,100180),
	(10,190170,'四大魔头之一奸奇的罪恶之源',2,100180,100190),
	(11,200170,'想不到古代武圣的肉身也被魔物侵占了',4,100190,100200),
	(12,210170,'聚集了太多怨气而化为妖魔的老树',4,100200,100210),
	(13,220170,'本是昆墟的守护，却在昆墟沦陷时一同堕入了深渊',4,100210,100220),
	(14,230170,'经过千年修炼的一条电鳗，可是不小心被阴影腐蚀了',4,100220,100230),
	(15,240170,'本是人间暴君，死后依旧本性不改',4,100230,100240),
	(16,250170,'三位上古仙灵中的飞羽肉身也已经沦陷了',5,100240,100250),
	(17,260170,'四大魔头之一色孽的罪恶之源',5,100250,100260),
	(18,270170,'操控风的大魔头，不知什么原因甘于屈居在徐福之下',5,100260,100270),
	(19,280170,'四大魔头之一纳垢的罪恶之源',5,100270,100280),
	(20,290170,'曾是著名方士，率众出海采仙药，一去不返',5,100280,100290),
	(21,300170,'四大魔头之一恐虐的罪恶之源',6,100290,100300),
	(22,310170,'藏匿在狂风巨浪的海岸边可怕的海妖',6,100300,100310),
	(23,320170,'带着不甘的第六天魔王织田信长',6,100310,100320),
	(24,330170,'天皇丑恶内心的罪恶之源',6,100320,100330),
	(25,340170,'死而不僵的骷髅将军',6,100330,100340),
	(26,350170,'一把马战用的关刀在虐杀手里平地上也能如臂使指',7,100340,100350),
	(27,360170,'龙虎门掌门，拳掌功夫相当了得',7,100350,100360),
	(28,370170,'先天的剑阵也堕入了魔道',7,100360,100370),
	(29,380170,'扫地的无名老僧武功竟已入化境',7,100370,100380),
	(30,390170,'手执狂刀似乎有着摧毁一切的无穷力量',7,100380,100390);

/*!40000 ALTER TABLE `hard_level` ENABLE KEYS */;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
");
?>
