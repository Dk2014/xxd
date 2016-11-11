<?php


	$export_table = array(
		'filename'          => 'equipment',
		'table' 			=> 'item',
		'export_order' 		=> 'WHERE `type_id` IN (3,4,5,6) ORDER BY `id`',
		'exclude_columns'	=> array(
		),
	);

	export_table($export_table);

?>