<?php


//现存数据可能咱脏数据先全部清除
$this->AddSQL("delete from `player_battle_pet_grid`");

$players = $this->NewQuery("select  * from player  inner join player_func_key on player.id=player_func_key.pid;");

while ($player = $players->GoNext()) {
	$pid = $player['id'];
	//检查玩家是否开启灵宠功能
	$openPetFunc = $player['key'] >= 1450; //FUNC_BATTLE_PET 常量如果被修改这里也要即使更新

	if ($openPetFunc) {
		//如果已开启功能则补全缺少的格子
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
	}else {
		//删除尚未开启灵宠功能的格子
		$this->AddSQL("delete from `player_battle_pet_grid` where pid={$pid}");
	}
}

$this->DropQuery($players);

?>
