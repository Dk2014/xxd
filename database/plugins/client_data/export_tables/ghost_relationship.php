<?php


	$export_table = array(
		'table' 			=> 'ghost_relationship',
		'export_order' 		=> ' order by `id` ',
		'exclude_columns'	=> array(
			'id',
		),
	);

	export_table($export_table);

?>
