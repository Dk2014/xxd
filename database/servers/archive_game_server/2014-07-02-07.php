<?php
db_execute($db, "

ALTER TABLE `ghost_passive_skill`
DROP COLUMN `quality`;

");
?>

