<?php
db_execute($db, "
alter table `player_item` add column `refine_fail_times` smallint(6) NOT NULL DEFAULT '0' COMMENT '精炼失败次数';

");
?>
