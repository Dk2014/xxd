
<?php
db_execute($db, 
	"
		alter table `quest` change column `award_item_id` `award_item1_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品1';
		alter table `quest` add column `award_item1_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品1数量' after `award_item1_id`;

		alter table `quest` add column `award_item2_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品2' after `award_item1_num`;
		alter table `quest` add column `award_item2_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品2数量' after `award_item2_id`;

		alter table `quest` add column `award_item3_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品3' after `award_item2_num`;
		alter table `quest` add column `award_item3_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品3数量' after `award_item3_id`;

		alter table `quest` add column `award_item4_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品4' after `award_item3_num`;
		alter table `quest` add column `award_item4_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品4数量' after `award_item4_id`;		


		CREATE TABLE `player_quest` (
		  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
		  `quest_id` smallint(6) NOT NULL COMMENT '当前任务ID',
		  `state` tinyint(4) DEFAULT NULL,
		  PRIMARY KEY (`pid`)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家任务';
	"
);

?>