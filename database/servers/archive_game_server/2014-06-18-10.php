<?php

db_execute($db, "

ALTER TABLE `player_sword_soul_state`
MODIFY COLUMN `fragment_num`  smallint(6) NOT NULL COMMENT '碎片数量' AFTER `pid`,
MODIFY COLUMN `last_is_ingot`  tinyint(4) NOT NULL COMMENT '上次是用为元宝开箱' AFTER `box_state`,
MODIFY COLUMN `protect_num`  tinyint(4) NOT NULL DEFAULT 0 COMMENT '概率保护次数' AFTER `is_first_time`,
ADD COLUMN `free_num`  tinyint(4) NOT NULL AFTER `num`;

");
?>
