<?php


$this->AddSQL("update `player_battle_pet` set parent_pet_id=91 where parent_pet_id=0");

$players = $this->NewQuery("select  player.id as pid, player_func_key.key as func_key  from player  inner join player_func_key on player.id=player_func_key.pid and player_func_key.key >= 1450;");

while ($player = $players->GoNext()) {
	$pid = $player['pid'];
	$id = $this->GetAutoId($pid);
	$petQuery = $this->NewQuery("select count(*) as num from player_battle_pet where pid={$pid} and parent_pet_id=91");
	$petDat = $petQuery->GoNext();
	$num = $petDat['num'];
	if ($num > 1 ) {
		$this->AddSQL("delete from player_battle_pet where pid={$pid} and battle_pet_id=91;");
	}
	$this->DropQuery($petQuery);
}

$this->DropQuery($players);

?>
