<?php
//总兑换日志
$this->AddSQL("
	CREATE TABLE `global_gift_card_record` (
	  `id` bigint(20) NOT NULL COMMENT '主键',
	  `pid` bigint(20) NOT NULL COMMENT '兑换玩家PID',
	  `code` char(10) NOT NULL COMMENT '兑换码',
	  `timestamp` bigint(20) NOT NULL COMMENT '兑换时间',
	  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='礼品卡兑换记录';
");

//避免玩家重复领取
$this->AddSQL("
	CREATE TABLE `player_global_gift_card_record` (
	  `id` bigint(20) NOT NULL COMMENT '主键',
	  `pid` bigint(20) NOT NULL COMMENT '兑换玩家PID',
	  `code` char(10) NOT NULL COMMENT '兑换码',
	  PRIMARY KEY (`id`),
  KEY `idx_pid` (`pid`)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家兑换记录';
");
?>
