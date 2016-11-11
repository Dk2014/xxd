<?php
//include "item_links.php";
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
	//'name' => array(
	//	'editor' => array('params' => array()),
	//	'render' => array('params' => array()),
	//	'data' => array(),
	//),
);



// 类型
$all_real_type = array(
	0 => '物品',
	1 => '铜币',
	2 => '元宝',
	3 => '爱心',
	4 => '上阵角色经验',
	//5 => '单个角色经验',
	6 => '魂侍',
	7 => '剑心',
	//8 => '体力',
);


$all_items = get_item();
$all_sword_souls = get_all_sword_soul();
$all_ghosts = get_all_ghost();
$all_pets = get_all_pet();
//function render_name($column_name, $column_val, $row, $data) {
//	return $column_val;
//}
//
//function editor_name($column_name, $column_val, $row, $data) {
//	$field_name = 'name';
//
//	$html = "<input class=\"{$field_name}\" name=\"{$field_name}[]\" value=\"{$column_val}\">";
//	return $html;
//}

function range_award_type(){
	global $all_real_type;
	return $all_real_type;
}

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
	$tmp = db_query($db, "select `id`, `name` from `item` where `type_id`!='10'");
	foreach ($tmp as $value)
	{
		$all_items[$value['id']] = $value['name'];
	}

	return $all_items;
}



function editor_award_id($column_name, $column_val, $row, $data){
	$item_name = '';
	$award_id = $row['award_id'];

	if ($row['award_type'] == 7) {
		global $all_sword_souls;
		if (isset($all_sword_souls[$award_id])) {
			$item_name = $all_sword_souls[$award_id];
		}
	} else if ($row['award_type'] == 6) {
		global $all_ghosts;
		if (isset($all_ghosts[$award_id])) {
			$item_name = $all_ghosts[$award_id];
		}
	} else if ($row['award_type'] == -1) {
		//不使用
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
	if ($row['award_type'] == 7) {
		global $all_sword_souls;
		return normal_render($all_sword_souls, $row['award_id']);
	} else if ($row['award_type'] == 6) {
		global $all_ghosts;
		return normal_render($all_ghosts, $row['award_id']);
	} else if ($row['award_type'] == -1) {
		global $all_pets;
		return normal_render($all_pets, $row['award_id']);
	} else {
		global $all_items;
		return normal_render($all_items, $row['award_id']);
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
			var award_type = parseInt($(this).parent().parent().find(".award_type").val());
			var award_id = parseInt($(this).siblings('input[class="real_value"]').val());
			var award_name = s['value'];
			switch (award_type) {
			case 1:
				award_name = '铜币';
				break;
			case 2:
				award_name = '元宝';
				break;
			case 3:
				award_name = '爱心';
				break;
			case 4:
				award_name = '上阵角色经验';
				break;
			case 6:
				award_name = '魂侍' + award_name;
				break;
			case 7:
				award_name = '剑心' + award_name;
				break;
			}
			$(this).parent().parent().find('input[name="name[]"]').val(award_name);

			console.log(award_id);
			console.log(award_name);
			console.log(award_type);
			console.log(autoopt_award_id.lookup);
		}
	};

	$(".display_{$column_name}").autocomplete(autoopt_{$column_name});

EOT;

	$html .= get_items_json($all_sword_souls, 'sword_souls');

	$html .= get_items_json($all_ghosts, 'ghosts');

	$html .= get_items_json($all_pets, 'pets');

	$nothing = array('无');
	$html .= get_items_json($nothing, 'nothing');



	$html .= <<<EOT

	$(".award_type").change(function() {
		switch (parseInt($(this).val())) {
		case 0:
			autoopt_award_id.lookup = items;
			break;
		case 6:
			autoopt_award_id.lookup = ghosts;
			break;
		case 7:
			autoopt_award_id.lookup = sword_souls;
			break;
		 default:
			autoopt_award_id.lookup = nothing;
			break;
		}
		console.log(autoopt_award_id.lookup);
		$(this).parent().parent().find(".display_award_id").autocomplete(autoopt_award_id);
	})

	//$(".real_value").focus(function() {
	//	var award_id = parseInt($(this).val());
	//	var award_name = autoopt_award_id.lookup[award_id]; 
	//	var award_type = parseInt($(this).parent().parent().find(".award_type").val());
	//	console.log(award_id);
	//	console.log(award_name);
	//	console.log(award_type);
	//})
EOT;

	return $html;
}

?>
