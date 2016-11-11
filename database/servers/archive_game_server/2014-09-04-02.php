<?php
db_execute($db, "
CREATE TABLE `player_rainbow_level` (
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`reset_timestamp` bigint(20) NOT NULL DEFAULT '0' COMMENT '重置彩虹关卡时间戳',
	`reset_num` tinyint(4) NOT NULL COMMENT '今日可重置次数',
  	`segment` smallint(6) NOT NULL COMMENT '段数',
	`order` tinyint(4) NOT NULL COMMENT '彩虹关段内第X关顺序[1,7]',
	`status` tinyint(4) NOT NULL COMMENT '状态 0--打败Boss  1--未领取奖励 2--已奖励未进入下一关卡',
	PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家彩虹关卡状态';

CREATE TABLE `player_rainbow_level_state_bin` (
	`pid` bigint(20) NOT NULL COMMENT '玩家ID',
	`bin` blob NOT NULL COMMENT '彩虹状态',
	PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家彩虹关卡状态';
");
?>
