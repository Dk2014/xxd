<?php


	$export_table = array(
		'table' 			=> 'quest',
		'export_order' 		=> ' ORDER BY `order`',
		'exclude_columns'	=> array(
			'award_mission_level_lock',
			'award_town_key',
			'award_func_key',
		),
	);

	export_table($export_table);

?>