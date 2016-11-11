<?php

db_execute($db, "

ALTER TABLE `player_friend` ADD COLUMN `block_mode` tinyint(1) NOT NULL DEFAULT '0' COMMENT '黑名单状态:0-否,1-是';

");
?>