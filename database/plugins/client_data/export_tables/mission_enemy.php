<?php


	$export_table = array(
		'table' 			=> 'mission_enemy',
		'export_order' 		=> ' ORDER BY `mission_level_id`, `shaded_mission_id` asc, `order` asc',
		'exclude_columns'	=> array(
			'monster_num',
			'monster1_chance',
			'monster2_chance',
			'monster3_chance',
			'monster4_chance',
			'monster5_chance',
		),
	);

	export_table($export_table);

?>
