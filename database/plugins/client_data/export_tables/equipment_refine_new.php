<?php


	$export_table = array(
		'table' 			=> 'equipment_refine_new',
		'export_order' 		=> ' ORDER BY `id` ASC',
		'exclude_columns'	=> array(
			'id',
		),
	);

	export_table($export_table);

?>