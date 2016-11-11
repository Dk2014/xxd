<?php
db_execute($db, "
create table `login_award` (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '奖励ID',
	`requrie_active_days` int(11) NOT NULL COMMENT '累计活跃天数',
	`award_type` tinyint(4) NOT NULL COMMENT '奖励类型',
	`award_id` int(11) NOT NULL COMMENT '奖励内容外键',
	`award_num` int(11) NOT NULL COMMENT '奖励数量',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='累计登录奖励';

create table `player_login_award_record` (
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`active_days` int(11) NOT NULL DEFAULT '0' COMMENT '累计活跃天数',
	`record` smallint(6) NOT NULL DEFAULT '0' COMMENT '七天奖励领取记录', 
	PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家奖励领取情况';
")
?>
