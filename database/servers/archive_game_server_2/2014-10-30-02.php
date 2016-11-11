<?php

//修复旧玩家没有灵宠数据
$players = $this->NewQuery("select `id` from player;");

while ($player = $players->GoNext()) {
	$pid = $player['id'];
	$anyPet = $this->NewQuery("select * from player_battle_pet limit 1");
	$havePet = false;
	if ($anyPet->Have()) {
		$havePet = true;
	}
	$this->DropQuery($anyPet);

	if ($havePet) {
		foreach(array(1,2,3,4,5) as $grid_id) {
			$grid = $this->NewQuery("select grid_id from `player_battle_pet_grid` where grid_id={$grid_id} and pid={$pid}");
			$haveGrid = false;
			if ($grid->Have()) {
				$haveGrid = true;
			}
			
			$this->DropQuery($grid);

			if (!$haveGrid) {
				$id = $this->GetAutoID($pid);
				$this->AddSQL("insert into `player_battle_pet_grid` (`id`, `pid`, `grid_id`) values('{$id}', '{$pid}', '{$grid_id}')");
			}
		}
	}
}

?>
