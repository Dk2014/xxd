<?php

db_execute($db, "

ALTER TABLE `sword_soul_quality`
DROP COLUMN `init_exp`;

");
?>
