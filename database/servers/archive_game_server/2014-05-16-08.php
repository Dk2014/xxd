<?php
db_execute($db, "

ALTER TABLE `sword_soul` drop column `kendo_level`;

");
?>