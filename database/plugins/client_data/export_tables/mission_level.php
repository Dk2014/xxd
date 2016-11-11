<?php


	$export_table = array(
		'table' 			=> 'mission_level',
		'export_order' 		=> ' ORDER BY `mission_id`, `lock` asc',
		'exclude_columns'	=> array(),
	);
    
	export_table($export_table);

?>
