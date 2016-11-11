<?php


	$export_table = array(
		'table' 			=> 'level_battle_pet',
		'export_order' 		=> ' ORDER BY `mission_enemy_id`',
		'exclude_columns'	=> array(
			'id',
			'rate',
		),
	);

	export_table($export_table);

?>
