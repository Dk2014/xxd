<?php

//查询深渊关卡数值
$hard_level_query = $this->NewQuery("select * from mission_level where parent_type=8");
$hard_level_dat = array();
while($row = $hard_level_query->GoNext()) {
	$hard_level_dat[ $row['id'] ] = $row;
}
$this->DropQuery($hard_level_query);

//查询出所有主线关卡数据
$mission_level_query = $this->NewQuery("select * from mission_level where parent_type=0 ");
$mission_level_dat = array();
while($row = $mission_level_query->GoNext()) {
	$mission_level_dat[ $row['id'] ] = $row;
}
$this->DropQuery($mission_level_query);

//查询出所有关卡评星规则
$mission_level_star_query = $this->NewQuery("select * from level_star where two_star_round>0");
$mission_level_star_dat = array();
while($row = $mission_level_star_query->GoNext()) {
	$mission_level_star_dat[$row['level_id']] = $row;
}
$this->DropQuery($mission_level_star_query);

//主线关卡的普通关卡和精英关卡原来没有记录分数，如果已通过则给3星。这里需要一个 award_lock 来判断是否通过
$player_data_query =  $this->NewQuery("
SELECT 	 r.id AS id
	,r.pid AS pid
	,r.mission_level_id AS mission_level_id
	,r.score AS score
	,r.round AS round
	,l.award_lock AS award_lock
FROM player_mission_level_record r
LEFT JOIN player_mission_level l ON r.pid = l.pid;
");

while($player_dat = $player_data_query->GoNext()) {
	$id = $player_dat['id'];
	$pid = $player_dat['pid'];
	$mission_level_id = $player_dat['mission_level_id'];
	$score = $player_dat['score'];
	$round = $player_dat['round'];
	$lock = $player_dat['award_lock'];
	if($round>0) {
		continue;
	}

	if(!isset($mission_level_dat[$mission_level_id])) {
		echo "can NOT find mission_level_data pid:{$pid} record id:{$id} mission_level_id:{$mission_level_id}\n";
		continue;
	}
	if(!isset($mission_level_star_dat[$mission_level_id])) {
		echo "can NOT find mission_star_dat pid:{$pid} record id:{$id} mission_level_id:{$mission_level_id}\n";
		continue;
	}
	$level_dat = $mission_level_dat[$mission_level_id];
	$star_dat = $mission_level_star_dat[$mission_level_id];
	
	//没有通关不需要修复
	if($lock<$level_dat['lock']) {
		continue;
	}

	$round = $star_dat['three_star_round'];
	if($score >=$star_dat['three_star_score']) {
		$round = $star_dat['three_star_round']-1;
	} else if ($score >= $star_dat['two_star_score']) {
		$round = $star_dat['two_star_round']-1;
	} else {
		$round = $star_dat['two_star_round']+1;
	}
        if($round <= 0 ) {
            $round = 1;
        }
	//echo "mission level record:pid {$pid}, mission_level_id {$mission_level_id}\n";
	$this->AddSQL("update player_mission_level_record set round='{$round}' where id={$id}");
}
$this->DropQuery($player_data_query);



//难度关卡全是boss关卡，只要有score就是已经通过
$player_data_query =  $this->NewQuery("select * from player_hard_level_record where score > 0 and round=0");

while($player_dat = $player_data_query->GoNext()) {
	$id = $player_dat['id'];
	$pid = $player_dat['pid'];
	$mission_level_id = $player_dat['level_id'];
	$score = $player_dat['score'];
	if(!isset($hard_level_dat[$mission_level_id])) {
		echo "hard_level: can NOT find mission_level_data pid:{$pid} record id:{$id} mission_level_id:{$mission_level_id}\n";
		continue;
	}
	if(!isset($mission_level_star_dat[$mission_level_id])) {
		echo "can NOT find mission_star_dat pid:{$pid} record id:{$id} mission_level_id:{$mission_level_id}\n";
		continue;
	}
	$level_dat = $hard_level_dat[$mission_level_id];
	$star_dat = $mission_level_star_dat[$mission_level_id];

	$round = $star_dat['three_star_round']-1;
	if($score >=$star_dat['three_star_score']) {
		$round = $star_dat['three_star_round']-1;
	} else if ($score >= $star_dat['two_star_score']) {
		$round = $star_dat['two_star_round']-1;
	} else {
		$round = $star_dat['two_star_round']+1;
	}
	//echo "hard level record:pid {$pid}, mission_level_id {$mission_level_id}\n";
	$this->AddSQL("update player_hard_level_record set round='{$round}' where id={$id}");
}
$this->DropQuery($player_data_query);
?>
