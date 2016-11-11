<?php

db_execute($db, "

ALTER TABLE `item`
MODIFY COLUMN `price`  bigint(20) NOT NULL DEFAULT 0 COMMENT '物品售价' AFTER `desc`;

");
?>
