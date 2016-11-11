<?php
db_execute($db, "

ALTER TABLE `player_physical` ADD COLUMN `daily_count` tinyint(1) DEFAULT '0' COMMENT '当天购买次数';
");
?>