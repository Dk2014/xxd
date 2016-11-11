<?php


	$export_table = array(
		'table' 			=> 'events_dinner_awards',
		'export_order' 		=> ' ORDER BY `id`',
		'exclude_columns'	=> array(
			'id',
		),
	);

	export_table($export_table);

?>