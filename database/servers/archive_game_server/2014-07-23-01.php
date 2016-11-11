<?php
db_execute($db,"


alter table `daily_quest` add column `class` smallint(6) NOT NULL COMMENT '任务类别';
alter table `player_daily_quest` add column `class` smallint(6) NOT NULL COMMENT '每日任务类别';
alter table `player_daily_quest` change column `last_finish_time` `last_update_time` bigint(20) NOT NULL COMMENT '最近一次更新时间';

"
);
?>
