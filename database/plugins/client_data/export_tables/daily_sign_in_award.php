<?php


	$export_table = array(
		'table' 			=> 'daily_sign_in_award',
		'export_order' 		=> ' ORDER BY `type`, `id` asc',
		'exclude_columns'	=> array(),
	);

	export_table($export_table);

?>
