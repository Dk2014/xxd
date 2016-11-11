<?php
//帮派送镖个人信息
$this->AddSQL("
drop table if exists player_global_clique_escort;
create table player_global_clique_escort (
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`daily_escort_timestamp` bigint(20) NOT NULL DEFAULT '0' COMMENT '今日首次送镖时间戳',
	`daily_escort_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '送镖次数',

	`escort_boat_type` smallint(6) NOT NULL DEFAULT '0' COMMENT '当前标船类型',
	`status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0--无运镖 1--运镖中 2--等待领取奖励',
	`hijacked` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0--没有被完成过劫持 1--被完成过劫持 ',

	PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家帮派运镖信息';
");

$this->AddSQL("
drop table if exists global_clique_boat;
create table global_clique_boat (
	`id` bigint(20) NOT NULL COMMENT 'ID',
	`clique_id` bigint(20) NOT NULL COMMENT '所属帮派ID',
	`boat_type` smallint(6) NOT NULL COMMENT '镖船类型',

	`owner_pid` bigint(20) NOT NULL COMMENT '护送人PID',
	`escort_start_timestamp` bigint(20) NOT NULL DEFAULT '0' COMMENT '本次送镖时间戳',
	`total_escort_time` smallint(6) NOT NULL DEFAULT '0' COMMENT '累计送镖时间',

	`recover_pid` bigint(20) NOT NULL DEFAULT '0' COMMENT '夺回玩家PID',
	
	`hijacker_pid` bigint(20) NOT NULL COMMENT '劫持人PID',
	`hijack_timestamp` bigint(20) NOT NULL DEFAULT '0' COMMENT '本次送镖时间戳',

	`status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态 0--运送中 1--劫持中 2--夺回中 3--劫持完成 ',

	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='镖船信息';
");

$this->AddSQL("
drop table if exists player_global_clique_hijack;
create table player_global_clique_hijack (
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`daily_hijack_timestamp` bigint(20) NOT NULL DEFAULT '0' COMMENT '劫持时间戳',
	`daily_hijack_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '劫持次数',

	`hijacked_boat_type` smallint(6) NOT NULL DEFAULT '0' COMMENT '待奖励的标船类型',
	`status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0--无劫镖 1--劫镖中 2--等待领取奖励',

	PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家劫镖信息';

");

$this->AddSQL("
create table if not exists clique_boat (
	`id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '主键',
	`name` varchar(20) NOT NULL DEFAULT '' COMMENT '镖船名称',
	`sign` varchar(50) NOT NULL COMMENT '资源标识',
	`rate` tinyint(4) NOT NULL DEFAULT '0' COMMENT '概率',
	`escort_time` smallint(6) NOT NULL DEFAULT '0' COMMENT '运送时间（单位分钟）',

	`select_cost_ingot` bigint(20) NOT NULL DEFAULT '0' COMMENT '直接选择话费元宝（0则不可直接选择）',

	`award_coins` bigint(20) NOT NULL DEFAULT '0' COMMENT '奖励铜钱',
	`award_fame` int(11) NOT NULL DEFAULT '0' COMMENT '奖励声望',
	`award_clique_contrib` bigint(20) NOT NULL DEFAULT '0' COMMENT '奖励贡献',

	`hijack_lose_coins` bigint(20) NOT NULL DEFAULT '0' COMMENT '抢劫损失铜钱',
	`hijack_lose_fame` int(11) NOT NULL DEFAULT '0' COMMENT '抢劫损失声望',
	`hijack_lose_clique_contrib` bigint(20) NOT NULL DEFAULT '0' COMMENT '抢劫损失贡献',

	PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='帮派运镖镖船配置';
");

$this->AddSQL("
	drop table if exists player_global_clique_escort_message;
create table player_global_clique_escort_message (
	`id` bigint(20) NOT NULL COMMENT '玩家ID',
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`tpl_id` smallint(6) NOT NULL COMMENT '模版ID',
	`parameters` varchar(1024) NOT NULL COMMENT '模版参数',
	`timestamp` bigint(20) NOT NULL DEFAULT '0' COMMENT '发送时间戳',

	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家帮派运镖消息';
");
//$this->AddSQL("
//
//CREATE TABLE `escort_message` (
//	`id` smallint(6) NOT NULL AUTO_INCREMENT COMMENT '消息模版ID',
//	`sign` varchar(30) DEFAULT NULL COMMENT '唯一标识',
//	`desc` varchar(30) NOT NULL COMMENT '描述',
//	`parameters` varchar(1024) NOT NULL COMMENT '参数',
//	`content` varchar(1024) NOT NULL COMMENT '内容',
//	PRIMARY KEY (`id`),
//	UNIQUE KEY `sign` (`sign`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='镖船消息模板';
//");
?>
