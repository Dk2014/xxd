<?php


	$export_table = array(
		'table' 			=> 'events_fight_power',
		'export_order' 		=> ' ORDER BY `fight` ASC',
		'exclude_columns'	=> array(
		),
	);

	export_table($export_table);

?>