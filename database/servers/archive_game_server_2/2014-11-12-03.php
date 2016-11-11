<?php

$this->AddSQL("
rename table player_frame to player_fame;
rename table frame_level to fame_level;
rename table frame_system to fame_system;

alter table fame_level change  `required_frame` `required_fame` int(11) NOT NULL COMMENT '要求声望';
alter table fame_system change `max_frame` `max_fame` int(11) NOT NULL COMMENT '最大产生声望';
alter table player_fame change `frame` `fame` int(11) NOT NULL DEFAULT '0' COMMENT '总声望';
alter table player_fame change `mult_level_frame` `mult_level_fame` int(11) NOT NULL DEFAULT '0' COMMENT '多人关卡声望';
alter table player_fame change `arena_frame` `arena_fame` int(11) NOT NULL DEFAULT '0' COMMENT '比武场关卡声望';
alter table player_arena change `daily_frame` `daily_fame` int(11) NOT NULL DEFAULT '0' COMMENT '每日奖励声望';


");

?>
