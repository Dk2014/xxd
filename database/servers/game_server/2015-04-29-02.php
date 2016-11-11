<?php
$this->AddSQL("


DROP TABLE IF EXISTS `world_channel_message`;

CREATE TABLE `world_channel_message` (
	`id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '消息模版ID',
	`desc` varchar(30) NOT NULL COMMENT '描述',
	`sign` varchar(30) DEFAULT NULL COMMENT '唯一标识',
	`parameters` varchar(1024) NOT NULL COMMENT '参数',
	`content` varchar(1024) NOT NULL COMMENT '内容',
	PRIMARY KEY (`id`),
UNIQUE KEY `sign` (`sign`)
	      ) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COMMENT='世界频道消息模版模板';


INSERT INTO `world_channel_message` (`id`, `desc`, `sign`, `parameters`, `content`)
	VALUES
	(1,'剑山拔剑获得','DrawSwordSoul','player,player,玩家;item,item,道具','{0}在<font color=#c503fd >【剑山拔剑】</font>中提手拔剑获得了{1}'),
(2,'命锁宝箱获得装备','FateBoxEquipment','player,player,玩家;item,item,道具','{0}在<font color=#c503fd >【命锁】</font>中好运连连获得了{1}'),
(3,'命锁宝箱获得魂侍碎片','FateBoxGhostFrame','player,player,玩家;item,item,道具','{0}在<font color=#c503fd >【命锁】</font>中好运连连获得了{1}'),
(4,'彩虹桥获得魂侍','RainbowLevelGhost','player,player,玩家;string,level,关卡;item,item,道具','{0}在{1}中突破挑战获得了金色魂侍{2}'),
(5,'合成金色魂侍','ComposeGhost','player,player,玩家;item,item,道具','{0}经过了长期累积获得了金色魂侍{1}'),
(6,'召唤金色阵印','CallTotem','player,player,玩家;item,item,道具','{0}使用<font color=#c503fd >【圣器召唤】</font>召唤出守护阵印{1}');



");
?>
