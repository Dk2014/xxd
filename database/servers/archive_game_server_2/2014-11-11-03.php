<?php

$this->AddSQL("alter table `player_arena` add column daily_frame int(11) not null default 0 comment '每日奖励声望';
"
);

$this->AddSQL("
alter table `multi_level` add column award_frame int(11) not null default 0 comment '奖励声望';
"
);

?>
