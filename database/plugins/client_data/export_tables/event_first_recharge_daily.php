<?php

	$export_table = array(
		'table' 			=> 'event_first_recharge_daily',
		'export_order' 		=> ' ORDER BY `require_day` ASC',
		'exclude_columns'	=> array(
			'id',
		),
	);

	export_table($export_table);

?>