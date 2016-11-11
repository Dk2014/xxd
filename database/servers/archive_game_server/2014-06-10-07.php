<?php

db_execute($db, "

ALTER TABLE `item` DROP COLUMN `cultivation`;
ALTER TABLE `item` DROP COLUMN `dodge_level`;
ALTER TABLE `item` DROP COLUMN `hit_level`;
ALTER TABLE `item` DROP COLUMN `block_level`;
ALTER TABLE `item` DROP COLUMN `critical_level`;
ALTER TABLE `item` DROP COLUMN `critical_hurt_level`;

");
?>