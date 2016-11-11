<?php
	$export_table = array(
		'table' 			=> 'vip_privilege_config',
		'export_order' 		=> ' ORDER BY `privilege_id`, `level`',
		'exclude_columns'	=> array('id'),
	);

	export_table($export_table);
?>
