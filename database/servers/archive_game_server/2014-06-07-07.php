<?php

db_execute($db, "

ALTER TABLE `quest` ADD COLUMN `auto_fight` tinyint(4) DEFAULT '0' COMMENT '自动打怪';

");
?>