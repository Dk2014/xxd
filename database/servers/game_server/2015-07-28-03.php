<?php 

$this->AddSQL("
	CREATE TABLE `player_month_card` (
		`pid` bigint(20) NOT NULL COMMENT '玩家ID',
		`expire_timestamp` bigint(20) NOT NULL DEFAULT '0' COMMENT '过期时间',
		`award_timestamp` bigint(20) NOT NULL DEFAULT '0' COMMENT '领取月卡奖励时间戳',
		PRIMARY KEY (`pid`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家月卡生效时间)';
");
?>
