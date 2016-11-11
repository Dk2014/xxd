<?php


	$export_table = array(
		'table' 			=> 'announcement',
		'export_order' 		=> ' ORDER BY `id`',
		'exclude_columns'	=> array(
			'sign',
			'name',
		),
	);

	export_table($export_table);

?>
