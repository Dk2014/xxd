<?php
	$export_table = array(
		'table' 			=> 'trader_refresh_price',
		'export_order' 		=> ' ORDER BY `trader_id`',
		'exclude_columns'	=> array('id'),
	);

	export_table($export_table);

?>
