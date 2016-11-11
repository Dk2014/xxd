<?php

db_execute($db, "

ALTER TABLE `item`
ADD COLUMN `panel`  tinyint(4) NOT NULL DEFAULT 0 COMMENT '指向面板' AFTER `can_use`;

");
?>
