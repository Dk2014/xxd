<?php


	$export_table = array(
		'table' 			=> 'item',
		'export_order' 		=> 'WHERE `type_id` NOT IN (3,4,5,6) ORDER BY `id`',
		'exclude_columns'	=> array(
			//'health',
			//'speed',
			//'attack',
			//'defence',
			//'equip_role_id',
			//'appendix_num',
			//'refine_base',
		),
	);

	export_table($export_table);

?>
