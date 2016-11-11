<?php
db_execute($db, "
alter table  `player_info` add column `new_arena_report_num` smallint(6) NOT NULL DEFAULT '0' COMMENT '玩家离线比武场战报数';
")
?>
