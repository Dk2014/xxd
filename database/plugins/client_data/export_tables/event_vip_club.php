<?php


	$export_table = array(
		'table' 			=> 'events_vip_club_awards',
		'export_order' 		=> ' ORDER BY `id`',
		'exclude_columns'	=> array(
			'id',
		),
	);

	export_table($export_table);

?>