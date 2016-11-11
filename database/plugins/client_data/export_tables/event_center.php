<?php


	$export_table = array(
		'table' 			=> 'quest_activity_center',
		'export_order' 		=> ' ORDER BY `id` ASC',
		'exclude_columns'	=> array(
			'is_mail',
			'mail_title',
			'mail_content',
		),
	);

	export_table($export_table);

?>