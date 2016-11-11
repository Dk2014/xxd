<?php


	$export_table = array(
		'table' 			=> 'events_arena_rank_awards',
		'export_order' 		=> ' ORDER BY `require_arena_rank` ASC',
		'exclude_columns'	=> array(
		),
	);

	export_table($export_table);

?>
