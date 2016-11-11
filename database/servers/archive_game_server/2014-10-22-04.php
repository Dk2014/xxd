<?php

db_execute($db, "
alter table `player_meditation_state` add column `start_timestamp` bigint(20) NOT NULL DEFAULT '0' COMMENT '打坐开始时间 0-未未打坐状态';

");

?>
