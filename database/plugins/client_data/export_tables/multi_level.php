<?php


	$export_table = array(
		'table' 			=> 'multi_level',
		'export_order' 		=> ' ORDER BY `require_level`',
		'exclude_columns'	=> array(
			'award_lock',
		),
	);

	export_table($export_table);

?>