<?php
$this->AddSQL("
drop table if exists `global_world_chat` ;

create table `global_world_chat` (
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '公告ID',
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
  	`nickname` varchar(50) NOT NULL COMMENT '玩家昵称',
	`timestamp` bigint(20) NOT NULL COMMENT '创建时间戳',
	`content` varchar(1024) NOT NULL COMMENT '内容',
PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='世界频道聊天队列';

drop table if exists `player_global_world_chat_state` ;

create table `player_global_world_chat_state` (
	  `pid` bigint(20) NOT NULL COMMENT '玩家ID',
	  `timestamp` bigint(20) NOT NULL DEFAULT '0' COMMENT '最近一次发送聊天时间戳',
	  `daily_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '今日次数',
PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家世界聊天状态';

");

?>
