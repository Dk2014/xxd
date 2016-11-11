<?php


	$export_table = array(
		'table' 			=> 'buy_resource_times_config',
		'export_order' 		=> ' ORDER BY `times` ASC',
		'exclude_columns'	=> array(
			'id',
		),
	);

	export_table($export_table);

?>
