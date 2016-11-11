<?php
/*
// 需要小宇宙爆发才能修好的数据，arm player's skill
$playerSkills=$this->NewQuery("

SELECT `player_skill`.*
	,`skill`.`child_kind`
	,`skill`.`required_level`
FROM `player_skill`
LEFT JOIN `skill` ON `player_skill`.`skill_id` = `skill`.`id`

-- only care about the main role
WHERE `player_skill`.`role_id` IN (
		1
		,2
		)

ORDER BY `pid`
	,`role_id`
	,`child_kind`
	,`required_level` DESC
	,`skill_id` DESC;

");

$grpd_pid=-1;
$grpd_roleid=-1;
$grpd_childkind=-1;
while($skillRow=$playerSkills->GoNext()){
	if($grpd_pid!=$skillRow['pid']
		|| $grpd_roleid!=$skillRow['role_id']
		|| $grpd_childkind!=$skillRow['child_kind']){
		$grpd_pid=$skillRow['pid'];
		$grpd_roleid=$skillRow['role_id'];
		$grpd_childkind=$skillRow['child_kind'];
		$this->AddSQL("
update `player_use_skill` set `skill_id{$grpd_childkind}`={$skillRow['skill_id']} where `pid`={$grpd_pid} and `role_id`={$grpd_roleid};
		", true);
	}
}

$this->DropQuery($playerSkills);
*/
