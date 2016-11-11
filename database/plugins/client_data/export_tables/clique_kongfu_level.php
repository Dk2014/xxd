<?php


	$export_table = array(
		'table' 			=> 'clique_kongfu_level',
		'export_order' 		=> ' ORDER BY `kongfu_id`, `level` asc',
		'exclude_columns'	=> array(),
	);

	export_table($export_table);

?>
