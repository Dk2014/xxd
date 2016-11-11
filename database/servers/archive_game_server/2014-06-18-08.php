<?php

db_execute($db, "

ALTER TABLE `sword_soul_quality`
MODIFY COLUMN `price`  bigint(20) NULL DEFAULT NULL COMMENT '售价' AFTER `init_exp`;

");
?>
