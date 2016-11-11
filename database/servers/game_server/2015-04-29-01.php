<?php

/*
 *  join_time: 用于48小时内推出帮派扣除铜币建筑贡献需求；
 *  contrib_clear_time 和 contrib : 每周一清零贡献
 *  donate_coins_time 和 daily_donate_coins : 贡献铜币时间戳，用于限制每天只能贡献有限的铜币
 *  donate_coins_X_building: 累计贡献铜币，假如不满48小时的玩家需要扣除各建筑投入的铜币
 *  worship_time:上香时间，限制玩家每天只能上香一次
 */
$this->AddSQL("
create table player_global_clique_info (
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`clique_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '帮派名称', 
	`join_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '加入帮派时间', 


	`contrib` bigint(20) NOT NULL DEFAULT '0' COMMENT '帮派贡献', 
	`contrib_clear_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '帮派贡献结算时间(每周一清理贡献)', 

	`donate_coins_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '贡献铜币时间戳(限制捐献大小)', 
	`daily_donate_coins` bigint(20) NOT NULL DEFAULT '0' COMMENT '当日贡献铜币', 

	`exchange_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '金银券兑换时间戳', 
	`gold_exchange_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '每日金劵兑换数', 
	`silver_exchange_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '每日银劵兑换数', 

	`donate_coins_center_building` bigint(20) NOT NULL DEFAULT '0' COMMENT '总部累积贡献铜币', 
	`donate_coins_temple_building` bigint(20) NOT NULL DEFAULT '0' COMMENT '宗祠累积贡献铜币', 
	`donate_coins_bank_building` bigint(20) NOT NULL DEFAULT '0' COMMENT '钱庄累积贡献铜币', 
	`donate_coins_health_building` bigint(20) NOT NULL DEFAULT '0' COMMENT '回春堂累积贡献铜币', 
	`donate_coins_attck_building` bigint(20) NOT NULL DEFAULT '0' COMMENT '神兵堂累积贡献铜币', 
	`donate_coins_defense_building` bigint(20) NOT NULL DEFAULT '0' COMMENT '金刚堂累积贡献铜币', 

	`health_level` smallint(6) NOT NULL DEFAULT '0' COMMENT '回春等级',
	`attack_level` smallint(6) NOT NULL DEFAULT '0' COMMENT '神兵等级',
	`defense_level` smallint(6) NOT NULL DEFAULT '0' COMMENT '金刚等级',

	`worship_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '上香时间', 

	PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家帮派信息';
");


//$this->AddSQL("
//create table player_clique_apply (
//	`id` bigint(20) NOT NULL COMMENT '主键',
//	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
//	`clique_id` bigint(20) NOT NULL COMMENT '帮派ID',
//	`apply_time` bigint(20) NOT NULL COMMENT '申请时间戳',
//	PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家帮派申请';
//");

$this->AddSQL("
create table global_clique (
	`id` bigint(20) NOT NULL COMMENT '主键',
	`name` varchar(100) NOT NULL  COMMENT '名字',
	`announce` varchar(200) NOT NULL DEFAULT '' COMMENT '公告',
	`total_donate_coins` bigint(20) NOT NULL DEFAULT '0' COMMENT '累积贡献铜币', 

	`owner_pid` bigint(20) NOT NULL COMMENT '帮主',
	`owner_login_time` bigint(20) NOT NULL COMMENT '帮主最近登录时间',
	`manger_pid1` bigint(20) NOT NULL DEFAULT '0' COMMENT '副帮主1',
	`manger_pid2` bigint(20) NOT NULL DEFAULT '0' COMMENT '副帮主2',

	`center_building_coins` bigint(20) NOT NULL DEFAULT '0' COMMENT '总部当前贡献铜币', 
	`temple_building_coins` bigint(20) NOT NULL DEFAULT '0' COMMENT '宗祠贡献铜币', 
	`health_building_coins` bigint(20) NOT NULL DEFAULT '0' COMMENT '回春堂当前贡献铜币', 
	`attck_building_coins` bigint(20) NOT NULL DEFAULT '0' COMMENT '神兵堂当前贡献铜币', 
	`defense_building_coins` bigint(20) NOT NULL DEFAULT '0' COMMENT '金刚堂当前贡献铜币', 
	`bank_building_coins` bigint(20) NOT NULL DEFAULT '0' COMMENT '钱庄当前贡献铜币', 

	`center_building_level` smallint(6) NOT NULL DEFAULT '1' COMMENT '总部当前等级', 
	`temple_building_level` smallint(6) NOT NULL DEFAULT '0' COMMENT '宗祠等级', 
	`health_building_level` smallint(6) NOT NULL DEFAULT '0' COMMENT '回春堂当前等级', 
	`attck_building_level` smallint(6) NOT NULL DEFAULT '0' COMMENT '神兵堂当前等级', 
	`defense_building_level` smallint(6) NOT NULL DEFAULT '0' COMMENT '金刚堂当前等级', 
	`bank_building_level` smallint(6) NOT NULL DEFAULT '0' COMMENT '钱庄当前等级', 

	`member_json` blob NOT NULL COMMENT '成员列表存储 pid 的json 数组',

PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='帮派信息';
"
);
?>
