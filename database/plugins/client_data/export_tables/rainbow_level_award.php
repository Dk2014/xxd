<?php


	$export_table = array(
		'table' 			=> 'rainbow_level_award',
		'export_order' 		=> ' ORDER BY `mission_level_id`, `order` ASC',
		'exclude_columns'	=> array(),
	);

	export_table($export_table);

?>
