<?php

	$export_table = array(
		'table' => 'driving_sword_exploring',
		'export_order' => ' ORDER BY `cloud_id`, `event_id`',
		'exclude_columns' => array(),
		'addition' => array(
			'driving_sword_exploring_item1_is_equipment' => 'int, item1是否为装备',
			'driving_sword_exploring_item2_is_equipment' => 'int, item2是否为装备',
			'driving_sword_exploring_item3_is_equipment' => 'int, item3是否为装备',
			'driving_sword_exploring_garrison_item_is_equipment' => 'int, 驻守奖励是否为装备',
		),
	);

	export_table($export_table);

	// additional columns' common logic
	function get_addition_driving_sword_exploring_item_is_equipment($db, $item_column) {
		$result = array();
		$rows = db_query($db, "

SELECT `driving_sword_exploring`.*
	,`item`.`type_id` IS NOT NULL
	AND `item`.`type_id` IN (
		3
		,4
		,5
		,6
		) `is_equipment`
FROM `driving_sword_exploring`
LEFT JOIN `item` ON `driving_sword_exploring`.`{$item_column}` = `item`.`id`
ORDER BY `driving_sword_exploring`.`cloud_id`
	,`driving_sword_exploring`.`event_id`;

		");
		foreach($rows as $row) {
			$result[] = $row['is_equipment'];
		}
		return $result;
	}

	function get_addition_driving_sword_exploring_item1_is_equipment($db) {
		return get_addition_driving_sword_exploring_item_is_equipment($db, 'award_item1');
	}
	function get_addition_driving_sword_exploring_item2_is_equipment($db) {
		return get_addition_driving_sword_exploring_item_is_equipment($db, 'award_item2');
	}
	function get_addition_driving_sword_exploring_item3_is_equipment($db) {
		return get_addition_driving_sword_exploring_item_is_equipment($db, 'award_item3');
	}

	function get_addition_driving_sword_exploring_garrison_item_is_equipment($db) {
		return get_addition_driving_sword_exploring_item_is_equipment($db, 'garrison_item');
	}

