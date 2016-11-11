<?php


	$export_table = array(
		'table' 			=> 'battle_pet',
		'export_order' 		=> ' ORDER BY `pet_id`',
		'exclude_columns'	=> array(
			'round_attack',
		),
	);

	export_table($export_table);

?>