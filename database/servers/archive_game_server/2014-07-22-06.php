<?php
db_execute($db,"
create table daily_sign_in_award(
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`type` tinyint(4) NOT NULL COMMENT '签到类型 0--首次登录奖励 1--全局循环奖励',
	`daily_sign_in_id` smallint(6) NOT NULL COMMENT '每日登录ID',
	`award_type` tinyint(4) NOT NULL COMMENT '奖励类型',
	`award_id` smallint(6) NOT NULL COMMENT '奖励物品ID',
	`num` int(11) NOT NULL COMMENT '奖励数量',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='每日签到奖励配置';

create table player_daily_sign_in_record(
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`pid` bigint(20) NOT NULL  COMMENT '玩家ID',
	`sign_in_time` bigint(20) NULL NULL COMMENT '签到时间',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家每日签到记录';

create table player_daily_sign_in_state(
	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`pid` bigint(20) NOT NULL  COMMENT '玩家ID',
	`current_bit`  tinyint(4) NOT NULL COMMENT '今日签到情况对应record的第几个位',
	`record` smallint(6) NOT NULL DEFAULT '0' COMMENT '签到记录',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家每日签到记录';

alter table player_info add column `first_login_time` bigint(20) NOT NULL COMMENT '玩家注册时间';
"
);
?>
