<?php


	$export_table = array(
		'table' 			=> 'role_level',
		'export_order' 		=> ' ORDER BY `level`, `role_id`',
		'exclude_columns'	=> array(
			'id',
			'critial',
			'block',
			'hit',
			'dodge',
			'critial_hurt',
			'sleep',
			'dizziness',
			'random',
			'disable_skill',
			'poisoning',
			'max_power',
			'init_power',
			'sunder_hurt_rate',
			'sunder_end_hurt_rate',
			'sunder_round_num',
			'sunder_dizziness',
		),
	);

	export_table($export_table);

?>