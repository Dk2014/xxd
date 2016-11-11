<?php
$this->AddSQL("
ALTER TABLE `player_rainbow_level` ADD COLUMN `auto_fight_time` bigint(20) DEFAULT 0 COMMENT '扫荡时间';
");

?>