<?php
db_execute($db, "

drop table player_resource_level_record;

alter table `player_resource_level` add column `coin_daily_num` tinyint(4) NOT NULL COMMENT '经验关卡今日进入次数';
alter table `player_resource_level` add column `exp_daily_num`  tinyint(4) NOT NULL COMMENT '铜钱关卡今日进入次数';

");
?>
