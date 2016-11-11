<?php


	$export_table = array(
		'table' 			=> 'func',
		'export_order' 		=> ' ORDER BY `lock`, `level`',
		'exclude_columns'	=> array(
			'unique_key',
		),
	);

	export_table($export_table);

?>
