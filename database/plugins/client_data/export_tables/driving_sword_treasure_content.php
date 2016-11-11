<?php

	$export_table = array(
		'table' => 'driving_sword_treasure_content',
		'export_order' => ' ORDER BY `treasure_id`, `order`',
		'exclude_columns' => array(),
		'addition' => array(
			'driving_sword_treasure_content_item_is_equipment' => 'int, item是否为装备',
		),
	);

	export_table($export_table);

	function get_addition_driving_sword_treasure_content_item_is_equipment($db) {
		$result = array();
		$rows = db_query($db, "

SELECT `driving_sword_treasure_content`.*
	,`item`.`type_id` IS NOT NULL
	AND `item`.`type_id` IN (
		3
		,4
		,5
		,6
		) `is_equipment`
FROM `driving_sword_treasure_content`
LEFT JOIN `item` ON `driving_sword_treasure_content`.`award_item` = `item`.`id`
ORDER BY `driving_sword_treasure_content`.`treasure_id`
	,`driving_sword_treasure_content`.`order`;

		");
		foreach($rows as $row) {
			$result[] = $row['is_equipment'];
		}
		return $result;
	}


