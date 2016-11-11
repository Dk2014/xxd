<?php
include "item_links.php";

$extend_columns = array(
/*   '字段' => 配置 */
	'mode' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array('直接获得','随机数量','概率获得','随机物品'),
	),

	'type' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),

	'get_item_id' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),

	'item_id_set' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
	),
);

$all_type = array(
	2 => '物品',
	7 => '剑心',
	6 => '魂侍',
	3 => '主角经验',
	8 => '灵宠',
	4 => '爱心',
	0 => '铜币',
	1 => '元宝',
	5 => '体力',
	9 => '剑心碎片',
);

$all_items = get_all_item();
$all_sword_souls = get_all_sword_soul();
$all_ghosts = get_all_ghost();
$all_pets = get_all_pet();

function render_mode($column_name, $column_val, $row, $data) {
	return $data[$column_val];
}

function editor_mode($column_name, $column_val, $row, $data) {
	return html_get_select($column_name,$data,$column_val);
}

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


function render_get_item_id($column_name, $column_val, $row, $data) {
	switch ($row['type']) {
	case 7:
		global $all_sword_souls;
		return normal_render($all_sword_souls, $row['get_item_id']); 
	case 6:
		global $all_ghosts;
		return normal_render($all_ghosts, $row['get_item_id']); 
	case 8:
		global $all_pets;
		return normal_render($all_pets, $row['get_item_id']); 
	case 2:
		global $all_items;
		return normal_render($all_items, $row['get_item_id']);
	}
}

function editor_get_item_id($column_name, $column_val, $row, $data) {
	$item_name = '';
	$item_id = $row['get_item_id'];

	global $all_items;
	global $all_pets;
	global $all_sword_souls;
	global $all_ghosts;

	switch ($row['type']) {
	case 7:
		$item_name = $all_sword_souls[$item_id];
		break;
	case 6:
		$item_name = $all_ghosts[$item_id];
		break;
	case 8:
		$item_name = $all_pets[$item_id];
		break;
	case 2:
		$item_name = $all_items[$item_id];
		break;
	}

	return editor_single_item($item_name, $item_id, "get_item_id");
}

function items_render($item_ids_str) {
	global $all_items;

	$items = json_decode($item_ids_str, true);

	if (!is_array($items)){
		$items = array();
	}

	$arr = array();
	foreach ($items as $value) {
		array_push($arr, "{$all_items[$value['ItemId']]}*{$value['Num']}");
	}

	return implode('</br>', $arr);
}

function render_item_id_set($column_name, $column_val, $row, $data) {
	return items_render($row['item_id_set']);
}

function editor_item_id_set($column_name, $column_val, $row, $data) {
	global $all_items;
	$column_name = 'item_id_set';

	$item_ids = json_decode($row[$column_name], true);

	if (!is_array($item_ids)){
		$item_ids = array();
	}

	$display = '';
	foreach ($item_ids as $value) {
		$display .= "<rowspan><input class='display_{$column_name}' value='{$all_items[$value['ItemId']]}' /input>*<input class='display_num' value='{$value['Num']}' /input></br></rowspan>";
	}

	$btn = "<a class='add_{$column_name}' href='javascript:void(0)'>+</a>";
	$real_value = "<input type='hidden' class='real_value' name='{$column_name}[]' value='{$row[$column_name]}' />";

	return $display.$btn.$real_value;
}

function location($params){
	global $db;

	$box_name = db_query($db, "select `name` from `item` where `id` = {$params['m']}");

	$html = '<a href="?p=item_box">'.$box_name[0]['name'].'</a> -> ';

	return $html;
}

function sql_where($params) {
	if (isset($params['m']))
		$where = " where `item_id`={$params['m']}";

	return $where;
}

function sql_insert($params) {
	return "`item_id` = {$params['m']}";
}

function js_function($params) {
	global $all_items;
	global $all_sword_souls;
	global $all_ghosts;
	global $all_pets;

	$html = get_items_json($all_items, 'items');
	$column_name = 'item_id_set';

	$html .= <<<EOT

	autoopt_{$column_name} = {
		lookup: items, 
		minChars: 0,
		onSelect: function(s){ 
			item_obj = $(this).parent().siblings('input[class="real_value"]');
			v = item_obj.val();
			if (v == ''){
				v = '[]';
			}
			items = JSON.parse(v);

			// 确定该位置序号
			inputs = $(this).parent().parent().find('.display_{$column_name}');
			index = null;

			for (var i = 0; i < inputs.length; i++) {
				if (inputs[i] == this){
					index = i ;
					break;
				}
			};

			if (index == null) {
				index = inputs.length;
			}

			id = parseInt(s.id); // 物品ID
			num = parseInt($(this).next().val()); // 物品数量

			// 删除
			if (id == 0) {
				$(this).parent().remove();
				items.splice(index, 1);
			} else {
				items[index] = {"ItemId":id,"Num":num};
			}

			item_obj.val(JSON.stringify(items));
		}
	};
	$(".display_{$column_name}").autocomplete(autoopt_{$column_name});

	function num_change() 
	{
		item_obj = $(this).parent().siblings('input[class="real_value"]');
		v = item_obj.val();
		if (v == ''){
			v = '[]';
		}
		items = JSON.parse(v);

		// 确定该位置序号
		inputs = $(this).parent().parent().find('.display_num');
		index = 0;

		for (var i = 0; i < inputs.length; i++) {
			if (inputs[i] == this){
				index = i ;
				break;
			}
		};

		num = parseInt($(this).val()); // 物品数量

		// 删除
		if (num == 0) {
			$(this).parent().remove();
			items.splice(index, 1);
		} else {
			if (items[index] == undefined) {
				alert("先选择物品！");
				return;
			}
			items[index]["Num"] = num;
		}

		item_obj.val(JSON.stringify(items));
	}
	$('.display_num').change(num_change);

	$(".add_{$column_name}").click(function(){
		$(this).before("<rowspan><input class='display_{$column_name}' value=''>*<input class='display_num' value='1'></br></rowspan>");

		$(this).prev().find('.display_num').change(num_change)

		$(this).prev().find('.display_{$column_name}').autocomplete(autoopt_{$column_name});
	})	

EOT;

	//$html .= choose_single_item($all_items, 'get_item_id');


	$html .= <<<EOT
	autoopt_get_item_id = {
		lookup: items, 
		minChars: 0,
		onSelect: function(s) {
			$(this).siblings('input[class="real_value"]').val(s.id);
		}
	};

	$(".display_get_item_id").autocomplete(autoopt_get_item_id);
EOT;

	$html .= get_items_json($all_sword_souls, 'sword_souls');
	$html .= get_items_json($all_ghosts, 'ghosts');
	$html .= get_items_json($all_pets, 'pets');
	$html .= get_items_json(array('无'), 'nothing');


	$html .= <<<EOT
	$(".type").change(function() {
		switch (parseInt($(this).val())) {
		case 2:
			autoopt_get_item_id.lookup = items;
			break;
		case 7:
			autoopt_get_item_id.lookup = sword_souls;
			break;
		case 6:
			autoopt_get_item_id.lookup = ghosts;
			break;
		case 8:
			autoopt_get_item_id.lookup = pets;
			break;
		default:
			autoopt_get_item_id.lookup =nothing;
		}
		$(this).parent().parent().find(".display_get_item_id").autocomplete(autoopt_get_item_id);
	})

EOT;

	return $html;
}
?>
