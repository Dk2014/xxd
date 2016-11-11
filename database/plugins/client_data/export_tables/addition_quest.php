<?php


	$export_table = array(
		'table' 			=> 'addition_quest',
		'export_order' 		=> ' order by `serial_number`, `require_lock` ',
		'exclude_columns'	=> array(
			'quest_item_rate',
		),
	);

	export_table($export_table);

?>
