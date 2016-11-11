<?php
$this->AddSQL("

alter table player_mission_level_record add column `buy_times` int(11)  default '0' comment 'boss关卡今日购买次数'; 
alter table player_mission_level_record add column `buy_update_time` bigint(20)  default '0' comment 'boss关卡上次购买时间戳'; 

");


?>
