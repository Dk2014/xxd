<?php

$this->AddSQL("
alter table `player_driving_sword_map` add `event_mask` blob NOT NULL COMMENT '是否可生成事件区';
");

