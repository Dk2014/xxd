<?php


	$export_table = array(
		'table' 			=> 'events_month_card_awards',
		'export_order' 		=> ' ORDER BY `id` ASC',
		'exclude_columns'	=> array(
			'id',
		),
	);

	export_table($export_table);

?>