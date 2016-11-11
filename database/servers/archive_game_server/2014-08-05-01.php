<?php
db_execute($db, "
delete from enemy_deploy_form where battle_type=1 and id in(32, 33, 34, 43, 44, 45, 47, 49);
");

?>
