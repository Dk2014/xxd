<?php

	$export_table = array(
		'table' => 'battle_pet_level_info',
		'export_order' => ' ORDER BY `pet_id`, `level` ASC',
		'exclude_columns' => array(),
	);

	export_table($export_table);

