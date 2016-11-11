<?php


	$export_table = array(
		'table' 			=> 'ghost_train_price',
		'export_order' 		=> ' ORDER BY `times` ASC',
		'exclude_columns'	=> array(
			'id',
		),
	);

	export_table($export_table);

?>
