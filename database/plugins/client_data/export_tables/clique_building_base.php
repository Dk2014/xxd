<?php


	$export_table = array(
		'table' 			=> 'clique_center_building_level_info',
		'export_order' 		=> ' ORDER BY `level`',
		'exclude_columns'	=> array(),
	);

	export_table($export_table);

?>