<?php

$this->AddSQL("
	ALTER TABLE `player_global_clique_info` ADD COLUMN `glod_exchange_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '金卷购买时间戳';
	");


?>