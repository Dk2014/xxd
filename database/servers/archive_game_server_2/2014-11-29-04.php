<?php
$players = $this->NewQuery("select  `pid` from `player_role` where `role_id` in (1,2) and `level`>=30 ;");

while($player = $players->GoNext()) {
	$pid = $player['pid'];
	$this->AddSQL("insert ignore `player_multi_level_info` (`pid`, `buddy_role_id`, `buddy_row`, `tactical_grid`, `daily_num`, `battle_time`, `lock`) values ({$pid}, 3, 2, 0, 0, 0, 0)"
	);
}

$this->DropQuery($players);

?>
