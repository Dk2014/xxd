<?php

	$export_table = array(
		'table' 			=> 'events_total_recharge_awards',
		'export_order' 		=> ' ORDER BY `require_total` ASC',
		'exclude_columns'	=> array(
			'id',
		),
	);

	export_table($export_table);

?>