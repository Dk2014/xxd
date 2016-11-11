<?php

$this->AddSQL("
ALTER TABLE `player_global_clique_info` ADD COLUMN `donate_coins_center_building_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '上次总舵捐钱时间戳';
	");

?>