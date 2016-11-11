<?php


	$export_table = array(
		'table' 			=> 'func_prediction',
		'export_order'	=> ' order by `order` desc',
		'exclude_columns'	=> array(
			'id', 
		),
	);

	export_table($export_table);

?>
