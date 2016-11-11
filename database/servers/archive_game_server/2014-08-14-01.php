<?php
db_execute($db, "

alter table `skill_content` drop column `recover_round_num`;
");

?>
