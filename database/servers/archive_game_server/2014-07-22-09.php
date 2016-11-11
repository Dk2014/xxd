<?php
db_execute($db,"

alter table daily_sign_in_award add column `vip_double` tinyint(4) not null default '0' comment 'vip用户获得双倍奖励';

drop table if exists player_daily_sign_in_state;
create table player_daily_sign_in_state(
	`pid` bigint(20) NOT NULL  COMMENT '玩家ID',
	`update_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最新签到时间',
	`index`  tinyint(4) NOT NULL COMMENT '今日签到情况对应record的第几个位[0,7]',
	`record` smallint(6) NOT NULL DEFAULT '0' COMMENT '签到记录',
	PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家最近七日每日签到状态';
"
);
?>
