<?php


$extend_columns = array(
	'award_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
	'award_type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'range' => array('params' => array()),
	),
	'vip_double' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array('否', '是'),
	),
);

function range_award_type(){
	global $all_real_type;
	return $all_real_type;
}
function location($params) {
	global $db;
	$sql = "select count(id) as c from daily_sign_in_award where `type` = {$params['m']}";
	$num = db_query($db, $sql)[0]['c'];
	return "一共{$num}条记录";
}


function sql_where($params) {
	return "where `type`={$params['m']}";
}

function sql_insert($params) {
	return "`type`={$params['m']}";
}

function render_vip_double($column_name, $column_val, $row, $data){
	return $data[$column_val];
}

function editor_vip_double($column_name, $column_val, $row, $data){
	return html_get_select($column_name, $data, $column_val);
}
// 类型
$all_real_type = array(
	1 => '物品',
	2 => '爱心',
	3 => '铜币',
	4 => '元宝',
	5 => '剑心',
	6 => '魂侍',
	//7 => '灵宠',
);

$all_items = get_item();
$all_sword_souls = get_all_sword_soul();
$all_ghosts = get_all_ghost();
$all_pets = array();

function render_award_type($column_name, $column_val, $row, $data) {
	global $all_real_type;
	return $all_real_type[$column_val];
}

function editor_award_type($column_name, $column_val, $row, $data) {
	global $all_real_type;

	$data = $all_real_type;
	$field_name = 'award_type';

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
function get_item() {
	global $db;
	$all_items = array();
	$tmp = db_query($db, "select `id`, `name` from `item` where `type_id`!=10");
	foreach ($tmp as $value)
	{
		$all_items[$value['id']] = $value['name'];
	}

	return $all_items;
}


function all_pet_ball() {
	global $db;
	global $all_enemy;
	$all_pets = array();

	$tmp = db_query($db, "select * from item where `type_id`='11'");
	foreach ($tmp as $value)
	{
		$all_pets[$value['id']] = $value['name'];
	}

	return $all_pets;
}


function editor_award_id($column_name, $column_val, $row, $data){
	$item_name = '';
	$award_id = $row['award_id'];

	if ($row['award_type'] == 5) {
		global $all_sword_souls;
		if (isset($all_sword_souls[$award_id])) {
			$item_name = $all_sword_souls[$award_id];
		}
	} else if ($row['award_type'] == 6) {
		global $all_ghosts;
		if (isset($all_ghosts[$award_id])) {
			$item_name = $all_ghosts[$award_id];
		}
	} else if ($row['award_type'] == 7) {
		global $all_pets;
		if (isset($all_pets[$award_id])) {
			$item_name = $all_pets[$award_id];
		}
	} else {
		global $all_items;
		if (isset($all_items[$award_id])) {
			$item_name = $all_items[$award_id];
		}
	}

	return editor_single_item($item_name, $award_id, "award_id");
}

function render_award_id($column_name, $column_val, $row, $data){
	if ($row['award_type'] == 5) {
		global $all_sword_souls;
		return normal_render($all_sword_souls, $row['award_id']);
	} else if ($row['award_type'] == 6) {
		global $all_ghosts;
		return normal_render($all_ghosts, $row['award_id']);
	} else if ($row['award_type'] == 7) {
		global $all_pets;
		return normal_render($all_pets, $row['award_id']);
	} else if ($row['award_type'] == 1) {
		global $all_items;
		return normal_render($all_items, $row['award_id']);
	} else {
		return normal_render(array('无'), 0);
	}
}

function js_function($params) {
	global $all_items;
	global $all_sword_souls;
	global $all_ghosts;
	global $all_pets;

	$column_name = 'award_id';

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
	
	$html .= get_items_json($all_ghosts, 'ghosts');
	
	$html .= get_items_json($all_pets, 'pets');


	$html .= <<<EOT

	$(".award_type").change(function() {
		switch (parseInt($(this).val())) {
		case 1:
			autoopt_award_id.lookup = items;
			break;
		case 6:
			autoopt_award_id.lookup = ghosts;
			break;
		case 5:
			autoopt_award_id.lookup = sword_souls;
			break;
		case 7:
			autoopt_award_id.lookup = pets;
			break;
		}
		console.log(autoopt_award_id.lookup);
		$(this).parent().parent().find(".display_award_id").autocomplete(autoopt_award_id);
	})

EOT;

return $html;
}

?>
