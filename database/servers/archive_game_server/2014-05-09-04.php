<?php

db_execute($db, "

ALTER TABLE `player_item` DROP COLUMN `pos`;
ALTER TABLE `player_item` ADD `is_dressed` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否被装备';

");
?>