<?php


	$export_table = array(
		'table' 			=> 'recruit_buddy',
		'export_order' 		=> ' ORDER BY `id`',
		'exclude_columns'	=> array(),
		'addition' => array(
			'required_main_quest_order' => 'int,以来主线任务顺序',
		)
	);
function get_addition_required_main_quest_order($db) {
	$result = array();
	$rows = db_query($db, " select q.order as `order` from extend_quest eq  left join quest q on eq.required_quest=q.id");
	$rows = db_query($db, " select * from recruit_buddy eq  left join quest q on quest_id=q.id");

	foreach($rows as $row) {
		if(isset($row['order']) && $row['order']!=null) {
			$result[] = $row['order'];
		} else {
			$result[] = 0;
		}
	}
	return $result;
}


	export_table($export_table);

?>
