<?php


	$export_table = array(
		'table' 			=> 'skill_content',
		'export_order' 		=> ' ORDER BY `skill_id`',
		'exclude_columns'	=> array(
			'id'
		),
	);

	export_table($export_table);

?>