<?php
db_execute($db, "
alter table `player_heart` modify column `recover_day_count` smallint(6) NOT NULL DEFAULT '0' COMMENT '今日恢复数量';
");

?>
