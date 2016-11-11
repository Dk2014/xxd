<?php
//include "item_links.php";
$extend_columns = array(
	'item_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),

	'origin_key' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
	'origin_type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
);

$all_items = get_all_item();
function editor_item_id($column_name, $column_val, $row, $data){
	global $all_items;

	$item_name = '';
	if (isset($all_items[$column_val])) {
		$item_name = $all_items[$column_val];
	}
	return editor_single_item($item_name, $column_val, $column_name);
}

function render_item_id($column_name, $column_val, $row, $data){
	global $all_items;
	return (isset($all_items[$column_val]) ? $all_items[$column_val] : "0");
}

$resource_origin_type = array(
	0 => '主线关卡',
	1 => '深渊关卡',
	2 => '彩虹关卡',
	3 => '灵宠幻境',
	4 => '魂侍关卡',
	5 => '灵宠关卡',
	6 => '伙伴关卡',
	7 => '多人关卡',
	8 => '经验关卡',
	9 => '铜钱关卡',
);



$all_mission_level = get_all_level_by_type(0);
$all_mission_level[-1] = '最近普通关卡';
$all_hard_level = get_all_level_by_type(8);
$all_buddy_level = get_all_level_by_type(9);
$all_battle_pet_level = get_all_level_by_type(10);
$all_ghost_level = get_all_level_by_type(11);
$all_pve_level = get_all_level_by_type(13);
//$all_rainbow_level = get_all_rainbow_level();
$all_rainbow_level = get_all_rainbow_segment();
$all_multi_level = get_all_multi_level();

function render_origin_type($column_name, $column_val, $row, $data) {
	global $resource_origin_type;
	return $resource_origin_type[$column_val];
}

function editor_origin_type($column_name, $column_val, $row, $data) {
	global $resource_origin_type;
	$data = $resource_origin_type;

	$field_name = 'origin_type';

	$html = "<select class={$field_name} name={$field_name}[]>";

	foreach ($data as $key => $value) {
		if (isset($row[$field_name]) && $row[$field_name] == $key) {
			$html .= '<option value="'.$key.'" selected="selected" >'.$value.'</option>';
		} else {
			$html .= '<option value="'.$key.'">'.$value.'</option>';
		}
	}

	$html .= '</select>';
	return $html;
}


function editor_origin_key($column_name, $column_val, $row, $data){
	$item_name = '';
	$origin_key = $row['origin_key'];

	if ($row['origin_type'] == 0) {
		global $all_mission_level;
		if (isset($all_mission_level[$origin_key])) {
			$item_name = $all_mission_level[$origin_key];
		}
	} else if ($row['origin_type'] == 1) {
		global $all_hard_level;
		if (isset($all_hard_level[$origin_key])) {
			$item_name = $all_hard_level[$origin_key];
		}
	} else if ($row['origin_type'] == 2) {
		global $all_rainbow_level;
		if (isset($all_rainbow_level[$origin_key])) {
			$item_name = $all_rainbow_level[$origin_key];
		}
	} else if ($row['origin_type'] == 3) {
		global $all_pve_level;
		if (isset($all_pve_level[$origin_key])) {
			$item_name = $all_pve_level[$origin_key];
		}
	} else if ($row['origin_key'] == 4) {
		global $all_ghost_level; 
		if (isset($all_ghost_level)) {
			$item_name = $all_ghost_level[$origin_key];
		}
	} else if ($row['origin_key'] == 5) {
		global $all_battle_pet_level;
		if (isset($all_battle_pet_level)) {
			$item_name = $all_battle_pet_level[$origin_key];
		}
	} else if ($row['origin_key'] == 6) {
		global $all_buddy_level;
		if (isset($all_buddy_level[$origin_key])) {
			$item_name = $all_buddy_level[$origin_key];
		}
	} else if ($row['origin_key'] == 7) {
		global $all_multi_level;
		if (isset($all_multi_level)) {
			$item_name = $all_multi_level[$origin_key];
		}
	}

	return editor_single_item($item_name, $origin_key, "origin_key");
}

function render_origin_key($column_name, $column_val, $row, $data){
	if ($row['origin_type'] == 0) {
		global $all_mission_level;
		return normal_render($all_mission_level, $row['origin_key']);
	} else if ($row['origin_type'] == 1) {
		global $all_hard_level;
		return normal_render($all_hard_level, $row['origin_key']);
	} else if ($row['origin_type'] == 2) {
		global $all_rainbow_level;
		return normal_render($all_rainbow_level, $row['origin_key']);
	} else if ($row['origin_type'] == 3) {
		global $all_pve_level;
		return normal_render($all_pve_level, $row['origin_key']);
	} else if ($row['origin_type'] == 4) {
		global $all_ghost_level;
		return normal_render($all_ghost_level, $row['origin_key']);
	} else if ($row['origin_type'] == 5) {
		global $all_battle_pet_level;
		return normal_render($all_battle_pet_level, $row['origin_key']);
	} else if ($row['origin_type'] == 6) {
		global $all_buddy_level;
		return normal_render($all_buddy_level, $row['origin_key']);
	} else if ($row['origin_type'] == 7) {
		global $all_multi_level;
		return normal_render($all_multi_level, $row['origin_key']);
	}
}

function js_function($params) {
	global $all_mission_level; 
	global $all_hard_level; 
	global $all_buddy_level;
	global $all_battle_pet_level;
	global $all_ghost_level; 
	global $all_pve_level;
	global $all_rainbow_level;
	global $all_multi_level;


	$column_name = 'origin_key';

	$html = '';
	$html .= get_items_json($all_mission_level, 'all_mission_level');
	$html .= get_items_json($all_hard_level, 'all_hard_level');
	$html .= get_items_json($all_buddy_level, 'all_buddy_level');
	$html .= get_items_json($all_battle_pet_level, 'all_battle_pet_level');
	$html .= get_items_json($all_ghost_level, 'all_ghost_level');
	$html .= get_items_json($all_pve_level, 'all_pve_level');
	$html .= get_items_json($all_rainbow_level, 'all_rainbow_level');
	$html .= get_items_json($all_multi_level, 'all_multi_level');

	$html .= <<<EOT

	autoopt_{$column_name} = {
		lookup: all_hard_level, 
		minChars: 0,
		onSelect: function(s) {
			$(this).siblings('input[class="real_value"]').val(s.id);
		}
	};

	$(".display_{$column_name}").autocomplete(autoopt_{$column_name});

EOT;


	$nothing = array('无');
	$html .= get_items_json($nothing, 'nothing');



	$html .= <<<EOT

	$(".origin_type").change(function() {
		switch (parseInt($(this).val())) {
		case 0:
			autoopt_origin_key.lookup = all_mission_level;
			break;
		case 1:
			autoopt_origin_key.lookup = all_hard_level;
			break;
		case 2:
			autoopt_origin_key.lookup = all_rainbow_level;
			break;
		case 3:
			autoopt_origin_key.lookup = all_pve_level;
			break;
		case 4:
			autoopt_origin_key.lookup = all_ghost_level;
			break;
		case 5:
			autoopt_origin_key.lookup = all_ghost_level;
			break;
		case 6:
			autoopt_origin_key.lookup = all_battle_pet_level;
			break;
		case 7:
			autoopt_origin_key.lookup = all_multi_level;
			break;
		 default:
			autoopt_origin_key.lookup = nothing;
			break;
		}
		console.log(autoopt_origin_key.lookup);
		$(this).parent().parent().find(".display_origin_key").autocomplete(autoopt_origin_key);
	})

EOT;
	global $all_items;
	$html .= choose_single_item($all_items, 'item_id');

	return $html;
}

?>
