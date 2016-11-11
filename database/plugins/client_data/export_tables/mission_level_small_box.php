<?php


	$export_table = array(
		'table' 			=> 'mission_level_small_box',
		'export_order' 		=> ' ORDER BY `mission_level_id`',
		'exclude_columns'	=> array(
			'probability',
			'items_kind',
		),
	);

	export_table($export_table);

?>