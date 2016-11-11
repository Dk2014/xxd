<?php

$this->AddSQL("
ALTER TABLE `player_hard_level_record` ADD COLUMN `buy_times` smallint(6) DEFAULT 0 COMMENT '深渊关卡今日购买次数';
ALTER TABLE `player_hard_level_record` ADD COLUMN `buy_update_time` bigint(20) DEFAULT 0 COMMENT '深渊关卡上次购买时间戳';
");

?>