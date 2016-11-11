<?php


	$export_table = array(
		'table' 			=> 'enemy_deploy_form',
		'export_order' 		=> ' ORDER BY `battle_type`, `parent_id`',
		'exclude_columns'	=> array(
			'id',
		),
	);

	export_table($export_table);

?>