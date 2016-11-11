<?php
$players = $this->NewQuery("select `pid`  from `player_func_key` where `key` >=1450;");
while ($player = $players->GoNext()) {
	$pid = $player['pid'];
	$this->AddSQL("insert ignore `player_battle_pet_state` (pid) values({$pid});");
}



?>
