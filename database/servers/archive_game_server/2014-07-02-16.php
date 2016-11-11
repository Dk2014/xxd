<?php
db_execute($db, "

ALTER TABLE `player_arena` ADD COLUMN `new_record_count`  smallint(6) NOT NULL COMMENT '新战报计数';

");
?>

