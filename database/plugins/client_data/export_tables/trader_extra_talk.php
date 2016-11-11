<?php


	$export_table = array(
		'table' 			=> 'trader_extra_talk',
		'export_order' 		=> ' ORDER BY `trader_id`',
		'exclude_columns'	=> array('id'),
	);

	export_table($export_table);

?>
