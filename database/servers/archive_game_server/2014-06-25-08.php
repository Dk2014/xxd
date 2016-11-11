<?php

db_execute($db, "

ALTER TABLE `chest_item`
ADD COLUMN `type`  tinyint(4) NOT NULL COMMENT '物品类型' AFTER `chest_id`;

");
?>
