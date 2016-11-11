<?php

db_execute($db, "

ALTER TABLE `multi_level` ADD COLUMN `name` varchar(20) NOT NULL DEFAULT '' COMMENT '关卡名称' after `music`;
ALTER TABLE `multi_level` DROP COLUMN `daily_num`;

");
?>