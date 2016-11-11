<?php

$this->AddSQL("

ALTER TABLE `player_pve_state` add  COLUMN `daily_num` tinyint(4) NOT NULL DEFAULT '0' COMMENT '今日进入次数';

");

?>
