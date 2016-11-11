<?php

//把所有玩家的功能权值设为0
$this->AddSQL("update player_func_key set `key`=0");

//资源关卡
$resource_level_key = 2048;
//彩虹关卡
$rainbow_level_key = 16384;
//打坐
$meditation_key = 65536;
//灵宠幻境
$pet_virtual_env_key = 131072;
//多人关卡
$multi_level_key = 8192;
//活动关卡（魂侍关卡 伙伴关卡 灵宠关卡）
$function_level_key = 4096;

$player_data_query =  $this->NewQuery("select f.pid as pid, f.unique_key as `unique_key`, r.level as level from player_func_key f left join player_role r on f.pid=r.pid and r.role_id in (1,2);");

while($player_dat = $player_data_query->GoNext()) {
	$pid = $player_dat['pid'];
	$level = intval($player_dat['level']);
	$unique_key = intval($player_dat['unique_key']);
	echo "{$pid} {$level} {$unique_key}\n";
	//资源关卡
	if($level >= 15) {
		$unique_key = ($unique_key | $resource_level_key);
	}
	//打坐
	if($level >= 21) {
		$unique_key = ($unique_key | $meditation_key);
	}
	//彩虹关卡
	if($level >= 25) {
		$unique_key = ($unique_key | $rainbow_level_key);
	}
	//灵宠幻境
	if(($unique_key & 32) > 0) {
		$unique_key = ($unique_key | $pet_virtual_env_key);
	}
	//多人关卡
	if($level >= 30) {
		$unique_key = ($unique_key | $multi_level_key);
	}
	//活动关卡
	if($level >= 35) {
		$unique_key = ($unique_key | $function_level_key);
	}
	$this->AddSQL("update player_func_key set `unique_key`={$unique_key} where pid={$pid}");
}

$this->DropQuery($player_data_query);

?>
