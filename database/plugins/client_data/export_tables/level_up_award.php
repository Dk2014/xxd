<?php


	$export_table = array(
		'table' 			=> 'events_level_up',
		'export_order' 		=> ' ORDER BY `require_level` ASC',
		'exclude_columns'	=> array(
		),
	);

	export_table($export_table);

?>