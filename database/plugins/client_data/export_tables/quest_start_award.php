<?php


	$export_table = array(
		'table' 			=> 'quest_start_award',
		'export_order' 		=> ' ORDER BY `stars` ASC',
		'exclude_columns'	=> array(
			'sign',
		),
	);

	export_table($export_table);

?>
