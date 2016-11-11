<?php
//include "item_links.php";
$extend_columns = array(
	'item_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
	'goods_type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
);


function sql_where($params) {
	return "where grid_id={$params['m']}";
}

function sql_insert($params) {
	return "grid_id={$params['m']}";
}

function location($params) {
	global $db;
	$sql = "select * from trader_grid where id={$params['m']}";
	$grid = db_query($db, $sql)[0];
	$sql = "select * from trader where id={$grid['trader_id']}";
	$trader = db_query($db, $sql)[0];

	$html = html_get_link("随机商人", "?p=trader", false) . "->";
	$html .= html_get_link($trader['name'], "?p=trader_grid&m={$trader['id']}", false) . "->";
	$html .= "格子物品配置";
	return $html;
}

// 类型
$trader_goods_type = array(
	0 => '物品',
	1 => '剑心',
	2 => '魂侍',
	3 => '灵宠',
	4 => '爱心',
	5 => '铜币',
	6 => '元宝',
	7 => '体力',
	8 => '等级装备',
);



$all_items = get_item();
$all_sword_souls = get_all_sword_soul();
$all_ghosts = get_all_ghost();
$all_pets = get_all_pet();

function render_goods_type($column_name, $column_val, $row, $data) {
	global $trader_goods_type;
	return $trader_goods_type[$column_val];
}

function editor_goods_type($column_name, $column_val, $row, $data) {
	global $trader_goods_type;

	$data = $trader_goods_type;
	$field_name = 'goods_type';

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
	$tmp = db_query($db, "select `id`, `name` from `item` where `type_id`!='10' and `type_id`!='11'");
	foreach ($tmp as $value)
	{
		$all_items[$value['id']] = $value['name'];
	}

	return $all_items;
}



function editor_item_id($column_name, $column_val, $row, $data){
	$item_name = '';
	$item_id = $row['item_id'];

	if ($row['goods_type'] == 1) {
		global $all_sword_souls;
		if (isset($all_sword_souls[$item_id])) {
			$item_name = $all_sword_souls[$item_id];
		}
	} else if ($row['goods_type'] == 2) {
		global $all_ghosts;
		if (isset($all_ghosts[$item_id])) {
			$item_name = $all_ghosts[$item_id];
		}
	} else if ($row['goods_type'] == 3) {
		global $all_pets;
		if (isset($all_pets[$item_id])) {
			$item_name = $all_pets[$item_id];
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
	if ($row['goods_type'] == 1) {
		global $all_sword_souls;
		return normal_render($all_sword_souls, $row['item_id']);
	} else if ($row['goods_type'] == 2) {
		global $all_ghosts;
		return normal_render($all_ghosts, $row['item_id']);
	} else if ($row['goods_type'] == 3) {
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
	
	$html .= get_items_json($all_ghosts, 'ghosts');
	
	$html .= get_items_json($all_pets, 'pets');

	$nothing = array('无');
	$html .= get_items_json($nothing, 'nothing');



	$html .= <<<EOT

	$(".goods_type").change(function() {
		switch (parseInt($(this).val())) {
		case 0:
			autoopt_item_id.lookup = items;
			break;
		case 2:
			autoopt_item_id.lookup = ghosts;
			break;
		case 1:
			autoopt_item_id.lookup = sword_souls;
			break;
		case 3:
			autoopt_item_id.lookup = pets;
			break;
		 default:
			autoopt_item_id.lookup = nothing;
			break;
		}
		console.log(autoopt_item_id.lookup);
		$(this).parent().parent().find(".display_item_id").autocomplete(autoopt_item_id);
	})

EOT;

return $html;
}

?>
