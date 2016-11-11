<?php

db_execute($db, "
create table player_meditation_state (
	`pid` bigint(20) NOT NULL AUTO_INCREMENT,
	`accumulate_time` int(11) NOT NULL COMMENT '光明钥匙奖励累积时间',
	PRIMARY KEY (`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='玩家打坐状态';

");

?>
