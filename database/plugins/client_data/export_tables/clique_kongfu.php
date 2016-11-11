<?php


	$export_table = array(
		'table' 			=> 'clique_kongfu',
		'export_order' 		=> ' ORDER BY `building`, `require_building_level` asc',
		'exclude_columns'	=> array(),
	);

	export_table($export_table);

?>
