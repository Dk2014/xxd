<?php
	$export_table = array(
		'table' 			=> 'vip_levelup_gift',
		'export_order' 		=> ' ORDER BY `vip_level`',
		'exclude_columns'	=> array(
			'id',
		),
	);

	export_table($export_table);
?>
