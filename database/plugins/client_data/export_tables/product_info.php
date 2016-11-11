<?php


	$export_table = array(
		'table' 			=> 'product_info',
		'export_order' 		=> ' order by `id` ',
		'exclude_columns'	=> array(),
	);

	export_table($export_table);

?>
