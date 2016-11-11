<?php
db_execute($db, "

update `mission_enemy` set `is_boss` = 1 where `is_boss` > 0;

");
?>