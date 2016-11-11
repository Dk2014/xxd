<?php
db_execute($db, "
alter table player_physical modify column `daily_count` smallint(6) DEFAULT '0' COMMENT '当天购买次数';
");

?>
