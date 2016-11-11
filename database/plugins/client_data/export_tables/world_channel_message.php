<?php


	$export_table = array(
		'table' 			=> 'world_channel_message',
		'export_order' 		=> ' ORDER BY `id`',
		'exclude_columns'	=> array(
			'sign',
			'desc',
		),
	);

	export_table($export_table);

?>
