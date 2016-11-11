<?php

	$export_table = array(
		'table' 			=> 'faq',
		'export_order' 		=> ' ORDER BY `order` asc',
		'exclude_columns'	=> array(
		),
	);

	export_table($export_table);

?>