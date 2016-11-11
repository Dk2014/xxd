<?php


	$export_table = array(
		'filename' => 'fate_box_config',
		'table' 			=> 'chest',
		'export_order' 		=> ' WHERE `type`>0 ORDER BY `id`,`type` ',
		'exclude_columns'	=> array(
			'quality',
			'probability',
		),
	);

	export_table($export_table);

?>
