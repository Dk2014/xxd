<?php


	$export_table = array(
		'table' 			=> 'mission_level_recovery_meng_yao',
		'export_order' 		=> ' ORDER BY `mission_level_id`',
		'exclude_columns'	=> array(
			'probability',
		),
	);

	export_table($export_table);

?>