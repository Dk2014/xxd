
<?php
db_execute($db, 
	"
		CREATE TABLE `quest` (
		  `id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '任务ID',
		  `order` int(11) NOT NULL COMMENT '任务排序，低的排前',
		  `name` varchar(10) NOT NULL COMMENT '任务标题',
		  `type` tinyint(4) DEFAULT '0' COMMENT '任务类型',
		  `desc` varchar(240) DEFAULT NULL COMMENT '简介',
		  `require_level` int(11) NOT NULL COMMENT '要求玩家等级',
		  `town_id` smallint(6) NOT NULL DEFAULT '-1' COMMENT '城镇ID',
		  `town_npc_id` int(11) NOT NULL COMMENT '完成任务所需的城镇NPC ID',
		  `mission_id` int(11) NOT NULL COMMENT '完成任务所需的区域ID',
		  `mission_level_id` int(11) DEFAULT '0' COMMENT '完成任务所需的关卡ID',
		  `enemy_num` tinyint(4) NOT NULL DEFAULT '0' COMMENT '敌人组数',
		  `enemy_id` smallint(6) DEFAULT '0' COMMENT '敌人ID',
		  `drama_mode` tinyint(4) DEFAULT '0' COMMENT '剧情模式(1--任务完成播放剧情)',
		  `award_exp` int(11) NOT NULL COMMENT '奖励经验',
		  `award_coins` bigint(20) NOT NULL DEFAULT '0' COMMENT '奖励铜钱',
		  `award_item_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品',
		  `award_func_key` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励功能权值',
		  `award_role_id` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '奖励角色ID',
		  `award_role_level` tinyint(4) NOT NULL DEFAULT '0' COMMENT '奖励角色等级',
		  `award_mission_key` int(11) NOT NULL DEFAULT '0' COMMENT '奖励区域权值',
		  `award_town_key` int(11) NOT NULL DEFAULT '0' COMMENT '奖励城镇权值',
		  `award_physical` tinyint(4) NOT NULL DEFAULT '0' COMMENT '奖励体力',
		  `auto_mission_level_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '自动进入关卡id',
		  PRIMARY KEY (`id`),
		  UNIQUE KEY `name` (`name`)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='主线任务';
	"
);

?>