<?php
db_execute($db, "
update `enemy_role` set `sunder_end_hurt_rate` = 150;
update `enemy_role` set `sunder_end_defend_rate` = 20;
");

?>
