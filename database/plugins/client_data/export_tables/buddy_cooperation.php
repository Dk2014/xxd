<?php
	$export_table = array(
		'table' 			=> 'buddy_cooperation',
		'export_order' 		=> ' order by `id` ',
		'exclude_columns'	=> array(
			//'health',
			//'attack',
			//'defence',
			//'skill_level',
			//'speed',
			//'cultivation',
			//'sunder',
			//'dodge_level',
			//'hit_level',
			//'block_level',
			//'tenacity_level',
			//'destroy_level',
			//'critical_level',
			//'critical_hurt_level',
			//'ghost_power',
		),
	);

	export_table($export_table);

?>
