<?php

db_execute($db, "

ALTER TABLE `equipment_refine_level`
MODIFY COLUMN `gain_pct`  int(11) NOT NULL COMMENT '增益百分比' AFTER `probability`;

");
?>
