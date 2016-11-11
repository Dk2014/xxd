<?php

	$export_table = array(
		'table' 			=> 'skill',
		'export_order' 		=> ' ORDER BY `id`',
		'exclude_columns'	=> array(
			'config',
		),
		'addition' => array(
			'DecPower' => 'int, 消耗精气',
			'TargetMode' => 'int, 攻击方式',
			'Skill_BuffTarget' =>'string,Buff目标',
		),
	);

function get_addition_Skill_BuffTarget($db) {
	$result = array();
	$row = db_query($db, "select id,config from skill");
	foreach ($row as $val) {
		$target = '"';
		if (isset($val['config'])) {
			$config = json_decode($val['config']);
			if ($config == NULL) {
				$result[] = '""';
			}else{
				//TODO
				if(isset($config->SelfBuffs) && count($config->SelfBuffs) > 0 ) {
					$target .= '-1,';
				}
				if(isset($config->BuddyBuffs)) {
					foreach($config->BuddyBuffs as $buff) {
						$target .= "{$buff->TargetMode},";
					}
				}
				$target .= '"';
				$result[] = $target;
			}
		} else {
			$result[] = '""';
		}
	}
	return $result;
}

function get_addition_DecPower($db) {
	$result = array();
	$row = db_query($db, "select id,config from skill");
	foreach ($row as $val) {
		if (isset($val['config'])) {
			$config = json_decode($val['config']);
			if ($config == NULL) {
				$result[] = 0;
			}else{
				$result[] = $config->DecPower;
			}
		} else {
			$result[] = 0;
		}
	}
	return $result;
}

function get_addition_TargetMode($db) {
	$result = array();
	$row = db_query($db, "select id,config from skill");
	foreach ($row as $val) {
		if (isset($val['config'])) {
			$config = json_decode($val['config']);
			if ($config == NULL) {
				$result[] = 0;
			}else{
				$result[] = $config->TargetMode;
			}
		} else {
			$result[] = 0;
		}
	}
	return $result;
}

export_table($export_table);

?>
