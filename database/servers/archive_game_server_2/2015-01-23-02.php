<?php

$this->AddSQL("
update `player_use_skill` set `skill_id1`=0, `skill_id2`=0, `skill_id3`=0, `skill_id4`=0 where `role_id`=1 or `role_id`=2;
");

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

function skill_kind2slot($kind) {
	switch($kind) {
		case 1:
			return 1;
		case 2:
			return 4;
		case 3:
			return 2;
		case 4:
			return 3;
	}
}

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
		$slot=skill_kind2slot($grpd_childkind);
		$this->AddSQL("
update `player_use_skill` set `skill_id{$slot}`={$skillRow['skill_id']} where `pid`={$grpd_pid} and `role_id`={$grpd_roleid};
		");
	}
}

$this->DropQuery($playerSkills);

