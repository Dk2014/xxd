<?php

	$export_table = array(
		'table' => 'role_friendship',
		'export_order' => ' ORDER BY `role_id` ASC, `friendship_level` ASC',
		'exclude_columns' => array(),
		'addition' => array(
			'role_friendship_name_postfix' => 'int, 名称表现后缀值',
		)
	);

	function get_addition_role_friendship_name_postfix($db) {
		$result = array();
		$rows = db_query($db, 'select `level_color` from `role_friendship` order by `role_id` asc, `friendship_level` asc');
		$previous_color = '';
		$same_color_count = 0;
		foreach($rows as $row) {
			$current_color = $row['level_color'];
			if($current_color == $previous_color) {
				$same_color_count++;
				$result[] = $same_color_count;
			} else {
				$same_color_count=0;
				$result[] = $same_color_count;
				$previous_color = $current_color;
			}
		}
		return $result;
	}

	export_table($export_table);

