<?php

$extend_columns = array(
	/*'字段' => 配置 */

	'item_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
	
	'type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
);

function sql_where($params) {
	if (!isset($params['m'])){
		return '';
	}

	return " where `chest_id` = {$params['m']}";
}

function sql_insert($params) {
	return "`chest_id` = {$params['m']}";
}

// 类型
$all_type = array(
	0 => '无',
	1 => '铜钱',
	2 => '元宝',
	3 => '物品（喜好品魂侍碎片不要选）',
	4 => '魂侍',
	5 => '剑心',
	6 => '灵宠',
	7 => '魂侍碎片',
	8 => '喜好品',
	9 => '装备',
);

function render_type($column_name, $column_val, $row, $data) {
	global $all_type;
	return $all_type[$column_val];
}

function editor_type($column_name, $column_val, $row, $data) {
	global $all_type;

	$data = $all_type;
	$field_name = 'type';

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


function get_all_ghost_fragment() {
	global $db;
	$all_ghost_fragment = array();

	$tmp = db_query($db, "select `id`, `name` from `item` where type_id = 13");
	foreach ($tmp as $value)
	{
		$all_ghost_fragment[$value['id']] = $value['name'];
	}

	return $all_ghost_fragment;
}

$all_perference = get_all_item_by_types(18); //喜好品
$all_items = get_all_item();
$all_equipments = get_all_item_by_types(array(3,4,5,6));
$all_sword_souls = get_all_sword_soul();
$all_ghosts = get_all_ghost();
$all_ghost_fragments = get_all_ghost_fragment();

$all_pets = get_all_pet();

function editor_item_id($column_name, $column_val, $row, $data){
	$item_name = '';
	$item_id = $row['item_id'];

	if ($row['type'] == 5) {
		global $all_sword_souls;
		if (isset($all_sword_souls[$item_id])) {
			$item_name = $all_sword_souls[$item_id];
		}
	} else if ($row['type'] == 4) {
		global $all_ghosts;
		if (isset($all_ghosts[$item_id])) {
			$item_name = $all_ghosts[$item_id];
		}
	} else if ($row['type'] == 6) {
		global $all_pets;
		if (isset($all_pets[$item_id])) {
			$item_name = $all_pets[$item_id];
		}
	} else if ($row['type'] == 7) {
		global $all_ghost_fragments;
		if (isset($all_ghost_fragments[$item_id])) {
			$item_name = $all_ghost_fragments[$item_id];
		}
	} else if ($row['type'] == 8) {
		global $all_perference;
		if (isset($all_perference[$item_id])) {
			$item_name = $all_perference[$item_id];
		}
	} else if ($row['type'] == 9) {
		global $all_equipments;
		if (isset($all_equipments[$item_id])) {
			$item_name = $all_equipments[$item_id];
		}
	} else {
		global $all_items;
		if (isset($all_items[$item_id])) {
			$item_name = $all_items[$item_id];
		}
	}

	return editor_single_item($item_name, $item_id, "item_id");
}

function render_item_id($column_name, $column_val, $row, $data){
	if ($row['type'] == 5) {
		global $all_sword_souls;
		return normal_render($all_sword_souls, $row['item_id']);
	} else if ($row['type'] == 4) {
		global $all_ghosts;
		return normal_render($all_ghosts, $row['item_id']);
	} else if ($row['type'] == 6) {
		global $all_pets;
		return normal_render($all_pets, $row['item_id']);
	} else {
		global $all_items;
		return normal_render($all_items, $row['item_id']);
	}
}

function js_function($params) {
	global $all_items;
	global $all_sword_souls;
	global $all_ghosts;
	global $all_pets;
	global $all_perference;
	global $all_ghost_fragments;
	global $all_equipments;

	$column_name = 'item_id';

	$html = get_items_json($all_items, 'items');

	$html .= <<<EOT

	autoopt_{$column_name} = {
		lookup: items, 
		minChars: 0,
		onSelect: function(s) {
			$(this).siblings('input[class="real_value"]').val(s.id);
		}
	};

	$(".display_{$column_name}").autocomplete(autoopt_{$column_name});

EOT;

	$html .= get_items_json($all_sword_souls, 'sword_souls');

	$html .= get_items_json($all_equipments, 'equipments');

	$html .= get_items_json($all_perference, 'perference');
	
	$html .= get_items_json($all_ghosts, 'ghosts');
	
	$html .= get_items_json($all_pets, 'pets');

	$html .= get_items_json($all_ghost_fragments, 'ghost_fragments');

	$html .= <<<EOT

	$(".type").change(function() {
		switch (parseInt($(this).val())) {
		case 3:
			autoopt_item_id.lookup = items;
			break;
		case 4:
			autoopt_item_id.lookup = ghosts;
			break;
		case 5:
			autoopt_item_id.lookup = sword_souls;
			break;
		case 6:
			autoopt_item_id.lookup = pets;
			break;
		case 7:
			autoopt_item_id.lookup = ghost_fragments;
			break;
		case 8:
			autoopt_item_id.lookup = perference;
			break;
		case 9:
			autoopt_item_id.lookup = equipments;
			break;
		default:
			autoopt_item_id.lookup = [{value:'无', id:'0'}];
		}
		$(this).parent().parent().find(".display_item_id").autocomplete(autoopt_item_id);
	})

EOT;

return $html;
}

?>
