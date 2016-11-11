<?php
db_execute($db, "

alter table `skill` drop column `skill_level`;

");

?>
