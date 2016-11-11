<?php


	$export_table = array(
		'table' 			=> 'hard_level',
		'export_order'	=> '',
		'exclude_columns'	=> array(
			'award_hard_level_lock',
		),
	);

	export_table($export_table);

?>
