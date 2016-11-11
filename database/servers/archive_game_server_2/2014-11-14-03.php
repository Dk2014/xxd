<?php
$this->AddSQL("
delete from level_battle_pet where mission_enemy_id not in (select id from mission_enemy);
");

?>
