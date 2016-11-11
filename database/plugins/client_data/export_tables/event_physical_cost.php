<?php


	$export_table = array(
		'table' 			=> 'events_physical_awards',
		'export_order' 		=> ' ORDER BY `require_physical`',
		'exclude_columns'	=> array(
			'id',
		),
	);

	export_table($export_table);

?>