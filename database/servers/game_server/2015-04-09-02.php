<?php
$this->AddSQL("alter table `player_formation` drop column `tactical_grid`");

$this->AddSQL("drop table `player_hard_level_state`");
$this->AddSQL("drop table `player_tower_level`");

?>
