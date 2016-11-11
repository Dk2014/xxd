<?php


	$export_table = array(
		'table' 			=> 'events_buy_partner',
		'export_order' 		=> ' ORDER BY `id` DESC',
		'exclude_columns'	=> array(
		),
	);

	export_table($export_table);

?>
