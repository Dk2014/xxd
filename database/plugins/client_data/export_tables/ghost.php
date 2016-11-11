<?php


	$export_table = array(
		'table' 			=> 'ghost',
		'export_order' 		=> ' ORDER BY `id`',
		'exclude_columns'	=> array(
			'role_id',
		),
	);

	export_table($export_table);

?>
