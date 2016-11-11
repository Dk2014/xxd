<?php


	$export_table = array(
		'table' 			=> 'town',
		'export_order' 		=> ' ORDER BY `id`',
		'exclude_columns'	=> array(
		),
	);

	export_table($export_table);

?>