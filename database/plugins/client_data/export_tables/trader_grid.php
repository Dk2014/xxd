<?php


	$export_table = array(
		'table' 			=> 'trader_grid',
		'export_order' 		=> ' ORDER BY `id`',
		'exclude_columns'	=> array('trader_id'),
	);

	export_table($export_table);

?>
