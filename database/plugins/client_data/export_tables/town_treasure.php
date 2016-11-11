<?php

	$export_table = array(
		'table'            => 'town_treasure',
		'export_order'     => ' ORDER BY `town_id`, `id` asc',
		'exclude_columns'  => array(),
		'addition'         => array(
			'town_treasure_item1_is_equipment' => 'int, item1是否为装备',
			'town_treasure_item2_is_equipment' => 'int, item2是否为装备',
			'town_treasure_item3_is_equipment' => 'int, item3是否为装备',
			'town_treasure_item4_is_equipment' => 'int, item4是否为装备',
		),
	);

	// additional columns' common logic
	function get_addition_town_treasure_item_is_equipment($db, $item_column) {
		$result = array();
		$rows = db_query($db, "

SELECT `town_treasure`.*
	,`item`.`type_id` IS NOT NULL
	AND `item`.`type_id` IN (
		3
		,4
		,5
		,6
		) `is_equipment`
FROM `town_treasure`
LEFT JOIN `item` ON `town_treasure`.`{$item_column}` = `item`.`id`
ORDER BY `town_treasure`.`town_id`
	,`town_treasure`.`id` ASC;

		");
		foreach($rows as $row) {
			$result[] = $row['is_equipment'];
		}
		return $result;
	}

	function get_addition_town_treasure_item1_is_equipment($db) {
		return get_addition_town_treasure_item_is_equipment($db, 'award_item1_id');
	}
	function get_addition_town_treasure_item2_is_equipment($db) {
		return get_addition_town_treasure_item_is_equipment($db, 'award_item2_id');
	}
	function get_addition_town_treasure_item3_is_equipment($db) {
		return get_addition_town_treasure_item_is_equipment($db, 'award_item3_id');
	}
	function get_addition_town_treasure_item4_is_equipment($db) {
		return get_addition_town_treasure_item_is_equipment($db, 'award_item4_id');
	}

	export_table($export_table);

