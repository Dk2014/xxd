<?php


	$export_table = array(
		'table' 			=> 'mission_level_box',
		'export_order' 		=> ' ORDER BY `mission_level_id`',
		'exclude_columns'	=> array(
			'order',
			'award_chance',
		),
	);

	export_table($export_table);

?>