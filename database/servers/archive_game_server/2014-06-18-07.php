<?php

db_execute($db, "

ALTER TABLE `multi_level`
MODIFY COLUMN `award_exp`  bigint(20) NOT NULL DEFAULT 0 COMMENT '奖励经验' AFTER `require_level`,
MODIFY COLUMN `award_coin`  bigint(20) NOT NULL DEFAULT 0 COMMENT '奖励铜钱' AFTER `award_exp`;

");
?>
