<?php

db_execute($db, "

ALTER TABLE `item`
ADD COLUMN `appendix_level` tinyint(4) NOT NULL AFTER `appendix_num`;

");
?>
