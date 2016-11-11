<?php

$this->AddSQL("

alter table `player_ghost_state` add `last_flush_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最近一次洗点时间';

alter table `player_info` add `last_skill_flush` bigint(20) NOT NULL DEFAULT '0' COMMENT '最近一次洗点时间';

");

