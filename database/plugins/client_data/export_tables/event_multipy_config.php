<?php


	$export_table = array(
		'table' 			=> 'event_multiply_config',
		'export_order' 		=> ' ORDER BY `id`',
		'exclude_columns'	=> array(
			'id',
		),
	);

	export_table($export_table);

?>