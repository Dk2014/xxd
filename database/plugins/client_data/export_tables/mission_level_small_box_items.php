<?php


	$export_table = array(
		'table' 			=> 'mission_level_small_box_items',
		'export_order' 		=> ' ORDER BY `small_box_id`',
		'exclude_columns'	=> array(
			'small_box_id',
			'probability',
		),
	);

	export_table($export_table);

?>