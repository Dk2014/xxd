<?php

	$export_table = array(
		'table' => 'battle_pet_exp',
		'export_order' => ' ORDER BY `level` ASC',
		'exclude_columns' => array('min_add_exp', 'max_add_exp'),
	);

	export_table($export_table);

