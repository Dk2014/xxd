<?php
$func_sword_soul = 16;
$func_battle_pet = 32;
$func_arena = 128;
$func_chest_draw = 256;
$func_ghost = 1;

$func_pve  = 131072;

$player_data_query =  $this->NewQuery("select * from player_func_key");

while($player_dat = $player_data_query->GoNext()) {
	$pid = $player_dat['pid'];
	$unique_key = intval($player_dat['unique_key']);

	$ghost_query = $this->NewQuery("select * from  player_ghost_state where pid={$pid}");
	if($ghost_query->Have()) {
		$unique_key = ($unique_key | $func_ghost);
	}
	$this->DropQuery($ghost_query);

	$battle_pet_query = $this->NewQuery("select * from  player_battle_pet_state where pid={$pid}");
	if($battle_pet_query->Have()) {
		$unique_key = ($unique_key | $func_battle_pet);
		$unique_key = ($unique_key | $func_pve);
	}
	$this->DropQuery($battle_pet_query);

	$arena_query = $this->NewQuery("select * from  player_arena where pid={$pid}");
	if($arena_query->Have()) {
		$unique_key = ($unique_key | $func_arena);
	}
	$this->DropQuery($arena_query);

	$chest_query = $this->NewQuery("select * from  player_chest_state where pid={$pid}");
	if($chest_query->Have()) {
		$unique_key = ($unique_key | $func_chest_draw);
	}
	$this->DropQuery($chest_query);

	$this->AddSQL("update player_func_key set `unique_key`={$unique_key} where pid={$pid}");
}

$this->DropQuery($player_data_query);

?>
