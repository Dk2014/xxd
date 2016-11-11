<?php


	$export_table = array(
		'table' 			=> 'events_seven_day_awards',
		'export_order' 		=> ' ORDER BY `require_day`',
		'exclude_columns'	=> array(
			'id',
		),
	);

	export_table($export_table);

?>