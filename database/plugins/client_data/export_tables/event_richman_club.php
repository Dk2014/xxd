<?php


	$export_table = array(
		'table' 			=> 'events_richman_club_awards',
		'export_order' 		=> ' ORDER BY `require_vip_level` ASC, `require_vip_count` ASC',
		'exclude_columns'	=> array(
			'id',
		),
	);

	export_table($export_table);

?>