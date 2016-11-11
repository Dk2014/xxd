<?php
db_execute($db, "
alter table `player_rainbow_level` drop column `segment_record`;
alter table `player_rainbow_level` add column `max_open_segment` smallint(6) NOT NULL COMMENT '可跳转的最大段数';
");
?>
