<?php


	$export_table = array(
		'table' 			=> 'ghost_passive_skill',
		'export_order' 		=> ' ORDER BY `id`',
		'exclude_columns'	=> array(
		),
	);

	export_table($export_table);

?>