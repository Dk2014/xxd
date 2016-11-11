<?php 
db_execute($db, "
create table IF NOT EXISTS `player_month_card_info` (
	`pid` bigint(20) NOT NULL COMMENT '用户ID',
	`starttime` bigint(20) NOT NULL DEFAULT 0 COMMENT '月卡开始时间',
	`endtime` bigint(20) NOT NULL DEFAULT 0 COMMENT '月卡结束时间',
	`buytimes` bigint(10) NOT NULL DEFAULT 0 COMMENT '购买次数',
	`presenttotal` bigint(10) NOT NULL DEFAULT 0 COMMENT '赠送总金额',
	PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家月卡信息'; 

");
?>
