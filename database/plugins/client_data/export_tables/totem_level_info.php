<?php


	$export_table = array(
		'table' 			=> 'totem_level_info',
		'export_order' 		=> ' ORDER BY `quality`, `level`',
		'exclude_columns'	=> array(),
	);

	export_table($export_table);

?>
