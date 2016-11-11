<?php

$this->AddSQL("
alter table `shaded_mission` add `order` tinyint(4) NOT NULL COMMENT '位序';

alter table `player_mission_level_record` add `empty_shadow_bits` smallint(6) NOT NULL DEFAULT '0' COMMENT '清剿过的影之间隙';

drop table `player_shaded_mission_record`;
");


