<?php


	$export_table = array(
		'table' 			=> 'heart_draw_award',
		'export_order' 		=> ' ORDER BY `heart_draw_id`',
		'exclude_columns'	=> array(
			'id',
			'chance',
		),
	);

	export_table($export_table);

?>