<?php

//玩家活跃度记录

$players = $this->NewQuery("select * from player;");

while ($player = $players->GoNext()) {
	$player_id = $player['id'];

	$pa = $this->NewQuery("SELECT * FROM `player_activity` WHERE pid={$player_id}");
	if (!$pa->Have()) {
		$this->AddSQL("INSERT INTO `player_activity` VALUES('{$player_id}',0,0)");
	}
	$this->DropQuery($pa);

}
$this->DropQuery($players);
?>