<?php

$this->AddSQL("
alter table skill change `required_frame_level` `required_fame_level`  int(11) NOT NULL DEFAULT '0' COMMENT '需要声望等级';

alter table multi_level change   `award_frame` `award_fame` int(11) NOT NULL DEFAULT '0' COMMENT '奖励声望';

alter table `arena_award_box` change   `frame` `fame` int(11) NOT NULL DEFAULT '0' COMMENT '奖励声望';


");

?>
