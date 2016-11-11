<?php


	$export_table = array(
		'table' 			=> 'platform_friend_award',
		'export_order'	=> 'order by `require_friend_num`',
		'exclude_columns'	=> array('id', 'name'),
	);

	export_table($export_table);

?>
