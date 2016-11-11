<?php


	$export_table = array(
		'table' 			=> 'arena_buy_cost_config',
		'export_order' 		=> ' ORDER BY `times` ASC',
		'exclude_columns'	=> array(
			'id',
		),
	);

	export_table($export_table);

?>