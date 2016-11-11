<?php


	$export_table = array(
		'table' 			=> 'ghost_level',
		'export_order' 		=> ' ORDER BY `id`',
		'exclude_columns'	=> array(
			'id',
			'min_add_exp',
			'max_add_exp',
		),
	);

	export_table($export_table);

?>