<?php
$this->AddSQL("
delete from player_pve_state where pid in (select pid from player_func_key where `key` < 1450);

insert ignore into player_pve_state (pid) (select pid from player_func_key where `key` >=1450);
");

?>
