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
);

function range_award_type(){
	global $all_real_type;
	return $all_real_type;
}

function sql_where($params) {
	return " order by `requrie_active_days` ASC";
}



// 类型
$all_real_type = array(
	1 => '物品',
	2 => '爱心',
	3 => '铜币',
	4 => '元宝',
	5 => '剑心',
	6 => '魂侍',
);

$all_items = get_item();
$all_sword_souls = get_all_sword_soul();
$all_ghosts = get_all_ghost();

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
	$tmp = db_query($db, "select `id`, `name` from `item` where `type_id`!=10 and `type_id`!=11");
	foreach ($tmp as $value)
	{
		$all_items[$value['id']] = $value['name'];
	}

	return $all_items;
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
		}
		console.log(autoopt_award_id.lookup);
		$(this).parent().parent().find(".display_award_id").autocomplete(autoopt_award_id);
	})

EOT;

return $html;
}

?>
