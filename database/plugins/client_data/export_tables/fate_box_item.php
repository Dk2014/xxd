<?php


	$export_table = array(
		'filename' => 'fate_box_item',
		'table' 			=> 'chest_item',
		'export_order' 		=> ' WHERE `chest_id` > 0 ORDER BY `chest_id`, `id` ',
		'exclude_columns'	=> array(
			'item_num',
		),
	);

	export_table($export_table);

?>
