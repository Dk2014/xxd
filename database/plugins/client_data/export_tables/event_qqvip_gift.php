<?php


	$export_table = array(
		'table' 			=> 'events_qqvip_gift_awards',
		'export_order' 		=> ' ORDER BY `require_login_days`',
		'exclude_columns'	=> array(
			'id',
		),
	);

	export_table($export_table);

?>