<?php


	$export_table = array(
		'table' 			=> 'coins_exchange',
		'export_order' 		=> ' ORDER BY `unique_key`',
		'exclude_columns'	=> array(
			'id',
		),
	);

	export_table($export_table);

?>
