<?php

$pid_result = $this->NewQuery("
select `pid` from `player_formation`;
");

while($pid_row = $pid_result->GoNext()){
	$pid = $pid_row['pid'];
	$this->AddSQL("
insert into `player_team_info`(`pid`, `relationship`, `health_lv`, `attack_lv`, `defence_lv`) values({$pid}, 0, 0, 0, 0);
	");
}

$this->DropQuery($pid_result);

