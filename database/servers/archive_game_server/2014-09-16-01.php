<?php
db_execute($db, "
alter table `rainbow_level` add column `award_id` smallint(16) not null default '0' comment '奖励魂侍ID';

alter table `player_rainbow_level` add column `segment_record` bigint(20) not null default '0' comment '彩虹关卡段首通记录';
");
?>
