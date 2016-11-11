<?php

$this->AddSQL("

ALTER TABLE `mission_level`
	ADD COLUMN `award_item` smallint(6) NOT NULL DEFAULT '0' COMMENT '关卡奖励物品',
	ADD COLUMN `award_item_num`  smallint(6) NOT NULL DEFAULT '0' COMMENT '关卡奖励物品数量';

");

