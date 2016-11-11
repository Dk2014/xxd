<?php


	$export_table = array(
		'table' 			=> 'item_costprops',
		'export_order' 		=> ' ORDER BY `item_id`',
		'exclude_columns'	=> array(
		),
	);

	export_table($export_table);

?>