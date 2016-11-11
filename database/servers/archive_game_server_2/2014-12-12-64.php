<?php 
$players = $this->NewQuery("select id from player");
while ($player = $players->GoNext()) {
	$pid = $player['id'];
	$formation = $this->NewQuery("SELECT * FROM `player_formation` WHERE pid={$pid}");
	$mainRole = $this->NewQuery("select role_id from player_role where pid={$pid} and role_id in (1,2)");
	$mainRoleData = $mainRole->GoNext();
	$mainRoleId = $mainRoleData['role_id'];
	$roleSet = array();
	$dupRole = false;
	$count = 0;


	$row = $formation->GoNext();
	$poses = array('pos0', 'pos1', 'pos2', 'pos3', 'pos4', 'pos5');
	foreach($poses as $key=>$pos) {
		if ($row[$pos] != -1) {
			if(isset($roleSet[$row[$pos]])) {
				$dupRole = true;
			}
			$roleSet[$row[$pos]] = true;
			$count += 1;
		}
		if ($row[$pos] == 1 || $row[$pos] == 2) {
			$mainRoleId = $row[$pos];
		}
	}
	if ($count > 3 || $dupRole) {
		$this->AddSQL("update player_formation set pos0={$mainRoleId}, pos1=-1, pos2=-1, pos3=-1, pos4=-1, pos5=-1 where pid={$pid}");
	}
	$this->DropQuery($formation);

}
$this->DropQuery($players);
?>

