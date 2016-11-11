<?php

	$export_table = array(
		'table' => 'shaded_mission',
		'export_order' => ' ORDER BY `mission_level_id`, `id` asc',
		'exclude_columns' => array(),
		'addition' => array(
			'shaded_mission_item1_is_equipment' => 'int, item1是否为装备',
			'shaded_mission_item2_is_equipment' => 'int, item2是否为装备',
			'shaded_mission_item3_is_equipment' => 'int, item3是否为装备',
		),
	);

	// additional columns' common logic
	function get_addition_shaded_mission_item_is_equipment($db, $item_column) {
		$result = array();
		$rows = db_query($db, "

SELECT `shaded_mission`.*
	,`item`.`type_id` IS NOT NULL
	AND `item`.`type_id` IN (
		3
		,4
		,5
		,6
		) `is_equipment`
FROM `shaded_mission`
LEFT JOIN `item` ON `shaded_mission`.`{$item_column}` = `item`.`id`
ORDER BY `shaded_mission`.`mission_level_id`
	,`shaded_mission`.`id` ASC;

		");
		foreach($rows as $row) {
			$result[] = $row['is_equipment'];
		}
		return $result;
	}

	function get_addition_shaded_mission_item1_is_equipment($db) {
		return get_addition_shaded_mission_item_is_equipment($db, 'award_item1');
	}
	function get_addition_shaded_mission_item2_is_equipment($db) {
		return get_addition_shaded_mission_item_is_equipment($db, 'award_item2');
	}
	function get_addition_shaded_mission_item3_is_equipment($db) {
		return get_addition_shaded_mission_item_is_equipment($db, 'award_item3');
	}

	export_table($export_table);

