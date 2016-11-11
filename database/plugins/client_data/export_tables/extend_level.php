<?php


	$export_table = array(
		'table' 			=> 'extend_level',
		'export_order' 		=> ' ORDER BY `level_type`,`max_level`',
		'exclude_columns'	=> array(
		),
	);

	export_table($export_table);

?>