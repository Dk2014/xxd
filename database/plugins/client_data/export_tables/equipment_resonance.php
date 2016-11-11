<?php

	$export_table = array(
		'table' => 'equipment_resonance',
		'export_order' => ' ORDER BY  `require_level` ASC',
		'exclude_columns' => array(),
	);

	export_table($export_table);

