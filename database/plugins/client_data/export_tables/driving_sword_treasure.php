<?php

	$export_table = array(
		'table' => 'driving_sword_treasure',
		'export_order' => ' ORDER BY `cloud_id`, `event_id`',
		'exclude_columns' => array(),
	);

	export_table($export_table);

