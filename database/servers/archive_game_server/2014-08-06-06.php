<?php
db_execute($db, "
alter table player_arena  change daily_award_longbi daily_award_item int(11) NOT NULL DEFAULT '0' COMMENT '今日获得暗影果实累计';
");

?>
