<?php
db_execute($db, "

alter table `player_resource_level` add column `exp_pass_time`  bigint(20) NOT NULL COMMENT '经验关卡通关时间';
alter table `player_resource_level` change column `pass_time` `coin_pass_time` bigint(20) NOT NULL COMMENT '铜钱关卡通关时间';

");
?>
