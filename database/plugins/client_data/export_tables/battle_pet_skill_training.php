<?php

	$export_table = array(
		'table' => 'battle_pet_skill_training',
		'export_order' => ' ORDER BY `level` ASC',
		'exclude_columns' => array(),
	);

	export_table($export_table);

