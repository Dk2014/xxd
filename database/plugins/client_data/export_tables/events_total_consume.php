<?php


	$export_table = array(
		'table' 			=> 'events_total_consume',
		'export_order' 		=> ' ORDER BY `require_cost`',
		'exclude_columns'	=> array(
			'id',
		),
	);

	export_table($export_table);

?>