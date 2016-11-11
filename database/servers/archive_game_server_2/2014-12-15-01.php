<?php

$this->AddSQL("
alter table `player_ghost` add column `role_id` tinyint(4) NOT NULL DEFAULT '0' COMMENT '魂侍装备在那个角色身上 0 未装备';
");

