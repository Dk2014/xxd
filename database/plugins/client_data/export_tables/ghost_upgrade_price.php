<?php


	$export_table = array(
		'table' 			=> 'ghost_upgrade_price',
		'export_order' 		=> ' ORDER BY `quality` ASC',
		'exclude_columns'	=> array(
			'id',
		),
	);

	export_table($export_table);

?>
