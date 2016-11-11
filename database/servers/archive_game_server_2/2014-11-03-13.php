<?php

$this->AddSQL("

ALTER TABLE `player_pve_state` add  COLUMN `enter_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最后一次进入关卡次数';

");

?>
