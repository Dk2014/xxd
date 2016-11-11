<?php
$this->AddSQL("
CREATE TABLE IF NOT EXISTS `player_qq_vip_gift`(
	`pid` bigint(20) NOT NULL,
	`qqvip` smallint(6) DEFAULT 0 COMMENT 'qq会员的礼包领取记录，1代表开通礼包已领取，2代表续费礼包已领取',
	`surper` smallint(6) DEFAULT 0 COMMENT 'qq超级会员礼包领取记录，值同上',
	primary key(`pid`)
	)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家qq服务礼包领取记录';
	");
?>