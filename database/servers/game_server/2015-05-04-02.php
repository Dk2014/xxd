<?php

$this->AddSQL("
create table player_global_clique_building (
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`silver_exchange_time` bigint(20) NOT NULL DEFAULT '0',
	`gold_exchange_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '每日金劵兑换数',
	`silver_exchange_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '每日银劵兑换数',
	`donate_coins_center_building` bigint(20) NOT NULL DEFAULT '0' COMMENT '总部累积贡献铜币',
	`donate_coins_temple_building` bigint(20) NOT NULL DEFAULT '0' COMMENT '宗祠累积贡献铜币',
	`donate_coins_bank_building` bigint(20) NOT NULL DEFAULT '0' COMMENT '钱庄累积贡献铜币',
	`donate_coins_health_building` bigint(20) NOT NULL DEFAULT '0' COMMENT '回春堂累积贡献铜币',
	`donate_coins_attack_building` bigint(20) NOT NULL DEFAULT '0' COMMENT '神兵堂累积贡献铜币',
	`donate_coins_defense_building` bigint(20) NOT NULL DEFAULT '0' COMMENT '金刚堂累积贡献铜币',
	`health_level` smallint(6) NOT NULL DEFAULT '0' COMMENT '回春等级',
	`attack_level` smallint(6) NOT NULL DEFAULT '0' COMMENT '神兵等级',
	`defense_level` smallint(6) NOT NULL DEFAULT '0' COMMENT '金刚等级',
	`worship_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '上香时间',
	`donate_coins_center_building_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '上次总舵捐钱时间戳',
	`glod_exchange_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '金卷购买时间戳',
	`worship_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上香类型',
	PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='玩家帮派建筑信息';

");

?>
