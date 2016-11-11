<?php
db_execute($db, "

alter table `player_info` drop column `init_global_srv`;

");

?>
