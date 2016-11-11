<?php

db_execute($db, "

ALTER TABLE `mission_level_small_box` ADD COLUMN `items_kind` tinyint(4) NOT NULL COMMENT '出现物品有几种';

");
?>
