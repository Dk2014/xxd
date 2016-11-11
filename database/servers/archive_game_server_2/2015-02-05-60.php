<?php

// 命锁功能
$fate = 1048576;

$player_func_key_query = $this->NewQuery("select * from `player_func_key`");

while($player_func_dat = $player_func_key_query->GoNext()){
	$pid = $player_func_dat['pid'];
	$new_key = $player_func_dat['unique_key'] & ($fate - 1);
	$query = $this->NewQuery("select * from `player_fate_box_state` where pid={$pid} LIMIT 1;");
	$query_dat = $query->GoNext();
	if(isset($query_dat) && $query_dat['pid'] > 0){
		$new_new_key = $new_key + $fate;
		$this->AddSQL("update `player_func_key` set unique_key={$new_new_key} where pid={$pid};");
	}else{
		$this->AddSQL("update `player_func_key` set unique_key={$new_key} where pid={$pid};");
	}
	$this->DropQuery($query);
}

$this->DropQuery($player_func_key_query);
?>