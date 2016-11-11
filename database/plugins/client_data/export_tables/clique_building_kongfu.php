<?php


	$export_table = array(
		'table' 			=> 'clique_building_kongfu',
		'export_order' 		=> ' ORDER BY `building`,`level`',
		'exclude_columns'	=> array(),
	);

	export_table($export_table);

?>
