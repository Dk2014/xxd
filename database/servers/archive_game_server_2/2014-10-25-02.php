<?php

// pr2版本玩家数据升级

$players = $this->NewQuery("select * from player;");

while ($player = $players->GoNext()) {
	$pid = $player['id'];

	// 活动中心玩家奖励记录数据
	$pa = $this->NewQuery("SELECT * FROM `player_event_award_record` WHERE pid={$pid}");
	if (!$pa->Have()) {
		$this->AddSQL("insert `player_event_award_record` set pid = {$pid};");
	}
	$this->DropQuery($pa);

	// 玩家通知中心开关
	$pa = $this->NewQuery("SELECT * FROM `player_push_notify_switch` WHERE pid={$pid}");
	if (!$pa->Have()) {
		$this->AddSQL("insert `player_push_notify_switch` set pid = {$pid};");
	}
	$this->DropQuery($pa);	

	// 玩家打坐信息
	$pa = $this->NewQuery("SELECT * FROM `player_meditation_state` WHERE pid={$pid}");
	if (!$pa->Have()) {
		$this->AddSQL("insert `player_meditation_state` set pid = {$pid}, accumulate_time = 0;");
	}

	$this->DropQuery($pa);
}

$this->DropQuery($players);

?>