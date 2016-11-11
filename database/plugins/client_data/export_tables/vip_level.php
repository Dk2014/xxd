<?php
	$export_table = array(
		'table' 			=> 'vip_level',
		'export_order' 		=> ' ORDER BY `level`',
		'exclude_columns'	=> array(),
	);

	export_table($export_table);
?>
