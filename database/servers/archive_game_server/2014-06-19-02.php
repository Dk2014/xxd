<?php

db_execute($db, "

ALTER TABLE `player_sword_soul_equipment`
ADD COLUMN `type_bits`  smallint(6) NOT NULL COMMENT '所有装备剑心类型的位标记' AFTER `is_equipped_xuanyuan`;

");
?>
