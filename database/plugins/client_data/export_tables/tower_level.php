<?php


	$export_table = array(
		'table' 			=> 'tower_level',
		'export_order' 		=> ' ORDER BY `floor`',
		'exclude_columns'	=> array(
		),
	);

	export_table($export_table);

?>