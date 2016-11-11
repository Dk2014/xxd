<?php

db_execute($db, "

ALTER TABLE `battle_pet` MODIFY COLUMN `live_pos` tinyint(4) NOT NULL COMMENT '召唤后出现的位置(1-前排；2-后排；3-左侧；4-右侧)';

");
?>

