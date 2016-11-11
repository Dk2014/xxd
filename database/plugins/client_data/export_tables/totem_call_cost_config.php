<?php


	$export_table = array(
		'table' 			=> 'totem_call_cost_config',
		'export_order' 		=> ' ORDER BY `times`',
		'exclude_columns'	=> array('id'),
	);

	export_table($export_table);

?>
