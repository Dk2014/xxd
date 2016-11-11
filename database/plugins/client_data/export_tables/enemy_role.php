<?php


	$export_table = array(
		'table' 			=> 'enemy_role',
		'export_order' 		=> ' ORDER BY `id`',
		'exclude_columns'	=> array(
			'cultivation',
			'dodge',
			'hit',
			'block',
			'critial',
			'toughness',
			'destroy',
			'critial_hurt',
			'sleep',
			'dizziness',
			'random',
			'disable_skill',
			'poisoning',
			'skill_id',
			'skill_force',
			'sunder_max_value',
			'sunder_min_hurt_rate',
			'sunder_end_hurt_rate',
			'sunder_attack',
			'skill_wait',
			'release_num',
			'recover_round_num',
			'common_attack_num',
			'skill_config',
		),
	);

	export_table($export_table);

?>