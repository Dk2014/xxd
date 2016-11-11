<?php
db_execute($db, "
drop table if exists npc_talk;
create table npc_talk (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '记录ID',
	`npc_id` int(11) NOT NULL COMMENT 'NPC ID',
	`town_id` smallint(6) NOT NULL COMMENT '相关城镇',
	`type` tinyint(4) NOT NULL COMMENT '对话类型 1--首次对话； 2--人物对话',
	`quest_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '关联主线任务',
	`conversion` varchar(1024) NOT NULL COMMENT '对话内容',
	`award_item_id`  smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品ID',
	`award_item_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '奖励物品数量',
	PRIMARY KEY (`id`)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='城镇NPC对话';


drop table if exists player_npc_talk_record;
create table player_npc_talk_record (
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '记录ID',
	`pid` bigint(20) not null  comment '玩家ID',
	`npc_id` int(11) not null comment 'NPC ID',
	`town_id` smallint(6) NOT NULL COMMENT '相关城镇',
	primary key (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家与NPC首次对话记录';
");
?>
