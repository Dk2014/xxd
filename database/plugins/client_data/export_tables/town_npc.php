<?php


	$export_table = array(
		'table' 			=> 'town_npc',
		'export_order' 		=> ' ORDER BY `town_id`',
		'exclude_columns'	=> array(),
	);

	export_table($export_table);

?>