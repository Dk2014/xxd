<?php


	$export_table = array(
		'table' 			=> 'battle_item_config',
		'export_order' 		=> ' ORDER BY `id`',
		'exclude_columns'	=> array(
			"effect_type",
			"config",
			"keep",
			"max_override",
			"can_use_level",
		),
		'addition' => array(
			'Value' => 'int,恢复血量',
			'Mod' => 'int,0--精气 1--生命 3--复活 4--捕捉 其他--buff',
		)
	);

function get_addition_Value($db) {
	$result = array();
	$rows = db_query($db, "select id, config from battle_item_config");
	foreach($rows as  $val) {
		if (!isset($val['config'])) {
			continue;
		}
		$configs = json_decode($val['config']);
		if(!$configs) {
			continue;
		}
		$found = false;
		foreach($configs as $config) {
			if ($config->mod == 1 && $config->val > 0) {
				$found = true;
				$result[] = $config->val;
			}
		}
		if(!$found) {
			$result[] = 0;
		}
	}
	return $result;
}

function get_addition_Mod($db) {
	$result = array();
	$rows = db_query($db, "select id, config from battle_item_config");
	foreach($rows as  $val) {
		if (!isset($val['config'])) {
			continue;
		}
		$configs = json_decode($val['config']);
		if(!$configs) {
			continue;
		}
		$found = false;
		foreach($configs as $config) {
			if (isset($config->mod) ) {
				$found = true;
				$result[] = $config->mod;
				break;
			}
		}
		if(!$found) {
			$result[] = -1;
		}
	}
	return $result;
}

export_table($export_table);

?>
