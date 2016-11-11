<?php


	$export_table = array(
		'table' 			=> 'daily_quest',
		'export_order' 		=> ' ORDER BY `require_min_level`, `require_max_level`',
		'exclude_columns'	=> array(
		),
	);

	export_table($export_table);

?>