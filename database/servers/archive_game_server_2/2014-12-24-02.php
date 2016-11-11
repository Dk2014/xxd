<?php

$this->AddSQL("

ALTER TABLE `mission_level` ADD COLUMN `award_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '关卡奖励物品类型（客户端用）';

");

