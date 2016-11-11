<?php

	$export_table = array(
		'table' => 'driving_sword',
		'export_order' => ' ORDER BY `level`',
		'exclude_columns' => array('hole_x', 'hole_y', 'obstacle_count'),
		'addition' => array(
			'driving_sword_teleport_count' => 'int, 传送点总数',
		),
	);

	export_table($export_table);

	function get_addition_driving_sword_teleport_count($db) {
		$result = array();
		$rows = db_query($db, "
SELECT *
FROM `driving_sword`
LEFT JOIN (
	SELECT `cloud_id`
		,count(`cloud_id`) `teleport_count`
	FROM `driving_sword_teleport`
	) `teleports` ON `driving_sword`.`id` = `teleports`.`cloud_id`
ORDER BY `driving_sword`.`level`;
		");
		foreach($rows as $row) {
			if (isset($row['teleport_count']))
				$result[] = $row['teleport_count'];
			else
				$result[] = 0;
		}
		return $result;
	}

