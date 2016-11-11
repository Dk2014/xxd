<?php
db_execute($db, "

ALTER TABLE `player_arena_record` ADD COLUMN `target_nick`  varchar(50) NOT NULL COMMENT '对手昵称' AFTER `target_pid`;

");
?>

