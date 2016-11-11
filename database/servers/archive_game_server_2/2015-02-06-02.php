<?php

/*
// flush player_shaded_mission_record and later we will regenerate these data
$this->AddSQL("
delete from `player_shaded_mission_record`;
");

$mission2shadows = array();

$shaded_mission = $this->NewQuery("
select `id`, `mission_level_id` from `shaded_mission`;
");

while($row = $shaded_mission->GoNext()) {
	$shaded_id = $row['id'];
	$mission_id = $row['mission_level_id'];
	if(!isset($mission2shadows[$mission_id])) {
		$mission2shadows[$mission_id] = array();
	}
	$mission2shadows[$mission_id][]=$shaded_id;
}

$this->DropQuery($shaded_mission);

$player_mission_record = $this->NewQuery("
select `pid`, `mission_level_id` from `player_mission_level_record`;
");

while($row = $player_mission_record->GoNext()) {
	$pid = $row['pid'];
	$mission_level_id = $row['mission_level_id'];
	foreach($mission2shadows[$mission_level_id] as $shaded_id) {
		$auto_id = $this->GetAutoID($pid, 'player_shaded_mission_record');
		$this->AddSQL("
insert into `player_shaded_mission_record`(`id`, `pid`, `waiting_shadow_id`, `mission_level_id`) values({$auto_id}, {$pid}, {$shaded_id}, {$mission_level_id});
		");
	}
}

$this->DropQuery($player_mission_record);
 */

?>

