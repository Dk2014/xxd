<?php

$role = array();
$dbRole = $this->NewQuery("select * from role");
while ($rs = $dbRole->GoNext()) {
	$role[$rs['id']] = $rs;
	$this->AddSQL("update player_use_skill set skill_id1 = {$rs['skill_id1']}, skill_id2 = {$rs['skill_id2']}, skill_id3 = 0 where role_id={$rs['id']};");
}
$this->DropQuery($dbRole);


$dbPlayerRole = $this->NewQuery("select * from player_role");
while ($rs = $dbPlayerRole->GoNext()) {
	$id1 = $this->GetAutoID($rs['pid']);
	$id2 = $this->GetAutoID($rs['pid']);
	$skill = $role[$rs['role_id']];
	$this->AddSQL("insert into player_skill (id, pid, role_id, skill_id)values({$id1}, {$rs['pid']}, {$rs['role_id']}, {$skill['skill_id1']}),({$id2}, {$rs['pid']}, {$rs['role_id']}, {$skill['skill_id2']});");
}
$this->DropQuery($dbPlayerRole);
?>
