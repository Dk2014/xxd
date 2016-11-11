<?php
db_execute($db, "
alter table `player_heart` modify column `add_day_count` int(11) NOT NULL DEFAULT '0' COMMENT '今日好友赠送数量';
alter table `player_heart` modify column `add_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '接受好友赠送爱心时间';
alter table `player_heart` add column `recover_day_count` int(11) NOT NULL DEFAULT '0' COMMENT '今日恢复数量';
");

?>
