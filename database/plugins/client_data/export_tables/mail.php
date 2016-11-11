<?php


	$export_table = array(
		'table' 			=> 'mail',
		'export_order' 		=> ' ORDER BY `id`',
		'exclude_columns'	=> array(
			'sign',
			'type',
		),
	);

	export_table($export_table);

?>