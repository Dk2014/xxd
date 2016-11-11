<?php
db_execute($db, "
alter table `player_rainbow_level` add column `max_pass_segment` smallint(6) NOT NULL DEFAULT '0' COMMENT '打通的最高段数';

alter table `player_rainbow_level` modify column `max_open_segment` smallint(6) NOT NULL DEFAULT '0' COMMENT '可跳转的最大段数';

");
?>
