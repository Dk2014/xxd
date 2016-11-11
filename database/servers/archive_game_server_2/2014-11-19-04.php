<?php
$players = $this->NewQuery("select `pid`  from `player_func_key` where `key` >=1000;");
while ($player = $players->GoNext()) {
	$pid = $player['pid'];
	$this->AddSQL("insert into `player_ghost_state` (pid) values({$pid});");
}

?>
