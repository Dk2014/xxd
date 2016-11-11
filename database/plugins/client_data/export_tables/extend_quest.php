<?php


	$export_table = array(
		'table' 			=> 'extend_quest',
		'export_order' 		=> ' ORDER BY `id`',
		'exclude_columns'	=> array(
			'related_npc',
			'required_quest',
		),
	);

	export_table($export_table);

?>
